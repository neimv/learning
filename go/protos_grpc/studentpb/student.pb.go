// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/student.proto

package studentpb

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Student struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_f512b57901e127c8, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type GetStudentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentRequest) Reset()         { *m = GetStudentRequest{} }
func (m *GetStudentRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentRequest) ProtoMessage()    {}
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f512b57901e127c8, []int{1}
}

func (m *GetStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentRequest.Unmarshal(m, b)
}
func (m *GetStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentRequest.Marshal(b, m, deterministic)
}
func (m *GetStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentRequest.Merge(m, src)
}
func (m *GetStudentRequest) XXX_Size() int {
	return xxx_messageInfo_GetStudentRequest.Size(m)
}
func (m *GetStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentRequest proto.InternalMessageInfo

func (m *GetStudentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SetStudentResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStudentResponse) Reset()         { *m = SetStudentResponse{} }
func (m *SetStudentResponse) String() string { return proto.CompactTextString(m) }
func (*SetStudentResponse) ProtoMessage()    {}
func (*SetStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f512b57901e127c8, []int{2}
}

func (m *SetStudentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStudentResponse.Unmarshal(m, b)
}
func (m *SetStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStudentResponse.Marshal(b, m, deterministic)
}
func (m *SetStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStudentResponse.Merge(m, src)
}
func (m *SetStudentResponse) XXX_Size() int {
	return xxx_messageInfo_SetStudentResponse.Size(m)
}
func (m *SetStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetStudentResponse proto.InternalMessageInfo

func (m *SetStudentResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Student)(nil), "student.Student")
	proto.RegisterType((*GetStudentRequest)(nil), "student.GetStudentRequest")
	proto.RegisterType((*SetStudentResponse)(nil), "student.SetStudentResponse")
}

func init() { proto.RegisterFile("proto/student.proto", fileDescriptor_f512b57901e127c8) }

var fileDescriptor_f512b57901e127c8 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2e, 0x29, 0x4d, 0x49, 0xcd, 0x2b, 0xd1, 0x03, 0xf3, 0x84, 0xd8, 0xa1, 0x5c,
	0x25, 0x7b, 0x2e, 0xf6, 0x60, 0x08, 0x53, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82,
	0x09, 0x2c, 0x02, 0x66, 0x0b, 0x09, 0x70, 0x31, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2b, 0x30, 0x6a,
	0xb0, 0x06, 0x81, 0x98, 0x4a, 0xca, 0x5c, 0x82, 0xee, 0xa9, 0x25, 0x50, 0x33, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x31, 0x8c, 0x52, 0x52, 0xe1, 0x12, 0x0a, 0x46, 0x52, 0x54, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x8a, 0xae, 0xca, 0xa8, 0x93, 0x91, 0x8b, 0x0f, 0xaa, 0x26, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0xc8, 0x8a, 0x8b, 0x0b, 0x61, 0xba, 0x90, 0x94, 0x1e, 0xcc, 0x17, 0x18, 0x56,
	0x4a, 0x09, 0xc0, 0xe5, 0x60, 0xaa, 0xad, 0xb9, 0xb8, 0x10, 0x96, 0x0a, 0x61, 0xc8, 0x4b, 0x49,
	0x23, 0x44, 0x30, 0xdc, 0xe6, 0x24, 0x14, 0x25, 0x90, 0x5e, 0x54, 0x90, 0xac, 0x9b, 0x67, 0x0d,
	0x55, 0x54, 0x90, 0x94, 0xc4, 0x06, 0x0e, 0x3b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0xfd, 0xa7, 0x61, 0x52, 0x01, 0x00, 0x00,
}
