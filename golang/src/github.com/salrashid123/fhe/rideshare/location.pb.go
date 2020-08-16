// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: src/github.com/salrashid123/fhe/rideshare/location.proto

package rideshare

import (
	proto "github.com/golang/protobuf/proto"
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

type EncryptedCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	X  []byte `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y  []byte `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	Pk []byte `protobuf:"bytes,4,opt,name=pk,proto3" json:"pk,omitempty"`
}

func (x *EncryptedCoordinate) Reset() {
	*x = EncryptedCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedCoordinate) ProtoMessage() {}

func (x *EncryptedCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedCoordinate.ProtoReflect.Descriptor instead.
func (*EncryptedCoordinate) Descriptor() ([]byte, []int) {
	return file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptedCoordinate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EncryptedCoordinate) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *EncryptedCoordinate) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *EncryptedCoordinate) GetPk() []byte {
	if x != nil {
		return x.Pk
	}
	return nil
}

type Distance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid  string `protobuf:"bytes,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Did  string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Dist []byte `protobuf:"bytes,3,opt,name=dist,proto3" json:"dist,omitempty"`
}

func (x *Distance) Reset() {
	*x = Distance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distance) ProtoMessage() {}

func (x *Distance) ProtoReflect() protoreflect.Message {
	mi := &file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distance.ProtoReflect.Descriptor instead.
func (*Distance) Descriptor() ([]byte, []int) {
	return file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescGZIP(), []int{1}
}

func (x *Distance) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *Distance) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Distance) GetDist() []byte {
	if x != nil {
		return x.Dist
	}
	return nil
}

type DecryptedCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	X  uint64 `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y  uint64 `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *DecryptedCoordinate) Reset() {
	*x = DecryptedCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptedCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptedCoordinate) ProtoMessage() {}

func (x *DecryptedCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptedCoordinate.ProtoReflect.Descriptor instead.
func (*DecryptedCoordinate) Descriptor() ([]byte, []int) {
	return file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescGZIP(), []int{2}
}

func (x *DecryptedCoordinate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DecryptedCoordinate) GetX() uint64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *DecryptedCoordinate) GetY() uint64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_src_github_com_salrashid123_fhe_rideshare_location_proto protoreflect.FileDescriptor

var file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x61, 0x6c, 0x72, 0x61, 0x73, 0x68, 0x69, 0x64, 0x31, 0x32, 0x33, 0x2f, 0x66, 0x68,
	0x65, 0x2f, 0x72, 0x69, 0x64, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x70, 0x6b, 0x22, 0x42, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x44,
	0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x79, 0x42, 0x27,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c,
	0x72, 0x61, 0x73, 0x68, 0x69, 0x64, 0x31, 0x32, 0x33, 0x2f, 0x66, 0x68, 0x65, 0x2f, 0x72, 0x69,
	0x64, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescOnce sync.Once
	file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescData = file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDesc
)

func file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescGZIP() []byte {
	file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescOnce.Do(func() {
		file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescData)
	})
	return file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDescData
}

var file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_github_com_salrashid123_fhe_rideshare_location_proto_goTypes = []interface{}{
	(*EncryptedCoordinate)(nil), // 0: location.EncryptedCoordinate
	(*Distance)(nil),            // 1: location.Distance
	(*DecryptedCoordinate)(nil), // 2: location.DecryptedCoordinate
}
var file_src_github_com_salrashid123_fhe_rideshare_location_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_github_com_salrashid123_fhe_rideshare_location_proto_init() }
func file_src_github_com_salrashid123_fhe_rideshare_location_proto_init() {
	if File_src_github_com_salrashid123_fhe_rideshare_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedCoordinate); i {
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
		file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distance); i {
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
		file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptedCoordinate); i {
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
			RawDescriptor: file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_github_com_salrashid123_fhe_rideshare_location_proto_goTypes,
		DependencyIndexes: file_src_github_com_salrashid123_fhe_rideshare_location_proto_depIdxs,
		MessageInfos:      file_src_github_com_salrashid123_fhe_rideshare_location_proto_msgTypes,
	}.Build()
	File_src_github_com_salrashid123_fhe_rideshare_location_proto = out.File
	file_src_github_com_salrashid123_fhe_rideshare_location_proto_rawDesc = nil
	file_src_github_com_salrashid123_fhe_rideshare_location_proto_goTypes = nil
	file_src_github_com_salrashid123_fhe_rideshare_location_proto_depIdxs = nil
}
