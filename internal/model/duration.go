package model

import (
	"gorm.io/gorm"
)

type Duration struct {
	gorm.Model
	Name     string `json:"name"`
	Cost     uint   `json:"cost"`
	Duration string `json:"duration"`
}
