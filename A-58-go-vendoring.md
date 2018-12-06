# A.58. Go Vendoring

Pada bagian ini kita akan belajar cara pemanfaatan vendoring untuk mempermudah manajemen package atau dependency dalam project Go.

Pada [Bab A.3. GOPATH Dan Workspace](/3-gopath-dan-workspace.html) sudah disinggung bahwa project Go harus ditempatkan didalam workspace, lebih spesifiknya dalam folder `$GOPATH/src/`. Aturan ini cukup memicu perdebatan di komunitas, karena menghasilkan efek negatif terhadap beberapa hal, yang salah satunya adalah: dependency management yang dirasa susah.

## A.58.1. Perbandingan Project Yang Menerapkan Vendoring vs Tidak

Penulis coba contohkan dengan sebuah kasus untuk mempermudah pembaca memahami permasalahan yang ada dalam dependency management di Go.

Dimisalkan, ada dua buah projek yang sedang di-develop, `project-one` dan `project-two`. Keduanya depend terhadap salah satu 3rd party library yg sama, [gubrak](https://github.com/novalagung/gubrak). Di dalam `project-one`, versi gubrak yang digunakan adalah `v0.9.1-alpha`, sedangkan di `project-two` versi `v1.0.0` digunakan. Pada `project-one` versi yang digunakan cukup tua karena proses pengembangannya sudah agak lama, dan aplikasinya sendiri sudah stabil, jika di upgrade paksa ke gubrak versi `v1.0.0` kemungkinan besar terjadi error dan panic.

Kedua projek tersebut pastinya akan lookup gubrak ke direktori yang sama, yaitu `$GOPATH/src/github.com/novalagung/gubrak`. Efeknya, ketika sedang bekerja pada `project-one`, harus dipastikan current revision pada repository gubrak di lokal adalah sesuai dengan versi `v1.0.0`. Dan, ketika mengerjakan `project-two` maka current revision gubrak harus sesuai dengan versi `v0.9.1-alpha`. Repot sekali bukan?

Setelah beberapa waktu, akhirnya `go1.6` rilis, dengan membawa kabar baik, yaitu rilisnya fasilitas baru **vendoring**. Vendoring ini berguna untuk meng-enkapsulasi packages atau dependencies atau 3rd party libraries yang digunakan dalam spesifik project.

Penggunaannya sendiri sangat mudah, cukup tempatkan 3rd party library ke-dalam folder `vendor`, yang berada di dalam masing-masing project. By default go akan memprioritaskan lookup pada folder `vendor`.

> Folder vendor ini kegunaannya sama seperti folder `node_modules` dalam project javascript.

Kita kembali ke contoh, library gubrak dipindahkan ke folder `vendor` dalam `project-one`, maka struktur project tersebut plus vendoring menjadi seperti berikut.

```bash
$GOPATH/src/project-one
$GOPATH/src/project-one/vendor/github.com/novalagung/gubrak
```

Source code library gubrak dalam folder `vendor` harus ditempatkan sesuai dengan struktur-nya.

```bash
$ tree .
.
├── main.go
└── vendor
    └── github.com
        └── novalagung
            └── gubrak
                ├── date.go
                ├── is.go
                └── ...

4 directories, N files
```

Isi `project-one/main.go` sendiri cukup sederhana, sebuah program kecil yang menampilkan angka random dengan range 10-20.

```go
package main

import (
	"fmt"
	"github.com/novalagung/gubrak"
)

func main() {
	fmt.Println(gubrak.RandomInt(10, 20))
}
```

Ketika di build, dengan flag `-v` terlihat perbedaan antara projek yang menerapkan vendoring maupun yang tidak.

- Untuk yang tidak menggunakan vendor folder, maka akan lookup ke folder library yang digunakan, di `$GOPATH`.
- Untuk project yang menerapkan vendoring, maka tetap lookup juga, tetapi dalam sub folder `vendor` yang ada di dalam projek tersebut.

![Using vendor vs not using vendor](images/A.58_1_vendor_vs_nope.png)

## A.58.2. Manajemen Dependencies Dalam Folder `vendor`

Cara menambahkan dependency ke folder `vendor` bisa dengan cara copy-paste, yang tetapi jelasnya adalah tidak praktis (dan bisa menimbulkan masalah). Cara yang lebih benar adalah dengan menggunakan package management tools.

Di go, ada sangat banyak sekali package management tools. Sangat. Banyak. Sekali. Silakan cek di [PackageManagementTools](https://github.com/golang/go/wiki/PackageManagementTools) untuk lebih detailnya.

Pembaca bebas memilih untuk menggunakan package management tools yang mana. Namun di buku ini, **Dep** akan kita bahas. Dep sendiri merupakan official package management tools untuk Go dari Go Team.

Pembahasan mengenai Dep ada di bab selanjutnya.
