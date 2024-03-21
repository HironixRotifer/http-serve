// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: proto/messaggio.proto

package messaggio

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type DeliveryStatus int32

const (
	DeliveryStatus_DELIVERY_STATUS_UNSPECIFIED DeliveryStatus = 0
	DeliveryStatus_DELIVERY_STATUS_SENT        DeliveryStatus = 1
	DeliveryStatus_DELIVERY_STATUS_DELIVERED   DeliveryStatus = 2
	DeliveryStatus_DELIVERY_STATUS_UNDELIVERED DeliveryStatus = 3
)

// Enum value maps for DeliveryStatus.
var (
	DeliveryStatus_name = map[int32]string{
		0: "DELIVERY_STATUS_UNSPECIFIED",
		1: "DELIVERY_STATUS_SENT",
		2: "DELIVERY_STATUS_DELIVERED",
		3: "DELIVERY_STATUS_UNDELIVERED",
	}
	DeliveryStatus_value = map[string]int32{
		"DELIVERY_STATUS_UNSPECIFIED": 0,
		"DELIVERY_STATUS_SENT":        1,
		"DELIVERY_STATUS_DELIVERED":   2,
		"DELIVERY_STATUS_UNDELIVERED": 3,
	}
)

func (x DeliveryStatus) Enum() *DeliveryStatus {
	p := new(DeliveryStatus)
	*p = x
	return p
}

func (x DeliveryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messaggio_proto_enumTypes[0].Descriptor()
}

func (DeliveryStatus) Type() protoreflect.EnumType {
	return &file_proto_messaggio_proto_enumTypes[0]
}

func (x DeliveryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryStatus.Descriptor instead.
func (DeliveryStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{0}
}

type ClaimMessageSendingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intents []*MessageSendingIntent `protobuf:"bytes,1,rep,name=intents,proto3" json:"intents,omitempty"`
}

func (x *ClaimMessageSendingRequest) Reset() {
	*x = ClaimMessageSendingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messaggio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimMessageSendingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimMessageSendingRequest) ProtoMessage() {}

func (x *ClaimMessageSendingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaggio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimMessageSendingRequest.ProtoReflect.Descriptor instead.
func (*ClaimMessageSendingRequest) Descriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimMessageSendingRequest) GetIntents() []*MessageSendingIntent {
	if x != nil {
		return x.Intents
	}
	return nil
}

type ReportMessageDeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reports []*MessageDeliveryReport `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
}

func (x *ReportMessageDeliveryRequest) Reset() {
	*x = ReportMessageDeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messaggio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportMessageDeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportMessageDeliveryRequest) ProtoMessage() {}

func (x *ReportMessageDeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaggio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportMessageDeliveryRequest.ProtoReflect.Descriptor instead.
func (*ReportMessageDeliveryRequest) Descriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{1}
}

func (x *ReportMessageDeliveryRequest) GetReports() []*MessageDeliveryReport {
	if x != nil {
		return x.Reports
	}
	return nil
}

type MessageSendingIntent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel        string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Country        string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Operator       string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	ProjectId      string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	GateId         string `protobuf:"bytes,5,opt,name=gate_id,json=gateId,proto3" json:"gate_id,omitempty"`
	GateSenderId   string `protobuf:"bytes,6,opt,name=gate_sender_id,json=gateSenderId,proto3" json:"gate_sender_id,omitempty"`
	ClientSenderId string `protobuf:"bytes,7,opt,name=client_sender_id,json=clientSenderId,proto3" json:"client_sender_id,omitempty"`
	MessageId      string `protobuf:"bytes,8,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageLabel   string `protobuf:"bytes,9,opt,name=message_label,json=messageLabel,proto3" json:"message_label,omitempty"`
	PhoneNumber    string `protobuf:"bytes,10,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SessionId      string `protobuf:"bytes,11,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	DeliveryTtl    int32  `protobuf:"varint,12,opt,name=delivery_ttl,json=deliveryTtl,proto3" json:"delivery_ttl,omitempty"`
}

func (x *MessageSendingIntent) Reset() {
	*x = MessageSendingIntent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messaggio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendingIntent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendingIntent) ProtoMessage() {}

func (x *MessageSendingIntent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaggio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendingIntent.ProtoReflect.Descriptor instead.
func (*MessageSendingIntent) Descriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{2}
}

