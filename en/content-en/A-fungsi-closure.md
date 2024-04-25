# A.21. Fungsi Closure

Definisi **Closure** adalah suatu *anonymous function* (atau fungsi tanpa nama) yang disimpan dalam variabel. Dengan adanya closure, kita bisa mendesain beberapa hal diantaranya seperti: membuat fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi. Closure biasa dimanfaatkan untuk membungkus suatu proses yang hanya dijalankan sekali saja atau hanya dipakai pada blok tertentu saja.

## A.21.1. Closure Disimpan Sebagai Variabel

Sebuah fungsi tanpa nama bisa disimpan dalam variabel. Variabel closure memiliki sifat seperti fungsi yang disimpannya.

Di bawah ini adalah contoh program sederhana yang menerapkan closure untuk pencarian nilai terendah dan tertinggi dari data array. Logika pencarian dibungkus dalam closure yang ditampung oleh variabel `getMinMax`.

```go
package main

import "fmt"

func main() {
    var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}
```

Bisa dilihat pada kode di atas bagaimana cara deklarasi closure dan cara pemanggilannya. Sedikit berbeda memang dibanding pembuatan fungsi biasa, pada closure fungsi ditulis tanpa memiliki nama lalu ditampung ke variabel.

```go
var getMinMax = func(n []int) (int, int) {
    // ...
}
```

Cara pemanggilan closure adalah dengan memperlakukan variabel closure seperti fungsi, dituliskan seperti pemanggilan fungsi.

```go
var min, max = getMinMax(numbers)
```

Output program:

![Penerapan closure](images/A_fungsi_closure_1_closure.png)

## A.21.2. Penjelasan tambahan

Berikut merupakan penjelasan tambahan untuk beberapa hal dari kode yang sudah dipraktekan:

#### ◉ Penggunaan Template String `%v`

Template `%v` digunakan untuk menampilkan data tanpa melihat tipe datanya. Jadi bisa digunakan untuk menampilkan data array, int, float, bool, dan lainnya. Bisa dilihat di contoh statement, data bertipe array dan numerik ditampilkan menggunakan `%v`.

```go
fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
```

Template `%v` ini biasa dimanfaatkan untuk menampilkan sebuah data yang tipe nya bisa dinamis atau belum diketahui. Biasa digunakan untuk keperluan debugging, misalnya untuk menampilkan data bertipe `any` atau `interface{}`.

> Pembahasan mengenai tipe data `any` atau `interface{}` ada di chapter [A.27. Interface](/A-interface.html)

#### ◉ Immediately-Invoked Function Expression (IIFE)

Closure jenis IIFE ini eksekusinya adalah langsung saat deklarasi. Teknik ini biasa diterapkan untuk membungkus proses yang hanya dilakukan sekali. IIFE bisa memiliki nilai balik atau bisa juga tidak.

Di bawah ini merupakan contoh sederhana penerapan metode IIFE untuk filtering data array.

```go
package main

import "fmt"

func main() {
    var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

    var newNumbers = func(min int) []int {
        var r []int
        for _, e := range numbers {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3)

    fmt.Println("original number :", numbers)
    fmt.Println("filtered number :", newNumbers)
}
```

Output program:

![Penerapan IIFE](images/A_fungsi_closure_2_iife.png)

Ciri khas dari penulisan IIFE adalah adanya tanda kurung parameter yang ditulis di akhir deklarasi closure. Jika IIFE memiliki parameter, maka argument-nya juga ditulis. Contoh:

```go
var newNumbers = func(min int) []int {
    // ...
}(3)
```

Di contoh sederhana di atas, IIFE menghasilkan nilai balik yang ditampung variabel `newNumber`. Perlu diperhatikan bahwa yang ditampung adalah **nilai kembaliannya** bukan body fungsi atau **closure**-nya.

> Closure bisa juga dengan gaya manifest typing, caranya dengan menuliskan skema closure-nya sebagai tipe data. Contoh:<br /><code>var closure (func (string, int, []string) int)</code><br /><code>closure = func (a string, b int, c []string) int {</code><br /><code>&nbsp;&nbsp;&nbsp;&nbsp;// ..</code><br /><code>}</code>
## A.21.3. Closure Sebagai Nilai Kembalian

Salah satu keunikan lain dari closure adalah: closure bisa dijadikan sebagai nilai balik fungsi. Cukup aneh, tapi pada kondisi tertentu teknik ini sangat berguna. Sebagai contoh, di bawah ini dideklarasikan sebuah fungsi bernama `findMax()` yang salah satu nilai kembaliannya adalah berupa closure.

```go
package main

import "fmt"

func findMax(numbers []int, max int) (int, func() []int) {
    var res []int
    for _, e := range numbers {
        if e <= max {
            res = append(res, e)
        }
    }
    return len(res), func() []int {
        return res
    }
}
```

Nilai kembalian ke-2 pada fungsi di atas adalah closure dengan skema `func() []int`. Bisa dilihat di bagian akhir, ada fungsi tanpa nama yang dikembalikan.

```go
return len(res), func() []int {
    return res
}
```

> Fungsi tanpa nama yang akan dikembalikan boleh disimpan pada variabel terlebih dahulu. Contohnya:<br /><code>var getNumbers = func() []int {</code><br /><code>&nbsp;&nbsp;&nbsp;&nbsp;return res</code><br /><code>}</code><br /><code>return len(res), getNumbers</code>
Tentang fungsi `findMax()` sendiri, fungsi ini dibuat untuk mempermudah pencarian angka-angka yang nilainya di bawah atau sama dengan angka tertentu. Fungsi ini mengembalikan dua buah nilai balik:

- Nilai balik pertama adalah jumlah angkanya.
- Nilai balik kedua berupa closure yang mengembalikan angka-angka yang dicari.

Berikut merupakan contoh implementasi fungsi tersebut:

```go
func main() {
    var max = 3
    var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var howMany, getNumbers = findMax(numbers, max)
    var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers)
    fmt.Printf("find \t: %d\n\n", max)

    fmt.Println("found \t:", howMany)    // 9
    fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}
```

Output program:

![Kombinasi parameter biasa dan variadic](images/A_fungsi_closure_3_combination.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.21-fungsi-closure">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.21...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
