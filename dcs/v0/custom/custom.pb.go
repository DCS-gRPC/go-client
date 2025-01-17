// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: dcs/custom/v0/custom.proto

package custom

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

type RequestMissionAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitName    string `protobuf:"bytes,1,opt,name=unit_name,json=unitName,proto3" json:"unit_name,omitempty"`
	MissionType string `protobuf:"bytes,2,opt,name=mission_type,json=missionType,proto3" json:"mission_type,omitempty"`
}

func (x *RequestMissionAssignmentRequest) Reset() {
	*x = RequestMissionAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMissionAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMissionAssignmentRequest) ProtoMessage() {}

func (x *RequestMissionAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMissionAssignmentRequest.ProtoReflect.Descriptor instead.
func (*RequestMissionAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMissionAssignmentRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *RequestMissionAssignmentRequest) GetMissionType() string {
	if x != nil {
		return x.MissionType
	}
	return ""
}

type RequestMissionAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestMissionAssignmentResponse) Reset() {
	*x = RequestMissionAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMissionAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMissionAssignmentResponse) ProtoMessage() {}

func (x *RequestMissionAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMissionAssignmentResponse.ProtoReflect.Descriptor instead.
func (*RequestMissionAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{1}
}

type JoinMissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitName    string `protobuf:"bytes,1,opt,name=unit_name,json=unitName,proto3" json:"unit_name,omitempty"`
	MissionCode int32  `protobuf:"varint,2,opt,name=mission_code,json=missionCode,proto3" json:"mission_code,omitempty"`
}

func (x *JoinMissionRequest) Reset() {
	*x = JoinMissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinMissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMissionRequest) ProtoMessage() {}

func (x *JoinMissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMissionRequest.ProtoReflect.Descriptor instead.
func (*JoinMissionRequest) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{2}
}

func (x *JoinMissionRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *JoinMissionRequest) GetMissionCode() int32 {
	if x != nil {
		return x.MissionCode
	}
	return 0
}

type JoinMissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinMissionResponse) Reset() {
	*x = JoinMissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinMissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMissionResponse) ProtoMessage() {}

func (x *JoinMissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMissionResponse.ProtoReflect.Descriptor instead.
func (*JoinMissionResponse) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{3}
}

type AbortMissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitName string `protobuf:"bytes,1,opt,name=unit_name,json=unitName,proto3" json:"unit_name,omitempty"`
}

func (x *AbortMissionRequest) Reset() {
	*x = AbortMissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortMissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortMissionRequest) ProtoMessage() {}

func (x *AbortMissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortMissionRequest.ProtoReflect.Descriptor instead.
func (*AbortMissionRequest) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{4}
}

func (x *AbortMissionRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

type AbortMissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AbortMissionResponse) Reset() {
	*x = AbortMissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortMissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortMissionResponse) ProtoMessage() {}

func (x *AbortMissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortMissionResponse.ProtoReflect.Descriptor instead.
func (*AbortMissionResponse) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{5}
}

type GetMissionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitName string `protobuf:"bytes,1,opt,name=unit_name,json=unitName,proto3" json:"unit_name,omitempty"`
}

func (x *GetMissionStatusRequest) Reset() {
	*x = GetMissionStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMissionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMissionStatusRequest) ProtoMessage() {}

func (x *GetMissionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMissionStatusRequest.ProtoReflect.Descriptor instead.
func (*GetMissionStatusRequest) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{6}
}

func (x *GetMissionStatusRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

type GetMissionStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMissionStatusResponse) Reset() {
	*x = GetMissionStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMissionStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMissionStatusResponse) ProtoMessage() {}

func (x *GetMissionStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMissionStatusResponse.ProtoReflect.Descriptor instead.
func (*GetMissionStatusResponse) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{7}
}

type EvalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lua string `protobuf:"bytes,1,opt,name=lua,proto3" json:"lua,omitempty"`
}

func (x *EvalRequest) Reset() {
	*x = EvalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalRequest) ProtoMessage() {}

func (x *EvalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalRequest.ProtoReflect.Descriptor instead.
func (*EvalRequest) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{8}
}

func (x *EvalRequest) GetLua() string {
	if x != nil {
		return x.Lua
	}
	return ""
}

type EvalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *EvalResponse) Reset() {
	*x = EvalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalResponse) ProtoMessage() {}

func (x *EvalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalResponse.ProtoReflect.Descriptor instead.
func (*EvalResponse) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{9}
}

func (x *EvalResponse) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type GetMagneticDeclinationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / Latitude in Decimal Degrees format
	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	// / Longitude in Decimal Degrees format
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	// / Altitude in Meters above Mean Sea Level (MSL)
	Alt float64 `protobuf:"fixed64,3,opt,name=alt,proto3" json:"alt,omitempty"`
}

func (x *GetMagneticDeclinationRequest) Reset() {
	*x = GetMagneticDeclinationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMagneticDeclinationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMagneticDeclinationRequest) ProtoMessage() {}

func (x *GetMagneticDeclinationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMagneticDeclinationRequest.ProtoReflect.Descriptor instead.
func (*GetMagneticDeclinationRequest) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{10}
}

func (x *GetMagneticDeclinationRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GetMagneticDeclinationRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *GetMagneticDeclinationRequest) GetAlt() float64 {
	if x != nil {
		return x.Alt
	}
	return 0
}

type GetMagneticDeclinationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / Magnetic declination in degrees. A negative value is an westerly
	// / declination, while a positive value is a easterly declination. `True
	// / North` + `declination` = `Magnetic North`
	Declination float64 `protobuf:"fixed64,1,opt,name=declination,proto3" json:"declination,omitempty"`
}

func (x *GetMagneticDeclinationResponse) Reset() {
	*x = GetMagneticDeclinationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_custom_v0_custom_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMagneticDeclinationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMagneticDeclinationResponse) ProtoMessage() {}

func (x *GetMagneticDeclinationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_custom_v0_custom_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMagneticDeclinationResponse.ProtoReflect.Descriptor instead.
func (*GetMagneticDeclinationResponse) Descriptor() ([]byte, []int) {
	return file_dcs_custom_v0_custom_proto_rawDescGZIP(), []int{11}
}

func (x *GetMagneticDeclinationResponse) GetDeclination() float64 {
	if x != nil {
		return x.Declination
	}
	return 0
}

var File_dcs_custom_v0_custom_proto protoreflect.FileDescriptor

var file_dcs_custom_v0_custom_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x63, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2f, 0x76, 0x30, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x63,
	0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x22, 0x61, 0x0a, 0x1f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x22,
	0x0a, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x54, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x32, 0x0a, 0x13, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x0a, 0x0b, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x75, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x75, 0x61,
	0x22, 0x22, 0x0a, 0x0c, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x67, 0x6e, 0x65,
	0x74, 0x69, 0x63, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x22, 0x42, 0x0a, 0x1e, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xe4, 0x04, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7d, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e,
	0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e,
	0x76, 0x30, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x41, 0x62, 0x6f, 0x72,
	0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64,
	0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x41, 0x62, 0x6f,
	0x72, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x45, 0x76,
	0x61, 0x6c, 0x12, 0x1a, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e,
	0x76, 0x30, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x45,
	0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x44, 0x65, 0x63, 0x6c,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x67, 0x6e, 0x65,
	0x74, 0x69, 0x63, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x69,
	0x63, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x43, 0x53, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x67, 0x6f,
	0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x64, 0x63, 0x73, 0x2f, 0x76, 0x30,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0xaa, 0x02, 0x1f, 0x52, 0x75, 0x72, 0x6f, 0x75, 0x6e,
	0x69, 0x4a, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x44, 0x63, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x30, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dcs_custom_v0_custom_proto_rawDescOnce sync.Once
	file_dcs_custom_v0_custom_proto_rawDescData = file_dcs_custom_v0_custom_proto_rawDesc
)

func file_dcs_custom_v0_custom_proto_rawDescGZIP() []byte {
	file_dcs_custom_v0_custom_proto_rawDescOnce.Do(func() {
		file_dcs_custom_v0_custom_proto_rawDescData = protoimpl.X.CompressGZIP(file_dcs_custom_v0_custom_proto_rawDescData)
	})
	return file_dcs_custom_v0_custom_proto_rawDescData
}

