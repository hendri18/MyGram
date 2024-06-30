package service

import (
	"MyGram/models"
	"MyGram/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func (u *UserService) Get() ([]*models.User, error) {
	return u.UserRepo.GetUser()
}

func (u *UserService) GetByEmail(email string) (*models.User, error) {
	return u.UserRepo.GetUserByEmail(email)
}

func (u *UserService) Create(user *models.User) (*models.User, error) {
	return u.UserRepo.CreateUser(user)
}

func (u *UserService) Update(id uint64, user *models.User) (*models.User, error) {
	return u.UserRepo.UpdateUser(id, user)
}

func (u *UserService) Delete(id uint64) error {
	return u.UserRepo.DeleteUser(id)
}
