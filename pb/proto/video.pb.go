// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/video.proto

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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Encoding string `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Frame    *Frame `protobuf:"bytes,3,opt,name=frame,proto3" json:"frame,omitempty"`
	DeviceId string `protobuf:"bytes,4,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Video) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *Video) GetFrame() *Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *Video) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number    int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	LastChunk bool   `protobuf:"varint,2,opt,name=lastChunk,proto3" json:"lastChunk,omitempty"`
	Chunk     []byte `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{1}
}

func (x *Frame) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Frame) GetLastChunk() bool {
	if x != nil {
		return x.LastChunk
	}
	return false
}

func (x *Frame) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type EmptyVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyVideoResponse) Reset() {
	*x = EmptyVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyVideoResponse) ProtoMessage() {}

func (x *EmptyVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyVideoResponse.ProtoReflect.Descriptor instead.
func (*EmptyVideoResponse) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{2}
}

type PullVideoStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MainUser string `protobuf:"bytes,2,opt,name=mainUser,proto3" json:"mainUser,omitempty"`
}

func (x *PullVideoStreamReq) Reset() {
	*x = PullVideoStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullVideoStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullVideoStreamReq) ProtoMessage() {}

func (x *PullVideoStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullVideoStreamReq.ProtoReflect.Descriptor instead.
func (*PullVideoStreamReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{3}
}

func (x *PullVideoStreamReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PullVideoStreamReq) GetMainUser() string {
	if x != nil {
		return x.MainUser
	}
	return ""
}

type PullVideoStreamResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video  *Video `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	Closed bool   `protobuf:"varint,2,opt,name=closed,proto3" json:"closed,omitempty"`
}

func (x *PullVideoStreamResp) Reset() {
	*x = PullVideoStreamResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullVideoStreamResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullVideoStreamResp) ProtoMessage() {}

func (x *PullVideoStreamResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullVideoStreamResp.ProtoReflect.Descriptor instead.
func (*PullVideoStreamResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{4}
}

func (x *PullVideoStreamResp) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *PullVideoStreamResp) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

type EndPullVideoStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MainUser string `protobuf:"bytes,2,opt,name=mainUser,proto3" json:"mainUser,omitempty"`
}

func (x *EndPullVideoStreamReq) Reset() {
	*x = EndPullVideoStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndPullVideoStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndPullVideoStreamReq) ProtoMessage() {}

func (x *EndPullVideoStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndPullVideoStreamReq.ProtoReflect.Descriptor instead.
func (*EndPullVideoStreamReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{5}
}

func (x *EndPullVideoStreamReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EndPullVideoStreamReq) GetMainUser() string {
	if x != nil {
		return x.MainUser
	}
	return ""
}

//a stream request to notify when the de1 need to start sending frames
type Streamrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request bool `protobuf:"varint,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *Streamrequest) Reset() {
	*x = Streamrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Streamrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Streamrequest) ProtoMessage() {}

func (x *Streamrequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Streamrequest.ProtoReflect.Descriptor instead.
func (*Streamrequest) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{6}
}

func (x *Streamrequest) GetRequest() bool {
	if x != nil {
		return x.Request
	}
	return false
}

//a request from de1 to first time set up the connection with server
type InitialConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setup bool `protobuf:"varint,1,opt,name=setup,proto3" json:"setup,omitempty"`
}

func (x *InitialConnection) Reset() {
	*x = InitialConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitialConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialConnection) ProtoMessage() {}

func (x *InitialConnection) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialConnection.ProtoReflect.Descriptor instead.
func (*InitialConnection) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{7}
}

func (x *InitialConnection) GetSetup() bool {
	if x != nil {
		return x.Setup
	}
	return false
}

var File_proto_video_proto protoreflect.FileDescriptor

var file_proto_video_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x77, 0x0a, 0x05, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40,
	0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x51, 0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x15, 0x45, 0x6e, 0x64, 0x50, 0x75, 0x6c, 0x6c, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x61, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x61, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x65, 0x74, 0x75, 0x70, 0x32, 0xb0,
	0x02, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0c, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x19, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x0f, 0x50, 0x75, 0x6c,
	0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x50, 0x75,
	0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x45, 0x6e, 0x64, 0x50, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x50, 0x45, 0x4e, 0x33, 0x39, 0x31, 0x2d, 0x54, 0x65, 0x61, 0x6d, 0x2d, 0x34, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_video_proto_rawDescOnce sync.Once
	file_proto_video_proto_rawDescData = file_proto_video_proto_rawDesc
)

func file_proto_video_proto_rawDescGZIP() []byte {
	file_proto_video_proto_rawDescOnce.Do(func() {
		file_proto_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_video_proto_rawDescData)
	})
	return file_proto_video_proto_rawDescData
}

var file_proto_video_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_video_proto_goTypes = []interface{}{
	(*Video)(nil),                 // 0: video.Video
	(*Frame)(nil),                 // 1: video.Frame
	(*EmptyVideoResponse)(nil),    // 2: video.EmptyVideoResponse
	(*PullVideoStreamReq)(nil),    // 3: video.PullVideoStreamReq
	(*PullVideoStreamResp)(nil),   // 4: video.PullVideoStreamResp
	(*EndPullVideoStreamReq)(nil), // 5: video.EndPullVideoStreamReq
	(*Streamrequest)(nil),         // 6: video.Streamrequest
	(*InitialConnection)(nil),     // 7: video.InitialConnection
}
var file_proto_video_proto_depIdxs = []int32{
	1, // 0: video.Video.frame:type_name -> video.Frame
	0, // 1: video.PullVideoStreamResp.video:type_name -> video.Video
	0, // 2: video.VideoRoute.StreamVideo:input_type -> video.Video
	3, // 3: video.VideoRoute.PullVideoStream:input_type -> video.PullVideoStreamReq
	5, // 4: video.VideoRoute.EndPullVideoStream:input_type -> video.EndPullVideoStreamReq
	7, // 5: video.VideoRoute.RequestToStream:input_type -> video.InitialConnection
	2, // 6: video.VideoRoute.StreamVideo:output_type -> video.EmptyVideoResponse
	4, // 7: video.VideoRoute.PullVideoStream:output_type -> video.PullVideoStreamResp
	2, // 8: video.VideoRoute.EndPullVideoStream:output_type -> video.EmptyVideoResponse
	6, // 9: video.VideoRoute.RequestToStream:output_type -> video.Streamrequest
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_video_proto_init() }
func file_proto_video_proto_init() {
	if File_proto_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_proto_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyVideoResponse); i {
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
		file_proto_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullVideoStreamReq); i {
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
		file_proto_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullVideoStreamResp); i {
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
		file_proto_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndPullVideoStreamReq); i {
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
		file_proto_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Streamrequest); i {
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
		file_proto_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitialConnection); i {
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
			RawDescriptor: file_proto_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_video_proto_goTypes,
		DependencyIndexes: file_proto_video_proto_depIdxs,
		MessageInfos:      file_proto_video_proto_msgTypes,
	}.Build()
	File_proto_video_proto = out.File
	file_proto_video_proto_rawDesc = nil
	file_proto_video_proto_goTypes = nil
	file_proto_video_proto_depIdxs = nil
}
