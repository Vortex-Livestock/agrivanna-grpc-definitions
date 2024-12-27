// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: production_record.proto

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

type ProductionRecordType int32

const (
	ProductionRecordType_PRODUCTION_RECORD_UNKNOWN ProductionRecordType = 0
	ProductionRecordType_PRODUCTION_RECORD_MILK    ProductionRecordType = 1
	ProductionRecordType_PRODUCTION_RECORD_CHEESE  ProductionRecordType = 2
	ProductionRecordType_PRODUCTION_RECORD_YOGURT  ProductionRecordType = 3
)

// Enum value maps for ProductionRecordType.
var (
	ProductionRecordType_name = map[int32]string{
		0: "PRODUCTION_RECORD_UNKNOWN",
		1: "PRODUCTION_RECORD_MILK",
		2: "PRODUCTION_RECORD_CHEESE",
		3: "PRODUCTION_RECORD_YOGURT",
	}
	ProductionRecordType_value = map[string]int32{
		"PRODUCTION_RECORD_UNKNOWN": 0,
		"PRODUCTION_RECORD_MILK":    1,
		"PRODUCTION_RECORD_CHEESE":  2,
		"PRODUCTION_RECORD_YOGURT":  3,
	}
)

func (x ProductionRecordType) Enum() *ProductionRecordType {
	p := new(ProductionRecordType)
	*p = x
	return p
}

func (x ProductionRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductionRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_production_record_proto_enumTypes[0].Descriptor()
}

func (ProductionRecordType) Type() protoreflect.EnumType {
	return &file_production_record_proto_enumTypes[0]
}

func (x ProductionRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductionRecordType.Descriptor instead.
func (ProductionRecordType) EnumDescriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{0}
}

type ProductionRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FarmId        string                 `protobuf:"bytes,2,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty"`
	AnimalId      string                 `protobuf:"bytes,3,opt,name=animal_id,json=animalId,proto3" json:"animal_id,omitempty"`
	Type          ProductionRecordType   `protobuf:"varint,4,opt,name=type,proto3,enum=proto.ProductionRecordType" json:"type,omitempty"`
	Quantity      string                 `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          string                 `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductionRecord) Reset() {
	*x = ProductionRecord{}
	mi := &file_production_record_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductionRecord) ProtoMessage() {}

func (x *ProductionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductionRecord.ProtoReflect.Descriptor instead.
func (*ProductionRecord) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{0}
}

func (x *ProductionRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductionRecord) GetFarmId() string {
	if x != nil {
		return x.FarmId
	}
	return ""
}

func (x *ProductionRecord) GetAnimalId() string {
	if x != nil {
		return x.AnimalId
	}
	return ""
}

func (x *ProductionRecord) GetType() ProductionRecordType {
	if x != nil {
		return x.Type
	}
	return ProductionRecordType_PRODUCTION_RECORD_UNKNOWN
}

func (x *ProductionRecord) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *ProductionRecord) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ProductionRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ProductionRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetProductionRecordByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductionRecordByIdRequest) Reset() {
	*x = GetProductionRecordByIdRequest{}
	mi := &file_production_record_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductionRecordByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionRecordByIdRequest) ProtoMessage() {}

func (x *GetProductionRecordByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionRecordByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProductionRecordByIdRequest) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductionRecordByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductionRecordsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FarmId        string                 `protobuf:"bytes,1,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductionRecordsRequest) Reset() {
	*x = GetProductionRecordsRequest{}
	mi := &file_production_record_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductionRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionRecordsRequest) ProtoMessage() {}

func (x *GetProductionRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionRecordsRequest.ProtoReflect.Descriptor instead.
func (*GetProductionRecordsRequest) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductionRecordsRequest) GetFarmId() string {
	if x != nil {
		return x.FarmId
	}
	return ""
}

type GetProductionRecordsResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ProductionRecords []*ProductionRecord    `protobuf:"bytes,1,rep,name=production_records,json=productionRecords,proto3" json:"production_records,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetProductionRecordsResponse) Reset() {
	*x = GetProductionRecordsResponse{}
	mi := &file_production_record_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductionRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionRecordsResponse) ProtoMessage() {}

