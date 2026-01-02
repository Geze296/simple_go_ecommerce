package models

import "time"


type Order struct{
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

