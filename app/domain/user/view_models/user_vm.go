package user

import (
	"time"

	user "github.com/Rifqi14/golang-ecommerce/app/domain/user/model"
)

type UserVm struct {
	ID         string `json:"user_id"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	VerifiedAt string `json:"verified_at"`
	Status     string `json:"status"`
}

func NewUserVm() UserVm {
	return UserVm{}
}

func (vm UserVm) Build(model *user.User) UserVm {
	return UserVm{
		ID:         model.ID.String(),
		Email:      model.Email,
		CreatedAt:  time.UnixMilli(model.CreatedAt).UTC().String(),
		UpdatedAt:  time.UnixMilli(model.UpdatedAt).UTC().String(),
		VerifiedAt: time.UnixMilli(model.VerifiedAt).UTC().String(),
		Status:     model.Status,
	}
}
