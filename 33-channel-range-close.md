# Channel - Range dan Close

Penerimaan data lewat channel yang dipakai oleh banyak goroutine, akan lebih mudah dengan memanfaatkan keyword `for` - `range`.

`for` - `range` jika diterapkan pada channel, akan melakukan perulangan tanpa henti. Perulangan tersebut tetap berjalan meski tidak ada transaksi pada channel, dan hanya akan berhenti jika status channel berubah menjadi **closed** atau sudah ditutup. Fungsi `close` digunakan utuk menutup channel.

Channel yang sudah ditutup tidak bisa digunakan lagi untuk menerima maupun mengirim data. Menjadikan penerimaan data menggunakan `for` - `range` juga ikut berhenti.

## Penerapan `for` - `range` - `close` Pada Channel

Berikut adalah contoh program yang menggunakan `for` - `range` untuk pengambilan data dari channel.

Pertama siapkan fungsi `sendMessage()` untuk handle pengiriman data. Didalam fungsi ini akan dijalankan perulangan sebanyak 20 kali, ditiap perulangannya data dikirim lewat channel. Setelah semua data terkirim, channel di-close.

```go
func sendMessage(ch chan<- string) {
    for i := 0; i < 20; i++ {
        ch <- fmt.Sprintf("data %d", i)
    }
    close(ch)
}
```

Siapkan juga fungsi `printMessage()` untuk handle penerimaan data. Didalamnya, channel akan di-looping menggunakan `for` - `range`, yang kemudian ditampilkan data-nya.

```go
func printMessage(ch <-chan string) {
    for message := range ch {
        fmt.Println(message)
    }
}
```

Bat channel baru di fungsi `main`, jalankan `sendMessage()` sebagai goroutine. Jalankan juga `printMessage()`. Dengan ini 20 data dikirimkan lewat goroutine baru, dan nantinya diterima di goroutine utama.

```go
func main() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan string)
    go sendMessage(messages)
    printMessage(messages)
}
```

Setelah 20 data sukses dikirim dan diterima, channel `ch` akan dimatikan (`close(ch)`). Membuat perulangan data channel dalam `printMessage()` juga akan berhenti.

![Penerapan for-range-close pada channel](images/33_1_for_range_close.png)

## Channel Direction

Ada yang unik dengan fitur parameter channel yang disediakan Golang. Level akses channel bisa ditentukan, apakah hanya sebagai penerima, pengirim, atau penerima sekaligus pengirim. Konsep ini disebut dengan **channel direction**.

Cara pemberian level akses adalah dengan menambahkan tanda `<-` sebelum atau setelah keyword `chan`. Untuk lebih jelasnya bisa dilihat di list berikut.

| Sintaks | Penjelasan |
| :------- | :--------- |
| `ch chan string` | Parameter `ch` bisa digunakan untuk **mengirim** dan **menerima** data |
| `ch chan<- string` | Parameter `ch` hanya bisa digunakan untuk **mengirim** data |
| `ch <-chan string` | Parameter `ch` hanya bisa digunakan untuk **menerima** data |

Pada kode di atas bisa dilihat bahwa secara default channel akan memiliki kemampuan untuk mengirim dan menerima data. Untuk mengubah channel tersebut agar hanya bisa mengirim atau menerima saja, dengan memanfaatkan simbol `<-`.

Sebagai contoh fungsi `sendMessage(ch chan<- string)` yang parameter `ch` dideklarasikan dengan level akses untuk pengiriman data saja. Channel tersebut hanya bisa digunakan untuk mengirim, contohnya: `ch <- fmt.Sprintf("data %d", i)`.

Dan sebaliknya pada fungsi `printMessage(ch <-chan string)`, channel `ch` hanya bisa digunakan untuk menerima data saja.
