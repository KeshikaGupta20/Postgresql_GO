package models

import (
	"github.com/jinzhu/gorm"
)


type Employee struct {
	
	gorm.Model

	Name  string `gorm:"not null" json:"name"`
	Empid string `gorm:"not null" json:"empid"`
	
}

type Response struct {
	Message string `json:"message"`
}

