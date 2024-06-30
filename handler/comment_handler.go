package handler

import (
	"MyGram/models"
	"MyGram/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (p *CommentHandler) Get(ctx *gin.Context) {
	comments, err := p.CommentService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments,
	})
}

func (p *CommentHandler) Create(ctx *gin.Context) {

	commentCreate := &models.Comment{}
	if err := ctx.Bind(commentCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))

	comment, err := p.CommentService.Create(&models.Comment{
		Message: commentCreate.Message,
		PhotoID: commentCreate.PhotoID,
		UserID:  userID,
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
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.CreatedAt,
		},
	})
}

func (p *CommentHandler) Update(ctx *gin.Context) {

	idx := ctx.Param("commentId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	commentUpdate := &models.Comment{}
	if err := ctx.Bind(commentUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	comment, err := p.CommentService.Update(uint64(id), &models.Comment{
		Message: commentUpdate.Message,
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
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"updated_at": comment.UpdatedAt,
		},
	})
}

func (p *CommentHandler) Delete(ctx *gin.Context) {

	idx := ctx.Param("commentId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	err := p.CommentService.Delete(uint64(id))

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
			"message": "Your comment has been successfully deleted",
		},
	})

}
