package entity

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name       string
	Age        string `valid:"int~Age is not interger"`
	CustomerID string `valid:"matches(^[C][MU]\\d{8}$)"`
}
