package repository

import (
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *entity.Users) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) IsUserRegistered(emailOrPhone string) (bool, error) {
	var count int64
	err := r.db.
		Model(entity.Users{}).
		Where("email = ? OR phone = ?", emailOrPhone, emailOrPhone).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.Users, error) {
	var user entity.Users
	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByPhone(phone string) (*entity.Users, error) {
	var user entity.Users
	err := r.db.
		Where("phone = ?", phone).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserById(id int) (*entity.Users, error) {
	var user entity.Users
	err := r.db.
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *entity.Users) error {
	return r.db.Save(user).Error
}
