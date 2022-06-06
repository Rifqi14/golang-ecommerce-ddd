package user

import (
	"time"

	auth "github.com/Rifqi14/golang-ecommerce/app/domain/auth/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// Column
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email        string    `gorm:"type:varchar(100);not null" json:"email"`
	Phone        string    `gorm:"type:varchar(20)" json:"phone"`
	PasswordHash string
	PasswordSalt string
	VerifiedAt   time.Time
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt

	// Relationship
	Session []auth.UserSession      `json:"sessions"`
	Login   []auth.UserLoginHistory `json:"logins"`
}
