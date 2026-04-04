package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============== HANDLERS ==============
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// GetProducts handler untuk mendapatkan semua produk
func GetProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// GetProductsByID handler untuk mendapatkan produk berdasarkan ID
func GetProductsByID(c *gin.Context){
	var product Product

	id := c.Param("id")
	if err := DB.First(&product, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"product not found"})
		return

	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct handler untuk mengupdate produk berdasarkan ID
func UpdateProduct(c *gin.Context){
	var product Product
	id := c.Param("id")

	if err := DB.First(&product, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"product not found"})
		return
	}
	c.ShouldBindJSON(&product)
	DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// UpdateProduct handler untuk mengupdate produk berdasarkan ID
func DeleteProduct(c *gin.Context){
		var product Product
	id := c.Param("id")

	if err := DB.First(&product, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"product not found"})
		return
	}
	DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "product has been deleted"})
}

//explanation:
// 1. Setiap handler menerima *gin.Context sebagai parameter untuk mengakses request dan response.
// 2. CreateProduct: Menggunakan ShouldBindJSON untuk binding data JSON ke struct Product, lalu menyimpan ke database.
// 3. GetProducts: Mengambil semua produk dari database dan mengembalikannya sebagai JSON.
// 4. GetProductsByID: Mengambil produk berdasarkan ID yang diberikan di URL, jika tidak ditemukan mengembalikan error 404.
// 5. UpdateProduct: Mengambil produk berdasarkan ID, jika ditemukan melakukan binding data JSON untuk update, lalu menyimpan perubahan ke database.
// 6. DeleteProduct: Mengambil produk berdasarkan ID, jika ditemukan menghapusnya dari database dan mengembalikan pesan sukses.