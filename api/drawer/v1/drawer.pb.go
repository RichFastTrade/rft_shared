// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: drawer/v1/drawer.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KlinesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbol        string                 `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Interval      string                 `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Klines        []*KlineData           `protobuf:"bytes,3,rep,name=klines,proto3" json:"klines,omitempty"`
	Fractals      []*FractalData         `protobuf:"bytes,4,rep,name=fractals,proto3" json:"fractals,omitempty"`
	UpCenters     []*CenterData          `protobuf:"bytes,5,rep,name=up_centers,json=upCenters,proto3" json:"up_centers,omitempty"`
	DownCenters   []*CenterData          `protobuf:"bytes,6,rep,name=down_centers,json=downCenters,proto3" json:"down_centers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KlinesRequest) Reset() {
	*x = KlinesRequest{}
	mi := &file_drawer_v1_drawer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KlinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KlinesRequest) ProtoMessage() {}

func (x *KlinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drawer_v1_drawer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KlinesRequest.ProtoReflect.Descriptor instead.
func (*KlinesRequest) Descriptor() ([]byte, []int) {
	return file_drawer_v1_drawer_proto_rawDescGZIP(), []int{0}
}

func (x *KlinesRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *KlinesRequest) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *KlinesRequest) GetKlines() []*KlineData {
	if x != nil {
		return x.Klines
	}
	return nil
}

func (x *KlinesRequest) GetFractals() []*FractalData {
	if x != nil {
		return x.Fractals
	}
	return nil
}

func (x *KlinesRequest) GetUpCenters() []*CenterData {
	if x != nil {
		return x.UpCenters
	}
	return nil
}

func (x *KlinesRequest) GetDownCenters() []*CenterData {
	if x != nil {
		return x.DownCenters
	}
	return nil
}

type KlineData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Time          int64                  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Open          string                 `protobuf:"bytes,2,opt,name=open,proto3" json:"open,omitempty"`
	High          string                 `protobuf:"bytes,3,opt,name=high,proto3" json:"high,omitempty"`
	Low           string                 `protobuf:"bytes,4,opt,name=low,proto3" json:"low,omitempty"`
	Close         string                 `protobuf:"bytes,5,opt,name=close,proto3" json:"close,omitempty"`
	Volume        string                 `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume,omitempty"`
	BollMid       *string                `protobuf:"bytes,7,opt,name=boll_mid,json=bollMid,proto3,oneof" json:"boll_mid,omitempty"`
	BollUpper     *string                `protobuf:"bytes,8,opt,name=boll_upper,json=bollUpper,proto3,oneof" json:"boll_upper,omitempty"`
	BollLower     *string                `protobuf:"bytes,9,opt,name=boll_lower,json=bollLower,proto3,oneof" json:"boll_lower,omitempty"`
	MacdDif       *string                `protobuf:"bytes,10,opt,name=macd_dif,json=macdDif,proto3,oneof" json:"macd_dif,omitempty"`
	MacdDea       *string                `protobuf:"bytes,11,opt,name=macd_dea,json=macdDea,proto3,oneof" json:"macd_dea,omitempty"`
	Macd          *string                `protobuf:"bytes,12,opt,name=macd,proto3,oneof" json:"macd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KlineData) Reset() {
	*x = KlineData{}
	mi := &file_drawer_v1_drawer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KlineData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KlineData) ProtoMessage() {}

func (x *KlineData) ProtoReflect() protoreflect.Message {
	mi := &file_drawer_v1_drawer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KlineData.ProtoReflect.Descriptor instead.
func (*KlineData) Descriptor() ([]byte, []int) {
	return file_drawer_v1_drawer_proto_rawDescGZIP(), []int{1}
}

func (x *KlineData) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *KlineData) GetOpen() string {
	if x != nil {
		return x.Open
	}
	return ""
}

func (x *KlineData) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *KlineData) GetLow() string {
	if x != nil {
		return x.Low
	}
	return ""
}

func (x *KlineData) GetClose() string {
	if x != nil {
		return x.Close
	}
	return ""
}

func (x *KlineData) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *KlineData) GetBollMid() string {
	if x != nil && x.BollMid != nil {
		return *x.BollMid
	}
	return ""
}

func (x *KlineData) GetBollUpper() string {
	if x != nil && x.BollUpper != nil {
		return *x.BollUpper
	}
	return ""
}

func (x *KlineData) GetBollLower() string {
	if x != nil && x.BollLower != nil {
		return *x.BollLower
	}
	return ""
}

func (x *KlineData) GetMacdDif() string {
	if x != nil && x.MacdDif != nil {
		return *x.MacdDif
	}
	return ""
}

func (x *KlineData) GetMacdDea() string {
	if x != nil && x.MacdDea != nil {
		return *x.MacdDea
	}
	return ""
}

func (x *KlineData) GetMacd() string {
	if x != nil && x.Macd != nil {
		return *x.Macd
	}
	return ""
}

type FractalData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Current       *KlineData             `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FractalData) Reset() {
	*x = FractalData{}
	mi := &file_drawer_v1_drawer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FractalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FractalData) ProtoMessage() {}

