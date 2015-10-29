# Konversi Data

Di bab-bab sebelumnya kita sudah mengaplikasikan beberapa cara konversi data, contohnya seperti konversi `string` ↔ `int` menggunakan `strconv`, dan `time.Time` ↔ `string`. Di bab ini kita akan belajar lebih banyak.

## Konversi Menggunakan `strconv`

`strconv` berisikan banyak fungsi yang sangat membantu untuk keperluan konversi data. Berikut merupakan beberapa fungsi dalam package tersebut yang bisa dimanfaatkan.

### Fungsi `strconv.Atoi()`

Fungsi ini digunakan untuk konversi data dari tipe `string` ke `int`. Mengembalikan 2 buah nilai balik, yaitu hasil konversi dan `error` (jika konversi sukses, maka `error` akan berisi `nil`).

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

### Fungsi `strconv.Itoa()`

Merupakan kebalikan dari `strconv.Atoi`, berguna untuk konversi `int` ke `string`.

```go
var num = 124
var str = strconv.Itoa(num)

fmt.Println(str) // "124"
```

### Fungsi `strconv.ParseInt()`

Digunakan untuk konversi `string` berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan.

Pada contoh berikut, string `"124"` ditentukan basis numeriknya 10, akan dikonversi ke jenis tipe data `int64`.

```go
var str = "124"
var num, err = strconv.ParseInt(str, 10, 64)

if err == nil {
    fmt.Println(num) // 124
}
```

Contoh lainnya, string `"1010"` ditentukan basis numeriknya 2 (biner), akan dikonversi ke jenis tipe data `int8`.

```go
var str = "1010"
var num, err = strconv.ParseInt(str, 2, 8)

if err == nil {
    fmt.Println(num) // 10
}
```

### Fungsi `strconv.FormatInt()`

Berguna untuk konversi data numerik `int64` ke `string` dengan basis numerik bisa ditentukan sendiri.

```go
var num = int64(24)
var str = strconv.FormatInt(num, 8)

fmt.Println(str) // 30
```

### Fungsi `strconv.ParseFloat()`

Digunakan untuk konversi `string` ke numerik desimal dengan lebar data bisa ditentukan.

```go
var str = "24.12"
var num, err = strconv.ParseFloat(str, 32)

if err == nil {
    fmt.Println(num) // 24.1200008392334
}
```

Pada contoh di atas, string `"24.12"` dikonversi ke float dengan lebar `float32`. Hasil konversi `strconv.ParseFloat` disesuaikan dengan standar [IEEE Standard for Floating-Point Arithmetic](https://en.wikipedia.org/wiki/IEEE_floating_point).

### Fungsi `strconv.FormatFloat()`

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

### Fungsi `strconv.ParseBool()`

Digunakan untuk konversi `string` ke `bool`.

```go
var str = "true"
var bul, err = strconv.ParseBool(str)

if err == nil {
    fmt.Println(bul) // true
}
```

### Fungsi `strconv.FormatBool()`

Digunakan untuk konversi `bool` ke `string`.

```go
var bul = true
var str = strconv.FormatBool(bul)

fmt.Println(str) // 124
```

## Konversi Data Menggunakan Casting

Keyword tipe data bisa digunakan untuk casting. Cara penggunaannya adalah dengan memanggilnya sebagai fungsi dan menyisipkan data yang akan dikonversi sebagai parameter.

```go
var a float64 = float64(24)
fmt.Println(a) // 24

var b int32 = int32(24.00)
fmt.Println(b) // 24
```

## Casting `string` ↔ `byte`

String sebenarnya adalah slice/array `byte`. Di Golang sebuah karakter biasa (bukan unicode) direpresentasikan oleh sebuah elemen slice byte. Nilai slice tersebut adalah data `int` yang (default-nya) ber-basis desimal, yang merupakan kode ASCII dari karakter biasa tersebut.

Cara mendapatkan slice byte dari sebuah data string adalah dengan meng-casting-nya ke tipe `[]byte`. Tiap elemen `byte` isinya adalah data numerik dengan basis desimal.

```go
var text1 = "halo"
var b = []byte(text1)

fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
// 104 97 108 111
```

Pada contoh di atas, string dalam variabel `text1` dikonversi ke `[]byte`. Tiap elemen slice byte tersebut kemudian ditampilkan satu-per-satu.

Contoh selanjutnya dibawah ini merupakan kebalikan dari contoh di atas, sebuah `[]byte` akan dicari bentuk `string`-nya.

```go
var byte1 = []byte{104, 97, 108, 111}
var s = string(byte1)

fmt.Printf("%s \n", s)
// halo
```

Beberapa kode byte string saya tuliskan sebagai dalam sebuah slice, yang ditampung oleh variabel `byte1`. Lalu, nilai variabel tersebut di-cast ke `string`, untuk kemudian ditampilkan.

Selain itu, tiap karakter string juga bisa di-casting ke bentuk `int`, hasilnya adalah sama yaitu data byte dalam bentuk numerik basis desimal, dengan ketentuan literal string yang digunakan adalah tanda petik satu (<code>'</code>).

Juga berlaku sebaliknya, data numerik jika di-casting ke bentuk string dideteksi sebagai kode byte dari karakter yang akan dihasilkan.

```go
var c int64 = int64('h')
fmt.Println(c) // 104

var d string = string(104)
fmt.Println(d) // h
```

## Konversi Data `interface{}` Menggunakan Teknik Type Assertions

**Type assertions** merupakan teknik casting data `interface{}` ke segala jenis tipe (dengan syarat data tersebut memang bisa di-casting).

Berikut merupakan contoh penerapannya. Disiapkan variabel `data` bertipe `map[string]interface{}` dengan value berbeda beda tipe datanya.

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

Statement `data["nama"].(string)` maksudnya adalah, nilai `data["nama"]` dicasting sebagai `string`.

Tipe asli data pada variabel `interface{}` bisa diketahui dengan cara meng-casting `interface{}` ke tipe `type`. Namun casting ini hanya bisa dilakukan pada `switch`.

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

Kombinasi `switch` - `case` bisa dimanfaatkan untuk deteksi tipe asli sebuah data bertipe `interface{}`, contoh penerapannya seperti pada kode di atas.
