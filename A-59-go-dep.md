# A.58. Dep - Go Dependency Management Tool

Dep adalah Go Official Package Management Tools, diperkenalkan pada go versi `go1.9`. Pada bab ini kita akan belajar penggunaannya.

Dengan menggunakan Dep, maka semua 3rd party libraries yang dipakai dalam suatu project akan dicatat dan ditempatkan di dalam sub folder `vendor` pada project tersebut.

## A.58.1. Instalasi Dep

Silakan ikuti petunjuk di [Dep Installation](https://golang.github.io/dep/docs/installation.html), ada beberapa cara yang bisa dipilih untuk meng-install Dep.

## A.58.2. Inisialisasi Dep Pada Project

Buat project baru, isi dengan kode sederhana untuk menampilkan angka random (sama seperti pada bab sebelumnya). Library yang kita gunakan adalah [gubrak](https://github.com/novalagung/gubrak).

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

Jalankan command `dep init` untuk meng-inisialisasi project agar ditandai sebagai project yang menggunakan package manager Dep. Command tersebut menghasilkan 3 buah file/folder baru.

- File `Gopkg.toml`, berisikan metada semua dependencies yang digunakan dalam project.
- File `Gopkg.lock`, isinya mirip seperti `Gopkg.toml` hanya saja lebih mendetail informasi tiap-tiap dependency yang disimpan dalam file ini.
- Folder `vendor`, akan berisi source code dari pada 3rd party yang digunakan.

Eksekusi command `dep init` akan secara otomatis diikuti dengan eksekusi `dep ensure`. Command `dep ensure` sendiri merupakan command paling penting dalam Dep, gunanya untuk sinkronisasi dan memastikan bahwa semua 3rd party libraries digunakan dalam project metadata-nya ter-mapping dengan baik dan benar dalam `Gopkg.*`, dan source code 3rd party sendiri ada dalam `vendor`.

```bash
$ tree -L 1 .
.
├── Gopkg.lock
├── Gopkg.toml
├── main.go
└── vendor

1 directory, 3 files
```

Selesai. Sesederhana ini. Cukup mudah bukan?

## A.58.3. Add & Update Dependency

Gunakan `dep ensure -add <library>` untuk menambahkan library. Gunakan `dep ensure -update <library>` untuk meng-update library yang sudah tercatat.

Penggunaan dua command tersebut jelasnya akan mengubah informasi dalam file `Gopkg.*`. Pastikan bahwa library yang didaftarkan adalah yang memang digunakan di project. Jika ragu, maka hapus saja folder `vendor`, lalu eksekusi command `dep ensure` untuk sinkronisasi ulang.

Contoh command update:

```bash
$ dep ensure -update github.com/novalagung/gubrak
```
