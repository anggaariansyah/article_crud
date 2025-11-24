# Article CRUD Backend - Go + GORM + PostgreSQL

Backend ini menyediakan **CRUD Artikel** dengan upload foto menggunakan **Go**, **GORM**, dan **PostgreSQL**.  
Endpoint dilindungi **Basic Auth** dan dapat dijalankan **manual** atau **menggunakan Docker**.  
Database tetap mengikuti konfigurasi di `config.yaml`.

---

## Fitur

- CRUD Artikel (Create, Read, Update, Delete)
- Upload banyak foto untuk setiap artikel
- Update foto akan mengganti foto lama **hanya jika ada foto baru**
- Basic Auth untuk mengamankan endpoint
- Konfigurasi via `config.yaml`
- Auto migrate tabel `articles` dan `article_photos`
- Bisa dijalankan manual atau dengan Docker (DB tetap dari config)

---

## Struktur Folder

article-crud/
├── main.go
├── go.mod
├── config/
│ └── config.go
├── models/
│ ├── article.go
│ └── article_photo.go
├── repository/
│ └── article_repository.go
├── handlers/
│ └── article_handler.go
├── routes/
│ └── routes.go
├── migrations/
│ └── migrate.go
├── utils/
│ └── upload.go
├── middlewares/
│ └── auth.go
├── uploads/ <-- folder untuk menyimpan foto
├── config.yaml
├── Dockerfile
├── docker-compose.yml
└── README.md


---

## Konfigurasi `config.yaml`

```yaml
database:
  host: localhost   # koneksi ke DB lokal atau remote, tidak pakai Docker
  port: 5432
  user: postgres
  password: postgres
  name: article_db
server:
  port: 8080
auth:
  username: admin
  password: secret123
Cara Menjalankan
1️⃣ Manual (Go + PostgreSQL)

Install Go & PostgreSQL

Buat database:

CREATE DATABASE article_db;


Sesuaikan config.yaml untuk database lokal/remote

Jalankan server:

go run main.go


Server akan berjalan di http://localhost:8080

Folder uploads/ menyimpan foto artikel
