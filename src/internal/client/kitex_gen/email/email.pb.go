// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.0--rc2
// source: email.proto

package email

import (
	context "context"
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

type Recipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Recipient) Reset() {
	*x = Recipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipient) ProtoMessage() {}

func (x *Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipient.ProtoReflect.Descriptor instead.
func (*Recipient) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{0}
}

func (x *Recipient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recipient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To        []*Recipient `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Cc        []*Recipient `protobuf:"bytes,2,rep,name=cc,proto3" json:"cc,omitempty"`
	Bcc       []*Recipient `protobuf:"bytes,3,rep,name=bcc,proto3" json:"bcc,omitempty"`
	Subject   string       `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      string       `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	FromName  string       `protobuf:"bytes,6,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	FromEmail string       `protobuf:"bytes,7,opt,name=from_email,json=fromEmail,proto3" json:"from_email,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailRequest) GetTo() []*Recipient {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendEmailRequest) GetCc() []*Recipient {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *SendEmailRequest) GetBcc() []*Recipient {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *SendEmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendEmailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SendEmailRequest) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *SendEmailRequest) GetFromEmail() string {
	if x != nil {
		return x.FromEmail
	}
	return ""
}

type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Success   bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SendEmailResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendEmailResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_email_proto protoreflect.FileDescriptor

var file_email_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x35, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe4, 0x01, 0x0a, 0x10,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x20, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x02, 0x63, 0x63, 0x12, 0x22, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x03, 0x62, 0x63, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x62, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x50, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x69, 0x6e, 0x6e, 0x54, 0x65, 0x77, 0x2f, 0x46,
	0x69, 0x6e, 0x6e, 0x45, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_proto_rawDescOnce sync.Once
	file_email_proto_rawDescData = file_email_proto_rawDesc
)

func file_email_proto_rawDescGZIP() []byte {
	file_email_proto_rawDescOnce.Do(func() {
		file_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_proto_rawDescData)
	})
	return file_email_proto_rawDescData
}

var file_email_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_email_proto_goTypes = []interface{}{
	(*Recipient)(nil),         // 0: email.Recipient
	(*SendEmailRequest)(nil),  // 1: email.SendEmailRequest
	(*SendEmailResponse)(nil), // 2: email.SendEmailResponse
}
var file_email_proto_depIdxs = []int32{
	0, // 0: email.SendEmailRequest.to:type_name -> email.Recipient
	0, // 1: email.SendEmailRequest.cc:type_name -> email.Recipient
	0, // 2: email.SendEmailRequest.bcc:type_name -> email.Recipient
	1, // 3: email.EmailService.SendEmail:input_type -> email.SendEmailRequest
	2, // 4: email.EmailService.SendEmail:output_type -> email.SendEmailResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_email_proto_init() }
func file_email_proto_init() {
	if File_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipient); i {
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
		file_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
			RawDescriptor: file_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_proto_goTypes,
		DependencyIndexes: file_email_proto_depIdxs,
		MessageInfos:      file_email_proto_msgTypes,
	}.Build()
	File_email_proto = out.File
	file_email_proto_rawDesc = nil
	file_email_proto_goTypes = nil
	file_email_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type EmailService interface {
	SendEmail(ctx context.Context, req *SendEmailRequest) (res *SendEmailResponse, err error)
}
