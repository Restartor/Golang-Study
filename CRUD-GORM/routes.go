package main

import (
	"github.com/gin-gonic/gin"
)

func jalurRouting(r *gin.Engine){

	r.Use(RateLimiter()) // Terapkan middleware rate limiter ke semua route




	product := r.Group("/products")
	{
		product.POST("/", RateLimiter(), CreateProduct) // Terapkan middleware rate limiter hanya untuk route POST /products
		product.GET("/", GetProducts)
		product.GET("/:id", GetProductsByID)
		product.PUT("/:id", UpdateProduct)
		product.DELETE("/:id", DeleteProduct)
	}
}