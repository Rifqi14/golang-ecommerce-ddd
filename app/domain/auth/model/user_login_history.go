package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLoginHistory struct {
	// Column
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Platform  string    `json:"platform"`
	Ip        string    `json:"ip"`
	OS        string    `json:"os"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
