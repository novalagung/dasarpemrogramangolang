# A.4. Instalasi Editor

Proses pembuatan aplikasi menggunakan Go akan lebih maksimal jika didukung oleh editor atau **IDE** yang pas. Ada cukup banyak pilihan bagus yang bisa dipertimbangkan, diantaranya: Brackets, JetBrains GoLand, Netbeans, Atom, Brackets, Visual Studio Code, Sublime Text, dan lainnya.

Penulis sarankan untuk memilih editor yang paling nyaman digunakan, preferensi masing-masing pastinya berbeda. Penulis sendiri sempat menggunakan Sublime Text 3, tapi sekarang pindah ke **Visual Studio Code**. Editor ini sangat ringan, mudah didapat, dan memiliki ekstensi yang bagus untuk bahasa Go. Jika pembaca ingin menggunakan editor yang sama, maka silakan melanjutkan guide berikut.

Di bab ini akan dijelaskan bagaimana cara instalasi editor Visual Studio Code.

## A.4.1. Instalasi Editor Visual Studio Code

 1. Download Visual Studio Code di [https://code.visualstudio.com/Download](https://code.visualstudio.com/Download), pilih sesuai dengan sistem operasi yang digunakan.
 2. Jalankan installer.
 3. Setelah selesai, jalankan aplikasi.

![Tampilan Visual Studio Code](images/A.4_1_visual_studio_code.png)

## A.4.2. Instalasi Extensi Go

Dengan meng-instal Go Extension, maka programming menggunakan bahasa ini akan menjadi sangat nyaman lewat VS Code. Banyak benefit yang didapat dengan meng-install ekstensi ini, beberapa diantaranya adalah integrasi dengan kompiler Go, auto lint on save, testing with coverage, fasilitas debugging with breakpoints, dan lainnya.

Cara instalasi ekstensi sendiri cukup mudah, klik `View -> Extension` atau klik ikon *Extension Marketplace* di sebelah kiri (silakan lihat gambar berikut, deretan button paling kiri yang dilingkari merah). Setelah itu ketikkan **Go** pada inputan search, silakan install ekstensi Go buatan Microsoft, biasanya muncul paling atas sendiri.

![VSCode Go extension](images/A.4_2_vscode_go_extension.png)

## A.4.3. Instalasi Editorconfig

[Editorconfig](https://editorconfig.org/) membantu kita supaya *coding style* kita konsisten untuk dibaca oleh banyak developer dan dimuat di berbagai macam **IDE**. Instalasinya di *VSCode* cukup mudah, kita bisa temukan *extension*-nya dan klik install seperti pada gambar berikut.

![VSCode Editorconfig extension](images/A.4_3_vscode_editorconfig_extension.png)

Editorconfig pada sebuah proyek (biasanya berada di root direktori proyek tersebut) berupa konfigurasi format file `.editorconfig` yang berisi definisi style penulisan yang menyesuaikan dengan standar penulisan masing-masing bahasa pemrograman. Misalnya untuk [*style guide* **GO**](https://golang.org/doc/effective_go.html) kita bisa mulai dengan menggunakan konfigurasi sederhana sebagai berikut:

```
root = true

[*]
insert_final_newline = true
charset = utf-8
trim_trailing_whitespace = true
indent_style = space
indent_size = 2

[{Makefile,go.mod,go.sum,*.go}]
indent_style = tab
indent_size = 8
```
