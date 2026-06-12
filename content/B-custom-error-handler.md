# B.26. Custom Error Handler

Secara default, ketika ada API call ke sebuah endpoint web server Go sedangkan endpoint tersebut tidak ada handlernya, maka response teks polos `404 page not found` dikembalikan.

Pada chapter ini kita belajar cara menampilkan halaman error yang lebih informatif dan bagaimana menangani error yang tidak terduga seperti panic.

## B.26.1. Custom Error Handler

`http.ServeMux` tidak menyediakan cara langsung untuk mengganti tampilan 404 default-nya. Satu-satunya cara untuk membuat custom error handler menggunakan stdlib mux milik Go adalah memanfaatkan sifat route `/` yang berfungsi sebagai catch-all: setiap request yang tidak cocok dengan rute manapun akan diteruskan ke handler `/`. Pendekatan ini paling tepat digunakan ketika aplikasi dibangun murni menggunakan stdlib tanpa framework tambahan.

```go
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        notFoundHandler(w, r)
        return
    }
    w.Write([]byte("Home"))
})
```

Karena route `/` juga menangani halaman home, `r.URL.Path` perlu diperiksa secara manual untuk membedakan keduanya. Jika path bukan `/`, fungsi `notFoundHandler()` dipanggil untuk mengirim status 404 dengan pesan custom.

```go
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 - Halaman tidak ditemukan"))
}

func internalErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Terjadi kesalahan pada server"))
}
```

Fungsi `internalErrorHandler()` dengan pola yang sama disiapkan untuk response 500, dan akan digunakan oleh panic recovery middleware.

## B.26.2. Panic Recovery Middleware

Ketika sebuah handler mengalami `panic` dan tidak ada yang menangkapnya, Go akan menghentikan goroutine tersebut dan seluruh server ikut crash. Solusi masalah tersebut adalah dengan membuat middleware untuk menangkap panic berikut.

```go
func panicMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic: %v", err)
                internalErrorHandler(w, r)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

Fungsi `recover()` hanya efektif jika dipanggil langsung di dalam fungsi yang di-`defer`. Itulah mengapa di kode di atas fungsi `recover()` dipanggil dengan skema pemanggilan `defer func() { recover() }()` dan bukan `defer recover()` karena bentuk kedua ini penerapan yang keliru.

Setelah `recover()` berhasil menangkap panic, eksekusi berlanjut normal dan response 500 dikirim ke client.

## B.26.3. Implementasi Lengkap

Ok, berikut adalah kode lengkap webserver dengan custom error handler dan panic recovery middleware. Silakan tulis ke `main.go`.

```go
package main

import (
    "log"
    "net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 - Halaman tidak ditemukan"))
}

func internalErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Terjadi kesalahan pada server"))
}

func panicMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic: %v", err)
                internalErrorHandler(w, r)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!"))
    })

    mux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
        panic("something went wrong")
    })

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            notFoundHandler(w, r)
            return
        }
        w.Write([]byte("Home"))
    })

    handler := panicMiddleware(mux)

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", handler)
    if err != nil {
        log.Fatal(err)
    }
}
```

Fungsi `panicMiddleware()` membungkus `mux` secara keseluruhan, bukan per-handler, sehingga semua panic dari handler manapun tetap tertangkap.

## B.26.4. Testing

Jalankan server lalu coba ketiga skenario berikut.

```bash
curl http://localhost:9000/hello
```

Response: `Hello!` dengan status 200.

```bash
curl -i http://localhost:9000/tidak-ada
```

Response: `HTTP/1.1 404 Not Found` dengan body `404 - Halaman tidak ditemukan`.

```bash
curl -i http://localhost:9000/panic
```

Response: `HTTP/1.1 500 Internal Server Error` dengan body `500 - Terjadi kesalahan pada server`. Coba akses `/panic` beberapa kali, server tetap berjalan karena panic ditangkap oleh middleware sebelum sempat menghentikan proses.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.26-custom-error-handler">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.26...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
