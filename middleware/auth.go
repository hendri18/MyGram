package middleware

import (
	"MyGram/helpers"
	"MyGram/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")

		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "invalid authorization header",
			})
			return
		}

		if splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "invalid authorization method",
			})
			return
		}

		valid, claims := helpers.ValidateUserJWT(splitToken[1])
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "malformed token",
			})
			return
		}

		ctx.Set("userData", claims)
		ctx.Next()
	}
}

func PhotoAuthorization(photoService *service.PhotoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idx := ctx.Param("photoId")

		if idx == "" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "ID not found",
			})
			return
		}
		id, _ := strconv.Atoi(idx)
		photo, err := photoService.GetById(uint64(id))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Photo not found",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint64(userData["id"].(float64))

		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Access denied, invalid authorization photo",
			})
			return
		}
		ctx.Next()
	}
}

func CommentAuthorization(commentService *service.CommentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idx := ctx.Param("commentId")

		if idx == "" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "ID not found",
			})
			return
		}
		id, _ := strconv.Atoi(idx)
		comment, err := commentService.GetById(uint64(id))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "comment not found",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint64(userData["id"].(float64))

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Access denied, invalid authorization comment",
			})
			return
		}
		ctx.Next()
	}
}

func SocialMediaAuthorization(socialMediaService *service.SocialMediaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idx := ctx.Param("socialMediaId")

		if idx == "" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "ID not found",
			})
			return
		}
		id, _ := strconv.Atoi(idx)
		socialMedia, err := socialMediaService.GetById(uint64(id))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "socialMedia not found",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint64(userData["id"].(float64))

		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Access denied, invalid authorization social media",
			})
			return
		}
		ctx.Next()
	}
}
