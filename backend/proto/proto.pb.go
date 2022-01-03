// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: proto.proto

package backend

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

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName      string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName       string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email          string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	NumberOfTicket uint32 `protobuf:"varint,4,opt,name=numberOfTicket,proto3" json:"numberOfTicket,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserData) GetNumberOfTicket() uint32 {
	if x != nil {
		return x.NumberOfTicket
	}
	return 0
}

type UserDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData *UserData `protobuf:"bytes,1,opt,name=userData,proto3" json:"userData,omitempty"`
}

func (x *UserDataReq) Reset() {
	*x = UserDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataReq) ProtoMessage() {}

func (x *UserDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataReq.ProtoReflect.Descriptor instead.
func (*UserDataReq) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{1}
}

func (x *UserDataReq) GetUserData() *UserData {
	if x != nil {
		return x.UserData
	}
	return nil
}

type UserDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData *UserData `protobuf:"bytes,1,opt,name=userData,proto3" json:"userData,omitempty"`
}

func (x *UserDataRes) Reset() {
	*x = UserDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataRes) ProtoMessage() {}

func (x *UserDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataRes.ProtoReflect.Descriptor instead.
func (*UserDataRes) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{2}
}

func (x *UserDataRes) GetUserData() *UserData {
	if x != nil {
		return x.UserData
	}
	return nil
}

type GetOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetOrderReq) Reset() {
	*x = GetOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReq) ProtoMessage() {}

func (x *GetOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReq.ProtoReflect.Descriptor instead.
func (*GetOrderReq) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfTicketd uint32 `protobuf:"varint,1,opt,name=numberOfTicketd,proto3" json:"numberOfTicketd,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrderRequest) GetNumberOfTicketd() uint32 {
	if x != nil {
		return x.NumberOfTicketd
	}
	return 0
}

func (x *UpdateOrderRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DeleteOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted int64 `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteOrderRes) Reset() {
	*x = DeleteOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRes) ProtoMessage() {}

func (x *DeleteOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRes.ProtoReflect.Descriptor instead.
func (*DeleteOrderRes) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteOrderRes) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type GetAllOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllOrderReq) Reset() {
	*x = GetAllOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrderReq) ProtoMessage() {}

func (x *GetAllOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrderReq.ProtoReflect.Descriptor instead.
func (*GetAllOrderReq) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{6}
}

type GetAllOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData []*UserData `protobuf:"bytes,1,rep,name=userData,proto3" json:"userData,omitempty"`
}

func (x *GetAllOrderRes) Reset() {
	*x = GetAllOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrderRes) ProtoMessage() {}

func (x *GetAllOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrderRes.ProtoReflect.Descriptor instead.
func (*GetAllOrderRes) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllOrderRes) GetUserData() []*UserData {
	if x != nil {
		return x.UserData
	}
	return nil
}

type TicketNumberRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfTicketd []uint32 `protobuf:"varint,1,rep,packed,name=numberOfTicketd,proto3" json:"numberOfTicketd,omitempty"`
}

func (x *TicketNumberRes) Reset() {
	*x = TicketNumberRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketNumberRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketNumberRes) ProtoMessage() {}

func (x *TicketNumberRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketNumberRes.ProtoReflect.Descriptor instead.
func (*TicketNumberRes) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{8}
}

func (x *TicketNumberRes) GetNumberOfTicketd() []uint32 {
	if x != nil {
		return x.NumberOfTicketd
	}
	return nil
}

var File_proto_proto protoreflect.FileDescriptor

var file_proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x3c, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x54, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x10,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x3b, 0x0a, 0x0f, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x64, 0x32, 0x89,
	0x03, 0x0a, 0x0c, 0x47, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x6e, 0x61, 0x63, 0x68, 0x65,
	0x6d, 0x35, 0x35, 0x34, 0x2f, 0x47, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_proto_rawDescOnce sync.Once
	file_proto_proto_rawDescData = file_proto_proto_rawDesc
)

func file_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_rawDescData)
	})
	return file_proto_proto_rawDescData
}

var file_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_proto_goTypes = []interface{}{
	(*UserData)(nil),           // 0: backend.UserData
	(*UserDataReq)(nil),        // 1: backend.UserDataReq
	(*UserDataRes)(nil),        // 2: backend.UserDataRes
	(*GetOrderReq)(nil),        // 3: backend.GetOrderReq
	(*UpdateOrderRequest)(nil), // 4: backend.UpdateOrderRequest
	(*DeleteOrderRes)(nil),     // 5: backend.DeleteOrderRes
	(*GetAllOrderReq)(nil),     // 6: backend.GetAllOrderReq
	(*GetAllOrderRes)(nil),     // 7: backend.GetAllOrderRes
	(*TicketNumberRes)(nil),    // 8: backend.TicketNumberRes
}
var file_proto_proto_depIdxs = []int32{
	0, // 0: backend.UserDataReq.userData:type_name -> backend.UserData
	0, // 1: backend.UserDataRes.userData:type_name -> backend.UserData
	0, // 2: backend.GetAllOrderRes.userData:type_name -> backend.UserData
	1, // 3: backend.GoConference.PostNewOrder:input_type -> backend.UserDataReq
	3, // 4: backend.GoConference.GetOrder:input_type -> backend.GetOrderReq
	4, // 5: backend.GoConference.UpdateOrder:input_type -> backend.UpdateOrderRequest
	3, // 6: backend.GoConference.DeleteOrder:input_type -> backend.GetOrderReq
	6, // 7: backend.GoConference.GetAllOrder:input_type -> backend.GetAllOrderReq
	6, // 8: backend.GoConference.GetTicketNumber:input_type -> backend.GetAllOrderReq
	2, // 9: backend.GoConference.PostNewOrder:output_type -> backend.UserDataRes
	2, // 10: backend.GoConference.GetOrder:output_type -> backend.UserDataRes
	2, // 11: backend.GoConference.UpdateOrder:output_type -> backend.UserDataRes
	5, // 12: backend.GoConference.DeleteOrder:output_type -> backend.DeleteOrderRes
	7, // 13: backend.GoConference.GetAllOrder:output_type -> backend.GetAllOrderRes
	8, // 14: backend.GoConference.GetTicketNumber:output_type -> backend.TicketNumberRes
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_proto_init() }
func file_proto_proto_init() {
	if File_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataReq); i {
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
		file_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataRes); i {
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
		file_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReq); i {
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
		file_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderRequest); i {
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
		file_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderRes); i {
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
		file_proto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrderReq); i {
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
		file_proto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrderRes); i {
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
		file_proto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketNumberRes); i {
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
			RawDescriptor: file_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_msgTypes,
	}.Build()
	File_proto_proto = out.File
	file_proto_proto_rawDesc = nil
	file_proto_proto_goTypes = nil
	file_proto_proto_depIdxs = nil
}
