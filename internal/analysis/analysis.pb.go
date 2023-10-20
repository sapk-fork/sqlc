// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: analysis/analysis.proto

package analysis

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

type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalog string `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	Schema  string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_analysis_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *Identifier) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *Identifier) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Identifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OriginalName string      `protobuf:"bytes,2,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	DataType     string      `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	NotNull      bool        `protobuf:"varint,4,opt,name=not_null,json=notNull,proto3" json:"not_null,omitempty"`
	Unsigned     bool        `protobuf:"varint,5,opt,name=unsigned,proto3" json:"unsigned,omitempty"`
	IsArray      bool        `protobuf:"varint,6,opt,name=is_array,json=isArray,proto3" json:"is_array,omitempty"`
	ArrayDims    int32       `protobuf:"varint,7,opt,name=array_dims,json=arrayDims,proto3" json:"array_dims,omitempty"`
	Comment      string      `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Length       int32       `protobuf:"varint,9,opt,name=length,proto3" json:"length,omitempty"` // *int
	IsNamedParam bool        `protobuf:"varint,10,opt,name=is_named_param,json=isNamedParam,proto3" json:"is_named_param,omitempty"`
	IsFuncCall   bool        `protobuf:"varint,11,opt,name=is_func_call,json=isFuncCall,proto3" json:"is_func_call,omitempty"`
	Scope        string      `protobuf:"bytes,12,opt,name=scope,proto3" json:"scope,omitempty"`
	Table        *Identifier `protobuf:"bytes,13,opt,name=table,proto3" json:"table,omitempty"`
	TableAlias   string      `protobuf:"bytes,14,opt,name=table_alias,json=tableAlias,proto3" json:"table_alias,omitempty"`
	Type         *Identifier `protobuf:"bytes,15,opt,name=type,proto3" json:"type,omitempty"`
	EmbedTable   *Identifier `protobuf:"bytes,16,opt,name=embed_table,json=embedTable,proto3" json:"embed_table,omitempty"`
	IsSqlcSlice  bool        `protobuf:"varint,17,opt,name=is_sqlc_slice,json=isSqlcSlice,proto3" json:"is_sqlc_slice,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_analysis_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *Column) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Column) GetNotNull() bool {
	if x != nil {
		return x.NotNull
	}
	return false
}

func (x *Column) GetUnsigned() bool {
	if x != nil {
		return x.Unsigned
	}
	return false
}

func (x *Column) GetIsArray() bool {
	if x != nil {
		return x.IsArray
	}
	return false
}

func (x *Column) GetArrayDims() int32 {
	if x != nil {
		return x.ArrayDims
	}
	return 0
}

func (x *Column) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Column) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Column) GetIsNamedParam() bool {
	if x != nil {
		return x.IsNamedParam
	}
	return false
}

func (x *Column) GetIsFuncCall() bool {
	if x != nil {
		return x.IsFuncCall
	}
	return false
}

func (x *Column) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Column) GetTable() *Identifier {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *Column) GetTableAlias() string {
	if x != nil {
		return x.TableAlias
	}
	return ""
}

func (x *Column) GetType() *Identifier {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Column) GetEmbedTable() *Identifier {
	if x != nil {
		return x.EmbedTable
	}
	return nil
}

func (x *Column) GetIsSqlcSlice() bool {
	if x != nil {
		return x.IsSqlcSlice
	}
	return false
}

type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Column *Column `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_analysis_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *Parameter) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Parameter) GetColumn() *Column {
	if x != nil {
		return x.Column
	}
	return nil
}

type Analysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []*Column    `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Params  []*Parameter `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *Analysis) Reset() {
	*x = Analysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_analysis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Analysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Analysis) ProtoMessage() {}

func (x *Analysis) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_analysis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Analysis.ProtoReflect.Descriptor instead.
func (*Analysis) Descriptor() ([]byte, []int) {
	return file_analysis_analysis_proto_rawDescGZIP(), []int{3}
}

func (x *Analysis) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Analysis) GetParams() []*Parameter {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_analysis_analysis_proto protoreflect.FileDescriptor

var file_analysis_analysis_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x22, 0x52, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb1, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x5f,
	0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x4e,
	0x75, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x44, 0x69, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x75, 0x6e, 0x63, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x35, 0x0a, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x65, 0x6d, 0x62,
	0x65, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x73, 0x71,
	0x6c, 0x63, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x53, 0x71, 0x6c, 0x63, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x09, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x63, 0x0a, 0x08, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0x89, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x42, 0x0d, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x71,
	0x6c, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x71, 0x6c, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0xa2, 0x02, 0x03,
	0x41, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0xca, 0x02,
	0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0xe2, 0x02, 0x14, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_analysis_analysis_proto_rawDescOnce sync.Once
	file_analysis_analysis_proto_rawDescData = file_analysis_analysis_proto_rawDesc
)

func file_analysis_analysis_proto_rawDescGZIP() []byte {
	file_analysis_analysis_proto_rawDescOnce.Do(func() {
		file_analysis_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_analysis_analysis_proto_rawDescData)
	})
	return file_analysis_analysis_proto_rawDescData
}

var file_analysis_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_analysis_analysis_proto_goTypes = []interface{}{
	(*Identifier)(nil), // 0: analysis.Identifier
	(*Column)(nil),     // 1: analysis.Column
	(*Parameter)(nil),  // 2: analysis.Parameter
	(*Analysis)(nil),   // 3: analysis.Analysis
}
var file_analysis_analysis_proto_depIdxs = []int32{
	0, // 0: analysis.Column.table:type_name -> analysis.Identifier
	0, // 1: analysis.Column.type:type_name -> analysis.Identifier
	0, // 2: analysis.Column.embed_table:type_name -> analysis.Identifier
	1, // 3: analysis.Parameter.column:type_name -> analysis.Column
	1, // 4: analysis.Analysis.columns:type_name -> analysis.Column
	2, // 5: analysis.Analysis.params:type_name -> analysis.Parameter
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_analysis_analysis_proto_init() }
func file_analysis_analysis_proto_init() {
	if File_analysis_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analysis_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
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
		file_analysis_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
		file_analysis_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
		file_analysis_analysis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Analysis); i {
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
			RawDescriptor: file_analysis_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_analysis_analysis_proto_goTypes,
		DependencyIndexes: file_analysis_analysis_proto_depIdxs,
		MessageInfos:      file_analysis_analysis_proto_msgTypes,
	}.Build()
	File_analysis_analysis_proto = out.File
	file_analysis_analysis_proto_rawDesc = nil
	file_analysis_analysis_proto_goTypes = nil
	file_analysis_analysis_proto_depIdxs = nil
}
