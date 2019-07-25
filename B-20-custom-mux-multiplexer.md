# B.20. Custom Multiplexer

Pada bab ini, kita akan belajar membuat custom multiplexer sendiri, lalu memanfaatkannya untuk mempermudah manajemen middleware.

Silakan salin projek sebelumnya, [Bab B.19. Middleware http.Handler](/B-19-middleware-using-http-handler.html), ke folder baru untuk keperluan pembelajaran.

## B.20.1. Pembuatan Custom Mux

Pada bab sebelumnya, default mux milik Go digunakan untuk routing dan implementasi middleware. Kali ini default mux tersebut tidak digunakan, kita akan buat mux baru.

Namun pembuatan mux baru tidaklah cukup, karena *naturally* mux baru tersebut tidak akan ada beda dengan default mux. Oleh karena itu agar lebih berguna, kita akan buat tipe mux baru, meng-embed `http.ServeMux` kedalamnya, lalu membuat beberapa hal dalam struct tersebut.

OK, langsung saja kita praktekan. Ubah isi fungsi main menjadi seperti berikut.

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

Objek `mux` dicetak dari struct `CustomMux` yang jelasnya akan di buat. Struct ini di dalamnya meng-embed `http.ServeMux`.

Registrasi middleware juga diubah, sekarang menggunakan method `.RegisterMiddleware()` milik mux.

Pada file `middleware.go`, siapkan struct `CustomMux`. Selain meng-embed objek mux milik Go, siapkan juga satu variabel bertipe slice-dari-tipe-fungsi-middleware.

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

Lalu buat method `ServeHTTP`. Method ini diperlukan dalam custom mux agar memenuhi kriteria interface `http.Handler`.

```go
func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var current http.Handler = &c.ServeMux

    for _, next := range c.middlewares {
        current = next(current)
    }

    current.ServeHTTP(w, r)
}
```

Method `ServeHTTP()` milik mux adalah method yang pasti dipanggil pada web server, di setiap request yang masuk.

Dengan perubahan di atas, setiap kali ada request masuk pasti akan melewati middleware-middleware terlebih dahulu secara berurutan. Jika lolos middleware ke-1, lanjut ke-2; jika lolos middleware ke-2, lanjut ke-3; dan seterusnya.

## B.20.2. Testing

Jalankan aplikasi.

![Run the server](images/B.18_2_run_server.png)

Lalu test menggunakan `curl`, hasilnya adalah sama dengan pada bab sebelumnya.

![Consume API](images/B.18_3_test_api.png)

Jika ada keperluan untuk menambahkan middleware baru lainnya, cukup registrasikan lewat `.RegisterMiddleware()`. Source code menjadi lebih rapi dan nyaman untuk dilihat.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang/tree/master/chapter-B.20-custom-mux-multiplexer">https://github.com/novalagung/dasarpemrogramangolang/.../chapter-B.20...</a>
</div>
