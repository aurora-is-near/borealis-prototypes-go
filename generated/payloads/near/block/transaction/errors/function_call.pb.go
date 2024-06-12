// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: payloads/near/block/transaction/errors/function_call.proto

package pb_errors

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

// A kind of a trap happened during execution of a binary
type WasmTrap int32

const (
	// An `unreachable` opcode was executed.
	WasmTrap_UNREACHABLE WasmTrap = 0
	// Call indirect incorrect signature trap.
	WasmTrap_INCORRECT_CALL_INDIRECT_SIGNATURE WasmTrap = 1
	// Memory out of bounds trap.
	WasmTrap_MEMORY_OUT_OF_BOUNDS WasmTrap = 2
	// Call indirect out of bounds trap.
	WasmTrap_CALL_INDIRECT_OOB WasmTrap = 3
	// An arithmetic exception = 0; e.g. divided by zero.
	WasmTrap_ILLEGAL_ARITHMETIC WasmTrap = 4
	// Misaligned atomic access trap.
	WasmTrap_MISALIGNED_ATOMIC_ACCESS WasmTrap = 5
	// Indirect call to null.
	WasmTrap_INDIRECT_CALL_TO_NULL WasmTrap = 6
	// Stack overflow.
	WasmTrap_STACK_OVERFLOW WasmTrap = 7
	// Generic trap.
	WasmTrap_GENERIC_TRAP WasmTrap = 8
)

// Enum value maps for WasmTrap.
var (
	WasmTrap_name = map[int32]string{
		0: "UNREACHABLE",
		1: "INCORRECT_CALL_INDIRECT_SIGNATURE",
		2: "MEMORY_OUT_OF_BOUNDS",
		3: "CALL_INDIRECT_OOB",
		4: "ILLEGAL_ARITHMETIC",
		5: "MISALIGNED_ATOMIC_ACCESS",
		6: "INDIRECT_CALL_TO_NULL",
		7: "STACK_OVERFLOW",
		8: "GENERIC_TRAP",
	}
	WasmTrap_value = map[string]int32{
		"UNREACHABLE":                       0,
		"INCORRECT_CALL_INDIRECT_SIGNATURE": 1,
		"MEMORY_OUT_OF_BOUNDS":              2,
		"CALL_INDIRECT_OOB":                 3,
		"ILLEGAL_ARITHMETIC":                4,
		"MISALIGNED_ATOMIC_ACCESS":          5,
		"INDIRECT_CALL_TO_NULL":             6,
		"STACK_OVERFLOW":                    7,
		"GENERIC_TRAP":                      8,
	}
)

func (x WasmTrap) Enum() *WasmTrap {
	p := new(WasmTrap)
	*p = x
	return p
}

func (x WasmTrap) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WasmTrap) Descriptor() protoreflect.EnumDescriptor {
	return file_payloads_near_block_transaction_errors_function_call_proto_enumTypes[0].Descriptor()
}

func (WasmTrap) Type() protoreflect.EnumType {
	return &file_payloads_near_block_transaction_errors_function_call_proto_enumTypes[0]
}

func (x WasmTrap) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WasmTrap.Descriptor instead.
func (WasmTrap) EnumDescriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0}
}

type MethodResolveError int32

const (
	MethodResolveError_METHOD_EMPTY_NAME        MethodResolveError = 0
	MethodResolveError_METHOD_NOT_FOUND         MethodResolveError = 1
	MethodResolveError_METHOD_INVALID_SIGNATURE MethodResolveError = 2
)

// Enum value maps for MethodResolveError.
var (
	MethodResolveError_name = map[int32]string{
		0: "METHOD_EMPTY_NAME",
		1: "METHOD_NOT_FOUND",
		2: "METHOD_INVALID_SIGNATURE",
	}
	MethodResolveError_value = map[string]int32{
		"METHOD_EMPTY_NAME":        0,
		"METHOD_NOT_FOUND":         1,
		"METHOD_INVALID_SIGNATURE": 2,
	}
)

func (x MethodResolveError) Enum() *MethodResolveError {
	p := new(MethodResolveError)
	*p = x
	return p
}

func (x MethodResolveError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodResolveError) Descriptor() protoreflect.EnumDescriptor {
	return file_payloads_near_block_transaction_errors_function_call_proto_enumTypes[1].Descriptor()
}

