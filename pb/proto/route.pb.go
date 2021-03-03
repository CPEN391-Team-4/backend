// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/route.proto

package backend

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Restricted bool   `protobuf:"varint,2,opt,name=restricted,proto3" json:"restricted,omitempty"`
	Photo      *Photo `protobuf:"bytes,3,opt,name=Photo,proto3" json:"Photo,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_route_proto_msgTypes[0]
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
	return file_proto_route_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetRestricted() bool {
	if x != nil {
		return x.Restricted
	}
	return false
}

func (x *User) GetPhoto() *Photo {
	if x != nil {
		return x.Photo
	}
	return nil
}

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image         []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	FileExtension string `protobuf:"bytes,2,opt,name=fileExtension,proto3" json:"fileExtension,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_proto_route_proto_rawDescGZIP(), []int{1}
}

func (x *Photo) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Photo) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

type UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usernames string `protobuf:"bytes,1,opt,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *UserName) Reset() {
	*x = UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserName) ProtoMessage() {}

func (x *UserName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserName.ProtoReflect.Descriptor instead.
func (*UserName) Descriptor() ([]byte, []int) {
	return file_proto_route_proto_rawDescGZIP(), []int{2}
}

func (x *UserName) GetUsernames() string {
	if x != nil {
		return x.Usernames
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_route_proto_rawDescGZIP(), []int{3}
}

var File_proto_route_proto protoreflect.FileDescriptor

var file_proto_route_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x5e, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x05, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x28, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x81, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x32, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x30, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x50, 0x45, 0x4e, 0x33, 0x39, 0x31, 0x2d, 0x54, 0x65, 0x61,
	0x6d, 0x2d, 0x34, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_route_proto_rawDescOnce sync.Once
	file_proto_route_proto_rawDescData = file_proto_route_proto_rawDesc
)

func file_proto_route_proto_rawDescGZIP() []byte {
	file_proto_route_proto_rawDescOnce.Do(func() {
		file_proto_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_route_proto_rawDescData)
	})
	return file_proto_route_proto_rawDescData
}

var file_proto_route_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_route_proto_goTypes = []interface{}{
	(*User)(nil),     // 0: route.User
	(*Photo)(nil),    // 1: route.Photo
	(*UserName)(nil), // 2: route.UserName
	(*Empty)(nil),    // 3: route.Empty
}
var file_proto_route_proto_depIdxs = []int32{
	1, // 0: route.User.Photo:type_name -> route.Photo
	0, // 1: route.Route.AddTrustedUser:input_type -> route.User
	0, // 2: route.Route.UpdateTrustedUser:input_type -> route.User
	0, // 3: route.Route.RemoveTrustedUser:input_type -> route.User
	0, // 4: route.Route.GetUserPhoto:input_type -> route.User
	3, // 5: route.Route.GetAllUserNames:input_type -> route.Empty
	3, // 6: route.Route.AddTrustedUser:output_type -> route.Empty
	3, // 7: route.Route.UpdateTrustedUser:output_type -> route.Empty
	3, // 8: route.Route.RemoveTrustedUser:output_type -> route.Empty
	1, // 9: route.Route.GetUserPhoto:output_type -> route.Photo
	2, // 10: route.Route.GetAllUserNames:output_type -> route.UserName
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_route_proto_init() }
func file_proto_route_proto_init() {
	if File_proto_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
		file_proto_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserName); i {
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
		file_proto_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_route_proto_goTypes,
		DependencyIndexes: file_proto_route_proto_depIdxs,
		MessageInfos:      file_proto_route_proto_msgTypes,
	}.Build()
	File_proto_route_proto = out.File
	file_proto_route_proto_rawDesc = nil
	file_proto_route_proto_goTypes = nil
	file_proto_route_proto_depIdxs = nil
}
