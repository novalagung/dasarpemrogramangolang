# A.44. Fungsi String

Go menyediakan package `strings`, isinya banyak fungsi untuk keperluan pengolahan data string. Chapter ini berisi pembahasan mengenai penggunaan fungsi yang ada di dalam package tersebut.

## A.44.1. Fungsi `strings.Contains()`

Dipakai untuk deteksi apakah string (parameter kedua) merupakan bagian dari string lain (parameter pertama). Nilai kembaliannya berupa `bool`.

```go
package main

import "fmt"
import "strings"

func main() {
    var isExists = strings.Contains("john wick", "wick")
    fmt.Println(isExists)
}
```

Variabel `isExists` akan bernilai `true`, karena string `"wick"` merupakan bagian dari `"john wick"`.

## A.44.2. Fungsi `strings.HasPrefix()`

Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).

```go
var isPrefix1 = strings.HasPrefix("john wick", "jo")
fmt.Println(isPrefix1) // true

var isPrefix2 = strings.HasPrefix("john wick", "wi")
fmt.Println(isPrefix2) // false
```

## A.44.3. Fungsi `strings.HasSuffix()`

Digunakan untuk deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua).

```go
var isSuffix1 = strings.HasSuffix("john wick", "ic")
fmt.Println(isSuffix1) // false

var isSuffix2 = strings.HasSuffix("john wick", "ck")
fmt.Println(isSuffix2) // true
```

## A.44.4. Fungsi `strings.Count()`

Memiliki kegunaan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama). Nilai kembalian fungsi ini adalah jumlah karakternya.

```go
var howMany = strings.Count("ethan hunt", "t")
fmt.Println(howMany) // 2
```

Nilai yang dikembalikan `2`, karena pada string `"ethan hunt"` terdapat dua buah karakter `"t"`.

## A.44.5. Fungsi `strings.Index()`

Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).

```go
var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2
```

String `"ha"` berada pada posisi ke `2` dalam string `"ethan hunt"` (indeks dimulai dari 0). Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:

```go
var index2 = strings.Index("ethan hunt", "n")
fmt.Println(index2) // 4
```

String `"n"` berada pada indeks `4` dan `8`. Yang dikembalikan adalah yang paling kiri (paling kecil), yaitu `4`.

## A.44.6. Fungsi `strings.Replace()`

Fungsi ini digunakan untuk replace atau mengganti bagian dari string dengan string tertentu. Jumlah substring yang di-replace bisa ditentukan, apakah hanya 1 string pertama, 2 string, atau seluruhnya.

```go
var text = "banana"
var find = "a"
var replaceWith = "o"

var newText1 = strings.Replace(text, find, replaceWith, 1)
fmt.Println(newText1) // "bonana"

var newText2 = strings.Replace(text, find, replaceWith, 2)
fmt.Println(newText2) // "bonona"

var newText3 = strings.Replace(text, find, replaceWith, -1)
fmt.Println(newText3) // "bonono"
```

Penjelasan:

 1. Pada contoh di atas, substring `"a"` pada string `"banana"` akan di-replace dengan string `"o"`.
 2. Pada `newText1`, hanya 1 huruf `o` saja yang tereplace karena maksimal substring yang ingin di-replace ditentukan 1.
 3. Angka `-1` akan menjadikan proses replace berlaku pada semua substring. Contoh bisa dilihat pada `newText3`.

## A.44.7. Fungsi `strings.Repeat()`

Digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).

```go
var str = strings.Repeat("na", 4)
fmt.Println(str) // "nananana"
```

Pada contoh di atas, string `"na"` diulang sebanyak 4 kali. Hasilnya adalah: `"nananana"`

## A.44.8. Fungsi `strings.Split()`

Digunakan untuk memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya berupa slice string.

```go
var string1 = strings.Split("the dark knight", " ")
fmt.Println(string1) // output: ["the", "dark", "knight"]

var string2 = strings.Split("batman", "")
fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]
```

