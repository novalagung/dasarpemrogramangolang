# C.2. Parsing HTTP Request Payload (Echo)

Pada bab ini kita akan belajar cara parsing beberapa variasi request payload.

Payload dalam HTTP request bisa dikirimkan dalam berbagai bentuk. Kita akan mempelajari cara untuk handle 4 jenis payload berikut.

 - Form Data
 - JSON Payload
 - XML Payload
 - Query String

Secara tidak sadar, kita sudah sering menggunakan jenis payload form data. Contohnya pada html form, ketika event submit di-trigger, data dikirim ke destinasi dalam bentuk form data. Indikatornya adalah data tersebut ditampung dalam request body, dan isi request header `Content-Type` adalah `application/x-www-form-urlencoded` (atau `multipart/form-data`).

JSON payload dan XML payload sebenarnya sama dengan Form Data, pembedanya adalah header`Content-Type` masing-masing request. Untuk JSON payload isi header tersebut adalah `application/json`, sedang untuk XML payload adalah `application/xml`.

Sedang jenis request query string adalah yang paling berbeda. Data tidak disisipkan dalam request body, melainkan pada url nya dalam bentuk key-value.

## C.2.1. Parsing Request Payload

Cara parsing payload request dalam echo sangat mudah, apapun jenis payload nya, API yang digunakan untuk parsing adalah sama.

Mari kita langsung praktekan, buat satu folder projek baru, buat `main.go`. Buat struct `User`, nantinya digunakan untuk menampung data payload yang dikirim.

```go
package main

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
)

type User struct {
    Name  string `json:"name" form:"name" query:"name"`
    Email string `json:"email" form:"email" query:"email"`
}

func main() {
    r := echo.New()

    // routes here

    fmt.Println("server started at :9000")
    r.Start(":9000")
}
```

Selanjutnya siapkan satu buah endpoint `/user` menggunakan `r.Any()`. Method `.Any()` menerima segala jenis request dengan method GET, POST, PUT, atau lainnya.

```go
r.Any("/user", func(c echo.Context) (err error) {
    u := new(User)
    if err = c.Bind(u); err != nil {
        return
    }

    return c.JSON(http.StatusOK, u)
})
```

Bisa dilihat dalam handler, method `.Bind()` milik `echo.Context` digunakan, dengan disisipi parameter pointer objek (hasil cetakan struct `User`). Parameter tersebut nantinya akan menampung payload yang dikirim, entah apapun jenis nya.

## C.2.2 Testing

Jalankan aplikasi, lakukan testing. Bisa gunakan `curl` ataupun API testing tools sejenis postman atau lainnya.

Di bawah ini shortcut untuk melakukan request menggunakan `curl` pada 4 jenis payload yang kita telah bahas. Response dari kesemua request adalah sama, menandakan bahwa data yang dikirim berhasil ditampung.

#### • Form Data

```bash
curl -X POST http://localhost:9000/user \
     -d 'name=Joe' \
     -d 'email=nope@novalagung.com'

# output => {"name":"Nope","email":"nope@novalagung.com"}
```

#### • JSON Payload

```bash
curl -X POST http://localhost:9000/user \
     -H 'Content-Type: application/json' \
     -d '{"name":"Nope","email":"nope@novalagung.com"}'

# output => {"name":"Nope","email":"nope@novalagung.com"}
```

#### • XML Payload

```bash
curl -X POST http://localhost:9000/user \
     -H 'Content-Type: application/xml' \
     -d '<?xml version="1.0"?>\
        <Data>\
            <Name>Joe</Name>\
            <Email>nope@novalagung.com</Email>\
        </Data>'

# output => {"name":"Nope","email":"nope@novalagung.com"}
```

#### • Query String

```bash
curl -X GET http://localhost:9000/user?name=Joe&email=nope@novalagung.com

# output => {"name":"Nope","email":"nope@novalagung.com"}
```

---

 - [Echo](https://github.com/labstack/echo), by Vishal Rana (Lab Stack), MIT license

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.2-parsing-http-request-payload-echo">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.2...</a>
</div>
