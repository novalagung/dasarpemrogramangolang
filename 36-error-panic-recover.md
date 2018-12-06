# A.36. Error, Panic, dan Recover

Error merupakan topik yang penting dalam pemrograman golang. Di bagian ini kita akan belajar mengenai pemanfaatan error dan cara membuat custom error sendiri.

Kita juga akan belajar tentang penggunaan **panic** untuk memunculkan panic error, dan **recover** untuk mengatasinya.

## A.36.1. Pemanfaatan Error

`error` adalah sebuah tipe. Error memiliki memiliki 1 buah property berupa method `Error()`, method ini mengembalikan detail pesan error. Error termasuk tipe yang isinya bisa kosong atau `nil`.

Di golang, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu. Biasanya, salah satu kembalian adalah bertipe `error`. Contohnya seperti pada fungsi `strconv.Atoi()`.

`strconv.Atoi()` berguna untuk mengkonversi data string menjadi numerik. Fungsi ini mengembalikan 2 nilai balik. Nilai balik pertama adalah hasil konversi, dan nilai balik kedua adalah `error`.

Ketika konversi berjalan mulus, nilai balik kedua akan bernilai `nil`. Sedangkan ketika konversi gagal, penyebabnya bisa langsung diketahui dari nilai balik kedua.

Dibawah ini merupakan contoh program sederhana untuk deteksi inputan dari user, apakah numerik atau bukan. Dari sini kita akan belajar mengenai pemanfaatan error.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var input string
    fmt.Print("Type some number: ")
    fmt.Scanln(&input)

    var number int
    var err error
    number, err = strconv.Atoi(input)

    if err == nil {
        fmt.Println(number, "is number")
    } else {
        fmt.Println(input, "is not number")
        fmt.Println(err.Error())
    }
}
```

Ketika program dijalankan, muncul tulisan `"Type some number: "`, ketik sebuah angka lalu enter.

`fmt.Scanln(&input)` bertugas mengambil inputan yang diketik user sebelum dia menekan enter, lalu menyimpannya sebagai string ke variabel `input`.

Selanjutnya variabel tersebut dikonversi ke tipe numerik menggunakan `strconv.Atoi()`. Fungsi tersebut mengembalikan 2 data, ditampung oleh `number` dan `err`.

Data pertama (`number`) berisi hasil konversi. Dan data kedua `err`, berisi informasi errornya (jika memang terjadi error ketika proses konversi).

Setelah itu dilakukan pengecekkan, ketika tidak ada error, `number` ditampilkan. Dan jika ada error, `input` ditampilkan beserta pesan errornya.

Pesan error bisa didapat dari method `Error()` milik tipe `error`.

![Penerapan error](images/A.36_1_error.png)

## A.36.2. Membuat Custom Error

Selain memanfaatkan error hasil kembalian fungsi, kita juga bisa membuat error sendiri dengan menggunakan fungsi `errors.New` (untuk menggunakannya harus import package `errors` terlebih dahulu).

Pada contoh berikut ditunjukan bagaimana cara membuat custom error. Pertama siapkan fungsi dengan nama `validate()`, yang nantinya digunakan untuk pengecekan input, apakah inputan kosong atau tidak. Ketika kosong, maka error baru akan dibuat.

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

func validate(input string) (bool, error) {
    if strings.TrimSpace(input) == "" {
        return false, errors.New("cannot be empty")
    }
    return true, nil
}
```

Selanjutnya di fungsi main, buat proses sederhana untuk capture inputan user. Manfaatkan fungsi `validate()` untuk mengecek inputannya.

```go
func main() {
    var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        fmt.Println(err.Error())
    }
}
```

Fungsi `validate()` mengembalikan 2 data. Data pertama adalah nilai `bool` yang menandakan inputan apakah valid atau tidak. Data ke-2 adalah pesan error-nya (jika inputan tidak valid).

Fungsi `strings.TrimSpace()` digunakan untuk menghilangkan karakter spasi sebelum dan sesudah string. Ini dibutuhkan karena user bisa saja menginputkan spasi lalu enter.

Ketika inputan tidak valid, maka error baru dibuat dengan memanfaatkan fungsi `errors.New()`.

![Custom error](images/A.36_2_custom_error.png)

## A.36.3. Penggunaan `panic`

Panic digunakan untuk menampilkan *trace* error sekaligus menghentikan flow goroutine (ingat, `main` juga merupakan goroutine). Setelah ada panic, proses selanjutnya tidak di-eksekusi (kecuali proses yang di-defer, akan tetap dijalankan tepat sebelum panic muncul).

Panic memnuculkan pesan di console, sama seperti `fmt.Println()` hanya saja informasi yang ditampilkan sangat mendetail dan heboh.

Pada program yang telah kita buat tadi, ubah `fmt.Println()` yang berada di dalam blok kondisi `else` pada fungsi main menjadi `panic()`, lalu tambahkan `fmt.Println()` setelahnya.

```go
func main() {
    var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        fmt.Println(err.Error())
        fmt.Println("end")
    }
}
```

Coba jalankan program, lalu langsung tekan enter, error panic akan muncul, dan baris kode setelahnya tidak dijalankan.

![Menampilkan error menggunakan panic](images/A.36_3_panic.png)

## A.36.4. Penggunaan `recover`

Recover berguna untuk meng-handle panic error. Pada saat panic muncul, recover men-take-over goroutine yang sedang panic (pesan panic tidak akan muncul).

Kita modif sedikit fungsi di-atas untuk mempraktekkan bagaimana cara penggunaan recover.

Tambahkan fungsi `catch()`, dalam fungsi ini terdapat statement `recover()` yang dia akan mengembalikan pesan error panic yang seharusnya muncul. Fungsi yang berisikan `recover()` harus di-defer, karena meskipun muncul panic, defer tetap di-eksekusi.

```go
func catch() {
    if r := recover(); r != nil {
        fmt.Println("Error occured", r)
    } else {
        fmt.Println("Application running perfectly")
    }
}

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
```

Output:

![Handle panic menggunakan recover](images/A.36_4_recover.png)

## A.36.5. Pemanfaatan `recover` pada IIFE

Recover akan dieksekusi (pada waktunya) di blok kode dimana blok fungsi untuk recover ditempatkan.

```go
func main() {

    // recover untuk panic dalam fungsi main
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic occured", r)
        } else {
            fmt.Println("Application running perfectly")
        }
    }()

    panic("some error happen")
}
```

Di real-world development, ada kalanya recover dibutuhkan tidak pada fungsi terluar, tetapi dalam blok kode yg lebih spesifik. Sebagai contoh: recover panic dalam perulangan tanpa memberhentikan perulangan itu sendiri.

```go
func main() {
    data := []string{"superman", "aquaman", "wonder woman"}

    for _, each := range data {

		func() {

            // recover untuk iife dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happen")
        }()
        
	}
}
```

Pada kode di atas, di dalam perulangan terdapat sebuah IIFE. Berisi operasi sesuai kebutuhan dan sebuah fungsi yang di-defer untuk recover panic. Ketika terjadi panic, maka perulangan tersebut tidak akan rusak, tetap lanjut.
