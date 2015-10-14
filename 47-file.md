# File

Ada beberapa cara yang bisa digunakan untuk keperluan operasi file di Golang. Pada bab ini kita akan mempelajari teknik yang paling dasar, dengan memanfaatkan `os.File`.

## Membuat File Baru

Pembuatan file di Golang sangatlah mudah, cukup dengan memanggil fungsi `os.Create` lalu memasukkan path file ingin dibuat sebagai parameter fungsi tersebut.

Jika file yang akan dibuat sudah ada, maka akan ditimpa. Bisa memanfaatkan `os.IsNotExist` untuk mendeteksi apakah file sudah dibuat atau belum.

Berikut merupakan contoh pembuatan file.

```go
import "fmt"
import "os"

var path = "/Users/novalagung/Documents/temp/test.txt"

func checkError(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
    }
}

func createFile() {
    // deteksi apakah file sudah ada
    var _, err = os.Stat(path)

    // buat file baru jika belum ada
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        checkError(err)
        defer file.Close()
    }
}

func main() {
    createFile()
}
```

Fungsi `os.Stat()` mengembalikan 2 data, yaitu informasi tetang path yang dicari, dan error (jika ada). Masukkan error kembalian fungsi `os.Stat()` sebagai parameter fungsi `os.IsNotExist()`, untuk mendeteksi apakah file yang akan dibuat sudah ada. Jika belum ada, maka fungsi tersebut akan mengembalikan nila `true`.

Fungsi `os.Create()` digunakan untuk membuat file pada path tertentu. Fungsi ini mengembalikan objek `*os.File` dari file yang bersangkutan. perlu diketahui bahwa file yang baru terbuat statusnya adalah otomatis **open**, jadi perlu untuk di-**close** menggunakan fungsi `file.Close()` setelah file tidak digunakan lagi. Membiarkan file terbuka ketika sudah tak lagi digunakan bukan hal yang baik, space memory akan teralokasi dengan sia-sia.

![Membuat file baru](images/47_1_create.png)

## Mengedit Isi File

Untuk mengedit file, yang perlu dilakukan pertama adalah membuka file-nya dengan level akses **write**. Setelah mendapatkan objek file tersebut, gunakan `file.WriteString()` untuk pengisian data. Terakhir panggil `file.Sync()` untuk menyimpan perubahan.

```go
func writeFile() {
    // buka file dengan level akses READ & WRITE
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    checkError(err)
    defer file.Close()

    // tulis data ke file
    _, err = file.WriteString("halo\n")
    checkError(err)
    _, err = file.WriteString("mari belajar golang\n")
    checkError(err)

    // simpan perubahan
    err = file.Sync()
    checkError(err)
}

func main() {
    writeFile()
}
```

Pada program di atas, file dibuka dengan level akses **read** dan **write** dengan kode permission **0664**.

Setelah itu, beberapa string diisikan kedalam file tersebut menggunakan `file.WriteString`.

Di akhir, semua perubahan terhadap file akan disimpan dengan dipanggilnya `file.Sync()`.

![Mengedit file](images/47_2_write.png)

## Membaca Isi File

File yang ingin dibaca harus dibuka terlebih dahulu dengan level akses minimal *read*.

Setelah itu, gunakan `file.Read()` dengan parameter variabel yang dimana hasil proses baca akan disimpan ke variabel tersebut.

```go
// tambahkan di bagian import package io
import "io"

func readFile() {
    // buka file
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    checkError(err)
    defer file.Close()

    // baca file
    var text = make([]byte, 1024)
    for {
        n, err := file.Read(text)
        if err != io.EOF {
            checkError(err)
        }
        if n == 0 {
            break
        }
    }
    fmt.Println(string(text))
    checkError(err)
}

func main() {
    readFile()
}
```

Variabel `text` adalah slice `[]byte` dengan alokasi elemen 1024. Variabel tersebut akan berisikan data hasil fungsi `file.Read()`.

Fungsi tersebut dieksekusi dalam perulangan, yang hanya berhenti ketika sudah tidak ada baris text yang bisa dibaca dari file bersangkutan.

Error yang muncul ketika eksekusi `file.Read()` akan difilter, ketika errornya adalah selain `io.EOF` maka proses baca file akan berlanjut.

![Membaca isi file](images/47_3_read.png)

## Menghapus File

Cara menghapus file sangatlah mudah, cukup panggil fungsi `os.Remove()`, masukkan path file yang ingin dihapus sebagai parameter.

```go
func deleteFile() {
    var err = os.Remove(path)
    checkError(err)
}

func main() {
    deleteFile()
}
```

![Menghapus file](images/47_4_delete.png)
