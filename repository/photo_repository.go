package repository

import (
	"MyGram/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepo struct {
	DB *gorm.DB
}

func (p *PhotoRepo) GetPhoto() ([]*models.Photo, error) {
	photos := []*models.Photo{}
	err := p.DB.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "username")
	}).Find(&photos).Error
	return photos, err
}

func (p *PhotoRepo) GetPhotoById(id uint64) (*models.Photo, error) {
	photo := &models.Photo{}
	result := p.DB.Debug().Where("id = ?", id).Find(&photo)
	err := result.Error
	if err == nil && result.RowsAffected < 1 {
		err = errors.New("photo not found")
	}
	return photo, err
}

func (p *PhotoRepo) CreatePhoto(photo *models.Photo) (*models.Photo, error) {
	err := p.DB.Debug().Create(&photo).Error
	return photo, err
}

func (p *PhotoRepo) UpdatePhoto(id uint64, photo *models.Photo) (*models.Photo, error) {
	result := p.DB.Debug().
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&models.Photo{
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
		}).Scan(&photo)

	err := result.Error
	if err == nil && result.RowsAffected < 1 {
		err = errors.New("photo not found")
	}
	return photo, err
}

func (p *PhotoRepo) DeletePhoto(id uint64) error {
	result := p.DB.
		Where("id = ?", id).
		Delete(&models.Photo{})

	err := result.Error

	if err == nil && result.RowsAffected < 1 {
		err = errors.New("photo not found")
	}
	return err
}
