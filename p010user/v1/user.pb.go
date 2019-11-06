// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p010user/v1/user.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserData struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	CreatedAt            int64    `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Oauth                string   `protobuf:"bytes,6,opt,name=Oauth,proto3" json:"Oauth,omitempty"`
	NativeLang           string   `protobuf:"bytes,7,opt,name=NativeLang,proto3" json:"NativeLang,omitempty"`
	IsValidated          bool     `protobuf:"varint,8,opt,name=IsValidated,proto3" json:"IsValidated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{1}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserData) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *UserData) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *UserData) GetOauth() string {
	if m != nil {
		return m.Oauth
	}
	return ""
}

func (m *UserData) GetNativeLang() string {
	if m != nil {
		return m.NativeLang
	}
	return ""
}

func (m *UserData) GetIsValidated() bool {
	if m != nil {
		return m.IsValidated
	}
	return false
}

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterReply struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{3}
}

func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (m *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(m, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

func (m *RegisterReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

///////// Register Validation //////////
type RegisterValidationRequest struct {
	Otp                  string   `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterValidationRequest) Reset()         { *m = RegisterValidationRequest{} }
func (m *RegisterValidationRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterValidationRequest) ProtoMessage()    {}
func (*RegisterValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{4}
}

func (m *RegisterValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterValidationRequest.Unmarshal(m, b)
}
func (m *RegisterValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterValidationRequest.Marshal(b, m, deterministic)
}
func (m *RegisterValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterValidationRequest.Merge(m, src)
}
func (m *RegisterValidationRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterValidationRequest.Size(m)
}
func (m *RegisterValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterValidationRequest proto.InternalMessageInfo

func (m *RegisterValidationRequest) GetOtp() string {
	if m != nil {
		return m.Otp
	}
	return ""
}

func (m *RegisterValidationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RegisterValidationReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterValidationReply) Reset()         { *m = RegisterValidationReply{} }
func (m *RegisterValidationReply) String() string { return proto.CompactTextString(m) }
func (*RegisterValidationReply) ProtoMessage()    {}
func (*RegisterValidationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{5}
}

func (m *RegisterValidationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterValidationReply.Unmarshal(m, b)
}
func (m *RegisterValidationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterValidationReply.Marshal(b, m, deterministic)
}
func (m *RegisterValidationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterValidationReply.Merge(m, src)
}
func (m *RegisterValidationReply) XXX_Size() int {
	return xxx_messageInfo_RegisterValidationReply.Size(m)
}
func (m *RegisterValidationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterValidationReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterValidationReply proto.InternalMessageInfo

func (m *RegisterValidationReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterValidationReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{6}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{7}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

///////// Auth Validation //////////
type AuthValidationRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	AuthType             string   `protobuf:"bytes,2,opt,name=AuthType,proto3" json:"AuthType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthValidationRequest) Reset()         { *m = AuthValidationRequest{} }
func (m *AuthValidationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthValidationRequest) ProtoMessage()    {}
func (*AuthValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{8}
}

func (m *AuthValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthValidationRequest.Unmarshal(m, b)
}
func (m *AuthValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthValidationRequest.Marshal(b, m, deterministic)
}
func (m *AuthValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthValidationRequest.Merge(m, src)
}
func (m *AuthValidationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthValidationRequest.Size(m)
}
func (m *AuthValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthValidationRequest proto.InternalMessageInfo

func (m *AuthValidationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthValidationRequest) GetAuthType() string {
	if m != nil {
		return m.AuthType
	}
	return ""
}

type AuthValidationReply struct {
	User                 *UserData `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Error                *Error    `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AuthValidationReply) Reset()         { *m = AuthValidationReply{} }
func (m *AuthValidationReply) String() string { return proto.CompactTextString(m) }
func (*AuthValidationReply) ProtoMessage()    {}
func (*AuthValidationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{9}
}

func (m *AuthValidationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthValidationReply.Unmarshal(m, b)
}
func (m *AuthValidationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthValidationReply.Marshal(b, m, deterministic)
}
func (m *AuthValidationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthValidationReply.Merge(m, src)
}
func (m *AuthValidationReply) XXX_Size() int {
	return xxx_messageInfo_AuthValidationReply.Size(m)
}
func (m *AuthValidationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthValidationReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthValidationReply proto.InternalMessageInfo

func (m *AuthValidationReply) GetUser() *UserData {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AuthValidationReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

///////// UpdateUser //////////
type SetNativeLangRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	NativeLang           string   `protobuf:"bytes,2,opt,name=NativeLang,proto3" json:"NativeLang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetNativeLangRequest) Reset()         { *m = SetNativeLangRequest{} }
func (m *SetNativeLangRequest) String() string { return proto.CompactTextString(m) }
func (*SetNativeLangRequest) ProtoMessage()    {}
func (*SetNativeLangRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{10}
}

func (m *SetNativeLangRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetNativeLangRequest.Unmarshal(m, b)
}
func (m *SetNativeLangRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetNativeLangRequest.Marshal(b, m, deterministic)
}
func (m *SetNativeLangRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNativeLangRequest.Merge(m, src)
}
func (m *SetNativeLangRequest) XXX_Size() int {
	return xxx_messageInfo_SetNativeLangRequest.Size(m)
}
func (m *SetNativeLangRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNativeLangRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetNativeLangRequest proto.InternalMessageInfo

func (m *SetNativeLangRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SetNativeLangRequest) GetNativeLang() string {
	if m != nil {
		return m.NativeLang
	}
	return ""
}

type SetNativeLangResponse struct {
	Error                *Error   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetNativeLangResponse) Reset()         { *m = SetNativeLangResponse{} }
func (m *SetNativeLangResponse) String() string { return proto.CompactTextString(m) }
func (*SetNativeLangResponse) ProtoMessage()    {}
func (*SetNativeLangResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30698c624460701e, []int{11}
}

func (m *SetNativeLangResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetNativeLangResponse.Unmarshal(m, b)
}
func (m *SetNativeLangResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetNativeLangResponse.Marshal(b, m, deterministic)
}
func (m *SetNativeLangResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNativeLangResponse.Merge(m, src)
}
func (m *SetNativeLangResponse) XXX_Size() int {
	return xxx_messageInfo_SetNativeLangResponse.Size(m)
}
func (m *SetNativeLangResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNativeLangResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetNativeLangResponse proto.InternalMessageInfo

func (m *SetNativeLangResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "user.Error")
	proto.RegisterType((*UserData)(nil), "user.UserData")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "user.RegisterReply")
	proto.RegisterType((*RegisterValidationRequest)(nil), "user.RegisterValidationRequest")
	proto.RegisterType((*RegisterValidationReply)(nil), "user.RegisterValidationReply")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "user.LoginReply")
	proto.RegisterType((*AuthValidationRequest)(nil), "user.AuthValidationRequest")
	proto.RegisterType((*AuthValidationReply)(nil), "user.AuthValidationReply")
	proto.RegisterType((*SetNativeLangRequest)(nil), "user.SetNativeLangRequest")
	proto.RegisterType((*SetNativeLangResponse)(nil), "user.SetNativeLangResponse")
}

func init() { proto.RegisterFile("p010user/v1/user.proto", fileDescriptor_30698c624460701e) }

var fileDescriptor_30698c624460701e = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6f, 0xd3, 0x3e,
	0x14, 0x5d, 0xd2, 0xa6, 0xeb, 0x6e, 0xf7, 0xdb, 0xaf, 0xf2, 0xda, 0x91, 0x65, 0x7c, 0x94, 0x08,
	0xa1, 0xbe, 0xd0, 0x6c, 0x05, 0x24, 0xc4, 0xdb, 0x80, 0x22, 0x22, 0x4d, 0x43, 0x84, 0x8e, 0x07,
	0xc4, 0x8b, 0xdb, 0x58, 0x6d, 0x44, 0x9b, 0x04, 0xdb, 0xe9, 0xd4, 0xff, 0x96, 0x07, 0xfe, 0x10,
	0xe4, 0xd8, 0xee, 0x47, 0x96, 0x49, 0x95, 0x78, 0x8a, 0xcf, 0x3d, 0xbe, 0x47, 0xe7, 0x7e, 0x38,
	0x70, 0x92, 0x9e, 0x5f, 0x9c, 0x67, 0x8c, 0x50, 0x6f, 0x71, 0xe1, 0x89, 0x6f, 0x2f, 0xa5, 0x09,
	0x4f, 0x50, 0x55, 0x9c, 0xdd, 0xd7, 0x60, 0x0d, 0x28, 0x4d, 0x28, 0x42, 0x50, 0x1d, 0x27, 0x21,
	0xb1, 0x8d, 0x8e, 0xd1, 0xb5, 0x82, 0xfc, 0x8c, 0x6c, 0xd8, 0x9f, 0x13, 0xc6, 0xf0, 0x84, 0xd8,
	0x66, 0xc7, 0xe8, 0x1e, 0x04, 0x1a, 0xba, 0xbf, 0x0d, 0xa8, 0xdf, 0x30, 0x42, 0x3f, 0x60, 0x8e,
	0xd1, 0x11, 0x98, 0x7e, 0xa8, 0x12, 0x4d, 0x3f, 0x44, 0x8e, 0xe4, 0x62, 0x3c, 0xd7, 0x79, 0x2b,
	0x8c, 0x5a, 0x60, 0x0d, 0xe6, 0x38, 0x9a, 0xd9, 0x95, 0x9c, 0x90, 0x00, 0x3d, 0x84, 0x83, 0xf7,
	0x94, 0x60, 0x4e, 0xc2, 0x4b, 0x6e, 0x57, 0x3b, 0x46, 0xb7, 0x12, 0xac, 0x03, 0x82, 0xbd, 0x49,
	0x43, 0xc5, 0x5a, 0x92, 0x5d, 0x05, 0x84, 0xe2, 0x67, 0x9c, 0xf1, 0xa9, 0x5d, 0x93, 0x8a, 0x39,
	0x40, 0x8f, 0x01, 0xae, 0x31, 0x8f, 0x16, 0xe4, 0x0a, 0xc7, 0x13, 0x7b, 0x3f, 0xa7, 0x36, 0x22,
	0xa8, 0x03, 0x0d, 0x9f, 0x7d, 0xc3, 0xb3, 0x28, 0x97, 0xb1, 0xeb, 0x1d, 0xa3, 0x5b, 0x0f, 0x36,
	0x43, 0xae, 0x0f, 0xff, 0x07, 0x64, 0x12, 0x31, 0x4e, 0x68, 0x40, 0x7e, 0x65, 0x84, 0x71, 0x51,
	0x58, 0xa6, 0x0b, 0x33, 0x64, 0x61, 0x1a, 0x0b, 0x2e, 0xc5, 0x8c, 0xdd, 0x26, 0x34, 0xd4, 0x45,
	0x6b, 0xec, 0xf6, 0xe1, 0xbf, 0xb5, 0x54, 0x3a, 0x5b, 0xa2, 0xa7, 0xaa, 0xeb, 0xb9, 0x4a, 0xa3,
	0xdf, 0xe8, 0xe5, 0x73, 0xc9, 0x43, 0x81, 0x64, 0x5c, 0x1f, 0x4e, 0x75, 0x8e, 0xf2, 0x14, 0x25,
	0xb1, 0x36, 0xd2, 0x84, 0x4a, 0xc2, 0x53, 0xe5, 0x41, 0x1c, 0xb7, 0xac, 0x99, 0xdb, 0xd6, 0xdc,
	0x00, 0x1e, 0x94, 0x49, 0x09, 0x23, 0x2d, 0xb0, 0x78, 0xf2, 0x93, 0xc4, 0x4a, 0x4a, 0x82, 0xb5,
	0x3d, 0xf3, 0x5e, 0x7b, 0x1f, 0xe1, 0xf0, 0x2a, 0x99, 0x44, 0xf1, 0xbf, 0xb6, 0x66, 0x00, 0xa0,
	0x74, 0x94, 0x9d, 0xe1, 0xa6, 0x9d, 0xe1, 0xae, 0x76, 0x7c, 0x68, 0x5f, 0x66, 0x7c, 0x7a, 0xb7,
	0x53, 0xe5, 0x8a, 0x0e, 0xd4, 0xc5, 0xf5, 0xe1, 0x32, 0x5d, 0x75, 0x4b, 0x63, 0xf7, 0x07, 0x1c,
	0x17, 0xa5, 0x84, 0x35, 0x17, 0xaa, 0x62, 0x89, 0xd5, 0xc4, 0x8e, 0xa4, 0x07, 0xfd, 0x04, 0x82,
	0x9c, 0xdb, 0xc5, 0xe8, 0x35, 0xb4, 0xbe, 0x12, 0xbe, 0x5e, 0x44, 0xed, 0xf3, 0x04, 0x6a, 0x42,
	0x62, 0xf5, 0x8e, 0x14, 0x2a, 0xec, 0xb1, 0x59, 0xdc, 0x63, 0xf7, 0x2d, 0xb4, 0x0b, 0x7a, 0x2c,
	0x4d, 0x62, 0x46, 0x76, 0xf0, 0xd2, 0xff, 0x63, 0xca, 0x9a, 0xd0, 0x1b, 0xa8, 0x53, 0xb5, 0x20,
	0xa8, 0x2d, 0x2f, 0x16, 0x56, 0xdf, 0x39, 0x2e, 0x86, 0xd3, 0xd9, 0xd2, 0xdd, 0x43, 0x1e, 0x58,
	0x33, 0x31, 0x3e, 0x84, 0x24, 0xbf, 0xb9, 0x13, 0x4e, 0x73, 0x2b, 0x26, 0x13, 0x86, 0xd0, 0x5c,
	0xa8, 0x27, 0xa6, 0xb5, 0xd0, 0x93, 0x6d, 0xed, 0x3b, 0x43, 0x74, 0x1e, 0xdd, 0x7f, 0x41, 0xaa,
	0x7e, 0x82, 0x43, 0xad, 0x2a, 0x66, 0x87, 0xce, 0x64, 0x42, 0xe9, 0x4a, 0x38, 0xa7, 0xe5, 0xa4,
	0x54, 0xfa, 0x02, 0x2d, 0x56, 0x36, 0x1f, 0x47, 0x26, 0x95, 0xcd, 0xce, 0x39, 0x2b, 0xe5, 0xe4,
	0x1c, 0xdc, 0xbd, 0x77, 0xcf, 0xbf, 0x3f, 0x9b, 0x44, 0x7c, 0x9a, 0x8d, 0x7a, 0xe3, 0x64, 0xee,
	0x2d, 0x6f, 0x5f, 0x8d, 0x3d, 0x4e, 0x71, 0xcc, 0x5e, 0xa4, 0x23, 0x6f, 0xe3, 0xdf, 0x3c, 0xaa,
	0xe5, 0xff, 0xe5, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xcc, 0x71, 0xa2, 0xb1, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	ValidateRegister(ctx context.Context, in *RegisterValidationRequest, opts ...grpc.CallOption) (*RegisterValidationReply, error)
	ValidateAuth(ctx context.Context, in *AuthValidationRequest, opts ...grpc.CallOption) (*AuthValidationReply, error)
	SetNativeLangRequest(ctx context.Context, in *SetNativeLangRequest, opts ...grpc.CallOption) (*SetNativeLangResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/user.User/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/user.User/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ValidateRegister(ctx context.Context, in *RegisterValidationRequest, opts ...grpc.CallOption) (*RegisterValidationReply, error) {
	out := new(RegisterValidationReply)
	err := c.cc.Invoke(ctx, "/user.User/validateRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ValidateAuth(ctx context.Context, in *AuthValidationRequest, opts ...grpc.CallOption) (*AuthValidationReply, error) {
	out := new(AuthValidationReply)
	err := c.cc.Invoke(ctx, "/user.User/validateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetNativeLangRequest(ctx context.Context, in *SetNativeLangRequest, opts ...grpc.CallOption) (*SetNativeLangResponse, error) {
	out := new(SetNativeLangResponse)
	err := c.cc.Invoke(ctx, "/user.User/setNativeLangRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	ValidateRegister(context.Context, *RegisterValidationRequest) (*RegisterValidationReply, error)
	ValidateAuth(context.Context, *AuthValidationRequest) (*AuthValidationReply, error)
	SetNativeLangRequest(context.Context, *SetNativeLangRequest) (*SetNativeLangResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUserServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServer) ValidateRegister(ctx context.Context, req *RegisterValidationRequest) (*RegisterValidationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateRegister not implemented")
}
func (*UnimplementedUserServer) ValidateAuth(ctx context.Context, req *AuthValidationRequest) (*AuthValidationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuth not implemented")
}
func (*UnimplementedUserServer) SetNativeLangRequest(ctx context.Context, req *SetNativeLangRequest) (*SetNativeLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNativeLangRequest not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ValidateRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ValidateRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/ValidateRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ValidateRegister(ctx, req.(*RegisterValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ValidateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ValidateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/ValidateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ValidateAuth(ctx, req.(*AuthValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetNativeLangRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNativeLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetNativeLangRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SetNativeLangRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetNativeLangRequest(ctx, req.(*SetNativeLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "validateRegister",
			Handler:    _User_ValidateRegister_Handler,
		},
		{
			MethodName: "validateAuth",
			Handler:    _User_ValidateAuth_Handler,
		},
		{
			MethodName: "setNativeLangRequest",
			Handler:    _User_SetNativeLangRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "p010user/v1/user.proto",
}
