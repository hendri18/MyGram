package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint64     `json:"id" gorm:"column:id;primaryKey;"`
	Name           string     `json:"name" form:"name" gorm:"not null;" valid:"required~Your name is required"`
	SocialMediaURL string     `json:"social_media_url" form:"social_media_url" gorm:"not null;" valid:"required~Your social media url is required"`
	UserID         uint64     `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt      *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	User           *User
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
