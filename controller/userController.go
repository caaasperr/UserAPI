package controller

import (
	"UserAPI/database"
	"UserAPI/model"
	"UserAPI/util"
)

// CRUD User
func CreateUser(id string, password string) bool {
	if IsIdDuplicated(id) {
		return false
	}
	hashedPassword := util.Pass2Hash(password)
	user := model.User{Username: id, Password: hashedPassword}
	database.DB.Create(&user)
	util.Logger.Trace("User created")
	return true
}

func GetUserById(id string) model.User {
	var temp model.User
	database.DB.Where("Username = ?", id).First(&temp)
	return temp
}

func DeleteUser(id string) error {
	var err error
	var temp model.User
	result := database.DB.Where("Username = ?", id).First(&temp)
	if result.Error != nil {
		return err
	}
	database.DB.Delete(&temp)
	return nil
}

// Check Functions
func IsIdDuplicated(id string) bool {
	var temp model.User
	result := database.DB.Where("Username = ?", id).First(&temp)
	if result.Error != nil {
		return false
	}
	return true
}
