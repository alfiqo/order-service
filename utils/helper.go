package utils

import (
	"time"
)

func ValidateDateFormat(date *string) bool {
	_, err := time.Parse(DateFormat, *date)
	return err == nil
}
