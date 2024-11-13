// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.6
// source: api/testing-service/v1/enums/testdata.enum.v1.proto

package enumv1

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

// TestdataInit TestdataInit enum
type TestdataInitEnum_TestdataInit int32

const (
	// UNSPECIFIED 未指定
	TestdataInitEnum_UNSPECIFIED TestdataInitEnum_TestdataInit = 0
)

// Enum value maps for TestdataInitEnum_TestdataInit.
var (
	TestdataInitEnum_TestdataInit_name = map[int32]string{
		0: "UNSPECIFIED",
	}
	TestdataInitEnum_TestdataInit_value = map[string]int32{
		"UNSPECIFIED": 0,
	}
)

func (x TestdataInitEnum_TestdataInit) Enum() *TestdataInitEnum_TestdataInit {
	p := new(TestdataInitEnum_TestdataInit)
	*p = x
	return p
}

func (x TestdataInitEnum_TestdataInit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestdataInitEnum_TestdataInit) Descriptor() protoreflect.EnumDescriptor {
	return file_api_testing_service_v1_enums_testdata_enum_v1_proto_enumTypes[0].Descriptor()
}

func (TestdataInitEnum_TestdataInit) Type() protoreflect.EnumType {
	return &file_api_testing_service_v1_enums_testdata_enum_v1_proto_enumTypes[0]
}

func (x TestdataInitEnum_TestdataInit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestdataInitEnum_TestdataInit.Descriptor instead.
func (TestdataInitEnum_TestdataInit) EnumDescriptor() ([]byte, []int) {
	return file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescGZIP(), []int{0, 0}
}

// TestdataInitEnum TestdataInitEnum enum
type TestdataInitEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestdataInitEnum) Reset() {
	*x = TestdataInitEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_testing_service_v1_enums_testdata_enum_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestdataInitEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestdataInitEnum) ProtoMessage() {}

func (x *TestdataInitEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_testing_service_v1_enums_testdata_enum_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestdataInitEnum.ProtoReflect.Descriptor instead.
func (*TestdataInitEnum) Descriptor() ([]byte, []int) {
	return file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_testing_service_v1_enums_testdata_enum_v1_proto protoreflect.FileDescriptor

var file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x22, 0x33, 0x0a,
	0x10, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x69, 0x74, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x1f, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x69,
	0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x42, 0x78, 0x0a, 0x16, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x42, 0x13, 0x4b, 0x69,
	0x74, 0x41, 0x70, 0x69, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x75, 0x6d, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescOnce sync.Once
	file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescData = file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDesc
)

func file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescGZIP() []byte {
	file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescOnce.Do(func() {
		file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescData)
	})
	return file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDescData
}

var file_api_testing_service_v1_enums_testdata_enum_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_testing_service_v1_enums_testdata_enum_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_testing_service_v1_enums_testdata_enum_v1_proto_goTypes = []interface{}{
	(TestdataInitEnum_TestdataInit)(0), // 0: kit.api.testing.enumv1.TestdataInitEnum.TestdataInit
	(*TestdataInitEnum)(nil),           // 1: kit.api.testing.enumv1.TestdataInitEnum
}
var file_api_testing_service_v1_enums_testdata_enum_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_testing_service_v1_enums_testdata_enum_v1_proto_init() }
func file_api_testing_service_v1_enums_testdata_enum_v1_proto_init() {
	if File_api_testing_service_v1_enums_testdata_enum_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_testing_service_v1_enums_testdata_enum_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestdataInitEnum); i {
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
			RawDescriptor: file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_testing_service_v1_enums_testdata_enum_v1_proto_goTypes,
		DependencyIndexes: file_api_testing_service_v1_enums_testdata_enum_v1_proto_depIdxs,
		EnumInfos:         file_api_testing_service_v1_enums_testdata_enum_v1_proto_enumTypes,
		MessageInfos:      file_api_testing_service_v1_enums_testdata_enum_v1_proto_msgTypes,
	}.Build()
	File_api_testing_service_v1_enums_testdata_enum_v1_proto = out.File
	file_api_testing_service_v1_enums_testdata_enum_v1_proto_rawDesc = nil
	file_api_testing_service_v1_enums_testdata_enum_v1_proto_goTypes = nil
	file_api_testing_service_v1_enums_testdata_enum_v1_proto_depIdxs = nil
}
