# A.28. Any / interface{} / Interface Kosong

Interface kosong atau *empty interface* yang dinotasikan dengan `interface{}` atau `any`, merupakan tipe data yang sangat spesial karena variabel bertipe ini bisa menampung segala jenis data, baik itu numerik, string, bahkan array, pointer, apapun.

> Dalam konsep pemrograman umum, konsep variabel yang bisa menampung banyak jenis tipe data disebut dengan **dynamic typing**.

## A.28.1. Penggunaan `any` / `interface{}`

`any` atau `interface{}` merupakan tipe data, sehingga cara penggunaannya sama seperti tipe data pada umumnya, perbedaannya pada variabel bertipe ini nilainya bisa diisi dengan apapun. Contoh:

```go
package main

import "fmt"

func main() {
    var secret interface{}

    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "mango", "banana"}
    fmt.Println(secret)

    secret = 12.4
    fmt.Println(secret)
}
```

Keyword `interface` seperti yang kita tau, digunakan untuk pembuatan interface. Tetapi ketika ditambahkan kurung kurawal (`{}`) di belakang-nya (menjadi `interface{}`), maka kegunaannya akan berubah, yaitu sebagai tipe data.

![Segala jenis data bisa ditampung <code>interface{}</code>](images/A_interface_kosong_1_empty_interface.png)

Agar tidak bingung, coba perhatikan kode berikut.

```go
var data map[string]interface{}

data = map[string]interface{}{
    "name":      "ethan hunt",
    "grade":     2,
    "breakfast": []string{"apple", "mango", "banana"},
}
```

Pada kode di atas, disiapkan variabel `data` dengan tipe `map[string]interface{}`, yaitu sebuah koleksi dengan key bertipe `string` dan nilai bertipe interface kosong `interface{}`.

Kemudian variabel tersebut di-inisialisasi, ditambahkan lagi kurung kurawal setelah keyword deklarasi untuk kebutuhan pengisian data, `map[string]interface{}{ /* data */ }`.

Dari situ terlihat bahwa `interface{}` bukanlah sebuah objek, melainkan tipe data.

## A.28.2. Type Alias `any`

Tipe `any` merupakan alias dari `interface{}`, keduanya adalah sama.

```go
var data map[string]any

data = map[string]any{
    "name":      "ethan hunt",
    "grade":     2,
    "breakfast": []string{"apple", "mango", "banana"},
}
```

## A.28.3. Casting Variabel Any / Interface Kosong

Variabel bertipe `interface{}` bisa ditampilkan ke layar sebagai `string` dengan memanfaatkan fungsi print, seperti `fmt.Println()`. Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah nilai asli, melainkan bentuk text dari nilai aslinya.

Hal ini penting diketahui, karena untuk melakukan operasi yang membutuhkan nilai asli pada variabel yang bertipe `interface{}`, diperlukan casting ke tipe aslinya. Contoh seperti pada kode berikut.

```go
package main

import "fmt"
import "strings"

func main() {
    var secret interface{}

    secret = 2
    var number = secret.(int) * 10
    fmt.Println(secret, "multiplied by 10 is :", number)

    secret = []string{"apple", "mango", "banana"}
    var fruits = strings.Join(secret.([]string), ", ")
    fmt.Println(fruits, "is my favorite fruits")
}
```

Pertama, variabel `secret` menampung nilai bertipe numerik. Ada kebutuhan untuk mengalikan nilai yang ditampung variabel tersebut dengan angka `10`. Maka perlu dilakukan casting ke tipe aslinya, yaitu `int`, setelahnya barulah nilai bisa dioperasikan, yaitu `secret.(int) * 10`.

Pada contoh kedua, `secret` berisikan slice string. Kita memerlukan string tersebut untuk digabungkan dengan pemisah tanda koma. Maka perlu di-casting ke `[]string` terlebih dahulu sebelum bisa digunakan di `strings.Join()`, contohnya pada `strings.Join(secret.([]string), ", ")`.

![Casting pada variabel bertipe <code>interface{}</code>](images/A_interface_kosong_2_interface_casting.png)

Teknik casting pada `any` disebut dengan **type assertions**.

## A.28.4. Casting Variabel Interface Kosong Ke Objek Pointer

Variabel `interface{}` bisa menyimpan data apa saja, termasuk data objek, pointer, ataupun gabungan keduanya. Di bawah ini merupakan contoh penerapan interface untuk menampung data objek pointer.