func (x *FractalData) ProtoReflect() protoreflect.Message {
	mi := &file_drawer_v1_drawer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FractalData.ProtoReflect.Descriptor instead.
func (*FractalData) Descriptor() ([]byte, []int) {
	return file_drawer_v1_drawer_proto_rawDescGZIP(), []int{2}
}

func (x *FractalData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FractalData) GetCurrent() *KlineData {
	if x != nil {
		return x.Current
	}
	return nil
}

type CenterData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         *FractalData           `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End           *FractalData           `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Highest       string                 `protobuf:"bytes,3,opt,name=highest,proto3" json:"highest,omitempty"`
	Lowest        string                 `protobuf:"bytes,4,opt,name=lowest,proto3" json:"lowest,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CenterData) Reset() {
	*x = CenterData{}
	mi := &file_drawer_v1_drawer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CenterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CenterData) ProtoMessage() {}

func (x *CenterData) ProtoReflect() protoreflect.Message {
	mi := &file_drawer_v1_drawer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CenterData.ProtoReflect.Descriptor instead.
func (*CenterData) Descriptor() ([]byte, []int) {
	return file_drawer_v1_drawer_proto_rawDescGZIP(), []int{3}
}

func (x *CenterData) GetStart() *FractalData {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *CenterData) GetEnd() *FractalData {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *CenterData) GetHighest() string {
	if x != nil {
		return x.Highest
	}
	return ""
}

func (x *CenterData) GetLowest() string {
	if x != nil {
		return x.Lowest
	}
	return ""
}

type PictureReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PngPicture    []byte                 `protobuf:"bytes,1,opt,name=png_picture,json=pngPicture,proto3" json:"png_picture,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PictureReply) Reset() {
	*x = PictureReply{}
	mi := &file_drawer_v1_drawer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PictureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PictureReply) ProtoMessage() {}

