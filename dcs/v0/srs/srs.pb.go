// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: dcs/srs/v0/srs.proto

package srs

import (
	common "github.com/DCS-gRPC/go-bindings/dcs/v0/common"
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

type TransmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The text that is synthesized to speech and transmitted to SRS. Supports SSML tags (you should
	// not wrap the text in the root `<speak>` tag though).
	Ssml string `protobuf:"bytes,1,opt,name=ssml,proto3" json:"ssml,omitempty"`
	// The plain text without any transformations made to it for the purpose of getting it spoken out
	// as desired (no SSML tags, no FOUR NINER instead of 49, ...). Even though this field is
	// optional, please consider providing it as it can be used to display the spoken text to players
	// with hearing impairments.
	Plaintext *string `protobuf:"bytes,2,opt,name=plaintext,proto3,oneof" json:"plaintext,omitempty"`
	// The radio frequency in Hz the transmission is send to. Example: 251000000
	// for 251.00MHz.
	Frequency uint64 `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Name of the SRS client. Defaults to "DCS-gRPC".
	SrsClientName *string `protobuf:"bytes,4,opt,name=srs_client_name,json=srsClientName,proto3,oneof" json:"srs_client_name,omitempty"`
	// The origin of the transmission. Relevant if the SRS server has "Line of
	// Sight" and/or "Distance Limit" enabled.
	Position *common.InputPosition `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`
	// The coalition of the transmission. Relevant if the SRS server has "Secure
	// Coalition Radios" enabled. Only Blue and Red are supported, all other
	// values will fallback to Spectator.
	Coalition common.Coalition `protobuf:"varint,6,opt,name=coalition,proto3,enum=dcs.common.v0.Coalition" json:"coalition,omitempty"`
	// Whether to keep the request open until the whole transmission was sent. If
	// enabled, you can send the next transmission after you've received the
	// response for the previous one and be sure that they don't overlap (talk
	// over each other). If disabled, you'll receive a response right away (kind
	// of fire and forget). You can use the returned duration as a spacing between
	// TTS requests to prevent the overlap of multiple playbacks yourself.
	Async bool `protobuf:"varint,7,opt,name=async,proto3" json:"async,omitempty"`
	// Optional TTS provider to be use. Defaults to the one configured in your
	// config or to Windows' built-in TTS.
	//
	// Types that are assignable to Provider:
	//
	//	*TransmitRequest_Aws_
	//	*TransmitRequest_Azure_
	//	*TransmitRequest_Gcloud
	//	*TransmitRequest_Win
	Provider isTransmitRequest_Provider `protobuf_oneof:"provider"`
}

func (x *TransmitRequest) Reset() {
	*x = TransmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitRequest) ProtoMessage() {}

func (x *TransmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitRequest.ProtoReflect.Descriptor instead.
func (*TransmitRequest) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{0}
}

func (x *TransmitRequest) GetSsml() string {
	if x != nil {
		return x.Ssml
	}
	return ""
}

func (x *TransmitRequest) GetPlaintext() string {
	if x != nil && x.Plaintext != nil {
		return *x.Plaintext
	}
	return ""
}

func (x *TransmitRequest) GetFrequency() uint64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *TransmitRequest) GetSrsClientName() string {
	if x != nil && x.SrsClientName != nil {
		return *x.SrsClientName
	}
	return ""
}

func (x *TransmitRequest) GetPosition() *common.InputPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *TransmitRequest) GetCoalition() common.Coalition {
	if x != nil {
		return x.Coalition
	}
	return common.Coalition(0)
}

func (x *TransmitRequest) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (m *TransmitRequest) GetProvider() isTransmitRequest_Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (x *TransmitRequest) GetAws() *TransmitRequest_Aws {
	if x, ok := x.GetProvider().(*TransmitRequest_Aws_); ok {
		return x.Aws
	}
	return nil
}

func (x *TransmitRequest) GetAzure() *TransmitRequest_Azure {
	if x, ok := x.GetProvider().(*TransmitRequest_Azure_); ok {
		return x.Azure
	}
	return nil
}

func (x *TransmitRequest) GetGcloud() *TransmitRequest_GCloud {
	if x, ok := x.GetProvider().(*TransmitRequest_Gcloud); ok {
		return x.Gcloud
	}
	return nil
}

func (x *TransmitRequest) GetWin() *TransmitRequest_Windows {
	if x, ok := x.GetProvider().(*TransmitRequest_Win); ok {
		return x.Win
	}
	return nil
}

type isTransmitRequest_Provider interface {
	isTransmitRequest_Provider()
}

type TransmitRequest_Aws_ struct {
	Aws *TransmitRequest_Aws `protobuf:"bytes,8,opt,name=aws,proto3,oneof"`
}

type TransmitRequest_Azure_ struct {
	Azure *TransmitRequest_Azure `protobuf:"bytes,9,opt,name=azure,proto3,oneof"`
}

type TransmitRequest_Gcloud struct {
	Gcloud *TransmitRequest_GCloud `protobuf:"bytes,10,opt,name=gcloud,proto3,oneof"`
}

type TransmitRequest_Win struct {
	Win *TransmitRequest_Windows `protobuf:"bytes,11,opt,name=win,proto3,oneof"`
}

func (*TransmitRequest_Aws_) isTransmitRequest_Provider() {}

func (*TransmitRequest_Azure_) isTransmitRequest_Provider() {}

func (*TransmitRequest_Gcloud) isTransmitRequest_Provider() {}

func (*TransmitRequest_Win) isTransmitRequest_Provider() {}

type TransmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The duration in milliseconds it roughly takes to speak the transmission.
	DurationMs uint32 `protobuf:"varint,1,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
}

func (x *TransmitResponse) Reset() {
	*x = TransmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitResponse) ProtoMessage() {}

func (x *TransmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitResponse.ProtoReflect.Descriptor instead.
func (*TransmitResponse) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{1}
}

func (x *TransmitResponse) GetDurationMs() uint32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

type GetClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClientsRequest) Reset() {
	*x = GetClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsRequest) ProtoMessage() {}

