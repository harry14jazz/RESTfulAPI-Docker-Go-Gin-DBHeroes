package models

import "github.com/jinzhu/gorm"

// Villan struct model
type Villan struct {
	gorm.Model
	VillanName  string `json:"villanName"`
	VillanPower string `json:"villanPower"`
}

// Villan payload mapper
type Villans struct {
	VillanName  string
	VillanPower string
}