func (MethodResolveError) Type() protoreflect.EnumType {
	return &file_payloads_near_block_transaction_errors_function_call_proto_enumTypes[1]
}

func (x MethodResolveError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodResolveError.Descriptor instead.
func (MethodResolveError) EnumDescriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{1}
}

type FunctionCallErrorSer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*FunctionCallErrorSer_CompilationError_
	//	*FunctionCallErrorSer_LinkError_
	//	*FunctionCallErrorSer_MethodResolveError_
	//	*FunctionCallErrorSer_WasmTrap_
	//	*FunctionCallErrorSer_WasmUnknownError_
	//	*FunctionCallErrorSer_HostError_
	//	*FunctionCallErrorSer_ExecutionError_
	Variant isFunctionCallErrorSer_Variant `protobuf_oneof:"variant"`
}

func (x *FunctionCallErrorSer) Reset() {
	*x = FunctionCallErrorSer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer) ProtoMessage() {}

func (x *FunctionCallErrorSer) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0}
}

func (m *FunctionCallErrorSer) GetVariant() isFunctionCallErrorSer_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *FunctionCallErrorSer) GetCompilationError() *FunctionCallErrorSer_CompilationError {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_CompilationError_); ok {
		return x.CompilationError
	}
	return nil
}

func (x *FunctionCallErrorSer) GetLinkError() *FunctionCallErrorSer_LinkError {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_LinkError_); ok {
		return x.LinkError
	}
	return nil
}

func (x *FunctionCallErrorSer) GetMethodResolveError() *FunctionCallErrorSer_MethodResolveError {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_MethodResolveError_); ok {
		return x.MethodResolveError
	}
	return nil
}

func (x *FunctionCallErrorSer) GetWasmTrap() *FunctionCallErrorSer_WasmTrap {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_WasmTrap_); ok {
		return x.WasmTrap
	}
	return nil
}

func (x *FunctionCallErrorSer) GetWasmUnknownError() *FunctionCallErrorSer_WasmUnknownError {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_WasmUnknownError_); ok {
		return x.WasmUnknownError
	}
	return nil
}

func (x *FunctionCallErrorSer) GetHostError() *FunctionCallErrorSer_HostError {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_HostError_); ok {
		return x.HostError
	}
	return nil
}

func (x *FunctionCallErrorSer) GetExecutionError() *FunctionCallErrorSer_ExecutionError {
	if x, ok := x.GetVariant().(*FunctionCallErrorSer_ExecutionError_); ok {
		return x.ExecutionError
	}
	return nil
}

type isFunctionCallErrorSer_Variant interface {
	isFunctionCallErrorSer_Variant()
}

type FunctionCallErrorSer_CompilationError_ struct {
	CompilationError *FunctionCallErrorSer_CompilationError `protobuf:"bytes,1,opt,name=compilation_error,json=compilationError,proto3,oneof"`
}

type FunctionCallErrorSer_LinkError_ struct {
	LinkError *FunctionCallErrorSer_LinkError `protobuf:"bytes,2,opt,name=link_error,json=linkError,proto3,oneof"`
}

type FunctionCallErrorSer_MethodResolveError_ struct {
	MethodResolveError *FunctionCallErrorSer_MethodResolveError `protobuf:"bytes,3,opt,name=method_resolve_error,json=methodResolveError,proto3,oneof"`
}

type FunctionCallErrorSer_WasmTrap_ struct {
	WasmTrap *FunctionCallErrorSer_WasmTrap `protobuf:"bytes,4,opt,name=wasm_trap,json=wasmTrap,proto3,oneof"`
}

type FunctionCallErrorSer_WasmUnknownError_ struct {
	WasmUnknownError *FunctionCallErrorSer_WasmUnknownError `protobuf:"bytes,5,opt,name=wasm_unknown_error,json=wasmUnknownError,proto3,oneof"`
}

type FunctionCallErrorSer_HostError_ struct {
	HostError *FunctionCallErrorSer_HostError `protobuf:"bytes,6,opt,name=host_error,json=hostError,proto3,oneof"`
}

type FunctionCallErrorSer_ExecutionError_ struct {
	ExecutionError *FunctionCallErrorSer_ExecutionError `protobuf:"bytes,7,opt,name=execution_error,json=executionError,proto3,oneof"`
}

func (*FunctionCallErrorSer_CompilationError_) isFunctionCallErrorSer_Variant() {}

func (*FunctionCallErrorSer_LinkError_) isFunctionCallErrorSer_Variant() {}

