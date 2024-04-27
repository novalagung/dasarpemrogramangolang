# B.20. Custom Multiplexer

Pada chapter ini, kita akan belajar cara membuat custom multiplexer sendiri, lalu memanfaatkannya untuk keperluan manajemen middleware.

Silakan salin project sebelumnya (chapter [B.19. Middleware http.Handler](/B-middleware-using-http-handler.html)) ke folder baru untuk keperluan pembelajaran.

## B.20.1. Pembuatan Custom Mux

Pada chapter sebelumnya, default mux milik Go digunakan untuk routing dan implementasi middleware. Kali ini default mux tersebut tidak digunakan karena mux baru akan dibuat.

Sebenarnya, pembuatan mux baru tidaklah cukup, karena mux baru tidak memiliki perbedaan signifikan dibanding default mux. Agar mux baru menjadi lebih berguna, mux baru tersebut perlu meng-embed `http.ServeMux` dan kita juga perlu mempersiapkan beberapa method.

OK, mari kita praktekan. Ubah isi fungsi `main()` menjadi seperti berikut.

```go
mux := new(CustomMux)

mux.HandleFunc("/student", ActionStudent)

mux.RegisterMiddleware(MiddlewareAuth)
mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

server := new(http.Server)
server.Addr = ":9000"
server.Handler = mux

fmt.Println("server started at localhost:9000")
server.ListenAndServe()
```

Objek `mux` dicetak dari struct `CustomMux` yang mana nantinya struct ini dibuat dengan meng-embed `http.ServeMux`.

Registrasi middleware juga perlu diubah, sekarang menggunakan method `.RegisterMiddleware()` milik `CustomMux`.

Selanjutnya, di file `middleware.go` siapkan struct `CustomMux`. Selain meng-embed objek mux milik Go, siapkan juga satu variabel bertipe `[]func(next http.Handler) http.Handler`.

```go
type CustomMux struct {
    http.ServeMux
    middlewares []func(next http.Handler) http.Handler
}
```

Buat fungsi `RegisterMiddleware()`. Middleware yang didaftarkan ditampung oleh slice `.middlewares`.

```go
func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
    c.middlewares = append(c.middlewares, next)
}
```

Lalu buat method `ServeHTTP()`. Method ini diperlukan agar custom mux memenuhi kriteria interface `http.Handler`.

```go
func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var current http.Handler = &c.ServeMux

    for _, next := range c.middlewares {
        current = next(current)
    }

    current.ServeHTTP(w, r)
}
```

Method `ServeHTTP()` milik mux dipanggil setiap kali ada HTTP request. Dengan perubahan di atas, maka setiap kali ada request masuk pasti akan melewati middleware-middleware terlebih dahulu secara berurutan.

- Jika lolos middleware ke-1, lanjut ke-2
- Jika lolos middleware ke-2, lanjut ke-3
- ... dan seterusnya

## B.20.2. Testing

Jalankan aplikasi.

![Run the server](images/B_http_basic_auth_2_run_server.png)

Lalu test menggunakan `curl`, hasilnya pasti sama dengan pada chapter sebelumnya.

![Consume API](images/B_http_basic_auth_3_test_api.png)

Jika ada keperluan untuk menambahkan middleware baru lainnya, cukup registrasikan lewat `.RegisterMiddleware()`. Pengaplikasian teknik custom mux ini membuat manajemen middleware menjadi lebih mudah.

> Fun fact: semua *3rd party* router di Go (seperti Gin, Chi, Gorilla Mux, dan lainnya) menerapkan teknik custom multiplexer

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.20-custom-mux-multiplexer">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.20...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
