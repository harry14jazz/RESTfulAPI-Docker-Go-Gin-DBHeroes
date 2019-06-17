package models

import "github.com/jinzhu/gorm"

// Hero struct model
type Hero struct {
	gorm.Model
	HeroName      string `json:"heroName"`
	HeroFullName  string `json:"heroFullName"`
	HeroPlace     string `json:"heroPlace"`
	HeroAlias     string `json:"heroAlias"`
	HeroAlignment string `json:"heroAlignment"`
	Team          []Team `gorm:"many2many:hero_teams"`
}

// Hero payload mapper
type Heroes struct {
	HeroName      string
	HeroFullName  string
	HeroPlace     string
	HeroAlias     string
	HeroAlignment string
	Team          []Team
}
