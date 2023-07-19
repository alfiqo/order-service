package customers

import (
	"net/http"
	"order-service/src/models"
)

type DetailResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    CustomerResponse `json:"data"`
}

func NewDetailResponse(model *models.Customer) DetailResponse {
	var response DetailResponse

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
