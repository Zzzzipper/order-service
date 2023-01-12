package entity

import (
	"time"
)

type (
	Order struct {
		ID              uint64        `json:"id" gorm:"id"`
		MerchantOrderId string        `json:"merchant_order_id" binding:"required" gorm:"merchant_order_id"`
		Currency        string        `json:"currency" binding:"required" gorm:"currency"`
		PaymentType     int32         `json:"payment_type" binding:"required" gorm:"payment_type"`
		Amount          int64         `json:"amount" binding:"required" gorm:"amount"`
		Lifetime        time.Duration `json:"lifetime" binding:"required" gorm:"lifetime"`
		Items           string        `json:"items" binding:"required" gorm:"items"`
		Buyer           string        `json:"buyer" binding:"required" gorm:"buyer"`
		UserData        string        `json:"user_data" binding:"required" gorm:"user_data"`
		ShopId          int64         `json:"shop_id" binding:"required" gorm:"shop_id"`
		MerchantId      int64         `json:"merchant_id" binding:"required" gorm:"merchant_id"`
		PartnerId       int64         `json:"partner_id" binding:"required" gorm:"partner_id"`
	}
)