func (x *GetClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsRequest.ProtoReflect.Descriptor instead.
func (*GetClientsRequest) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{2}
}

type GetClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*GetClientsResponse_Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetClientsResponse) Reset() {
	*x = GetClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsResponse) ProtoMessage() {}

func (x *GetClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsResponse.ProtoReflect.Descriptor instead.
func (*GetClientsResponse) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{3}
}

func (x *GetClientsResponse) GetClients() []*GetClientsResponse_Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type TransmitRequest_Aws struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The voice the text is synthesized in, see:
	// https://docs.aws.amazon.com/polly/latest/dg/voicelist.html
	Voice *string `protobuf:"bytes,1,opt,name=voice,proto3,oneof" json:"voice,omitempty"`
}

func (x *TransmitRequest_Aws) Reset() {
	*x = TransmitRequest_Aws{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitRequest_Aws) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitRequest_Aws) ProtoMessage() {}

func (x *TransmitRequest_Aws) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitRequest_Aws.ProtoReflect.Descriptor instead.
func (*TransmitRequest_Aws) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TransmitRequest_Aws) GetVoice() string {
	if x != nil && x.Voice != nil {
		return *x.Voice
	}
	return ""
}

type TransmitRequest_Azure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The voice the text is synthesized in, see:
	// https://learn.microsoft.com/azure/cognitive-services/speech-service/language-support
	Voice *string `protobuf:"bytes,1,opt,name=voice,proto3,oneof" json:"voice,omitempty"`
}

func (x *TransmitRequest_Azure) Reset() {
	*x = TransmitRequest_Azure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitRequest_Azure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitRequest_Azure) ProtoMessage() {}

func (x *TransmitRequest_Azure) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitRequest_Azure.ProtoReflect.Descriptor instead.
func (*TransmitRequest_Azure) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TransmitRequest_Azure) GetVoice() string {
	if x != nil && x.Voice != nil {
		return *x.Voice
	}
	return ""
}

type TransmitRequest_GCloud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The voice the text is synthesized in, see:
	// https://cloud.google.com/text-to-speech/docs/voices
	Voice *string `protobuf:"bytes,1,opt,name=voice,proto3,oneof" json:"voice,omitempty"`
}

func (x *TransmitRequest_GCloud) Reset() {
	*x = TransmitRequest_GCloud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitRequest_GCloud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitRequest_GCloud) ProtoMessage() {}

func (x *TransmitRequest_GCloud) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitRequest_GCloud.ProtoReflect.Descriptor instead.
func (*TransmitRequest_GCloud) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{0, 2}
}

func (x *TransmitRequest_GCloud) GetVoice() string {
	if x != nil && x.Voice != nil {
		return *x.Voice
	}
	return ""
}

