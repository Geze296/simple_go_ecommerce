package models

import "time"

type Product struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}