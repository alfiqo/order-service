package controllers

import (
	"errors"
	"net/http"
	"order-service/config"
	"order-service/src/models"
	response "order-service/src/responses/users"
	"order-service/utils"
	"order-service/utils/constant"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController() UserController {
	return UserController{
		DB: config.NewDB(),
	}
}

func (c *UserController) Login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := user.Password

	err := c.DB.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(&user, password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *UserController) Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eg, _ := errgroup.WithContext(context)
	eg.Go(func() error {
		exist := c.DB.Where("username = ?", user.Username).First(&user)
		if exist.Error != nil && exist.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return errors.New(constant.UsernameTaken)
	})

	if err := eg.Wait(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = password
	err = c.DB.Create(&user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewCreateOrUpdateResponse(&user))
}
