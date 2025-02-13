// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: proto/livestock/v1/weight_record.proto

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

type WeightRecordUnit int32

const (
	WeightRecordUnit_WEIGHT_RECORD_UNIT_UNSPECIFIED WeightRecordUnit = 0
	WeightRecordUnit_WEIGHT_RECORD_UNIT_GRAMS       WeightRecordUnit = 1
	WeightRecordUnit_WEIGHT_RECORD_UNIT_KG          WeightRecordUnit = 2
	WeightRecordUnit_WEIGHT_RECORD_UNIT_LB          WeightRecordUnit = 3
)

// Enum value maps for WeightRecordUnit.
var (
	WeightRecordUnit_name = map[int32]string{
		0: "WEIGHT_RECORD_UNIT_UNSPECIFIED",
		1: "WEIGHT_RECORD_UNIT_GRAMS",
		2: "WEIGHT_RECORD_UNIT_KG",
		3: "WEIGHT_RECORD_UNIT_LB",
	}
	WeightRecordUnit_value = map[string]int32{
		"WEIGHT_RECORD_UNIT_UNSPECIFIED": 0,
		"WEIGHT_RECORD_UNIT_GRAMS":       1,
		"WEIGHT_RECORD_UNIT_KG":          2,
		"WEIGHT_RECORD_UNIT_LB":          3,
	}
)

func (x WeightRecordUnit) Enum() *WeightRecordUnit {
	p := new(WeightRecordUnit)
	*p = x
	return p
}

func (x WeightRecordUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeightRecordUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_livestock_v1_weight_record_proto_enumTypes[0].Descriptor()
}

func (WeightRecordUnit) Type() protoreflect.EnumType {
	return &file_proto_livestock_v1_weight_record_proto_enumTypes[0]
}

func (x WeightRecordUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeightRecordUnit.Descriptor instead.
func (WeightRecordUnit) EnumDescriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{0}
}

type WeightRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LivestockId   string                 `protobuf:"bytes,2,opt,name=livestock_id,json=livestockId,proto3" json:"livestock_id,omitempty"`
	RecordedAt    string                 `protobuf:"bytes,3,opt,name=recorded_at,json=recordedAt,proto3" json:"recorded_at,omitempty"`
	Weight        float32                `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Unit          WeightRecordUnit       `protobuf:"varint,5,opt,name=unit,proto3,enum=proto.livestock.v1.WeightRecordUnit" json:"unit,omitempty"`
	Notes         string                 `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WeightRecord) Reset() {
	*x = WeightRecord{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeightRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightRecord) ProtoMessage() {}

func (x *WeightRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightRecord.ProtoReflect.Descriptor instead.
func (*WeightRecord) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{0}
}

func (x *WeightRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WeightRecord) GetLivestockId() string {
	if x != nil {
		return x.LivestockId
	}
	return ""
}

func (x *WeightRecord) GetRecordedAt() string {
	if x != nil {
		return x.RecordedAt
	}
	return ""
}

func (x *WeightRecord) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *WeightRecord) GetUnit() WeightRecordUnit {
	if x != nil {
		return x.Unit
	}
	return WeightRecordUnit_WEIGHT_RECORD_UNIT_UNSPECIFIED
}

func (x *WeightRecord) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *WeightRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WeightRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateWeightRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LivestockId   string                 `protobuf:"bytes,1,opt,name=livestock_id,json=livestockId,proto3" json:"livestock_id,omitempty"`
	RecordedAt    string                 `protobuf:"bytes,2,opt,name=recorded_at,json=recordedAt,proto3" json:"recorded_at,omitempty"`
	Weight        float32                `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Unit          WeightRecordUnit       `protobuf:"varint,4,opt,name=unit,proto3,enum=proto.livestock.v1.WeightRecordUnit" json:"unit,omitempty"`
	Notes         string                 `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWeightRecordRequest) Reset() {
	*x = CreateWeightRecordRequest{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWeightRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWeightRecordRequest) ProtoMessage() {}

func (x *CreateWeightRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWeightRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateWeightRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWeightRecordRequest) GetLivestockId() string {
	if x != nil {
		return x.LivestockId
	}
	return ""
}

func (x *CreateWeightRecordRequest) GetRecordedAt() string {
	if x != nil {
		return x.RecordedAt
	}
	return ""
}

func (x *CreateWeightRecordRequest) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CreateWeightRecordRequest) GetUnit() WeightRecordUnit {
	if x != nil {
		return x.Unit
	}
	return WeightRecordUnit_WEIGHT_RECORD_UNIT_UNSPECIFIED
}

