# A.3. Go Modules

Pada bagian ini kita akan belajar cara pembuatan project baru menggunakan Go Modules.

## A.3.1. Penjelasan

Go modules merupakan tools untuk manajemen dependensi resmi milik Go. Modules digunakan untuk menginisialisasi sebuah project, sekaligus melakukan manajemen terhadap *3rd party* atau *library* atau *dependency* yang digunakan dalam project.

Modules penggunaannya adalah via CLI. Jika pembaca sudah sukses meng-*install* Go, maka otomatis bisa menggunakan operasi CLI Go Modules.

> Di Go, istilah modules (atau module) maknanya adalah sama dengan project. Jadi gak perlu bingung

## A.3.2. Inisialisasi Project Menggunakan Go Modules

Command `go mod init` digunakan untuk menginisialisasi project baru.

Mari langsung praktekan saja. Buat folder baru, bisa via CLI atau lewat browser/finder.

```bash
mkdir project-pertama
cd project-pertama
go mod init project-pertama
```

Bisa dilihat pada *command* di atas ada direktori `project-pertama`, dibuat. Setelah masuk ke direktori tersebut, perintah `go mod init project-pertama` dijalankan. Dengan ini maka kita telah menginisialisasi direktori/folder `project-pertama` sebagai sebuah project Go dengan nama `project-pertama`.

![Init project](images/A_go_modules_1_initmodule.png)

Skema penulisan command `go mod`:

```
go mod init <nama-project>
go mod init project-pertama
```

Di sini kita tentukan nama project adalah sama dengan nama folder, ini merupakan *best practice* di Go.

> Nama project dan nama module merupakan artinya adalah sama. Ingat, module adalah sama dengan project

Eksekusi perintah `go mod init` menghasilkan satu buah file baru bernama `go.mod`. File ini digunakan oleh Go toolchain untuk menandai bahwa folder di mana file tersebut berada adalah folder project. Jadi pastikan untuk tidak menghapus file tersebut.

Ok, sekian. Cukup itu saja cara inisialisasi project di Go.

---

O iya, sebenarnya selain Go Modules, setup project di Go juga bisa menggunakan `$GOPATH` yang pembahasannya ada di chapter ([A.4. Setup GOPATH Dan Workspace](/A-gopath-dan-workspace.html)).

Namun metode inisialisasi project via GOPATH sudah outdate dan kurang dianjurkan untuk project-project yang dikembangkan menggunakan Go versi terbaru (1.14 ke atas).

Jadi setelah chapter ini, penulis boleh langsung lompat ke pembahasan di chapter [A.5. Instalasi Editor](/A-instalasi-editor.html).

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
