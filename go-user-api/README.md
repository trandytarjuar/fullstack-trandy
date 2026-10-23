# Go User API

REST API sederhana untuk mengelola data user. Aplikasi dibangun menggunakan Go, Gin, GORM, dan SQLite, serta menyediakan dokumentasi interaktif melalui Swagger UI.

## Teknologi

- Go 1.23
- Gin Web Framework
- GORM
- SQLite dengan driver pure Go (tidak memerlukan GCC/CGO)
- Swagger/OpenAPI

## Menjalankan aplikasi

Pastikan Go 1.23 atau versi yang lebih baru sudah terpasang, lalu jalankan:

```powershell
cd go-user-api
go mod download
go run .
```

Server akan berjalan di:

```text
http://localhost:8081
```

File database SQLite `users.db` akan dibuat otomatis di direktori kerja apabila belum tersedia.

Untuk menggunakan lokasi database lain, atur environment variable `DB_PATH`:

```powershell
$env:DB_PATH = "C:\data\users.db"
go run .
```

## Swagger UI

Setelah server berjalan, buka:

```text
http://localhost:8081/swagger/index.html
```

Pilih endpoint, klik **Try it out**, isi parameter atau request body, lalu klik **Execute**.

## Endpoint API

| Method | Endpoint | Keterangan |
| --- | --- | --- |
| `GET` | `/api/users` | Mengambil semua user |
| `GET` | `/api/users/{userId}` | Mengambil user berdasarkan ID |
| `POST` | `/api/users` | Membuat user baru |
| `PUT` | `/api/users/{userId}` | Memperbarui user |
| `DELETE` | `/api/users/{userId}` | Menghapus user |

Contoh request body untuk `POST` dan `PUT`:

```json
{
  "firstName": "Budi",
  "lastName": "Santoso",
  "email": "budi@example.com"
}
```

Semua field wajib diisi. Alamat email harus valid dan unik.

## Status respons

| Status | Keterangan |
| --- | --- |
| `200 OK` | Permintaan berhasil |
| `201 Created` | User berhasil dibuat |
| `204 No Content` | User berhasil dihapus |
| `400 Bad Request` | Data request tidak valid |
| `404 Not Found` | User tidak ditemukan |
| `409 Conflict` | Email sudah digunakan |
| `500 Internal Server Error` | Terjadi kesalahan pada server |

## Memperbarui dokumentasi Swagger

Jika anotasi Swagger pada kode diubah, hasilkan ulang folder `docs` dengan:

```powershell
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init
```

## Pengujian

Jalankan pemeriksaan build dan test dengan:

```powershell
go test ./...
```

## Struktur proyek

```text
go-user-api/
|-- docs/             # Dokumentasi Swagger yang dihasilkan
|-- database.go       # Koneksi dan migrasi database
|-- main.go           # Entry point aplikasi
|-- model.go          # Model dan request DTO
|-- router.go         # Route API dan Swagger
|-- user_handler.go   # Handler CRUD user
|-- go.mod
`-- go.sum
```
