package user

import (
	userModel "github.com/Rifqi14/golang-ecommerce/app/domain/user/model"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(user userModel.User) (res *userModel.User, err error)

	Update(user userModel.User) (res *userModel.User, err error)

	Delete(user userModel.User) (err error)

	ForceDelete(model userModel.User) (err error)

	GetAll() (res []*userModel.User, err error)

	GetByColumn(column string, value interface{}) (res []*userModel.User, err error)

	Detail(id uuid.UUID) (res *userModel.User, err error)

	List(search string, order, sort []string, page, limit int) (res []*userModel.User, total int, err error)
}
