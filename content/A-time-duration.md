# A.42. Time Duration

Pada chapter ini kita akan belajar tentang tipe waktu durasi yaitu `time.Duration`.

Tipe `time.Duration` ini merepresentasikan durasi, contohnya seperti 1 menit, 2 jam 5 detik, dst. Data dengan tipe ini bisa dihasilkan dari operasi pencarian delta atau selisih dari dua buah objek `time.Time`, atau bisa juga kita buat sendiri.

Tipe durasi ini sangat berguna untuk banyak hal, seperti *benchmarking* durasi ataupun operasi-operasi lainnya yang membutuhkan informasi durasi.

## A.42.1. Praktek 

Mari kita bahas sambil praktek. Silakan tulis kode berikut lalu jalankan.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	time.Sleep(5 * time.Second)
	
	duration := time.Since(start)
	
	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())
}
```

Pada kode di atas, sebuah objek waktu bernama `start` dibuat. Tepat setelah baris tersebut, ada statement `time.Sleep()` yang digunakan untuk menghentikan proses selama X, yang durasinya di-set lewat parameter fungsi tersebut. Bisa dilihat durasi yang dipilih adalah `5 * time.Second`.

Tipe data durasi adalah `time.Duration`, yang sebenarnya tipe ini merupakan tipe buatan baru dari `int64`.

Ada beberapa *predefined* konstanta durasi yang perlu kita ketahui:

- `time.Nanosecond` yang nilainya adalah `1`
- `time.Microsecond` yang nilainya adalah `1000`, atau `1000` x `time.Nanosecond`
- `time.Millisecond` yang nilainya adalah `1000000`, atau `1000` x `time.Microsecond`
- `time.Second` yang nilainya adalah `1000000000`, atau `1000` x `time.Millisecond`
- `time.Minute` yang nilainya adalah `1000000000000`, atau `1000` x `time.Second`
- `time.Hour` yang nilainya adalah `1000000000000000`, atau `1000` x `time.Minute`

Dari list di atas bisa dicontohkan bahwa sebuah data dengan tipe `time.Duration` yang nilainya `1`, maka artinya durasi adalah **1 nanosecond**.

Kembali ke pembahasan fungsi `time.Sleep()`, fungsi ini membutuhkan argumen/parameter durasi dalam bentuk `time.Duration`. Misalnya saya tulis `time.Sleep(1)` maka yang terjadi adalah, waktu statement tersebut hanya akan menghentikan proses selama **1 nanosecond** saja. Jika ingin menghentikan selama 1 detik, maka harus ditulis `time.Sleep(1000000000)`. Nah daripada menulis angka sepanjang itu, cukup saja tulis dengan `1 * time.Second`, artinya adalah 1 detik. Cukup mudah bukan.

Di atas kita gunakan `5 * time.Second` sebagai argumen `time.Sleep()`, maka dengan itu proses akan diberhentikan selama 5 detik.

Sekarang jalankan program yang sudah dibuat.

![Time Duration](images/A_time_duration_1.png)

Bisa dilihat, hasilnya adalah semua statement di bawah `time.Sleep()` dieksekusi setelah 5 detik berlalu. Ini merupakan contoh penggunaan tipe data durasi pada fungsi `time.Sleep()`.

## A.42.2. Kalkulasi Durasi Menggunakan `time.Since()`.

Pada kode di atas, variabel `duration` berisi durasi atau lama waktu antara kapan variabel `start` di-inisialisasi hingga kapan variabel `duration` ini statement-nya dieksekusi.

Cara menghitung durasi bisa menggunakan `time.Since()`. Isi argumen fungsi tersebut dengan variabel bertipe waktu, maka durasi antara waktu pada argument vs ketika statement `time.Since()` akan dihitung.

Pada contoh di atas, karena ada statement `time.Sleep(5 * time.Second)` maka idealnya `time.Since(start)` isinya adalah 5 detik (mungkin lebih sedikit, sekian mili/micro/nano-second, karena eksekusi statement juga butuh waktu).

## A.42.3. Method `time.Duration`

Tipe `time.Duration` memiliki beberapa method yang sangat-sangat berguna untuk keperluan mengambil nilai durasinya dalam unit tertentu. Misalnya, objek durasi tersebut ingin di-ambil nilainya dalam satuan unit detik, maka gunakan `.Seconds()`. Jika ingin dalam bentuk menit, maka gunakan `.Minutes()`, dan lainnya.

Pada contoh di atas, kita mengambil nilai durasi waktu dalam tiga bentuk, yaitu detik, menit, dan jam. Caranya cukup akses saja method-nya, maka kita akan langsung dapat nilainya, tanpa perlu memikirkan operasi aritmatik konversinya. Cukup mudah bukan.

## A.42.4. Kalkulasi Durasi Antara 2 Objek Waktu

Di atas kita sudah membahas cara hitung durasi menggunakan `time.Since()` antara sebuah objek waktu vs kapan statement di-eksekusi. Pada bagian ini, masih mirip, perbedannya adalah hitung durasi dilakukan pada 2 objek waktu.

Silakan perhatikan contoh berikut. Kode berikut esensinya adalah sama dengan kode di atas.

```go
t1 := time.Now()
time.Sleep(5 * time.Second)
t2 := time.Now()

duration := t2.Sub(t1)

fmt.Println("time elapsed in seconds:", duration.Seconds())
fmt.Println("time elapsed in minutes:", duration.Minutes())
fmt.Println("time elapsed in hours:", duration.Hours())
```

Method `.Sub()` milik objek `time.Time` digunakan untuk mencari selisih waktu. Pada contoh di atas, durasi antara waktu `t1` dan waktu `t2` dihitung. Method `.Sub()` ini menghasilkan nilai balik bertipe `time.Duration`.

## A.42.5. Konversi Angka ke `time.Duration` 

Kita bisa mengalikan angka literal dengan konstanta `time.Duration` untuk menciptakan variabel/objek bertipe durasi. Contohnya seperti yang sudah kita terapkan sebelumnya, yaitu `5 * time.Second` yang menghasilkan data durasi 5 detik. Contoh lainnya:

```go
12 * time.Minute 			// 12 menit
65 * time.Hour 				// 65 jam
150000 * time.Milisecond 	// 150k milidetik atau 150 detik
45 * time.Microsecond 		// 45 microdetik
233 * time.Nanosecond 		// 233 nano detik
```

Sedikit kembali ke pembahasan dasar di awal-awal chapter, operasi aritmatika di golang hanya bisa dilakukan ketika data adalah 1 tipe. Selebihnya harus ada casting atau konversi tipe data agar bisa dioperasikan.

Tipe `time.Duration` diciptakan menggunakan tipe `ìnt64`. Jadi jika ingin mengalikan `time.Duration` dengan suatu angka, maka pastikan tipe-nya juga sama yaitu `time.Duration`. Jika angka tersebut tidak ditampung dalam variabel terlebih dahulu (contohnya seperti di atas) maka bisa langsung kalikan saja. Jika ditampung ke variabel terlebih dahulu, maka pastikan tipe variabelnya adalah `time.Duration`. Contoh:


```go
var n time.Duration = 5
duration := n * time.Second
```

Atau bisa manfaatkan casting untuk mengkonversi data numerik ke tipe `time.Duration`. Contoh:

```go
n := 5
duration := time.Duration(n) * time.Second
```

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.42-time-duration">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.42...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="360px" frameborder="0" scrolling="no"></iframe>
