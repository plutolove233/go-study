// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: services/hello/helloService.proto

package hello

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	models "grpc-demo/pb/models"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_hello_helloService_proto protoreflect.FileDescriptor

var file_services_hello_helloService_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x44, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x61, 0x79, 0x48, 0x69,
	0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x44, 0x0a,
	0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x34, 0x0a,
	0x05, 0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_services_hello_helloService_proto_goTypes = []interface{}{
	(*models.HelloRequest)(nil),  // 0: hello.HelloRequest
	(*models.HelloResponse)(nil), // 1: hello.HelloResponse
}
var file_services_hello_helloService_proto_depIdxs = []int32{
	0, // 0: ServerStream.SayHi:input_type -> hello.HelloRequest
	0, // 1: ClientStream.SayHi:input_type -> hello.HelloRequest
	1, // 2: ServerStream.SayHi:output_type -> hello.HelloResponse
	1, // 3: ClientStream.SayHi:output_type -> hello.HelloResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_hello_helloService_proto_init() }
func file_services_hello_helloService_proto_init() {
	if File_services_hello_helloService_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_hello_helloService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_services_hello_helloService_proto_goTypes,
		DependencyIndexes: file_services_hello_helloService_proto_depIdxs,
	}.Build()
	File_services_hello_helloService_proto = out.File
	file_services_hello_helloService_proto_rawDesc = nil
	file_services_hello_helloService_proto_goTypes = nil
	file_services_hello_helloService_proto_depIdxs = nil
}