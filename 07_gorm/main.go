package main

import (
	"07_gorm/database"
	"07_gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	// create data
	err := db.Create(&User).Error
	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("New user data: ", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	// get the first data
	err := db.First(&user, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user", err)
	}
	fmt.Printf("User data: %+v \n", user)
}

func main() {
	database.StartDB()

	// createUser("hello@doe.com")
	getUserById(1)
}
