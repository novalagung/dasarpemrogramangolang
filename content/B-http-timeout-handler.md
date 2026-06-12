# B.28. HTTP Timeout Handler

Pada chapter [B.22. Simple Configuration](/B-simple-configuration.html) kita telah mempelajari cara penggunaan `ReadTimeout` dan `WriteTimeout` pada `http.Server`. Timeout tersebut bersifat global untuk semua endpoint. Go juga menyediakan `http.TimeoutHandler` yang memungkinkan kita menentukan timeout **per handler**, sehingga endpoint yang lambat tidak mempengaruhi endpoint lain.

Pada chapter ini kita akan belajar timeout spesifik per handler.

## B.28.1. Perbedaan Server Timeout dan `http.TimeoutHandler`

| | Server Timeout (`ReadTimeout`/`WriteTimeout`) | `http.TimeoutHandler` |
| --- | --- | --- |
| Scope | Seluruh server | Per handler |
| Aksi saat timeout | Koneksi TCP diputus paksa | Response `503 Service Unavailable` dikirim |
| Goroutine handler | Tetap berjalan | Tetap berjalan (tidak di-cancel) |

Poin penting: `http.TimeoutHandler` **tidak membatalkan goroutine** handler yang sedang berjalan. Handler tetap berjalan di background, hanya response ke client yang dipotong. Untuk membatalkan proses di handler, perlu dikombinasikan dengan context dari `r.Context().Done()` (lihat chapter [B.30. Server Handler HTTP Request Cancellation](/B-server-handler-http-request-cancellation.html)).

## B.28.2. Penggunaan `http.TimeoutHandler`

Cara menggunakan `http.TimeoutHandler()` adalah dengan menjadikannya sebagai handler func untuk membungkus handler sebenarnya. Fungsi ini skema-nya seperti ini:

```go
http.TimeoutHandler(handler, timeout, message)
```

Penjelasan parameter:

 1. `handler`: handler yang dibungkus, bertipe `http.Handler`
 2. `timeout`: durasi timeout bertipe `time.Duration`
 3. `message`: pesan yang dikirim ke client saat timeout

## B.28.3. Implementasi

Mari coba praktikkan, buat file `main.go` dengan dua handler: satu cepat dan satu lambat, keduanya dibungkus `http.TimeoutHandler` dengan durasi 3 detik.

```go
package main

import (
    "log"
    "net/http"
    "time"
)

func handleFast(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("response cepat"))
}

func handleSlow(w http.ResponseWriter, r *http.Request) {
    time.Sleep(5 * time.Second)
    w.Write([]byte("response lambat selesai"))
}

func main() {
    mux := http.NewServeMux()

    mux.Handle("GET /fast", http.TimeoutHandler(
        http.HandlerFunc(handleFast),
        3*time.Second,
        "request timeout",
    ))

    mux.Handle("GET /slow", http.TimeoutHandler(
        http.HandlerFunc(handleSlow),
        3*time.Second,
        "request timeout: proses terlalu lama",
    ))

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", mux)
    if err != nil {
        log.Fatal(err)
    }
}
```

`handleFast()` langsung menulis response tanpa jeda, sehingga selesai jauh sebelum batas 3 detik. `handleSlow()` sebaliknya: ada `time.Sleep(5 * time.Second)` yang membuatnya selalu melewati timeout.

Masing-masing handler dibungkus dengan `http.TimeoutHandler` saat didaftarkan ke mux. `http.HandlerFunc(handleFast)` digunakan untuk mengonversi fungsi biasa ke tipe `http.Handler` yang dibutuhkan sebagai parameter pertama `http.TimeoutHandler`. Parameter ketiga adalah pesan yang dikirim ke client saat timeout terjadi.

## B.28.4. Testing

Jalankan server lalu coba dua skenario berikut.

```bash
curl http://localhost:9000/fast
```

Response: `response cepat` dengan status 200.

```bash
curl -i http://localhost:9000/slow
```

Response: `HTTP/1.1 503 Service Unavailable` dengan body `request timeout: proses terlalu lama` setelah menunggu 3 detik.

## B.28.5. Kombinasi dengan Context Cancellation

Untuk benar-benar membatalkan proses di dalam handler saat timeout, pantau `r.Context().Done()` di dalam handler.

```go
func handleSlow(w http.ResponseWriter, r *http.Request) {
    select {
    case <-time.After(5 * time.Second):
        w.Write([]byte("selesai"))
    case <-r.Context().Done():
        log.Println("handler dibatalkan karena timeout")
        return
    }
}
```

Dengan pola ini, ketika `http.TimeoutHandler` meng-cancel context setelah timeout, handler juga ikut berhenti sehingga tidak membuang resource.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.28-http-timeout-handler">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.28...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