func (*FunctionCallErrorSer_MethodResolveError_) isFunctionCallErrorSer_Variant() {}

func (*FunctionCallErrorSer_WasmTrap_) isFunctionCallErrorSer_Variant() {}

func (*FunctionCallErrorSer_WasmUnknownError_) isFunctionCallErrorSer_Variant() {}

func (*FunctionCallErrorSer_HostError_) isFunctionCallErrorSer_Variant() {}

func (*FunctionCallErrorSer_ExecutionError_) isFunctionCallErrorSer_Variant() {}

// Wasm compilation error
type FunctionCallErrorSer_CompilationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *CompilationError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FunctionCallErrorSer_CompilationError) Reset() {
	*x = FunctionCallErrorSer_CompilationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_CompilationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_CompilationError) ProtoMessage() {}

func (x *FunctionCallErrorSer_CompilationError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_CompilationError.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_CompilationError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FunctionCallErrorSer_CompilationError) GetError() *CompilationError {
	if x != nil {
		return x.Error
	}
	return nil
}

// Wasm binary env link error
type FunctionCallErrorSer_LinkError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FunctionCallErrorSer_LinkError) Reset() {
	*x = FunctionCallErrorSer_LinkError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_LinkError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_LinkError) ProtoMessage() {}

func (x *FunctionCallErrorSer_LinkError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_LinkError.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_LinkError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FunctionCallErrorSer_LinkError) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// Import/export resolve error
type FunctionCallErrorSer_MethodResolveError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error MethodResolveError `protobuf:"varint,1,opt,name=error,proto3,enum=borealis.payloads.near.MethodResolveError" json:"error,omitempty"`
}

func (x *FunctionCallErrorSer_MethodResolveError) Reset() {
	*x = FunctionCallErrorSer_MethodResolveError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_MethodResolveError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_MethodResolveError) ProtoMessage() {}

func (x *FunctionCallErrorSer_MethodResolveError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_MethodResolveError.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_MethodResolveError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 2}
}

func (x *FunctionCallErrorSer_MethodResolveError) GetError() MethodResolveError {
	if x != nil {
		return x.Error
	}
	return MethodResolveError_METHOD_EMPTY_NAME
}

// A trap happened during execution of a binary
type FunctionCallErrorSer_WasmTrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error WasmTrap `protobuf:"varint,1,opt,name=error,proto3,enum=borealis.payloads.near.WasmTrap" json:"error,omitempty"`
}

func (x *FunctionCallErrorSer_WasmTrap) Reset() {
	*x = FunctionCallErrorSer_WasmTrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_WasmTrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_WasmTrap) ProtoMessage() {}

func (x *FunctionCallErrorSer_WasmTrap) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_WasmTrap.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_WasmTrap) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 3}
}

func (x *FunctionCallErrorSer_WasmTrap) GetError() WasmTrap {
	if x != nil {
		return x.Error
	}
	return WasmTrap_UNREACHABLE
}

type FunctionCallErrorSer_WasmUnknownError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FunctionCallErrorSer_WasmUnknownError) Reset() {
	*x = FunctionCallErrorSer_WasmUnknownError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_WasmUnknownError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_WasmUnknownError) ProtoMessage() {}

func (x *FunctionCallErrorSer_WasmUnknownError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_WasmUnknownError.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_WasmUnknownError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 4}
}

type FunctionCallErrorSer_HostError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *HostError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FunctionCallErrorSer_HostError) Reset() {
	*x = FunctionCallErrorSer_HostError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_HostError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_HostError) ProtoMessage() {}

func (x *FunctionCallErrorSer_HostError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_HostError.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_HostError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 5}
}

func (x *FunctionCallErrorSer_HostError) GetError() *HostError {
	if x != nil {
		return x.Error
	}
	return nil
}

type FunctionCallErrorSer_ExecutionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FunctionCallErrorSer_ExecutionError) Reset() {
	*x = FunctionCallErrorSer_ExecutionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallErrorSer_ExecutionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallErrorSer_ExecutionError) ProtoMessage() {}

