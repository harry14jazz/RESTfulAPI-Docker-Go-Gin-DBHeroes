package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/harry/jla/config"
	controllers "github.com/harry/jla/controller"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/teams", inDB.GetTeams)
	router.GET("/villans", inDB.GetVillans)
	router.GET("/heroes", inDB.GetHeroes)
	router.POST("/villan", inDB.CreateVillan)

	router.Run(":3000")
}
