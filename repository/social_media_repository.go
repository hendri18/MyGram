package repository

import (
	"MyGram/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepo struct {
	DB *gorm.DB
}

func (s *SocialMediaRepo) GetSocialMedia() ([]*models.SocialMedia, error) {
	socialMedias := []*models.SocialMedia{}
	err := s.DB.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "profile_image_url")
	}).Find(&socialMedias).Error
	return socialMedias, err
}

func (s *SocialMediaRepo) GetSocialMediaById(id uint64) (*models.SocialMedia, error) {
	socialMedia := &models.SocialMedia{}
	result := s.DB.Debug().Where("id = ?", id).Find(&socialMedia)
	err := result.Error
	if err == nil && result.RowsAffected < 1 {
		err = errors.New("social media not found")
	}
	return socialMedia, err
}

func (s *SocialMediaRepo) CreateSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := s.DB.Debug().Create(&socialMedia).Error
	return socialMedia, err
}

func (s *SocialMediaRepo) UpdateSocialMedia(id uint64, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	result := s.DB.Debug().
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&models.SocialMedia{
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
		}).Scan(&socialMedia)

	err := result.Error
	if err == nil && result.RowsAffected < 1 {
		err = errors.New("social media not found")
	}
	return socialMedia, err
}

func (s *SocialMediaRepo) DeleteSocialMedia(id uint64) error {
	result := s.DB.
		Where("id = ?", id).
		Delete(&models.SocialMedia{})

	err := result.Error

	if err == nil && result.RowsAffected < 1 {
		err = errors.New("social media not found")
	}
	return err
}
