// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/timber.proto

package timber

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type Timber struct {
	Context              *TimberContext       `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content              map[string]*any.Any  `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Timber) Reset()         { *m = Timber{} }
func (m *Timber) String() string { return proto.CompactTextString(m) }
func (*Timber) ProtoMessage()    {}
func (*Timber) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_4721bf1355109cc7, []int{0}
}
func (m *Timber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timber.Unmarshal(m, b)
}
func (m *Timber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timber.Marshal(b, m, deterministic)
}
func (dst *Timber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timber.Merge(dst, src)
}
func (m *Timber) XXX_Size() int {
	return xxx_messageInfo_Timber.Size(m)
}
func (m *Timber) XXX_DiscardUnknown() {
	xxx_messageInfo_Timber.DiscardUnknown(m)
}

var xxx_messageInfo_Timber proto.InternalMessageInfo

func (m *Timber) GetContext() *TimberContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Timber) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Timber) GetContent() map[string]*any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

type TimberCollection struct {
	Context              *TimberContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Items                []*Timber      `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TimberCollection) Reset()         { *m = TimberCollection{} }
func (m *TimberCollection) String() string { return proto.CompactTextString(m) }
func (*TimberCollection) ProtoMessage()    {}
func (*TimberCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_4721bf1355109cc7, []int{1}
}
func (m *TimberCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimberCollection.Unmarshal(m, b)
}
func (m *TimberCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimberCollection.Marshal(b, m, deterministic)
}
func (dst *TimberCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimberCollection.Merge(dst, src)
}
func (m *TimberCollection) XXX_Size() int {
	return xxx_messageInfo_TimberCollection.Size(m)
}
func (m *TimberCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_TimberCollection.DiscardUnknown(m)
}

var xxx_messageInfo_TimberCollection proto.InternalMessageInfo

func (m *TimberCollection) GetContext() *TimberContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TimberCollection) GetItems() []*Timber {
	if m != nil {
		return m.Items
	}
	return nil
}

type TimberContext struct {
	KafkaTopic             string   `protobuf:"bytes,1,opt,name=kafka_topic,json=kafkaTopic,proto3" json:"kafka_topic,omitempty"`
	KafkaPartition         int32    `protobuf:"varint,2,opt,name=kafka_partition,json=kafkaPartition,proto3" json:"kafka_partition,omitempty"`
	KafkaReplicationFactor int32    `protobuf:"varint,3,opt,name=kafka_replication_factor,json=kafkaReplicationFactor,proto3" json:"kafka_replication_factor,omitempty"`
	EsIndexPrefix          string   `protobuf:"bytes,4,opt,name=es_index_prefix,json=esIndexPrefix,proto3" json:"es_index_prefix,omitempty"`
	EsDocumentType         string   `protobuf:"bytes,5,opt,name=es_document_type,json=esDocumentType,proto3" json:"es_document_type,omitempty"`
	AppMaxTps              int32    `protobuf:"varint,6,opt,name=app_max_tps,json=appMaxTps,proto3" json:"app_max_tps,omitempty"`
	AppSecret              string   `protobuf:"bytes,7,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TimberContext) Reset()         { *m = TimberContext{} }
func (m *TimberContext) String() string { return proto.CompactTextString(m) }
func (*TimberContext) ProtoMessage()    {}
func (*TimberContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_4721bf1355109cc7, []int{2}
}
func (m *TimberContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimberContext.Unmarshal(m, b)
}
func (m *TimberContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimberContext.Marshal(b, m, deterministic)
}
func (dst *TimberContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimberContext.Merge(dst, src)
}
func (m *TimberContext) XXX_Size() int {
	return xxx_messageInfo_TimberContext.Size(m)
}
func (m *TimberContext) XXX_DiscardUnknown() {
	xxx_messageInfo_TimberContext.DiscardUnknown(m)
}

var xxx_messageInfo_TimberContext proto.InternalMessageInfo

func (m *TimberContext) GetKafkaTopic() string {
	if m != nil {
		return m.KafkaTopic
	}
	return ""
}

func (m *TimberContext) GetKafkaPartition() int32 {
	if m != nil {
		return m.KafkaPartition
	}
	return 0
}

func (m *TimberContext) GetKafkaReplicationFactor() int32 {
	if m != nil {
		return m.KafkaReplicationFactor
	}
	return 0
}

func (m *TimberContext) GetEsIndexPrefix() string {
	if m != nil {
		return m.EsIndexPrefix
	}
	return ""
}

func (m *TimberContext) GetEsDocumentType() string {
	if m != nil {
		return m.EsDocumentType
	}
	return ""
}

func (m *TimberContext) GetAppMaxTps() int32 {
	if m != nil {
		return m.AppMaxTps
	}
	return 0
}

func (m *TimberContext) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

type ProduceResult struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResult) Reset()         { *m = ProduceResult{} }
func (m *ProduceResult) String() string { return proto.CompactTextString(m) }
func (*ProduceResult) ProtoMessage()    {}
func (*ProduceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_4721bf1355109cc7, []int{3}
}
func (m *ProduceResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResult.Unmarshal(m, b)
}
func (m *ProduceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResult.Marshal(b, m, deterministic)
}
func (dst *ProduceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResult.Merge(dst, src)
}
func (m *ProduceResult) XXX_Size() int {
	return xxx_messageInfo_ProduceResult.Size(m)
}
func (m *ProduceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResult.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResult proto.InternalMessageInfo

func (m *ProduceResult) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*Timber)(nil), "timber.Timber")
	proto.RegisterMapType((map[string]*any.Any)(nil), "timber.Timber.ContentEntry")
	proto.RegisterType((*TimberCollection)(nil), "timber.TimberCollection")
	proto.RegisterType((*TimberContext)(nil), "timber.TimberContext")
	proto.RegisterType((*ProduceResult)(nil), "timber.ProduceResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BaritoProducerClient is the client API for BaritoProducer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BaritoProducerClient interface {
	Produce(ctx context.Context, in *Timber, opts ...grpc.CallOption) (*ProduceResult, error)
	ProduceBatch(ctx context.Context, in *TimberCollection, opts ...grpc.CallOption) (*ProduceResult, error)
}