func (x *MessageSendingIntent) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *MessageSendingIntent) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MessageSendingIntent) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *MessageSendingIntent) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *MessageSendingIntent) GetGateId() string {
	if x != nil {
		return x.GateId
	}
	return ""
}

func (x *MessageSendingIntent) GetGateSenderId() string {
	if x != nil {
		return x.GateSenderId
	}
	return ""
}

func (x *MessageSendingIntent) GetClientSenderId() string {
	if x != nil {
		return x.ClientSenderId
	}
	return ""
}

func (x *MessageSendingIntent) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *MessageSendingIntent) GetMessageLabel() string {
	if x != nil {
		return x.MessageLabel
	}
	return ""
}

func (x *MessageSendingIntent) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *MessageSendingIntent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *MessageSendingIntent) GetDeliveryTtl() int32 {
	if x != nil {
		return x.DeliveryTtl
	}
	return 0
}

type MessageDeliveryReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel        string              `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Country        string              `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Operator       string              `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	ProjectId      string              `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	GateId         string              `protobuf:"bytes,5,opt,name=gate_id,json=gateId,proto3" json:"gate_id,omitempty"`
	GateSenderId   string              `protobuf:"bytes,6,opt,name=gate_sender_id,json=gateSenderId,proto3" json:"gate_sender_id,omitempty"`
	ClientSenderId string              `protobuf:"bytes,7,opt,name=client_sender_id,json=clientSenderId,proto3" json:"client_sender_id,omitempty"`
	MessageId      string              `protobuf:"bytes,8,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageLabel   string              `protobuf:"bytes,9,opt,name=message_label,json=messageLabel,proto3" json:"message_label,omitempty"`
	PhoneNumber    string              `protobuf:"bytes,10,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SessionId      string              `protobuf:"bytes,11,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionCreated *wrappers.BoolValue `protobuf:"bytes,12,opt,name=session_created,json=sessionCreated,proto3" json:"session_created,omitempty"`
	DeliveryStatus DeliveryStatus      `protobuf:"varint,13,opt,name=delivery_status,json=deliveryStatus,proto3,enum=messaggio.DeliveryStatus" json:"delivery_status,omitempty"`
}

func (x *MessageDeliveryReport) Reset() {
	*x = MessageDeliveryReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messaggio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDeliveryReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDeliveryReport) ProtoMessage() {}

func (x *MessageDeliveryReport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaggio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDeliveryReport.ProtoReflect.Descriptor instead.
func (*MessageDeliveryReport) Descriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{3}
}

func (x *MessageDeliveryReport) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *MessageDeliveryReport) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MessageDeliveryReport) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *MessageDeliveryReport) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *MessageDeliveryReport) GetGateId() string {
	if x != nil {
		return x.GateId
	}
	return ""
}

func (x *MessageDeliveryReport) GetGateSenderId() string {
	if x != nil {
		return x.GateSenderId
	}
	return ""
}

func (x *MessageDeliveryReport) GetClientSenderId() string {
	if x != nil {
		return x.ClientSenderId
	}
	return ""
}

func (x *MessageDeliveryReport) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *MessageDeliveryReport) GetMessageLabel() string {
	if x != nil {
		return x.MessageLabel
	}
	return ""
}

func (x *MessageDeliveryReport) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *MessageDeliveryReport) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *MessageDeliveryReport) GetSessionCreated() *wrappers.BoolValue {
	if x != nil {
		return x.SessionCreated
	}
	return nil
}

func (x *MessageDeliveryReport) GetDeliveryStatus() DeliveryStatus {
	if x != nil {
		return x.DeliveryStatus
	}
	return DeliveryStatus_DELIVERY_STATUS_UNSPECIFIED
}

type BareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Ok    bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BareResponse) Reset() {
	*x = BareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messaggio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BareResponse) ProtoMessage() {}

func (x *BareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaggio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BareResponse.ProtoReflect.Descriptor instead.
func (*BareResponse) Descriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{4}
}

