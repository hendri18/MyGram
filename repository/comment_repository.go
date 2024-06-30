package repository

import (
	"MyGram/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepo struct {
	DB *gorm.DB
}

func (c *CommentRepo) GetComment() ([]*models.Comment, error) {
	comments := []*models.Comment{}
	err := c.DB.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "title", "caption", "photo_url", "user_id")
	}).Find(&comments).Error
	return comments, err
}

func (c *CommentRepo) GetCommentById(id uint64) (*models.Comment, error) {
	comment := &models.Comment{}
	result := c.DB.Debug().Where("id = ?", id).Find(&comment)
	err := result.Error
	if result.RowsAffected < 1 {
		err = errors.New("comment not found")
	}
	return comment, err
}

func (c *CommentRepo) CreateComment(comment *models.Comment) (*models.Comment, error) {
	err := c.DB.Debug().Create(&comment).Error
	return comment, err
}

func (c *CommentRepo) UpdateComment(id uint64, comment *models.Comment) (*models.Comment, error) {
	result := c.DB.Debug().
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&models.Comment{
			Message: comment.Message,
		}).Scan(&comment)

	err := result.Error
	if result.RowsAffected < 1 {
		err = errors.New("comment not found")
	}
	return comment, err
}

func (c *CommentRepo) DeleteComment(id uint64) error {
	result := c.DB.
		Where("id = ?", id).
		Delete(&models.Comment{})

	err := result.Error

	if result.RowsAffected < 1 {
		err = errors.New("comment not found")
	}
	return err
}
