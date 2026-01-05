package models

import "time"

type Merchant struct {
	MerchantID uint `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	Balance float64 `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
