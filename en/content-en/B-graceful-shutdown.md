# B.24. Graceful Shutdown

Saat server dihentikan secara paksa (misalnya dengan `Ctrl+C`), request yang sedang diproses oleh web server akan langsung terputus tanpa sempat mengirim response. Pada chapter ini kita belajar cara menghentikan server secara _graceful_, yaitu server menunggu semua request yang sedang berjalan selesai sebelum benar-benar berhenti.

> Graceful shutdown sangat penting untuk diterapkan terutama di production pada aplikasi dengan traffic tinggi. Tanpa mekanisme ini, setiap proses deployment atau restart server berpotensi memutus ratusan request yang sedang berjalan sekaligus, yang bisa berujung pada data tidak konsisten, transaksi gagal di tengah jalan, atau pengalaman buruk di sisi pengguna.

## B.24.1. Masalah Shutdown Paksa

Seperti yang sudah disampaikan sekilas di atas, jika server dihentikan dengan cara biasa, proses di back end yang masih berjalan dan belum selesai akan langsung diputus. Hal ini bisa menyebabkan:

- Response tidak terkirim ke client
- Transaksi database tidak ter-_commit_
- File yang sedang ditulis menjadi korup
- Masalah lainnya tergantung seberapa kompleks proses di dalam web server

Solusi agar masalah tersebut tidak terjadi adalah memanfaatkan `server.Shutdown()` yang disediakan oleh Go.

## B.24.2. Cara Kerja Graceful Shutdown

Proses graceful shutdown Go bekerja dengan urutan seperti ini:

1. Server berhenti menerima koneksi baru
2. Server menunggu semua request yang sedang aktif selesai
3. Setelah semua request selesai (atau timeout tercapai), server benar-benar berhenti

## B.24.3. Praktik

Ok, mari kita praktikkan. Siapkan file `main.go` dengan kode berikut.

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(2 * time.Second)
        w.Write([]byte("Hello!"))
    })

    server := &http.Server{
        Addr:    ":9000",
        Handler: mux,
    }

    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    go func() {
        log.Println("server started at localhost:9000")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()

    <-ctx.Done()
    log.Println("shutdown signal received")

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := server.Shutdown(shutdownCtx); err != nil {
        log.Fatal("forced shutdown:", err)
    }

    log.Println("server exited gracefully")
}
```

Kode terkait mux dan handler sepertinya cukup clear, jadi kita fokus saja ke penjelasan bagian yang terasa baru.

#### ◉ `signal.NotifyContext()`

Fungsi `signal.NotifyContext()` (tersedia sejak Go 1.16) berfungsi untuk membuat context yang otomatis di-cancel ketika sinyal `SIGINT` (Ctrl+C) atau `SIGTERM` diterima. Ini adalah cara modern menangkap sinyal OS di Go.

```go
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()
```

> Lebih jelasnya mengenai context dibahas di chapter [A.65. Concurrency Pattern: Context Cancellation Pipeline](/A-pipeline-context-cancellation.html)

#### ◉ Sinyal `syscall.SIGINT`, `syscall.SIGTERM`, dan lainnya.

Fungsi `signal.NotifyContext()` parameter ke-2 adalah variadic, jadi bisa diisi banyak jenis sinyal. Go mendukung banyak jenis sinyal, beberapa yang umum digunakan pada web server:

| Sinyal            | Nilai | Keterangan                                                                                         |
| ----------------- | ----- | -------------------------------------------------------------------------------------------------- |
| `syscall.SIGINT`  | 2     | Dikirim saat pengguna menekan `Ctrl+C` di terminal                                                 |
| `syscall.SIGTERM` | 15    | Sinyal terminasi standar, dikirim oleh `kill <PID>` atau container orchestrator seperti Kubernetes |
| `syscall.SIGHUP`  | 1     | Hangup signal, sering digunakan untuk meminta proses reload konfigurasi                            |
| `syscall.SIGQUIT` | 3     | Quit signal (`Ctrl+\`), serupa `SIGTERM` namun juga menghasilkan core dump     |

Selain itu ada jenis sinyal lainnya seperti `syscall.SIGKILL (9)` dan `syscall.SIGSTOP`, namun keduanya tidak bisa ditangkap lewat `signal.NotifyContext` karena langsung ditangani oleh kernel.

#### ◉ Server di-run dalam goroutine

Server dijalankan di goroutine terpisah agar goroutine utama bisa menunggu sinyal shutdown.

```go
go func() {
    log.Println("server started at localhost:9000")
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()
```

Error `http.ErrServerClosed` bukan error sungguhan, melainkan indikator informasi bahwa server berhasil di-shutdown. Oleh karena itu error ini perlu dikecualikan dari pengecekan.

#### ◉ Menunggu sinyal

Di main goroutine, statement blocked pada baris berikut:

```go
<-ctx.Done()
```

Goroutine utama diblokir di sini hingga sinyal shutdown diterima. Sinyal shutdown yang dimaksud adalah `syscall.SIGINT` (Ctrl+C) atau `syscall.SIGTERM`.

#### ◉ `server.Shutdown()`

Setelah baris `<-ctx.Done()` jalan, maka bisa disimpulkan sinyal shutdown sudah diterima. Eksekusi program kemudian lanjut hingga kode berikut, di mana ada proses pembuatan konteks untuk shutdown `shutdownCtx` serta operasi `server.Shutdown()`.

```go
shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

if err := server.Shutdown(shutdownCtx); err != nil {
    log.Fatal("forced shutdown:", err)
}
```

`server.Shutdown()` menghentikan server secara graceful. Context dengan timeout 10 detik digunakan di sini untuk memastikan bahwa jika ada request yang menggantung terlalu lama, server tetap akan berhenti setelah 10 detik.

## B.24.4. Pengujian

Jalankan server, lalu kirim request yang butuh waktu lama:

```bash
curl http://localhost:9000/
```

Sebelum request selesai, tekan `Ctrl+C`. Server akan menunggu request tersebut selesai terlebih dahulu sebelum benar-benar berhenti.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div><a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.24-graceful-shutdown">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.24...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
