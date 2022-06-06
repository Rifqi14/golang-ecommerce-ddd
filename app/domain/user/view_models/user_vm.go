package user

import user "github.com/Rifqi14/golang-ecommerce/app/domain/user/model"

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
		CreatedAt:  model.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  model.UpdatedAt.Format("2006-01-02 15:04:05"),
		VerifiedAt: model.VerifiedAt.Format("2006-01-02 15:04:05"),
		Status:     model.Status,
	}
}
