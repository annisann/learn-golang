package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	ENDPOINT TO CREATE NEW CAR DATA
*/

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var Cars = []Car{}

// endpoint untuk create data process
func CreateCar(ctx *gin.Context) {
	var newCar Car

	// to bind JSON data (by client) as a request body to server.
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		// return error if error
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// generate ID for new car
	newCar.CarID = fmt.Sprintf("c%d", len(Cars)+1)
	Cars = append(Cars, newCar)

	// send response to client
	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID") // get request parameter from client
	condition := false

	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// find which car wanted to be updated.
	for i, car := range Cars {
		if carID == car.CarID {
			condition = true
			Cars[i] = updatedCar
			Cars[i].CarID = carID
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Car with ID %v is not found.", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with ID %v has been successfully updated.", carID),
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var carData Car

	for i, car := range Cars {
		if carID == car.CarID {
			condition = true
			carData = Cars[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Car with ID %v is not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var carIndex int

	for i, car := range Cars {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Car with ID %v is not found", carID),
		})
		return
	}

	copy(Cars[carIndex:], Cars[carIndex+1:])
	Cars[len(Cars)-1] = Car{}
	Cars = Cars[:len(Cars)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with ID %v has been successfully deleted", carID),
	})
}
