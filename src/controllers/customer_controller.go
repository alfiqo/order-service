package controllers

import (
	"errors"
	"net/http"
	"order-service/config"
	"order-service/src/models"
	"order-service/src/responses"
	response "order-service/src/responses/customers"
	"order-service/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/sync/errgroup"
)

var db *gorm.DB = config.NewDB()

func GetAll(context *gin.Context) {
	var customers, data []*models.Customer
	var meta responses.Meta
	search := context.Query("search")

	if err := context.Bind(&meta); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if search != "" {
		db = db.Where("fullname LIKE ? or email LIKE ?", "%"+search+"%", "%"+search+"%").Session(&gorm.Session{})
	}

	eg, _ := errgroup.WithContext(context)

	eg.Go(func() error {
		err := db.Scopes(utils.Paginate(customers, &meta, db)).Find(&customers).Error
		if err != nil {
			return errors.New(err.Error())
		}
		return nil
	})

	eg.Go(func() error {
		err := db.Find(&data).Error
		if err != nil {
			return errors.New(err.Error())
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewListResponse(customers, &meta))
}

func Create(context *gin.Context) {
	var customer models.Customer

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eg, _ := errgroup.WithContext(context)

	if customer.DOB != nil {
		eg.Go(func() error {
			ok := utils.ValidateDateFormat(*customer.DOB)
			if !ok {
				return errors.New(utils.DateFormatInvalid)
			}
			return nil
		})
	}
	eg.Go(func() error {
		exist := db.Where("email = ?", customer.Email).First(&customer)
		if exist.Error != nil && exist.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return errors.New(utils.EmailRegistered)
	})

	if err := eg.Wait(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&customer)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	context.JSON(http.StatusOK, response.NewCreateResponse(&customer))
}
