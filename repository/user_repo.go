package repository

import (
	"fmt"

	"github.com/maabodi/goApp/domain"
	"github.com/maabodi/goApp/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserAdapterRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUsers(user model.User) error {
	res := ur.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert!")
	}

	return nil
}

func (ur *userRepository) GetAll() []model.User {
	users := []model.User{}
	ur.DB.Find(&users)

	return users
}

func (ur *userRepository) GetOneByEmail(email string) (user model.User, err error) {
	res := ur.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}
