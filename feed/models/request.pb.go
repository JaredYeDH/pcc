// Code generated by protoc-gen-go.
// source: request.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestMethod int32

const (
	RequestMethod__      RequestMethod = 0
	RequestMethod_Add    RequestMethod = 1
	RequestMethod_Delete RequestMethod = 2
	RequestMethod_Update RequestMethod = 3
)

var RequestMethod_name = map[int32]string{
	0: "_",
	1: "Add",
	2: "Delete",
	3: "Update",
}
var RequestMethod_value = map[string]int32{
	"_":      0,
	"Add":    1,
	"Delete": 2,
	"Update": 3,
}

func (x RequestMethod) String() string {
	return proto.EnumName(RequestMethod_name, int32(x))
}
func (RequestMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type Request struct {
	Id     int64         `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Method RequestMethod `protobuf:"varint,2,opt,name=method,enum=models.RequestMethod" json:"method,omitempty"`
	Data   []byte        `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetMethod() RequestMethod {
	if m != nil {
		return m.Method
	}
	return RequestMethod__
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "models.Request")
	proto.RegisterEnum("models.RequestMethod", RequestMethod_name, RequestMethod_value)
}

func init() { proto.RegisterFile("request.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8e, 0x41, 0x0b, 0x82, 0x40,
	0x10, 0x46, 0x9b, 0xb5, 0x56, 0x18, 0x52, 0x96, 0x81, 0xc0, 0xa3, 0x74, 0x92, 0x20, 0x0f, 0x75,
	0xe8, 0x1c, 0x74, 0xed, 0xb2, 0xd0, 0x2d, 0x08, 0x63, 0x06, 0x12, 0x94, 0x35, 0x9d, 0xfe, 0x7f,
	0xa4, 0x5e, 0xba, 0x7d, 0xf0, 0xde, 0x83, 0x0f, 0x93, 0x5e, 0xde, 0x1f, 0x19, 0xb4, 0xec, 0xfa,
	0xa0, 0x81, 0x6c, 0x1b, 0x58, 0x9a, 0x61, 0x7b, 0xc7, 0xd8, 0x4f, 0x80, 0x52, 0x34, 0x35, 0x67,
	0x90, 0x43, 0x11, 0x79, 0x53, 0x33, 0xed, 0xd1, 0xb6, 0xa2, 0xaf, 0xc0, 0x99, 0xc9, 0xa1, 0x48,
	0x0f, 0x9b, 0x72, 0x6a, 0xca, 0x39, 0xb8, 0x8e, 0xd0, 0xcf, 0x12, 0x11, 0x2e, 0xb9, 0xd2, 0x2a,
	0xb3, 0x39, 0x14, 0x6b, 0x3f, 0xee, 0xdd, 0x09, 0x93, 0x3f, 0x99, 0x56, 0x08, 0x0f, 0xb7, 0xa0,
	0x18, 0xa3, 0x33, 0xb3, 0x03, 0x42, 0xb4, 0x17, 0x69, 0x44, 0xc5, 0x99, 0xdf, 0xbe, 0x75, 0x5c,
	0xa9, 0xb8, 0xe8, 0x69, 0xc7, 0x97, 0xc7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xb0, 0x08,
	0x87, 0xb6, 0x00, 0x00, 0x00,
}