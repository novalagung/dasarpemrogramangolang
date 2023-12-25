# A.7. Program Pertama: Hello World

Semua persiapan sudah selesai, saatnya masuk pada sesi programming. Program pertama yang akan kita buat adalah aplikasi kecil yang menampilkan text **Hello world**.

Pada chapter ini akan dijelaskan secara komprehensif *step-by-step* mulai dari awal. Mulai dari pembuatan project, pembuatan file program, sesi penulisan kode (coding), hingga eksekusi program.

## A.7.1. Inisialisasi Project

Buat direktori bernama `hello-world` bebas ditempatkan di mana. Lalu via CLI, masuk ke direktori tersebut dan jalankan *command* untuk inisialisasi project.

```
mkdir hello-world
cd hello-world
go mod init hello-world
```

![Inisialisasi project](images/A_hello_world_1_init_project.png)

## A.7.2. Load Project Folder ke Editor

Buka editor, di sini penulis menggunakan VSCode. Cari menu untuk menambahkan project, lalu pilih project folder `hello-world`. Untuk beberapa jenis editor, cara load project bisa cukup dengan klik-drag folder tersebut ke editor.

![Load project folder ke editor](images/A_hello_world_2_load_project_to_editor.png)

## A.7.3. Menyiapkan File Program

File program di sini maksudnya adalah file yang isinya *source code* Go. File ini berekstensi `.go`.

Di dalam project yang telah dibuat, siapkan sebuah file dengan nama bebas, yang jelas harus ber-ekstensi `.go`. Pada contoh ini saya menggunakan nama file `main.go`.

Pembuatan file program bisa dilakukan lewat CLI atau browser, atau juga lewat editor. Pastikan file dibuat dalam project folder ya.

![File program](images/A_hello_world_3_new_file_on_editor.png)

## A.7.4. Program Pertama: Hello Word

Setelah project folder dan file program sudah siap, saatnya untuk *programming*.

Di bawah ini merupakan contoh kode program sederhana untuk memunculkan text **Hello world** ke layar output command prompt. Silakan salin kode berikut ke file program yang telah dibuat. Sebisa mungkin jangan copy paste. Biasakan untuk menulis dari awal, agar cepat terbiasa dan familiar dengan Go.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

Setelah kode disalin, buka terminal (atau CMD bagi pengguna Windows), lalu masuk ke direktori proyek, kemudian jalankan program menggunakan perintah `go run`.

```bash
cd hello-world
go run main.go
```

Hasilnya, muncul tulisan **hello world** di layar console.

![Menjalankan program](images/A_hello_world_4_execute_hello_world.png)

Selamat! Anda telah berhasil membuat program Go!

---

Berikut merupakan pembahasan untuk tiap baris kode yang sudah ditulis di atas.

## A.7.5. Penggunaan Keyword `package`

Setiap file program harus memiliki **package**. Setiap project harus ada minimal satu file dengan nama *package* `main`. File yang ber-*package* `main`, akan dieksekusi pertama kali ketika program di jalankan.

Cara menentukan *package* dengan menggunakan keyword `package`, berikut adalah contoh penulisannya.

```go
package <nama-package>
package main
```

## A.7.6. Penggunaan Keyword `import`

Keyword `import` digunakan untuk meng-*import* atau memasukan *package* lain ke dalam file program, agar isi dari package yang di-*import* bisa dimanfaatkan.

*Package* `fmt` merupakan salah satu *package* bawaan yang disediakan oleh Go, isinya banyak fungsi untuk keperluan **I/O** yang berhubungan dengan text.

Berikut adalah skema penulisan keyword `import`:

```go
import "<nama-package>"
import "fmt"
```

## A.7.7. Penggunaan Fungsi `main()`

Dalam sebuah proyek harus ada file program yang di dalamnya berisi sebuah fungsi bernama `main()`. Fungsi tersebut harus berada di file yang package-nya bernama `main`.

Fungsi `main()` adalah yang dipanggil pertama kali pada saat eksekusi program. Contoh penulisan fungsi `main`:

```go
func main() {

}
```

## A.7.8. Penggunaan Fungsi `fmt.Println()`

Fungsi `fmt.Println()` digunakan untuk memunculkan text ke layar (pada konteks ini, terminal atau CMD). Di program pertama yang telah kita buat, fungsi ini memunculkan tulisan **Hello world**.

Skema penulisan keyword `fmt.Println()` bisa dilihat pada contoh berikut.

```go
fmt.Println("<isi-pesan>")
fmt.Println("Hello world")
```

Fungsi `fmt.Println()` berada dalam package `fmt`, maka untuk menggunakannya perlu package tersebut untuk di-import terlebih dahulu.

Fungsi `fmt.Println()` dapat menampung parameter yang tidak terbatas jumlahnya. Semua data parameter akan dimunculkan dengan pemisah tanda spasi.

```go
fmt.Println("Hello", "world!", "how", "are", "you")
```

Contoh statement di atas akan menghasilkan output: **Hello world! how are you**.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.7-hello-world">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.7...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="430px" frameborder="0" scrolling="no"></iframe>
