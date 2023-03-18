package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gdsc-ys/21days-gin/controllers"
	"github.com/gdsc-ys/21days-gin/database"
)

func main() {
	router := gin.Default()

	database.Connect()

	user := router.Group("/user")
	{
		user.GET("/:id", controllers.ReadUser)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	marker := router.Group("/marker")
	{
		marker.GET("/:id", controllers.ReadMarker)
		marker.POST("/", controllers.CreateMarker)
		marker.PUT("/:id", controllers.UpdateMarker)
		marker.DELETE("/:id", controllers.DeleteMarker)
	}

	router.Run(":8080")
}
