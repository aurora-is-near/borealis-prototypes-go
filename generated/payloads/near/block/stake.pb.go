// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: payloads/near/block/stake.proto

package pb_block

import (
	common "github.com/aurora-is-near/borealis-prototypes-go/generated/payloads/near/common"
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

type ValidatorStakeView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*ValidatorStakeView_V1
	Variant isValidatorStakeView_Variant `protobuf_oneof:"variant"`
}

func (x *ValidatorStakeView) Reset() {
	*x = ValidatorStakeView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_stake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorStakeView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorStakeView) ProtoMessage() {}

func (x *ValidatorStakeView) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_stake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorStakeView.ProtoReflect.Descriptor instead.
func (*ValidatorStakeView) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_stake_proto_rawDescGZIP(), []int{0}
}

func (m *ValidatorStakeView) GetVariant() isValidatorStakeView_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *ValidatorStakeView) GetV1() *ValidatorStakeView_ValidatorStakeViewV1 {
	if x, ok := x.GetVariant().(*ValidatorStakeView_V1); ok {
		return x.V1
	}
	return nil
}

type isValidatorStakeView_Variant interface {
	isValidatorStakeView_Variant()
}

type ValidatorStakeView_V1 struct {
	V1 *ValidatorStakeView_ValidatorStakeViewV1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*ValidatorStakeView_V1) isValidatorStakeView_Variant() {}

type ValidatorStakeView_ValidatorStakeViewV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string            `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	PublicKey *common.PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	U128Stake []byte            `protobuf:"bytes,3,opt,name=u128_stake,json=u128Stake,proto3" json:"u128_stake,omitempty"`
}

func (x *ValidatorStakeView_ValidatorStakeViewV1) Reset() {
	*x = ValidatorStakeView_ValidatorStakeViewV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_stake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorStakeView_ValidatorStakeViewV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorStakeView_ValidatorStakeViewV1) ProtoMessage() {}

func (x *ValidatorStakeView_ValidatorStakeViewV1) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_stake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorStakeView_ValidatorStakeViewV1.ProtoReflect.Descriptor instead.
func (*ValidatorStakeView_ValidatorStakeViewV1) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_stake_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ValidatorStakeView_ValidatorStakeViewV1) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ValidatorStakeView_ValidatorStakeViewV1) GetPublicKey() *common.PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *ValidatorStakeView_ValidatorStakeViewV1) GetU128Stake() []byte {
	if x != nil {
		return x.U128Stake
	}
	return nil
}

var File_payloads_near_block_stake_proto protoreflect.FileDescriptor

var file_payloads_near_block_stake_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x1a, 0x21, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a,
	0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x51, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x56, 0x69, 0x65, 0x77, 0x56, 0x31,
	0x48, 0x00, 0x52, 0x02, 0x76, 0x31, 0x1a, 0x96, 0x01, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x56, 0x69, 0x65, 0x77, 0x56, 0x31, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x40,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x31, 0x32, 0x38, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x75, 0x31, 0x32, 0x38, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payloads_near_block_stake_proto_rawDescOnce sync.Once
	file_payloads_near_block_stake_proto_rawDescData = file_payloads_near_block_stake_proto_rawDesc
)

func file_payloads_near_block_stake_proto_rawDescGZIP() []byte {
	file_payloads_near_block_stake_proto_rawDescOnce.Do(func() {
		file_payloads_near_block_stake_proto_rawDescData = protoimpl.X.CompressGZIP(file_payloads_near_block_stake_proto_rawDescData)
	})
	return file_payloads_near_block_stake_proto_rawDescData
}

var file_payloads_near_block_stake_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payloads_near_block_stake_proto_goTypes = []any{
	(*ValidatorStakeView)(nil),                      // 0: borealis.payloads.near.ValidatorStakeView
	(*ValidatorStakeView_ValidatorStakeViewV1)(nil), // 1: borealis.payloads.near.ValidatorStakeView.ValidatorStakeViewV1
	(*common.PublicKey)(nil),                        // 2: borealis.payloads.near.PublicKey
}
var file_payloads_near_block_stake_proto_depIdxs = []int32{
	1, // 0: borealis.payloads.near.ValidatorStakeView.v1:type_name -> borealis.payloads.near.ValidatorStakeView.ValidatorStakeViewV1
	2, // 1: borealis.payloads.near.ValidatorStakeView.ValidatorStakeViewV1.public_key:type_name -> borealis.payloads.near.PublicKey
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payloads_near_block_stake_proto_init() }
func file_payloads_near_block_stake_proto_init() {
	if File_payloads_near_block_stake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payloads_near_block_stake_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ValidatorStakeView); i {
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
		file_payloads_near_block_stake_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ValidatorStakeView_ValidatorStakeViewV1); i {
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
	file_payloads_near_block_stake_proto_msgTypes[0].OneofWrappers = []any{
		(*ValidatorStakeView_V1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payloads_near_block_stake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payloads_near_block_stake_proto_goTypes,
		DependencyIndexes: file_payloads_near_block_stake_proto_depIdxs,
		MessageInfos:      file_payloads_near_block_stake_proto_msgTypes,
	}.Build()
	File_payloads_near_block_stake_proto = out.File
	file_payloads_near_block_stake_proto_rawDesc = nil
	file_payloads_near_block_stake_proto_goTypes = nil
	file_payloads_near_block_stake_proto_depIdxs = nil
}
