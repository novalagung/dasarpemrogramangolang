# C.3. Echo Framework & Routing

Pada chapter ini kita akan belajar cara mudah routing menggunakan [Echo Framework](https://echo.labstack.com/). 

Mulai chapter **C1** hingga **C6** kita akan mempelajari banyak aspek dalam framework Echo dan mengkombinasikannya dengan beberapa library lain.

# C.3.1 Echo Framework

Echo adalah framework bahasa golang untuk pengembangan aplikasi web. Framework ini cukup terkenal di komunitas. Echo merupakan framework besar, di dalamnya terdapat banyak sekali dependensi.

Salah satu dependensi yang ada di dalamnya adalah router, dan pada chapter ini kita akan mempelajarinya.

Dari banyak routing library yang sudah penulis gunakan, hampir seluruhnya mempunyai kemiripan dalam hal penggunaannya, cukup panggil fungsi/method yang dipilih (biasanya namanya sama dengan HTTP Method), lalu sisipkan rute pada parameter pertama dan handler pada parameter kedua.

Berikut contoh sederhana penggunaan echo framework.

```go
r := echo.New()
r.GET("/", handler)
r.Start(":9000")
```

Sebuah objek router `r` dicetak lewat `echo.New()`. Lalu lewat objek router tersebut, dilakukan registrasi rute untuk `/` dengan method GET dan handler adalah closure `handler`. Terakhir, dari objek router di-start-lah sebuah web server pada port 9000.

> Echo router mengadopsi konsep [radix tree](https://en.wikipedia.org/wiki/Radix_tree), membuat performa lookup nya begitu cepat. Tak juga itu, pemanfaatan sync pool membuat penggunaan memory lebih hemat, dan aman dari GC overhead.

## C.3.2. Praktek

Mari kita pelajari lebih lanjut dengan praktek langsung. Buat folder proyek baru, buat `main.go`, isi dengan kode berikut, kemudian jalankan aplikasi.

```go
package main

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
    "strings"
)

type M map[string]interface{}

func main() {
    r := echo.New()

    r.GET("/", func(ctx echo.Context) error {
        data := "Hello from /index"
        return ctx.String(http.StatusOK, data)
    })

    r.Start(":9000")
}
```

Kode di atas adalah contoh sederhana penerapan echo router.

![Preview](images/C_echo_routing_1_routing_slash_test.png)

Routing dengan memanfaatkan package `net/http` dalam penerapannya adalah menggunakan `http.HandleFunc()` atau `http.Handle()`. Berbeda dengan Echo, routingnya adalah method-based, tidak hanya endpoint dan handler yang di-registrasi, method juga.

Statement `echo.New()` mengembalikan objek mux/router. Pada kode di atas rute `/` dengan method `GET` di-daftarkan. Selain `r.GET()` ada banyak lagi method lainnya, semua method dalam [spesifikasi REST](https://en.wikipedia.org/wiki/Representational_state_transfer) seperti PUT, POST, dan lainnya bisa digunakan.

Handler dari method routing milik echo membutuhkan satu argument saja, dengan tipe adalah `echo.Context`. Dari argumen tersebut objek `http.ResponseWriter` dan `http.Request` bisa di-akses. Namun kedua objek tersebut akan jarang kita gunakan karena `echo.Context` memiliki banyak method yang beberapa tugasnya sudah meng-cover operasi umum yang biasanya kita lakukan lewat objek request dan response, di antara seperti: 

 - Render output (dalam bentuk html, plain text, json, atau lainnya).
 - Parsing request data (json payload, form data, query string).
 - URL Redirection.
 - ... dan lainnya.

> Untuk mengakses objek `http.Request` gunakan `ctx.Request()`.<br />Sedang untuk objek `http.ResponseWriter` gunakan `ctx.Response()`.

Salah satu alasan lain kenapa penulis memilih framework ini, adalah karena desain route-handler-nya menarik. Dalam handler cukup kembalikan objek error ketika memang ada kesalahan terjadi, sedangkan jika tidak ada error maka kembalikan nilai `nil`.

Ketika terjadi error pada saat mengakses endpoint, idealnya [HTTP Status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) error dikembalikan sesuai dengan jenis errornya. Tapi terkadang juga ada kebutuhan dalam kondisi tertentu `http.StatusOK` atau status 200 dikembalikan dengan disisipi informasi error dalam response body-nya. Kasus sejenis ini menjadikan standar error reporting menjadi kurang bagus. Pada konteks ini echo unggul menurut penulis, karena default-nya semua error dikembalikan sebagai response dalam bentuk yang sama.

Method `ctx.String()` dari objek context milik handler digunakan untuk mempermudah rendering data string sebagai output. Method ini mengembalikan objek error, jadi bisa digunakan langsung sebagai nilai balik handler. Argumen pertama adalah http status dan argumen ke-2 adalah data yang dijadikan output.

## C.3.3. Response Method milik `ctx`

Selain `ctx.String()` ada banyak method sejenis lainnya, berikut selengkapnya.

#### • Method `.String()`

Digunakan untuk render plain text sebagai output (isi response header `Content-Type` adalah `text/plain`). Method ini tugasnya sama dengan method `.Write()` milik objek `http.ResponseWriter`.

```go
r.GET("/index", func(ctx echo.Context) error {
    data := "Hello from /index"
    return ctx.String(http.StatusOK, data)
})
```

#### • Method `.HTML()`

Digunakan untuk render html sebagai output. Isi response header `Content-Type` adalah `text/html`.

```go
r.GET("/html", func(ctx echo.Context) error {
    data := "Hello from /html"
    return ctx.HTML(http.StatusOK, data)
})
```

#### • Method `.Redirect()`

Digunakan untuk redirect, pengganti `http.Redirect()`.

```go
r.GET("/index", func(ctx echo.Context) error {
    return ctx.Redirect(http.StatusTemporaryRedirect, "/")
})
```

#### • Method `.JSON()`

Digunakan untuk render data JSON sebagai output. Isi response header `Content-Type` adalah `application/json`.

```go
r.GET("/json", func(ctx echo.Context) error {
    data := M{"Message": "Hello", "Counter": 2}
    return ctx.JSON(http.StatusOK, data)
})
```

## C.3.4. Parsing Request

Echo juga menyediakan beberapa method untuk keperluan parsing request, di antaranya:

#### • Parsing Query String

Method `.QueryParam()` digunakan untuk mengambil data pada query string request, sesuai dengan key yang diinginkan.

```go
r.GET("/page1", func(ctx echo.Context) error {
    name := ctx.QueryParam("name")
    data := fmt.Sprintf("Hello %s", name)

    return ctx.String(http.StatusOK, data)
})
```

Test menggunakan curl:

```bash
curl -X GET http://localhost:9000/page1?name=grayson
```

#### • Parsing URL Path Param

Method `.Param()` digunakan untuk mengambil data path parameter sesuai skema rute.

```go
r.GET("/page2/:name", func(ctx echo.Context) error {
    name := ctx.Param("name")
    data := fmt.Sprintf("Hello %s", name)

    return ctx.String(http.StatusOK, data)
})
```

Bisa dilihat, terdapat `:name` pada pendeklarasian rute. Nantinya url apapun yang ditulis sesuai skema di-atas akan bisa diambil path parameter-nya. Misalkan `/page2/halo` maka `ctx.Param("name")` mengembalikan string `halo`.

Test menggunakan curl:

```bash
curl -X GET http://localhost:9000/page2/grayson
```

#### • Parsing URL Path Param dan Setelahnya

Selain mengambil parameter sesuai spesifik path, kita juga bisa mengambil data **parameter path dan setelahnya**.

```go
r.GET("/page3/:name/*", func(ctx echo.Context) error {
    name := ctx.Param("name")
    message := ctx.Param("*")

    data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

    return ctx.String(http.StatusOK, data)
})
```

Statement `ctx.Param("*")` mengembalikan semua path sesuai dengan skema url-nya. Misal url adalah `/page3/tim/a/b/c/d/e/f/g/h` maka yang dikembalikan adalah `a/b/c/d/e/f/g/h`.

Test menggunakan curl:

```bash
curl -X GET http://localhost:9000/page3/tim/need/some/sleep
```

#### • Parsing Form Data

Data yang dikirim sebagai request body dengan jenis adalah Form Data bisa di-ambil dengan mudah menggunakan `ctx.FormValue()`.

```go
r.POST("/page4", func(ctx echo.Context) error {
    name := ctx.FormValue("name")
    message := ctx.FormValue("message")

    data := fmt.Sprintf(
        "Hello %s, I have message for you: %s",
        name,
        strings.Replace(message, "/", "", 1),
    )

    return ctx.String(http.StatusOK, data)
})
```

Test menggunakan curl:

```bash
curl -X POST -F name=damian -F message=angry http://localhost:9000/page4
```

Pada chapter selanjutnya kita akan belajar teknik parsing request data yang lebih advance.

## C.3.5. Penggunaan `echo.WrapHandler` Untuk Routing Handler Bertipe `func(http.ResponseWriter,*http.Request)` atau `http.HandlerFunc`

Echo bisa dikombinasikan dengan handler ber-skema *NON-echo-handler* seperti `func(http.ResponseWriter,*http.Request)` atau `http.HandlerFunc`.

Caranya dengan memanfaatkan fungsi `echo.WrapHandler` untuk mengkonversi handler tersebut menjadi echo-compatible. Lebih jelasnya silakan lihat kode berikut.

```go
var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("from action home"))
    },
)

var ActionAbout = echo.WrapHandler(
    http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("from action about"))
        },
    ),
)

func main() {
    r := echo.New()

    r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
    r.GET("/home", echo.WrapHandler(ActionHome))
    r.GET("/about", ActionAbout)

    r.Start(":9000")
}
```

Untuk routing handler dengan skema `func(http.ResponseWriter,*http.Request)`, maka harus dibungkus dua kali, pertama menggunakan `http.HandlerFunc`, lalu dengan `echo.WrapHandler`.

Sedangkan untuk handler yang sudah bertipe `http.HandlerFunc`, bungkus langsung menggunakan `echo.WrapHandler`.

## C.3.6. Routing Static Assets

Cara routing static assets di echo sangatlah mudah. Gunakan method `.Static()`, isi parameter pertama dengan prefix rute yang di-inginkan, dan parameter ke-2 dengan path folder tujuan.

Buat sub folder dengan nama `assets` dalam folder projek. Dalam folder tersebut buat sebuah file `layout.js`, isinya bebas.

Pada `main.go` tambahkan routing static yang mengarah ke path folder `assets`.

```go
r.Static("/static", "assets")
```

Jalankan aplikasi, lalu coba akses `http://localhost:9000/static/layout.js`.

![Routing static assets](images/C_echo_routing_2_routing_static_assets.png)

---

 - [Echo](https://github.com/labstack/echo), by Vishal Rana (Lab Stack), MIT license

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.3-echo-routing">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.3...</a>
</div>