func (x *CreateWeightRecordRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type CreateWeightRecordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WeightRecord  *WeightRecord          `protobuf:"bytes,1,opt,name=weight_record,json=weightRecord,proto3" json:"weight_record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWeightRecordResponse) Reset() {
	*x = CreateWeightRecordResponse{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWeightRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWeightRecordResponse) ProtoMessage() {}

func (x *CreateWeightRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWeightRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateWeightRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWeightRecordResponse) GetWeightRecord() *WeightRecord {
	if x != nil {
		return x.WeightRecord
	}
	return nil
}

type GetWeightRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWeightRecordRequest) Reset() {
	*x = GetWeightRecordRequest{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWeightRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeightRecordRequest) ProtoMessage() {}

func (x *GetWeightRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeightRecordRequest.ProtoReflect.Descriptor instead.
func (*GetWeightRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetWeightRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWeightRecordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WeightRecord  *WeightRecord          `protobuf:"bytes,1,opt,name=weight_record,json=weightRecord,proto3" json:"weight_record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWeightRecordResponse) Reset() {
	*x = GetWeightRecordResponse{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWeightRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeightRecordResponse) ProtoMessage() {}

func (x *GetWeightRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeightRecordResponse.ProtoReflect.Descriptor instead.
func (*GetWeightRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetWeightRecordResponse) GetWeightRecord() *WeightRecord {
	if x != nil {
		return x.WeightRecord
	}
	return nil
}

type UpdateWeightRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LivestockId   string                 `protobuf:"bytes,2,opt,name=livestock_id,json=livestockId,proto3" json:"livestock_id,omitempty"`
	RecordedAt    string                 `protobuf:"bytes,3,opt,name=recorded_at,json=recordedAt,proto3" json:"recorded_at,omitempty"`
	Weight        float32                `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Unit          WeightRecordUnit       `protobuf:"varint,5,opt,name=unit,proto3,enum=proto.livestock.v1.WeightRecordUnit" json:"unit,omitempty"`
	Notes         string                 `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWeightRecordRequest) Reset() {
	*x = UpdateWeightRecordRequest{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWeightRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWeightRecordRequest) ProtoMessage() {}

func (x *UpdateWeightRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWeightRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateWeightRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateWeightRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateWeightRecordRequest) GetLivestockId() string {
	if x != nil {
		return x.LivestockId
	}
	return ""
}

func (x *UpdateWeightRecordRequest) GetRecordedAt() string {
	if x != nil {
		return x.RecordedAt
	}
	return ""
}

func (x *UpdateWeightRecordRequest) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UpdateWeightRecordRequest) GetUnit() WeightRecordUnit {
	if x != nil {
		return x.Unit
	}
	return WeightRecordUnit_WEIGHT_RECORD_UNIT_UNSPECIFIED
}

func (x *UpdateWeightRecordRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type UpdateWeightRecordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WeightRecord  *WeightRecord          `protobuf:"bytes,1,opt,name=weight_record,json=weightRecord,proto3" json:"weight_record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWeightRecordResponse) Reset() {
	*x = UpdateWeightRecordResponse{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWeightRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWeightRecordResponse) ProtoMessage() {}

func (x *UpdateWeightRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWeightRecordResponse.ProtoReflect.Descriptor instead.
func (*UpdateWeightRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateWeightRecordResponse) GetWeightRecord() *WeightRecord {
	if x != nil {
		return x.WeightRecord
	}
	return nil
}

type DeleteWeightRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWeightRecordRequest) Reset() {
	*x = DeleteWeightRecordRequest{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWeightRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWeightRecordRequest) ProtoMessage() {}

func (x *DeleteWeightRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWeightRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteWeightRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteWeightRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteWeightRecordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWeightRecordResponse) Reset() {
	*x = DeleteWeightRecordResponse{}
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWeightRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWeightRecordResponse) ProtoMessage() {}

func (x *DeleteWeightRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_livestock_v1_weight_record_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWeightRecordResponse.ProtoReflect.Descriptor instead.
func (*DeleteWeightRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_livestock_v1_weight_record_proto_rawDescGZIP(), []int{8}
}

var File_proto_livestock_v1_weight_record_proto protoreflect.FileDescriptor

var file_proto_livestock_v1_weight_record_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x88, 0x02, 0x0a,
	0x0c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76,
	0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x38, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x22, 0x63, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x60, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x38, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x0c, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c,
	0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x8a, 0x01, 0x0a,
	0x10, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x69,
	0x74, 0x12, 0x22, 0x0a, 0x1e, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x47, 0x52, 0x41, 0x4d,
	0x53, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4b, 0x47, 0x10, 0x02, 0x12, 0x19,
	0x0a, 0x15, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4c, 0x42, 0x10, 0x03, 0x32, 0xe0, 0x03, 0x0a, 0x13, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x73, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x73, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x72, 0x69, 0x76,
	0x61, 0x6e, 0x6e, 0x61, 0x2f, 0x61, 0x67, 0x72, 0x69, 0x76, 0x61, 0x6e, 0x6e, 0x61, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_livestock_v1_weight_record_proto_rawDescOnce sync.Once
	file_proto_livestock_v1_weight_record_proto_rawDescData = file_proto_livestock_v1_weight_record_proto_rawDesc
)

func file_proto_livestock_v1_weight_record_proto_rawDescGZIP() []byte {
	file_proto_livestock_v1_weight_record_proto_rawDescOnce.Do(func() {
		file_proto_livestock_v1_weight_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_livestock_v1_weight_record_proto_rawDescData)
	})
	return file_proto_livestock_v1_weight_record_proto_rawDescData
}

var file_proto_livestock_v1_weight_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_livestock_v1_weight_record_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_livestock_v1_weight_record_proto_goTypes = []any{
	(WeightRecordUnit)(0),              // 0: proto.livestock.v1.WeightRecordUnit
	(*WeightRecord)(nil),               // 1: proto.livestock.v1.WeightRecord
	(*CreateWeightRecordRequest)(nil),  // 2: proto.livestock.v1.CreateWeightRecordRequest
	(*CreateWeightRecordResponse)(nil), // 3: proto.livestock.v1.CreateWeightRecordResponse
	(*GetWeightRecordRequest)(nil),     // 4: proto.livestock.v1.GetWeightRecordRequest
	(*GetWeightRecordResponse)(nil),    // 5: proto.livestock.v1.GetWeightRecordResponse
	(*UpdateWeightRecordRequest)(nil),  // 6: proto.livestock.v1.UpdateWeightRecordRequest
	(*UpdateWeightRecordResponse)(nil), // 7: proto.livestock.v1.UpdateWeightRecordResponse
	(*DeleteWeightRecordRequest)(nil),  // 8: proto.livestock.v1.DeleteWeightRecordRequest
	(*DeleteWeightRecordResponse)(nil), // 9: proto.livestock.v1.DeleteWeightRecordResponse
}
var file_proto_livestock_v1_weight_record_proto_depIdxs = []int32{
	0,  // 0: proto.livestock.v1.WeightRecord.unit:type_name -> proto.livestock.v1.WeightRecordUnit
	0,  // 1: proto.livestock.v1.CreateWeightRecordRequest.unit:type_name -> proto.livestock.v1.WeightRecordUnit
	1,  // 2: proto.livestock.v1.CreateWeightRecordResponse.weight_record:type_name -> proto.livestock.v1.WeightRecord
	1,  // 3: proto.livestock.v1.GetWeightRecordResponse.weight_record:type_name -> proto.livestock.v1.WeightRecord
	0,  // 4: proto.livestock.v1.UpdateWeightRecordRequest.unit:type_name -> proto.livestock.v1.WeightRecordUnit
	1,  // 5: proto.livestock.v1.UpdateWeightRecordResponse.weight_record:type_name -> proto.livestock.v1.WeightRecord
	2,  // 6: proto.livestock.v1.WeightRecordService.CreateWeightRecord:input_type -> proto.livestock.v1.CreateWeightRecordRequest
	4,  // 7: proto.livestock.v1.WeightRecordService.GetWeightRecord:input_type -> proto.livestock.v1.GetWeightRecordRequest
	6,  // 8: proto.livestock.v1.WeightRecordService.UpdateWeightRecord:input_type -> proto.livestock.v1.UpdateWeightRecordRequest
	8,  // 9: proto.livestock.v1.WeightRecordService.DeleteWeightRecord:input_type -> proto.livestock.v1.DeleteWeightRecordRequest
	3,  // 10: proto.livestock.v1.WeightRecordService.CreateWeightRecord:output_type -> proto.livestock.v1.CreateWeightRecordResponse
	5,  // 11: proto.livestock.v1.WeightRecordService.GetWeightRecord:output_type -> proto.livestock.v1.GetWeightRecordResponse
	7,  // 12: proto.livestock.v1.WeightRecordService.UpdateWeightRecord:output_type -> proto.livestock.v1.UpdateWeightRecordResponse
	9,  // 13: proto.livestock.v1.WeightRecordService.DeleteWeightRecord:output_type -> proto.livestock.v1.DeleteWeightRecordResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_livestock_v1_weight_record_proto_init() }
func file_proto_livestock_v1_weight_record_proto_init() {
	if File_proto_livestock_v1_weight_record_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_livestock_v1_weight_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_livestock_v1_weight_record_proto_goTypes,
		DependencyIndexes: file_proto_livestock_v1_weight_record_proto_depIdxs,
		EnumInfos:         file_proto_livestock_v1_weight_record_proto_enumTypes,
		MessageInfos:      file_proto_livestock_v1_weight_record_proto_msgTypes,
	}.Build()
	File_proto_livestock_v1_weight_record_proto = out.File
	file_proto_livestock_v1_weight_record_proto_rawDesc = nil
	file_proto_livestock_v1_weight_record_proto_goTypes = nil
	file_proto_livestock_v1_weight_record_proto_depIdxs = nil
}
