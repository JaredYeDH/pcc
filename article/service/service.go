package service

import (
	"fmt"
	"time"

	"github.com/chideat/glog"
	"github.com/chideat/pcc/article/models"
	. "github.com/chideat/pcc/article/modules/config"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
)

var (
	consumer *nsq.Consumer
)

func init() {
	var err error

	// config
	config := nsq.NewConfig()
	config.MaxAttempts = 2
	consumer, err = nsq.NewConsumer("pcc.action_like", "article", config)
	if err != nil {
		panic(err)
	}

	consumer.AddHandler(&ActionHandler{errorMessages: map[int64]error{}})
	consumer.SetLogger(nil, nsq.LogLevelError)
	err = consumer.ConnectToNSQLookupd(Conf.Queue.LookupdAddress)
	if err != nil {
		panic(err)
	}
}

type ActionHandler struct {
	errorMessages map[int64]error
}

func (handler *ActionHandler) HandleMessage(msg *nsq.Message) error {
	req := models.Request{}
	err := proto.Unmarshal(msg.Body, &req)
	if err != nil {
		glog.Error(err)
		return nil
	}
	likeAction := models.LikeAction{}
	err = proto.Unmarshal(req.Data, &likeAction)
	if err != nil {
		glog.Error(err)
		return nil
	}
	if likeAction.Id == 0 {
		glog.Error("invalid like action with empty id")
		return nil
	}

	article, err := models.GetarticleById(likeAction.Target)
	if err != nil {
		glog.Error(err)
		return nil
	}
	if article == nil {
		glog.Error("article with id %d not exists", likeAction.Target)
		return nil
	}

	switch req.Method {
	case models.RequestMethod_Add:
		err = article.Like()
	case models.RequestMethod_Update:
		if likeAction.Deleted {
			err = article.Like()
		}
	case models.RequestMethod_Delete:
		err = article.CancelLike()
	default:
		err = fmt.Errorf("unknow type %s", req.Method)
	}

	if err != nil {
		glog.Error(err)
		handler.errorMessages[likeAction.Id] = fmt.Errorf("%s %s", req.Method, err.Error())
		msg.Requeue(1 * time.Second)
	}
	return nil
}

func (handler *ActionHandler) LogFailedMessage(msg *nsq.Message) {
	req := models.Request{}
	err := proto.Unmarshal(msg.Body, &req)
	if err != nil {
		glog.Error(err)
		return
	}
	likeAction := models.LikeAction{}
	err = proto.Unmarshal(req.Data, &likeAction)
	if err != nil {
		glog.Error(err)
		return
	}

	err, ok := handler.errorMessages[likeAction.Id]
	if ok {
		glog.Error("process like action %d failed with error %s", likeAction.Id, err.Error())
	} else {
		glog.Error("process like action %d failed with unknown error", likeAction.Id)
	}
}