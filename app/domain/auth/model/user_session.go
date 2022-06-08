package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSession struct {
	// Column
	ID               uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID           uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Token            string    `gorm:"type:varchar(100);not null" json:"token"`
	RefreshToken     string    `gorm:"type:varchar(100);not null" json:"refresh_token"`
	ExpiredAt        int64     `gorm:"type:bigint;not null" json:"expired_at"`
	RefreshExpiredAt int64     `gorm:"type:bigint;not null" json:"refresh_expired_at"`
	FirstIp          string    `gorm:"type:varchar(100);not null" json:"first_ip"`
	LastIp           string    `gorm:"type:varchar(100);not null" json:"last_ip"`
	State            bool      `gorm:"type:boolean;not null" json:"state"`
	Platform         string    `json:"platform"`
	CreatedAt        int64
	UpdatedAt        int64
	DeletedAt        gorm.DeletedAt
}
