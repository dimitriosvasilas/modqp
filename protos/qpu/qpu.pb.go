// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qpu.proto

package qpu // import "github.com/dvasilas/proteus/protos/qpu"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import utils "github.com/dvasilas/proteus/protos/utils"

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

type ResponseStreamRecord_StreamRecordType int32

const (
	ResponseStreamRecord_UPDATEOP    ResponseStreamRecord_StreamRecordType = 0
	ResponseStreamRecord_UPDATEDELTA ResponseStreamRecord_StreamRecordType = 1
	ResponseStreamRecord_STATE       ResponseStreamRecord_StreamRecordType = 2
	ResponseStreamRecord_HEARTBEAT   ResponseStreamRecord_StreamRecordType = 3
)

var ResponseStreamRecord_StreamRecordType_name = map[int32]string{
	0: "UPDATEOP",
	1: "UPDATEDELTA",
	2: "STATE",
	3: "HEARTBEAT",
}
var ResponseStreamRecord_StreamRecordType_value = map[string]int32{
	"UPDATEOP":    0,
	"UPDATEDELTA": 1,
	"STATE":       2,
	"HEARTBEAT":   3,
}

func (x ResponseStreamRecord_StreamRecordType) String() string {
	return proto.EnumName(ResponseStreamRecord_StreamRecordType_name, int32(x))
}
func (ResponseStreamRecord_StreamRecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{1, 0}
}

type ConfigResponse_QPUType int32

const (
	ConfigResponse_DBDRIVER              ConfigResponse_QPUType = 0
	ConfigResponse_FILTER                ConfigResponse_QPUType = 1
	ConfigResponse_INDEX                 ConfigResponse_QPUType = 2
	ConfigResponse_CACHE                 ConfigResponse_QPUType = 3
	ConfigResponse_FEDERATION_DISPATCHER ConfigResponse_QPUType = 4
)

var ConfigResponse_QPUType_name = map[int32]string{
	0: "DBDRIVER",
	1: "FILTER",
	2: "INDEX",
	3: "CACHE",
	4: "FEDERATION_DISPATCHER",
}
var ConfigResponse_QPUType_value = map[string]int32{
	"DBDRIVER":              0,
	"FILTER":                1,
	"INDEX":                 2,
	"CACHE":                 3,
	"FEDERATION_DISPATCHER": 4,
}

func (x ConfigResponse_QPUType) String() string {
	return proto.EnumName(ConfigResponse_QPUType_name, int32(x))
}
func (ConfigResponse_QPUType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{5, 0}
}

type RequestStream struct {
	// Types that are valid to be assigned to Payload:
	//	*RequestStream_Request
	//	*RequestStream_Ack
	Payload              isRequestStream_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RequestStream) Reset()         { *m = RequestStream{} }
func (m *RequestStream) String() string { return proto.CompactTextString(m) }
func (*RequestStream) ProtoMessage()    {}
func (*RequestStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{0}
}
func (m *RequestStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestStream.Unmarshal(m, b)
}
func (m *RequestStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestStream.Marshal(b, m, deterministic)
}
func (dst *RequestStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestStream.Merge(dst, src)
}
func (m *RequestStream) XXX_Size() int {
	return xxx_messageInfo_RequestStream.Size(m)
}
func (m *RequestStream) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestStream.DiscardUnknown(m)
}

var xxx_messageInfo_RequestStream proto.InternalMessageInfo

type isRequestStream_Payload interface {
	isRequestStream_Payload()
}

type RequestStream_Request struct {
	Request *QueryRequest `protobuf:"bytes,1,opt,name=request,oneof"`
}
type RequestStream_Ack struct {
	Ack *AckMsg `protobuf:"bytes,2,opt,name=ack,oneof"`
}

func (*RequestStream_Request) isRequestStream_Payload() {}
func (*RequestStream_Ack) isRequestStream_Payload()     {}

func (m *RequestStream) GetPayload() isRequestStream_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RequestStream) GetRequest() *QueryRequest {
	if x, ok := m.GetPayload().(*RequestStream_Request); ok {
		return x.Request
	}
	return nil
}

func (m *RequestStream) GetAck() *AckMsg {
	if x, ok := m.GetPayload().(*RequestStream_Ack); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RequestStream) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RequestStream_OneofMarshaler, _RequestStream_OneofUnmarshaler, _RequestStream_OneofSizer, []interface{}{
		(*RequestStream_Request)(nil),
		(*RequestStream_Ack)(nil),
	}
}

