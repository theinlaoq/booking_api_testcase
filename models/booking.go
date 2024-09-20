package models

import (
	"time"
)

type Booking struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `gorm:"index" json:"user_id" validate:"required"`
	StartTime time.Time `gorm:"not null" json:"start_time" validate:"required"`
	EndTime   time.Time `gorm:"not null" json:"end_time" validate:"required,gtfield=StartTime"`
}
