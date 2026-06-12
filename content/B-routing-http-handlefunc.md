# B.2. Routing `http.HandleFunc`

Routing di Go bisa dilakukan dengan beberapa cara, di antaranya:

 1. Dengan memanfaatkan fungsi `http.HandleFunc()`
 2. Mengimplementasikan interface `http.Handler` pada suatu struct, untuk kemudian digunakan pada fungsi `http.Handle()`
 3. Membuat multiplexer sendiri dengan memanfaatkan struct `http.ServeMux`

Pada buku ini, semua cara tersebut akan dibahas, namun khusus di chapter ini hanya `http.HandleFunc()` yang kita pelajari.

> Metode routing cara pertama dan cara kedua memiliki kesamaan yaitu sama-sama menggunakan `DefaultServeMux` sebagai router. Mengenai apa itu `DefaultServeMux` akan kita bahas lebih mendetail pada chapter lain.

## B.2.1. Penggunaan `http.HandleFunc()`

Seperti yang sudah dijelaskan sekilas pada chapter sebelumnya, fungsi `http.HandleFunc()` digunakan untuk registrasi rute/endpoint beserta handler-nya. Penggunaan fungsi ini cukup mudah, panggil saja fungsi lalu isi dua parameternya.

 1. Parameter ke-1, adalah rute (atau endpoint). Sebagai contoh: `/`, `/index`, `/about`.
 2. Parameter ke-2, berisikan handler untuk rute bersangkutan. Sebagai contoh handler untuk rute `/` bertugas untuk menampilkan output berupa html `<p>hello</p>`.

Agar lebih mudah dipahami mari langsung praktik. Siapkan file `main.go` dengan package adalah `main`, dan import package `net/http` di dalamnya.

```go
package main

import (
    "log"
    "net/http"
)
```

Buat fungsi `main()`, di dalamnya siapkan sebuah closure `handlerIndex`, lalu gunakan closure tersebut sebagai handler dari dua rute baru yang sebentar lagi disiapkan, yaitu rute `/` dan `/index`.

```go
func main() {
    handlerIndex := func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello"))
    }

    http.HandleFunc("/", handlerIndex)
    http.HandleFunc("/index", handlerIndex)
}
```

Selanjutnya, masih dalam fungsi `main()`, tambahkan rute `/data` dengan handler adalah anonymous function.

```go
func main() {
    // ...

    http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello again"))
    })
}
```

Terakhir, jalankan web server.

```go
func main() {
    // ...

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

## B.2.2. Run & Test

Tes dan lihat hasilnya.

![Rute `/data` mengembalikan data](images/B_routing_http_handlefunc_1_routing.png)

Handler bisa berupa fungsi, closure, ataupun anonymous function, intinya bebas, yang terpenting adalah skema fungsi-nya harus sesuai dengan `func (http.ResponseWriter, *http.Request)`.

## B.2.3. Enhanced Routing Pattern (Go 1.22+)

Sejak Go 1.22, pola routing pada `http.HandleFunc()` mendukung dua fitur baru yang sangat berguna: penentuan HTTP method langsung di pola rute, dan wildcard pada segmen path.

#### ◉ Method-Based Routing

Sebelumnya, satu handler bisa dipanggil dari method GET, POST, DELETE, maupun method HTTP lainnya tanpa bisa dibedakan langsung di level rute. Sekarang, kita bisa menentukannya langsung di parameter pertama:

```go
http.HandleFunc("GET /articles", handlerListArticles)
http.HandleFunc("POST /articles", handlerCreateArticle)
```

Rute `GET /articles` hanya akan cocok dengan request GET, sedangkan `POST /articles` hanya cocok dengan POST. Request dengan method lain akan mendapat response `405 Method Not Allowed` secara otomatis.

#### ◉ Wildcard Path

Segmen path bisa dijadikan wildcard menggunakan kurung kurawal `{nama}`. Nilai wildcard kemudian diambil lewat `r.PathValue("nama")`:

```go
http.HandleFunc("GET /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    w.Write([]byte("article id: " + id))
})
```

Akses rute `/articles/42`, maka `id` bernilai `"42"`.

Untuk menangkap sisa path (termasuk `/`), gunakan `{nama...}`:

```go
http.HandleFunc("/static/{path...}", handlerStatic)
```

Berikut contoh lengkap penggunaan kedua fitur ini:

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("GET /articles", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("list semua artikel"))
    })

    http.HandleFunc("POST /articles", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("buat artikel baru"))
    })

    http.HandleFunc("GET /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")
        w.Write([]byte("article id: " + id))
    })

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

> Fitur ini hanya tersedia mulai Go 1.22. Pada Go versi sebelumnya, pola routing seperti `"GET /articles/{id}"` tidak dikenali dan perlu menggunakan library routing pihak ketiga (seperti gorilla/mux atau chi).

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.2-routing-http-handlefunc">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.2...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
