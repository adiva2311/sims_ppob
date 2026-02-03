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