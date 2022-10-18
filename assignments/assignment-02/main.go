package main

import (
	"assignment-02/config"
	"assignment-02/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:orderID", inDB.UpdateOrder)
	router.DELETE("/orders/:orderID", inDB.DeleteOrder)

	router.Run(":3000")
}
