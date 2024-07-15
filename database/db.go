package database

import (
	"UserAPI/model"
	"UserAPI/util"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() error {
	var err error
	dsn := "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	connection := fmt.Sprintf(dsn, os.Getenv("DB-Username"), os.Getenv("DB-Password"), os.Getenv("DB-Name"))
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		util.Logger.Error("DB Conncection Error!")
		return err
	}
	DB.AutoMigrate(&model.User{})
	util.Logger.Info("DB Connected Succesfully!")
	return nil
}