func _RequestStream_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RequestStream)
	// payload
	switch x := m.Payload.(type) {
	case *RequestStream_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *RequestStream_Ack:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RequestStream.Payload has unexpected type %T", x)
	}
	return nil
}

func _RequestStream_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RequestStream)
	switch tag {
	case 1: // payload.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QueryRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &RequestStream_Request{msg}
		return true, err
	case 2: // payload.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AckMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &RequestStream_Ack{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RequestStream_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RequestStream)
	// payload
	switch x := m.Payload.(type) {
	case *RequestStream_Request:
		s := proto.Size(x.Request)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestStream_Ack:
		s := proto.Size(x.Ack)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResponseStreamRecord struct {
	SequenceId           int64                                 `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId" json:"sequence_id,omitempty"`
	Type                 ResponseStreamRecord_StreamRecordType `protobuf:"varint,2,opt,name=type,enum=qpu.ResponseStreamRecord_StreamRecordType" json:"type,omitempty"`
	LogOp                *utils.LogOperation                   `protobuf:"bytes,3,opt,name=logOp" json:"logOp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ResponseStreamRecord) Reset()         { *m = ResponseStreamRecord{} }
func (m *ResponseStreamRecord) String() string { return proto.CompactTextString(m) }
func (*ResponseStreamRecord) ProtoMessage()    {}
func (*ResponseStreamRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{1}
}
func (m *ResponseStreamRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseStreamRecord.Unmarshal(m, b)
}
func (m *ResponseStreamRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseStreamRecord.Marshal(b, m, deterministic)
}
func (dst *ResponseStreamRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseStreamRecord.Merge(dst, src)
}
func (m *ResponseStreamRecord) XXX_Size() int {
	return xxx_messageInfo_ResponseStreamRecord.Size(m)
}
func (m *ResponseStreamRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseStreamRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseStreamRecord proto.InternalMessageInfo

func (m *ResponseStreamRecord) GetSequenceId() int64 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *ResponseStreamRecord) GetType() ResponseStreamRecord_StreamRecordType {
	if m != nil {
		return m.Type
	}
	return ResponseStreamRecord_UPDATEOP
}

func (m *ResponseStreamRecord) GetLogOp() *utils.LogOperation {
	if m != nil {
		return m.LogOp
	}
	return nil
}

type QueryRequest struct {
	// Timestamp is part of the attributes
	// Desclared explicitly here for easier parsing
	Clock                *utils.SnapshotTimePredicate `protobuf:"bytes,1,opt,name=clock" json:"clock,omitempty"`
	Predicate            []*utils.AttributePredicate  `protobuf:"bytes,2,rep,name=predicate" json:"predicate,omitempty"`
	Ops                  bool                         `protobuf:"varint,3,opt,name=ops" json:"ops,omitempty"`
	Sync                 bool                         `protobuf:"varint,4,opt,name=sync" json:"sync,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{2}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetClock() *utils.SnapshotTimePredicate {
	if m != nil {
		return m.Clock
	}
	return nil
}

func (m *QueryRequest) GetPredicate() []*utils.AttributePredicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *QueryRequest) GetOps() bool {
	if m != nil {
		return m.Ops
	}
	return false
}

func (m *QueryRequest) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type AckMsg struct {
	SequenceId           int64    `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId" json:"sequence_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckMsg) Reset()         { *m = AckMsg{} }
func (m *AckMsg) String() string { return proto.CompactTextString(m) }
func (*AckMsg) ProtoMessage()    {}
func (*AckMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{3}
}
func (m *AckMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckMsg.Unmarshal(m, b)
}
func (m *AckMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckMsg.Marshal(b, m, deterministic)
}
func (dst *AckMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckMsg.Merge(dst, src)
}
func (m *AckMsg) XXX_Size() int {
	return xxx_messageInfo_AckMsg.Size(m)
}
func (m *AckMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_AckMsg.DiscardUnknown(m)
}

var xxx_messageInfo_AckMsg proto.InternalMessageInfo

func (m *AckMsg) GetSequenceId() int64 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

type ConfigRequest struct {
	Clock                *utils.SnapshotTimePredicate `protobuf:"bytes,1,opt,name=clock" json:"clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{4}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(dst, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetClock() *utils.SnapshotTimePredicate {
	if m != nil {
		return m.Clock
	}
	return nil
}

type ConfigResponse struct {
	QpuType              ConfigResponse_QPUType      `protobuf:"varint,1,opt,name=qpu_type,json=qpuType,enum=qpu.ConfigResponse_QPUType" json:"qpu_type,omitempty"`
	SupportedQueries     []*utils.AttributePredicate `protobuf:"bytes,2,rep,name=supportedQueries" json:"supportedQueries,omitempty"`
	Dataset              *DataSet                    `protobuf:"bytes,3,opt,name=dataset" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{5}
}
func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(dst, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetQpuType() ConfigResponse_QPUType {
	if m != nil {
		return m.QpuType
	}
	return ConfigResponse_DBDRIVER
}

func (m *ConfigResponse) GetSupportedQueries() []*utils.AttributePredicate {
	if m != nil {
		return m.SupportedQueries
	}
	return nil
}

func (m *ConfigResponse) GetDataset() *DataSet {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type DataSet struct {
	Databases            map[string]*DataSet_DB `protobuf:"bytes,1,rep,name=databases" json:"databases,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DataSet) Reset()         { *m = DataSet{} }
func (m *DataSet) String() string { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()    {}
func (*DataSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{6}
}
func (m *DataSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSet.Unmarshal(m, b)
}
func (m *DataSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSet.Marshal(b, m, deterministic)
}
func (dst *DataSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSet.Merge(dst, src)
}
func (m *DataSet) XXX_Size() int {
	return xxx_messageInfo_DataSet.Size(m)
}
func (m *DataSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSet.DiscardUnknown(m)
}

var xxx_messageInfo_DataSet proto.InternalMessageInfo

func (m *DataSet) GetDatabases() map[string]*DataSet_DB {
	if m != nil {
		return m.Databases
	}
	return nil
}

type DataSet_DB struct {
	Datacenters          map[string]*DataSet_DC `protobuf:"bytes,1,rep,name=datacenters" json:"datacenters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DataSet_DB) Reset()         { *m = DataSet_DB{} }
func (m *DataSet_DB) String() string { return proto.CompactTextString(m) }
func (*DataSet_DB) ProtoMessage()    {}
func (*DataSet_DB) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{6, 0}
}
func (m *DataSet_DB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSet_DB.Unmarshal(m, b)
}
func (m *DataSet_DB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSet_DB.Marshal(b, m, deterministic)
}
func (dst *DataSet_DB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSet_DB.Merge(dst, src)
}
func (m *DataSet_DB) XXX_Size() int {
	return xxx_messageInfo_DataSet_DB.Size(m)
}
func (m *DataSet_DB) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSet_DB.DiscardUnknown(m)
}

var xxx_messageInfo_DataSet_DB proto.InternalMessageInfo

func (m *DataSet_DB) GetDatacenters() map[string]*DataSet_DC {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

type DataSet_DC struct {
	Shards               []string `protobuf:"bytes,1,rep,name=shards" json:"shards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataSet_DC) Reset()         { *m = DataSet_DC{} }
func (m *DataSet_DC) String() string { return proto.CompactTextString(m) }
func (*DataSet_DC) ProtoMessage()    {}
func (*DataSet_DC) Descriptor() ([]byte, []int) {
	return fileDescriptor_qpu_f7dee21d5154f47b, []int{6, 1}
}
func (m *DataSet_DC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSet_DC.Unmarshal(m, b)
}
func (m *DataSet_DC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSet_DC.Marshal(b, m, deterministic)
}
func (dst *DataSet_DC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSet_DC.Merge(dst, src)
}
func (m *DataSet_DC) XXX_Size() int {
	return xxx_messageInfo_DataSet_DC.Size(m)
}
func (m *DataSet_DC) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSet_DC.DiscardUnknown(m)
}

var xxx_messageInfo_DataSet_DC proto.InternalMessageInfo

func (m *DataSet_DC) GetShards() []string {
	if m != nil {
		return m.Shards
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestStream)(nil), "qpu.RequestStream")
	proto.RegisterType((*ResponseStreamRecord)(nil), "qpu.ResponseStreamRecord")
	proto.RegisterType((*QueryRequest)(nil), "qpu.QueryRequest")
	proto.RegisterType((*AckMsg)(nil), "qpu.AckMsg")
	proto.RegisterType((*ConfigRequest)(nil), "qpu.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "qpu.ConfigResponse")
	proto.RegisterType((*DataSet)(nil), "qpu.DataSet")
	proto.RegisterMapType((map[string]*DataSet_DB)(nil), "qpu.DataSet.DatabasesEntry")
	proto.RegisterType((*DataSet_DB)(nil), "qpu.DataSet.DB")
	proto.RegisterMapType((map[string]*DataSet_DC)(nil), "qpu.DataSet.DB.DatacentersEntry")
	proto.RegisterType((*DataSet_DC)(nil), "qpu.DataSet.DC")
	proto.RegisterEnum("qpu.ResponseStreamRecord_StreamRecordType", ResponseStreamRecord_StreamRecordType_name, ResponseStreamRecord_StreamRecordType_value)
	proto.RegisterEnum("qpu.ConfigResponse_QPUType", ConfigResponse_QPUType_name, ConfigResponse_QPUType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for QPU service

type QPUClient interface {
	Query(ctx context.Context, opts ...grpc.CallOption) (QPU_QueryClient, error)
	GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type qPUClient struct {
	cc *grpc.ClientConn
}

func NewQPUClient(cc *grpc.ClientConn) QPUClient {
	return &qPUClient{cc}
}

func (c *qPUClient) Query(ctx context.Context, opts ...grpc.CallOption) (QPU_QueryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_QPU_serviceDesc.Streams[0], c.cc, "/qpu.QPU/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &qPUQueryClient{stream}
	return x, nil
}

type QPU_QueryClient interface {
	Send(*RequestStream) error
	Recv() (*ResponseStreamRecord, error)
	grpc.ClientStream
}

type qPUQueryClient struct {
	grpc.ClientStream
}

func (x *qPUQueryClient) Send(m *RequestStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *qPUQueryClient) Recv() (*ResponseStreamRecord, error) {
	m := new(ResponseStreamRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *qPUClient) GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/qpu.QPU/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QPU service

type QPUServer interface {
	Query(QPU_QueryServer) error
	GetConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

func RegisterQPUServer(s *grpc.Server, srv QPUServer) {
	s.RegisterService(&_QPU_serviceDesc, srv)
}

func _QPU_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QPUServer).Query(&qPUQueryServer{stream})
}

type QPU_QueryServer interface {
	Send(*ResponseStreamRecord) error
	Recv() (*RequestStream, error)
	grpc.ServerStream
}

type qPUQueryServer struct {
	grpc.ServerStream
}

func (x *qPUQueryServer) Send(m *ResponseStreamRecord) error {
	return x.ServerStream.SendMsg(m)
}

func (x *qPUQueryServer) Recv() (*RequestStream, error) {
	m := new(RequestStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _QPU_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPUServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qpu.QPU/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPUServer).GetConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QPU_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qpu.QPU",
	HandlerType: (*QPUServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _QPU_GetConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _QPU_Query_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "qpu.proto",
}

func init() { proto.RegisterFile("qpu.proto", fileDescriptor_qpu_f7dee21d5154f47b) }

var fileDescriptor_qpu_f7dee21d5154f47b = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x6d, 0x6b, 0xe3, 0x46,
	0x10, 0xb6, 0x24, 0xbf, 0x44, 0xe3, 0xbc, 0xa8, 0x9b, 0xb6, 0x38, 0x4e, 0x20, 0x41, 0xd0, 0xe0,
	0x14, 0xea, 0x14, 0x17, 0xd2, 0x17, 0x4a, 0x41, 0xb2, 0x94, 0xda, 0x34, 0x89, 0x9d, 0xb5, 0x52,
	0x4a, 0xbf, 0x84, 0xb5, 0xb4, 0x75, 0x54, 0x3b, 0xd2, 0x5a, 0xbb, 0x0a, 0x98, 0xfb, 0x0d, 0xf7,
	0x03, 0xee, 0xe3, 0xfd, 0xa0, 0xfb, 0x2d, 0xf7, 0x17, 0x0e, 0xad, 0x64, 0x62, 0xe7, 0x72, 0x84,
	0xe3, 0x3e, 0x69, 0x76, 0xf6, 0x99, 0x79, 0x66, 0xe7, 0x19, 0x0d, 0xe8, 0x73, 0x96, 0xb6, 0x59,
	0x12, 0x8b, 0x18, 0x69, 0x73, 0x96, 0x36, 0xeb, 0xa9, 0x08, 0x67, 0x3c, 0xf7, 0x98, 0xff, 0xc3,
	0x16, 0xa6, 0xf3, 0x94, 0x72, 0x31, 0x12, 0x09, 0x25, 0xf7, 0xe8, 0x07, 0xa8, 0x25, 0xb9, 0xa3,
	0xa1, 0x1c, 0x29, 0xad, 0x7a, 0xe7, 0xab, 0x76, 0x16, 0x7f, 0x9d, 0xd2, 0x64, 0x51, 0x20, 0x7b,
	0x25, 0xbc, 0xc4, 0xa0, 0x43, 0xd0, 0x88, 0x3f, 0x6d, 0xa8, 0x12, 0x5a, 0x97, 0x50, 0xcb, 0x9f,
	0x5e, 0xf2, 0x49, 0xaf, 0x84, 0xb3, 0x1b, 0x5b, 0x87, 0x1a, 0x23, 0x8b, 0x59, 0x4c, 0x02, 0xf3,
	0xbd, 0x02, 0x5f, 0x63, 0xca, 0x59, 0x1c, 0x71, 0x9a, 0xb3, 0x61, 0xea, 0xc7, 0x49, 0x80, 0x0e,
	0xa1, 0xce, 0xb3, 0x7c, 0x91, 0x4f, 0x6f, 0xc3, 0x40, 0xf2, 0x6a, 0x18, 0x96, 0xae, 0x7e, 0x80,
	0xfe, 0x80, 0xb2, 0x58, 0x30, 0x2a, 0x69, 0xb6, 0x3b, 0xdf, 0x4b, 0x9a, 0xe7, 0x32, 0xb5, 0x57,
	0x0f, 0xde, 0x82, 0x51, 0x2c, 0xe3, 0xd0, 0x09, 0x54, 0x66, 0xf1, 0x64, 0xc0, 0x1a, 0x9a, 0xac,
	0x73, 0xb7, 0x9d, 0xb7, 0xe0, 0x22, 0xf3, 0xd1, 0x84, 0x88, 0x30, 0x8e, 0x70, 0x8e, 0x30, 0xff,
	0x02, 0xe3, 0x69, 0x12, 0xb4, 0x09, 0x1b, 0x37, 0x43, 0xc7, 0xf2, 0xdc, 0xc1, 0xd0, 0x28, 0xa1,
	0x1d, 0xa8, 0xe7, 0x27, 0xc7, 0xbd, 0xf0, 0x2c, 0x43, 0x41, 0x3a, 0x54, 0x46, 0x9e, 0xe5, 0xb9,
	0x86, 0x8a, 0xb6, 0x40, 0xef, 0xb9, 0x16, 0xf6, 0x6c, 0xd7, 0xf2, 0x0c, 0xcd, 0x7c, 0xab, 0xc0,
	0xe6, 0x6a, 0xe7, 0x50, 0x07, 0x2a, 0xfe, 0x2c, 0xf6, 0xa7, 0x45, 0x6f, 0x0f, 0x8a, 0x42, 0x46,
	0x11, 0x61, 0xfc, 0x2e, 0x16, 0x5e, 0x78, 0x4f, 0x87, 0x09, 0x0d, 0x42, 0x9f, 0x08, 0x8a, 0x73,
	0x28, 0xfa, 0x19, 0x74, 0xb6, 0xf4, 0x35, 0xd4, 0x23, 0xad, 0x55, 0xef, 0xec, 0x15, 0x71, 0x96,
	0x10, 0x49, 0x38, 0x4e, 0xc5, 0x4a, 0xd0, 0x23, 0x16, 0x19, 0xa0, 0xc5, 0x8c, 0xcb, 0x37, 0x6f,
	0xe0, 0xcc, 0x44, 0x08, 0xca, 0x7c, 0x11, 0xf9, 0x8d, 0xb2, 0x74, 0x49, 0xdb, 0x3c, 0x81, 0x6a,
	0xae, 0xd8, 0x8b, 0x32, 0x98, 0x5d, 0xd8, 0xea, 0xc6, 0xd1, 0x7f, 0xe1, 0xe4, 0x0b, 0x9e, 0x63,
	0xbe, 0x56, 0x61, 0x7b, 0x99, 0x25, 0x57, 0x10, 0x9d, 0xc1, 0xc6, 0x9c, 0xa5, 0xb7, 0x52, 0x62,
	0x45, 0x4a, 0xbc, 0x2f, 0x25, 0x5e, 0x87, 0xb5, 0xaf, 0x87, 0x37, 0x52, 0xd3, 0xda, 0x9c, 0xa5,
	0x52, 0x17, 0x17, 0x0c, 0x9e, 0x32, 0x16, 0x27, 0x82, 0x06, 0x59, 0x9b, 0x43, 0xca, 0x5f, 0x6e,
	0xd0, 0x47, 0x21, 0xe8, 0x18, 0x6a, 0x01, 0x11, 0x84, 0x53, 0x51, 0xcc, 0xc7, 0xa6, 0x64, 0x77,
	0x88, 0x20, 0x23, 0x2a, 0xf0, 0xf2, 0xd2, 0xf4, 0xa0, 0x56, 0x94, 0x90, 0x4d, 0x84, 0x63, 0x3b,
	0xb8, 0xff, 0xb7, 0x8b, 0x8d, 0x12, 0x02, 0xa8, 0x9e, 0xf7, 0x2f, 0x3c, 0x17, 0xe7, 0xc3, 0xd0,
	0xbf, 0x72, 0xdc, 0x7f, 0x0c, 0x35, 0x33, 0xbb, 0x56, 0xb7, 0xe7, 0x1a, 0x1a, 0xda, 0x83, 0x6f,
	0xce, 0x5d, 0xc7, 0xc5, 0x96, 0xd7, 0x1f, 0x5c, 0xdd, 0x3a, 0xfd, 0xd1, 0xd0, 0xf2, 0xba, 0x3d,
	0x17, 0x1b, 0x65, 0xf3, 0x9d, 0x0a, 0xb5, 0x82, 0x0a, 0xfd, 0x0a, 0x7a, 0x46, 0x36, 0x26, 0x9c,
	0xf2, 0x86, 0x22, 0x5f, 0xb2, 0xbf, 0x5a, 0x8b, 0xfc, 0xca, 0x5b, 0x37, 0x12, 0xc9, 0x02, 0x3f,
	0xa2, 0x9b, 0x6f, 0x14, 0x50, 0x1d, 0x1b, 0xd9, 0x50, 0xcf, 0x7c, 0x3e, 0x8d, 0x04, 0x4d, 0x96,
	0x39, 0x8e, 0xd6, 0x73, 0xd8, 0xd2, 0x2c, 0x20, 0x79, 0xa2, 0xd5, 0xa0, 0xe6, 0x00, 0x8c, 0xa7,
	0x80, 0x6c, 0x96, 0xa6, 0x74, 0x21, 0xd5, 0xd1, 0x71, 0x66, 0xa2, 0xef, 0xa0, 0xf2, 0x40, 0x66,
	0x29, 0x2d, 0xfe, 0xfd, 0x9d, 0x75, 0x8e, 0x2e, 0xce, 0x6f, 0x7f, 0x53, 0x7f, 0x51, 0x9a, 0x07,
	0xa0, 0x3a, 0x5d, 0xf4, 0x2d, 0x54, 0xf9, 0x1d, 0x49, 0x82, 0xbc, 0x2a, 0x1d, 0x17, 0xa7, 0xe6,
	0x25, 0x6c, 0xaf, 0x3f, 0xeb, 0xb3, 0xc9, 0xec, 0x15, 0xb2, 0xce, 0x2b, 0xd0, 0xae, 0x87, 0x37,
	0xe8, 0x77, 0xa8, 0xc8, 0x3f, 0x0f, 0xa1, 0x62, 0x5b, 0xac, 0x2c, 0xb9, 0xe6, 0xde, 0x27, 0x37,
	0x88, 0x59, 0x6a, 0x29, 0x3f, 0x2a, 0xe8, 0x0c, 0xf4, 0x3f, 0xa9, 0xc8, 0xe7, 0xaf, 0xc8, 0xb0,
	0x36, 0xf9, 0xcd, 0xdd, 0x67, 0x06, 0xd4, 0x2c, 0xd9, 0xad, 0x7f, 0x8f, 0x27, 0xa1, 0xb8, 0x4b,
	0xc7, 0x6d, 0x3f, 0xbe, 0x3f, 0x0d, 0x1e, 0x08, 0x0f, 0x67, 0x84, 0x9f, 0x66, 0xbb, 0x96, 0xa6,
	0xf9, 0x37, 0xe6, 0xa7, 0x73, 0x96, 0x8e, 0xab, 0xd2, 0xfe, 0xe9, 0x43, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x89, 0xa6, 0xfd, 0x60, 0x9e, 0x05, 0x00, 0x00,
}