func (x *FunctionCallErrorSer_ExecutionError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallErrorSer_ExecutionError.ProtoReflect.Descriptor instead.
func (*FunctionCallErrorSer_ExecutionError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP(), []int{0, 6}
}

func (x *FunctionCallErrorSer_ExecutionError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_payloads_near_block_transaction_errors_function_call_proto protoreflect.FileDescriptor

var file_payloads_near_block_transaction_errors_function_call_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x6f,
	0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e,
	0x6e, 0x65, 0x61, 0x72, 0x1a, 0x31, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e,
	0x65, 0x61, 0x72, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf8, 0x08, 0x0a, 0x14, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61,
	0x6c, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x57, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x62,
	0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61,
	0x6c, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x73, 0x0a, 0x14, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x12, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x54, 0x0a, 0x09, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x74,
	0x72, 0x61, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x6f, 0x72, 0x65,
	0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65,
	0x61, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x73, 0x6d, 0x54, 0x72, 0x61, 0x70,
	0x48, 0x00, 0x52, 0x08, 0x77, 0x61, 0x73, 0x6d, 0x54, 0x72, 0x61, 0x70, 0x12, 0x6d, 0x0a, 0x12,
	0x77, 0x61, 0x73, 0x6d, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61,
	0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61,
	0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x73, 0x6d, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x10, 0x77, 0x61, 0x73, 0x6d, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x57, 0x0a, 0x0a, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x66, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x61, 0x6c, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x52, 0x0a, 0x10,
	0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x3e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x1a, 0x1d, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x1a,
	0x56, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x42, 0x0a, 0x08, 0x57, 0x61, 0x73, 0x6d, 0x54,
	0x72, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x6e, 0x65, 0x61, 0x72, 0x2e, 0x57, 0x61, 0x73, 0x6d,
	0x54, 0x72, 0x61, 0x70, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x12, 0x0a, 0x10, 0x57,
	0x61, 0x73, 0x6d, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a,
	0x44, 0x0a, 0x09, 0x48, 0x6f, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6f,
	0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e,
	0x6e, 0x65, 0x61, 0x72, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x2a, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0xea, 0x01, 0x0a,
	0x08, 0x57, 0x61, 0x73, 0x6d, 0x54, 0x72, 0x61, 0x70, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x52,
	0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x49, 0x4e,
	0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x55, 0x54, 0x5f,
	0x4f, 0x46, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x4f, 0x42,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x41, 0x52,
	0x49, 0x54, 0x48, 0x4d, 0x45, 0x54, 0x49, 0x43, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x49,
	0x53, 0x41, 0x4c, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x4f, 0x4d, 0x49, 0x43, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x55, 0x4c,
	0x4c, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x49, 0x43, 0x5f, 0x54, 0x52, 0x41, 0x50, 0x10, 0x08, 0x2a, 0x5f, 0x0a, 0x12, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x15, 0x0a, 0x11, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53,
	0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payloads_near_block_transaction_errors_function_call_proto_rawDescOnce sync.Once
	file_payloads_near_block_transaction_errors_function_call_proto_rawDescData = file_payloads_near_block_transaction_errors_function_call_proto_rawDesc
)

func file_payloads_near_block_transaction_errors_function_call_proto_rawDescGZIP() []byte {
	file_payloads_near_block_transaction_errors_function_call_proto_rawDescOnce.Do(func() {
		file_payloads_near_block_transaction_errors_function_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_payloads_near_block_transaction_errors_function_call_proto_rawDescData)
	})
	return file_payloads_near_block_transaction_errors_function_call_proto_rawDescData
}

