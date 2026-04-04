# Notes CRUD GORM (Gin + PostgreSQL)

Catatan ini dibuat berdasarkan implementasi di folder ini:
- main.go
- db.go
- models.go
- routes.go
- handlers.go

## 1. Gambaran Alur Dari Nol

1. Aplikasi start di `main()`.
2. Fungsi `ConnectDB()` dipanggil untuk konek ke PostgreSQL.
3. `AutoMigrate(&Product{})` membuat/update tabel `products` otomatis.
4. Gin router dibuat dengan `gin.Default()`.
5. Route CRUD didaftarkan lewat `jalurRouting(r)`.
6. Server berjalan di port `1111`.

Alur request saat API dipanggil:
1. Request masuk ke route tertentu (misalnya `POST /products/`).
2. Handler membaca data dari JSON atau param URL.
3. Handler akses variabel global `DB` (GORM) untuk query database.
4. Handler kirim response JSON.

## 2. Setup Environment (.env)

Project ini membaca konfigurasi DB dari environment variable:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_PORT=5432
DB_SSLMODE=disable
```

Contoh menjalankan (PowerShell):

```powershell
$env:DB_HOST="localhost"
$env:DB_USER="postgres"
$env:DB_PASSWORD="your_password"
$env:DB_NAME="your_database"
$env:DB_PORT="5432"
$env:DB_SSLMODE="disable"
go run .\CRUD-GORM
```

## 3. Model Data

Model utama:

```go
type Product struct {
		gorm.Model
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		Picture     string  `json:"picture"`
		Description string  `json:"description"`
}
```

`gorm.Model` otomatis memberi field:
- `ID`
- `CreatedAt`
- `UpdatedAt`
- `DeletedAt` (soft delete)

## 4. Daftar Endpoint CRUD

Base path: `/products`

1. `POST /products/` -> Create product
2. `GET /products/` -> Ambil semua product
3. `GET /products/:id` -> Ambil product by id
4. `PUT /products/:id` -> Update product by id
5. `DELETE /products/:id` -> Delete product by id

## 5. Contoh Request Response

### Create Product

Request:

```http
POST /products/
Content-Type: application/json

{
	"name": "Keyboard Mechanical",
	"price": 750000,
	"stock": 10,
	"picture": "https://example.com/keyboard.jpg",
	"description": "Keyboard hot-swappable"
}
```

Response sukses: `201 Created`

### Get All Products

Request:

```http
GET /products/
```

Response sukses: `200 OK` + array product.

### Get Product By ID

Request:

```http
GET /products/1
```

Jika tidak ada data: `404 Not Found`

### Update Product

Request:

```http
PUT /products/1
Content-Type: application/json

{
	"name": "Keyboard Mechanical V2",
	"price": 800000,
	"stock": 8,
	"picture": "https://example.com/keyboard-v2.jpg",
	"description": "Switch linear"
}
```

Response sukses: `200 OK`

### Delete Product

Request:

```http
DELETE /products/1
```

Response sukses: `200 OK`

## 6. Catatan Penting Implementasi Saat Ini

1. Di `CreateProduct`, error dari `DB.Create` belum dicek.
2. Di `GetProducts`, error dari `DB.Find` belum dicek.
3. Di `UpdateProduct`, hasil `ShouldBindJSON` belum dicek (sebaiknya divalidasi).
4. Belum ada validasi bisnis (misalnya `price >= 0`, `stock >= 0`).
5. Belum ada pagination/filter/sort untuk `GET /products/`.

## 7. Checklist Improvement Berikutnya

1. Tambah validasi input di struct pakai tag `binding`.
2. Cek semua error query database dan kirim status code yang tepat.
3. Tambah response format konsisten (misalnya selalu ada `message`, `data`, `error`).
4. Tambah pagination: `?page=1&limit=10`.
5. Tambah logging request dan error handling middleware.

## 8. Ringkasan Singkat

CRUD-GORM di folder ini sudah berjalan dengan pola standar:
- Gin untuk routing HTTP
- GORM untuk ORM PostgreSQL
- Satu model `Product`
- Lima endpoint CRUD dasar

Ini sudah cukup bagus untuk fondasi belajar, lalu bisa ditingkatkan bertahap ke validasi, pagination, dan struktur response yang lebih production-ready.
