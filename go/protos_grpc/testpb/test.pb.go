// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testpb/test.proto

package testpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "grpc-n/studentpb"
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

type Test struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Test) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetTestRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTestRequest) Reset()         { *m = GetTestRequest{} }
func (m *GetTestRequest) String() string { return proto.CompactTextString(m) }
func (*GetTestRequest) ProtoMessage()    {}
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{1}
}

func (m *GetTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTestRequest.Unmarshal(m, b)
}
func (m *GetTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTestRequest.Marshal(b, m, deterministic)
}
func (m *GetTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTestRequest.Merge(m, src)
}
func (m *GetTestRequest) XXX_Size() int {
	return xxx_messageInfo_GetTestRequest.Size(m)
}
func (m *GetTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTestRequest proto.InternalMessageInfo

func (m *GetTestRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SetTestResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTestResponse) Reset()         { *m = SetTestResponse{} }
func (m *SetTestResponse) String() string { return proto.CompactTextString(m) }
func (*SetTestResponse) ProtoMessage()    {}
func (*SetTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{2}
}

func (m *SetTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTestResponse.Unmarshal(m, b)
}
func (m *SetTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTestResponse.Marshal(b, m, deterministic)
}
func (m *SetTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTestResponse.Merge(m, src)
}
func (m *SetTestResponse) XXX_Size() int {
	return xxx_messageInfo_SetTestResponse.Size(m)
}
func (m *SetTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTestResponse proto.InternalMessageInfo

func (m *SetTestResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetTestResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Question struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question             string   `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	TestId               string   `protobuf:"bytes,4,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{3}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Question) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *Question) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Question) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

type SetQuestionResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetQuestionResponse) Reset()         { *m = SetQuestionResponse{} }
func (m *SetQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*SetQuestionResponse) ProtoMessage()    {}
func (*SetQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{4}
}

func (m *SetQuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetQuestionResponse.Unmarshal(m, b)
}
func (m *SetQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetQuestionResponse.Marshal(b, m, deterministic)
}
func (m *SetQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetQuestionResponse.Merge(m, src)
}
func (m *SetQuestionResponse) XXX_Size() int {
	return xxx_messageInfo_SetQuestionResponse.Size(m)
}
func (m *SetQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetQuestionResponse proto.InternalMessageInfo

func (m *SetQuestionResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type EnrollmentRequest struct {
	StudentId            string   `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TestId               string   `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnrollmentRequest) Reset()         { *m = EnrollmentRequest{} }
func (m *EnrollmentRequest) String() string { return proto.CompactTextString(m) }
func (*EnrollmentRequest) ProtoMessage()    {}
func (*EnrollmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{5}
}

func (m *EnrollmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnrollmentRequest.Unmarshal(m, b)
}
func (m *EnrollmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnrollmentRequest.Marshal(b, m, deterministic)
}
func (m *EnrollmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnrollmentRequest.Merge(m, src)
}
func (m *EnrollmentRequest) XXX_Size() int {
	return xxx_messageInfo_EnrollmentRequest.Size(m)
}
func (m *EnrollmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnrollmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnrollmentRequest proto.InternalMessageInfo

func (m *EnrollmentRequest) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

func (m *EnrollmentRequest) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

type GetStudentsPerTestRequest struct {
	TestId               string   `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentsPerTestRequest) Reset()         { *m = GetStudentsPerTestRequest{} }
func (m *GetStudentsPerTestRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentsPerTestRequest) ProtoMessage()    {}
func (*GetStudentsPerTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{6}
}

func (m *GetStudentsPerTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentsPerTestRequest.Unmarshal(m, b)
}
func (m *GetStudentsPerTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentsPerTestRequest.Marshal(b, m, deterministic)
}
func (m *GetStudentsPerTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentsPerTestRequest.Merge(m, src)
}
func (m *GetStudentsPerTestRequest) XXX_Size() int {
	return xxx_messageInfo_GetStudentsPerTestRequest.Size(m)
}
func (m *GetStudentsPerTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentsPerTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentsPerTestRequest proto.InternalMessageInfo

func (m *GetStudentsPerTestRequest) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

type TakeTestRequest struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeTestRequest) Reset()         { *m = TakeTestRequest{} }
func (m *TakeTestRequest) String() string { return proto.CompactTextString(m) }
func (*TakeTestRequest) ProtoMessage()    {}
func (*TakeTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{7}
}

func (m *TakeTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeTestRequest.Unmarshal(m, b)
}
func (m *TakeTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeTestRequest.Marshal(b, m, deterministic)
}
func (m *TakeTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeTestRequest.Merge(m, src)
}
func (m *TakeTestRequest) XXX_Size() int {
	return xxx_messageInfo_TakeTestRequest.Size(m)
}
func (m *TakeTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TakeTestRequest proto.InternalMessageInfo

func (m *TakeTestRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "test.Test")
	proto.RegisterType((*GetTestRequest)(nil), "test.GetTestRequest")
	proto.RegisterType((*SetTestResponse)(nil), "test.SetTestResponse")
	proto.RegisterType((*Question)(nil), "test.Question")
	proto.RegisterType((*SetQuestionResponse)(nil), "test.SetQuestionResponse")
	proto.RegisterType((*EnrollmentRequest)(nil), "test.EnrollmentRequest")
	proto.RegisterType((*GetStudentsPerTestRequest)(nil), "test.GetStudentsPerTestRequest")
	proto.RegisterType((*TakeTestRequest)(nil), "test.TakeTestRequest")
}

func init() { proto.RegisterFile("testpb/test.proto", fileDescriptor_41c67e33ca9d1f26) }

var fileDescriptor_41c67e33ca9d1f26 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6f, 0xda, 0x30,
	0x10, 0xc7, 0xe5, 0x2c, 0x82, 0x70, 0xdb, 0xc2, 0xb8, 0xfd, 0x20, 0x44, 0x9a, 0x86, 0x22, 0x4d,
	0x62, 0xbf, 0x00, 0x6d, 0xe3, 0x69, 0x6f, 0x93, 0x5a, 0x84, 0xfa, 0xd2, 0x12, 0x9e, 0xfa, 0x82,
	0x80, 0x58, 0x28, 0x02, 0x9c, 0x10, 0x9b, 0xf6, 0xcf, 0xee, 0xbf, 0x50, 0x39, 0x76, 0x52, 0x13,
	0x50, 0xd5, 0xa7, 0x9c, 0xcf, 0xdf, 0xaf, 0x2f, 0xf7, 0x39, 0x1b, 0x5a, 0x82, 0x72, 0x91, 0x2e,
	0x07, 0xf2, 0xd3, 0x4f, 0xb3, 0x44, 0x24, 0x68, 0xcb, 0xd8, 0x6f, 0x73, 0x71, 0x88, 0x28, 0x93,
	0x7b, 0x3a, 0x52, 0xdb, 0xc1, 0x77, 0xb0, 0x67, 0x94, 0x0b, 0x74, 0xc1, 0x8a, 0x23, 0x8f, 0x74,
	0x49, 0xaf, 0x31, 0xb5, 0xe2, 0x08, 0x11, 0x6c, 0xb6, 0xd8, 0x51, 0xcf, 0xca, 0x33, 0x79, 0x1c,
	0x74, 0xc1, 0x1d, 0x53, 0x21, 0xe5, 0x53, 0xba, 0x3f, 0x9c, 0x71, 0x05, 0x23, 0x68, 0x86, 0x85,
	0x82, 0xa7, 0x09, 0xe3, 0xf4, 0x45, 0x07, 0xaf, 0xc1, 0xb9, 0x91, 0xe7, 0xc5, 0x09, 0x3b, 0xd1,
	0x7f, 0x82, 0xda, 0x82, 0xf1, 0x7b, 0x9a, 0x69, 0x87, 0x5e, 0xa1, 0x0f, 0xce, 0x5e, 0x7b, 0xbc,
	0x57, 0xf9, 0x4e, 0xb9, 0xc6, 0x36, 0xd4, 0x65, 0xd7, 0xf3, 0x38, 0xf2, 0x6c, 0x65, 0x92, 0xcb,
	0x49, 0x14, 0x7c, 0x85, 0xf7, 0x21, 0x15, 0x45, 0x2d, 0xf3, 0x1f, 0x93, 0x4d, 0x5e, 0xd3, 0x99,
	0x5a, 0xc9, 0x26, 0xb8, 0x82, 0xd6, 0x05, 0xcb, 0x92, 0xed, 0x76, 0x47, 0x59, 0xd9, 0xeb, 0x67,
	0x00, 0x8d, 0x6e, 0x5e, 0xfe, 0x60, 0x43, 0x67, 0x26, 0x91, 0x59, 0xd3, 0x3a, 0xaa, 0xf9, 0x17,
	0x3a, 0x63, 0x2a, 0x42, 0x25, 0xe4, 0xd7, 0x34, 0x33, 0x01, 0x1a, 0x2e, 0x72, 0xe4, 0xfa, 0x06,
	0xcd, 0xd9, 0x62, 0x43, 0x4d, 0xed, 0x13, 0x09, 0x62, 0x92, 0xf8, 0xfd, 0x60, 0xc1, 0x6b, 0xa9,
	0x0b, 0x69, 0x76, 0x17, 0xaf, 0x28, 0xfe, 0x80, 0xba, 0x1e, 0x13, 0x7e, 0xe8, 0xe7, 0x37, 0xe1,
	0x78, 0x6a, 0x3e, 0xa8, 0x6c, 0xae, 0xf8, 0x09, 0x75, 0x3d, 0x31, 0x34, 0xd2, 0xfe, 0x47, 0x15,
	0x57, 0x87, 0xf9, 0x0f, 0xde, 0x18, 0xfc, 0x38, 0xba, 0x4a, 0x56, 0x24, 0xfc, 0x4e, 0x69, 0xab,
	0x32, 0xee, 0x11, 0xbc, 0x04, 0x57, 0x51, 0x2d, 0x58, 0x60, 0x5b, 0xc9, 0x4f, 0x58, 0x3f, 0x7f,
	0xce, 0x04, 0xf0, 0x14, 0x28, 0x7e, 0x29, 0x5b, 0x3d, 0x8f, 0xda, 0x7f, 0xd7, 0x2f, 0x6e, 0xbe,
	0x16, 0x0c, 0x09, 0x8e, 0xc0, 0x29, 0x28, 0xa3, 0x6e, 0xb9, 0x42, 0xdd, 0xaf, 0xb4, 0xd8, 0x23,
	0x43, 0xf2, 0xbf, 0x79, 0xfb, 0x76, 0x9d, 0xa5, 0xab, 0x5f, 0x6c, 0xa0, 0xde, 0xdb, 0xb2, 0x96,
	0x3f, 0xa6, 0x3f, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x7e, 0x40, 0xc3, 0x80, 0x03, 0x00,
	0x00,
}
