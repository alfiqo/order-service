package users

import (
	"net/http"
	"order-service/src/models"
)

type CreateResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

func NewCreateOrUpdateResponse(model *models.User) CreateResponse {
	var response CreateResponse

	response.Code = http.StatusOK
	response.Message = "success"

	var user UserResponse
	user.ID = model.ID
	user.Username = model.Username
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt

	response.Data = user

	return response
}
