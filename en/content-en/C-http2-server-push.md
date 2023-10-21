# C.25. HTTP/2 dan HTTP/2 Server Push

HTTP/2 adalah versi terbaru protokol HTTP, dikembangkan dari protokol [SPDY](https://tools.ietf.org/html/draft-mbelshe-httpbis-spdy-00) yang diinisiasi oleh Google.

Protokol ini sekarang sudah kompatibel dengan banyak browser di antaranya: Chrome, Opera, Firefox 9, IE 11, Safari, Silk, dan Edge.

Kelebihan HTTP/2 dibanding HTTP 1.1 (protokol yang umumnya digunakan) sebagian besar adalah pada performa dan sekuriti. Berikut merupakan beberapa point yang menjadi kelebihan dari protokol baru ini.

 - Backward compatible dengan HTTP 1.1
 - Kompresi data pada HTTP Headers
 - Multiplexing banyak request (dalam satu koneksi TCP)
 - HTTP/2 Server Push

Pada chapter ini kita akan belajar cara menerapkan HTTP/2 dan salah satu fitur milik protokol ini yaitu HTTP/2 Server Push.

> Mengenai multiplexing banyak request tidak akan kita bahas pada buku ini, silakan coba pelajari sendiri jika tertarik, menggunakan library cmux.

## C.25.1. HTTP/2 di Golang

Golang memiliki dukungan sangat baik terhadap HTTP/2. Dengan cukup meng-enable fasilitas TLS/HTTPS maka aplikasi golang secara otomatis menggunakan HTTP/2.

Untuk memastikan mari kita langsung praktekkan, coba duplikat project pada chapter sebelumnya (**A.23. HTTPS/TLS Web Server**) sebagai project baru, jalankan aplikasinya lalu cek di browser chrome. Gunakan chrome extension [HTTP/2 and SPDY indicator](https://chrome.google.com/webstore/detail/http2-and-spdy-indicator/mpbpobfflnpcgagjijhmgnchggcjblin?hl=en) untuk menge-test apakah HTTP/2 sudah enabled.

![SPDY checker](images/C_http2_server_push_1_spdy_checker.png)

Perlu diketahui untuk golang versi sebelum **1.6** ke bawah, secara default HTTP/2 tidak akan di-enable. Perlu memanggil fungsi `http2.ConfigureServer()` secara eksplist untuk meng-enable HTTP/2. Fungsi tersebut tersedia dalam package `golang.org/x/net/http2`. Lebih jelasnya silakan baca [laman dokumentasi](https://godoc.org/golang.org/x/net/http2).

## C.25.2. HTTP/2 Server Push

HTTP/2 Server Push adalah salah satu fitur pada HTTP/2, berguna untuk mempercepat response dari request, dengan cara data yang akan di-response dikirim terlebih dahulu oleh server.

Fitur server push ini cocok digunakan untuk push data assets, seperti: css, gambar, js, dan file assets lainnya.

Lalu apakah server push ini bisa dimanfaatkan untuk push data JSON, XML, atau sejenisnya? Sebenarnya bisa, hanya saja ini akan menyalahi tujuan dari penciptaan server push sendiri dan hasilnya tidak akan optimal, karena sebenernya server push tidak murni bidirectional, masih perlu adanya request ke server untuk mendapatkan data yg sudah di push oleh server itu sendiri.

> HTTP/2 server push bukanlah pengganti dari websocket. Gunakan websocket untuk melakukan komunikasi bidirectional antara server dan client.

Untuk mengecek suport-tidak-nya server push, lakukan casting pada objek `http.ResponseWriter` milik handler ke interface `http.Pusher`, lalu manfaatkan method `Push()` milik interface ini untuk push data dari server.

> Fasilitas server push ini hanya bisa digunakan pada golang versi 1.8 ke-atas.

## C.25.3. Praktek

Mari kita praktekan. Buat project baru, buat file `main.go`, isi dengan kode berikut.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.Handle("/static/", 
        http.StripPrefix("/static/", 
            http.FileServer(http.Dir("."))))

    log.Println("Server started at :9000")
    err := http.ListenAndServeTLS(":9000", "server.crt", "server.key", nil)
    if err != nil {
        panic(err)
    }
}
```

Dalam folder proyek baru di atas, siapkan juga beberapa file lain:

 - File `app.js`
 - File `app.css`
 - File `server.crt`, salin dari proyek sebelumnya
 - File `server.key`, salin dari proyek sebelumnya

Selanjutnya siapkan string html, nantinya akan dijadikan sebagai output rute `/`.

```go
const indexHTML = `
    <!DOCTYPE html>
    <html>
        <head>
            <title>Hello World</title>
            <script src="/static/app.js"></script>
            <link rel="stylesheet" href="/static/app.css"">
        </head>
        <body>
        Hello, gopher!<br>
        <img src="https://blog.golang.org/go-brand/logos.jpg" height="100">
        </body>
    </html>
`
```

Siapkan rute `/`, render string html tadi sebagai output dari endpoint ini.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    if pusher, ok := w.(http.Pusher); ok {
        if err := pusher.Push("/static/app.js", nil); err != nil {
            log.Printf("Failed to push: %v", err)
        }

        if err := pusher.Push("/static/app.css", nil); err != nil {
            log.Printf("Failed to push: %v", err)
        }
    }

    fmt.Fprintf(w, indexHTML)
})
```

Pada handler di atas bisa kita lihat bahwa selain me-render string html, rute ini juga memiliki tugas lain yaitu push assets.

Cara untuk mendeteksi apakah server push di-support, adalah dengan meng-casting objek `http.ResponseWriter` ke tipe `http.Pusher`. Proses casting tersebut mengembalikan dua buah data.

 1. Data objek yang sudah di casting.
 2. Variabel bertipe `bool` penanda aplikasi kita mendukung mendukung server push atau tidak.

Jika nilai variabel `ok` adalah `true`, maka server push di-support.

Method `Push()` milik `http.Pusher` digunakan untuk untuk push data. Endpoint yang disisipkan sebagai argumen, datanya akan di push ke front end oleh server, dalam contoh di atas adalah `/static/app.js` dan `/static/app.css`.

Server push ini adalah per endpoint atau rute. Jika ada rute lain, maka dua assets di atas tidak akan di push, kecuali method `Push()` dipanggil lagi dalam rute lain tersebut.

> Daripada memanggil method push satu-per-satu pada banyak handler, lebih baik jadikan sebagai middleware terpisah.

Kegunaan dari fungsi `fmt.Fprintf()` adalah untuk render html, sama seperti `w.Write()`.

OK, jalankan aplikasi lalu test.

## C.25.4. Testing

Perbedaan antara aplikasi yang menerapkan HTTP/2 dan tidak, atau yang menerapkan server push atau tidak; adalah tidak terasa bedanya jika hanya di-test lewat lokal saja.

Untuk mengecek HTTP/2 diterapkan atau tidak, kita bisa gunakan Chrome extension **HTTP/2 and SPDY indicator**.

Untuk mengecek server push pada tiap request sebenernya bisa hanya cukup menggunakan chrome dev tools, namun fitur ini hanya tersedia pada [Chrome Canary](https://www.google.com/chrome/browser/canary.html). Download browser tersebut lalu install, gunakan untuk testing aplikasi kita.

Pada saat mengakses `https://localhost:9000` pastikan developer tools sudah aktif (klik kanan, inspect element), lalu buka tab **Network**.

![SPDY indicator](images/C_http2_server_push_2_spdy_indicator.png)


Untuk endpoint yang menggunakan server push, pada kolom **Protocol** nilainya adalah **spdy**. Pada screenshot di atas terlihat bahwa assets `app.js` dan `app.css` dikirim lewat server push.

> Jika kolom Protocol tidak muncul, klik kanan pada kolom, lalu centang **Protocol**.

Berikut merupakan variasi nilai pada kolom protocol.

 - **http/1.1**, protokol yang kita gunakan mulai tahun 1999.
 - **spdy**, versi pertama spesifikasi HTTP/2 dari google, mengawali terciptanya HTTP/2.
 - **h2**, kependekan dari "HTTP 2".
 - **h2c**, kependekan dari "HTTP 2 Cleartext". HTTP/2 lewat kanal yang tidak ter-enkripsi.

Selain dari kolom protocol, penanda server push bisa dilihat juga lewat grafik **Waterfall**. Garis warna ungu muda pada grafik tersebut adalah start time. Untuk endpoint yang di push maka bar chart akan muncul sebelum garis ungu atau start time.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.25-http2-server-push">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.25...</a>
</div>

---

<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>
