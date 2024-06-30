package service

import (
	"MyGram/models"
	"MyGram/repository"
)

type SocialMediaService struct {
	SocialMediaRepo *repository.SocialMediaRepo
}

func (s *SocialMediaService) Get() ([]*models.SocialMedia, error) {
	return s.SocialMediaRepo.GetSocialMedia()
}

func (s *SocialMediaService) GetById(id uint64) (*models.SocialMedia, error) {
	return s.SocialMediaRepo.GetSocialMediaById(id)
}

func (s *SocialMediaService) Create(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	return s.SocialMediaRepo.CreateSocialMedia(socialMedia)
}

func (s *SocialMediaService) Update(id uint64, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	return s.SocialMediaRepo.UpdateSocialMedia(id, socialMedia)
}

func (s *SocialMediaService) Delete(id uint64) error {
	return s.SocialMediaRepo.DeleteSocialMedia(id)
}
