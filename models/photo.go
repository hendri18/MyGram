package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey;"`
	Title     string     `json:"title" form:"title" gorm:"not null" valid:"required~Your title is required"`
	Caption   string     `json:"caption" form:"caption" valid:"required~Your caption is required"`
	PhotoURL  string     `json:"photo_url" form:"photo_url" gorm:"not null" valid:"required~Your photo_url is required"`
	UserID    uint64     `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	User      *User      `json:",omitempty" gorm:"foreignKey:UserID"`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
