package main

import (
	"supplier-portal/controllers"
	"supplier-portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	res := models.ConnectDatabase()

	if res == true {
		r.GET("/", controllers.Home)
		r.GET("/api/user", controllers.GetGoUsers)
		r.GET("/api/user/manipulated", controllers.GoUsersAndManipulateThem)
		r.POST("/api/user", controllers.CreateGoUsers)
		r.PUT("/api/user", controllers.UpdateGoUsers)
	} else {
		r.GET("/", controllers.HomeError)
	}
	r.Run()
}
