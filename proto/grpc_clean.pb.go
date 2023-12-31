// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/grpc_clean.proto

package proto

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

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{0}
}

func (x *SuccessResponse) GetStatus() string {
	if x != nil {
		return x.Status
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
		mi := &file_proto_grpc_clean_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[1]
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
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{1}
}

type TasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Tasks []string `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TasksRequest) Reset() {
	*x = TasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksRequest) ProtoMessage() {}

func (x *TasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksRequest.ProtoReflect.Descriptor instead.
func (*TasksRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{2}
}

func (x *TasksRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TasksRequest) GetTasks() []string {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{3}
}

func (x *TaskRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start  int32  `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End    int32  `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Entity string `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Label  string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{4}
}

func (x *Entity) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Entity) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Entity) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *Entity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Task1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessedText string    `protobuf:"bytes,1,opt,name=processed_text,json=processedText,proto3" json:"processed_text,omitempty"`
	Entities      []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *Task1Response) Reset() {
	*x = Task1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task1Response) ProtoMessage() {}

func (x *Task1Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task1Response.ProtoReflect.Descriptor instead.
func (*Task1Response) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{5}
}

func (x *Task1Response) GetProcessedText() string {
	if x != nil {
		return x.ProcessedText
	}
	return ""
}

func (x *Task1Response) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type Task2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task2 string `protobuf:"bytes,1,opt,name=task2,proto3" json:"task2,omitempty"`
}

func (x *Task2Response) Reset() {
	*x = Task2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task2Response) ProtoMessage() {}

func (x *Task2Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task2Response.ProtoReflect.Descriptor instead.
func (*Task2Response) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{6}
}

func (x *Task2Response) GetTask2() string {
	if x != nil {
		return x.Task2
	}
	return ""
}

type Task3Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task3 int32 `protobuf:"varint,1,opt,name=task3,proto3" json:"task3,omitempty"`
}

func (x *Task3Response) Reset() {
	*x = Task3Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task3Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task3Response) ProtoMessage() {}

func (x *Task3Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task3Response.ProtoReflect.Descriptor instead.
func (*Task3Response) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{7}
}

func (x *Task3Response) GetTask3() int32 {
	if x != nil {
		return x.Task3
	}
	return 0
}

type TasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task1 *Task1Response `protobuf:"bytes,1,opt,name=task1,proto3" json:"task1,omitempty"`
	Task2 string         `protobuf:"bytes,2,opt,name=task2,proto3" json:"task2,omitempty"`
	Task3 int32          `protobuf:"varint,3,opt,name=task3,proto3" json:"task3,omitempty"`
}

func (x *TasksResponse) Reset() {
	*x = TasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_clean_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksResponse) ProtoMessage() {}

func (x *TasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_clean_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksResponse.ProtoReflect.Descriptor instead.
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_clean_proto_rawDescGZIP(), []int{8}
}

func (x *TasksResponse) GetTask1() *Task1Response {
	if x != nil {
		return x.Task1
	}
	return nil
}

func (x *TasksResponse) GetTask2() string {
	if x != nil {
		return x.Task2
	}
	return ""
}

func (x *TasksResponse) GetTask3() int32 {
	if x != nil {
		return x.Task3
	}
	return 0
}

var File_proto_grpc_clean_proto protoreflect.FileDescriptor

var file_proto_grpc_clean_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x65,
	0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x0c,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x5e, 0x0a, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x5b, 0x0a, 0x0d, 0x54, 0x61, 0x73,
	0x6b, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x32, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x32,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x32, 0x22, 0x25, 0x0a,
	0x0d, 0x54, 0x61, 0x73, 0x6b, 0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x33, 0x22, 0x61, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x33, 0x32, 0xd2, 0x01, 0x0a, 0x10, 0x47, 0x52, 0x50, 0x43,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0d, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x31,
	0x12, 0x0c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x32, 0x12, 0x0c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x32, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x33, 0x12, 0x0c,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_clean_proto_rawDescOnce sync.Once
	file_proto_grpc_clean_proto_rawDescData = file_proto_grpc_clean_proto_rawDesc
)

func file_proto_grpc_clean_proto_rawDescGZIP() []byte {
	file_proto_grpc_clean_proto_rawDescOnce.Do(func() {
		file_proto_grpc_clean_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_clean_proto_rawDescData)
	})
	return file_proto_grpc_clean_proto_rawDescData
}

var file_proto_grpc_clean_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_grpc_clean_proto_goTypes = []interface{}{
	(*SuccessResponse)(nil), // 0: SuccessResponse
	(*Empty)(nil),           // 1: Empty
	(*TasksRequest)(nil),    // 2: TasksRequest
	(*TaskRequest)(nil),     // 3: TaskRequest
	(*Entity)(nil),          // 4: Entity
	(*Task1Response)(nil),   // 5: Task1Response
	(*Task2Response)(nil),   // 6: Task2Response
	(*Task3Response)(nil),   // 7: Task3Response
	(*TasksResponse)(nil),   // 8: TasksResponse
}
var file_proto_grpc_clean_proto_depIdxs = []int32{
	4, // 0: Task1Response.entities:type_name -> Entity
	5, // 1: TasksResponse.task1:type_name -> Task1Response
	1, // 2: GRPCCleanService.Hello:input_type -> Empty
	2, // 3: GRPCCleanService.Tasks:input_type -> TasksRequest
	3, // 4: GRPCCleanService.Task1:input_type -> TaskRequest
	3, // 5: GRPCCleanService.Task2:input_type -> TaskRequest
	3, // 6: GRPCCleanService.Task3:input_type -> TaskRequest
	0, // 7: GRPCCleanService.Hello:output_type -> SuccessResponse
	8, // 8: GRPCCleanService.Tasks:output_type -> TasksResponse
	5, // 9: GRPCCleanService.Task1:output_type -> Task1Response
	6, // 10: GRPCCleanService.Task2:output_type -> Task2Response
	7, // 11: GRPCCleanService.Task3:output_type -> Task3Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_grpc_clean_proto_init() }
func file_proto_grpc_clean_proto_init() {
	if File_proto_grpc_clean_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_clean_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_proto_grpc_clean_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_grpc_clean_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksRequest); i {
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
		file_proto_grpc_clean_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_proto_grpc_clean_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_proto_grpc_clean_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task1Response); i {
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
		file_proto_grpc_clean_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task2Response); i {
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
		file_proto_grpc_clean_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task3Response); i {
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
		file_proto_grpc_clean_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksResponse); i {
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
			RawDescriptor: file_proto_grpc_clean_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_clean_proto_goTypes,
		DependencyIndexes: file_proto_grpc_clean_proto_depIdxs,
		MessageInfos:      file_proto_grpc_clean_proto_msgTypes,
	}.Build()
	File_proto_grpc_clean_proto = out.File
	file_proto_grpc_clean_proto_rawDesc = nil
	file_proto_grpc_clean_proto_goTypes = nil
	file_proto_grpc_clean_proto_depIdxs = nil
}
