# A.3. Go Modules

Pada bagian ini kita akan belajar cara pemanfaatan Go Modules (atau Modules) untuk manajemen project dan dependensi.

## A.3.1. Penjelasan

Go modules merupakan manajemen dependensi resmi untuk Go. Modules ini diperkenalkan pertama kali di `go1.11`, sebelum itu pengembangan projek Go dilakukan dalam `GOPATH`.

Modules digunakan untuk menginisialisasi sebuah projek, sekaligus melakukan manajemen terhadap *3rd party* atau *library* lain yang dipergunakan.

Modules penggunaannya adalah lewat CLI. Dan jika temen-temen sudah sukses meng-*install* Go, maka otomatis bisa mempergunakan Go Modules.

> Modules atau Module disini merupakan istilah untuk project ya. Jadi jangan bingung.

## A.3.2. Inisialisasi Projek Menggunakan Go Modules

Command `go mod init` digunakan untuk menginisalisasi projek baru. Buat direktori kosong, lalu masuk ke direktori tersebut lewat CLI, kemudian jalankan perintah `go mod init namaproject`. Untuk nama project silakan ganti sesuai dengan nama projek yang ditentukan.

Command tersebut akan menghasilkan file baru di dalam *project* folder dengan nama `go.mod`, yang isinya merupakan konfigurasi module untuk project tersebut.

Pada contoh berikut, saya siapkan sebuah direktori bernama `chapter-A.60-go-modules`, lalu saya inisialisasi folder tersebut sebagai projek Go menggunakan Go Modules

> Nama folder `chapter-A.60-go-modules` itu hanya nama, temen-temen bisa gunakan nama apapun.

Ok, mungkin akan muncul beberapa pertanyaan. Salah satunya adalah kenapa command-nya harus `go mod init chapter-A.60-go-modules`, kenapa tidak cukup hanya `go mod init` saja? Hal ini karena mulai awal, projek ini sudah tidak didalam `$GOPATH`. Argument `chapter-A.60-go-modules` setelah command inisialisasi berguna untuk menentukan package path dari project, hal ini penting untuk keperluan import package lain yang merupakan sub folder project.

Agar lebih mudah untuk dipahami, silakan lihat contoh berikut.

![Example Module Name 1](images/A.60_3_module_name_1.png)

Coba perhatikan gambar di atas. Nama folder adalah `chapter-A.60-go-modules`, tetapi pada saat import `models`, package path yang digunakan adalah `myproject/models`. Harusnya ini error kan? Tetapi tidak, *hmmmm*.

Dengan menggunakan Go Modules kita bisa menentukan base package path suatu project, tidak harus sama dengan nama folder-nya. Pada contoh ini nama folder adalah `chapter-A.60-go-modules` tapi base package path adalah `myproject`.

Agar lebih jelas lagi, lanjut ke contoh ke-dua.

![Example Module Name 2](images/A.60_4_module_name_2.png)

Pada contoh ini, nama module di-set panjang, `github.com/novalagung/myproject`. Dengan nama folder masih tetap sama. Hasilnya aplikasi tetap berjalan lancar.

Nama package path bisa di set dalam bentuk arbitrary path seperti pada gambar. Sangat membantu bukan sekali, tidak perlu membuat nested folder terlebih dahulu seperti yang biasa dilakukan dalam `$GOPATH`.

## A.60.4. Vendoring Dengan Go Modules

Vendoring, seperti yang kita tau berguna untuk men-centralize 3rd party libraries. Sedangkan dari yang sudah dipelajari, kita bisa membuat sebuah project untuk tidak tergantung dengan `$GOPATH` lewat Go Modules.

Sebenarnya, Go Modules, selain untuk membuat sebuah project tidak tergantung dengan `$GOPATH`, adalah juga berguna untuk manajemen dependencies project. Jadi dengan Go Modules tidak perlu menggunakan `dep`, karena di dalam Go Modules sudah include kapabilitas yang sama dengan `dep` yaitu untuk me-manage dependencies.

Cara untuk meng-enable vendoring dengan Go Modules, adalah lewat command `go mod vendor`.

![Go Modules + Vendoring](images/A.60_5_mod_vendor.png)

Jika pada sebuah projek sudah enabled `dep`, dan ingin di enable Go Modules. Maka file `Gopkg.lock` yang sudah ada akan dikonversi ke-dalam bentuk `go.mod` dan `go.sum`.

## A.60.5. Sinkronisasi Dependencies

Gunakan command `go mod tidy` untuk sinkronisasi dependencies yang digunakan dalam project. Dengan command tersebut, secara otomatis 3rd party yang belum ditambahkan akan dicatat dan ditambahkan; dan yang tidak digunakan tapi terlanjut tercatat akan dihapuskan.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.60-go-modules">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.60...</a>
</div>
