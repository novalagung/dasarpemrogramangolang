# Seleksi Kondisi

Seleksi kondisi digunakan untuk mengontrol alur program. Kalau dianalogikan, fungsinya mirip seperti rambu lalu lintas di jalan raya. Kapan kendaraan diperbolehkan melaju dan kapan harus berhenti, diatur oleh rambu tersebut. Sama seperti pada seleksi kondisi, kapan sebuah blok kode akan dieksekusi juga akan diatur.

Yang dijadikan acuan oleh seleksi kondisi adalah nilai bertipe `bool`, bisa berasa dari variabel, ataupun hasil operasi perbandingan. Nilai tersebut akan menentukan blok kode mana yang akan dieksekusi.

Golang memiliki 2 macam keyword untuk seleksi kondisi, yaitu **if else** dan **switch**. Di bab ini kita akan mempelajarinya satu-persatu.

> Golang tidak mendukung seleksi kondisi menggunakan **ternary**.<br />Statement seperti: `var data = (isExist ? "ada" : "tidak ada")` akan menghasilkan error.

## Seleksi Kondisi Menggunakan Keyword `if`, `else if`, & `else`

Cara penerapan if-else di Golang sama dengan pada bahasa pemrograman lain. Yang membedakan hanya tanda kurung kurawalnya, di Golang tidak perlu ditulis. Kode berikut merupakan contoh penerapan seleksi kondisi if else, dengan jumlah kondisi 4 buah.

```go
var point = 8

if point == 10 {
    fmt.Println("lulus dengan nilai sempurna")
} else if point > 5 {
    fmt.Println("lulus")
} else if point == 4 {
    fmt.Println("hampir lulus")
} else {
    fmt.Printf("tidak lulus. nilai anda %d\n", point)
}
```

Dari ke-empat kondisi di atas, yang terpenuhi adalah `if point > 5` karena nilai variabel `point` memang lebih besar dari `5`. Maka blok kode tepat dibawah kondisi tersebut akan dieksekusi (blok kode ditandai kurung kurawal buka dan tutup), text `"lulus"` akan muncul di console.

![Seleksi kondisi `if` - `else`](images/12_1_if_else.png)

Skema if else Golang sama seperti pada pemrograman umumnya. Yaitu di awal seleksi kondisi menggunakan `if`, dan ketika kondisinya tidak terpenuhi akan menuju ke `else` (jika ada). Ketika ada banyak kondisi, gunakan `else if`.

I> #### Seleksi kondisi dengan blok kode 1 baris
I>
I> Di bahasa pemrograman lain, ketika ada seleksi kondisi yang isi blok-nya hanya 1 baris saja, kurung kurawal boleh tidak dituliskan. Berbeda dengan aturan di Golang, kurung kurawal harus tetap dituliskan meski isinya hanya 1 blok satement.

## Variabel Temporary Pada `if` - `else`

Variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi dimana ia ditempatkan saja. Penggunaan variabel ini membawa beberapa manfaat, antara lain:

 - Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
 - Kode menjadi lebih rapi
 - Ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan di dalam blok masing-masing kondisi.

Berikut merupakan contoh penerapannya.

```go
var point = 8840.0

if percent := point / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
} else if percent >= 70 {
    fmt.Printf("%.1f%s good\n", percent, "%")
} else {
    fmt.Printf("%.1f%s not bad\n", percent, "%")
}
```

Variabel `percent` nilainya didapat dari hasil perhitungan, dan hanya bisa digunakan di deretan blok seleksi kondisi itu saja.

> Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda `:=`. Penggunaan keyword `var` disitu tidak diperbolehkan karena akan menyebabkan error.

## Seleksi Kondisi Menggunakan Keyword `switch`

Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel. Contoh sederhananya seperti penentuan apakah nilai variabel `x` adalah: `1`, `2`, `3`, atau lainnya. Agar lebih jelas, silakan melihat contoh di bawah ini.

```go
var point = 6

switch point {
case 8:
    fmt.Println("perfect")
case 7:
    fmt.Println("awesome")
default:
    fmt.Println("not bad")
}
```

Pada kode di atas, tidak ada kondisi atau `case` yang terpenuhi karena nilai variabel `point` adalah `6`. Ketika hal seperti ini terjadi, blok kondisi `default` akan dipanggil. Bisa dibilang bahwa `default` merupakan `else` dalam sebuah switch.