func (x *BareResponse) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BareResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *BareResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message          string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ValidationErrors map[string]string `protobuf:"bytes,3,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messaggio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaggio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_messaggio_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetValidationErrors() map[string]string {
	if x != nil {
		return x.ValidationErrors
	}
	return nil
}

var File_proto_messaggio_proto protoreflect.FileDescriptor

var file_proto_messaggio_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67,
	0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x57, 0x0a, 0x1a, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5a, 0x0a, 0x1c, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x97, 0x03, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x74,
	0x6c, 0x22, 0xfe, 0x03, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x42,
	0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x67, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x5c, 0x0a, 0x0c, 0x42, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x67, 0x69, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0xcf, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x43, 0x0a,
	0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x8b, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52,
	0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45,
	0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03,
	0x32, 0xc8, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x57, 0x0a, 0x13, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x2e,
	0x42, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5b,
	0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x67, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x67, 0x69, 0x6f, 0x2e, 0x42, 0x61, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x2e, 0x61, 0x71, 0x71, 0x2e, 0x6d, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67,
	0x62, 0x69, 0x6c, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x67, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_messaggio_proto_rawDescOnce sync.Once
	file_proto_messaggio_proto_rawDescData = file_proto_messaggio_proto_rawDesc
)

func file_proto_messaggio_proto_rawDescGZIP() []byte {
	file_proto_messaggio_proto_rawDescOnce.Do(func() {
		file_proto_messaggio_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messaggio_proto_rawDescData)
	})
	return file_proto_messaggio_proto_rawDescData
}

var file_proto_messaggio_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_messaggio_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_messaggio_proto_goTypes = []interface{}{
	(DeliveryStatus)(0),                  // 0: messaggio.DeliveryStatus
	(*ClaimMessageSendingRequest)(nil),   // 1: messaggio.ClaimMessageSendingRequest
	(*ReportMessageDeliveryRequest)(nil), // 2: messaggio.ReportMessageDeliveryRequest
	(*MessageSendingIntent)(nil),         // 3: messaggio.MessageSendingIntent
	(*MessageDeliveryReport)(nil),        // 4: messaggio.MessageDeliveryReport
	(*BareResponse)(nil),                 // 5: messaggio.BareResponse
	(*Error)(nil),                        // 6: messaggio.Error
	nil,                                  // 7: messaggio.Error.ValidationErrorsEntry
	(*wrappers.BoolValue)(nil),           // 8: google.protobuf.BoolValue
}
var file_proto_messaggio_proto_depIdxs = []int32{
	3, // 0: messaggio.ClaimMessageSendingRequest.intents:type_name -> messaggio.MessageSendingIntent
	4, // 1: messaggio.ReportMessageDeliveryRequest.reports:type_name -> messaggio.MessageDeliveryReport
	8, // 2: messaggio.MessageDeliveryReport.session_created:type_name -> google.protobuf.BoolValue
	0, // 3: messaggio.MessageDeliveryReport.delivery_status:type_name -> messaggio.DeliveryStatus
	6, // 4: messaggio.BareResponse.error:type_name -> messaggio.Error
	7, // 5: messaggio.Error.validation_errors:type_name -> messaggio.Error.ValidationErrorsEntry
	1, // 6: messaggio.MessaggioBilling.ClaimMessageSending:input_type -> messaggio.ClaimMessageSendingRequest
	2, // 7: messaggio.MessaggioBilling.ReportMessageDelivery:input_type -> messaggio.ReportMessageDeliveryRequest
	5, // 8: messaggio.MessaggioBilling.ClaimMessageSending:output_type -> messaggio.BareResponse
	5, // 9: messaggio.MessaggioBilling.ReportMessageDelivery:output_type -> messaggio.BareResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_messaggio_proto_init() }
func file_proto_messaggio_proto_init() {
	if File_proto_messaggio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_messaggio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimMessageSendingRequest); i {
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
		file_proto_messaggio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportMessageDeliveryRequest); i {
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
		file_proto_messaggio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendingIntent); i {
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
		file_proto_messaggio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDeliveryReport); i {
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
		file_proto_messaggio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BareResponse); i {
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
		file_proto_messaggio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_proto_messaggio_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_messaggio_proto_goTypes,
		DependencyIndexes: file_proto_messaggio_proto_depIdxs,
		EnumInfos:         file_proto_messaggio_proto_enumTypes,
		MessageInfos:      file_proto_messaggio_proto_msgTypes,
	}.Build()
	File_proto_messaggio_proto = out.File
	file_proto_messaggio_proto_rawDesc = nil
	file_proto_messaggio_proto_goTypes = nil
	file_proto_messaggio_proto_depIdxs = nil
}
