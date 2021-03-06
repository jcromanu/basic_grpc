// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: records.proto

package record

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

type User_Type int32

const (
	User_FREE       User_Type = 0
	User_ENTERPRISE User_Type = 1
	User_COMPANY    User_Type = 1
)

// Enum value maps for User_Type.
var (
	User_Type_name = map[int32]string{
		0: "FREE",
		1: "ENTERPRISE",
		// Duplicate value: 1: "COMPANY",
	}
	User_Type_value = map[string]int32{
		"FREE":       0,
		"ENTERPRISE": 1,
		"COMPANY":    1,
	}
)

func (x User_Type) Enum() *User_Type {
	p := new(User_Type)
	*p = x
	return p
}

func (x User_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_records_proto_enumTypes[0].Descriptor()
}

func (User_Type) Type() protoreflect.EnumType {
	return &file_records_proto_enumTypes[0]
}

func (x User_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Type.Descriptor instead.
func (User_Type) EnumDescriptor() ([]byte, []int) {
	return file_records_proto_rawDescGZIP(), []int{3, 0}
}

type RecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RecordRequest) Reset() {
	*x = RecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRequest) ProtoMessage() {}

func (x *RecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_records_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRequest.ProtoReflect.Descriptor instead.
func (*RecordRequest) Descriptor() ([]byte, []int) {
	return file_records_proto_rawDescGZIP(), []int{0}
}

func (x *RecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Error  *Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RecordResponse) Reset() {
	*x = RecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResponse) ProtoMessage() {}

func (x *RecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_records_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResponse.ProtoReflect.Descriptor instead.
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return file_records_proto_rawDescGZIP(), []int{1}
}

func (x *RecordResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *RecordResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId string  `protobuf:"bytes,1,opt,name=recordId,proto3" json:"recordId,omitempty"`
	UserId   string  `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Volume   float64 `protobuf:"fixed64,3,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_records_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_records_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *Record) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Record) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string    `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Type   User_Type `protobuf:"varint,2,opt,name=type,proto3,enum=record.User_Type" json:"type,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_records_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_records_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetType() User_Type {
	if x != nil {
		return x.Type
	}
	return User_FREE
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_records_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_records_proto_rawDescGZIP(), []int{4}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_records_proto protoreflect.FileDescriptor

var file_records_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x78, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x31, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x52, 0x45, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x50,
	0x52, 0x49, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e,
	0x59, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf9,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0c, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x28, 0x01, 0x12, 0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_records_proto_rawDescOnce sync.Once
	file_records_proto_rawDescData = file_records_proto_rawDesc
)

func file_records_proto_rawDescGZIP() []byte {
	file_records_proto_rawDescOnce.Do(func() {
		file_records_proto_rawDescData = protoimpl.X.CompressGZIP(file_records_proto_rawDescData)
	})
	return file_records_proto_rawDescData
}

var file_records_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_records_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_records_proto_goTypes = []interface{}{
	(User_Type)(0),         // 0: record.User.Type
	(*RecordRequest)(nil),  // 1: record.RecordRequest
	(*RecordResponse)(nil), // 2: record.RecordResponse
	(*Record)(nil),         // 3: record.Record
	(*User)(nil),           // 4: record.User
	(*Error)(nil),          // 5: record.Error
}
var file_records_proto_depIdxs = []int32{
	3, // 0: record.RecordResponse.record:type_name -> record.Record
	5, // 1: record.RecordResponse.error:type_name -> record.Error
	0, // 2: record.User.type:type_name -> record.User.Type
	1, // 3: record.RecordService.GetRecord:input_type -> record.RecordRequest
	4, // 4: record.RecordService.ListRecords:input_type -> record.User
	1, // 5: record.RecordService.SetRecords:input_type -> record.RecordRequest
	1, // 6: record.RecordService.RecordPong:input_type -> record.RecordRequest
	2, // 7: record.RecordService.GetRecord:output_type -> record.RecordResponse
	2, // 8: record.RecordService.ListRecords:output_type -> record.RecordResponse
	5, // 9: record.RecordService.SetRecords:output_type -> record.Error
	2, // 10: record.RecordService.RecordPong:output_type -> record.RecordResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_records_proto_init() }
func file_records_proto_init() {
	if File_records_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_records_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRequest); i {
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
		file_records_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResponse); i {
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
		file_records_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_records_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_records_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_records_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_records_proto_goTypes,
		DependencyIndexes: file_records_proto_depIdxs,
		EnumInfos:         file_records_proto_enumTypes,
		MessageInfos:      file_records_proto_msgTypes,
	}.Build()
	File_records_proto = out.File
	file_records_proto_rawDesc = nil
	file_records_proto_goTypes = nil
	file_records_proto_depIdxs = nil
}
