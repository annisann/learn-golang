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

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Error updating user data: ", err)
		return
	}

	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(userID uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error
	if err != nil {
		fmt.Println("Error creating product data: ", err.Error())
		return
	}

	fmt.Println("New product data: ", Product)
}

func main() {
	database.StartDB()

	// createUser("hello@doe.com")
	// getUserById(1)
	// updateUserById(1, "johndoe@gmail.com")
	createProduct(1, "YLO", "ABC")
}
