# A.3.A. Go Modules

Pada bagian ini kita akan belajar cara inisialisasi projek menggunakan Go Modules (atau Modules).

## A.3.A.1. Penjelasan

Go modules merupakan manajemen dependensi resmi untuk Go. Modules ini diperkenalkan pertama kali di `go1.11`, sebelum itu pengembangan projek Go dilakukan dalam `GOPATH`.

Modules digunakan untuk menginisialisasi sebuah projek, sekaligus melakukan manajemen terhadap *3rd party* atau *library* lain yang dipergunakan.

Modules penggunaannya adalah lewat CLI. Dan jika temen-temen sudah sukses meng-*install* Go, maka otomatis bisa mempergunakan Go Modules.

> Modules atau Module disini merupakan istilah untuk project ya. Jadi jangan bingung.

## A.3.A.2. Inisialisasi Projek Menggunakan Go Modules

Command `go mod init` digunakan untuk menginisalisasi projek baru.

Mari kita praktekan, buat folder baru, bisa via CLI atau lewat browser/finder.

```bash
mkdir projek-pertama
cd project-pertama
go mod init project-pertama
dir
```

Bisa dilihat pada *command* di atas ada direktori `projek-pertama`, dibuat. Setelah masuk ke direktori tersebut, perintah `go mod init project-pertama` dijalankan. Degan ini maka kita telah menginisialisasi direktori `projek-pertama` sebagai sebuah project Go dengan nama `project-pertama` (kebetulan di sini nama projek sama dengan nama direktori-nya).

![Init project](images/A.3_1_initmodule.png)

Skema penulisan command `go mod`:

```
go mod init <nama-project>
go mod init project-pertama
```

Untuk nama project, umumnya adalah disamakan dengan nama direktori, tapi saja sebenarnya menggunakan nama yang lain.

> Nama project dan Nama module merupakan istilah yang sama.

Eksekusi perintah `go mod init` menghasilkan satu buah file baru bernama `go.mod`. File ini digunakan oleh Go toolchain untuk menandai bahwa folder dimana file tersebut berada adalah folder projek. Jadi jangan di hapus ya command tersebut.

---

Ok, sekian. Cukup itu saja cara inisalisasi projek di Go.

O iya, sebenarnya selain Go Modules, setup projek di Go juga bisa menggunakan [GOPATH](/3-setup-go-project-dengan-go-modules.html). Tapi inisalisasi porjek dengan GOPATH sudah outdate dan kurang dianjurkan untuk projek-projek yang dikembangkan menggunakan Go versi terbaru (1.14 ke atas).

Jadi setelah bab ini, silakan langsung lanjut ke bab [A.4. Instalasi Editor](4-instalasi-editor.md).
