# Version Changelogs & Updates

## 📝 Release v4.1.20260510 (2026-05-10)

#### ◉ Perbaikan konsep & konten

- [A.16. Slice](/A-slice.html)
  - Pemindahan section `make` dari chapter array ke slice agar sesuai konsep
  - Penjelasan reference type dan operasi slice diperbaiki (termasuk typo "pinnaple" → "pineapple")
- [A.17. Map](/A-map.html)
  - Klarifikasi penggunaan `make` dan `new` pada map
- [A.31. Channel](/A-channel.html)
  - Koreksi penjelasan "pass by reference" → channel adalah reference type yang di-copy by value
- [A.47. Hash SHA1](/A-hash-sha1.html)
  - Koreksi konsep hash ("enkripsi satu arah" → "fungsi satu arah")
  - Penambahan catatan keamanan untuk penggunaan production (SHA1 tidak lagi aman, gunakan SHA256/SHA3)
- [A.54. Web Service API Server](/A-web-service-api.html)
  - Perbaikan penjelasan request method, response code, dan contoh cURL
- [A.65. Go Generics](/A-golang-generics.html)
  - Koreksi penjelasan `comparable` (bukan untuk "semua tipe data", melainkan tipe yang mendukung `==`/`!=`)
  - Rekomendasi penggunaan `any` untuk semua tipe data
- [B.23. Server Handler HTTP Request Cancellation](/B-server-handler-http-request-cancellation.html)
  - Perbaikan section number (B.32 → B.23), duplikat import dihapus
- [C.25. HTTP/2 dan HTTP/2 Server Push](/C-http2-server-push.html)
  - Penyesuaian narasi untuk browser modern yang sudah tidak mendukung server push
  - Catatan historis dan rekomendasi preload/103 Early Hints sebagai alternatif
- [C.32. JSON Web Token (JWT)](/C-golang-jwt.html)
  - Koreksi konsep signing vs encryption (JWT di-sign, bukan di-enkripsi)
  - Migrasi dari `jwt.StandardClaims` (deprecated) ke `jwt.RegisteredClaims`

#### ◉ Perbaikan bug pada kode

- Perbaikan variabel name mismatch pada kode FTP (`connt` → `conn`) agar bisa dikompilasi
- Perbaikan typo `pckage` → `package` pada kode SSO/SAML
- Perbaikan `go get` URL pada chapter CORS (menghapus `https://` prefix)
- Perbaikan error handling pada AWS S3 presign URL (dari `req, _` menjadi `req, err` dengan pengecekan)
- Perbaikan missing `return` setelah `http.Error` pada WebSocket handler (D.3)
- Perbaikan `http.Error` ganda pada singleflight handler (C.37)

#### ◉ Pembaruan dependency & safety notes

- [Redis](/C-golang-redis.html): update import path `go-redis/redis/v8` → `redis/go-redis/v9`
- [Excel](/C-read-write-excel-xlsx-file.html): update import path `360EntSecGroup-Skylar/excelize` → `xuri/excelize/v2`
- [Session MongoDB](/C-session.html): penggantian `gopkg.in/mgo.v2` → `github.com/globalsign/mgo` dengan catatan deprecation
- [gRPC](/C-golang-grpc-protobuf.html): penambahan catatan versi gRPC yang lebih baru
- [Cobra/Kingpin](/C-flag-parser.html): koreksi referensi author Cobra (spf13) dan Kingpin (alecthomas)
- [Validator](/C-http-request-payload-validation.html): update referensi v9 → v10
- Penambahan catatan keamanan pada chapter secure cookie, session, secure middleware, send email, dan insecure TLS request

#### ◉ Perbaikan typo & narasi di banyak chapter

- [A.1. Berkenalan Dengan Golang](/1-berkenalan-dengan-golang.html)
- [A.7. Hello World](/A-hello-world.html) (typo "Hello Word" → "Hello World")
- [A.8. Komentar](/A-komentar.html) (typo double "untuk untuk")
- [A.11. Konstanta](/A-konstanta.html) (tipe data `float` → `float64`)
- [A.21. Fungsi Closure](/A-fungsi-closure.html) (typo "diantaranya" → "di antaranya", variable name fix)
- [A.25. Method](/A-method.html) (section number fix, "array" → "slice")
- [A.34. Channel Range Close](/A-channel-range-close.html) (penjelasan close channel)
- [A.35. Channel Timeout](/A-channel-timeout.html) (typo function name `retreiveData` → `retrieveData`)
- [A.36. Defer & Exit](/A-defer-exit.html) (typo "ekseusi" → "eksekusi")
- [A.37. Error, Panic, & Recover](/A-error-panic-recover.html) (typo "occured" → "occurred")
- [A.53. JSON](/A-json.html) (typo "konverstri" → "konversi")
- [A.55. Client HTTP Request](/A-client-http-request-simple.html) (typo, penjelasan response body)
- [A.61. Go Vendoring](/A-go-vendoring.html) (section number duplikat diperbaiki)
- [A.6. Go Command](/A-go-command.html) (penjelasan `go get` untuk Go modern)
- [B.1. Golang Web Hello World](/B-golang-web-hello-world.html) (typo `ResponseWrite`, `seperta` → `seperti`, dll)
- [B.2. Routing](/B-routing-http-handlefunc.html) (caption gambar diperbaiki)
- [B.3. Routing Static Assets](/B-routing-static-assets.html) (typo "didefiniskan" → "didefinisikan")
- [B.8. Template Custom Functions](/B-template-custom-functions.html) (typo "Perbadaan" → "Perbedaan")
- [B.19. Middleware](/B-middleware-using-http-handler.html) (typo "pengencekan" → "pengecekan")
- [B.5. Render Partial HTML](/B-template-render-partial-html.html) (section number fix)
- [B.15. AJAX JSON Response](/B-ajax-json-response.html) (koreksi `io.Reader` → `io.Writer`)
- [C.16. Secure Middleware](/C-secure-middleware.html) (section number fix)
- [C.28. FTP](/C-golang-ftp.html) (variable name fix)
- [D.3. WebSocket Chatting App](/D-golang-web-socket-chatting-app.html) (missing return ditambahkan)

#### ◉ General update

- Penambahan GitHub Actions workflow (`release.yml`) untuk membuat release otomatis saat push tag `v*`
- Penambahan fitur darkmode pada stylesheet website
- Perbaikan tampilan font-settings dropdown (sekarang selalu terlihat)
- Pembaruan README (jumlah chapter 120 → 160+, versi e-book v4.0.20241115 → v4.0.20251111)
- Penambahan kontributor baru: Rofid (alimurrofid) dan M. Gusti Maulana Z (mgustimz)
- Sinkronisasi submodule `examples` ke revisi terbaru
- Penghapusan duplikasi entry changelog v4.0.20250422

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