type baritoProducerClient struct {
	cc *grpc.ClientConn
}

func NewBaritoProducerClient(cc *grpc.ClientConn) BaritoProducerClient {
	return &baritoProducerClient{cc}
}

func (c *baritoProducerClient) Produce(ctx context.Context, in *Timber, opts ...grpc.CallOption) (*ProduceResult, error) {
	out := new(ProduceResult)
	err := c.cc.Invoke(ctx, "/timber.BaritoProducer/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baritoProducerClient) ProduceBatch(ctx context.Context, in *TimberCollection, opts ...grpc.CallOption) (*ProduceResult, error) {
	out := new(ProduceResult)
	err := c.cc.Invoke(ctx, "/timber.BaritoProducer/ProduceBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaritoProducerServer is the server API for BaritoProducer service.
type BaritoProducerServer interface {
	Produce(context.Context, *Timber) (*ProduceResult, error)
	ProduceBatch(context.Context, *TimberCollection) (*ProduceResult, error)
}

func RegisterBaritoProducerServer(s *grpc.Server, srv BaritoProducerServer) {
	s.RegisterService(&_BaritoProducer_serviceDesc, srv)
}

func _BaritoProducer_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaritoProducerServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timber.BaritoProducer/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaritoProducerServer).Produce(ctx, req.(*Timber))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaritoProducer_ProduceBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimberCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaritoProducerServer).ProduceBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timber.BaritoProducer/ProduceBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaritoProducerServer).ProduceBatch(ctx, req.(*TimberCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _BaritoProducer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timber.BaritoProducer",
	HandlerType: (*BaritoProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _BaritoProducer_Produce_Handler,
		},
		{
			MethodName: "ProduceBatch",
			Handler:    _BaritoProducer_ProduceBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/timber.proto",
}

func init() { proto.RegisterFile("proto/timber.proto", fileDescriptor_timber_4721bf1355109cc7) }

var fileDescriptor_timber_4721bf1355109cc7 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x8f, 0x12, 0x31,
	0x14, 0xc7, 0x03, 0x08, 0x84, 0xc7, 0xc2, 0x92, 0x66, 0x35, 0x15, 0xa3, 0xbb, 0x21, 0xfe, 0x20,
	0x1e, 0xc0, 0x60, 0x4c, 0x88, 0x17, 0xe3, 0xae, 0x9a, 0x78, 0x30, 0x21, 0x23, 0xf7, 0x49, 0x19,
	0x1e, 0x6b, 0xc3, 0xcc, 0xb4, 0x69, 0xdf, 0x98, 0x99, 0xb3, 0xff, 0x8a, 0x7f, 0xa6, 0x07, 0x33,
	0xed, 0xcc, 0xee, 0x82, 0xf1, 0xb0, 0xb7, 0xf6, 0xfb, 0xfd, 0xf4, 0x7d, 0xdf, 0x6b, 0x0b, 0x4c,
	0x1b, 0x45, 0x6a, 0x4e, 0x32, 0xd9, 0xa0, 0x99, 0xb9, 0x0d, 0xeb, 0xf8, 0xdd, 0xf8, 0xf1, 0xb5,
	0x52, 0xd7, 0x31, 0xce, 0x9d, 0xba, 0xc9, 0x76, 0x73, 0x91, 0x16, 0x1e, 0x19, 0x9f, 0x1f, 0x5b,
	0x24, 0x13, 0xb4, 0x24, 0x12, 0xed, 0x81, 0xc9, 0x9f, 0x06, 0x74, 0xd6, 0xae, 0x0c, 0x9b, 0x43,
	0x37, 0x52, 0x29, 0x61, 0x4e, 0xbc, 0x71, 0xd1, 0x98, 0xf6, 0x17, 0x0f, 0x67, 0x55, 0x9c, 0x07,
	0xae, 0xbc, 0x19, 0xd4, 0x14, 0x5b, 0x42, 0xef, 0xa6, 0x1c, 0x6f, 0xba, 0x23, 0xe3, 0x99, 0x0f,
	0x9c, 0xd5, 0x81, 0xe5, 0x59, 0x4f, 0x04, 0xb7, 0x30, 0x7b, 0x57, 0x45, 0xa5, 0xc4, 0x5b, 0x17,
	0xad, 0x69, 0x7f, 0xf1, 0xe4, 0x30, 0x6a, 0x76, 0xe5, 0xdd, 0xcf, 0x29, 0x99, 0x22, 0xa8, 0xd9,
	0xf1, 0x0a, 0x4e, 0xee, 0x1a, 0x6c, 0x04, 0xad, 0x3d, 0x16, 0xae, 0xdb, 0x5e, 0x50, 0x2e, 0xd9,
	0x6b, 0x68, 0xff, 0x14, 0x71, 0x86, 0x55, 0x3b, 0x67, 0xff, 0xb4, 0xf3, 0x31, 0x2d, 0x02, 0x8f,
	0xbc, 0x6f, 0x2e, 0x1b, 0x13, 0x09, 0xa3, 0x7a, 0xb8, 0x38, 0xc6, 0x88, 0xa4, 0x4a, 0xef, 0x7f,
	0x0f, 0xcf, 0xa1, 0x2d, 0x09, 0x13, 0xcb, 0x9b, 0x6e, 0x96, 0xe1, 0x21, 0x1e, 0x78, 0x73, 0xf2,
	0xbb, 0x09, 0x83, 0x83, 0x02, 0xec, 0x1c, 0xfa, 0x7b, 0xb1, 0xdb, 0x8b, 0x90, 0x94, 0x96, 0x51,
	0x35, 0x06, 0x38, 0x69, 0x5d, 0x2a, 0xec, 0x15, 0x9c, 0x7a, 0x40, 0x0b, 0x43, 0xb2, 0x6c, 0xce,
	0xcd, 0xd5, 0x0e, 0x86, 0x4e, 0x5e, 0xd5, 0x2a, 0x5b, 0x02, 0xf7, 0xa0, 0x41, 0x1d, 0xcb, 0x48,
	0x94, 0x62, 0xb8, 0x13, 0x11, 0x29, 0xc3, 0x5b, 0xee, 0xc4, 0x23, 0xe7, 0x07, 0xb7, 0xf6, 0x17,
	0xe7, 0xb2, 0x97, 0x70, 0x8a, 0x36, 0x94, 0xe9, 0x16, 0xf3, 0x50, 0x1b, 0xdc, 0xc9, 0x9c, 0x3f,
	0x70, 0x7d, 0x0c, 0xd0, 0x7e, 0x2d, 0xd5, 0x95, 0x13, 0xd9, 0x14, 0x46, 0x68, 0xc3, 0xad, 0x8a,
	0xb2, 0x04, 0x53, 0x0a, 0xa9, 0xd0, 0xc8, 0xdb, 0x0e, 0x1c, 0xa2, 0xfd, 0x54, 0xc9, 0xeb, 0x42,
	0x23, 0x7b, 0x06, 0x7d, 0xa1, 0x75, 0x98, 0x88, 0x3c, 0x24, 0x6d, 0x79, 0xc7, 0xc5, 0xf7, 0x84,
	0xd6, 0xdf, 0x44, 0xbe, 0xd6, 0x96, 0x3d, 0x05, 0x28, 0x7d, 0x8b, 0x91, 0x41, 0xe2, 0x5d, 0x57,
	0xa3, 0xb4, 0xbf, 0x3b, 0x61, 0xf2, 0x02, 0x06, 0x2b, 0xa3, 0xb6, 0x59, 0x84, 0x01, 0xda, 0x2c,
	0x26, 0x76, 0x06, 0xed, 0xbb, 0xf7, 0xe3, 0x37, 0x8b, 0x5f, 0x0d, 0x18, 0x5e, 0x0a, 0x23, 0x49,
	0x55, 0xb4, 0x61, 0x6f, 0xa0, 0x5b, 0xad, 0xd9, 0xd1, 0x13, 0x8c, 0x6f, 0x5e, 0xf0, 0xb0, 0xf4,
	0x07, 0x38, 0xa9, 0x84, 0x4b, 0x41, 0xd1, 0x0f, 0xc6, 0x8f, 0x1f, 0xba, 0xfe, 0x13, 0xff, 0x29,
	0xb0, 0xe9, 0xb8, 0x7f, 0xf5, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x09, 0x37, 0x69,
	0x9e, 0x03, 0x00, 0x00,
}