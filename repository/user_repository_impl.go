package repository

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func (r *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Order("id desc").Find(&users).Error
	helper.PanicIfError(err)

	return users, nil
}

func (r *UserRepositoryImpl) FindByID(id int) (domain.User, error) {
	user := domain.User{}
	err := r.DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := r.DB.Where("email=?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}
