# Timer

Ada beberapa fungsi dalam package `time` yang memiliki kegunaan sebagai timer. Dengan memanfaatkan fungsi-fungsi tersebut, kita bisa menunda eksekusi sebuah proses dengan durasi waktu tertentu.

## Fungsi `time.Sleep()`

Fungsi ini digunakan untuk menghentikan program sejenak. `time.Sleep()` bersifat **blocking**, sehingga statement dibawahnya tidak akan dieksekusi sampai waktu pemberhentian usai. Contoh sederhana penerapan `time.Sleep()`:

```go
import "fmt"
import "time"

func main () {
    fmt.Println("start")
    time.Sleep(time.Second * 4)
    fmt.Println("after 4 seconds")
}
```

Tulisan `"start"` muncul, lalu 4 detik kemudian tulisan `"after 4 seconds"` muncul.

## Fungsi `time.NewTimer()`

Fungsi ini sedikit berbeda dengan `time.Sleep()`. Fungsi `time.NewTimer()` mengembalikan sebuah objek `*time.Timer` yang memiliki method `C`. Method ini mengembalikan sebuah channel dan akan dieksekusi dalam waktu yang sudah ditentukan. Contoh penerapannya bisa dilihat pada kode berikut.

```go
var timer = time.NewTimer(4 * time.Second)
fmt.Println("start")
<-timer.C
fmt.Println("finish")
```

Tulisan `"finish"` akan muncul setelah delay **4 detik**.

## Fungsi `time.AfterFunc()`

Fungsi `time.AfterFunc()` memiliki 2 parameter. Parameter pertama adalah durasi timer, dan parameter kedua adalah *callback* nya. Callback tersebut akan dieksekusi jika waktu sudah memenuhi durasi timer.

```go
var ch = make(chan bool)

time.AfterFunc(4*time.Second, func() {
    fmt.Println("expired")
    ch <- true
})

fmt.Println("start")
<-ch
fmt.Println("finish")
```

Tulisan `"start"` akan muncul di awal. Diikuti 4 detik kemudian tulisan `"expired"`.

Didalam callback terdapat proses transfer data lewat channel, mengakibatkan tulisan `"finish"` akan muncul tepat setelah tulisan `"expired"` muncul.

Beberapa hal yang perlu diketahui dalam menggunakan fungsi ini:

 - Jika tidak ada serah terima data lewat channel, maka eksekusi `time.AfterFunc()` adalah asynchronous dan tidak blocking.
 - Jika ada serah terima data lewat channel, maka fungsi akan tetap berjalan asynchronous dan tidak blocking hingga baris kode dimana penerimaan data channel dilakukan.

## Fungsi `time.After()`

Kegunaan fungsi ini mirip seperti `time.Sleep()`. Perbedaannya adalah, fungsi `timer.After()` akan mengembalikan data channel, sehingga perlu menggunakan tanda `<-` dalam penerapannya.

```go
<-time.After(4 * time.Second)
fmt.Println("expired")
```

Tulisan `"expired"` akan muncul setelah 4 detik.

## Kombinasi Timer & Goroutine

Berikut merupakan contoh penerapan timer dan goroutine. Program di bawah ini adalah program tanya-jawab sederhana. User harus menginputkan jawaban dalam waktu tidak lebih dari 5 detik. Jika lebih dari waktu tersebut belum ada jawabamn, maka akan muncul pesan *time out*.

Pertama, siapkan import package yang diperlukan.

```go
import "fmt"
import "os"
import "time"
```

Buat fungsi `timer()`, yang nantinya akan dieksekusi sebagai goroutine. Di dalam fungsi ini akan ada proses pengiriman data lewat channel `ch` ketika waktu sudah mencapai `timeout`.

```go
func timer(timeout int, ch chan<- bool) {
    time.AfterFunc(time.Duration(timeout)*time.Second, func() {
        ch <- true
    })
}
```

Siapkan juga fungsi `watcher()`. Fungsi ini juga akan dieksekusi sebagai goroutine. Tugasnya cukup sederhana, yaitu ketika sebuah data diterima dari channel `ch` maka akan ditampilkan tulisan penanda waktu habis.

```go
func watcher(timeout int, ch <-chan bool) {
    <-ch
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
    os.Exit(0)
}
```

Terakhir, buat implementasi di fungsi `main`.

```go
func main() {
    var timeout = 5
    var ch = make(chan bool)

    go timer(timeout, ch)
    go watcher(timeout, ch)

    var input string
    fmt.Print("what is 725/25 ? ")
    fmt.Scan(&input)

    if input == "29" {
        fmt.Println("the answer is right!")
    } else {
        fmt.Println("the answer is wrong!")
    }
}
```

Ketika user tidak menginputkan apa-apa dalam kurun waktu 5 detik, maka akan muncul pesan timeout, lalu program dihentikan.

![Penerapan timer dalam goroutine](images/39_1_timer.png)
