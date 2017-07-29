// Code generated by protoc-gen-go.
// source: collector.proto
// DO NOT EDIT!

/*
Package collector is a generated protocol buffer package.

It is generated from these files:
	collector.proto
	collector_status.proto

It has these top-level messages:
	HostID
	OboeSetting
	MessageRequest
	MessageResult
	SettingsRequest
	SettingsResult
	ConnectionInfo
	ConnectionStatusRequest
	ConnectionStatusResult
*/
package collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type ResultCode int32

const (
	ResultCode_OK              ResultCode = 0
	ResultCode_TRY_LATER       ResultCode = 1
	ResultCode_INVALID_API_KEY ResultCode = 2
	ResultCode_LIMIT_EXCEEDED  ResultCode = 3
	ResultCode_REDIRECT        ResultCode = 4
)

var ResultCode_name = map[int32]string{
	0: "OK",
	1: "TRY_LATER",
	2: "INVALID_API_KEY",
	3: "LIMIT_EXCEEDED",
	4: "REDIRECT",
}
var ResultCode_value = map[string]int32{
	"OK":              0,
	"TRY_LATER":       1,
	"INVALID_API_KEY": 2,
	"LIMIT_EXCEEDED":  3,
	"REDIRECT":        4,
}

func (x ResultCode) String() string {
	return proto.EnumName(ResultCode_name, int32(x))
}
func (ResultCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EncodingType int32

const (
	EncodingType_BSON     EncodingType = 0
	EncodingType_PROTOBUF EncodingType = 1
)

var EncodingType_name = map[int32]string{
	0: "BSON",
	1: "PROTOBUF",
}
var EncodingType_value = map[string]int32{
	"BSON":     0,
	"PROTOBUF": 1,
}

func (x EncodingType) String() string {
	return proto.EnumName(EncodingType_name, int32(x))
}
func (EncodingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type OboeSettingType int32

const (
	OboeSettingType_DEFAULT_SAMPLE_RATE        OboeSettingType = 0
	OboeSettingType_LAYER_SAMPLE_RATE          OboeSettingType = 1
	OboeSettingType_LAYER_APP_SAMPLE_RATE      OboeSettingType = 2
	OboeSettingType_LAYER_HTTPHOST_SAMPLE_RATE OboeSettingType = 3
	OboeSettingType_CONFIG_STRING              OboeSettingType = 4
	OboeSettingType_CONFIG_INT                 OboeSettingType = 5
)

var OboeSettingType_name = map[int32]string{
	0: "DEFAULT_SAMPLE_RATE",
	1: "LAYER_SAMPLE_RATE",
	2: "LAYER_APP_SAMPLE_RATE",
	3: "LAYER_HTTPHOST_SAMPLE_RATE",
	4: "CONFIG_STRING",
	5: "CONFIG_INT",
}
var OboeSettingType_value = map[string]int32{
	"DEFAULT_SAMPLE_RATE":        0,
	"LAYER_SAMPLE_RATE":          1,
	"LAYER_APP_SAMPLE_RATE":      2,
	"LAYER_HTTPHOST_SAMPLE_RATE": 3,
	"CONFIG_STRING":              4,
	"CONFIG_INT":                 5,
}

func (x OboeSettingType) String() string {
	return proto.EnumName(OboeSettingType_name, int32(x))
}
func (OboeSettingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type HostID struct {
	Hostname    string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	IpAddresses []string `protobuf:"bytes,2,rep,name=ip_addresses,json=ipAddresses" json:"ip_addresses,omitempty"`
	Uuid        string   `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *HostID) Reset()                    { *m = HostID{} }
func (m *HostID) String() string            { return proto.CompactTextString(m) }
func (*HostID) ProtoMessage()               {}
func (*HostID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HostID) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostID) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

func (m *HostID) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type OboeSetting struct {
	Type      OboeSettingType   `protobuf:"varint,1,opt,name=type,enum=collector.OboeSettingType" json:"type,omitempty"`
	Flags     []byte            `protobuf:"bytes,2,opt,name=flags,proto3" json:"flags,omitempty"`
	Timestamp int64             `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Value     int64             `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	Layer     []byte            `protobuf:"bytes,5,opt,name=layer,proto3" json:"layer,omitempty"`
	Arguments map[string][]byte `protobuf:"bytes,7,rep,name=arguments" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ttl       int64             `protobuf:"varint,8,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *OboeSetting) Reset()                    { *m = OboeSetting{} }
func (m *OboeSetting) String() string            { return proto.CompactTextString(m) }
func (*OboeSetting) ProtoMessage()               {}
func (*OboeSetting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OboeSetting) GetType() OboeSettingType {
	if m != nil {
		return m.Type
	}
	return OboeSettingType_DEFAULT_SAMPLE_RATE
}

func (m *OboeSetting) GetFlags() []byte {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *OboeSetting) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *OboeSetting) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *OboeSetting) GetLayer() []byte {
	if m != nil {
		return m.Layer
	}
	return nil
}

func (m *OboeSetting) GetArguments() map[string][]byte {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *OboeSetting) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type MessageRequest struct {
	ApiKey   string       `protobuf:"bytes,1,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	Messages [][]byte     `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Encoding EncodingType `protobuf:"varint,3,opt,name=encoding,enum=collector.EncodingType" json:"encoding,omitempty"`
}

func (m *MessageRequest) Reset()                    { *m = MessageRequest{} }
func (m *MessageRequest) String() string            { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()               {}
func (*MessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MessageRequest) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *MessageRequest) GetMessages() [][]byte {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *MessageRequest) GetEncoding() EncodingType {
	if m != nil {
		return m.Encoding
	}
	return EncodingType_BSON
}

type MessageResult struct {
	Result ResultCode `protobuf:"varint,1,opt,name=result,enum=collector.ResultCode" json:"result,omitempty"`
	Arg    string     `protobuf:"bytes,2,opt,name=arg" json:"arg,omitempty"`
}

func (m *MessageResult) Reset()                    { *m = MessageResult{} }
func (m *MessageResult) String() string            { return proto.CompactTextString(m) }
func (*MessageResult) ProtoMessage()               {}
func (*MessageResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MessageResult) GetResult() ResultCode {
	if m != nil {
		return m.Result
	}
	return ResultCode_OK
}

func (m *MessageResult) GetArg() string {
	if m != nil {
		return m.Arg
	}
	return ""
}

type SettingsRequest struct {
	ApiKey        string  `protobuf:"bytes,1,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	Identity      *HostID `protobuf:"bytes,2,opt,name=identity" json:"identity,omitempty"`
	ClientVersion string  `protobuf:"bytes,3,opt,name=clientVersion" json:"clientVersion,omitempty"`
}

func (m *SettingsRequest) Reset()                    { *m = SettingsRequest{} }
func (m *SettingsRequest) String() string            { return proto.CompactTextString(m) }
func (*SettingsRequest) ProtoMessage()               {}
func (*SettingsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SettingsRequest) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *SettingsRequest) GetIdentity() *HostID {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *SettingsRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type SettingsResult struct {
	Result   ResultCode     `protobuf:"varint,1,opt,name=result,enum=collector.ResultCode" json:"result,omitempty"`
	Arg      string         `protobuf:"bytes,2,opt,name=arg" json:"arg,omitempty"`
	Settings []*OboeSetting `protobuf:"bytes,3,rep,name=settings" json:"settings,omitempty"`
}

func (m *SettingsResult) Reset()                    { *m = SettingsResult{} }
func (m *SettingsResult) String() string            { return proto.CompactTextString(m) }
func (*SettingsResult) ProtoMessage()               {}
func (*SettingsResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SettingsResult) GetResult() ResultCode {
	if m != nil {
		return m.Result
	}
	return ResultCode_OK
}

func (m *SettingsResult) GetArg() string {
	if m != nil {
		return m.Arg
	}
	return ""
}

func (m *SettingsResult) GetSettings() []*OboeSetting {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*HostID)(nil), "collector.HostID")
	proto.RegisterType((*OboeSetting)(nil), "collector.OboeSetting")
	proto.RegisterType((*MessageRequest)(nil), "collector.MessageRequest")
	proto.RegisterType((*MessageResult)(nil), "collector.MessageResult")
	proto.RegisterType((*SettingsRequest)(nil), "collector.SettingsRequest")
	proto.RegisterType((*SettingsResult)(nil), "collector.SettingsResult")
	proto.RegisterEnum("collector.ResultCode", ResultCode_name, ResultCode_value)
	proto.RegisterEnum("collector.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("collector.OboeSettingType", OboeSettingType_name, OboeSettingType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TraceCollector service

type TraceCollectorClient interface {
	PostEvents(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error)
	PostMetrics(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error)
	PostStatus(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error)
	GetSettings(ctx context.Context, in *SettingsRequest, opts ...grpc.CallOption) (*SettingsResult, error)
}

type traceCollectorClient struct {
	cc *grpc.ClientConn
}

func NewTraceCollectorClient(cc *grpc.ClientConn) TraceCollectorClient {
	return &traceCollectorClient{cc}
}

func (c *traceCollectorClient) PostEvents(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := grpc.Invoke(ctx, "/collector.TraceCollector/postEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) PostMetrics(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := grpc.Invoke(ctx, "/collector.TraceCollector/postMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) PostStatus(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := grpc.Invoke(ctx, "/collector.TraceCollector/postStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) GetSettings(ctx context.Context, in *SettingsRequest, opts ...grpc.CallOption) (*SettingsResult, error) {
	out := new(SettingsResult)
	err := grpc.Invoke(ctx, "/collector.TraceCollector/getSettings", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TraceCollector service

type TraceCollectorServer interface {
	PostEvents(context.Context, *MessageRequest) (*MessageResult, error)
	PostMetrics(context.Context, *MessageRequest) (*MessageResult, error)
	PostStatus(context.Context, *MessageRequest) (*MessageResult, error)
	GetSettings(context.Context, *SettingsRequest) (*SettingsResult, error)
}

func RegisterTraceCollectorServer(s *grpc.Server, srv TraceCollectorServer) {
	s.RegisterService(&_TraceCollector_serviceDesc, srv)
}

func _TraceCollector_PostEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).PostEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/PostEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).PostEvents(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_PostMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).PostMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/PostMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).PostMetrics(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_PostStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).PostStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/PostStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).PostStatus(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/GetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).GetSettings(ctx, req.(*SettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceCollector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "collector.TraceCollector",
	HandlerType: (*TraceCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "postEvents",
			Handler:    _TraceCollector_PostEvents_Handler,
		},
		{
			MethodName: "postMetrics",
			Handler:    _TraceCollector_PostMetrics_Handler,
		},
		{
			MethodName: "postStatus",
			Handler:    _TraceCollector_PostStatus_Handler,
		},
		{
			MethodName: "getSettings",
			Handler:    _TraceCollector_GetSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}

func init() { proto.RegisterFile("collector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x6f, 0xe2, 0x48,
	0x10, 0xc6, 0x98, 0x10, 0x28, 0xc0, 0x38, 0x95, 0xcd, 0xc6, 0x41, 0xab, 0x55, 0x16, 0xed, 0xae,
	0x50, 0xa4, 0x70, 0x20, 0x97, 0xd5, 0x6a, 0x2f, 0x0e, 0x98, 0xc4, 0x0a, 0x2f, 0x35, 0x4e, 0xb4,
	0xd9, 0x3d, 0x58, 0x0e, 0xf4, 0x30, 0xd6, 0x18, 0xdb, 0xe3, 0x6e, 0x22, 0x31, 0x97, 0x39, 0xcd,
	0xef, 0x98, 0xeb, 0x9c, 0xe7, 0x17, 0x8e, 0xdc, 0x36, 0xaf, 0x08, 0x69, 0xa4, 0x68, 0x6e, 0x55,
	0x5f, 0x7d, 0xfd, 0xd5, 0xa3, 0xab, 0x1b, 0xaa, 0x93, 0xc0, 0xf3, 0xe8, 0x84, 0x07, 0x51, 0x33,
	0x8c, 0x02, 0x1e, 0x60, 0x71, 0x0d, 0xd4, 0xff, 0x87, 0xfc, 0x6d, 0xc0, 0xb8, 0xd9, 0xc1, 0x1a,
	0x14, 0xde, 0x06, 0x8c, 0xfb, 0xce, 0x9c, 0x6a, 0xd2, 0xb9, 0xd4, 0x28, 0x92, 0xb5, 0x8f, 0xbf,
	0x41, 0xd9, 0x0d, 0x6d, 0x67, 0x3a, 0x8d, 0x28, 0x63, 0x94, 0x69, 0xd9, 0x73, 0xb9, 0x51, 0x24,
	0x25, 0x37, 0xd4, 0x57, 0x10, 0x22, 0xe4, 0x16, 0x0b, 0x77, 0xaa, 0xc9, 0xe2, 0xa8, 0xb0, 0xeb,
	0x5f, 0xb3, 0x50, 0x1a, 0x3e, 0x05, 0x74, 0x4c, 0x39, 0x77, 0xfd, 0x19, 0x36, 0x21, 0xc7, 0x97,
	0x61, 0x22, 0xaf, 0xb4, 0x6a, 0xcd, 0x4d, 0x5d, 0x5b, 0x2c, 0x6b, 0x19, 0x52, 0x22, 0x78, 0xf8,
	0x13, 0x1c, 0xbc, 0xf1, 0x9c, 0x59, 0x9c, 0x4f, 0x6a, 0x94, 0x49, 0xe2, 0xe0, 0x2f, 0x50, 0xe4,
	0xee, 0x9c, 0x32, 0xee, 0xcc, 0x43, 0x91, 0x4e, 0x26, 0x1b, 0x20, 0x3e, 0xf3, 0xec, 0x78, 0x0b,
	0xaa, 0xe5, 0x44, 0x24, 0x71, 0x62, 0xd4, 0x73, 0x96, 0x34, 0xd2, 0x0e, 0x12, 0x25, 0xe1, 0x60,
	0x1b, 0x8a, 0x4e, 0x34, 0x5b, 0xcc, 0xa9, 0xcf, 0x99, 0x76, 0x78, 0x2e, 0x37, 0x4a, 0xad, 0x3f,
	0xf6, 0x17, 0xd5, 0xd4, 0x57, 0x3c, 0xc3, 0xe7, 0xd1, 0x92, 0x6c, 0xce, 0xa1, 0x0a, 0x32, 0xe7,
	0x9e, 0x56, 0x10, 0xe9, 0x62, 0xb3, 0xf6, 0x0f, 0x28, 0xbb, 0xf4, 0x98, 0xf3, 0x8e, 0x2e, 0xd3,
	0xb1, 0xc6, 0xe6, 0xa6, 0xcc, 0xb4, 0x35, 0xe1, 0xfc, 0x9d, 0xfd, 0x4b, 0xaa, 0x7f, 0x00, 0xa5,
	0x4f, 0x19, 0x73, 0x66, 0x94, 0xd0, 0xf7, 0x0b, 0xca, 0x38, 0x9e, 0xc2, 0xa1, 0x13, 0xba, 0xf6,
	0x46, 0x21, 0xef, 0x84, 0xee, 0x1d, 0x5d, 0xc6, 0x57, 0x36, 0x4f, 0xa8, 0xc9, 0x95, 0x94, 0xc9,
	0xda, 0xc7, 0x2b, 0x28, 0x50, 0x7f, 0x12, 0x4c, 0x5d, 0x7f, 0x26, 0x86, 0xa4, 0xb4, 0x4e, 0xb7,
	0x5a, 0x33, 0xd2, 0x90, 0x18, 0xf6, 0x9a, 0x58, 0x1f, 0x41, 0x65, 0x9d, 0x9b, 0x2d, 0x3c, 0x8e,
	0x97, 0x90, 0x8f, 0x84, 0x95, 0xde, 0xd9, 0xc9, 0x96, 0x46, 0x42, 0x69, 0x07, 0x53, 0x4a, 0x52,
	0x52, 0xdc, 0xa7, 0x13, 0xcd, 0x44, 0x4f, 0x45, 0x12, 0x9b, 0xf5, 0x8f, 0x50, 0x4d, 0x47, 0xc8,
	0xbe, 0xdb, 0xce, 0x25, 0x14, 0xdc, 0x29, 0xf5, 0xb9, 0xcb, 0x97, 0x42, 0xa2, 0xd4, 0x3a, 0xda,
	0x4a, 0x97, 0xac, 0x29, 0x59, 0x53, 0xf0, 0x77, 0xa8, 0x4c, 0x3c, 0x97, 0xfa, 0xfc, 0x81, 0x46,
	0xcc, 0x0d, 0xfc, 0x74, 0xf5, 0x76, 0xc1, 0xfa, 0x27, 0x09, 0x94, 0x4d, 0x05, 0x3f, 0xa4, 0x29,
	0x6c, 0x41, 0x81, 0xa5, 0x92, 0x9a, 0x2c, 0xd6, 0xe6, 0xe7, 0xfd, 0x6b, 0x43, 0xd6, 0xbc, 0x8b,
	0xff, 0x00, 0x36, 0xda, 0x98, 0x87, 0xec, 0xf0, 0x4e, 0xcd, 0x60, 0x05, 0x8a, 0x16, 0x79, 0xb4,
	0x7b, 0xba, 0x65, 0x10, 0x55, 0xc2, 0x63, 0xa8, 0x9a, 0x83, 0x07, 0xbd, 0x67, 0x76, 0x6c, 0x7d,
	0x64, 0xda, 0x77, 0xc6, 0xa3, 0x9a, 0x45, 0x04, 0xa5, 0x67, 0xf6, 0x4d, 0xcb, 0x36, 0xfe, 0x6d,
	0x1b, 0x46, 0xc7, 0xe8, 0xa8, 0x32, 0x96, 0xa1, 0x40, 0x8c, 0x8e, 0x49, 0x8c, 0xb6, 0xa5, 0xe6,
	0x2e, 0xfe, 0x84, 0xf2, 0xf6, 0x85, 0x62, 0x01, 0x72, 0xd7, 0xe3, 0xe1, 0x40, 0xcd, 0xc4, 0xbc,
	0x11, 0x19, 0x5a, 0xc3, 0xeb, 0xfb, 0xae, 0x2a, 0x5d, 0x7c, 0x96, 0xa0, 0xfa, 0xe2, 0xa5, 0xe1,
	0x29, 0x1c, 0x77, 0x8c, 0xae, 0x7e, 0xdf, 0xb3, 0xec, 0xb1, 0xde, 0x1f, 0xf5, 0x0c, 0x9b, 0xe8,
	0x96, 0xa1, 0x66, 0xf0, 0x04, 0x8e, 0x7a, 0xfa, 0xa3, 0x41, 0x76, 0x60, 0x09, 0xcf, 0xe0, 0x24,
	0x81, 0xf5, 0xd1, 0x68, 0x27, 0x94, 0xc5, 0x5f, 0xa1, 0x96, 0x84, 0x6e, 0x2d, 0x6b, 0x74, 0x3b,
	0x1c, 0xef, 0x2a, 0xca, 0x78, 0x04, 0x95, 0xf6, 0x70, 0xd0, 0x35, 0x6f, 0xec, 0xb1, 0x45, 0xcc,
	0xc1, 0x8d, 0x9a, 0x43, 0x05, 0x20, 0x85, 0xcc, 0x81, 0xa5, 0x1e, 0xb4, 0xbe, 0x64, 0x41, 0xb1,
	0x22, 0x67, 0x42, 0xdb, 0xab, 0x71, 0x62, 0x1b, 0x20, 0x0c, 0x18, 0x37, 0x9e, 0xc5, 0x6b, 0x3b,
	0xdb, 0x1a, 0xf4, 0xee, 0x33, 0xa9, 0x69, 0xfb, 0x42, 0xf1, 0xc4, 0xeb, 0x19, 0xec, 0x40, 0x29,
	0x16, 0xe9, 0x53, 0x1e, 0xb9, 0x93, 0x57, 0xab, 0xa4, 0xa5, 0x8c, 0xb9, 0xc3, 0x17, 0xaf, 0x16,
	0xe9, 0x42, 0x69, 0x46, 0xf9, 0x6a, 0x25, 0x71, 0xfb, 0x17, 0x7c, 0xf1, 0x52, 0x6a, 0x67, 0x7b,
	0x63, 0x89, 0xce, 0x53, 0x5e, 0xfc, 0xe5, 0x57, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x82,
	0xa1, 0x55, 0xde, 0x05, 0x00, 0x00,
}
