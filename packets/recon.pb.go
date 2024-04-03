// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: recon.proto

package packets

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

type BankReconRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode               string                 `protobuf:"bytes,1,opt,name=TenantCode,proto3" json:"TenantCode,omitempty"`
	BankCode                 string                 `protobuf:"bytes,2,opt,name=BankCode,proto3" json:"BankCode,omitempty"`
	TargetBankCode           string                 `protobuf:"bytes,3,opt,name=TargetBankCode,proto3" json:"TargetBankCode,omitempty"`
	ReconcileType            string                 `protobuf:"bytes,4,opt,name=ReconcileType,proto3" json:"ReconcileType,omitempty"`
	TransactionNumber        string                 `protobuf:"bytes,5,opt,name=TransactionNumber,proto3" json:"TransactionNumber,omitempty"`
	PaymentNumber            string                 `protobuf:"bytes,6,opt,name=PaymentNumber,proto3" json:"PaymentNumber,omitempty"`
	HomingAccountNumber      string                 `protobuf:"bytes,7,opt,name=HomingAccountNumber,proto3" json:"HomingAccountNumber,omitempty"`
	BeneficiaryAccountNumber string                 `protobuf:"bytes,8,opt,name=BeneficiaryAccountNumber,proto3" json:"BeneficiaryAccountNumber,omitempty"`
	BeneficiaryAccountName   string                 `protobuf:"bytes,9,opt,name=BeneficiaryAccountName,proto3" json:"BeneficiaryAccountName,omitempty"`
	BankAccountNumber        string                 `protobuf:"bytes,10,opt,name=BankAccountNumber,proto3" json:"BankAccountNumber,omitempty"`
	BankTransactionNumber    string                 `protobuf:"bytes,11,opt,name=BankTransactionNumber,proto3" json:"BankTransactionNumber,omitempty"`
	BankTransactionId        string                 `protobuf:"bytes,12,opt,name=BankTransactionId,proto3" json:"BankTransactionId,omitempty"`
	BankReferenceNumber      string                 `protobuf:"bytes,13,opt,name=BankReferenceNumber,proto3" json:"BankReferenceNumber,omitempty"`
	BankBranchCode           string                 `protobuf:"bytes,14,opt,name=BankBranchCode,proto3" json:"BankBranchCode,omitempty"`
	PayRecvAccountNumber     string                 `protobuf:"bytes,15,opt,name=PayRecvAccountNumber,proto3" json:"PayRecvAccountNumber,omitempty"`
	MerchantId               string                 `protobuf:"bytes,16,opt,name=MerchantId,proto3" json:"MerchantId,omitempty"`
	TransactionAmount        float64                `protobuf:"fixed64,17,opt,name=TransactionAmount,proto3" json:"TransactionAmount,omitempty"`
	CurrencyCode             string                 `protobuf:"bytes,18,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	BankReferenceInfo        string                 `protobuf:"bytes,19,opt,name=BankReferenceInfo,proto3" json:"BankReferenceInfo,omitempty"`
	TransactionAt            *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=TransactionAt,proto3" json:"TransactionAt,omitempty"`
	TransactionMethod        string                 `protobuf:"bytes,21,opt,name=TransactionMethod,proto3" json:"TransactionMethod,omitempty"`
	TransactionStatus        string                 `protobuf:"bytes,22,opt,name=TransactionStatus,proto3" json:"TransactionStatus,omitempty"`
	TransactionNote          string                 `protobuf:"bytes,23,opt,name=TransactionNote,proto3" json:"TransactionNote,omitempty"`
}

func (x *BankReconRegisterRequest) Reset() {
	*x = BankReconRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankReconRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankReconRegisterRequest) ProtoMessage() {}

func (x *BankReconRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankReconRegisterRequest.ProtoReflect.Descriptor instead.
func (*BankReconRegisterRequest) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{0}
}

