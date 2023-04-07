package main

import (
	"supplier-portal/controllers"
	"supplier-portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/user", controllers.GetGoUsers)
	r.GET("/api/user/manipulated", controllers.GoUsersAndManipulateThem)
	r.POST("/api/user", controllers.CreateGoUsers)
	r.PUT("/api/user", controllers.UpdateGoUsers)

	r.Run()
}
