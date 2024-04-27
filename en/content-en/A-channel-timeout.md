# A.35. Channel - Timeout

Teknik channel timeout digunakan untuk kontrol waktu penerimaan data pada channel, berapa lama channel tersebut harus menunggu hingga akhirnya suatu penerimaan data dianggap timeout.

Durasi penerimaan kita tentukan sendiri. Ketika tidak ada aktivitas penerimaan data dalam durasi tersebut, blok timeout dijalankan.

## A.35.1. Penerapan Channel Timeout

Berikut adalah program sederhana contoh pengaplikasian timeout pada channel. Sebuah goroutine dijalankan dengan tugas adalah mengirimkan data secara berulang dalam interval tertentu, dengan durasi interval-nya sendiri adalah acak/random.

```go
package main

import "fmt"
import "math/rand"
import "runtime"
import "time"

func sendData(ch chan<- int) {
    randomizer := rand.New(rand.NewSource(time.Now().Unix()))

    for i := 0; true; i++ {
        ch <- i
        time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
    }
}
```

Selanjutnya, disiapkan perulangan tanpa henti, yang di setiap perulangan ada seleksi kondisi channel menggunakan `select`.

```go
func retreiveData(ch <-chan int) {
    loop:
    for {
        select {
        case data := <-ch:
            fmt.Print(`receive data "`, data, `"`, "\n")
        case <-time.After(time.Second * 5):
            fmt.Println("timeout. no activities under 5 seconds")
            break loop
        }
    }
}
```

Ada 2 blok kondisi pada `select` tersebut.

- Kondisi `case data := <-messages:`, akan terpenuhi ketika ada serah terima data pada channel `messages`.
- Kondisi `case <-time.After(time.Second * 5):`, akan terpenuhi ketika tidak ada aktivitas penerimaan data dari channel dalam durasi 5 detik. Blok inilah yang kita sebut sebagai blok timeout.

Terakhir, kedua fungsi tersebut dipanggil di `main()`.

```go
func main() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan int)

    go sendData(messages)
    retreiveData(messages)
}
```

Muncul output setiap kali ada penerimaan data dengan delay waktu acak. Ketika dalam durasi 5 detik tidak ada aktivitas penerimaan sama sekali, maka dianggap timeout dan perulangan pengecekkan channel dihentikan.

![Channel timeout](images/A_channel_timeout_1_channel_delay.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.35-channel-timeout">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.35...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