String `"the dark knight"` dipisah oleh karakter spasi `" "`, hasilnya kemudian ditampung oleh `string1`.

Untuk memisah string menjadi slice tiap 1 string, gunakan pemisah string kosong `""`. Bisa dilihat contohnya pada variabel `string2`.

## A.44.9. Fungsi `strings.Join()`

Memiliki kegunaan berkebalikan dengan `strings.Split()`. Digunakan untuk menggabungkan slice string (parameter pertama) menjadi sebuah string dengan pemisah tertentu (parameter kedua).

```go
var data = []string{"banana", "papaya", "tomato"}
var str = strings.Join(data, "-")
fmt.Println(str) // "banana-papaya-tomato"
```

Slice `data` digabungkan menjadi satu dengan pemisah tanda *dash* (`-`).

## A.44.10. Fungsi `strings.ToLower()`

Mengubah huruf-huruf string menjadi huruf kecil.

```go
var str = strings.ToLower("aLAy")
fmt.Println(str) // "alay"
```

## A.44.11. Fungsi `strings.ToUpper()`

Mengubah huruf-huruf string menjadi huruf besar.

```go
var str = strings.ToUpper("eat!")
fmt.Println(str) // "EAT!"
```

## A.44.12. Fungsi `strings.ReplaceAll()` (Go 1.12+)

Tersedia sejak Go 1.12, fungsi `strings.ReplaceAll()` merupakan alternatif yang lebih singkat dari `strings.Replace(..., -1)`. Keduanya menghasilkan output yang sama, namun `strings.ReplaceAll()` tidak memerlukan parameter jumlah penggantian.

```go
var text = "banana"
var find = "a"
var replaceWith = "o"

var result = strings.ReplaceAll(text, find, replaceWith)
fmt.Println(result) // "bonono"
```

`strings.ReplaceAll(text, find, replaceWith)` ekuivalen dengan `strings.Replace(text, find, replaceWith, -1)`.

## A.44.13. Fungsi `strings.Cut()` (Go 1.18+)

Sejak Go 1.18, tersedia fungsi `strings.Cut()` untuk memotong string menjadi dua bagian berdasarkan pemisah tertentu. Fungsi ini mengembalikan tiga nilai: bagian sebelum pemisah, bagian setelah pemisah, dan `bool` yang menandakan apakah pemisah ditemukan atau tidak.

```go
before, after, found := strings.Cut("user:password", ":")
fmt.Println(before) // user
fmt.Println(after)  // password
fmt.Println(found)  // true

before2, after2, found2 := strings.Cut("nocohere", ":")
fmt.Println(before2) // nocohere
fmt.Println(after2)  // (string kosong)
fmt.Println(found2)  // false
```

`strings.Cut()` berguna untuk parsing string yang memiliki format `key:value` atau sejenisnya, dan lebih ekspresif dibanding kombinasi `strings.Index()` + slicing manual.

## A.44.14. Fungsi `strings.CutPrefix()` dan `strings.CutSuffix()` (Go 1.20+)

Sejak Go 1.20, tersedia dua fungsi baru untuk memotong prefix dan suffix dari sebuah string.

`strings.CutPrefix(s, prefix)` mengembalikan string setelah prefix dipotong dan `bool` apakah prefix ditemukan. `strings.CutSuffix(s, suffix)` bekerja dengan cara yang sama namun untuk suffix.

```go
after, found := strings.CutPrefix("foobar", "foo")
fmt.Println(after) // bar
fmt.Println(found) // true

before, found2 := strings.CutSuffix("foobar", "bar")
fmt.Println(before) // foo
fmt.Println(found2) // true
```

Berbeda dengan `strings.TrimPrefix()` dan `strings.TrimSuffix()`, fungsi `Cut*` juga memberi tahu apakah prefix/suffix benar-benar ditemukan lewat nilai `bool`-nya.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.44-fungsi-string">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.44...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
