// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: proto/v1/farm.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Farm struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location      string                 `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	OwnerId       string                 `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	ContactNumber string                 `protobuf:"bytes,6,opt,name=contact_number,json=contactNumber,proto3" json:"contact_number,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Farm) Reset() {
	*x = Farm{}
	mi := &file_proto_v1_farm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Farm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Farm) ProtoMessage() {}

func (x *Farm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Farm.ProtoReflect.Descriptor instead.
func (*Farm) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{0}
}

func (x *Farm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Farm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Farm) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Farm) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Farm) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Farm) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *Farm) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Farm) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateFarmRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location      string                 `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	OwnerId       string                 `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ContactNumber string                 `protobuf:"bytes,5,opt,name=contact_number,json=contactNumber,proto3" json:"contact_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFarmRequest) Reset() {
	*x = CreateFarmRequest{}
	mi := &file_proto_v1_farm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFarmRequest) ProtoMessage() {}

func (x *CreateFarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFarmRequest.ProtoReflect.Descriptor instead.
func (*CreateFarmRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFarmRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFarmRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateFarmRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateFarmRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateFarmRequest) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

type CreateFarmResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Farm          *Farm                  `protobuf:"bytes,1,opt,name=farm,proto3" json:"farm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFarmResponse) Reset() {
	*x = CreateFarmResponse{}
	mi := &file_proto_v1_farm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFarmResponse) ProtoMessage() {}

func (x *CreateFarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFarmResponse.ProtoReflect.Descriptor instead.
func (*CreateFarmResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFarmResponse) GetFarm() *Farm {
	if x != nil {
		return x.Farm
	}
	return nil
}

type GetFarmRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFarmRequest) Reset() {
	*x = GetFarmRequest{}
	mi := &file_proto_v1_farm_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFarmRequest) ProtoMessage() {}

