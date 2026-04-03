package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}

// terdapat GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS
///GET adalah untuk mengambil data, 
// POST untuk membuat data baru, 
// PUT untuk memperbarui data yang sudah ada, 
// DELETE untuk menghapus data, 
// PATCH untuk memperbarui sebagian data, 
// HEAD untuk mengambil header respons, 
// dan OPTIONS untuk mengambil informasi tentang metode HTTP yang didukung oleh server.