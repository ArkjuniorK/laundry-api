# Laundry REST API

Repository ini mengandung implementasi REST API untuk layanan laundry sederhana menggunakan Go sebagai salah test yang diajukan oleh PT. Cita Rasa Mendunia.

### Tech:

- Fiber Framework
- GORM
- SQLite
- Swagger

### How to run

- Buat sebuah database local dengan nama `laundry.db`
- Unduh semua dependencies dengan perintah `go mod tidy` atau `go mod download`
- Jalankan API menggunakan perintah `go run . api`, secara default API akan berjalan pada host 0.0.0.0 dan port 8080

### Docs

Untuk mengakses dokumentasi API bisa menggunakan endpoint `/docs`
