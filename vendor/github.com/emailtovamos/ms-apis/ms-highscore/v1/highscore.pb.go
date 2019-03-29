// Code generated by protoc-gen-go. DO NOT EDIT.
// source: highscore.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SetHighScoreRequest struct {
	HighScore            float64  `protobuf:"fixed64,1,opt,name=highScore,proto3" json:"highScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetHighScoreRequest) Reset()         { *m = SetHighScoreRequest{} }
func (m *SetHighScoreRequest) String() string { return proto.CompactTextString(m) }
func (*SetHighScoreRequest) ProtoMessage()    {}
func (*SetHighScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bfe52306d72d22, []int{0}
}

func (m *SetHighScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetHighScoreRequest.Unmarshal(m, b)
}
func (m *SetHighScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetHighScoreRequest.Marshal(b, m, deterministic)
}
func (m *SetHighScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetHighScoreRequest.Merge(m, src)
}
func (m *SetHighScoreRequest) XXX_Size() int {
	return xxx_messageInfo_SetHighScoreRequest.Size(m)
}
func (m *SetHighScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetHighScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetHighScoreRequest proto.InternalMessageInfo

func (m *SetHighScoreRequest) GetHighScore() float64 {
	if m != nil {
		return m.HighScore
	}
	return 0
}

type SetHighScoreResponse struct {
	Set                  bool     `protobuf:"varint,1,opt,name=set,proto3" json:"set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetHighScoreResponse) Reset()         { *m = SetHighScoreResponse{} }
func (m *SetHighScoreResponse) String() string { return proto.CompactTextString(m) }
func (*SetHighScoreResponse) ProtoMessage()    {}
func (*SetHighScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bfe52306d72d22, []int{1}
}

func (m *SetHighScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetHighScoreResponse.Unmarshal(m, b)
}
func (m *SetHighScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetHighScoreResponse.Marshal(b, m, deterministic)
}
func (m *SetHighScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetHighScoreResponse.Merge(m, src)
}
func (m *SetHighScoreResponse) XXX_Size() int {
	return xxx_messageInfo_SetHighScoreResponse.Size(m)
}
func (m *SetHighScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetHighScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetHighScoreResponse proto.InternalMessageInfo

func (m *SetHighScoreResponse) GetSet() bool {
	if m != nil {
		return m.Set
	}
	return false
}

type GetHighScoreRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHighScoreRequest) Reset()         { *m = GetHighScoreRequest{} }
func (m *GetHighScoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetHighScoreRequest) ProtoMessage()    {}
func (*GetHighScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bfe52306d72d22, []int{2}
}

