package entity

import "gorm.io/gorm"

// Member struct with validation
type Member struct {
	PhoneNumber string `valid:"length(10),matches(^[0-9]+$)~Phone Number length is not 10 digits."`
	Password 	string
	Url         string `valid:"url~Url is invalid"`
}