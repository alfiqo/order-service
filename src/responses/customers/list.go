package customers

import (
	"log"
	"net/http"
	"order-service/src/models"
	"order-service/src/responses"
	"time"
)

type ListResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Meta    interface{}        `json:"meta"`
	Data    []CustomerResponse `json:"data"`
}

func NewListResponse(models []*models.Customer, meta *responses.Meta) ListResponse {
	var response ListResponse

	var customers []CustomerResponse
	for _, customer := range models {

		customerDOB := ""
		if customer.DOB != nil {
			dob, err := time.Parse(time.RFC3339, *customer.DOB)
			if err != nil {
				log.Fatalln(err)
			}
			customerDOB = dob.Format(time.DateOnly)
		}

		customers = append(customers, CustomerResponse{
			ID:        customer.ID,
			Fullname:  customer.Fullname,
			Email:     customer.Email,
			Gender:    customer.Gender,
			DOB:       customerDOB,
			Phone:     customer.Phone,
			Address:   customer.Address,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		})
	}

	response.Code = http.StatusOK
	response.Message = "success"
	response.Meta = meta
	response.Data = customers

	return response
}
