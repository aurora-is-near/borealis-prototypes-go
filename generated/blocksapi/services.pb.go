// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: blocksapi/services.proto

package pb_blocksapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_blocksapi_services_proto protoreflect.FileDescriptor

var file_blocksapi_services_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6f, 0x72, 0x65,
	0x61, 0x6c, 0x69, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x1a, 0x1c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xd3, 0x02, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x6d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61,
	0x6c, 0x69, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69,
	0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x66, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x28, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x6f, 0x72,
	0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_blocksapi_services_proto_goTypes = []any{
	(*ListBlockStreamsRequest)(nil),  // 0: borealis.blocksapi.ListBlockStreamsRequest
	(*GetBlockMessageRequest)(nil),   // 1: borealis.blocksapi.GetBlockMessageRequest
	(*ReceiveBlocksRequest)(nil),     // 2: borealis.blocksapi.ReceiveBlocksRequest
	(*ListBlockStreamsResponse)(nil), // 3: borealis.blocksapi.ListBlockStreamsResponse
	(*GetBlockMessageResponse)(nil),  // 4: borealis.blocksapi.GetBlockMessageResponse
	(*ReceiveBlocksResponse)(nil),    // 5: borealis.blocksapi.ReceiveBlocksResponse
}
var file_blocksapi_services_proto_depIdxs = []int32{
	0, // 0: borealis.blocksapi.BlocksProvider.ListBlockStreams:input_type -> borealis.blocksapi.ListBlockStreamsRequest
	1, // 1: borealis.blocksapi.BlocksProvider.GetBlockMessage:input_type -> borealis.blocksapi.GetBlockMessageRequest
	2, // 2: borealis.blocksapi.BlocksProvider.ReceiveBlocks:input_type -> borealis.blocksapi.ReceiveBlocksRequest
	3, // 3: borealis.blocksapi.BlocksProvider.ListBlockStreams:output_type -> borealis.blocksapi.ListBlockStreamsResponse
	4, // 4: borealis.blocksapi.BlocksProvider.GetBlockMessage:output_type -> borealis.blocksapi.GetBlockMessageResponse
	5, // 5: borealis.blocksapi.BlocksProvider.ReceiveBlocks:output_type -> borealis.blocksapi.ReceiveBlocksResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blocksapi_services_proto_init() }
func file_blocksapi_services_proto_init() {
	if File_blocksapi_services_proto != nil {
		return
	}
	file_blocksapi_read_message_proto_init()
	file_blocksapi_stream_list_proto_init()
	file_blocksapi_stream_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocksapi_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blocksapi_services_proto_goTypes,
		DependencyIndexes: file_blocksapi_services_proto_depIdxs,
	}.Build()
	File_blocksapi_services_proto = out.File
	file_blocksapi_services_proto_rawDesc = nil
	file_blocksapi_services_proto_goTypes = nil
	file_blocksapi_services_proto_depIdxs = nil
}
