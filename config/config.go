package config

import (
	models "github.com/harry/jla/model"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:kurkur14@/jlago?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(models.Villan{}, models.Team{}, models.Hero{})
	return db
}