var file_dcs_custom_v0_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_dcs_custom_v0_custom_proto_goTypes = []interface{}{
	(*RequestMissionAssignmentRequest)(nil),  // 0: dcs.custom.v0.RequestMissionAssignmentRequest
	(*RequestMissionAssignmentResponse)(nil), // 1: dcs.custom.v0.RequestMissionAssignmentResponse
	(*JoinMissionRequest)(nil),               // 2: dcs.custom.v0.JoinMissionRequest
	(*JoinMissionResponse)(nil),              // 3: dcs.custom.v0.JoinMissionResponse
	(*AbortMissionRequest)(nil),              // 4: dcs.custom.v0.AbortMissionRequest
	(*AbortMissionResponse)(nil),             // 5: dcs.custom.v0.AbortMissionResponse
	(*GetMissionStatusRequest)(nil),          // 6: dcs.custom.v0.GetMissionStatusRequest
	(*GetMissionStatusResponse)(nil),         // 7: dcs.custom.v0.GetMissionStatusResponse
	(*EvalRequest)(nil),                      // 8: dcs.custom.v0.EvalRequest
	(*EvalResponse)(nil),                     // 9: dcs.custom.v0.EvalResponse
	(*GetMagneticDeclinationRequest)(nil),    // 10: dcs.custom.v0.GetMagneticDeclinationRequest
	(*GetMagneticDeclinationResponse)(nil),   // 11: dcs.custom.v0.GetMagneticDeclinationResponse
}
var file_dcs_custom_v0_custom_proto_depIdxs = []int32{
	0,  // 0: dcs.custom.v0.CustomService.RequestMissionAssignment:input_type -> dcs.custom.v0.RequestMissionAssignmentRequest
	2,  // 1: dcs.custom.v0.CustomService.JoinMission:input_type -> dcs.custom.v0.JoinMissionRequest
	4,  // 2: dcs.custom.v0.CustomService.AbortMission:input_type -> dcs.custom.v0.AbortMissionRequest
	6,  // 3: dcs.custom.v0.CustomService.GetMissionStatus:input_type -> dcs.custom.v0.GetMissionStatusRequest
	8,  // 4: dcs.custom.v0.CustomService.Eval:input_type -> dcs.custom.v0.EvalRequest
	10, // 5: dcs.custom.v0.CustomService.GetMagneticDeclination:input_type -> dcs.custom.v0.GetMagneticDeclinationRequest
	1,  // 6: dcs.custom.v0.CustomService.RequestMissionAssignment:output_type -> dcs.custom.v0.RequestMissionAssignmentResponse
	3,  // 7: dcs.custom.v0.CustomService.JoinMission:output_type -> dcs.custom.v0.JoinMissionResponse
	5,  // 8: dcs.custom.v0.CustomService.AbortMission:output_type -> dcs.custom.v0.AbortMissionResponse
	7,  // 9: dcs.custom.v0.CustomService.GetMissionStatus:output_type -> dcs.custom.v0.GetMissionStatusResponse
	9,  // 10: dcs.custom.v0.CustomService.Eval:output_type -> dcs.custom.v0.EvalResponse
	11, // 11: dcs.custom.v0.CustomService.GetMagneticDeclination:output_type -> dcs.custom.v0.GetMagneticDeclinationResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_dcs_custom_v0_custom_proto_init() }
func file_dcs_custom_v0_custom_proto_init() {
	if File_dcs_custom_v0_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dcs_custom_v0_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMissionAssignmentRequest); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMissionAssignmentResponse); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinMissionRequest); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinMissionResponse); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbortMissionRequest); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbortMissionResponse); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMissionStatusRequest); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMissionStatusResponse); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvalRequest); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvalResponse); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMagneticDeclinationRequest); i {
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
		file_dcs_custom_v0_custom_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMagneticDeclinationResponse); i {
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
			RawDescriptor: file_dcs_custom_v0_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dcs_custom_v0_custom_proto_goTypes,
		DependencyIndexes: file_dcs_custom_v0_custom_proto_depIdxs,
		MessageInfos:      file_dcs_custom_v0_custom_proto_msgTypes,
	}.Build()
	File_dcs_custom_v0_custom_proto = out.File
	file_dcs_custom_v0_custom_proto_rawDesc = nil
	file_dcs_custom_v0_custom_proto_goTypes = nil
	file_dcs_custom_v0_custom_proto_depIdxs = nil
}
