package v1_repository

import (
	"strings"

	userModel "github.com/Rifqi14/golang-ecommerce/app/domain/user/model"
	userIRepo "github.com/Rifqi14/golang-ecommerce/app/domain/user/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userIRepo.IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo UserRepository) Create(user userModel.User) (res *userModel.User, err error) {
	tx := repo.DB
	err = tx.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) Update(user userModel.User) (res *userModel.User, err error) {
	tx := repo.DB
	err = tx.Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) Delete(user userModel.User) (err error) {
	tx := repo.DB
	err = tx.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepository) ForceDelete(model userModel.User) (err error) {
	tx := repo.DB
	err = tx.Unscoped().Delete(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) GetAll() (res []*userModel.User, err error) {
	tx := repo.DB

	err = tx.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo UserRepository) GetByColumn(column string, value interface{}) (res []*userModel.User, err error) {
	tx := repo.DB

	err = tx.Find(&res, map[string]interface{}{column: value}).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo UserRepository) Detail(id uuid.UUID) (res *userModel.User, err error) {
	tx := repo.DB

	err = tx.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo UserRepository) List(search string, order, sort []string, page, limit int) (res []*userModel.User, total int, err error) {
	tx := repo.DB

	total64 := int64(total)
	err = tx.Count(&total64).Where("lower(email) LIKE ?", "%"+strings.ToLower(search)+"%").Order(order).Offset(page).Limit(limit).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}
