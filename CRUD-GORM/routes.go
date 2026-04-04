package main

import (
	"github.com/gin-gonic/gin"
)

func jalurRouting(r *gin.Engine){
	product := r.Group("/products")
	{
		product.POST("/", CreateProduct)
		product.GET("/", GetProducts)
		product.GET("/:id", GetProductsByID)
		product.PUT("/:id", UpdateProduct)
		product.DELETE("/:id", DeleteProduct)
	}
}