Perlu diketahui, switch pada pemrograman Golang memiliki perbedaan dibanding bahasa lain. Di Golang, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekkan case selanjutnya, meskipun tidak ada keyword `break` di situ. Konsep ini berkebalikan dengan switch pada umumnya, yang ketika sebuah case terpenuhi, maka akan tetap dilanjut mengecek case selanjutnya kecuali ada keyword `break`.-

## Pemanfaatan 1 `case` Untuk Banyak Kondisi

Satu buah `case` bisa menampung banyak kondisi. Cara untuk menerapkannya yaitu dengan menuliskan nilai pembanding-pembandingnya setelah keyword `case` dipisah tanda koma (`,`). Contoh bisa dilihat pada kode berikut.

```go
var point = 6

switch point {
case 8:
    fmt.Println("perfect")
case 7, 6, 5, 4:
    fmt.Println("awesome")
default:
    fmt.Println("not bad")
}
```

Kondisi `case 7, 6, 5, 4:` akan terpenuhi ketika nilai variabel `point` adalah 7 atau 6 atau 5 atau 4.

## Kurung Kurawal Pada Keyword `case` & `default`

Tanda kurung kurawal atau brackets (`{ }`) bisa diterapkan pada keyword `case` dan `default`. Tanda ini opsional, boleh dipakai boleh tidak. Bagus jika dipakai pada blok kondisi yang didalamnya ada banyak statement, kode akan terlihat lebih rapi dan mudah di-maintain.

Berikut adalah contoh penggunaan brackets dalam switch. Bisa dilihat pada keyword `default` terdapat kurung kurawal yang mengapit 2 statement didalamnya.

```go
var point = 6

switch point {
case 8:
    fmt.Println("perfect")
case 7, 6, 5, 4:
    fmt.Println("awesome")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you can be better!")
    }
}
```

## Switch Dengan Gaya `if` - `else`

Uniknya di Golang, switch bisa digunakan dengan gaya ala if-else. Nilai yang akan dibandingkan tidak dituliskan setelah keyword `switch`, melainkan akan ditulis langsung dalam bentuk perbandingan dalam keyword `case`.

Pada kode di bawah ini, kode program switch di atas diubah ke dalam gaya `if-else`. Variabel `point` dihilangkan dari keyword `switch`, lalu kondisi-kondisinya dituliskan di tiap `case`.

```go
var point = 6

switch {
case point == 8:
    fmt.Println("perfect")
case (point < 8) && (point > 3):
    fmt.Println("awesome")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}
```

## Penggunaan Keyword `fallthrough` Dalam `switch`

Seperti yang kita sudah singgung di atas, bahwa switch pada Golang memiliki beberapa perbedaan dengan bahasa lain. Ketika sebuah `case` terpenuhi, pengecekkan kondisi tidak akan diteruskan ke case-case setelahnya.

Keyword `fallthrough` digunakan untuk memaksa proses pengecekkan diteruskan ke `case` selanjutnya, tanpa melihat case tersebut terpenuhi atau tidak. Contoh berikut merupakan penerapan keyword ini.

```go
var point = 6

switch {
case point == 8:
    fmt.Println("perfect")
case (point < 8) && (point > 3):
    fmt.Println("awesome")
    fallthrough
case point < 5:
    fmt.Println("you need to learn more")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}
```


Setelah pengecekkan `case (point < 8) && (point > 3)` selesai, akan dilanjut ke pengecekkan `case point < 5`, karena ada `fallthrough` di situ.

![Penggunaan `fallthrough` dalam `switch`](images/12_2_fallthrough.png)

## Seleksi Kondisi Bersarang

Seleksi kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi kondisi, yang mungkin juga berada dalam seleksi kondisi, dan seterusnya. *Nested loop* atau seleksi kondisi bersarang bisa dilakukan pada `if` - `else`, `switch`, ataupun kombinasi keduanya. Contohnya:

```go
var point = 10

if point > 7 {
    switch point {
    case 10:
        fmt.Println("perfect!")
    default:
        fmt.Println("nice!")
    }
} else {
    if point == 5 {
        fmt.Println("not bad")
    } else if point == 3 {
        fmt.Println("keep trying")
    } else {
        fmt.Println("you can do it")
        if point == 0 {
            fmt.Println("try harder!")
        }
    }
}
```