func (x *PictureReply) ProtoReflect() protoreflect.Message {
	mi := &file_drawer_v1_drawer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PictureReply.ProtoReflect.Descriptor instead.
func (*PictureReply) Descriptor() ([]byte, []int) {
	return file_drawer_v1_drawer_proto_rawDescGZIP(), []int{4}
}

func (x *PictureReply) GetPngPicture() []byte {
	if x != nil {
		return x.PngPicture
	}
	return nil
}

var File_drawer_v1_drawer_proto protoreflect.FileDescriptor

var file_drawer_v1_drawer_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x72, 0x61, 0x77,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x95, 0x02, 0x0a, 0x0d, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x06, 0x6b, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x72, 0x61, 0x77,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x06, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74,
	0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x77,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x75,
	0x70, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x75, 0x70, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x38, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b,
	0x64, 0x6f, 0x77, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x96, 0x03, 0x0a, 0x09,
	0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70, 0x65,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x62, 0x6f, 0x6c, 0x6c, 0x5f, 0x6d, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6c, 0x6c, 0x4d,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x6f, 0x6c, 0x6c, 0x5f, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x62, 0x6f, 0x6c,
	0x6c, 0x55, 0x70, 0x70, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x6f, 0x6c,
	0x6c, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x09, 0x62, 0x6f, 0x6c, 0x6c, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x6d, 0x61, 0x63, 0x64, 0x5f, 0x64, 0x69, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x03, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x64, 0x44, 0x69, 0x66, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x6d, 0x61, 0x63, 0x64, 0x5f, 0x64, 0x65, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x64, 0x44, 0x65, 0x61, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x6d, 0x61, 0x63, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x04, 0x6d,
	0x61, 0x63, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x6f, 0x6c, 0x6c, 0x5f,
	0x6d, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x6f, 0x6c, 0x6c, 0x5f, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x6f, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x61, 0x63, 0x64, 0x5f, 0x64, 0x69, 0x66, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6d, 0x61, 0x63, 0x64, 0x5f, 0x64, 0x65, 0x61, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6d, 0x61, 0x63, 0x64, 0x22, 0x51, 0x0a, 0x0b, 0x46, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72,
	0x61, 0x63, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x77, 0x65,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74,
	0x22, 0x2f, 0x0a, 0x0c, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6e, 0x67, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x6e, 0x67, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x32, 0x48, 0x0a, 0x06, 0x44, 0x72, 0x61, 0x77, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x5a, 0x65, 0x6e, 0x50, 0x69, 0x63, 0x12, 0x18, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x69, 0x63, 0x68, 0x46, 0x61,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x72, 0x66, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_drawer_v1_drawer_proto_rawDescOnce sync.Once
	file_drawer_v1_drawer_proto_rawDescData []byte
)

func file_drawer_v1_drawer_proto_rawDescGZIP() []byte {
	file_drawer_v1_drawer_proto_rawDescOnce.Do(func() {
		file_drawer_v1_drawer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_drawer_v1_drawer_proto_rawDesc), len(file_drawer_v1_drawer_proto_rawDesc)))
	})
	return file_drawer_v1_drawer_proto_rawDescData
}

var file_drawer_v1_drawer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_drawer_v1_drawer_proto_goTypes = []any{
	(*KlinesRequest)(nil), // 0: drawer.v1.KlinesRequest
	(*KlineData)(nil),     // 1: drawer.v1.KlineData
	(*FractalData)(nil),   // 2: drawer.v1.FractalData
	(*CenterData)(nil),    // 3: drawer.v1.CenterData
	(*PictureReply)(nil),  // 4: drawer.v1.PictureReply
}
var file_drawer_v1_drawer_proto_depIdxs = []int32{
	1, // 0: drawer.v1.KlinesRequest.klines:type_name -> drawer.v1.KlineData
	2, // 1: drawer.v1.KlinesRequest.fractals:type_name -> drawer.v1.FractalData
	3, // 2: drawer.v1.KlinesRequest.up_centers:type_name -> drawer.v1.CenterData
	3, // 3: drawer.v1.KlinesRequest.down_centers:type_name -> drawer.v1.CenterData
	1, // 4: drawer.v1.FractalData.current:type_name -> drawer.v1.KlineData
	2, // 5: drawer.v1.CenterData.start:type_name -> drawer.v1.FractalData
	2, // 6: drawer.v1.CenterData.end:type_name -> drawer.v1.FractalData
	0, // 7: drawer.v1.Drawer.GetZenPic:input_type -> drawer.v1.KlinesRequest
	4, // 8: drawer.v1.Drawer.GetZenPic:output_type -> drawer.v1.PictureReply
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_drawer_v1_drawer_proto_init() }
func file_drawer_v1_drawer_proto_init() {
	if File_drawer_v1_drawer_proto != nil {
		return
	}
	file_drawer_v1_drawer_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_drawer_v1_drawer_proto_rawDesc), len(file_drawer_v1_drawer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drawer_v1_drawer_proto_goTypes,
		DependencyIndexes: file_drawer_v1_drawer_proto_depIdxs,
		MessageInfos:      file_drawer_v1_drawer_proto_msgTypes,
	}.Build()
	File_drawer_v1_drawer_proto = out.File
	file_drawer_v1_drawer_proto_goTypes = nil
	file_drawer_v1_drawer_proto_depIdxs = nil
}
