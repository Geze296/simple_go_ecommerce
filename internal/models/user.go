package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"unique;size:150;not null"`
	Password  string    `json:"-" gorm:"not null"` 
	Phone     string    `json:"phone" gorm:"size:20"`
	Role      string    `json:"role" gorm:"size:20;default:'customer'"` // customer/admin
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

