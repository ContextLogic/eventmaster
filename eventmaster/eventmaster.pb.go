// Code generated by protoc-gen-go.
// source: eventmaster.proto
// DO NOT EDIT!

/*
Package eventmaster is a generated protocol buffer package.

It is generated from these files:
	eventmaster.proto

It has these top-level messages:
	Event
	Query
	Topic
	UpdateTopicRequest
	DeleteTopicRequest
	Dc
	UpdateDcRequest
	WriteResponse
*/
package eventmaster

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

type Event struct {
	EventId       string   `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	ParentEventId string   `protobuf:"bytes,2,opt,name=parent_event_id,json=parentEventId" json:"parent_event_id,omitempty"`
	EventTime     int64    `protobuf:"varint,3,opt,name=event_time,json=eventTime" json:"event_time,omitempty"`
	Dc            string   `protobuf:"bytes,4,opt,name=dc" json:"dc,omitempty"`
	TopicName     string   `protobuf:"bytes,5,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	TagSet        []string `protobuf:"bytes,6,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	Host          string   `protobuf:"bytes,7,opt,name=host" json:"host,omitempty"`
	TargetHostSet []string `protobuf:"bytes,8,rep,name=target_host_set,json=targetHostSet" json:"target_host_set,omitempty"`
	User          string   `protobuf:"bytes,9,opt,name=user" json:"user,omitempty"`
	Data          string   `protobuf:"bytes,10,opt,name=data" json:"data,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetParentEventId() string {
	if m != nil {
		return m.ParentEventId
	}
	return ""
}

func (m *Event) GetEventTime() int64 {
	if m != nil {
		return m.EventTime
	}
	return 0
}

func (m *Event) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *Event) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *Event) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *Event) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Event) GetTargetHostSet() []string {
	if m != nil {
		return m.TargetHostSet
	}
	return nil
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Query struct {
	Dc                []string `protobuf:"bytes,1,rep,name=dc" json:"dc,omitempty"`
	Host              []string `protobuf:"bytes,2,rep,name=host" json:"host,omitempty"`
	TargetHostSet     []string `protobuf:"bytes,3,rep,name=target_host_set,json=targetHostSet" json:"target_host_set,omitempty"`
	TopicName         []string `protobuf:"bytes,4,rep,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	TagSet            []string `protobuf:"bytes,5,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	Data              string   `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	StartEventTime    int64    `protobuf:"varint,7,opt,name=start_event_time,json=startEventTime" json:"start_event_time,omitempty"`
	EndEventTime      int64    `protobuf:"varint,8,opt,name=end_event_time,json=endEventTime" json:"end_event_time,omitempty"`
	StartReceivedTime int64    `protobuf:"varint,9,opt,name=start_received_time,json=startReceivedTime" json:"start_received_time,omitempty"`
	EndReceivedTime   int64    `protobuf:"varint,10,opt,name=end_received_time,json=endReceivedTime" json:"end_received_time,omitempty"`
	SortField         []string `protobuf:"bytes,11,rep,name=sort_field,json=sortField" json:"sort_field,omitempty"`
	SortAscending     []bool   `protobuf:"varint,12,rep,packed,name=sort_ascending,json=sortAscending" json:"sort_ascending,omitempty"`
	Start             int32    `protobuf:"varint,13,opt,name=start" json:"start,omitempty"`
	Limit             int32    `protobuf:"varint,14,opt,name=limit" json:"limit,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Query) GetDc() []string {
	if m != nil {
		return m.Dc
	}
	return nil
}

func (m *Query) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Query) GetTargetHostSet() []string {
	if m != nil {
		return m.TargetHostSet
	}
	return nil
}

func (m *Query) GetTopicName() []string {
	if m != nil {
		return m.TopicName
	}
	return nil
}

func (m *Query) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *Query) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Query) GetStartEventTime() int64 {
	if m != nil {
		return m.StartEventTime
	}
	return 0
}

