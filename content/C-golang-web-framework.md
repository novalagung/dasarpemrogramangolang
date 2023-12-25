# C.2. Go Web Framework

Salah satu kelebihan bahasa Go adalah dukungan dari komunitas. Banyak sekali library dan framework untuk bahasa ini yang kesiapannya production-ready dan gratis untuk dipergunakan.

Di Go, sama seperti bahasa pemrograman lainnya, ada banyak library dan framework yang siap pakai. Ada framework yang sifatnya sudah komplit, lengkap isinya dari ujung ke ujung, mulai dari setup project hingga testing dan build/deployment sudah ada semua tooling-nya. Ada juga framework yg scope-nya lebih spesifik (biasa disebut library), seperti lib untuk mempermudah operasi di data layer, lib untuk routing, dan lainnya.

Pada chapter ini kita tidak akan membahas satu persatu mengenai library/framework yang ada. Penulis hanya akan menuliskan yang pernah penulis pakai saja.

## C.2.1. Web Framework

Untuk opsi web framewok, ada cukup banyak pilihan. Author sendiri pernah menggunakan 3 pilihan berikut:

- [Beego](https://github.com/astaxie/beego)
- [Echo](https://github.com/labstack/echo)
- [Gin](https://github.com/gin-gonic/gin)
- custom, menggunakan kombinasi dari banyak library sesuai kebutuhan dan selera.
- atau bisa menggunakan pilihan alternatif web framework lainnya.
- atau bisa juga menggunakan salah satu web framework dan di-combine dengan library lain.

Untuk opsi custom framework sendiri, pembaca bisa menggunakan kombinasi dari beberapa library berikut:

## C.2.2. Routing Library

Untuk opsi router, ada cukup banyak pilihan yg tersedia, sebagian di antaranya:

- [Chi](https://github.com/go-chi/chi)
- [FastHttp](https://github.com/valyala/fasthttp) atau [FastHttpRouter](https://github.com/buaazp/fasthttprouter)
- [Gorilla Mux](https://github.com/gorilla/mux)
- dan lainnya

## C.2.3. HTTP Middlewares

Untuk middlewares biasanya include sebagai dependensi router library. Tapi ada juga middleware independen. Contohnya:

- [CORS](https://github.com/rs/cors)
- [JWT](https://github.com/golang-jwt/jwt)
- [Rate Limiter](https://github.com/ulule/limiter)
- [Secure](https://github.com/unrolled/secure)
- dan lainnya

## C.2.4. Form & Validator

Validator library berfungsi untuk mempermudah parsing payload dan parameter dari objek http request. Rekomendasi validator library salah satunya:

- [Validator by go-playground](https://github.com/go-playground/validator/tree/v9)
- dan lainnya

## C.2.5. Database / ORM

ORM adalah salah satu pattern yg cukup sering dipakai di data layer. Beberapa library yang tersedia di antaranya:

- [Gorm](https://github.com/jinzhu/gorm)
- [Gorp](https://github.com/go-gorp/gorp)
- dan lainnya

---

Silakan mencoba-coba dan memilih kombinasi library yang cocok sesuai kebutuhan dan keinginan kawan-kawan. Semua opsi ada kelebihan dan kekurangannya. Begitu juga pada implementasi library, akan sangat mempengaruhi hasil.

Saya sangat menganjurkan pembaca untuk mencoba banyak library dan framework.

Ok, saya rasa cukup untuk pembahasan kali ini. Semoga bermanfaat

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
