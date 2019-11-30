# C.14. Secure Middleware

Pada bab ini kita akan belajar menggunakan library [secure](https://github.com/unrolled/secure) untuk meningkatkan keamanan aplikasi web.

## C.14.1. Keamanan Web Server

Jika berbicara mengenai keamanan aplikasi web, sangat luas sebenarnya cakupannya, ada banyak hal yang perlu diperhatian dan disiapkan. Mungkin tiga diantaranya sudah kita pelajari sebelumnya, yaitu penerapan Secure Cookie, CORS, dan CSRF.

Selain 3 topik tersebut masih terdapat banyak lagi. Beruntungnya ada library [secure](https://github.com/unrolled/secure). Sesuai tagline-nya, secure library digunakan untuk membantu mengatasi beberapa masalah keamanan aplikasi.

Secure library merupakan middleware, penggunaannya sama seperti middleware pada umumnya.

## C.15.2. Praktek

Mari langsung kita praktekan. Buat folder projek baru. Di file main tulis kode berikut. Sebuah aplikasi dibuat, isinya satu buah rute `/index` yang bisa diakses dari mana saja.

```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/index", func(c echo.Context) error {
        c.Response().Header().Set("Access-Control-Allow-Origin", "*")

        return c.String(http.StatusOK, "Hello")
    })

    e.Logger.Fatal(e.StartTLS(":9000", "server.crt", "server.key"))
}
```

Perlu diketahui, aplikasi di atas di-start dengan SSL/TLS enabled. Dua buah file dibutuhkan, yaitu file certificate `server.crt` dan file private key `server.key`. Silakan unduh kedua file tersebut dari source code di
[github, bab B14](https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.14). Pada bab [B.22 HTTPS/TLS Web Server](/B-22-https-tls.html) nantinya akan kita pelajari lebih lanjut mengenai cara generate kedua file di atas hingga cara penggunannya.

Kembali ke pembahasan, sekarang tambahkan secure middleware. Import package-nya, buat instance middleware, lalu registrasikan ke echo.

```go
import (
    // ...
    "github.com/unrolled/secure"
)

func main() {
    // ...

    secureMiddleware := secure.New(secure.Options{
        AllowedHosts:            []string{"localhost:9000", "www.google.com"},
        FrameDeny:               true,
        CustomFrameOptionsValue: "SAMEORIGIN",
        ContentTypeNosniff:      true,
        BrowserXssFilter:        true,
    })

    e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

    // ...
}
```

Pembuatan objek secure middleware dilakukan menggunakan `secure.New()` dengan isi parameter adalah konfigurasi. Bisa dilihat ada 5 buah property konfigurasi di-set. Berikut merupakan penjelasan tiap-tiap property tersebut.

#### • Konfigurasi `AllowedHosts`

```go
AllowedHosts: []string{"localhost:9000", "www.google.com"}
```

Host yang diperbolehkan mengakses web server ditentukan hanya 2, yaitu localhost:9000 yang merupakan web server itu sendiri, dan google.com. Silakan coba mengakses aplikasi kita ini menggunakan AJAX lewat google.com dan domainnya lainnya untuk mengetes apakah fungsionalitas nya berjalan.

#### • Konfigurasi `FrameDeny`

```go
FrameDeny: true
```

Secara default sebuah aplikasi web adalah bisa di-load di dalam iframe yang berada host nya berbeda. Misalnya di salah satu laman web www.kalipare.com ada iframe yang atribut src nya berisi www.novalagung.com, hal seperti ini diperbolehkan.

Perijinan apakah website boleh di-load lewat iframe atau tidak, dikontrol lewat header [X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options).

Di library secure, untuk men-disable ijin akses aplikasi dari dalam iframe, bisa dilakukan cukup dengan mengeset proerty `FrameDeny` dengan nilai `true`.

Untuk mengetes, silakan buat aplikasi web terpisah yang mer-render sebuah view. Dalam view tersebut siapkan satu buah iframe yang mengarah ke `https://localhost:9000/index`.

#### • Konfigurasi `CustomFrameOptionsValue`

```go
CustomFrameOptionsValue: "SAMEORIGIN"
```

Jika `FrameDeny` di-set sebagai `true`, maka semua host (termasuk aplikasi itu sendiri) tidak akan bisa me-load url lewat iframe. 

Dengan menambahkan satu buah property lagi yaitu `CustomFrameOptionsValue: "SAMEORIGIN"` maka ijin pengaksesan url lewat iframe menjadi eksklusif hanya untuk aplikasi sendiri. 

Untuk mengetes, buat rute baru yang me-render sebuah view. Dalam view tersebut siapkan satu buah iframe yang mengarah ke `/index`.

#### • Konfigurasi `ContentTypeNosniff`

```go
ContentTypeNosniff: true
```

Property `ContentTypeNosniff: true` digunakan untuk disable MIME-sniffing yang dilakukan oleh browser IE. Lebih jelasnya silakan baca [X-Content-Type-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options).

#### • Konfigurasi `BrowserXssFilter`

```go
BrowserXssFilter: true
```

Property di atas digunakan untuk mengaktifkan header [X-XSS-Protection](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection), dengan isi header adalah `1; mode=block`.

## C.15.3. Property Library Secure

Selain 5 property yang kita telah pelajari di atas, masih ada banyak lagi konfigurasi yang bisa digunakan.

 - AllowedHosts
 - HostsProxyHeaders
 - SSLRedirect
 - SSLTemporaryRedirect
 - SSLHost
 - SSLHostFunc
 - SSLProxyHeaders
 - STSSeconds
 - STSIncludeSubdomains
 - STSPreload
 - ForceSTSHeader
 - FrameDeny
 - CustomFrameOptionsValue
 - ContentTypeNosniff
 - BrowserXssFilter
 - CustomBrowserXssValue
 - ContentSecurityPolicy
 - PublicKey
 - ReferrerPolicy

Lebih mendetailnya silakan langsung cek halaman official library secure di https://github.com/unrolled/secure.

---

 - [Secure](https://github.com/unrolled/secure), by Cory Jacobsen, MIT license
 - [Echo](https://github.com/labstack/echo), by Vishal Rana (Lab Stack), MIT license

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.14-secure-middleware">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.14...</a>
</div>
