package customers

import (
	"net/http"
	"order-service/src/models"
)

type CreateResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    CustomerResponse `json:"data"`
}

func NewCreateOrUpdateResponse(model *models.Customer) CreateResponse {
	var response CreateResponse

	response.Code = http.StatusOK
	response.Message = "success"

	var customer CustomerResponse
	customer.ID = model.ID
	customer.Fullname = model.Fullname
	customer.Email = model.Email
	customer.Gender = model.Gender
	customer.DOB = *model.DOB
	customer.Phone = model.Phone
	customer.Address = model.Address
	customer.CreatedAt = model.CreatedAt
	customer.UpdatedAt = model.UpdatedAt

	response.Data = customer

	return response
}
