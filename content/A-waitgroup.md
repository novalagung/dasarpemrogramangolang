# A.59. WaitGroup

Sebelumnya kita telah belajar banyak mengenai channel, yang dimana fungsi utama-nya adalah untuk sharing/kirim data antar goroutine. Selain untuk komunikasi data, channel secara tidak langsung bisa dimanfaatkan untuk kontrol goroutine.

Go menyediakan package `sync`, berisi cukup banyak API untuk manajemen operasi multiprocessing (goroutine), salah satunya di antaranya adalah yang kita bahas pada chapter ini, yaitu `sync.WaitGroup`.

Kegunaan `sync.WaitGroup` adalah untuk sinkronisasi goroutine. Berbeda dengan channel, `sync.WaitGroup` memang dirancang khusus untuk pengelolahan goroutine, dengan penggunaan relatif lebih mudah dan efektif dibanding channel.

> Sebenarnya kurang pas jika membandingkan `sync.WaitGroup` dan channel, karena fungsi utama dari keduanya adalah berbeda. Channel untuk keperluan sharing data antar goroutine, sedangkan `sync.WaitGroup` untuk sinkronisasi goroutine.

## A.59.1. Penerapan `sync.WaitGroup`

`sync.WaitGroup` digunakan untuk menunggu goroutine. Cara pengaplikasiannya sangat mudah, tinggal masukan jumlah goroutine yang dieksekusi, sebagai parameter method `Add()` pada object cetakan `sync.WaitGroup`, kemudian di akhir setiap goroutine pastikan untuk memanggil method `Done()`. Lalu gunakan method `Wait()` untuk menunggu eksekusi semua goroutine selesai.

Agar lebih jelas, silakan coba kode berikut.

```go
package main

import "sync"
import "runtime"
import "fmt"

func doPrint(wg *sync.WaitGroup, message string) {
    defer wg.Done()
    fmt.Println(message)
}

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        var data = fmt.Sprintf("data %d", i)

        wg.Add(1)
        go doPrint(&wg, data)
    }

    wg.Wait()
}
```

Kode di atas merupakan contoh penerapan `sync.WaitGroup` untuk pengelolahan goroutine. Fungsi `doPrint()` akan dijalankan sebagai goroutine, dengan tugas menampilkan isi variabel `message`.

Variabel `wg` dibuat dengan tipe data `sync.WaitGroup`. Variabel ini digunakan sebagai kontrol dan sinkronisasi goroutines yang dijalankan.

Di tiap perulangan statement `wg.Add(1)` dipanggil. Kode tersebut akan memberikan informasi kepada `wg` bahwa jumlah goroutine yang sedang di proses ditambah 1 (karena dipanggil 5 kali, maka `wg` akan sadar bahwa terdapat 5 buah goroutine sedang berjalan).

Di baris selanjutnya, fungsi `doPrint()` dieksekusi sebagai goroutine. Di dalam fungsi tersebut, sebuah method bernama `Done()` dipanggil. Method ini digunakan untuk memberikan informasi kepada `wg` bahwa goroutine di mana method itu dipanggil sudah selesai. Sejumlah 5 buah goroutine dijalankan, maka method tersebut harus dipanggil 5 kali.

Statement `wg.Wait()` bersifat blocking, proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum sejumlah 5 goroutine selesai. Jika `Add(1)` dipanggil 5 kali, maka `Done()` juga harus dipanggil 5 kali.

Output program di atas:

![Contoh penerapan `sync.WaitGroup`](images/A_waitgroup_1_waitgroup.png)

> `sync.WaitGroup` merupakan salah satu tipe yang *thread safe*. Kita tidak perlu khawatir terhadap potensi *race condition* karena variabel bertipe ini aman untuk digunakan di banyak goroutine secara paralel.

## A.59.2. Perbedaan WaitGroup Dengan Channel

Bukan sebuah perbandingan yang *fair*, tapi jika dilihat perbedaan antara channel dan `sync.WaitGroup` kurang lebih ada di bagian ini:

 - Channel tergantung kepada goroutine tertentu dalam penggunaannya, tidak seperti `sync.WaitGroup` yang dia tidak perlu tahu goroutine mana saja yang dijalankan, cukup tahu jumlah goroutine yang harus selesai
 - Penerapan `sync.WaitGroup` lebih mudah dibanding channel
 - Kegunaan utama channel adalah untuk komunikasi data antar goroutine. Sifatnya yang blocking bisa kita manfaatkan untuk manage goroutine; sedangkan WaitGroup khusus digunakan untuk sinkronisasi goroutine
 - Performa `sync.WaitGroup` lebih baik dibanding channel, sumber: https://groups.google.com/forum/#!topic/golang-nuts/whpCEk9yLhc

Kombinasi yang tepat antara `sync.WaitGroup` dan channel sangat penting, keduanya diperlukan dalam concurrent programming program performansinya bisa maksimal.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.59-waitgroup">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.59...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
