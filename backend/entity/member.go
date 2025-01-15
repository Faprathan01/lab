package entity

import "gorm.io/gorm"

// Member struct with validation
type Member struct {
	gorm.Model
	PhoneNumber string `valid:"length(10)~Phone Number must be exactly 10 digits, matches(^[0-9]+$)~Phone Number must contain only numbers"`
	Password    string `valid:"required~Password is required"`
	Url         string `valid:"url~Url is invalid"`
}
