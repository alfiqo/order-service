package controllers

import (
	"errors"
	"net/http"
	"order-service/config"
	"order-service/src/models"
	"order-service/src/responses"
	response "order-service/src/responses/customers"
	"order-service/utils"
	"order-service/utils/constant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/sync/errgroup"
)

type CustomerController struct {
	DB *gorm.DB
}

func NewCustomerController() CustomerController {
	return CustomerController{
		DB: config.NewDB(),
	}
}

func (c *CustomerController) GetAll(context *gin.Context) {
	var customers, data []*models.Customer
	var meta responses.Meta
	search := context.Query("search")
	db := c.DB

	if err := context.Bind(&meta); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eg, _ := errgroup.WithContext(context)

	eg.Go(func() error {
		if search != "" {
			db = c.DB.Where("fullname LIKE ? or email LIKE ?", "%"+search+"%", "%"+search+"%")
		}
		err := db.Scopes(utils.Paginate(customers, &meta, c.DB)).Find(&customers).Error
		if err != nil {
			return errors.New(err.Error())
		}
		return nil
	})

	eg.Go(func() error {
		if search != "" {
			db = c.DB.Where("fullname LIKE ? or email LIKE ?", "%"+search+"%", "%"+search+"%")
		}
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

func (c *CustomerController) Create(context *gin.Context) {
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
				return errors.New(constant.DateFormatInvalid)
			}
			return nil
		})
	}
	eg.Go(func() error {
		exist := c.DB.Where("email = ?", customer.Email).First(&customer)
		if exist.Error != nil && exist.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return errors.New(constant.EmailRegistered)
	})

	if err := eg.Wait(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := c.DB.Create(&customer)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	context.JSON(http.StatusOK, response.NewCreateOrUpdateResponse(&customer))
}

func (c *CustomerController) Detail(context *gin.Context) {
	var customer models.Customer

	id := context.Param("id")

	err := c.DB.Where("id = ?", id).First(&customer).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewDetailResponse(&customer))
}

func (c *CustomerController) Update(context *gin.Context) {
	var customer, data models.Customer

	id := context.Param("id")
	err := c.DB.Where("id = ?", id).First(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eg, _ := errgroup.WithContext(context)
	if customer.DOB != nil {
		eg.Go(func() error {
			ok := utils.ValidateDateFormat(*customer.DOB)
			if !ok {
				return errors.New(constant.DateFormatInvalid)
			}
			return nil
		})
	}
	if customer.Email != data.Email {
		eg.Go(func() error {
			db := c.DB.Where("email = ?", customer.Email).Session(&gorm.Session{})
			err := db.First(&data).Error
			if err != nil && err == gorm.ErrRecordNotFound {
				return nil
			}
			return errors.New(constant.EmailRegistered)
		})
	}

	if err := eg.Wait(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = data.ID
	result := c.DB.Updates(&customer)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	context.JSON(http.StatusOK, response.NewCreateOrUpdateResponse(&customer))
}

func (c *CustomerController) Delete(context *gin.Context) {
	var customer models.Customer

	id := context.Param("id")

	err := c.DB.Where("id = ?", id).First(&customer).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = c.DB.Delete(&customer).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewDeleteResponse())
}