type TransmitRequest_Windows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The voice the text is synthesized in.
	Voice *string `protobuf:"bytes,1,opt,name=voice,proto3,oneof" json:"voice,omitempty"`
}

func (x *TransmitRequest_Windows) Reset() {
	*x = TransmitRequest_Windows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitRequest_Windows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitRequest_Windows) ProtoMessage() {}

func (x *TransmitRequest_Windows) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitRequest_Windows.ProtoReflect.Descriptor instead.
func (*TransmitRequest_Windows) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{0, 3}
}

func (x *TransmitRequest_Windows) GetVoice() string {
	if x != nil && x.Voice != nil {
		return *x.Voice
	}
	return ""
}

type GetClientsResponse_Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unit that is connected to SRS.
	Unit *common.Unit `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	// The radio frequencies in Hz the unit is connected to.
	Frequencies []uint64 `protobuf:"varint,2,rep,packed,name=frequencies,proto3" json:"frequencies,omitempty"`
}

func (x *GetClientsResponse_Client) Reset() {
	*x = GetClientsResponse_Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_srs_v0_srs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsResponse_Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsResponse_Client) ProtoMessage() {}

func (x *GetClientsResponse_Client) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_srs_v0_srs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsResponse_Client.ProtoReflect.Descriptor instead.
func (*GetClientsResponse_Client) Descriptor() ([]byte, []int) {
	return file_dcs_srs_v0_srs_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetClientsResponse_Client) GetUnit() *common.Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

func (x *GetClientsResponse_Client) GetFrequencies() []uint64 {
	if x != nil {
		return x.Frequencies
	}
	return nil
}

var File_dcs_srs_v0_srs_proto protoreflect.FileDescriptor

var file_dcs_srs_v0_srs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x63, 0x73, 0x2f, 0x73, 0x72, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x63, 0x73, 0x2e, 0x73, 0x72, 0x73, 0x2e,
	0x76, 0x30, 0x1a, 0x1a, 0x64, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9,
	0x05, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x73, 0x6d, 0x6c, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x0f, 0x73, 0x72, 0x73, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0d, 0x73, 0x72, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x09, 0x63, 0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x43, 0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x61,
	0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x33, 0x0a, 0x03,
	0x61, 0x77, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x63, 0x73, 0x2e,
	0x73, 0x72, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x77, 0x73, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77,
	0x73, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x73, 0x72, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x7a,
	0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x06,
	0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64,
	0x63, 0x73, 0x2e, 0x73, 0x72, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x48, 0x00, 0x52, 0x06, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x37, 0x0a, 0x03, 0x77, 0x69,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x73, 0x72,
	0x73, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x48, 0x00, 0x52, 0x03,
	0x77, 0x69, 0x6e, 0x1a, 0x2a, 0x0a, 0x03, 0x41, 0x77, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a,
	0x2c, 0x0a, 0x05, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x2d, 0x0a,
	0x06, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x2e, 0x0a, 0x07,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x72, 0x73, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x22,
	0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64,
	0x63, 0x73, 0x2e, 0x73, 0x72, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x53, 0x0a, 0x06,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x32, 0xa4, 0x01, 0x0a, 0x0a, 0x53, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x2e, 0x64,
	0x63, 0x73, 0x2e, 0x73, 0x72, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x63, 0x73, 0x2e,
	0x73, 0x72, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x73, 0x72,
	0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x73, 0x72, 0x73,
	0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4b, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x43, 0x53, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2f,
	0x67, 0x6f, 0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x64, 0x63, 0x73, 0x2f,
	0x76, 0x30, 0x2f, 0x73, 0x72, 0x73, 0xaa, 0x02, 0x1c, 0x52, 0x75, 0x72, 0x6f, 0x75, 0x6e, 0x69,
	0x4a, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x44, 0x63, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56,
	0x30, 0x2e, 0x53, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dcs_srs_v0_srs_proto_rawDescOnce sync.Once
	file_dcs_srs_v0_srs_proto_rawDescData = file_dcs_srs_v0_srs_proto_rawDesc
)

func file_dcs_srs_v0_srs_proto_rawDescGZIP() []byte {
	file_dcs_srs_v0_srs_proto_rawDescOnce.Do(func() {
		file_dcs_srs_v0_srs_proto_rawDescData = protoimpl.X.CompressGZIP(file_dcs_srs_v0_srs_proto_rawDescData)
	})
	return file_dcs_srs_v0_srs_proto_rawDescData
}

var file_dcs_srs_v0_srs_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_dcs_srs_v0_srs_proto_goTypes = []interface{}{
	(*TransmitRequest)(nil),           // 0: dcs.srs.v0.TransmitRequest
	(*TransmitResponse)(nil),          // 1: dcs.srs.v0.TransmitResponse
	(*GetClientsRequest)(nil),         // 2: dcs.srs.v0.GetClientsRequest
	(*GetClientsResponse)(nil),        // 3: dcs.srs.v0.GetClientsResponse
	(*TransmitRequest_Aws)(nil),       // 4: dcs.srs.v0.TransmitRequest.Aws
	(*TransmitRequest_Azure)(nil),     // 5: dcs.srs.v0.TransmitRequest.Azure
	(*TransmitRequest_GCloud)(nil),    // 6: dcs.srs.v0.TransmitRequest.GCloud
	(*TransmitRequest_Windows)(nil),   // 7: dcs.srs.v0.TransmitRequest.Windows
	(*GetClientsResponse_Client)(nil), // 8: dcs.srs.v0.GetClientsResponse.Client
	(*common.InputPosition)(nil),      // 9: dcs.common.v0.InputPosition
	(common.Coalition)(0),             // 10: dcs.common.v0.Coalition
	(*common.Unit)(nil),               // 11: dcs.common.v0.Unit
}
var file_dcs_srs_v0_srs_proto_depIdxs = []int32{
	9,  // 0: dcs.srs.v0.TransmitRequest.position:type_name -> dcs.common.v0.InputPosition
	10, // 1: dcs.srs.v0.TransmitRequest.coalition:type_name -> dcs.common.v0.Coalition
	4,  // 2: dcs.srs.v0.TransmitRequest.aws:type_name -> dcs.srs.v0.TransmitRequest.Aws
	5,  // 3: dcs.srs.v0.TransmitRequest.azure:type_name -> dcs.srs.v0.TransmitRequest.Azure
	6,  // 4: dcs.srs.v0.TransmitRequest.gcloud:type_name -> dcs.srs.v0.TransmitRequest.GCloud
	7,  // 5: dcs.srs.v0.TransmitRequest.win:type_name -> dcs.srs.v0.TransmitRequest.Windows
	8,  // 6: dcs.srs.v0.GetClientsResponse.clients:type_name -> dcs.srs.v0.GetClientsResponse.Client
	11, // 7: dcs.srs.v0.GetClientsResponse.Client.unit:type_name -> dcs.common.v0.Unit
	0,  // 8: dcs.srs.v0.SrsService.Transmit:input_type -> dcs.srs.v0.TransmitRequest
	2,  // 9: dcs.srs.v0.SrsService.GetClients:input_type -> dcs.srs.v0.GetClientsRequest
	1,  // 10: dcs.srs.v0.SrsService.Transmit:output_type -> dcs.srs.v0.TransmitResponse
	3,  // 11: dcs.srs.v0.SrsService.GetClients:output_type -> dcs.srs.v0.GetClientsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_dcs_srs_v0_srs_proto_init() }
func file_dcs_srs_v0_srs_proto_init() {
	if File_dcs_srs_v0_srs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dcs_srs_v0_srs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitRequest); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitResponse); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsRequest); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsResponse); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitRequest_Aws); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitRequest_Azure); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitRequest_GCloud); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitRequest_Windows); i {
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
		file_dcs_srs_v0_srs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsResponse_Client); i {
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
	file_dcs_srs_v0_srs_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TransmitRequest_Aws_)(nil),
		(*TransmitRequest_Azure_)(nil),
		(*TransmitRequest_Gcloud)(nil),
		(*TransmitRequest_Win)(nil),
	}
	file_dcs_srs_v0_srs_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_dcs_srs_v0_srs_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_dcs_srs_v0_srs_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_dcs_srs_v0_srs_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dcs_srs_v0_srs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dcs_srs_v0_srs_proto_goTypes,
		DependencyIndexes: file_dcs_srs_v0_srs_proto_depIdxs,
		MessageInfos:      file_dcs_srs_v0_srs_proto_msgTypes,
	}.Build()
	File_dcs_srs_v0_srs_proto = out.File
	file_dcs_srs_v0_srs_proto_rawDesc = nil
	file_dcs_srs_v0_srs_proto_goTypes = nil
	file_dcs_srs_v0_srs_proto_depIdxs = nil
}
