package routers

import (
	"learn-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	/*
		to run app server.
	*/
	router := gin.Default()

	// connect client with endpoint CreateCar
	router.POST("/cars", controllers.CreateCar)

	// regist endpoint `UpdateCar`
	router.PUT("/cars/:carID", controllers.UpdateCar)

	// regist endpoint `GetCar`
	router.GET("/cars/:carID", controllers.GetCar)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
