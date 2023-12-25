# A.3. Go Modules

Pada bagian ini kita akan belajar cara inisialisasi project menggunakan Go Modules (atau Modules).

## A.3.1. Penjelasan

Go modules merupakan manajemen dependensi resmi untuk Go. Modules ini diperkenalkan pertama kali di `go1.11`, sebelum itu pengembangan project Go dilakukan dalam `GOPATH`.

Modules digunakan untuk menginisialisasi sebuah project, sekaligus melakukan manajemen terhadap *3rd party* atau *library* lain yang dipergunakan.

Modules penggunaannya adalah lewat CLI. Dan jika temen-temen sudah sukses meng-*install* Go, maka otomatis bisa mempergunakan Go Modules.

> Modules atau Module di sini merupakan istilah untuk project ya. Jadi jangan bingung.

## A.3.2. Inisialisasi Project Menggunakan Go Modules

Command `go mod init` digunakan untuk menginisialisasi project baru.

Mari kita praktekan, buat folder baru, bisa via CLI atau lewat browser/finder.

```bash
mkdir project-pertama
cd project-pertama
go mod init project-pertama
dir
```

Bisa dilihat pada *command* di atas ada direktori `project-pertama`, dibuat. Setelah masuk ke direktori tersebut, perintah `go mod init project-pertama` dijalankan. Dengan ini maka kita telah menginisialisasi direktori `project-pertama` sebagai sebuah project Go dengan nama `project-pertama` (kebetulan di sini nama project sama dengan nama direktori-nya).

![Init project](images/A_go_modules_1_initmodule.png)

Skema penulisan command `go mod`:

```
go mod init <nama-project>
go mod init project-pertama
```

Untuk nama project, umumnya adalah disamakan dengan nama direktori, tapi bisa saja sebenarnya menggunakan nama yang lain.

> Nama project dan Nama module merupakan istilah yang sama.

Eksekusi perintah `go mod init` menghasilkan satu buah file baru bernama `go.mod`. File ini digunakan oleh Go toolchain untuk menandai bahwa folder di mana file tersebut berada adalah folder project. Jadi jangan di hapus ya file tersebut.

---

Ok, sekian. Cukup itu saja cara inisialisasi project di Go.

O iya, sebenarnya selain Go Modules, setup project di Go juga bisa menggunakan `$GOPATH` ([A.4. Setup GOPATH Dan Workspace](/A-gopath-dan-workspace.html)). Tapi inisialisasi project dengan GOPATH sudah outdate dan kurang dianjurkan untuk project-project yang dikembangkan menggunakan Go versi terbaru (1.14 ke atas). Jadi setelah chapter ini, bisa langsung lanjut ke [A. Instalasi Editor](/A-instalasi-editor.html).

---

<iframe src="partial/ebooks.html" class="partial-ebooks-wrapper" frameborder="0" scrolling="no"></iframe>
