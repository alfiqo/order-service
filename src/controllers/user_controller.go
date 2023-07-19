package controllers

import (
	"errors"
	"net/http"
	"order-service/src/models"
	response "order-service/src/responses/users"
	"order-service/utils/constant"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

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

	err := user.BeforeSave()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewCreateOrUpdateResponse(&user))
}