```go
type person struct {
    name string
    age  int
}

var secret interface{} = &person{name: "wick", age: 27}
var name = secret.(*person).name
fmt.Println(name)
```

Variabel `secret` dideklarasikan bertipe `interface{}` menampung referensi objek cetakan struct `person`. Cara casting dari `interface{}` ke struct pointer adalah dengan menuliskan nama struct-nya dan ditambahkan tanda asterisk (`*`) di awal, contohnya seperti `secret.(*person)`. Setelah itu barulah nilai asli bisa diakses.

![Casting <code>interface{}</code> ke variabel objek](images/A_interface_kosong_3_interface_pointer.png)

## A.28.5. Kombinasi Slice, `map`, dan `interface{}`

Tipe `[]map[string]interface{}` adalah salah satu tipe yang paling sering digunakan untuk menyimpan sekumpulan data berbasis *key-value*. Tipe tersebut merupakan alternatif dari slice struct.

Pada contoh berikut, variabel `person` dideklarasikan berisi data slice `map` berisikan 2 item dengan key adalah `name` dan `age`.

```go
var person = []map[string]interface{}{
    {"name": "Wick", "age": 23},
    {"name": "Ethan", "age": 23},
    {"name": "Bourne", "age": 22},
}

for _, each := range person {
    fmt.Println(each["name"], "age is", each["age"])
}
```

Dengan memanfaatkan slice dan `interface{}`, kita bisa membuat data array yang isinya adalah bisa apa saja. Silakan perhatikan contoh berikut.

```go
var fruits = []interface{}{
    map[string]interface{}{"name": "strawberry", "total": 10},
    []string{"mango", "pineapple", "papaya"},
    "orange",
}

for _, each := range fruits {
    fmt.Println(each)
}
```

## A.28.6. Interface `nil` vs Pointer `nil`

Ada kode yang sering sekali membingungkan pemula, yaitu: variabel interface yang menyimpan pointer `nil` **tidak sama** dengan interface `nil` itu sendiri.

```go
package main

import "fmt"

type MyError struct {
    msg string
}

func (e *MyError) Error() string {
    return e.msg
}

func getError(fail bool) error {
    var err *MyError // pointer nil
    if fail {
        err = &MyError{"something went wrong"}
    }
    return err // mengembalikan interface yang menyimpan (*MyError)(nil)
}

func main() {
    err := getError(false)
    if err != nil {
        fmt.Println("error:", err) // baris ini TETAP dieksekusi!
    } else {
        fmt.Println("no error")
    }
}
```

Pada contoh di atas, `getError(false)` mengembalikan variabel `err` bertipe `*MyError` yang nilainya `nil`. Namun ketika dikembalikan sebagai `error` (interface), interface tersebut **tidak nil** karena interface menyimpan informasi tipe (`*MyError`) meskipun nilainya `nil`. Akibatnya pengecekan `err != nil` bernilai `true`.

Cara yang benar agar fungsi benar-benar mengembalikan interface nil:

```go
func getError(fail bool) error {
    if fail {
        return &MyError{"something went wrong"}
    }
    return nil // kembalikan nil langsung, bukan variabel typed nil
}
```

#### ◉ Kenapa Bisa Terjadi?

Di balik layar, sebuah nilai interface sebenarnya menyimpan dua hal sekaligus:

| Bagian    | Isi                                   |
| --------- | ------------------------------------- |
| **Tipe**  | Tipe konkret dari nilai yang disimpan |
| **Nilai** | Nilai itu sendiri                     |

Interface baru dianggap `nil` kalau **kedua bagian** tersebut kosong. Kalau salah satu sudah terisi, interface tidak lagi `nil`.

Nah, ketika kita menulis `return err` di mana `err` bertipe `*MyError` dan nilainya `nil`, Go tetap menyimpan informasi tipenya (`*MyError`) ke dalam interface. Bagian tipe sudah terisi, jadi interface tidak `nil`, meskipun nilainya sendiri memang `nil`.

Sebaliknya, `return nil` langsung mengembalikan interface yang benar-benar kosong, tanpa tipe dan tanpa nilai.

Aturan praktisnya: **kalau fungsi mengembalikan tipe interface, selalu kembalikan `nil` secara langsung, bukan lewat variabel bertipe konkret.**

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.28-interface-kosong">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.28...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
