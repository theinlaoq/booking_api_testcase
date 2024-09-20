package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username" validate:"required,min=3,max=20"`
	Password  string    `gorm:"not null" json:"password" validate:"required,min=8"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Booking   []Booking `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE" json:"booking"`
}
