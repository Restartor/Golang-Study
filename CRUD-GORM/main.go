package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	// Inisialisasi database
	ConnectDB()

	// Inisialisasi router
	r := gin.Default()

	// Setup routes
	jalurRouting(r)


	// Jalankan server di port 1111 
	r.Run(":1111")
}