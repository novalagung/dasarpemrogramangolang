# A.43. Konversi Antar Tipe Data

Di beberapa chapter sebelum ini kita telah menerapkan beberapa cara konversi data, contohnya seperti konversi `string` ↔ `int` menggunakan `strconv`, dan `time.Time` ↔ `string`. Pada chapter ini kita akan mempelajarinya lebih detail.

## A.43.1. Konversi Menggunakan `strconv`

Package `strconv` berisi banyak fungsi yang sangat membantu kita untuk melakukan konversi. Berikut merupakan beberapa fungsi yang dalam package tersebut.

#### • Fungsi `strconv.Atoi()`

Fungsi ini digunakan untuk konversi data dari tipe `string` ke `int`. `strconv.Atoi()` menghasilkan 2 buah nilai kembalian, yaitu hasil konversi dan `error` (jika konversi sukses, maka `error` berisi `nil`).

```go
package main

import "fmt"
import "strconv"

func main() {
    var str = "124"
    var num, err = strconv.Atoi(str)

    if err == nil {
        fmt.Println(num) // 124
    }
}
```

#### • Fungsi `strconv.Itoa()`

Merupakan kebalikan dari `strconv.Atoi`, berguna untuk konversi `int` ke `string`.

```go
var num = 124
var str = strconv.Itoa(num)

fmt.Println(str) // "124"
```

#### • Fungsi `strconv.ParseInt()`

Digunakan untuk konversi `string` berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan.

Pada contoh berikut, string `"124"` dikonversi ke tipe numerik dengan ketentuan basis yang digunakan `10` dan lebar datanya mengikuti tipe `int64` (lihat parameter ketiga).

```go
var str = "124"
var num, err = strconv.ParseInt(str, 10, 64)

if err == nil {
    fmt.Println(num) // 124
}
```

Contoh lainnya, string `"1010"` dikonversi ke basis 2 (biner) dengan tipe data hasil adalah `int8`.

```go
var str = "1010"
var num, err = strconv.ParseInt(str, 2, 8)

if err == nil {
    fmt.Println(num) // 10
}
```

#### • Fungsi `strconv.FormatInt()`

Berguna untuk konversi data numerik `int64` ke `string` dengan basis numerik bisa ditentukan sendiri.

```go
var num = int64(24)
var str = strconv.FormatInt(num, 8)

fmt.Println(str) // 30
```

#### • Fungsi `strconv.ParseFloat()`

Digunakan untuk konversi `string` ke numerik desimal dengan lebar data bisa ditentukan.

```go
var str = "24.12"
var num, err = strconv.ParseFloat(str, 32)

if err == nil {
    fmt.Println(num) // 24.1200008392334
}
```

