package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	// Inisialisasi database
	ConnectDB()

	// Inisialisasi router
	r := gin.Default()

	// Setup routes
	jalurRouting(r)


	// Jalankan server di port 1111 
	r.Run(":1111")
}