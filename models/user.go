package models

import (
	"MyGram/helpers"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID              uint64     `json:"id" gorm:"column:id;primaryKey;"`
	Username        string     `json:"username" form:"username" gorm:"not null;uniqueIndex" valid:"required~Your username is required"`
	Email           string     `json:"email,omitempty" form:"email" gorm:"not null;uniqueIndex" valid:"required~Your email is required,email~Invalid email format"`
	Password        string     `json:"password,omitempty" form:"password" gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age             uint       `json:"age,omitempty" form:"age" gorm:"not null" valid:"required~Your Age is required"`
	ProfileImageURL string     `json:"profile_image_url,omitempty" form:"profile_image_url" gorm:"column:profile_image_url"`
	CreatedAt       *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age <= 8 {
		err = errors.New("age must be more than 8 years")
		return
	}

	hashedPassword, errHash := helpers.HasPass(u.Password)
	u.Password = hashedPassword

	if errHash != nil {
		err = errHash
		return
	}

	err = nil
	return
}
