package main

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "hello REST API with Gin",
		})
		
	})
}
// SetupRoutes is a function that sets up the routes for the Gin web server. It takes a pointer to a gin.Engine as an argument and defines the routes for the application. In this example, it defines a single route for the root URL ("/") that responds with a JSON message when accessed via a GET request.

// disini kamu bisa menambahkan lebih banyak route sesuai kebutuhan aplikasi kamu, seperti route untuk CRUD (Create, Read, Update, Delete) operasi pada resource tertentu.