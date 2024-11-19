# A.11. Konstanta

Konstanta adalah jenis variabel yang nilainya tidak bisa diubah setelah dideklarasikan. Inisialisasi nilai konstanta hanya dilakukan sekali saja di awal, setelah itu variabel tidak bisa diubah nilainya.

## A.11.1. Penggunaan Konstanta

Data seperti **pi** (22/7), kecepatan cahaya (299.792.458 m/s), adalah contoh data yang tepat untuk dideklarasikan sebagai konstanta (daripada variabel), karena nilainya sudah pasti dan tidak akan berubah.

Cara penerapan konstanta sama seperti deklarasi variabel biasa, perbedaannya ada pada keyword yang digunakan, yaitu `const` (bukan `var`).

```go
const firstName string = "john"
fmt.Print("halo ", firstName, "!\n")
```

Teknik type inference bisa diterapkan pada konstanta, caranya cukup dengan menghilangkan tipe data pada saat deklarasi.

```go
const lastName = "wick"
fmt.Print("nice to meet you ", lastName, "!\n")
```

#### ◉ Penggunaan Fungsi `fmt.Print()`

Fungsi ini memiliki peran yang sama seperti fungsi `fmt.Println()`, perbedaannya fungsi `fmt.Print()` tidak menghasilkan baris baru di akhir output-nya.

Perbedaan lainnya: nilai argument parameter yang ditulis saat pemanggilan fungsi akan di-print tanpa pemisah. Tidak seperti pada fungsi `fmt.Println()` yang nilai argument paremeternya dipisah menggunakan karakter spasi.

```go
fmt.Println("john wick")
fmt.Println("john", "wick")

fmt.Print("john wick\n")
fmt.Print("john ", "wick\n")
fmt.Print("john", " ", "wick\n")
```

Kode di atas menunjukkan perbedaan antara `fmt.Println()` dan `fmt.Print()`. Output yang dihasilkan oleh 5 statement di atas adalah sama, meski cara yang digunakan berbeda.

Bila menggunakan `fmt.Println()`, maka tidak perlu menambahkan spasi di tiap kata, karena fungsi tersebut akan secara otomatis menambahkannya di sela-sela text. Berbeda dengan `fmt.Print()` yang perlu ditambahkan spasi, karena fungsi ini tidak menambahkan spasi secara otomatis di sela-sela nilai text yang digabungkan.

## A.11.2. Deklarasi Multi Konstanta

Sama seperti variabel, konstanta juga dapat dideklarasikan secara bersamaan. Berikut adalah contoh deklarasi konstanta dengan tipe data dan nilai yang berbeda.

```go
const (
    square          = "kotak"
    isToday bool    = true
    numeric uint8   = 1
    floatNum        = 2.2
)
```

- `square`, dideklarasikan dengan metode _type inference_ dengan tipe data **string** dan nilainya **"kotak"**
- `isToday`, dideklarasikan dengan metode _manifest typing_ dengan tipe data **bool** dan nilainya **true**
- `numeric`, dideklarasikan dengan metode _manifest typing_ dengan tipe data **uint8** dan nilainya **1**
- `floatNum`, dideklarasikan dengan metode _type inference_ dengan tipe data **float** dan nilainya **2.2**

Contoh deklarasi konstanta dengan tipe data dan nilai yang sama:

```go
const (
    a = "konstanta"
    b
)
```

> Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.

- `a` dideklarasikan dengan metode _type inference_ dengan tipe data **string** dan nilainya **"konstanta"**
- `b` dideklarasikan dengan metode _type inference_ dengan tipe data **string** dan nilainya **"konstanta"**

Berikut contoh gabungan dari keduanya:

```go
const (
    today string = "senin"
    sekarang
    isToday2 = true
)
```

- `today` dideklarasikan dengan metode _manifest typing_ dengan tipe data **string** dan nilainya **"senin"**
- `sekarang` dideklarasikan dengan metode _manifest typing_ dengan tipe data **string** dan nilainya **"senin"**
- `isToday2` dideklarasikan dengan metode _type inference_ dengan tipe data **bool** dan nilainya **true**

Berikut contoh deklrasi _multiple_ konstanta dalam satu baris:

```go
const satu, dua = 1, 2
const three, four string = "tiga", "empat"
```

- `satu`, dideklarasikan dengan metode  _type inference_ dengan tipe data **int** dan nilainya **1**
- `dua`, dideklarasikan dengan metode _type inference_ dengan tipe data **int** dan nilainya **2**
- `three`, dideklarasikan dengan metode _manifest typing_ dengan tipe data **string** dan nilainya **"tiga"**
- `four`, dideklarasikan dengan metode _manifest typing_ dengan tipe data **string** dan nilainya **"empat"**

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.11-konstanta">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.11...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
