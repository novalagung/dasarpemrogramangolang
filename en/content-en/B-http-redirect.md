# B.25. HTTP Redirect

Saat mengakses sebuah URL, kadang browser berpindah ke halaman lain secara otomatis tanpa interaksi apapun dari pengguna. Ini terjadi karena server mengirimkan response dengan status code khusus beserta header `Location` yang berisi URL tujuan, lalu browser mengikutinya secara otomatis. Kejadian ini disebut dengan HTTP redirect.

Pada chapter ini kita belajar cara melakukan redirect di Go menggunakan fungsi `http.Redirect()` dan memahami kapan harus menggunakan masing-masing jenis status code redirect.

## B.25.1. Fungsi `http.Redirect()`

Go menyediakan fungsi `http.Redirect()` untuk mengirimkan response redirect ke client.

```go
http.Redirect(w, r, "/tujuan", http.StatusMovedPermanently)
```

Parameter fungsi ini:

1. `w`: objek `http.ResponseWriter`
2. `r`: objek `*http.Request`
3. URL tujuan redirect (bisa berupa path relatif atau URL absolut)
4. HTTP status code redirect

## B.25.2. Jenis-Jenis Status Code Redirect

| Status Code | Konstanta                      | Keterangan                                                                                                                                                    |
| ----------- | ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 301         | `http.StatusMovedPermanently`  | Redirect permanen. Browser dan search engine akan memperbarui cache. Method request bisa berubah menjadi GET. |
| 302         | `http.StatusFound`             | Redirect sementara. Browser tidak memperbarui cache. Method request bisa berubah menjadi GET.                 |
| 303         | `http.StatusSeeOther`          | Redirect ke URL lain menggunakan GET, terlepas dari method aslinya. Umum digunakan setelah POST form berhasil.                |
| 307         | `http.StatusTemporaryRedirect` | Redirect sementara. Method request **dijamin** tidak berubah (POST tetap POST).                            |
| 308         | `http.StatusPermanentRedirect` | Redirect permanen. Method request **dijamin** tidak berubah (POST tetap POST).                             |

Perbedaan paling penting antara 301/302 dan 307/308 adalah pada method HTTP: 301 dan 302 bisa mengubah POST menjadi GET saat redirect, sedangkan 307 dan 308 mempertahankan method aslinya. Sedangkan 303 secara eksplisit selalu menggunakan GET di tujuan redirect.

## B.25.3. Praktik

Ok, langsung kita praktikkan. Buat file `main.go` berisi web server sederhana, lalu pada handler-nya kita terapkan _redirection_.

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/old-page", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/new-page", http.StatusMovedPermanently)
    })

    mux.HandleFunc("/new-page", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the new page!"))
    })

    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/dashboard", http.StatusFound)
    })

    mux.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            http.Redirect(w, r, "/thank-you", http.StatusSeeOther)
            return
        }
        w.Write([]byte(`<form method="POST"><button type="submit">Submit</button></form>`))
    })

    mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Dashboard"))
    })

    mux.HandleFunc("/thank-you", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Thank you!"))
    })

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", mux)
    if err != nil {
        log.Fatal(err)
    }
}
```

Mengacu ke kode di atas, ada tiga skenario redirect yang dipraktikkan. Penjelasannya sebagai berikut.

#### ◉ Redirect 301: URL Lama (/old-page) ke URL Baru (/new-page)

```go
mux.HandleFunc("/old-page", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/new-page", http.StatusMovedPermanently)
})
```

Route `/old-page` diredirect ke `/new-page` menggunakan status 301. Digunakan ketika sebuah halaman atau endpoint sudah dipindahkan secara permanen. Browser dan search engine akan memperbarui cache mereka sehingga request berikutnya langsung ke URL baru.

#### ◉ Redirect 302: Redirect Sementara (/login -> /dashboard)

```go
mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/dashboard", http.StatusFound)
})
```

Route `/login` diredirect ke `/dashboard` menggunakan status 302. Cocok untuk redirect yang bersifat kondisional atau sementara, misalnya setelah proses login. Browser tidak memperbarui cache sehingga `/login` tetap dicatat sebagai URL aslinya.

#### ◉ Redirect 303: Pola Post/Redirect/Get (PRG) Ke /thank-you

```go
mux.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        http.Redirect(w, r, "/thank-you", http.StatusSeeOther)
        return
    }
    w.Write([]byte(`<form method="POST"><button type="submit">Submit</button></form>`))
})
```

Route `/form` menangani dua kondisi: jika request adalah GET maka tampilkan form HTML, jika POST maka redirect ke `/thank-you` menggunakan status 303.

Ini biasa disebut dengan pola _Post/Redirect/Get_ (PRG). Setelah form berhasil disubmit via POST, browser diarahkan ke halaman konfirmasi menggunakan GET. Manfaatnya: kalau pengguna refresh halaman `/thank-you`, browser hanya melakukan GET ulang ke halaman itu, bukan mengirim ulang data POST yang bisa menyebabkan duplikasi data.

## B.25.4. Testing

Jalankan server lalu coba beberapa skenario berikut.

```bash
# Lihat response headers termasuk Location header
curl -i http://localhost:9000/old-page
```

Output-nya akan menampilkan response headers dari server, di antaranya ada `Location` yang berisi URL tujuan redirect.

```http
HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: /new-page
```

Secara default curl tidak mengikuti redirect. Untuk mengikutinya secara otomatis gunakan flag `-L`.

```bash
# Ikuti redirect secara otomatis
curl -L http://localhost:9000/old-page
```

Dengan flag `-L`, curl akan mengikuti header `Location` dan menampilkan response akhir dari `/new-page`: `Welcome to the new page!`

```bash
# Test POST ke /form, lihat Location header pada redirect 303
curl -i -X POST http://localhost:9000/form
```

Output-nya menampilkan `HTTP/1.1 303 See Other` dengan `Location: /thank-you`. Pola PRG bekerja: server tidak memproses ulang POST, melainkan meminta client mengakses `/thank-you` via GET.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div><a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.25-http-redirect">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.25...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
