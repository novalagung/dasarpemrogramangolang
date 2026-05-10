# Version Changelogs & Updates

## 📝 Release v4.1.20260510 (2026-05-10)

#### ◉ Chapter update

- Perbaikan narasi, typo, dan kejelasan penjelasan di banyak chapter A, B, dan C
- [A.16. Slice](/A-slice.html)
  - Penataan ulang urutan pembahasan agar `make` muncul setelah konsep inti slice
- [A.17. Map](/A-map.html)
  - Klarifikasi penggunaan `make` dan `new` pada map
- [A.6. Go Command](/A-go-command.html)
  - Penjelasan `go get` diperbarui agar sesuai perilaku Go modern, termasuk dampak saat dijalankan di luar module
- [A.47. Hash SHA1](/A-hash-sha1.html)
  - Koreksi konsep hash vs enkripsi dan penambahan catatan keamanan untuk penggunaan production
- [A.54. Web Service API Server](/A-web-service-api.html)
  - Perbaikan penjelasan request method, response code, dan contoh cURL
- [A.65. Go Generics](/A-golang-generics.html)
  - Penjelasan `comparable` dan generic type constraint dibuat lebih mudah dipahami
- [C.25. HTTP/2 dan HTTP/2 Server Push](/C-http2-server-push.html)
  - Penyesuaian narasi agar selaras dengan kondisi browser modern yang sudah tidak mendukung server push
- [C.32. JSON Web Token (JWT)](/C-golang-jwt.html)
  - Koreksi konsep signing vs encryption dan sinkronisasi ke `RegisteredClaims`
- Pembaruan dependency, safety note, dan referensi pada chapter pilihan lain seperti session, redis, excel, send mail, secure cookie, secure middleware, dan SSO/SAML

#### ◉ General update

- Penambahan GitHub Actions workflow untuk membuat release otomatis saat tag `v*` dipush
- Sinkronisasi submodule `examples` ke revisi terbaru
- Pembaruan README dan SUMMARY untuk mencerminkan jumlah chapter terbaru
- Perbaikan kecil pada tampilan website, termasuk font-settings dropdown dan sidebar/mobile view
- Penambahan nama kontributor baru pada halaman kontribusi

## 📝 Release v4.0.20251111 (2025-11-11)

#### ◉ Chapter update

- [A.31. Channel](/A-channel.html)
    - Perbaikan inkonsistensi gambar dan contoh
- [A.19. Fungsi Multiple Return](/A-fungsi-multiple-return.html)
    - Perbaikan typo
- [B.5. Template: Render Partial HTML Template](/B-template-render-partial-html.html)
    - Refactoring kode
- [C.36. Redis](/C-golang-redis.html)
    - Perbaikan syntax error
- [C.25. HTTP/2 dan HTTP/2 Server Push](/C-http2-server-push.html)
    - Perbaikan link chapter relevan

#### ◉ General update

- Sidebar improvement untuk mobile app view

## 📝 Release v4.0.20250422 (2025-04-22)

#### ◉ Chapter update

- [A.35. Channel - Timeout](/A-channel-timeout.html)
    - Peningkatan konten & perbaikan typo
- [C.35. Dockerize Aplikasi Golang](/C-dockerize-golang.html)
    - Peningkatan konten & perbaikan typo

## 📝 Release v4.0.20250422 (2025-04-22)

#### ◉ Chapter update

- [A.35. Channel - Timeout](/A-channel-timeout.html)
    - Peningkatan konten & perbaikan typo
- [C.35. Dockerize Aplikasi Golang](/C-dockerize-golang.html)
    - Peningkatan konten & perbaikan typo

## 📝 Release v4.0.20241115 (2024-11-15)

#### ◉ General update

- UI updates

## 📝 Release v4.0.20240830 (2024-08-30)

#### ◉ Chapter update

- [A.2. Instalasi Golang (Stable & Unstable)](/2-instalasi-golang.html)
    - Update command instalasi
- [A.3. Go Modules](/A-setup-go-project-dengan-go-modules.html)
    - Peningkatan konten & perbaikan typo
- [A.10. Tipe Data](A-tipe-data.html)
    - Peningkatan konten & perbaikan typo
- [A.14. Perulangan](/A-perulangan.html)
    - Penambahan penjelasan tentang `for i := range N`
- [A.18. Fungsi](/A-fungsi.html)
    - Peningkatan konten & perbaikan typo
- [A.32. Buffered Channel](/A-buffered-channel.html)
    - Peningkatan konten & perbaikan typo
- [A.42. Time Duration](/A-time-duration.html)
    - Perbaikan kesalahan penjelasan pada `time.Duration`
- [C.34. SSO SAML 2.0 (Service Provider)](/C-golang-sso-saml-sp.html)
    - Peningkatan gambar
- Perbaikan narasi konten semua chapter di section A
- Perbaikan narasi konten semua chapter di section B

#### ◉ General update

- Penerapan manual versioning
- Penambahan halaman changelogs
- Penambahan halaman download file
- Improvisasi keyword untuk keperluan SEO
- Penyesuaian resolusi gambar konten
