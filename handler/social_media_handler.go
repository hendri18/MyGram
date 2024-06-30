package handler

import (
	"MyGram/models"
	"MyGram/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandler struct {
	SocialMediaService *service.SocialMediaService
}

func (p *SocialMediaHandler) Get(ctx *gin.Context) {
	socialMedias, err := p.SocialMediaService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   socialMedias,
	})
}

func (p *SocialMediaHandler) Create(ctx *gin.Context) {

	socialMediaCreate := &models.SocialMedia{}
	if err := ctx.Bind(socialMediaCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))

	socialMedia, err := p.SocialMediaService.Create(&models.SocialMedia{
		Name:           socialMediaCreate.Name,
		SocialMediaURL: socialMediaCreate.SocialMediaURL,
		UserID:         userID,
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
			"id":               socialMedia.ID,
			"name":             socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaURL,
			"user_id":          socialMedia.UserID,
			"created_at":       socialMedia.CreatedAt,
		},
	})
}

func (p *SocialMediaHandler) Update(ctx *gin.Context) {

	idx := ctx.Param("socialMediaId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	socialMediaUpdate := &models.SocialMedia{}
	if err := ctx.Bind(socialMediaUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	socialMedia, err := p.SocialMediaService.Update(uint64(id), &models.SocialMedia{
		Name:           socialMediaUpdate.Name,
		SocialMediaURL: socialMediaUpdate.SocialMediaURL,
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
			"id":               socialMedia.ID,
			"name":             socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaURL,
			"user_id":          socialMedia.UserID,
			"updated_at":       socialMedia.UpdatedAt,
		},
	})
}

func (p *SocialMediaHandler) Delete(ctx *gin.Context) {

	idx := ctx.Param("socialMediaId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	err := p.SocialMediaService.Delete(uint64(id))

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
			"message": "Your social media has been successfully deleted",
		},
	})

}
