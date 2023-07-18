package customers

import (
	"net/http"
	"order-service/src/models"
	"time"
)

type CustomerResponse struct {
	ID        uint      `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	DOB       *string   `json:"dob"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    CustomerResponse `json:"data"`
}

func NewCreateResponse(model *models.Customer) CreateResponse {
	var response CreateResponse

	response.Code = http.StatusOK
	response.Message = "success"

	var customer CustomerResponse
	customer.ID = model.ID
	customer.Fullname = model.Fullname
	customer.Email = model.Email
	customer.Gender = model.Gender
	customer.DOB = model.DOB
	customer.Phone = model.Phone
	customer.Address = model.Address
	customer.CreatedAt = model.CreatedAt
	customer.UpdatedAt = model.UpdatedAt

	response.Data = customer

	return response
}
