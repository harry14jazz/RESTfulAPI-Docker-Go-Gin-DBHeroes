package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/harry/jla/model"
)

//get all heroes
func (idb *InDB) GetHeroes(c *gin.Context) {
	var (
		heroes []models.Hero
		result gin.H
	)

	idb.DB.Find(&heroes)
	if len(heroes) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "There are not heroes in db",
		}
	} else {
		result = gin.H{
			"result":  heroes,
			"count":   len(heroes),
			"message": "These are the heroes",
		}
	}
	c.JSON(http.StatusOK, result)
}
