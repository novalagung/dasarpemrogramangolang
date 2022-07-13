# A.11. Konstanta

Konstanta adalah jenis variabel yang nilainya tidak bisa diubah. Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya.

## A.11.1. Penggunaan Konstanta

Data seperti **pi** (22/7), kecepatan cahaya (299.792.458 m/s), adalah contoh data yang tepat jika dideklarasikan sebagai konstanta daripada variabel, karena nilainya sudah pasti dan tidak berubah.

Cara penerapan konstanta sama seperti deklarasi variabel biasa, selebihnya tinggal ganti keyword `var` dengan `const`.

```go
const firstName string = "john"
fmt.Print("halo ", firstName, "!\n")
```

Teknik type inference bisa diterapkan pada konstanta, caranya yaitu cukup dengan menghilangkan tipe data pada saat deklarasi.

```go
const lastName = "wick"
fmt.Print("nice to meet you ", lastName, "!\n")
```

#### • Penggunaan Fungsi `fmt.Print()`

Fungsi ini memiliki peran yang sama seperti fungsi `fmt.Println()`, pembedanya fungsi `fmt.Print()` tidak menghasilkan baris baru di akhir outputnya.

Perbedaan lainnya adalah, nilai pada parameter-parameter yang dimasukkan ke fungsi tersebut digabungkan tanpa pemisah. Tidak seperti pada fungsi `fmt.Println()` yang nilai paremeternya digabung menggunakan penghubung spasi.

```go
fmt.Println("john wick")
fmt.Println("john", "wick")

fmt.Print("john wick\n")
fmt.Print("john ", "wick\n")
fmt.Print("john", " ", "wick\n")
```

Kode di atas menunjukkan perbedaan antara `fmt.Println()` dan `fmt.Print()`. Output yang dihasilkan oleh 5 statement di atas adalah sama, meski cara yang digunakan berbeda.

Bila menggunakan `fmt.Println()` tidak perlu menambahkan spasi di tiap kata, karena fungsi tersebut akan secara otomatis menambahkannya di sela-sela nilai. Berbeda dengan `fmt.Print()`, perlu ditambahkan spasi, karena fungsi ini tidak menambahkan spasi di sela-sela nilai parameter yang digabungkan.

## A.11.2. Deklarasi multi konstanta

Sama seperti variabel, konstanta juga dapat di deklarasi secara bersamaan

Berikut contoh deklarasi konstanta dengan nilai dan tipe yang berbeda

```go
const (
    square          = "kotak"
    isToday bool    = true
    numeric uint8   = 1
    floatNum        = 2.2
)
```

- isToday, termasuk _type inference_ dengan tipe data **bool** dan nilai nya **true**
- square, termasuk _manifest typing_ dengan tipe data **string** dan nilai nya **"kotak"**
- numeric, termasuk _manifest typing_ dengan tipe data **uint8** dan nilai nya **1**
- floatNum, termasuk _type inference_ dengan tipe data **float** dan nilai nya **2.2**

Contoh deklarasi konstanta dengan nilai dan tipe yang sama

```go
const (
    a = "konstanta"
    b
)
```

> ketika tipe dan nilai _const_ tidak diberikan, maka tipe dan nilai nya didapat dari deklarasi sebelumnya

- a, termasuk _type inference_ dengan tipe data **string** dan nilai nya **"konstanta"**
- b, termasuk _type inference_ dengan tipe data **string** dan nilai nya **"konstanta"**

Berikut contoh gabungan dari keduanya

```go
const (
    today string = "senin"
    sekarang
    isToday2 = true
)
```

- today, termasuk _manifest typing_ dengan tipe data **string** dan nilai nya **"senin"**
- sekarang, termasuk _manifest typing_ dengan tipe data **string** dan nilai nya **"senin"**
- isToday2, termasuk _type inference_ dengan tipe data **bool** dan nilai nya **true**

Berikut contoh deklrasi _multiple_ konstanta dalam satu baris

```go
const satu, dua = 1, 2
const three, four string = "tiga", "empat"
```

- satu, termasuk _type inference_ dengan tipe data **int** dan nilai nya **1**
- dua, termasuk _type inference_ dengan tipe data **int** dan nilai nya **2**
- three, termasuk _manifest typing_ dengan tipe data **string** dan nilai nya **"tiga"**
- four, termasuk _manifest typing_ dengan tipe data **string** dan nilai nya **"empat"**

sumber [klik disini](https://golangbyexample.com/multiple-constant-declarations-go/)


---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.11-konstanta">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.11...</a>
</div>
