# B.3. Routing Static Assets

Pada bagian ini kita akan mempelajari cara routing static assets / static contents. Static assets yang dimaksud adalah seperti file statis css, js, gambar, dan lainnya.

Ok, mari belajar sambil praktek.

## B.3.1. Struktur Aplikasi

Buat project baru, siapkan file dan folder dengan struktur sesuai dengan gambar berikut.

![Structure](images/B_routing_static_assets_1_structure.png)

Dalam folder `assets`, isi dengan file apapun bebas, bisa gambar atau file js. Selanjutnya masuk ke bagian routing static assets.

## B.3.2. Routing

Berbeda dengan routing menggunakan `http.HandleFunc()`, routing static assets lebih mudah. Silakan tulis kode berikut dalam `main.go`, setelahnya kita akan bahas secara mendetail.

```go
package main

import "fmt"
import "net/http"

func main() {
    http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
```

Syarat yang dibutuhkan untuk routing static assets masih sama dengan routing handler, yaitu perlu didefiniskan rute beserta handler-nya. Hanya saja pembedanya di sini adalah dalam routing static assets yang digunakan adalah `http.Handle()`, bukan `http.HandleFunc()`.

 1. Rute terpilih adalah `/static/`, maka nantinya semua request yang di awali dengan `/static/` akan diarahkan ke sini. Registrasi rute menggunakan `http.Handle()` adalah berbeda dengan routing menggunakan `http.HandleFunc()`, lebih jelasnya akan ada sedikit penjelasan pada chapter lain.

 2. Sedang untuk handler-nya bisa di-lihat, ada pada parameter ke-2 yang isinya statement `http.StripPrefix()`. Sebenarnya actual handler nya berada pada `http.FileServer()`. Fungsi `http.StripPrefix()` hanya digunakan untuk membungkus actual handler.

Fungsi `http.FileServer()` mengembalikan objek ber-tipe `http.Handler`. Fungsi ini berguna untuk merespon http request dengan konten yang ada di dalam folder `assets` sesuai permintaan.

Jalankan `main.go`, lalu test hasilnya di browser `http://localhost:9000/static/`.

![Structure](images/B_routing_static_assets_2_preview.png)

## B.3.3. Penjelasan

Penjelasan akan lebih mudah dipahami jika disajikan juga contoh praktek, maka sejenak kita coba bahas menggunakan contoh sederhana berikut.

```go
http.Handle("/", http.FileServer(http.Dir("assets")))
```

Jika dilihat pada struktur folder yang sudah di-buat, di dalam folder `assets` terdapat file bernama `site.css`. Maka dengan bentuk routing pada contoh sederhana di atas, request ke `/site.css` akan diarahkan ke path `./site.css` (relatif dari folder `assets`). Permisalan contoh lainnya:

 * Request ke `/site.css` mengarah path `./site.css` relatif dari folder `assets`
 * Request ke `/script.js` mengarah path `./script.js` relatif dari folder `assets`
 * Request ke `/some/folder/test.png` mengarah path `./some/folder/test.png` relatif dari folder `assets`
 * ... dan seterusnya

> Fungsi `http.Dir()` berguna untuk *adjustment path parameter*. Separator dari path yang di-definisikan otomatis di-konversi ke path separator sesuai sistem operasi.

Sekarang coba perhatikan kode berikut.

```go
http.Handle("/static", http.FileServer(http.Dir("assets")))
```

Dengan skema routing di atas, maka:

 * Request ke `/static/site.css` mengarah ke `./static/site.css` relatif dari folder `assets`
 * Request ke `/static/script.js` mengarah ke `./static/script.js` relatif dari folder `assets`
 * Request ke `/static/some/folder/test.png` mengarah ke `./static/some/folder/test.png` relatif dari folder `assets`
 * ... dan seterusnya

Bisa dilihat bahwa rute yang didaftarkan juga akan digabung dengan path destinasi file yang dicari, dan ini menjadikan path tidak valid. File `site.css` berada pada path `assets/site.css`, sedangkan dari routing di atas pencarian file mengarah ke path `assets/static/site.css`. Di sinilah kegunaan dari fungsi `http.StripPrefix()`.

Fungsi `http.StripPrefix()` berguna untuk menghapus prefix dari endpoint yang diakses. Request ke URL yang di awali dengan `/static/` akan diambil informasi endpoint-nya tanpa prefix `/static/`.

 * Request ke `/static/site.css` mengarah ke `site.css`
 * Request ke `/static/script.js` mengarah ke `script.js`
 * Request ke `/static/some/folder/test.png` mengarah ke `some/folder/test.png`
 * ... dan seterusnya

Dengan penerapan `http.StripPrefix()` maka routing static assets menjadi valid, karena file yang di-request akan cocok dengan path folder yang telah dibuat.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.3-routing-static-assets">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.3...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
