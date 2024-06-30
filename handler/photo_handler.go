package handler

import (
	"MyGram/models"
	"MyGram/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	PhotoService *service.PhotoService
}

func (p *PhotoHandler) Get(ctx *gin.Context) {
	photos, err := p.PhotoService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   photos,
	})
}

func (p *PhotoHandler) Create(ctx *gin.Context) {

	photoCreate := &models.Photo{}
	if err := ctx.Bind(photoCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))

	photo, err := p.PhotoService.Create(&models.Photo{
		Title:    photoCreate.Title,
		Caption:  photoCreate.Caption,
		PhotoURL: photoCreate.PhotoURL,
		UserID:   userID,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoURL,
			"user_id":    photo.UserID,
			"created_at": photo.CreatedAt,
		},
	})
}

func (p *PhotoHandler) Update(ctx *gin.Context) {

	idx := ctx.Param("photoId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	photoUpdate := &models.Photo{}

	if err := ctx.Bind(photoUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	photo, err := p.PhotoService.Update(uint64(id), &models.Photo{
		Title:    photoUpdate.Title,
		Caption:  photoUpdate.Caption,
		PhotoURL: photoUpdate.PhotoURL,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoURL,
			"user_id":    photo.UserID,
			"updated_at": photo.UpdatedAt,
		},
	})
}

func (p *PhotoHandler) Delete(ctx *gin.Context) {

	idx := ctx.Param("photoId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	err := p.PhotoService.Delete(uint64(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your photo has been successfully deleted",
		},
	})

}
