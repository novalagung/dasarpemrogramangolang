# Layout Format String

Di bab-bab sebelumnya kita telah banyak menggunakan layout format string seperti `%s`, `%d`, `%.2f`, dan lainnya; untuk keperluan menampilkan output ke layar ataupun untuk memformat string.

Layout format string digunakan pada konversi data ke string dengan bentuk tertentu. Contohnya pada perubahan nilai data bertipe `double` menjadi string dengan batas desimal 3 digit menggunakan `%.3f`.

Pada bab ini kita akan mempelajari satu per satu layout format string yang tersedia di Golang. Sample data yang digunakan sebagai contoh adalah kode berikut.

```go
type student struct {
    name        string
    height      float64
    age         int32
    isGraduated bool
    hobbies     []string
}

var data = student{
    name:        "wick",
    height:      182.5,
    age:         26,
    isGraduated: false,
    hobbies:     []string{"eating", "sleeping"},
}
```

## Layout Format `%b`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 2 (biner).

```go
fmt.Printf("%t\n", data.age)
// 11010
```

## Layout Format `%c`

Digunakan untuk memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicode-nya.

```go
fmt.Printf("%c\n", 1400)
// ո

fmt.Printf("%c\n", 1235)
// ӓ
```

## Layout Format `%d`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 10 (basis bilangan yang kita gunakan).

```go
fmt.Printf("%t\n", data.age)
// 26
```

## Layout Format `%e` atau `%E`

Digunakan untuk memformat data numerik desimal ke dalam bentuk notasi standar [Scientific notation](https://en.wikipedia.org/wiki/Scientific_notation).

```go
fmt.Printf("%e\n", data.height)
// 1.825000e+02

fmt.Printf("%E\n", data.height)
// 1.825000E+02
```

**1.825000E+02** maksudnya adalah **1.825 x 10^2**, dan hasil operasi tersebut adalah sesuai dengan data asli = **182.5**.

Perbedaan antara `%e` dan `%E` hanya huruf besar kecil karakter `e` pada hasil.

## Layout Format `%f` atau `%F`

`%F` adalah alias dari `%f`. Keduanya memiliki fungsi yang sama.

Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.

```go
fmt.Printf("%f\n", data.height)
// 182.500000

fmt.Printf("%.9f\n", data.height)
// 182.500000000

fmt.Printf("%.2f\n", data.height)
// 182.50

fmt.Printf("%.f\n", data.height)
// 182
```

## Layout Format `%g` atau `%G`

`%G` adalah alias dari `%g`. Keduanya memiliki fungsi yang sama.

Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Lebar kapasitasnya sangat besar, pas digunakan untuk data yang jumlah digit desimalnya cukup banyak.

Bisa dilihat pada kode berikut perbandingan antara `%e`, `%f`, dan `%g`.

```go
fmt.Printf("%e\n", 0.123123123123)
// 1.231231e-01

fmt.Printf("%f\n", 0.123123123123)
// 0.123123

fmt.Printf("%g\n", 0.123123123123)
// 0.123123123123
```

Perbedaan lainnya adalah pada `%g`, lebar digit desimal adalah sesuai dengan datanya, tidak bisa dicustom seperti pada `%f`.

```go
fmt.Printf("%g\n", 0.12)
// 0.12

fmt.Printf("%.5g\n", 0.12)
// 0.12
```

## Layout Format `%o`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 8 (oktal).

```go
fmt.Printf("%o\n", data.age)
// 32
```

## Layout Format `%p`

Digunakan untuk memformat data pointer, mengembalikan alamat pointer referensi variabel-nya.

Alamat pointer dituliskan dalam bentuk numerik berbasis 16 dengan prefix `0x`.

```go
fmt.Printf("%p\n", &data.name)
// 0x2081be0c0
```

## Layout Format `%q`

Digunakan untuk *escape* string. Meskipun string yang dipakai menggunakan literal `\`` akan tetap di-escape.

```go
fmt.Printf("%q\n", `" name \ height "`)
// "\" name \\\\ height \""
```

## Layout Format `%s`

Digunakan untuk memformat data string.

```go
fmt.Printf("%s\n", data.name)
// wick
```

## Layout Format `%t`

Digunakan untuk memformat data boolean, menampilkan nilai bool-nya.

```go
fmt.Printf("%t\n", data.isGraduated)
// false
```

## Layout Format `%T`

Berfungsi untuk mengambil tipe variabel yang akan diformat.

```go
fmt.Printf("%T\n", data.name)
// string

fmt.Printf("%T\n", data.height)
// float64

fmt.Printf("%T\n", data.age)
// int32

fmt.Printf("%T\n", data.isGraduated)
// bool

fmt.Printf("%T\n", data.hobbies)
// []string
```

## Layout Format `%v`

Digunakan untuk memformat data apa saja (termasuk data bertipe `interface{}`). Hasil kembaliannya adalah string nilai data aslinya.

Jika data adalah objek cetakan `struct`, maka akan ditampilkan semua secara property berurutan.

```go
fmt.Printf("%v\n", data)
// {wick 182.5 26 false [eating sleeping]}
```

## Layout Format `%+v`

Digunakan untuk memformat struct, mengembalikan nama tiap property dan nilainya berurutan sesuai dengan struktur struct.

```go
fmt.Printf("%+v\n", data)
// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}
```

## Layout Format `%#v`

Digunakan untuk memformat struct, mengembalikan nama dan nilai tiap property sesuai dengan struktur struct dan juga bagaimana objek tersebut dideklarasikan.

```go
fmt.Printf("%#v\n", data)
// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string{"eating", "sleeping"}}
```

Ketika menampilkan objek yang deklarasinya adalah menggunakan teknik *anonymous struct*, maka akan muncul juga struktur anonymous struct nya.

```go
var data = struct {
    name   string
    height float64
}{
    name:   "wick",
    height: 182.5,
}

fmt.Printf("%#v\n", data)
// struct { name string; height float64 }{name:"wick", height:182.5}
```

Format ini juga bisa digunakan untuk menampilkan tipe data lain, dan akan dimunculkan strukturnya juga.

## Layout Format `%x` atau `%X`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 16 (heksadesimal).

```go
fmt.Printf("%x\n", data.age)
// 1a
```

Jika digunakan pada tipe data string, maka akan mengembalikan kode heksadesimal tiap karakter.

```go
var d = data.name

fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
// 7769636b

fmt.Printf("%x\n", d)
// 7769636b
```

`%x` dan `%X` memiliki fungsi yang sama. Perbedaannya adalah `%X` akan mengembalikan string dalam bentuk *uppercase* atau huruf kapital.

## Layout Format `%%`

Cara untuk menulis karakter `%` pada string format.

```go
fmt.Printf("%%\n")
// %
```
