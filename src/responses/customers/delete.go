package customers

import (
	"net/http"
)

type DeleteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewDeleteResponse() DeleteResponse {
	var response DeleteResponse

	response.Code = http.StatusOK
	response.Message = "success"

	return response
}
