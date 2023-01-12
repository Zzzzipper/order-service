// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: payment/api/init_types.proto

package api

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type PayType int32

const (
	PayType_CARD    PayType = 0
	PayType_APPLE   PayType = 1
	PayType_GOOGLE  PayType = 2
	PayType_SAMSUNG PayType = 3
	PayType_IPS     PayType = 4
	PayType_YANDEX  PayType = 5
)

// Enum value maps for PayType.
var (
	PayType_name = map[int32]string{
		0: "CARD",
		1: "APPLE",
		2: "GOOGLE",
		3: "SAMSUNG",
		4: "IPS",
		5: "YANDEX",
	}
	PayType_value = map[string]int32{
		"CARD":    0,
		"APPLE":   1,
		"GOOGLE":  2,
		"SAMSUNG": 3,
		"IPS":     4,
		"YANDEX":  5,
	}
)

func (x PayType) Enum() *PayType {
	p := new(PayType)
	*p = x
	return p
}

func (x PayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayType) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_api_init_types_proto_enumTypes[0].Descriptor()
}

func (PayType) Type() protoreflect.EnumType {
	return &file_payment_api_init_types_proto_enumTypes[0]
}

func (x PayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayType.Descriptor instead.
func (PayType) EnumDescriptor() ([]byte, []int) {
	return file_payment_api_init_types_proto_rawDescGZIP(), []int{0}
}

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор платежа в системе Продавца
	MerchantOrderId string `protobuf:"bytes,1,opt,name=merchant_order_id,json=merchantOrderId,proto3" json:"merchant_order_id,omitempty"`
	// Код валюты ISO 4217
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	// Тип оплаты: CARD, APPLE, GOOGLE, SAMSUNG, IPS
	PaymentType PayType `protobuf:"varint,3,opt,name=payment_type,json=paymentType,proto3,enum=payment.api.PayType" json:"payment_type,omitempty"`
	// Сумма блокировки в минимальных единицах валюты (копейках)
	Amount uint32 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// Срок действия сессии (в секундах), по истечении которого оплата по данной сессии будет невозможна.
	// Если не передан, время жизни сессии устанавливается равным одной неделе
	Lifetime *duration.Duration `protobuf:"bytes,5,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// Список наименований товаров/услуг для отправки в ОФД (54-ФЗ).
	// Общая сумма всех товаров должна соответствовать amount.
	Items []*Item `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	// Идентификация покупателя
	Buyer *Buyer `protobuf:"bytes,7,opt,name=buyer,proto3" json:"buyer,omitempty"`
	// Дополнительные данные пользователя
	UserData *_struct.Struct `protobuf:"bytes,8,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	// Индекс магазина
	ShopId uint64 `protobuf:"varint,9,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// Индекс мерчанта
	MerchantId uint64 `protobuf:"varint,10,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// Индекс партнера
	PartnerId uint64 `protobuf:"varint,11,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_api_init_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_api_init_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_payment_api_init_types_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetMerchantOrderId() string {
	if x != nil {
		return x.MerchantOrderId
	}
	return ""
}

func (x *OrderRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *OrderRequest) GetPaymentType() PayType {
	if x != nil {
		return x.PaymentType
	}
	return PayType_CARD
}

func (x *OrderRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderRequest) GetLifetime() *duration.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *OrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderRequest) GetBuyer() *Buyer {
	if x != nil {
		return x.Buyer
	}
	return nil
}

func (x *OrderRequest) GetUserData() *_struct.Struct {
	if x != nil {
		return x.UserData
	}
	return nil
}

