# A.18. Fungsi

Fungsi merupakan aspek penting dalam pemrograman. Definisi fungsi sendiri adalah sekumpulan blok kode yang dibungkus dengan nama tertentu. Penerapan fungsi yang tepat akan menjadikan kode lebih modular dan juga *dry* (kependekan dari *don't repeat yourself*), tak perlu menuliskan banyak kode yang kegunaannya berkali-kali, cukup sekali saja lalu panggil sesuai kebutuhan.

Pada chapter ini kita akan belajar tentang penggunaan fungsi di Go.

## A.18.1. Penerapan Fungsi

Sebenarnya tanpa sadar, kita sudah menerapkan fungsi pada pembahasan-pembahasan sebelum ini, yaitu pada fungsi `main`. Fungsi `main` merupakan fungsi yang paling utama pada program Go.

Cara membuat fungsi cukup mudah, yaitu dengan menuliskan keyword `func`, diikuti setelahnya nama fungsi, kurung yang berisikan parameter, dan kurung kurawal untuk membungkus blok kode.

Parameter sendiri adalah variabel yang disisipkan pada saat pemanggilan fungsi.

Silakan lihat dan praktekan kode tentang implementasi fungsi berikut.

```go
package main

import "fmt"
import "strings"

func main() {
    var names = []string{"John", "Wick"}
    printMessage("halo", names)
}

func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}
```

Pada kode di atas, sebuah fungsi baru dibuat dengan nama `printMessage` memiliki 2 buah parameter yaitu string `message` dan slice string `arr`.

Fungsi tersebut dipanggil dalam `main`, dengan disisipkan 2 buah data sebagai parameter, data pertama adalah string `"hallo"` yang ditampung parameter `message`, dan parameter ke 2 adalah slice string `names` yang nilainya ditampung oleh parameter `arr`.

Di dalam `printMessage`, nilai `arr` yang merupakan slice string digabungkan menjadi sebuah string dengan pembatas adalah karakter **spasi**. Penggabungan slice dapat dilakukan dengan memanfaatkan fungsi `strings.Join()` (berada di dalam package `strings`).

![Contoh penggunaan fungsi](images/A_fungsi_1_function.png)

## A.18.2. Fungsi Dengan Return Value / Nilai Balik

Sebuah fungsi bisa dirancang tidak mengembalikan nilai balik (*void*), atau bisa mengembalikan suatu nilai. Fungsi yang memiliki nilai kembalian, harus ditentukan tipe data nilai baliknya pada saat deklarasi.

Program berikut merupakan contoh penerapan fungsi yang memiliki return value.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
}

func randomWithRange(min, max int) int {
    var value = randomizer.Int()%(max-min+1) + min
    return value
}
```

Fungsi `randomWithRange` bertugas untuk *generate* angka acak sesuai dengan range yang ditentukan, yang kemudian angka tersebut dijadikan nilai kembalian fungsi.

![Fungsi dengan nilai balik](images/A_fungsi_2_function_return_type.png)

Cara menentukan tipe data nilai balik fungsi adalah dengan menuliskan tipe data yang diinginkan setelah kurung parameter. Bisa dilihat pada kode di atas, bahwa `int` merupakan tipe data nilai balik fungsi `randomWithRange`.

```go
func randomWithRange(min, max int) int
```

Sedangkan cara untuk mengembalikan nilai itu sendiri adalah dengan menggunakan keyword `return` diikuti data yang ingin dikembalikan. Pada contoh di atas, `return value` artinya nilai variabel `value` dijadikan nilai kembalian fungsi.

Eksekusi keyword `return` akan menjadikan proses dalam blok fungsi berhenti pada saat itu juga. Semua statement setelah keyword tersebut tidak akan dieksekusi.

---

Dari kode di atas mungkin ada beberapa hal yang belum pernah kita lakukan pada pembahasan-pembahasan sebelumnya, kita akan bahas satu-persatu.

## A.18.3. Penggunaan Fungsi `rand.New()`

Fungsi ini digunakan untuk membuat object randomizer, yang dari object tersebut nilai random/acak bisa di-generate. Dalam penerapannya, fungsi `rand.New()` membutuhkan argument yaitu random source seed, yang bisa kita buat lewat statement `rand.NewSource(time.Now().Unix())`.

```go
var randomizer = rand.New(rand.NewSource(time.Now().Unix()))
```

> Dalam penggunaan fungsi `rand.NewSource`, argument bisa diisi dengan nilai apapun, salah satunya adalah `time.Now().Unix()`.
> 
> Lebih detailnya mengenai random dibahas pada chapter [A.39. Random](A-random.html).

Fungsi `rand.New()` berada dalam package `math/rand`, yang harus di-import terlebih dahulu sebelum bisa dimanfaatkan. Package `time` juga perlu di-import karena kita menggunakan fungsi `(time.Now().Unix())` di situ.

## A.18.4. Import Banyak Package

Penulisan keyword `import` untuk banyak package bisa dilakukan dengan dua cara, dengan menuliskannya di tiap package, atau cukup sekali saja, bebas.

```go
import "fmt"
import "math/rand"
import "time"

// atau

import (
    "fmt"
    "math/rand"
    "time"
)
```

## A.18.5. Deklarasi Parameter Bertipe Data Sama

Khusus untuk fungsi yang tipe data parameternya sama, bisa ditulis dengan gaya yang unik. Tipe datanya dituliskan cukup sekali saja di akhir. Contohnya bisa dilihat pada kode berikut.

```go
func nameOfFunc(paramA type, paramB type, paramC type) returnType
func nameOfFunc(paramA, paramB, paramC type) returnType

func randomWithRange(min int, max int) int
func randomWithRange(min, max int) int
```

## A.18.6. Penggunaan Keyword `return` Untuk Menghentikan Proses Dalam Fungsi

Selain sebagai penanda nilai balik, keyword `return` juga bisa dimanfaatkan untuk menghentikan proses dalam blok fungsi di mana ia dipakai. Contohnya bisa dilihat pada kode berikut.

```go
package main

import "fmt"

func main() {
    divideNumber(10, 2)
    divideNumber(4, 0)
    divideNumber(8, -4)
}

func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}
```

Fungsi `divideNumber` dirancang tidak memiliki nilai balik. Fungsi ini dibuat untuk membungkus proses pembagian 2 bilangan, lalu menampilkan hasilnya.

Di dalamnya terdapat proses validasi nilai variabel pembagi, jika nilainya adalah 0, maka akan ditampilkan pesan bahwa pembagian tidak bisa dilakukan, lalu proses dihentikan pada saat itu juga (dengan memanfaatkan keyword `return`). Jika nilai pembagi valid, maka proses pembagian diteruskan.

![Keyword return menjadikan proses dalam fungsi berhenti](images/A_fungsi_3_function_return_as_break.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.18-fungsi">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.18...</a>
</div>

---

<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>
