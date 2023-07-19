package controllers

import (
	"errors"
	"net/http"
	"order-service/src/models"
	response "order-service/src/responses/users"
	"order-service/utils"
	"order-service/utils/constant"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func Login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := user.Password

	err := db.Where("username = ?", user.Username).First(&user).Error
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

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eg, _ := errgroup.WithContext(c)
	eg.Go(func() error {
		exist := db.Where("username = ?", user.Username).First(&user)
		if exist.Error != nil && exist.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return errors.New(constant.UsernameTaken)
	})

	if err := eg.Wait(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = password
	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewCreateOrUpdateResponse(&user))
}
