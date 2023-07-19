package customers

import "time"

type CustomerResponse struct {
	ID        uint      `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	DOB       string    `json:"dob"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