func (x *GetFarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFarmRequest.ProtoReflect.Descriptor instead.
func (*GetFarmRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{3}
}

func (x *GetFarmRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFarmResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Farm          *Farm                  `protobuf:"bytes,1,opt,name=farm,proto3" json:"farm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFarmResponse) Reset() {
	*x = GetFarmResponse{}
	mi := &file_proto_v1_farm_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFarmResponse) ProtoMessage() {}

func (x *GetFarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFarmResponse.ProtoReflect.Descriptor instead.
func (*GetFarmResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{4}
}

func (x *GetFarmResponse) GetFarm() *Farm {
	if x != nil {
		return x.Farm
	}
	return nil
}

type GetFarmByOwnerIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerId       string                 `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFarmByOwnerIdRequest) Reset() {
	*x = GetFarmByOwnerIdRequest{}
	mi := &file_proto_v1_farm_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFarmByOwnerIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFarmByOwnerIdRequest) ProtoMessage() {}

func (x *GetFarmByOwnerIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFarmByOwnerIdRequest.ProtoReflect.Descriptor instead.
func (*GetFarmByOwnerIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{5}
}

func (x *GetFarmByOwnerIdRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetFarmByOwnerIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Farms         []*Farm                `protobuf:"bytes,1,rep,name=farms,proto3" json:"farms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFarmByOwnerIdResponse) Reset() {
	*x = GetFarmByOwnerIdResponse{}
	mi := &file_proto_v1_farm_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFarmByOwnerIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFarmByOwnerIdResponse) ProtoMessage() {}

func (x *GetFarmByOwnerIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFarmByOwnerIdResponse.ProtoReflect.Descriptor instead.
func (*GetFarmByOwnerIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{6}
}

func (x *GetFarmByOwnerIdResponse) GetFarms() []*Farm {
	if x != nil {
		return x.Farms
	}
	return nil
}

type UpdateFarmRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location      string                 `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	OwnerId       string                 `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	ContactNumber string                 `protobuf:"bytes,6,opt,name=contact_number,json=contactNumber,proto3" json:"contact_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFarmRequest) Reset() {
	*x = UpdateFarmRequest{}
	mi := &file_proto_v1_farm_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFarmRequest) ProtoMessage() {}

func (x *UpdateFarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFarmRequest.ProtoReflect.Descriptor instead.
func (*UpdateFarmRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateFarmRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateFarmRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFarmRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateFarmRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *UpdateFarmRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateFarmRequest) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

type UpdateFarmResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Farm          *Farm                  `protobuf:"bytes,1,opt,name=farm,proto3" json:"farm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFarmResponse) Reset() {
	*x = UpdateFarmResponse{}
	mi := &file_proto_v1_farm_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFarmResponse) ProtoMessage() {}

func (x *UpdateFarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFarmResponse.ProtoReflect.Descriptor instead.
func (*UpdateFarmResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateFarmResponse) GetFarm() *Farm {
	if x != nil {
		return x.Farm
	}
	return nil
}

type DeleteFarmRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFarmRequest) Reset() {
	*x = DeleteFarmRequest{}
	mi := &file_proto_v1_farm_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFarmRequest) ProtoMessage() {}

func (x *DeleteFarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFarmRequest.ProtoReflect.Descriptor instead.
func (*DeleteFarmRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteFarmRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFarmResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFarmResponse) Reset() {
	*x = DeleteFarmResponse{}
	mi := &file_proto_v1_farm_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFarmResponse) ProtoMessage() {}

func (x *DeleteFarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_farm_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFarmResponse.ProtoReflect.Descriptor instead.
func (*DeleteFarmResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_farm_proto_rawDescGZIP(), []int{10}
}

var File_proto_v1_farm_proto protoreflect.FileDescriptor

var file_proto_v1_farm_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x94, 0x02, 0x0a, 0x04, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x66,
	0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x22,
	0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x72, 0x6d, 0x52, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x22, 0x34, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x72, 0x6d, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x66, 0x61,
	0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x05, 0x66, 0x61, 0x72, 0x6d, 0x73,
	0x22, 0xab, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x72, 0x6d, 0x52, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x83, 0x03, 0x0a, 0x0b, 0x46, 0x61, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x72,
	0x6d, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x72, 0x6d, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x72, 0x69, 0x76, 0x61, 0x6e, 0x6e,
	0x61, 0x2f, 0x61, 0x67, 0x72, 0x69, 0x76, 0x61, 0x6e, 0x6e, 0x61, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_farm_proto_rawDescOnce sync.Once
	file_proto_v1_farm_proto_rawDescData = file_proto_v1_farm_proto_rawDesc
)

func file_proto_v1_farm_proto_rawDescGZIP() []byte {
	file_proto_v1_farm_proto_rawDescOnce.Do(func() {
		file_proto_v1_farm_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_farm_proto_rawDescData)
	})
	return file_proto_v1_farm_proto_rawDescData
}

var file_proto_v1_farm_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_v1_farm_proto_goTypes = []any{
	(*Farm)(nil),                     // 0: proto.v1.Farm
	(*CreateFarmRequest)(nil),        // 1: proto.v1.CreateFarmRequest
	(*CreateFarmResponse)(nil),       // 2: proto.v1.CreateFarmResponse
	(*GetFarmRequest)(nil),           // 3: proto.v1.GetFarmRequest
	(*GetFarmResponse)(nil),          // 4: proto.v1.GetFarmResponse
	(*GetFarmByOwnerIdRequest)(nil),  // 5: proto.v1.GetFarmByOwnerIdRequest
	(*GetFarmByOwnerIdResponse)(nil), // 6: proto.v1.GetFarmByOwnerIdResponse
	(*UpdateFarmRequest)(nil),        // 7: proto.v1.UpdateFarmRequest
	(*UpdateFarmResponse)(nil),       // 8: proto.v1.UpdateFarmResponse
	(*DeleteFarmRequest)(nil),        // 9: proto.v1.DeleteFarmRequest
	(*DeleteFarmResponse)(nil),       // 10: proto.v1.DeleteFarmResponse
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_proto_v1_farm_proto_depIdxs = []int32{
	11, // 0: proto.v1.Farm.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: proto.v1.Farm.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: proto.v1.CreateFarmResponse.farm:type_name -> proto.v1.Farm
	0,  // 3: proto.v1.GetFarmResponse.farm:type_name -> proto.v1.Farm
	0,  // 4: proto.v1.GetFarmByOwnerIdResponse.farms:type_name -> proto.v1.Farm
	0,  // 5: proto.v1.UpdateFarmResponse.farm:type_name -> proto.v1.Farm
	1,  // 6: proto.v1.FarmService.CreateFarm:input_type -> proto.v1.CreateFarmRequest
	3,  // 7: proto.v1.FarmService.GetFarm:input_type -> proto.v1.GetFarmRequest
	5,  // 8: proto.v1.FarmService.GetFarmByOwnerId:input_type -> proto.v1.GetFarmByOwnerIdRequest
	7,  // 9: proto.v1.FarmService.UpdateFarm:input_type -> proto.v1.UpdateFarmRequest
	9,  // 10: proto.v1.FarmService.DeleteFarm:input_type -> proto.v1.DeleteFarmRequest
	2,  // 11: proto.v1.FarmService.CreateFarm:output_type -> proto.v1.CreateFarmResponse
	4,  // 12: proto.v1.FarmService.GetFarm:output_type -> proto.v1.GetFarmResponse
	6,  // 13: proto.v1.FarmService.GetFarmByOwnerId:output_type -> proto.v1.GetFarmByOwnerIdResponse
	8,  // 14: proto.v1.FarmService.UpdateFarm:output_type -> proto.v1.UpdateFarmResponse
	10, // 15: proto.v1.FarmService.DeleteFarm:output_type -> proto.v1.DeleteFarmResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_v1_farm_proto_init() }
func file_proto_v1_farm_proto_init() {
	if File_proto_v1_farm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_farm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_farm_proto_goTypes,
		DependencyIndexes: file_proto_v1_farm_proto_depIdxs,
		MessageInfos:      file_proto_v1_farm_proto_msgTypes,
	}.Build()
	File_proto_v1_farm_proto = out.File
	file_proto_v1_farm_proto_rawDesc = nil
	file_proto_v1_farm_proto_goTypes = nil
	file_proto_v1_farm_proto_depIdxs = nil
}
