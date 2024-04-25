# A.6. Command

Pengembangan aplikasi Go pastinya tak akan jauh dari hal-hal yang berbau CLI atau *Command Line Interface*. Di Go, proses inisialisasi project, kompilasi, testing, eksekusi program, semuanya dilakukan lewat command line.

Go menyediakan command `go`, dan pada chapter ini kita akan mempelajari beberapa di antaranya.

> Pada pembelajaran chapter ini, pembaca tidak harus menghafal dan mempraktekan semuanya, cukup ikuti saja pembelajaran agar mulai familiar. Perihal prakteknya sendiri akan dimulai pada chapter selanjutnya, yaitu [A.7. Program Pertama: Hello World](/A-hello-world.html).

## A.6.1. Command `go mod init`

*Command* `go mod init` digunakan untuk inisialisasi project pada Go yang menggunakan Go Modules. Untuk nama project bisa menggunakan apapun, tapi umumnya disamakan dengan nama direktori/folder.

Nama project ini penting karena nantinya berpengaruh pada *import path sub packages* yang ada dalam project tersebut.

```
mkdir <nama-project>
cd <nama-project>
go mod init <nama-project>
```

## A.6.2. Command `go run`

*Command* `go run` digunakan untuk eksekusi file program, yaitu file yang ber-ekstensi `.go`. Cara penggunaannya dengan menuliskan *command* tersebut diikuti argumen nama file.

Berikut adalah contoh penerapan `go run` untuk eksekusi file program `main.go` yang tersimpan di path `project-pertama` yang path tersebut sudah diinisialisasi menggunakan `go mod init`.

```bash
cd project-pertama
go run main.go
```

![Eksekusi file program menggunakan `go run`](images/A_go_command_1_go_run.png)

*Command* `go run` hanya bisa digunakan pada file yang nama package-nya adalah `main`. Lebih jelasnya dibahas pada chapter selanjutnya, yaitu ([A.7. Program Pertama: Hello World](/A-hello-world.html)).

Jika ada banyak file yang package-nya `main` dan file-file tersebut berada pada satu direktori level dengan file utama, maka eksekusinya adalah dengan menuliskan semua file sebagai argument *command* `go run`. Contohnya bisa dilihat pada kode berikut.

```bash
go run main.go library.go
```

## A.6.3. Command `go test`

Go menyediakan package `testing`, berguna untuk keperluan pembuatan file test. Pada penerapannya, ada aturan yang wajib diikuti yaitu nama file test harus berakhiran `_test.go`.

Berikut adalah contoh penggunaan *command* `go test` untuk testing file `main_test.go`.

```bash
go test main_test.go
```

![Unit testing menggunakan `go test`](images/A_go_command_3_go_test.png)

## A.6.4. Command `go build`

*Command* ini digunakan untuk mengkompilasi file program.

Sebenarnya ketika eksekusi program menggunakan `go run` didalamnya terjadi proses kompilasi juga. File hasil kompilasi kemudian disimpan pada folder temporary untuk selanjutnya langsung dieksekusi.

Berbeda dengan `go build`, *command* ini menghasilkan file *executable* atau *binary* pada folder yang sedang aktif. Contoh praktiknya bisa dilihat di bawah ini.

![Kompilasi file program menghasilkan file executable](images/A_go_command_4_go_build.png)

Di contoh, project `project-pertama` di-build, hasilnya adalah file baru bernama `project-pertama.exe` berada di folder yang sama. File *executable* tersebut kemudian dieksekusi.

*Default* nama file binary atau executable adalah sesuai dengan nama project. Untuk mengubah nama file executable, gunakan flag `-o`. Contoh:

```
go build -o <nama-executable>
go build -o program.exe
```

> Khusus untuk sistem operasi non-windows, tidak perlu menambahkan akhiran `.exe` pada nama *binary*

## A.6.5. Command `go get`

*Command* `go get` digunakan untuk men-download package atau *dependency*. Sebagai contoh, penulis ingin men-download package Kafka driver untuk Go pada project `project-pertama`, maka command-nya kurang lebih seperti berikut:

```bash
cd project-pertama
go get github.com/segmentio/kafka-go
```

![Download package menggunakan `go get`](images/A_go_command_6_go_get.png)

Pada contoh di atas, bisa dilihat bahwa URL `github.com/segmentio/kafka-go` merupakan URL package kafka-go. Package yang sudah terunduh tersimpan dalam temporary folder yang ter-link dengan project folder di mana *command* `go get` dieksekusi, menjadikan project tersebut bisa meng-*import* package yang telah di-download.

Untuk mengunduh package/dependency versi terbaru, gunakan flag `-u` pada command `go get`, contohnya:

```
go get -u github.com/segmentio/kafka-go
```

Command `go get` **harus dijalankan dalam folder project**. Jika dijalankan di-luar path project maka dependency yang ter-unduh akan ter-link dengan GOPATH, bukan dengan project.

## A.6.6. Command `go mod download`

*Command* `go mod download` digunakan untuk men-download dependency.

## A.6.7. Command `go mod tidy`

*Command* `go mod tidy` digunakan untuk memvalidasi dependency sekaligus men-download-nya jika memang belum ter-download.

## A.6.8. Command `go mod vendor`

Command ini digunakan untuk vendoring. Lebih detailnya akan dibahas di akhir serial chapter A, pada chapter [A.61. Go Vendoring](/A-go-vendoring.html).

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