Pada contoh di atas, string `"24.12"` dikonversi ke float dengan lebar tipe data `float32`. Hasil konversi `strconv.ParseFloat` adalah sesuai dengan standar [IEEE Standard for Floating-Point Arithmetic](https://en.wikipedia.org/wiki/IEEE_floating_point).

#### • Fungsi `strconv.FormatFloat()`

Berguna untuk konversi data bertipe `float64` ke `string` dengan format eksponen, lebar digit desimal, dan lebar tipe data bisa ditentukan.

```go
var num = float64(24.12)
var str = strconv.FormatFloat(num, 'f', 6, 64)

fmt.Println(str) // 24.120000
```

Pada kode di atas, Data `24.12` yang bertipe `float64` dikonversi ke string dengan format eksponen `f` atau tanpa eksponen, lebar digit desimal 6 digit, dan lebar tipe data `float64`.

Ada beberapa format eksponen yang bisa digunakan. Detailnya bisa dilihat di tabel berikut.

| Format&nbsp;Eksponen | Penjelasan |
| :-------------: | :-------- |
| `b` | -ddddp±ddd, a, eksponen biner (basis 2) |
| `e` | -d.dddde±dd, a, eksponen desimal (basis 10) |
| `E` | -d.ddddE±dd, a, eksponen desimal (basis 10) |
| `f` | -ddd.dddd, tanpa eksponen |
| `g` | Akan menggunakan format eksponen `e` untuk eksponen besar dan `f` untuk selainnya |
| `G` | Akan menggunakan format eksponen `E` untuk eksponen besar dan `f` untuk selainnya |

#### • Fungsi `strconv.ParseBool()`

Digunakan untuk konversi `string` ke `bool`.

```go
var str = "true"
var bul, err = strconv.ParseBool(str)

if err == nil {
    fmt.Println(bul) // true
}
```

#### • Fungsi `strconv.FormatBool()`

Digunakan untuk konversi `bool` ke `string`.

```go
var bul = true
var str = strconv.FormatBool(bul)

fmt.Println(str) // true
```

## A.43.2. Konversi Data Menggunakan Teknik Casting

Cara penerapannya adalah dengan menggunakan keyword tipe data sebagai nama fungsi, kemudiaan argument pemanggilannya diisi dengan data yang ingin dikonversi tipenya.

```go
// konversi nilai 24 bertipe int ke float64
var a float64 = float64(24)
fmt.Println(a) // 24

// konversi nilai 24.00 bertipe float32 ke int32
var b int32 = int32(24.00)
fmt.Println(b) // 24
```

## A.43.3. Casting `string` ↔ `byte`

String sebenarnya adalah slice/array `byte`. Di Go sebuah karakter biasa (bukan unicode) direpresentasikan oleh sebuah elemen slice byte. Tiap elemen slice berisi data `int` dengan basis desimal, yang merupakan kode ASCII dari karakter dalam string.

Cara mendapatkan slice byte dari sebuah data string adalah dengan meng-casting-nya ke tipe `[]byte`.

```go
var text1 = "halo"
var b = []byte(text1)

fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
// 104 97 108 111
```

Pada contoh di atas, string dalam variabel `text1` dikonversi ke `[]byte`. Tiap elemen slice byte tersebut kemudian ditampilkan satu-per-satu.

Contoh berikut ini merupakan kebalikan dari contoh di atas, data bertipe `[]byte` akan dicari bentuk `string`-nya.

```go
var byte1 = []byte{104, 97, 108, 111}
var s = string(byte1)

fmt.Printf("%s \n", s)
// halo
```

Di contoh ke-2 di-atas, beberapa kode byte dituliskan dalam bentuk slice, ditampung variabel `byte1`. Lalu, nilai variabel tersebut di-cast ke `string`, untuk kemudian ditampilkan.

Selain itu, tiap karakter string juga bisa di-casting ke bentuk `int`, hasilnya adalah sama yaitu data byte dalam bentuk numerik basis desimal, dengan ketentuan literal string yang digunakan adalah tanda petik satu (<code>'</code>).

Juga berlaku sebaliknya, data numerik jika di-casting ke bentuk string dideteksi sebagai kode ASCII dari karakter yang akan dihasilkan.

```go
var c int64 = int64('h')
fmt.Println(c) // 104

var d string = string(104)
fmt.Println(d) // h
```

## A.43.4. Type Assertions Pada Tipe `any` atau Interface Kosong (`interface{}`)

**Type assertions** merupakan teknik untuk mengambil tipe data konkret dari data yang terbungkus dalam `interface{}` atau `any`. Lebih jelasnya silakan cek contoh berikut.

Variabel `data` disiapkan bertipe `map[string]interface{}`, map tersebut berisikan beberapa item dengan tipe data value-nya berbeda satu sama lain, sementara tipe data untuk key-nya sama yaitu `string`.

```go
var data = map[string]interface{}{
    "nama":    "john wick",
    "grade":   2,
    "height":  156.5,
    "isMale":  true,
    "hobbies": []string{"eating", "sleeping"},
}

fmt.Println(data["nama"].(string))
fmt.Println(data["grade"].(int))
fmt.Println(data["height"].(float64))
fmt.Println(data["isMale"].(bool))
fmt.Println(data["hobbies"].([]string))
```

Statement `data["nama"].(string)` maksudnya adalah, nilai `data["nama"]` yang bertipe `interface{}` diambil nilai konkretnya dalam bentuk string `string`.

Pada kode di atas, tidak akan terjadi panic error, karena semua operasi type assertion adalah dilakukan menggunakan tipe data yang sudah sesuai dengan tipe data nilai aslinya. Seperti `data["nama"]` yang merupakan `string` pasti bisa di-asertasi ke tipe `string`.

Coba lakukan asertasi ke tipe yang tidak sesuai dengan tipe nilai aslinya, seperti `data["nama"].(int)`, hasilnya adalah panic error.

Nah, dari penjelasan di atas, terlihat bahwa kita harus tau terlebih dahulu apa tipe data asli dari data yang tersimpan dalam interface. Jika misal tidak tau, maka bisa gunakan teknik di bawah ini untuk pengecekan sukses tidaknya proses asertasi.

Tipe asli data pada variabel `interface{}` bisa diketahui dengan cara meng-casting ke tipe `type`, namun casting ini hanya bisa dilakukan pada `switch`.

```go
for _, val := range data {
    switch val.(type) {
    case string:
        fmt.Println(val.(string))
    case int:
        fmt.Println(val.(int))
    case float64:
        fmt.Println(val.(float64))
    case bool:
        fmt.Println(val.(bool))
    case []string:
        fmt.Println(val.([]string))
    default:
        fmt.Println(val.(int))
    }
}
```

Kombinasi `switch` - `case` bisa dimanfaatkan untuk deteksi tipe konkret data yang bertipe `interface{}`, contoh penerapannya seperti pada kode di atas.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.43-konversi-data">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.43...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
