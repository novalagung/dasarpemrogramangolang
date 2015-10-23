# Defer & Exit

**Defer** digunakan untuk mengakhirkan eksekusi sebuah statement. Sedangkan **Exit** digunakan untuk menghentikan program. 2 topik ini sengaja digabung agar hubungan antara keduanya lebih mudah dipahami.

## Penerapan keyword `defer`

Seperti yang sudah dijelaskan secara singkat di atas, bahwa defer digunakan untuk mengakhirkan eksekusi baris kode. Ketika eksekusi sudah sampai pada akhir blok fungsi, statement yang di defer baru akan dijalankan.

Defer bisa ditempatkan di mana saja, awal maupun akhir blok.

```go
package main

import "fmt"

func main() {
    defer fmt.Println("halo")
    fmt.Println("selamat datang")
}
```

Keyword `defer` digunakan untuk mengakhirkan statement. Pada kode di atas, `fmt.Println("halo")` di-defer, hasilnya string `"halo"` akan muncul setelah `"selamat datang"`.

![Penerapan `defer`](images/35_1_defer.png)

Ketika ada banyak statement yang di-defer, maka statement tersebut akan dieksekusi di akhir secara berurutan.

## Penerapan Fungsi `os.Exit()`

Exit digunakan untuk menghentikan program secara paksa pada saat itu juga. Semua statement setelah exit tidak akan di eksekusi, termasuk juga defer.

Fungsi `os.Exit()` berada dalam package `os`. Fungsi ini memiliki sebuah parameter bertipe numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai **exit status** ketika program berhenti.

```go
package main

import "fmt"
import "os"

func main() {
    defer fmt.Println("halo")
    os.Exit(1)
    fmt.Println("selamat datang")
}
```

Meskipun `defer fmt.Println("halo")` ditempatkan sebelum `os.Exit()`, statement tersebut tidak akan dieksekusi, karena di-tengah fungsi program dihentikan secara paksa.

![Penerapan `exit`](images/35_2_exit.png)
