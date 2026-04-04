/*
================== JWT AUTHENTICATION MOVED ==================

File JWT yang awalnya 373 lines sudah dipindahkan ke folder "5. AUTH"

Struktur terbaru:
📁 5. AUTH/
  ├── jwt.go        → Penjelasan JWT + fungsi token
  ├── models.go     → Struct User, Request, Claims
  ├── password.go   → Hash & Verify password
  ├── handlers.go   → Register, Login, Logout
  ├── middleware.go → AuthMiddleware & GetProfile
  └── README.go     → Panduan penggunaan & contoh

ALASAN SPLIT:
✓ Setiap file punya tanggung jawab tunggal
✓ Lebih mudah di-maintain dan di-update
✓ Struktur code lebih rapi
✓ Lebih mudah untuk testing

UNTUK MENGGUNAKAN:
1. Pastikan package di 5. AUTH/ sudah berada di satu folder dengan main.go
   atau import sesuai path Anda
2. Pastikan database sudah ter-connect
3. Panggil InitUserTable() di main()
4. Register routes sesuai contoh di 5. AUTH/README.go

QUICK START:
- Lihat file: 5. AUTH/README.go untuk contoh lengkap penggunaan
- Lihat file: 5. AUTH/jwt.go untuk penjelasan konsep JWT
- Lihat file: 5. AUTH/handlers.go untuk logic Register/Login/Logout

Terima kasih sudah mengikuti perjalanan belajar JWT! 🚀
*/