func (x *OrderRequest) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *OrderRequest) GetMerchantId() uint64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *OrderRequest) GetPartnerId() uint64 {
	if x != nil {
		return x.PartnerId
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Индекс заказа в системе продавца
	MerchantOrderId string `protobuf:"bytes,1,opt,name=MerchantOrderId,proto3" json:"MerchantOrderId,omitempty"`
	// Код валюты ISO 4217
	Currency string `protobuf:"bytes,2,opt,name=Currency,proto3" json:"Currency,omitempty"`
	// Сумма резервирования
	Amount uint32 `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	// Тип оплаты: CARD, APPLE, GOOGLE, SAMSUNG, IPS
	PaymentType PayType `protobuf:"varint,4,opt,name=PaymentType,proto3,enum=payment.api.PayType" json:"PaymentType,omitempty"`
	// Список наименований товаров/услуг для отправки в ОФД (54-ФЗ).
	// Общая сумма всех товаров должна соответствовать amount.
	Items []*Item `protobuf:"bytes,5,rep,name=Items,proto3" json:"Items,omitempty"`
	// Идентификация покупателя
	Buyer *Buyer `protobuf:"bytes,6,opt,name=Buyer,proto3" json:"Buyer,omitempty"`
	// Индекс заказа в order-service
	OrderId uint64 `protobuf:"varint,7,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_api_init_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_payment_api_init_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_payment_api_init_types_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetMerchantOrderId() string {
	if x != nil {
		return x.MerchantOrderId
	}
	return ""
}

func (x *Order) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Order) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetPaymentType() PayType {
	if x != nil {
		return x.PaymentType
	}
	return PayType_CARD
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetBuyer() *Buyer {
	if x != nil {
		return x.Buyer
	}
	return nil
}

func (x *Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_payment_api_init_types_proto protoreflect.FileDescriptor

var file_payment_api_init_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e,
	0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x37, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x28, 0x0a, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75,
	0x79, 0x65, 0x72, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x42, 0x75, 0x79, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x05, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x4c, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x50, 0x50, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x41, 0x4d, 0x53, 0x55, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x50, 0x53, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x41, 0x4e, 0x44,
	0x45, 0x58, 0x10, 0x05, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6d,
	0x61, 0x70, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2d, 0x6d, 0x61, 0x70, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_api_init_types_proto_rawDescOnce sync.Once
	file_payment_api_init_types_proto_rawDescData = file_payment_api_init_types_proto_rawDesc
)

func file_payment_api_init_types_proto_rawDescGZIP() []byte {
	file_payment_api_init_types_proto_rawDescOnce.Do(func() {
		file_payment_api_init_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_api_init_types_proto_rawDescData)
	})
	return file_payment_api_init_types_proto_rawDescData
}

var file_payment_api_init_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_api_init_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payment_api_init_types_proto_goTypes = []interface{}{
	(PayType)(0),              // 0: payment.api.PayType
	(*OrderRequest)(nil),      // 1: payment.api.OrderRequest
	(*Order)(nil),             // 2: payment.api.Order
	(*duration.Duration)(nil), // 3: google.protobuf.Duration
	(*Item)(nil),              // 4: payment.api.Item
	(*Buyer)(nil),             // 5: payment.api.Buyer
	(*_struct.Struct)(nil),    // 6: google.protobuf.Struct
}
var file_payment_api_init_types_proto_depIdxs = []int32{
	0, // 0: payment.api.OrderRequest.payment_type:type_name -> payment.api.PayType
	3, // 1: payment.api.OrderRequest.lifetime:type_name -> google.protobuf.Duration
	4, // 2: payment.api.OrderRequest.items:type_name -> payment.api.Item
	5, // 3: payment.api.OrderRequest.buyer:type_name -> payment.api.Buyer
	6, // 4: payment.api.OrderRequest.user_data:type_name -> google.protobuf.Struct
	0, // 5: payment.api.Order.PaymentType:type_name -> payment.api.PayType
	4, // 6: payment.api.Order.Items:type_name -> payment.api.Item
	5, // 7: payment.api.Order.Buyer:type_name -> payment.api.Buyer
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_payment_api_init_types_proto_init() }
func file_payment_api_init_types_proto_init() {
	if File_payment_api_init_types_proto != nil {
		return
	}
	file_payment_api_api_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_payment_api_init_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_payment_api_init_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_payment_api_init_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_api_init_types_proto_goTypes,
		DependencyIndexes: file_payment_api_init_types_proto_depIdxs,
		EnumInfos:         file_payment_api_init_types_proto_enumTypes,
		MessageInfos:      file_payment_api_init_types_proto_msgTypes,
	}.Build()
	File_payment_api_init_types_proto = out.File
	file_payment_api_init_types_proto_rawDesc = nil
	file_payment_api_init_types_proto_goTypes = nil
	file_payment_api_init_types_proto_depIdxs = nil
}