func (m *Query) GetEndEventTime() int64 {
	if m != nil {
		return m.EndEventTime
	}
	return 0
}

func (m *Query) GetStartReceivedTime() int64 {
	if m != nil {
		return m.StartReceivedTime
	}
	return 0
}

func (m *Query) GetEndReceivedTime() int64 {
	if m != nil {
		return m.EndReceivedTime
	}
	return 0
}

func (m *Query) GetSortField() []string {
	if m != nil {
		return m.SortField
	}
	return nil
}

func (m *Query) GetSortAscending() []bool {
	if m != nil {
		return m.SortAscending
	}
	return nil
}

func (m *Query) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Query) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Topic struct {
	TopicName  string `protobuf:"bytes,1,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	DataSchema string `protobuf:"bytes,2,opt,name=data_schema,json=dataSchema" json:"data_schema,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Topic) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *Topic) GetDataSchema() string {
	if m != nil {
		return m.DataSchema
	}
	return ""
}

type UpdateTopicRequest struct {
	OldName   string `protobuf:"bytes,1,opt,name=old_name,json=oldName" json:"old_name,omitempty"`
	NewName   string `protobuf:"bytes,2,opt,name=new_name,json=newName" json:"new_name,omitempty"`
	NewSchema string `protobuf:"bytes,3,opt,name=new_schema,json=newSchema" json:"new_schema,omitempty"`
}

func (m *UpdateTopicRequest) Reset()                    { *m = UpdateTopicRequest{} }
func (m *UpdateTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTopicRequest) ProtoMessage()               {}
func (*UpdateTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateTopicRequest) GetOldName() string {
	if m != nil {
		return m.OldName
	}
	return ""
}

func (m *UpdateTopicRequest) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

func (m *UpdateTopicRequest) GetNewSchema() string {
	if m != nil {
		return m.NewSchema
	}
	return ""
}

type DeleteTopicRequest struct {
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteTopicRequest) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

type Dc struct {
	Dc string `protobuf:"bytes,1,opt,name=dc" json:"dc,omitempty"`
}

func (m *Dc) Reset()                    { *m = Dc{} }
func (m *Dc) String() string            { return proto.CompactTextString(m) }
func (*Dc) ProtoMessage()               {}
func (*Dc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Dc) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

type UpdateDcRequest struct {
	OldDc string `protobuf:"bytes,1,opt,name=old_dc,json=oldDc" json:"old_dc,omitempty"`
	NewDc string `protobuf:"bytes,2,opt,name=new_dc,json=newDc" json:"new_dc,omitempty"`
}

func (m *UpdateDcRequest) Reset()                    { *m = UpdateDcRequest{} }
func (m *UpdateDcRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDcRequest) ProtoMessage()               {}
func (*UpdateDcRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateDcRequest) GetOldDc() string {
	if m != nil {
		return m.OldDc
	}
	return ""
}

func (m *UpdateDcRequest) GetNewDc() string {
	if m != nil {
		return m.NewDc
	}
	return ""
}

