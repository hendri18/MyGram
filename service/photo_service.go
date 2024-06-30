package service

import (
	"MyGram/models"
	"MyGram/repository"
)

type PhotoService struct {
	PhotoRepo *repository.PhotoRepo
}

func (p *PhotoService) Get() ([]*models.Photo, error) {
	return p.PhotoRepo.GetPhoto()
}

func (p *PhotoService) GetById(id uint64) (*models.Photo, error) {
	return p.PhotoRepo.GetPhotoById(id)
}

func (p *PhotoService) Create(photo *models.Photo) (*models.Photo, error) {
	return p.PhotoRepo.CreatePhoto(photo)
}

func (p *PhotoService) Update(id uint64, photo *models.Photo) (*models.Photo, error) {
	return p.PhotoRepo.UpdatePhoto(id, photo)
}

func (p *PhotoService) Delete(id uint64) error {
	return p.PhotoRepo.DeletePhoto(id)
}