func (x *GetProductionRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionRecordsResponse.ProtoReflect.Descriptor instead.
func (*GetProductionRecordsResponse) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductionRecordsResponse) GetProductionRecords() []*ProductionRecord {
	if x != nil {
		return x.ProductionRecords
	}
	return nil
}

type CreateProductionRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FarmId        string                 `protobuf:"bytes,1,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty"`
	AnimalId      string                 `protobuf:"bytes,2,opt,name=animal_id,json=animalId,proto3" json:"animal_id,omitempty"`
	Type          ProductionRecordType   `protobuf:"varint,3,opt,name=type,proto3,enum=proto.ProductionRecordType" json:"type,omitempty"`
	Quantity      string                 `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          string                 `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductionRecordRequest) Reset() {
	*x = CreateProductionRecordRequest{}
	mi := &file_production_record_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductionRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductionRecordRequest) ProtoMessage() {}

func (x *CreateProductionRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductionRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateProductionRecordRequest) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProductionRecordRequest) GetFarmId() string {
	if x != nil {
		return x.FarmId
	}
	return ""
}

func (x *CreateProductionRecordRequest) GetAnimalId() string {
	if x != nil {
		return x.AnimalId
	}
	return ""
}

func (x *CreateProductionRecordRequest) GetType() ProductionRecordType {
	if x != nil {
		return x.Type
	}
	return ProductionRecordType_PRODUCTION_RECORD_UNKNOWN
}

func (x *CreateProductionRecordRequest) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *CreateProductionRecordRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type UpdateProductionRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FarmId        string                 `protobuf:"bytes,2,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty"`
	AnimalId      string                 `protobuf:"bytes,3,opt,name=animal_id,json=animalId,proto3" json:"animal_id,omitempty"`
	Type          ProductionRecordType   `protobuf:"varint,4,opt,name=type,proto3,enum=proto.ProductionRecordType" json:"type,omitempty"`
	Quantity      string                 `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          string                 `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductionRecordRequest) Reset() {
	*x = UpdateProductionRecordRequest{}
	mi := &file_production_record_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductionRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductionRecordRequest) ProtoMessage() {}

func (x *UpdateProductionRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductionRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductionRecordRequest) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProductionRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductionRecordRequest) GetFarmId() string {
	if x != nil {
		return x.FarmId
	}
	return ""
}

func (x *UpdateProductionRecordRequest) GetAnimalId() string {
	if x != nil {
		return x.AnimalId
	}
	return ""
}

func (x *UpdateProductionRecordRequest) GetType() ProductionRecordType {
	if x != nil {
		return x.Type
	}
	return ProductionRecordType_PRODUCTION_RECORD_UNKNOWN
}

func (x *UpdateProductionRecordRequest) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *UpdateProductionRecordRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type DeleteProductionRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductionRecordRequest) Reset() {
	*x = DeleteProductionRecordRequest{}
	mi := &file_production_record_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductionRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductionRecordRequest) ProtoMessage() {}

func (x *DeleteProductionRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductionRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductionRecordRequest) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProductionRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductionRecordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductionRecordResponse) Reset() {
	*x = DeleteProductionRecordResponse{}
	mi := &file_production_record_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductionRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductionRecordResponse) ProtoMessage() {}

