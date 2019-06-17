package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/harry/jla/model"
)

//get all villan
func (idb *InDB) GetVillans(c *gin.Context) {
	var (
		villans []models.Villan
		result  gin.H
	)

	idb.DB.Find(&villans)
	if len(villans) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": villans,
			"count":  len(villans),
		}
	}
	c.JSON(http.StatusOK, result)

}
