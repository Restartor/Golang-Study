# Golang-Study

Repository ini berisi perjalanan belajar backend dengan Go, dari dasar sampai mini project API dengan Gin + GORM.

## Roadmap Belajar

1. Basic
- Input output
- Tipe data
- Operator
- If else
- Loop
- Function
- Slice
- Struct

2. Another Basic
- Pointer
- Interface
- Goroutine dan channel (dasar)

3. Web Server
- Pakai net/http
- REST API sederhana

4. Framework
- Belajar Gin

5. Feature dan Database
- Koneksi PostgreSQL
- JWT authentication (register, login, logout)
- CRUD dengan GORM

6. Next Step (opsional)
- Docker
- Deploy ke VPS atau cloud

## Struktur Folder Saat Ini

- 1. Basic
Contoh konsep dasar Go.

- 1.5 Another Basic
Materi lanjutan dasar seperti pointer, interface, dan concurrency.

- 2. Web Server
Contoh API menggunakan package net/http.

- 3. Gin Framework
Contoh server menggunakan Gin.

- FITUR
Eksperimen fitur seperti koneksi DB, route, dan JWT versi awal.

- AUTH
Implementasi auth yang dipisah per file agar lebih rapi.

- CRUD-GORM
Contoh CRUD Product menggunakan Gin + GORM + PostgreSQL.

## Prasyarat

- Go terpasang (cek dengan: go version)
- PostgreSQL aktif (untuk folder yang butuh database)
- Environment variable database diset:
- DB_HOST
- DB_USER
- DB_PASSWORD
- DB_NAME
- DB_PORT
- DB_SSLMODE

## Cara Menjalankan

Karena setiap folder adalah bahan belajar mandiri, jalankan sesuai folder yang ingin dipelajari.

Contoh menjalankan CRUD-GORM:

```powershell
cd CRUD-GORM
go run .
```

Contoh menjalankan Gin dasar:

```powershell
cd "3. Gin Framework"
go run Gin.go
```

## Catatan

- Beberapa file dibuat independen untuk tujuan belajar, jadi tidak semua folder harus saling terhubung.
- Fokus utama repository ini adalah pembelajaran bertahap dengan contoh yang mudah dipahami.

Selamat belajar Go backend.