func (x *BankReconRegisterRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankCode() string {
	if x != nil {
		return x.BankCode
	}
	return ""
}

func (x *BankReconRegisterRequest) GetTargetBankCode() string {
	if x != nil {
		return x.TargetBankCode
	}
	return ""
}

func (x *BankReconRegisterRequest) GetReconcileType() string {
	if x != nil {
		return x.ReconcileType
	}
	return ""
}

func (x *BankReconRegisterRequest) GetTransactionNumber() string {
	if x != nil {
		return x.TransactionNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetPaymentNumber() string {
	if x != nil {
		return x.PaymentNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetHomingAccountNumber() string {
	if x != nil {
		return x.HomingAccountNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBeneficiaryAccountNumber() string {
	if x != nil {
		return x.BeneficiaryAccountNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBeneficiaryAccountName() string {
	if x != nil {
		return x.BeneficiaryAccountName
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankAccountNumber() string {
	if x != nil {
		return x.BankAccountNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankTransactionNumber() string {
	if x != nil {
		return x.BankTransactionNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankTransactionId() string {
	if x != nil {
		return x.BankTransactionId
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankReferenceNumber() string {
	if x != nil {
		return x.BankReferenceNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankBranchCode() string {
	if x != nil {
		return x.BankBranchCode
	}
	return ""
}

func (x *BankReconRegisterRequest) GetPayRecvAccountNumber() string {
	if x != nil {
		return x.PayRecvAccountNumber
	}
	return ""
}

func (x *BankReconRegisterRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *BankReconRegisterRequest) GetTransactionAmount() float64 {
	if x != nil {
		return x.TransactionAmount
	}
	return 0
}

func (x *BankReconRegisterRequest) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *BankReconRegisterRequest) GetBankReferenceInfo() string {
	if x != nil {
		return x.BankReferenceInfo
	}
	return ""
}

func (x *BankReconRegisterRequest) GetTransactionAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionAt
	}
	return nil
}

func (x *BankReconRegisterRequest) GetTransactionMethod() string {
	if x != nil {
		return x.TransactionMethod
	}
	return ""
}

func (x *BankReconRegisterRequest) GetTransactionStatus() string {
	if x != nil {
		return x.TransactionStatus
	}
	return ""
}

func (x *BankReconRegisterRequest) GetTransactionNote() string {
	if x != nil {
		return x.TransactionNote
	}
	return ""
}

type BankReconRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *PStatus `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *BankReconRegisterResponse) Reset() {
	*x = BankReconRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankReconRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankReconRegisterResponse) ProtoMessage() {}

func (x *BankReconRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankReconRegisterResponse.ProtoReflect.Descriptor instead.
func (*BankReconRegisterResponse) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{1}
}

func (x *BankReconRegisterResponse) GetStatus() *PStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type BankReconModifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode        string `protobuf:"bytes,1,opt,name=TenantCode,proto3" json:"TenantCode,omitempty"`
	BankCode          string `protobuf:"bytes,2,opt,name=BankCode,proto3" json:"BankCode,omitempty"`
	ReconcileType     string `protobuf:"bytes,3,opt,name=ReconcileType,proto3" json:"ReconcileType,omitempty"`
	TransactionNumber string `protobuf:"bytes,4,opt,name=TransactionNumber,proto3" json:"TransactionNumber,omitempty"`
	TransactionStatus string `protobuf:"bytes,5,opt,name=TransactionStatus,proto3" json:"TransactionStatus,omitempty"`
}

func (x *BankReconModifyRequest) Reset() {
	*x = BankReconModifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankReconModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankReconModifyRequest) ProtoMessage() {}

func (x *BankReconModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankReconModifyRequest.ProtoReflect.Descriptor instead.
func (*BankReconModifyRequest) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{2}
}

func (x *BankReconModifyRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *BankReconModifyRequest) GetBankCode() string {
	if x != nil {
		return x.BankCode
	}
	return ""
}

func (x *BankReconModifyRequest) GetReconcileType() string {
	if x != nil {
		return x.ReconcileType
	}
	return ""
}

func (x *BankReconModifyRequest) GetTransactionNumber() string {
	if x != nil {
		return x.TransactionNumber
	}
	return ""
}

func (x *BankReconModifyRequest) GetTransactionStatus() string {
	if x != nil {
		return x.TransactionStatus
	}
	return ""
}

type BankReconModifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *PStatus `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *BankReconModifyResponse) Reset() {
	*x = BankReconModifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankReconModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankReconModifyResponse) ProtoMessage() {}

func (x *BankReconModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankReconModifyResponse.ProtoReflect.Descriptor instead.
func (*BankReconModifyResponse) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{3}
}

func (x *BankReconModifyResponse) GetStatus() *PStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ProviderReconRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantCode           string                 `protobuf:"bytes,1,opt,name=TenantCode,proto3" json:"TenantCode,omitempty"`
	ProviderCode         string                 `protobuf:"bytes,2,opt,name=ProviderCode,proto3" json:"ProviderCode,omitempty"`
	PaymentNumber        string                 `protobuf:"bytes,3,opt,name=PaymentNumber,proto3" json:"PaymentNumber,omitempty"`
	PaymentLinkNumber    string                 `protobuf:"bytes,4,opt,name=PaymentLinkNumber,proto3" json:"PaymentLinkNumber,omitempty"`
	PaymentMethod        string                 `protobuf:"bytes,5,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	ProviderReferenceId  string                 `protobuf:"bytes,6,opt,name=ProviderReferenceId,proto3" json:"ProviderReferenceId,omitempty"`
	PartnerReferenceId   string                 `protobuf:"bytes,7,opt,name=PartnerReferenceId,proto3" json:"PartnerReferenceId,omitempty"`
	TransactionAmount    float64                `protobuf:"fixed64,8,opt,name=TransactionAmount,proto3" json:"TransactionAmount,omitempty"`
	CurrencyCode         string                 `protobuf:"bytes,9,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	TransactionAt        *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=TransactionAt,proto3" json:"TransactionAt,omitempty"`
	PaymentRequestStatus string                 `protobuf:"bytes,11,opt,name=PaymentRequestStatus,proto3" json:"PaymentRequestStatus,omitempty"`
}

func (x *ProviderReconRegisterRequest) Reset() {
	*x = ProviderReconRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderReconRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderReconRegisterRequest) ProtoMessage() {}

func (x *ProviderReconRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderReconRegisterRequest.ProtoReflect.Descriptor instead.
func (*ProviderReconRegisterRequest) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{4}
}

func (x *ProviderReconRegisterRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetProviderCode() string {
	if x != nil {
		return x.ProviderCode
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetPaymentNumber() string {
	if x != nil {
		return x.PaymentNumber
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetPaymentLinkNumber() string {
	if x != nil {
		return x.PaymentLinkNumber
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetProviderReferenceId() string {
	if x != nil {
		return x.ProviderReferenceId
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetPartnerReferenceId() string {
	if x != nil {
		return x.PartnerReferenceId
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetTransactionAmount() float64 {
	if x != nil {
		return x.TransactionAmount
	}
	return 0
}

func (x *ProviderReconRegisterRequest) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *ProviderReconRegisterRequest) GetTransactionAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionAt
	}
	return nil
}

func (x *ProviderReconRegisterRequest) GetPaymentRequestStatus() string {
	if x != nil {
		return x.PaymentRequestStatus
	}
	return ""
}

type ProviderReconRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *PStatus `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *ProviderReconRegisterResponse) Reset() {
	*x = ProviderReconRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderReconRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderReconRegisterResponse) ProtoMessage() {}

func (x *ProviderReconRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderReconRegisterResponse.ProtoReflect.Descriptor instead.
func (*ProviderReconRegisterResponse) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{5}
}

func (x *ProviderReconRegisterResponse) GetStatus() *PStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_recon_proto protoreflect.FileDescriptor

var file_recon_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x08, 0x0a, 0x18, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x61,
	0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x13, 0x48, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x48,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x18, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x36,
	0x0a, 0x16, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x42, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x42, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x61,
	0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x42, 0x61,
	0x6e, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x52, 0x65, 0x63, 0x76, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x50, 0x61, 0x79, 0x52, 0x65, 0x63, 0x76, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x45,
	0x0a, 0x19, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x16, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43,
	0x0a, 0x17, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2e, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x86, 0x04, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x1d,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa7, 0x02, 0x0a, 0x05, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x12, 0x5c, 0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x56, 0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x42, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x42, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x25, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recon_proto_rawDescOnce sync.Once
	file_recon_proto_rawDescData = file_recon_proto_rawDesc
)

func file_recon_proto_rawDescGZIP() []byte {
	file_recon_proto_rawDescOnce.Do(func() {
		file_recon_proto_rawDescData = protoimpl.X.CompressGZIP(file_recon_proto_rawDescData)
	})
	return file_recon_proto_rawDescData
}

var file_recon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_recon_proto_goTypes = []interface{}{
	(*BankReconRegisterRequest)(nil),      // 0: packets.BankReconRegisterRequest
	(*BankReconRegisterResponse)(nil),     // 1: packets.BankReconRegisterResponse
	(*BankReconModifyRequest)(nil),        // 2: packets.BankReconModifyRequest
	(*BankReconModifyResponse)(nil),       // 3: packets.BankReconModifyResponse
	(*ProviderReconRegisterRequest)(nil),  // 4: packets.ProviderReconRegisterRequest
	(*ProviderReconRegisterResponse)(nil), // 5: packets.ProviderReconRegisterResponse
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(*PStatus)(nil),                       // 7: packets.PStatus
}
var file_recon_proto_depIdxs = []int32{
	6, // 0: packets.BankReconRegisterRequest.TransactionAt:type_name -> google.protobuf.Timestamp
	7, // 1: packets.BankReconRegisterResponse.Status:type_name -> packets.PStatus
	7, // 2: packets.BankReconModifyResponse.Status:type_name -> packets.PStatus
	6, // 3: packets.ProviderReconRegisterRequest.TransactionAt:type_name -> google.protobuf.Timestamp
	7, // 4: packets.ProviderReconRegisterResponse.Status:type_name -> packets.PStatus
	0, // 5: packets.Recon.BankReconRegister:input_type -> packets.BankReconRegisterRequest
	2, // 6: packets.Recon.BankReconModify:input_type -> packets.BankReconModifyRequest
	4, // 7: packets.Recon.ProviderReconRegister:input_type -> packets.ProviderReconRegisterRequest
	1, // 8: packets.Recon.BankReconRegister:output_type -> packets.BankReconRegisterResponse
	3, // 9: packets.Recon.BankReconModify:output_type -> packets.BankReconModifyResponse
	5, // 10: packets.Recon.ProviderReconRegister:output_type -> packets.ProviderReconRegisterResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_recon_proto_init() }
func file_recon_proto_init() {
	if File_recon_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_recon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankReconRegisterRequest); i {
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
		file_recon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankReconRegisterResponse); i {
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
		file_recon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankReconModifyRequest); i {
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
		file_recon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankReconModifyResponse); i {
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
		file_recon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderReconRegisterRequest); i {
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
		file_recon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderReconRegisterResponse); i {
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
			RawDescriptor: file_recon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recon_proto_goTypes,
		DependencyIndexes: file_recon_proto_depIdxs,
		MessageInfos:      file_recon_proto_msgTypes,
	}.Build()
	File_recon_proto = out.File
	file_recon_proto_rawDesc = nil
	file_recon_proto_goTypes = nil
	file_recon_proto_depIdxs = nil
}
