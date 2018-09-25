// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mutant/mutant.proto

package go_micro_srv_mutant

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6eac06ca9f6551, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Dna                  []string `protobuf:"bytes,1,rep,name=dna,proto3" json:"dna,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6eac06ca9f6551, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetDna() []string {
	if m != nil {
		return m.Dna
	}
	return nil
}

type Response struct {
	IsMutant             bool     `protobuf:"varint,1,opt,name=isMutant,proto3" json:"isMutant,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6eac06ca9f6551, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetIsMutant() bool {
	if m != nil {
		return m.IsMutant
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.mutant.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.mutant.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.mutant.Response")
}

func init() { proto.RegisterFile("proto/mutant/mutant.proto", fileDescriptor_ad6eac06ca9f6551) }

var fileDescriptor_ad6eac06ca9f6551 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x2d, 0x2d, 0x49, 0xcc, 0x2b, 0x81, 0x52, 0x7a, 0x60, 0x31, 0x21, 0xe1, 0xf4,
	0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0x3d, 0x88, 0x94, 0x92, 0x34,
	0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x71, 0x62, 0xa5,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x09, 0x92, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x01, 0x49, 0xa6, 0xe4, 0x25, 0x4a, 0x30, 0x2a, 0x30, 0x83, 0x24, 0x53, 0xf2, 0x12, 0x95,
	0x2c, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0x32,
	0x8b, 0x7d, 0xc1, 0x26, 0x82, 0xf5, 0x73, 0x04, 0xc1, 0xf9, 0x20, 0x9d, 0xb9, 0xc5, 0xe9, 0x12,
	0x4c, 0x10, 0x63, 0x73, 0x8b, 0xd3, 0x8d, 0x02, 0xb8, 0x98, 0x5d, 0xf2, 0x12, 0x85, 0x3c, 0xb9,
	0x38, 0x3c, 0x61, 0x8a, 0x64, 0xf4, 0xb0, 0x38, 0x4e, 0x0f, 0x6a, 0xb9, 0x94, 0x2c, 0x0e, 0x59,
	0x88, 0xed, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x1f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b,
	0x37, 0xcd, 0xcf, 0xfe, 0x00, 0x00, 0x00,
}