func (m *GetHighScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHighScoreRequest.Unmarshal(m, b)
}
func (m *GetHighScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHighScoreRequest.Marshal(b, m, deterministic)
}
func (m *GetHighScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHighScoreRequest.Merge(m, src)
}
func (m *GetHighScoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetHighScoreRequest.Size(m)
}
func (m *GetHighScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHighScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHighScoreRequest proto.InternalMessageInfo

type GetHighScoreResponse struct {
	HighScore            float64  `protobuf:"fixed64,1,opt,name=highScore,proto3" json:"highScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHighScoreResponse) Reset()         { *m = GetHighScoreResponse{} }
func (m *GetHighScoreResponse) String() string { return proto.CompactTextString(m) }
func (*GetHighScoreResponse) ProtoMessage()    {}
func (*GetHighScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1bfe52306d72d22, []int{3}
}

func (m *GetHighScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHighScoreResponse.Unmarshal(m, b)
}
func (m *GetHighScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHighScoreResponse.Marshal(b, m, deterministic)
}
func (m *GetHighScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHighScoreResponse.Merge(m, src)
}
func (m *GetHighScoreResponse) XXX_Size() int {
	return xxx_messageInfo_GetHighScoreResponse.Size(m)
}
func (m *GetHighScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHighScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHighScoreResponse proto.InternalMessageInfo

func (m *GetHighScoreResponse) GetHighScore() float64 {
	if m != nil {
		return m.HighScore
	}
	return 0
}

func init() {
	proto.RegisterType((*SetHighScoreRequest)(nil), "ms.highscore.v1.SetHighScoreRequest")
	proto.RegisterType((*SetHighScoreResponse)(nil), "ms.highscore.v1.SetHighScoreResponse")
	proto.RegisterType((*GetHighScoreRequest)(nil), "ms.highscore.v1.GetHighScoreRequest")
	proto.RegisterType((*GetHighScoreResponse)(nil), "ms.highscore.v1.GetHighScoreResponse")
}

func init() { proto.RegisterFile("highscore.proto", fileDescriptor_a1bfe52306d72d22) }

var fileDescriptor_a1bfe52306d72d22 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xc8, 0x4c, 0xcf,
	0x28, 0x4e, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0x2d, 0xd6,
	0x43, 0x88, 0x95, 0x19, 0x2a, 0x19, 0x73, 0x09, 0x07, 0xa7, 0x96, 0x78, 0x64, 0xa6, 0x67, 0x04,
	0x83, 0x84, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0x33, 0x60, 0x62,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x08, 0x01, 0x25, 0x0d, 0x2e, 0x11, 0x54, 0x4d, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x02, 0x5c, 0xcc, 0xc5, 0xa9, 0x25, 0x60, 0xf5, 0x1c, 0x41,
	0x20, 0xa6, 0x92, 0x28, 0x97, 0xb0, 0x3b, 0xa6, 0xf1, 0x4a, 0x26, 0x5c, 0x22, 0xee, 0xd8, 0x0c,
	0xc0, 0x6b, 0xad, 0xd1, 0x01, 0x46, 0x2e, 0x16, 0xf7, 0xc4, 0xdc, 0x54, 0xa1, 0x68, 0x2e, 0x1e,
	0x64, 0xfb, 0x85, 0x54, 0xf4, 0xd0, 0xbc, 0xa5, 0x87, 0xc5, 0x4f, 0x52, 0xaa, 0x04, 0x54, 0x41,
	0xdd, 0x10, 0xcd, 0xc5, 0xe3, 0x8e, 0xdf, 0x70, 0x77, 0xa2, 0x0c, 0xc7, 0xe6, 0x41, 0x27, 0xb6,
	0x28, 0x96, 0xf4, 0xc4, 0xdc, 0xd4, 0x24, 0x36, 0x70, 0x74, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7e, 0x72, 0x86, 0x32, 0xa1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	SetHighScore(ctx context.Context, in *SetHighScoreRequest, opts ...grpc.CallOption) (*SetHighScoreResponse, error)
	GetHighScore(ctx context.Context, in *GetHighScoreRequest, opts ...grpc.CallOption) (*GetHighScoreResponse, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) SetHighScore(ctx context.Context, in *SetHighScoreRequest, opts ...grpc.CallOption) (*SetHighScoreResponse, error) {
	out := new(SetHighScoreResponse)
	err := c.cc.Invoke(ctx, "/ms.highscore.v1.Game/SetHighScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetHighScore(ctx context.Context, in *GetHighScoreRequest, opts ...grpc.CallOption) (*GetHighScoreResponse, error) {
	out := new(GetHighScoreResponse)
	err := c.cc.Invoke(ctx, "/ms.highscore.v1.Game/GetHighScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	SetHighScore(context.Context, *SetHighScoreRequest) (*SetHighScoreResponse, error)
	GetHighScore(context.Context, *GetHighScoreRequest) (*GetHighScoreResponse, error)
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_SetHighScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHighScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SetHighScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ms.highscore.v1.Game/SetHighScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SetHighScore(ctx, req.(*SetHighScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetHighScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHighScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetHighScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ms.highscore.v1.Game/GetHighScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetHighScore(ctx, req.(*GetHighScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ms.highscore.v1.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHighScore",
			Handler:    _Game_SetHighScore_Handler,
		},
		{
			MethodName: "GetHighScore",
			Handler:    _Game_GetHighScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "highscore.proto",
}
