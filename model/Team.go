package models

import "github.com/jinzhu/gorm"

// Team struct model
type Team struct {
	gorm.Model
	TeamName   string `json:"teamName"`
	TeamBirth  string `json:"teamBirth"`
	TeamSlogan string `json:"teamSlogan"`
}

// Team Payload mapper
type Teams struct {
	TeamName   string
	TeamBirth  string
	TeamSlogan string
}
