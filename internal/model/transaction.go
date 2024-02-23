package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Service       Service   `json:"service" gorm:"foreignKey:ID"`
	Duration      Duration  `json:"duration" gorm:"foreignKey:ID"`
	TotalPrice    int       `json:"total_price"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	TakeOn        time.Time `json:"taken_on"`
}
