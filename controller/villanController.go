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

//create villan
func (idb *InDB) CreateVillan(c *gin.Context) {
	var (
		villan models.Villan
		result gin.H
	)

	villanName := c.PostForm("VillanName")
	villanPower := c.PostForm("VillanPower")
	villan.VillanName = villanName
	villan.VillanPower = villanPower
	idb.DB.Create(&villan)
	result = gin.H{
		"result": villan,
	}
	c.JSON(http.StatusOK, result)
}
