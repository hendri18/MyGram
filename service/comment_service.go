package service

import (
	"MyGram/models"
	"MyGram/repository"
)

type CommentService struct {
	CommentRepo *repository.CommentRepo
}

func (c *CommentService) Get() ([]*models.Comment, error) {
	return c.CommentRepo.GetComment()
}

func (c *CommentService) GetById(id uint64) (*models.Comment, error) {
	return c.CommentRepo.GetCommentById(id)
}

func (c *CommentService) Create(comment *models.Comment) (*models.Comment, error) {
	return c.CommentRepo.CreateComment(comment)
}

func (c *CommentService) Update(id uint64, comment *models.Comment) (*models.Comment, error) {
	return c.CommentRepo.UpdateComment(id, comment)
}

func (c *CommentService) Delete(id uint64) error {
	return c.CommentRepo.DeleteComment(id)
}
