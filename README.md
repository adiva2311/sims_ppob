# SIMS_PPOB API

## Deskripsi

SIMS_PPOB adalah proyek backend RESTful API yang dikembangkan sebagai bagian dari Home Test Nusatech. Aplikasi ini mensimulasikan sistem Payment Point Online Bank (PPOB) sederhana, di mana pengguna dapat melakukan registrasi, login, melihat profil, top up saldo, dan melakukan pembayaran layanan.
API ini dibangun menggunakan Golang dengan framework Echo, serta mengimplementasikan JWT Authentication, validasi input, dan middleware otorisasi.

---

## Teknologi yang Digunakan

- Golang (Echo Framework) — untuk membangun REST API
- MySQL — untuk penyimpanan data
- JWT (JSON Web Token) — untuk autentikasi pengguna
- Validator (go-playground/validator) — untuk validasi data request
- Railway — untuk deployment ke cloud

---

## Struktur Proyek

```
SIMS_PPOB/
├── config/               # Konfigurasi database dan environment
├── controllers/          # Handler untuk setiap endpoint
├── db/migrations/        # File migrasi database
├── dto/                  # Data Transfer Object (request/response struct)
├── img/profile_image/    # Folder penyimpanan gambar profil
├── middlewares/          # Middleware JWT dan lainnya
├── models/               # Struktur model database
├── repositories/         # Query dan interaksi dengan database
├── routes/               # Inisialisasi routing aplikasi
├── services/             # Logika bisnis utama
├── utils/                # Helper umum (misalnya untuk hashing, response)
├── .env                  # File konfigurasi environment
├── env-example           # Contoh konfigurasi environment
├── migrate.sh            # Script untuk menjalankan migrasi
├── go.mod / go.sum       # Dependency management
└── main.go               # Entry point aplikasi
```

---

## Endpoint API

### User Routes

| Method | Endpoint        | Deskripsi                   | Auth |
| ------ | --------------- | --------------------------- | ---- |
| POST   | /registration   | Registrasi user baru        | -    |
| POST   | /login          | Login dan mendapatkan JWT   | -    |
| GET    | /profile        | Mendapatkan profil pengguna | ✅   |
| PUT    | /profile/update | Update data profil          | ✅   |
| PUT    | /profile/image  | Update foto profil          | ✅   |

### Information Routes

| Method | Endpoint  | Deskripsi                       | Auth |
| ------ | --------- | ------------------------------- | ---- |
| GET    | /banner   | Menampilkan semua banner        | -    |
| GET    | /services | Menampilkan daftar layanan PPOB | ✅   |

### Transaction Routes

| Method | Endpoint             | Deskripsi                    | Auth |
| ------ | -------------------- | ---------------------------- | ---- |
| GET    | /balance             | Melihat saldo pengguna       | ✅   |
| POST   | /topup               | Melakukan top up saldo       | ✅   |
| POST   | /transaction         | Melakukan pembayaran layanan | ✅   |
| GET    | /transaction/history | Melihat riwayat transaksi    | ✅   |

---

## Cara Menjalankan Proyek Secara Lokal

1. Clone repository

   ```bash
   git clone https://github.com/adiva2311/sims_ppob.git
   cd sims_ppob
   ```

2. Buat file `.env`
   Salin dari `env-example` dan isi sesuai konfigurasi lokal, misalnya:

   ```env
    API_HOST=127.0.0.1
    API_PORT=8080

    DB_NAME=sims_ppob
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=password

    JWT_SECRET=your_secret_key
   ```

3. Jalankan migrasi database

   ```bash
   migrate -database "mysql://root:password@tcp(127.0.0.1:3306)/sims_ppob" -path db/migrations up #run all up migrations
   migrate -database "mysql://root:password@tcp(127.0.0.1:3306)/sims_ppob" -path db/migrations down #run all down migrations
   ```

4. Jalankan aplikasi

   ```bash
   go run main.go
   ```

5. Akses API
   Aplikasi akan berjalan di http://localhost:8080 (atau port sesuai konfigurasi .env)

---

## Deployment

Aplikasi ini telah dideploy ke Railway dan dapat diakses melalui:

https://simsppob-production-177e.up.railway.app/

---

## Fitur Utama

- Registrasi dan Login dengan JWT Authentication
- Update profil dan foto pengguna
- Top up saldo dan transaksi pembayaran
- Middleware untuk proteksi endpoint
- Validasi input menggunakan go-playground/validator
- Struktur kode modular dengan layer repository, service, controller