var file_payloads_near_block_transaction_errors_function_call_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_payloads_near_block_transaction_errors_function_call_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_payloads_near_block_transaction_errors_function_call_proto_goTypes = []any{
	(WasmTrap)(0),                                   // 0: borealis.payloads.near.WasmTrap
	(MethodResolveError)(0),                         // 1: borealis.payloads.near.MethodResolveError
	(*FunctionCallErrorSer)(nil),                    // 2: borealis.payloads.near.FunctionCallErrorSer
	(*FunctionCallErrorSer_CompilationError)(nil),   // 3: borealis.payloads.near.FunctionCallErrorSer.CompilationError
	(*FunctionCallErrorSer_LinkError)(nil),          // 4: borealis.payloads.near.FunctionCallErrorSer.LinkError
	(*FunctionCallErrorSer_MethodResolveError)(nil), // 5: borealis.payloads.near.FunctionCallErrorSer.MethodResolveError
	(*FunctionCallErrorSer_WasmTrap)(nil),           // 6: borealis.payloads.near.FunctionCallErrorSer.WasmTrap
	(*FunctionCallErrorSer_WasmUnknownError)(nil),   // 7: borealis.payloads.near.FunctionCallErrorSer.WasmUnknownError
	(*FunctionCallErrorSer_HostError)(nil),          // 8: borealis.payloads.near.FunctionCallErrorSer.HostError
	(*FunctionCallErrorSer_ExecutionError)(nil),     // 9: borealis.payloads.near.FunctionCallErrorSer.ExecutionError
	(*CompilationError)(nil),                        // 10: borealis.payloads.near.CompilationError
	(*HostError)(nil),                               // 11: borealis.payloads.near.HostError
}
var file_payloads_near_block_transaction_errors_function_call_proto_depIdxs = []int32{
	3,  // 0: borealis.payloads.near.FunctionCallErrorSer.compilation_error:type_name -> borealis.payloads.near.FunctionCallErrorSer.CompilationError
	4,  // 1: borealis.payloads.near.FunctionCallErrorSer.link_error:type_name -> borealis.payloads.near.FunctionCallErrorSer.LinkError
	5,  // 2: borealis.payloads.near.FunctionCallErrorSer.method_resolve_error:type_name -> borealis.payloads.near.FunctionCallErrorSer.MethodResolveError
	6,  // 3: borealis.payloads.near.FunctionCallErrorSer.wasm_trap:type_name -> borealis.payloads.near.FunctionCallErrorSer.WasmTrap
	7,  // 4: borealis.payloads.near.FunctionCallErrorSer.wasm_unknown_error:type_name -> borealis.payloads.near.FunctionCallErrorSer.WasmUnknownError
	8,  // 5: borealis.payloads.near.FunctionCallErrorSer.host_error:type_name -> borealis.payloads.near.FunctionCallErrorSer.HostError
	9,  // 6: borealis.payloads.near.FunctionCallErrorSer.execution_error:type_name -> borealis.payloads.near.FunctionCallErrorSer.ExecutionError
	10, // 7: borealis.payloads.near.FunctionCallErrorSer.CompilationError.error:type_name -> borealis.payloads.near.CompilationError
	1,  // 8: borealis.payloads.near.FunctionCallErrorSer.MethodResolveError.error:type_name -> borealis.payloads.near.MethodResolveError
	0,  // 9: borealis.payloads.near.FunctionCallErrorSer.WasmTrap.error:type_name -> borealis.payloads.near.WasmTrap
	11, // 10: borealis.payloads.near.FunctionCallErrorSer.HostError.error:type_name -> borealis.payloads.near.HostError
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_payloads_near_block_transaction_errors_function_call_proto_init() }
func file_payloads_near_block_transaction_errors_function_call_proto_init() {
	if File_payloads_near_block_transaction_errors_function_call_proto != nil {
		return
	}
	file_payloads_near_block_transaction_errors_host_proto_init()
	file_payloads_near_block_transaction_errors_compilation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_CompilationError); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_LinkError); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_MethodResolveError); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_WasmTrap); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_WasmUnknownError); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_HostError); i {
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
		file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionCallErrorSer_ExecutionError); i {
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
	file_payloads_near_block_transaction_errors_function_call_proto_msgTypes[0].OneofWrappers = []any{
		(*FunctionCallErrorSer_CompilationError_)(nil),
		(*FunctionCallErrorSer_LinkError_)(nil),
		(*FunctionCallErrorSer_MethodResolveError_)(nil),
		(*FunctionCallErrorSer_WasmTrap_)(nil),
		(*FunctionCallErrorSer_WasmUnknownError_)(nil),
		(*FunctionCallErrorSer_HostError_)(nil),
		(*FunctionCallErrorSer_ExecutionError_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payloads_near_block_transaction_errors_function_call_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payloads_near_block_transaction_errors_function_call_proto_goTypes,
		DependencyIndexes: file_payloads_near_block_transaction_errors_function_call_proto_depIdxs,
		EnumInfos:         file_payloads_near_block_transaction_errors_function_call_proto_enumTypes,
		MessageInfos:      file_payloads_near_block_transaction_errors_function_call_proto_msgTypes,
	}.Build()
	File_payloads_near_block_transaction_errors_function_call_proto = out.File
	file_payloads_near_block_transaction_errors_function_call_proto_rawDesc = nil
	file_payloads_near_block_transaction_errors_function_call_proto_goTypes = nil
	file_payloads_near_block_transaction_errors_function_call_proto_depIdxs = nil
}
