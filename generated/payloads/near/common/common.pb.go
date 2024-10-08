// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: payloads/near/common/common.proto

package pb_common

import (
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

type OptionalSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Signature `protobuf:"bytes,1,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *OptionalSignature) Reset() {
	*x = OptionalSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalSignature) ProtoMessage() {}

func (x *OptionalSignature) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalSignature.ProtoReflect.Descriptor instead.
func (*OptionalSignature) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *OptionalSignature) GetValue() *Signature {
	if x != nil {
		return x.Value
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*Signature_Ed25519
	//	*Signature_Secp256K1
	Variant isSignature_Variant `protobuf_oneof:"variant"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{1}
}

func (m *Signature) GetVariant() isSignature_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *Signature) GetEd25519() *Signature_ED25519 {
	if x, ok := x.GetVariant().(*Signature_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (x *Signature) GetSecp256K1() *Signature_SECP256K1 {
	if x, ok := x.GetVariant().(*Signature_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

type isSignature_Variant interface {
	isSignature_Variant()
}

type Signature_Ed25519 struct {
	Ed25519 *Signature_ED25519 `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}

type Signature_Secp256K1 struct {
	Secp256K1 *Signature_SECP256K1 `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof"`
}

func (*Signature_Ed25519) isSignature_Variant() {}

func (*Signature_Secp256K1) isSignature_Variant() {}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*PublicKey_Ed25519
	//	*PublicKey_Secp256K1
	Variant isPublicKey_Variant `protobuf_oneof:"variant"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{2}
}

func (m *PublicKey) GetVariant() isPublicKey_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *PublicKey) GetEd25519() *PublicKey_ED25519 {
	if x, ok := x.GetVariant().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (x *PublicKey) GetSecp256K1() *PublicKey_SECP256K1 {
	if x, ok := x.GetVariant().(*PublicKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

type isPublicKey_Variant interface {
	isPublicKey_Variant()
}

type PublicKey_Ed25519 struct {
	Ed25519 *PublicKey_ED25519 `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}

type PublicKey_Secp256K1 struct {
	Secp256K1 *PublicKey_SECP256K1 `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_Variant() {}

func (*PublicKey_Secp256K1) isPublicKey_Variant() {}

type Signature_ED25519 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	H512Value []byte `protobuf:"bytes,1,opt,name=h512_value,json=h512Value,proto3" json:"h512_value,omitempty"`
}

func (x *Signature_ED25519) Reset() {
	*x = Signature_ED25519{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature_ED25519) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_ED25519) ProtoMessage() {}

func (x *Signature_ED25519) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_ED25519.ProtoReflect.Descriptor instead.
func (*Signature_ED25519) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Signature_ED25519) GetH512Value() []byte {
	if x != nil {
		return x.H512Value
	}
	return nil
}

type Signature_SECP256K1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	H520Value []byte `protobuf:"bytes,1,opt,name=h520_value,json=h520Value,proto3" json:"h520_value,omitempty"`
}

func (x *Signature_SECP256K1) Reset() {
	*x = Signature_SECP256K1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature_SECP256K1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_SECP256K1) ProtoMessage() {}

func (x *Signature_SECP256K1) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_SECP256K1.ProtoReflect.Descriptor instead.
func (*Signature_SECP256K1) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Signature_SECP256K1) GetH520Value() []byte {
	if x != nil {
		return x.H520Value
	}
	return nil
}

type PublicKey_ED25519 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	H256Value []byte `protobuf:"bytes,1,opt,name=h256_value,json=h256Value,proto3" json:"h256_value,omitempty"`
}

func (x *PublicKey_ED25519) Reset() {
	*x = PublicKey_ED25519{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey_ED25519) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey_ED25519) ProtoMessage() {}

func (x *PublicKey_ED25519) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey_ED25519.ProtoReflect.Descriptor instead.
func (*PublicKey_ED25519) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PublicKey_ED25519) GetH256Value() []byte {
	if x != nil {
		return x.H256Value
	}
	return nil
}

type PublicKey_SECP256K1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	H512Value []byte `protobuf:"bytes,1,opt,name=h512_value,json=h512Value,proto3" json:"h512_value,omitempty"`
}

func (x *PublicKey_SECP256K1) Reset() {
	*x = PublicKey_SECP256K1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey_SECP256K1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey_SECP256K1) ProtoMessage() {}

func (x *PublicKey_SECP256K1) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey_SECP256K1.ProtoReflect.Descriptor instead.
func (*PublicKey_SECP256K1) Descriptor() ([]byte, []int) {
	return file_payloads_near_common_common_proto_rawDescGZIP(), []int{2, 1}
}

func (x *PublicKey_SECP256K1) GetH512Value() []byte {
	if x != nil {
		return x.H512Value
	}
	return nil
}

var File_payloads_near_common_common_proto protoreflect.FileDescriptor

var file_payloads_near_common_common_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x22, 0x5b, 0x0a, 0x11, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31,
	0x39, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c,
	0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x45, 0x44, 0x32, 0x35, 0x35,
	0x31, 0x39, 0x48, 0x00, 0x52, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x4b, 0x0a,
	0x09, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x53, 0x45, 0x43, 0x50, 0x32, 0x35, 0x36, 0x4b, 0x31, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x1a, 0x28, 0x0a, 0x07, 0x45, 0x44,
	0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x35, 0x31, 0x32, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x68, 0x35, 0x31, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x2a, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x50, 0x32, 0x35, 0x36, 0x4b,
	0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x35, 0x32, 0x30, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x68, 0x35, 0x32, 0x30, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x09,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x07, 0x65, 0x64, 0x32,
	0x35, 0x35, 0x31, 0x39, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6f, 0x72,
	0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e,
	0x65, 0x61, 0x72, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x2e, 0x45, 0x44,
	0x32, 0x35, 0x35, 0x31, 0x39, 0x48, 0x00, 0x52, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39,
	0x12, 0x4b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x2e, 0x53, 0x45, 0x43, 0x50, 0x32, 0x35, 0x36, 0x4b, 0x31,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x1a, 0x28, 0x0a,
	0x07, 0x45, 0x44, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x32, 0x35, 0x36,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x68, 0x32,
	0x35, 0x36, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x2a, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x50, 0x32,
	0x35, 0x36, 0x4b, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x35, 0x31, 0x32, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x68, 0x35, 0x31, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payloads_near_common_common_proto_rawDescOnce sync.Once
	file_payloads_near_common_common_proto_rawDescData = file_payloads_near_common_common_proto_rawDesc
)

func file_payloads_near_common_common_proto_rawDescGZIP() []byte {
	file_payloads_near_common_common_proto_rawDescOnce.Do(func() {
		file_payloads_near_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_payloads_near_common_common_proto_rawDescData)
	})
	return file_payloads_near_common_common_proto_rawDescData
}

var file_payloads_near_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_payloads_near_common_common_proto_goTypes = []any{
	(*OptionalSignature)(nil),   // 0: borealis.payloads.near.OptionalSignature
	(*Signature)(nil),           // 1: borealis.payloads.near.Signature
	(*PublicKey)(nil),           // 2: borealis.payloads.near.PublicKey
	(*Signature_ED25519)(nil),   // 3: borealis.payloads.near.Signature.ED25519
	(*Signature_SECP256K1)(nil), // 4: borealis.payloads.near.Signature.SECP256K1
	(*PublicKey_ED25519)(nil),   // 5: borealis.payloads.near.PublicKey.ED25519
	(*PublicKey_SECP256K1)(nil), // 6: borealis.payloads.near.PublicKey.SECP256K1
}
var file_payloads_near_common_common_proto_depIdxs = []int32{
	1, // 0: borealis.payloads.near.OptionalSignature.value:type_name -> borealis.payloads.near.Signature
	3, // 1: borealis.payloads.near.Signature.ed25519:type_name -> borealis.payloads.near.Signature.ED25519
	4, // 2: borealis.payloads.near.Signature.secp256k1:type_name -> borealis.payloads.near.Signature.SECP256K1
	5, // 3: borealis.payloads.near.PublicKey.ed25519:type_name -> borealis.payloads.near.PublicKey.ED25519
	6, // 4: borealis.payloads.near.PublicKey.secp256k1:type_name -> borealis.payloads.near.PublicKey.SECP256K1
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_payloads_near_common_common_proto_init() }
func file_payloads_near_common_common_proto_init() {
	if File_payloads_near_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payloads_near_common_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OptionalSignature); i {
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
		file_payloads_near_common_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Signature); i {
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
		file_payloads_near_common_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PublicKey); i {
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
		file_payloads_near_common_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Signature_ED25519); i {
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
		file_payloads_near_common_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Signature_SECP256K1); i {
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
		file_payloads_near_common_common_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PublicKey_ED25519); i {
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
		file_payloads_near_common_common_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PublicKey_SECP256K1); i {
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
	file_payloads_near_common_common_proto_msgTypes[0].OneofWrappers = []any{}
	file_payloads_near_common_common_proto_msgTypes[1].OneofWrappers = []any{
		(*Signature_Ed25519)(nil),
		(*Signature_Secp256K1)(nil),
	}
	file_payloads_near_common_common_proto_msgTypes[2].OneofWrappers = []any{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_Secp256K1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payloads_near_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payloads_near_common_common_proto_goTypes,
		DependencyIndexes: file_payloads_near_common_common_proto_depIdxs,
		MessageInfos:      file_payloads_near_common_common_proto_msgTypes,
	}.Build()
	File_payloads_near_common_common_proto = out.File
	file_payloads_near_common_common_proto_rawDesc = nil
	file_payloads_near_common_common_proto_goTypes = nil
	file_payloads_near_common_common_proto_depIdxs = nil
}
