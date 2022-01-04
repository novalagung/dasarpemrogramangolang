# A.39. Random

Pada chapter ini kita akan belajar cara untuk mengutilisasi package `math/rand` untuk menciptakan data acak atau random.

## A.39.1. Definisi

Random Number Generator (RNG) merupakan sebuah perangkat (bisa software, bisa hardware) yang menghasilkan data deret/urutan angka yang sifatnya acak.

RNG bisa berupa hardware yang murni bisa menghasilkan data angka acak, atau bisa saja sebuah [pseudo-random](https://en.wikipedia.org/wiki/Pseudorandom_number_generator) yang menghasilkan deret angka-angka yang **terlihat acak** tetapi sebenarnya tidak benar-benar acak, yang deret angka tersebut sebenarnya merupakan hasil kalkulasi algoritma deterministik dan probabilitas. Jadi untuk pseudo-random ini, asalkan kita tau *state*-nya maka kita akan bisa menebak hasil deret angka random-nya.

Dalam per-randoman-duniawi terdapat istilah **seed** atau titik mulai (*starting point*). Seed ini digunakan oleh RNG dalam peng-generate-an angka random di tiap urutannya.

Sedikit ilustrasi mengenai korelasi antara seed dengan RNG, agar lebih jelas.

- Dimisalkan saya menggunakan seed yaitu angka `10`, maka ketika fungsi RNG dijalankan untuk pertama kalinya, output angka yang dihasilkan pasti `5221277731205826435`. Angka random tersebut pasti *fix* dan akan selalu menjadi output angka random pertama yang dihasilkan, ketika seed yang digunakan adalah angka `10`.
- Misalnya lagi, fungsi RNG di-eksekusi untuk ke-dua kalinya, maka angka random kedua yang dihasilkan adalah pasti `3852159813000522384`. Dan seterusnya.
- Misalkan lagi, fungsi RNG di-eksekusi lagi, maka angka random ketiga pasti `8532807521486154107`.
- Jadi untuk seed angka `10`, akan selalu menghasilkan angka random ke-1: `5221277731205826435`, ke-2: `3852159813000522384`, ke-3 `8532807521486154107`. Meskipun fungsi random dijalankan di program yang berbeda, di waktu yang berbeda, di environment yang berbeda, jika seed adalah `10` maka deret angka random yang dihasilkan pasti sama seperti contoh di atas.

## A.39.2. Package `math/rand`

Di Go terdapat sebuah package yaitu `math/rand` yang isinya banyak sekali API untuk keperluan penciptaan angka random. Package ini mengadopsi **PRNG** atau *pseudo-random* number generator. Deret angka random yang dihasilkan sangat tergantung dengan angka **seed** yang digunakan.

Cara menggunakan package ini sangat mudah, yaitu cukup import `math/rand`, lalu set seed-nya, lalu panggil fungsi untuk generate angka random-nya. Lebih jelasnya silakan cek contoh berikut.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("random ke-1:", rand.Int()) // 5221277731205826435
	fmt.Println("random ke-2:", rand.Int()) // 3852159813000522384
	fmt.Println("random ke-3:", rand.Int()) // 8532807521486154107
}
```

Fungsi `rand.Seed()` digunakan untuk penentuan nilai seed. Fungsi `rand.Int()` digunakan untuk generate angka random dalam bentuk numerik bertipe `int`. Fungsi `rand.Int()` ini setiap kali dipanggil akan menghasilkan angka berbeda, tapi pasti hasilnya akan selalu tetap jika mengacu ke deret.

- Angka random ke-1 akan selalu `5221277731205826435`
- Angka random ke-2 akan selalu `3852159813000522384`
- Angka random ke-3 akan selalu `8532807521486154107`
- Dan seterusnya ...

Coba jalankan program di atas beberapa kali, hasilnya pasti selalu sama untuk angka random ke-1, ke-2, dan seterusnya.

![Random Golang](images/A_random_1.png)

## A.39.3. Unique Seed

Lalu bagaimana cara agar angka yang dihasilkan selalu berbeda setiap kali dipanggil? Apakah harus set ulang seed-nya? Jangan, karena kalau seed di-set ulang maka urutan deret random akan berubah. Seed hanya perlu di set sekali di awal. Lha, terus bagaimana?

Jadi begini, setiap kali `rand.Int()` dipanggil, hasilnya itu selalu berbeda, tapi sangat bisa diprediksi jika kita tau seed-nya, dan ini adalah masalah besar. Nah, ada cara agar angka random yang dihasilkan tidak berulang-ulang selalu contoh di-atas, caranya adalah menggunakan angka yang *unique*/unik sebagai seed, contohnya seperti angka [unix nano](https://en.wikipedia.org/wiki/GNU_nano) dari waktu sekarang.

Coba modifikasi program dengan kode berikut, lalu jalankan ulang. Jangan lupa meng-import package `time` ya.

```go
rand.Seed(time.Now().UTC().UnixNano())
fmt.Println(rand.Int())
fmt.Println(rand.Int())
fmt.Println(rand.Int())
```

![Random Golang with unix nano seed](images/A_random_2.png)

Bisa dilihat, setiap program dieksekusi angka random nya selalu berbeda, hal ini karena seed yang digunakan pasti berbeda satu sama lain saat program dijalankan. Seed-nya adalah angka unix nano dari waktu sekarang.

## A.39.4. Random Tipe Data Numerik Lainnya

Di dalam package `math/rand`, ada banyak fungsi untuk generate angka random. Fungsi `rand.Int()` hanya salah satu dari fungsi yang tersedia didalam package tersebut, yang gunanya adalah menghasilkan angka random bertipe `int`.

Selain itu, ada juga `rand.Float32()` yang menghasilkan angka random bertipe `float32`. Ada juga `rand.Uint32()` yang menghasilkan angka random bertipe *unsigned* int, dan lainnya.

lebih detailnya silakan merujuk ke https://golang.org/pkg/math/rand/

## A.39.5. Angka Random Index Tertentu

Gunakan fungsi `rand.Intn(n)` untuk mendapatkan angka random pada indeks ke `n`. Dengan ini tidak perlu memanggil `rand.Int()` tiga kali untuk mendapatkan angka random ke-tiga, melainkan cukup gunakan `rand.Intn(2)` (indeks dari 0 ya).

## A.39.6. Random Tipe Data String

Untuk menghasilkan data random string, ada banyak cara yang bisa digunakan, salah satunya adalah dengan memafaatkan alfabet dan hasil random numerik.

```go
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
```

Dengan fungsi di atas kita bisa dengan mudah meng-generate string random dengan panjang karakter yang sudah ditentukan, misal `randomString(10)` akan menghasilkan random string 10 karakter.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.39-random">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.39...</a>
</div>
