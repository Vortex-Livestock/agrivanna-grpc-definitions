// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: proto/livestock/v1/breeding_record.proto

package generated

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

type BreedingRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LivestockId   string                 `protobuf:"bytes,2,opt,name=livestock_id,json=livestockId,proto3" json:"livestock_id,omitempty"`
	BreedingDate  string                 `protobuf:"bytes,3,opt,name=breeding_date,json=breedingDate,proto3" json:"breeding_date,omitempty"`
	SireId        string                 `protobuf:"bytes,4,opt,name=sire_id,json=sireId,proto3" json:"sire_id,omitempty"`
	DamId         string                 `protobuf:"bytes,5,opt,name=dam_id,json=damId,proto3" json:"dam_id,omitempty"`
	Notes         string                 `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BreedingRecord) Reset() {
	*x = BreedingRecord{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BreedingRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreedingRecord) ProtoMessage() {}

func (x *BreedingRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreedingRecord.ProtoReflect.Descriptor instead.
func (*BreedingRecord) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{0}
}

func (x *BreedingRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BreedingRecord) GetLivestockId() string {
	if x != nil {
		return x.LivestockId
	}
	return ""
}

func (x *BreedingRecord) GetBreedingDate() string {
	if x != nil {
		return x.BreedingDate
	}
	return ""
}

func (x *BreedingRecord) GetSireId() string {
	if x != nil {
		return x.SireId
	}
	return ""
}

func (x *BreedingRecord) GetDamId() string {
	if x != nil {
		return x.DamId
	}
	return ""
}

func (x *BreedingRecord) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *BreedingRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BreedingRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateBreedingRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LivestockId   string                 `protobuf:"bytes,1,opt,name=livestock_id,json=livestockId,proto3" json:"livestock_id,omitempty"`
	BreedingDate  string                 `protobuf:"bytes,2,opt,name=breeding_date,json=breedingDate,proto3" json:"breeding_date,omitempty"`
	SireId        string                 `protobuf:"bytes,3,opt,name=sire_id,json=sireId,proto3" json:"sire_id,omitempty"`
	DamId         string                 `protobuf:"bytes,4,opt,name=dam_id,json=damId,proto3" json:"dam_id,omitempty"`
	Notes         string                 `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBreedingRecordRequest) Reset() {
	*x = CreateBreedingRecordRequest{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBreedingRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBreedingRecordRequest) ProtoMessage() {}

func (x *CreateBreedingRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBreedingRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateBreedingRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBreedingRecordRequest) GetLivestockId() string {
	if x != nil {
		return x.LivestockId
	}
	return ""
}

func (x *CreateBreedingRecordRequest) GetBreedingDate() string {
	if x != nil {
		return x.BreedingDate
	}
	return ""
}

func (x *CreateBreedingRecordRequest) GetSireId() string {
	if x != nil {
		return x.SireId
	}
	return ""
}

func (x *CreateBreedingRecordRequest) GetDamId() string {
	if x != nil {
		return x.DamId
	}
	return ""
}

func (x *CreateBreedingRecordRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type CreateBreedingRecordResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	BreedingRecord *BreedingRecord        `protobuf:"bytes,1,opt,name=breeding_record,json=breedingRecord,proto3" json:"breeding_record,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateBreedingRecordResponse) Reset() {
	*x = CreateBreedingRecordResponse{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBreedingRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBreedingRecordResponse) ProtoMessage() {}

func (x *CreateBreedingRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBreedingRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateBreedingRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBreedingRecordResponse) GetBreedingRecord() *BreedingRecord {
	if x != nil {
		return x.BreedingRecord
	}
	return nil
}

type GetBreedingRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBreedingRecordRequest) Reset() {
	*x = GetBreedingRecordRequest{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBreedingRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBreedingRecordRequest) ProtoMessage() {}

func (x *GetBreedingRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBreedingRecordRequest.ProtoReflect.Descriptor instead.
func (*GetBreedingRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetBreedingRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBreedingRecordResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	BreedingRecord *BreedingRecord        `protobuf:"bytes,1,opt,name=breeding_record,json=breedingRecord,proto3" json:"breeding_record,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetBreedingRecordResponse) Reset() {
	*x = GetBreedingRecordResponse{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBreedingRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBreedingRecordResponse) ProtoMessage() {}

func (x *GetBreedingRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBreedingRecordResponse.ProtoReflect.Descriptor instead.
func (*GetBreedingRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetBreedingRecordResponse) GetBreedingRecord() *BreedingRecord {
	if x != nil {
		return x.BreedingRecord
	}
	return nil
}

type UpdateBreedingRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LivestockId   string                 `protobuf:"bytes,2,opt,name=livestock_id,json=livestockId,proto3" json:"livestock_id,omitempty"`
	BreedingDate  string                 `protobuf:"bytes,3,opt,name=breeding_date,json=breedingDate,proto3" json:"breeding_date,omitempty"`
	SireId        string                 `protobuf:"bytes,4,opt,name=sire_id,json=sireId,proto3" json:"sire_id,omitempty"`
	DamId         string                 `protobuf:"bytes,5,opt,name=dam_id,json=damId,proto3" json:"dam_id,omitempty"`
	Notes         string                 `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBreedingRecordRequest) Reset() {
	*x = UpdateBreedingRecordRequest{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBreedingRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBreedingRecordRequest) ProtoMessage() {}

func (x *UpdateBreedingRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBreedingRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateBreedingRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBreedingRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBreedingRecordRequest) GetLivestockId() string {
	if x != nil {
		return x.LivestockId
	}
	return ""
}

func (x *UpdateBreedingRecordRequest) GetBreedingDate() string {
	if x != nil {
		return x.BreedingDate
	}
	return ""
}

func (x *UpdateBreedingRecordRequest) GetSireId() string {
	if x != nil {
		return x.SireId
	}
	return ""
}

func (x *UpdateBreedingRecordRequest) GetDamId() string {
	if x != nil {
		return x.DamId
	}
	return ""
}

func (x *UpdateBreedingRecordRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type UpdateBreedingRecordResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	BreedingRecord *BreedingRecord        `protobuf:"bytes,1,opt,name=breeding_record,json=breedingRecord,proto3" json:"breeding_record,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateBreedingRecordResponse) Reset() {
	*x = UpdateBreedingRecordResponse{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBreedingRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBreedingRecordResponse) ProtoMessage() {}

func (x *UpdateBreedingRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBreedingRecordResponse.ProtoReflect.Descriptor instead.
func (*UpdateBreedingRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBreedingRecordResponse) GetBreedingRecord() *BreedingRecord {
	if x != nil {
		return x.BreedingRecord
	}
	return nil
}

type DeleteBreedingRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBreedingRecordRequest) Reset() {
	*x = DeleteBreedingRecordRequest{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBreedingRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBreedingRecordRequest) ProtoMessage() {}

func (x *DeleteBreedingRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBreedingRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteBreedingRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBreedingRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBreedingRecordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBreedingRecordResponse) Reset() {
	*x = DeleteBreedingRecordResponse{}
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBreedingRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBreedingRecordResponse) ProtoMessage() {}

func (x *DeleteBreedingRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_breeding_record_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBreedingRecordResponse.ProtoReflect.Descriptor instead.
func (*DeleteBreedingRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_breeding_record_proto_rawDescGZIP(), []int{8}
}

var File_proto_livestock_v1_breeding_record_proto protoreflect.FileDescriptor

var file_proto_livestock_v1_breeding_record_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0xec,
	0x01, 0x0a, 0x0e, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x72, 0x65,
	0x65, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xab, 0x01,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x72, 0x65, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x64, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x62,
	0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0e, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x72, 0x65, 0x65, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0e,
	0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0xbb,
	0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x65, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x64, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x1c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f,
	0x62, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x65, 0x65, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0e, 0x62, 0x72, 0x65, 0x65, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x2d, 0x0a, 0x1b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfa, 0x03, 0x0a, 0x15, 0x42, 0x72, 0x65,
	0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x79, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x65, 0x65, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x79, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72,
	0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x72, 0x69, 0x76, 0x61, 0x6e, 0x6e, 0x61, 0x2f, 0x61, 0x67,
	0x72, 0x69, 0x76, 0x61, 0x6e, 0x6e, 0x61, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_livestock_v1_breeding_record_proto_rawDescOnce sync.Once
	file_proto_livestock_v1_breeding_record_proto_rawDescData = file_proto_livestock_v1_breeding_record_proto_rawDesc
)

func file_proto_livestock_v1_breeding_record_proto_rawDescGZIP() []byte {
	file_proto_livestock_v1_breeding_record_proto_rawDescOnce.Do(func() {
		file_proto_livestock_v1_breeding_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_livestock_v1_breeding_record_proto_rawDescData)
	})
	return file_proto_livestock_v1_breeding_record_proto_rawDescData
}

var file_proto_livestock_v1_breeding_record_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_livestock_v1_breeding_record_proto_goTypes = []any{
	(*BreedingRecord)(nil),               // 0: proto.livestock.v1.BreedingRecord
	(*CreateBreedingRecordRequest)(nil),  // 1: proto.livestock.v1.CreateBreedingRecordRequest
	(*CreateBreedingRecordResponse)(nil), // 2: proto.livestock.v1.CreateBreedingRecordResponse
	(*GetBreedingRecordRequest)(nil),     // 3: proto.livestock.v1.GetBreedingRecordRequest
	(*GetBreedingRecordResponse)(nil),    // 4: proto.livestock.v1.GetBreedingRecordResponse
	(*UpdateBreedingRecordRequest)(nil),  // 5: proto.livestock.v1.UpdateBreedingRecordRequest
	(*UpdateBreedingRecordResponse)(nil), // 6: proto.livestock.v1.UpdateBreedingRecordResponse
	(*DeleteBreedingRecordRequest)(nil),  // 7: proto.livestock.v1.DeleteBreedingRecordRequest
	(*DeleteBreedingRecordResponse)(nil), // 8: proto.livestock.v1.DeleteBreedingRecordResponse
}
var file_proto_livestock_v1_breeding_record_proto_depIdxs = []int32{
	0, // 0: proto.livestock.v1.CreateBreedingRecordResponse.breeding_record:type_name -> proto.livestock.v1.BreedingRecord
	0, // 1: proto.livestock.v1.GetBreedingRecordResponse.breeding_record:type_name -> proto.livestock.v1.BreedingRecord
	0, // 2: proto.livestock.v1.UpdateBreedingRecordResponse.breeding_record:type_name -> proto.livestock.v1.BreedingRecord
	1, // 3: proto.livestock.v1.BreedingRecordService.CreateBreedingRecord:input_type -> proto.livestock.v1.CreateBreedingRecordRequest
	3, // 4: proto.livestock.v1.BreedingRecordService.GetBreedingRecord:input_type -> proto.livestock.v1.GetBreedingRecordRequest
	5, // 5: proto.livestock.v1.BreedingRecordService.UpdateBreedingRecord:input_type -> proto.livestock.v1.UpdateBreedingRecordRequest
	7, // 6: proto.livestock.v1.BreedingRecordService.DeleteBreedingRecord:input_type -> proto.livestock.v1.DeleteBreedingRecordRequest
	2, // 7: proto.livestock.v1.BreedingRecordService.CreateBreedingRecord:output_type -> proto.livestock.v1.CreateBreedingRecordResponse
	4, // 8: proto.livestock.v1.BreedingRecordService.GetBreedingRecord:output_type -> proto.livestock.v1.GetBreedingRecordResponse
	6, // 9: proto.livestock.v1.BreedingRecordService.UpdateBreedingRecord:output_type -> proto.livestock.v1.UpdateBreedingRecordResponse
	8, // 10: proto.livestock.v1.BreedingRecordService.DeleteBreedingRecord:output_type -> proto.livestock.v1.DeleteBreedingRecordResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_livestock_v1_breeding_record_proto_init() }
func file_proto_livestock_v1_breeding_record_proto_init() {
	if File_proto_livestock_v1_breeding_record_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_livestock_v1_breeding_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_livestock_v1_breeding_record_proto_goTypes,
		DependencyIndexes: file_proto_livestock_v1_breeding_record_proto_depIdxs,
		MessageInfos:      file_proto_livestock_v1_breeding_record_proto_msgTypes,
	}.Build()
	File_proto_livestock_v1_breeding_record_proto = out.File
	file_proto_livestock_v1_breeding_record_proto_rawDesc = nil
	file_proto_livestock_v1_breeding_record_proto_goTypes = nil
	file_proto_livestock_v1_breeding_record_proto_depIdxs = nil
}
