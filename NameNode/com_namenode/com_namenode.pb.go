// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: com_namenode.proto

package com_namenode

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titulo string `protobuf:"bytes,1,opt,name=Titulo,proto3" json:"Titulo,omitempty"`
	Texto  string `protobuf:"bytes,2,opt,name=Texto,proto3" json:"Texto,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_namenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_com_namenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_com_namenode_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetTitulo() string {
	if x != nil {
		return x.Titulo
	}
	return ""
}

func (x *Log) GetTexto() string {
	if x != nil {
		return x.Texto
	}
	return ""
}

type EstadoEscritura struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado bool   `protobuf:"varint,1,opt,name=Estado,proto3" json:"Estado,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *EstadoEscritura) Reset() {
	*x = EstadoEscritura{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_namenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstadoEscritura) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoEscritura) ProtoMessage() {}

func (x *EstadoEscritura) ProtoReflect() protoreflect.Message {
	mi := &file_com_namenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoEscritura.ProtoReflect.Descriptor instead.
func (*EstadoEscritura) Descriptor() ([]byte, []int) {
	return file_com_namenode_proto_rawDescGZIP(), []int{1}
}

func (x *EstadoEscritura) GetEstado() bool {
	if x != nil {
		return x.Estado
	}
	return false
}

func (x *EstadoEscritura) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Disponibilidad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disponible bool `protobuf:"varint,1,opt,name=Disponible,proto3" json:"Disponible,omitempty"`
}

func (x *Disponibilidad) Reset() {
	*x = Disponibilidad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_namenode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disponibilidad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disponibilidad) ProtoMessage() {}

func (x *Disponibilidad) ProtoReflect() protoreflect.Message {
	mi := &file_com_namenode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disponibilidad.ProtoReflect.Descriptor instead.
func (*Disponibilidad) Descriptor() ([]byte, []int) {
	return file_com_namenode_proto_rawDescGZIP(), []int{2}
}

func (x *Disponibilidad) GetDisponible() bool {
	if x != nil {
		return x.Disponible
	}
	return false
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_namenode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_com_namenode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_com_namenode_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_com_namenode_proto protoreflect.FileDescriptor

var file_com_namenode_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f,
	0x64, 0x65, 0x22, 0x33, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x74,
	0x75, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x69, 0x74, 0x75, 0x6c,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x65, 0x78, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x65, 0x78, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x45, 0x73, 0x63, 0x72, 0x69, 0x74, 0x75, 0x72, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x6e,
	0x69, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x44, 0x69, 0x73, 0x70,
	0x6f, 0x6e, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x32, 0x8b, 0x01, 0x0a,
	0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x3f,
	0x0a, 0x0b, 0x45, 0x73, 0x63, 0x72, 0x69, 0x62, 0x69, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x6f, 0x67,
	0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x45, 0x73, 0x63, 0x72, 0x69, 0x74, 0x75, 0x72, 0x61, 0x12,
	0x39, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x64, 0x1a, 0x1c, 0x2e, 0x63,
	0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70,
	0x6f, 0x6e, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b,
	0x63, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_namenode_proto_rawDescOnce sync.Once
	file_com_namenode_proto_rawDescData = file_com_namenode_proto_rawDesc
)

func file_com_namenode_proto_rawDescGZIP() []byte {
	file_com_namenode_proto_rawDescOnce.Do(func() {
		file_com_namenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_namenode_proto_rawDescData)
	})
	return file_com_namenode_proto_rawDescData
}

var file_com_namenode_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_namenode_proto_goTypes = []interface{}{
	(*Log)(nil),             // 0: com_namenode.Log
	(*EstadoEscritura)(nil), // 1: com_namenode.EstadoEscritura
	(*Disponibilidad)(nil),  // 2: com_namenode.Disponibilidad
	(*Id)(nil),              // 3: com_namenode.Id
}
var file_com_namenode_proto_depIdxs = []int32{
	0, // 0: com_namenode.Interacciones.EscribirLog:input_type -> com_namenode.Log
	3, // 1: com_namenode.Interacciones.Request:input_type -> com_namenode.Id
	1, // 2: com_namenode.Interacciones.EscribirLog:output_type -> com_namenode.EstadoEscritura
	2, // 3: com_namenode.Interacciones.Request:output_type -> com_namenode.Disponibilidad
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_namenode_proto_init() }
func file_com_namenode_proto_init() {
	if File_com_namenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_namenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_namenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstadoEscritura); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_namenode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disponibilidad); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_namenode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_namenode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_namenode_proto_goTypes,
		DependencyIndexes: file_com_namenode_proto_depIdxs,
		MessageInfos:      file_com_namenode_proto_msgTypes,
	}.Build()
	File_com_namenode_proto = out.File
	file_com_namenode_proto_rawDesc = nil
	file_com_namenode_proto_goTypes = nil
	file_com_namenode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InteraccionesClient is the client API for Interacciones service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InteraccionesClient interface {
	EscribirLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*EstadoEscritura, error)
	Request(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Disponibilidad, error)
}

type interaccionesClient struct {
	cc grpc.ClientConnInterface
}

func NewInteraccionesClient(cc grpc.ClientConnInterface) InteraccionesClient {
	return &interaccionesClient{cc}
}

func (c *interaccionesClient) EscribirLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*EstadoEscritura, error) {
	out := new(EstadoEscritura)
	err := c.cc.Invoke(ctx, "/com_namenode.Interacciones/EscribirLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interaccionesClient) Request(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Disponibilidad, error) {
	out := new(Disponibilidad)
	err := c.cc.Invoke(ctx, "/com_namenode.Interacciones/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteraccionesServer is the server API for Interacciones service.
type InteraccionesServer interface {
	EscribirLog(context.Context, *Log) (*EstadoEscritura, error)
	Request(context.Context, *Id) (*Disponibilidad, error)
}

// UnimplementedInteraccionesServer can be embedded to have forward compatible implementations.
type UnimplementedInteraccionesServer struct {
}

func (*UnimplementedInteraccionesServer) EscribirLog(context.Context, *Log) (*EstadoEscritura, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscribirLog not implemented")
}
func (*UnimplementedInteraccionesServer) Request(context.Context, *Id) (*Disponibilidad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

func RegisterInteraccionesServer(s *grpc.Server, srv InteraccionesServer) {
	s.RegisterService(&_Interacciones_serviceDesc, srv)
}

func _Interacciones_EscribirLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteraccionesServer).EscribirLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com_namenode.Interacciones/EscribirLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteraccionesServer).EscribirLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interacciones_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteraccionesServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com_namenode.Interacciones/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteraccionesServer).Request(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Interacciones_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com_namenode.Interacciones",
	HandlerType: (*InteraccionesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EscribirLog",
			Handler:    _Interacciones_EscribirLog_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _Interacciones_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com_namenode.proto",
}