func (x *DeleteProductionRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_record_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductionRecordResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductionRecordResponse) Descriptor() ([]byte, []int) {
	return file_production_record_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProductionRecordResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_production_record_proto protoreflect.FileDescriptor

var file_production_record_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf7, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x30, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61,
	0x72, 0x6d, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xb6, 0x01, 0x0a,
	0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x69, 0x6d,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x61, 0x72, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x2f,
	0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x30, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x2a, 0x8d, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d,
	0x49, 0x4c, 0x4b, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x48, 0x45, 0x45, 0x53,
	0x45, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x59, 0x4f, 0x47, 0x55, 0x52, 0x54, 0x10,
	0x03, 0x32, 0xee, 0x03, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x57, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x65, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x67, 0x72, 0x69, 0x76, 0x61, 0x6e, 0x6e, 0x61, 0x2f, 0x61, 0x67, 0x72, 0x69, 0x76,
	0x61, 0x6e, 0x6e, 0x61, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_production_record_proto_rawDescOnce sync.Once
	file_production_record_proto_rawDescData = file_production_record_proto_rawDesc
)

func file_production_record_proto_rawDescGZIP() []byte {
	file_production_record_proto_rawDescOnce.Do(func() {
		file_production_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_production_record_proto_rawDescData)
	})
	return file_production_record_proto_rawDescData
}

var file_production_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_production_record_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_production_record_proto_goTypes = []any{
	(ProductionRecordType)(0),              // 0: proto.ProductionRecordType
	(*ProductionRecord)(nil),               // 1: proto.ProductionRecord
	(*GetProductionRecordByIdRequest)(nil), // 2: proto.GetProductionRecordByIdRequest
	(*GetProductionRecordsRequest)(nil),    // 3: proto.GetProductionRecordsRequest
	(*GetProductionRecordsResponse)(nil),   // 4: proto.GetProductionRecordsResponse
	(*CreateProductionRecordRequest)(nil),  // 5: proto.CreateProductionRecordRequest
	(*UpdateProductionRecordRequest)(nil),  // 6: proto.UpdateProductionRecordRequest
	(*DeleteProductionRecordRequest)(nil),  // 7: proto.DeleteProductionRecordRequest
	(*DeleteProductionRecordResponse)(nil), // 8: proto.DeleteProductionRecordResponse
}
var file_production_record_proto_depIdxs = []int32{
	0, // 0: proto.ProductionRecord.type:type_name -> proto.ProductionRecordType
	1, // 1: proto.GetProductionRecordsResponse.production_records:type_name -> proto.ProductionRecord
	0, // 2: proto.CreateProductionRecordRequest.type:type_name -> proto.ProductionRecordType
	0, // 3: proto.UpdateProductionRecordRequest.type:type_name -> proto.ProductionRecordType
	2, // 4: proto.ProductionRecordService.GetProductionRecordById:input_type -> proto.GetProductionRecordByIdRequest
	3, // 5: proto.ProductionRecordService.GetProductionRecords:input_type -> proto.GetProductionRecordsRequest
	5, // 6: proto.ProductionRecordService.CreateProductionRecord:input_type -> proto.CreateProductionRecordRequest
	6, // 7: proto.ProductionRecordService.UpdateProductionRecord:input_type -> proto.UpdateProductionRecordRequest
	7, // 8: proto.ProductionRecordService.DeleteProductionRecord:input_type -> proto.DeleteProductionRecordRequest
	1, // 9: proto.ProductionRecordService.GetProductionRecordById:output_type -> proto.ProductionRecord
	4, // 10: proto.ProductionRecordService.GetProductionRecords:output_type -> proto.GetProductionRecordsResponse
	1, // 11: proto.ProductionRecordService.CreateProductionRecord:output_type -> proto.ProductionRecord
	1, // 12: proto.ProductionRecordService.UpdateProductionRecord:output_type -> proto.ProductionRecord
	8, // 13: proto.ProductionRecordService.DeleteProductionRecord:output_type -> proto.DeleteProductionRecordResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_production_record_proto_init() }
func file_production_record_proto_init() {
	if File_production_record_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_production_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_production_record_proto_goTypes,
		DependencyIndexes: file_production_record_proto_depIdxs,
		EnumInfos:         file_production_record_proto_enumTypes,
		MessageInfos:      file_production_record_proto_msgTypes,
	}.Build()
	File_production_record_proto = out.File
	file_production_record_proto_rawDesc = nil
	file_production_record_proto_goTypes = nil
	file_production_record_proto_depIdxs = nil
}