type WriteResponse struct {
	Errcode int32  `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Id      string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WriteResponse) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *WriteResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *WriteResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "eventmaster.Event")
	proto.RegisterType((*Query)(nil), "eventmaster.Query")
	proto.RegisterType((*Topic)(nil), "eventmaster.Topic")
	proto.RegisterType((*UpdateTopicRequest)(nil), "eventmaster.UpdateTopicRequest")
	proto.RegisterType((*DeleteTopicRequest)(nil), "eventmaster.DeleteTopicRequest")
	proto.RegisterType((*Dc)(nil), "eventmaster.Dc")
	proto.RegisterType((*UpdateDcRequest)(nil), "eventmaster.UpdateDcRequest")
	proto.RegisterType((*WriteResponse)(nil), "eventmaster.WriteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventMaster service

type EventMasterClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*WriteResponse, error)
	GetEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (EventMaster_GetEventsClient, error)
	AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*WriteResponse, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	AddDc(ctx context.Context, in *Dc, opts ...grpc.CallOption) (*WriteResponse, error)
	UpdateDc(ctx context.Context, in *UpdateDcRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type eventMasterClient struct {
	cc *grpc.ClientConn
}

func NewEventMasterClient(cc *grpc.ClientConn) EventMasterClient {
	return &eventMasterClient{cc}
}

func (c *eventMasterClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/AddEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) GetEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (EventMaster_GetEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventMaster_serviceDesc.Streams[0], c.cc, "/eventmaster.EventMaster/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventMasterGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventMaster_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventMasterGetEventsClient struct {
	grpc.ClientStream
}

func (x *eventMasterGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventMasterClient) AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/AddTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/UpdateTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/DeleteTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) AddDc(ctx context.Context, in *Dc, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/AddDc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) UpdateDc(ctx context.Context, in *UpdateDcRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/UpdateDc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventMaster service

type EventMasterServer interface {
	AddEvent(context.Context, *Event) (*WriteResponse, error)
	GetEvents(*Query, EventMaster_GetEventsServer) error
	AddTopic(context.Context, *Topic) (*WriteResponse, error)
	UpdateTopic(context.Context, *UpdateTopicRequest) (*WriteResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*WriteResponse, error)
	AddDc(context.Context, *Dc) (*WriteResponse, error)
	UpdateDc(context.Context, *UpdateDcRequest) (*WriteResponse, error)
}

func RegisterEventMasterServer(s *grpc.Server, srv EventMasterServer) {
	s.RegisterService(&_EventMaster_serviceDesc, srv)
}

func _EventMaster_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventMasterServer).GetEvents(m, &eventMasterGetEventsServer{stream})
}

type EventMaster_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventMasterGetEventsServer struct {
	grpc.ServerStream
}

func (x *eventMasterGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventMaster_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/AddTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).AddTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_AddDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).AddDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/AddDc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).AddDc(ctx, req.(*Dc))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_UpdateDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).UpdateDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/UpdateDc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).UpdateDc(ctx, req.(*UpdateDcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventmaster.EventMaster",
	HandlerType: (*EventMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _EventMaster_AddEvent_Handler,
		},
		{
			MethodName: "AddTopic",
			Handler:    _EventMaster_AddTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _EventMaster_UpdateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _EventMaster_DeleteTopic_Handler,
		},
		{
			MethodName: "AddDc",
			Handler:    _EventMaster_AddDc_Handler,
		},
		{
			MethodName: "UpdateDc",
			Handler:    _EventMaster_UpdateDc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _EventMaster_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eventmaster.proto",
}

func init() { proto.RegisterFile("eventmaster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0x6d, 0xe2, 0x3a, 0x4e, 0x26, 0x4d, 0x72, 0xbb, 0xb7, 0xc0, 0x52, 0x81, 0x5a, 0x59, 0x80,
	0x2a, 0x1e, 0x2a, 0x44, 0x25, 0x78, 0x41, 0x42, 0x95, 0xd2, 0x16, 0x24, 0x40, 0xaa, 0x5b, 0xc4,
	0x63, 0x64, 0xec, 0x21, 0xb5, 0x94, 0xd8, 0xc1, 0xbb, 0x29, 0xe2, 0x43, 0xf8, 0x1c, 0x7e, 0x85,
	0x6f, 0x61, 0x76, 0xd6, 0x49, 0x6d, 0x37, 0x10, 0xde, 0x76, 0xce, 0x9c, 0x39, 0x73, 0x66, 0x3d,
	0xb6, 0x61, 0x1b, 0xaf, 0x31, 0xd5, 0xd3, 0x50, 0x69, 0xcc, 0x0f, 0x67, 0x79, 0xa6, 0x33, 0xd1,
	0x2d, 0x41, 0xfe, 0x8f, 0x26, 0xb8, 0x27, 0x26, 0x16, 0xf7, 0xa1, 0xcd, 0x89, 0x51, 0x12, 0xcb,
	0xc6, 0x7e, 0xe3, 0xa0, 0x13, 0x78, 0x1c, 0xbf, 0x8d, 0xc5, 0x13, 0x18, 0xcc, 0xc2, 0xdc, 0xe4,
	0x96, 0x8c, 0x26, 0x33, 0x7a, 0x16, 0x3e, 0x29, 0x78, 0x0f, 0x01, 0x2c, 0x41, 0x27, 0x53, 0x94,
	0x0e, 0x51, 0x9c, 0xa0, 0xc3, 0xc8, 0x25, 0x01, 0xa2, 0x0f, 0xcd, 0x38, 0x92, 0x9b, 0x5c, 0x49,
	0x27, 0x43, 0xd7, 0xd9, 0x2c, 0x89, 0x46, 0x69, 0x48, 0x74, 0x97, 0xf1, 0x0e, 0x23, 0x1f, 0x08,
	0x10, 0xf7, 0xc0, 0xd3, 0xe1, 0x78, 0xa4, 0x50, 0xcb, 0xd6, 0xbe, 0x43, 0xb9, 0x16, 0x85, 0x17,
	0xa8, 0x85, 0x80, 0xcd, 0xab, 0x4c, 0x69, 0xe9, 0x71, 0x05, 0x9f, 0x8d, 0x45, 0x1d, 0xe6, 0x63,
	0xd4, 0x23, 0x13, 0x72, 0x51, 0x9b, 0x8b, 0x7a, 0x16, 0x7e, 0x43, 0x68, 0x51, 0x3b, 0x57, 0x98,
	0xcb, 0x8e, 0xad, 0x35, 0x67, 0x83, 0xc5, 0xa1, 0x0e, 0x25, 0x58, 0xcc, 0x9c, 0xfd, 0x9f, 0x0e,
	0xb8, 0xe7, 0x73, 0xcc, 0xbf, 0x17, 0xae, 0x1b, 0x2c, 0x66, 0x5c, 0x2f, 0xba, 0x37, 0x19, 0xf9,
	0x63, 0x77, 0x67, 0x55, 0xf7, 0xea, 0xc4, 0x9b, 0x4c, 0x59, 0x3d, 0xb1, 0x5b, 0x9f, 0x98, 0x1d,
	0xb6, 0x6e, 0x1c, 0x8a, 0x03, 0xf8, 0x4f, 0x91, 0xfa, 0xe2, 0x99, 0xf0, 0x95, 0x7b, 0x7c, 0xe5,
	0x7d, 0xc6, 0x4f, 0x96, 0xf7, 0xfe, 0x08, 0xfa, 0x98, 0xc6, 0x65, 0x5e, 0x9b, 0x79, 0x5b, 0x84,
	0xde, 0xb0, 0x0e, 0xe1, 0x7f, 0xab, 0x97, 0x63, 0x84, 0xc9, 0x35, 0xc6, 0x96, 0xda, 0x61, 0xea,
	0x36, 0xa7, 0x82, 0x22, 0xc3, 0xfc, 0xa7, 0xb4, 0x5b, 0xa4, 0x5a, 0x65, 0x03, 0xb3, 0x07, 0x94,
	0xa8, 0x70, 0x69, 0x6e, 0x95, 0x91, 0xf4, 0x97, 0x04, 0x27, 0xb1, 0xec, 0xda, 0xb9, 0x0d, 0x72,
	0x6a, 0x00, 0xf1, 0x18, 0xfa, 0x9c, 0x0e, 0x55, 0x44, 0x95, 0x49, 0x3a, 0x96, 0x5b, 0x44, 0x69,
	0x07, 0x3d, 0x83, 0x1e, 0x2f, 0x40, 0xb1, 0x03, 0x2e, 0xdb, 0x90, 0x3d, 0xea, 0xe2, 0x06, 0x36,
	0x30, 0xe8, 0x24, 0x99, 0x26, 0x5a, 0xf6, 0x2d, 0xca, 0x81, 0x7f, 0x06, 0xee, 0xa5, 0xb9, 0xd7,
	0xda, 0x95, 0x37, 0xea, 0x4b, 0xb6, 0x07, 0x5d, 0x73, 0x9b, 0x23, 0x15, 0x5d, 0xe1, 0x34, 0x2c,
	0xd6, 0x1a, 0x0c, 0x74, 0xc1, 0x88, 0x9f, 0x80, 0xf8, 0x38, 0xa3, 0x18, 0x59, 0x2e, 0xc0, 0xaf,
	0x73, 0x54, 0xfc, 0xb2, 0x64, 0x93, 0xb8, 0xac, 0xe9, 0x51, 0xcc, 0x8a, 0x94, 0x4a, 0xf1, 0x9b,
	0x4d, 0x59, 0x39, 0x8f, 0x62, 0x4e, 0x91, 0x17, 0x93, 0x2a, 0x7a, 0x39, 0xd6, 0x0b, 0x21, 0x45,
	0xab, 0x23, 0x10, 0x43, 0x9c, 0x60, 0xad, 0xd5, 0xdf, 0x07, 0xf0, 0x77, 0xa0, 0x39, 0x8c, 0x96,
	0x4b, 0x5a, 0xbc, 0x5a, 0xfe, 0x6b, 0x18, 0x58, 0xd7, 0xc3, 0xa5, 0xce, 0x1d, 0x68, 0x19, 0xcb,
	0x4b, 0x9a, 0x4b, 0x11, 0x55, 0x12, 0x6c, 0x3c, 0x11, 0x6c, 0xcd, 0xba, 0x14, 0x0d, 0x23, 0xff,
	0x1c, 0x7a, 0x9f, 0xf2, 0x44, 0x63, 0x80, 0x6a, 0x96, 0xa5, 0x0a, 0x85, 0x04, 0x0f, 0xf3, 0x3c,
	0xca, 0x62, 0xeb, 0xc1, 0x0d, 0x16, 0xa1, 0xb8, 0x0b, 0x2d, 0x3a, 0x4e, 0xd5, 0xb8, 0x50, 0x28,
	0x22, 0xe3, 0x89, 0x3e, 0x14, 0x76, 0x4a, 0x3a, 0x3d, 0xff, 0xe5, 0x40, 0x97, 0xd7, 0xed, 0x3d,
	0x7f, 0x7a, 0xc4, 0x2b, 0x68, 0x1f, 0xc7, 0x76, 0x01, 0x85, 0x38, 0x2c, 0x7f, 0xa7, 0x18, 0xdb,
	0xdd, 0xad, 0x60, 0x15, 0x37, 0xfe, 0x86, 0x78, 0x09, 0x9d, 0x33, 0xb4, 0x4b, 0xae, 0x6a, 0xe5,
	0xfc, 0xde, 0xee, 0xae, 0x90, 0xf4, 0x37, 0x9e, 0x35, 0x8a, 0xb6, 0x76, 0x39, 0xaa, 0x1c, 0xc6,
	0xd6, 0xb4, 0x7d, 0x07, 0xdd, 0xd2, 0x3a, 0x88, 0xbd, 0x0a, 0xf9, 0xf6, 0xa2, 0xac, 0x57, 0x2b,
	0x3d, 0xf1, 0x9a, 0xda, 0xed, 0x5d, 0x58, 0xa3, 0xf6, 0x02, 0x5c, 0x9a, 0x8c, 0x9e, 0xe9, 0xa0,
	0xaa, 0xb3, 0x6e, 0xa6, 0x53, 0x68, 0x2f, 0x96, 0x45, 0x3c, 0x58, 0x31, 0xd0, 0xf0, 0xdf, 0xfa,
	0x7f, 0x6e, 0xf1, 0xff, 0xe5, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x09, 0x0d, 0xb1,
	0x74, 0x06, 0x00, 0x00,
}
