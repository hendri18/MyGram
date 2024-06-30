package handler

import (
	"MyGram/helpers"
	"MyGram/models"
	"MyGram/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func (u *UserHandler) Login(ctx *gin.Context) {
	userLogin := &models.User{}

	if err := ctx.Bind(userLogin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := u.UserService.GetByEmail(userLogin.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	comparePass := helpers.CheckPasswordHash(userLogin.Password, user.Password)

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "invalid email or password",
		})
		return
	}

	token, _ := helpers.GenerateUserJWT(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"token": token,
		},
	})
}

func (u *UserHandler) Register(ctx *gin.Context) {
	userRegister := &models.User{}

	if err := ctx.Bind(userRegister); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := u.UserService.Create(userRegister)

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
			"age":      user.Age,
			"email":    user.Email,
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func (p *UserHandler) Update(ctx *gin.Context) {

	idx := ctx.Param("userId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	userUpdate := &models.User{}
	if err := ctx.Bind(userUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := p.UserService.Update(uint64(id), &models.User{
		Email:           userUpdate.Email,
		Username:        userUpdate.Username,
		ProfileImageURL: userUpdate.ProfileImageURL,
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
			"id":         user.ID,
			"email":      user.Email,
			"username":   user.Username,
			"age":        user.Age,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (p *UserHandler) Delete(ctx *gin.Context) {

	idx := ctx.Param("userId")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	err := p.UserService.Delete(uint64(id))

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
			"message": "Yout account has been successfully deleted",
		},
	})

}

func (p *UserHandler) DeleteWithoutID(ctx *gin.Context) {

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))

	err := p.UserService.Delete(userID)

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
			"message": "Yout account has been successfully deleted",
		},
	})

}
