# WaitGroup

Sebelumnya kita telah belajar banyak mengenai channel, bagaimana sebuah/banyak goroutine dapat dikontrol dengan baik. Pada bab ini topik yang akan dipelajari masih relevan, yaitu tentang manajemen goroutine tetapi tidak menggunakan channel, melainkan `sync.WaitGroup`.

Golang menyediakan package `sync`, yang berisikan cukup banyak API untuk handle masalah threading (dalam kasus ini goroutine). Dari sekian banyak fitur yang ada, pada bab ini kita hanya akan mempelajari `sync.WaitGroup`.

`sync.WaitGroup` sendiri digunakan menunggu selesainya goroutine yang sedang berjalan. Cara penggunaannya sangat mudah, tinggal masukan jumlah goroutine yang dieksekusi sebagai parameter method `Add()` object cetakan `sync.WaitGroup`. Dan pada akhir tiap-tiap goroutine, panggil method `Done()`. Juga, pada baris kode setelah eksekusi goroutine, panggil method `Wait()`.

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

Kode di atas merupakan contoh penerapan `sync.WaitGroup` untuk pengelolahan goroutine. Fungsi `doPrint()` akan dijalankan sebagai goroutine, bertugas untuk menampilkan isi variabel `message`.

Variabel `wg` disiapkan bertipe `sync.WaitGroup`, nantinya digunakan untuk mengontrol proses sinkronisasi goroutines yang dijalankan. 

Di tiap perulangan statement `wg.Add(1)` dipanggil. Kode tersebut akan memberikan informasi kepada `wg` bahwa jumlah goroutine yang sedang diproses ditambah 1 (karena dipanggil 5 kali, maka `wg` akan sadar bahwa terdapat 5 buah goroutine sedang berjalan).

Di baris selanjutnya, fungsi `doPrint()` dieksekusi sebagai goroutine. Didalam fungsi tersebut, sebuah method bernama `Done()` dipanggil. Method ini digunakan untuk memberikan informasi kepada `wg` bahwa goroutine dimana method itu dipanggil sudah selesai. Sejumlah 5 buah goroutine dijalankan, maka method tersebut harus dipanggil 5 kali.

Statement `wg.Wait()` bersifat blocking, proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum sejumlah 5 goroutine selesai. Jika `Add(1)` dipanggil 5 kali, maka `Done()` juga harus dipanggil 5 kali.

![Contoh penerapan `sync.WaitGroup`](images/56_1_waitgroup.png)

Beberapa perbedaan antara channel dan `sync.WaitGroup`:
 - Channel tergantung kepada goroutine tertentu dalam penggunaannya, tidak seperti `sync.WaitGroup` yang dia tidak perlu tahu goroutine mana saja yang dijalankan, cukup tahu jumlah goroutine yang harus selesai
 - Penerapan `sync.WaitGroup` lebih mudah dibanding channel
 - Channel selain digunakan untuk manajemen goroutine, digunakan juga untuk komunikasi data antar goroutine, sedangkan WaitGroup digunakan untuk pengontrolan eksekusi goroutine saja.
 - Performa `sync.WaitGroup` lebih baik dibanding channel, sumber: https://groups.google.com/forum/#!topic/golang-nuts/whpCEk9yLhc

Jika muncul pertanyaan manakah yang lebih baik, apakah WaitGroup atau channel, maka kembali ke kebutuhan, karena kedua API tersebut memiliki kelebihan masing-masing.
