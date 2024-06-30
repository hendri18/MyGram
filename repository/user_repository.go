package repository

import (
	"MyGram/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) GetUser() ([]*models.User, error) {
	users := []*models.User{}
	err := u.DB.Debug().Find(&users).Error
	return users, err
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := u.DB.Debug().Where("email = ?", email).Find(&user)

	err := result.Error

	if result.RowsAffected < 1 {
		err = errors.New("user not found")
	}

	return user, err
}

func (u *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	err := u.DB.Debug().Create(&user).Error
	return user, err
}

func (u *UserRepo) UpdateUser(id uint64, user *models.User) (*models.User, error) {
	result := u.DB.Debug().
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&models.User{
			Email:           user.Email,
			Username:        user.Username,
			ProfileImageURL: user.ProfileImageURL,
		}).Scan(&user)

	err := result.Error
	if result.RowsAffected < 1 {
		err = errors.New("user not found")
	}
	return user, err
}

func (u *UserRepo) DeleteUser(id uint64) error {
	result := u.DB.
		Where("id = ?", id).
		Delete(&models.User{})

	err := result.Error

	if result.RowsAffected < 1 {
		err = errors.New("user not found")
	}
	return err
}
