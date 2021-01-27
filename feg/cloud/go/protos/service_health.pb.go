//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: feg/protos/service_health.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protos "magma/orc8r/lib/go/protos"
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

type DisableMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisablePeriodSecs uint64 `protobuf:"varint,1,opt,name=disablePeriodSecs,proto3" json:"disablePeriodSecs,omitempty"`
}

func (x *DisableMessage) Reset() {
	*x = DisableMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_service_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableMessage) ProtoMessage() {}

func (x *DisableMessage) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_service_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableMessage.ProtoReflect.Descriptor instead.
func (*DisableMessage) Descriptor() ([]byte, []int) {
	return file_feg_protos_service_health_proto_rawDescGZIP(), []int{0}
}

func (x *DisableMessage) GetDisablePeriodSecs() uint64 {
	if x != nil {
		return x.DisablePeriodSecs
	}
	return 0
}

var File_feg_protos_service_health_proto protoreflect.FileDescriptor

var file_feg_protos_service_health_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x65, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x1a, 0x19, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x66, 0x65, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3e, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x73,
	0x32, 0xbd, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x39, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65,
	0x67, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x66, 0x65, 0x67, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feg_protos_service_health_proto_rawDescOnce sync.Once
	file_feg_protos_service_health_proto_rawDescData = file_feg_protos_service_health_proto_rawDesc
)

func file_feg_protos_service_health_proto_rawDescGZIP() []byte {
	file_feg_protos_service_health_proto_rawDescOnce.Do(func() {
		file_feg_protos_service_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_feg_protos_service_health_proto_rawDescData)
	})
	return file_feg_protos_service_health_proto_rawDescData
}

var file_feg_protos_service_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_feg_protos_service_health_proto_goTypes = []interface{}{
	(*DisableMessage)(nil), // 0: magma.feg.DisableMessage
	(*protos.Void)(nil),    // 1: magma.orc8r.Void
	(*HealthStatus)(nil),   // 2: magma.feg.HealthStatus
}
var file_feg_protos_service_health_proto_depIdxs = []int32{
	0, // 0: magma.feg.ServiceHealth.Disable:input_type -> magma.feg.DisableMessage
	1, // 1: magma.feg.ServiceHealth.Enable:input_type -> magma.orc8r.Void
	1, // 2: magma.feg.ServiceHealth.GetHealthStatus:input_type -> magma.orc8r.Void
	1, // 3: magma.feg.ServiceHealth.Disable:output_type -> magma.orc8r.Void
	1, // 4: magma.feg.ServiceHealth.Enable:output_type -> magma.orc8r.Void
	2, // 5: magma.feg.ServiceHealth.GetHealthStatus:output_type -> magma.feg.HealthStatus
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feg_protos_service_health_proto_init() }
func file_feg_protos_service_health_proto_init() {
	if File_feg_protos_service_health_proto != nil {
		return
	}
	file_feg_protos_health_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_feg_protos_service_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableMessage); i {
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
			RawDescriptor: file_feg_protos_service_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feg_protos_service_health_proto_goTypes,
		DependencyIndexes: file_feg_protos_service_health_proto_depIdxs,
		MessageInfos:      file_feg_protos_service_health_proto_msgTypes,
	}.Build()
	File_feg_protos_service_health_proto = out.File
	file_feg_protos_service_health_proto_rawDesc = nil
	file_feg_protos_service_health_proto_goTypes = nil
	file_feg_protos_service_health_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceHealthClient is the client API for ServiceHealth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceHealthClient interface {
	// Disable service functionality for the time specified in the request
	Disable(ctx context.Context, in *DisableMessage, opts ...grpc.CallOption) (*protos.Void, error)
	// Enable service functionality
	Enable(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Void, error)
	// Get health status of the service
	GetHealthStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*HealthStatus, error)
}

type serviceHealthClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceHealthClient(cc grpc.ClientConnInterface) ServiceHealthClient {
	return &serviceHealthClient{cc}
}

func (c *serviceHealthClient) Disable(ctx context.Context, in *DisableMessage, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.ServiceHealth/Disable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceHealthClient) Enable(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.ServiceHealth/Enable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceHealthClient) GetHealthStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*HealthStatus, error) {
	out := new(HealthStatus)
	err := c.cc.Invoke(ctx, "/magma.feg.ServiceHealth/GetHealthStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceHealthServer is the server API for ServiceHealth service.
type ServiceHealthServer interface {
	// Disable service functionality for the time specified in the request
	Disable(context.Context, *DisableMessage) (*protos.Void, error)
	// Enable service functionality
	Enable(context.Context, *protos.Void) (*protos.Void, error)
	// Get health status of the service
	GetHealthStatus(context.Context, *protos.Void) (*HealthStatus, error)
}

// UnimplementedServiceHealthServer can be embedded to have forward compatible implementations.
type UnimplementedServiceHealthServer struct {
}

func (*UnimplementedServiceHealthServer) Disable(context.Context, *DisableMessage) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (*UnimplementedServiceHealthServer) Enable(context.Context, *protos.Void) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (*UnimplementedServiceHealthServer) GetHealthStatus(context.Context, *protos.Void) (*HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthStatus not implemented")
}

func RegisterServiceHealthServer(s *grpc.Server, srv ServiceHealthServer) {
	s.RegisterService(&_ServiceHealth_serviceDesc, srv)
}

func _ServiceHealth_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceHealthServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.ServiceHealth/Disable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceHealthServer).Disable(ctx, req.(*DisableMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceHealth_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceHealthServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.ServiceHealth/Enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceHealthServer).Enable(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceHealth_GetHealthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceHealthServer).GetHealthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.ServiceHealth/GetHealthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceHealthServer).GetHealthStatus(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceHealth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.ServiceHealth",
	HandlerType: (*ServiceHealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Disable",
			Handler:    _ServiceHealth_Disable_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _ServiceHealth_Enable_Handler,
		},
		{
			MethodName: "GetHealthStatus",
			Handler:    _ServiceHealth_GetHealthStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/service_health.proto",
}
