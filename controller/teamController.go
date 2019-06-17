package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/harry/jla/model"
)

//get all team
func (idb *InDB) GetTeams(c *gin.Context) {
	var (
		teams  []models.Team
		result gin.H
	)

	idb.DB.Find(&teams)
	if len(teams) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "There are no teams in db",
		}
	} else {
		result = gin.H{
			"result":  teams,
			"count":   len(teams),
			"message": "These are the teams",
		}
	}
	c.JSON(http.StatusOK, result)
}
