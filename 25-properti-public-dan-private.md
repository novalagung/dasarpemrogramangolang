# Property Public Dan Private

Bab ini membahas mengenai *modifier* public dan private dalam Golang. Kapan sebuah struct, fungsi, atau method bisa diakses dari package lain dan kapan tidak.

## Package Public & Private

Pengembangan aplikasi dalam *real development* pasti membutuhkan banyak sekali file program. Dan tidak mungkin semuanya di set sebagai package `main`. Dengan pertimbangan tersebut biasanya file-file tersebut dipisah sebagai package baru.

Folder proyek selain berisikan file-file `.go` juga bisa berisikan folder. Subfolder tersebut nantinya akan menjadi package baru. Di Golang, 1 subfolder adalah 1 package (kecuali package main yang berada langsung didalam folder proyek). Menjadikan file yang ada di dalam suatu folder memiliki nama package berbeda dengan di folder lainnya. Bahkan antara folder dan subfolder juga bisa memiliki nama package yang berbeda.

Fungsi, struct, dan variabel yang dibuat di package lain, jika diakses dari package `main` caranya tidak seperti biasanya. Perlu adanya penentuan hak akses yang tepat (apakah public atau private) agar kode tidak kacau balau. Package public, berarti bisa diakses dari package berbeda, sedangkan private berarti hanya bisa di akses dari package yang sama.

Penentuan level akses atau modifier sendiri di golang sangat mudah, ditandai dengan **character case** nama fungsi/struct/variabel yang ingin di akses. Ketika namanya diawali dengan huruf kapital menandakan bahwa modifier-nya public. Dan sebaliknya, jika diawali huruf kecil, berarti private.

## Penggunaan Package, Import, Dan Modifier Public & Private

Agar lebih mudah dipahami, maka langsung saja kita praktekan.

Pertama buat folder proyek baru bernama `belajar-golang-level-akses` dalam folder `$GOPATH/src`, lalu buat file baru bernama `main.go` didalamnya. File ini kita tentukan package-nya sebagai **main**.

Selanjutnya buat folder baru didalam folder yang sudah dibuat dengan nama `library`, isinya sebuah file bernama `library.go`. File ini ditentukan package-nya adalah **library**.

![Struktur folder dan file](images/25_1_folder_structure.png)

Buka file `library.go` lalu isi dengan kode berikut.

```go
package library

import "fmt"

func SayHello() {
    fmt.Println("hello")
}

func introduce(name string) {
    fmt.Println("nama saya", name)
}
```

File `library.go` yang telah dibuat ditentukan nama package-nya adalah `library` (sesuai dengan nama folder). Dalam package tersebut terdapat dua fungsi: `SayHello()` dan `introduce()`. Fungsi `SayHello()` adalah publik, bisa diakses dari package lain. Sedang fungsi `introduce()` adalah private, ditandai dengan huruf kecil di huruf pertama nama fungsi.

Selanjutnya akan di-tes apakah memang fungsi yang ber-modifier private tidak bisa diakses dari package lain. Buka file `main.go`, lalu tulis kode berikut.

```go
package main

import "belajar-golang-level-akses/library"

func main() {
    library.SayHello()
    library.introduce("ethan")
}
```

Package yang telah dibuat di-import ke dalam package `main`. Pada saat import, ditulis dengan `"belajar-golang-level-akses/library"` karena lokasi foldernya merupakan subfolder dari proyek `belajar-golang-level-akses`. Dengan ini fungsi-fungsi dalam package tersebut bisa digunakan.

```go
library.SayHello()
library.introduce("ethan")
```

Cara pemanggilan fungsi yang berada dalam package lain adalah dengan menulis nama package target diikut dengan nama fungsi menggunakan *dot notation* atau tanda titik.

OK, sekarang coba jalankan kode yang sudah disiapkan di atas. Harusnya akan ada error seperti pada gambar di bawah ini.

![Error saat menjalankan program](images/25_2_error.png)

Error tersebut disebabkan karena fungsi `introduce()` yang berada di package `library` adalah **private**, fungsi jenis ini tidak bisa diakses dari package lain (pada kasus ini `main`). Agar fungsi tersebut bisa diakses, solusinya bisa dengan menjadikannya public, atau diubah cara pemanggilannya. Disini kita menggunakan cara ke-2.

Tambahkan parameter `name` pada fungsi `SayHello()`, lalu panggil fungsi `introduce()` dengan menyisipkan parameter `name` dari dalam fungsi `SayHello()`.

```go
func SayHello(name string) {
    fmt.Println("hello")
    introduce(name)
}
```

Lalu pada main, cukup panggil fungsi `SayHello()` saja, jangan lupa menyisipkan pesan string sebagai parameter-nya.

```go
func main() {
    library.SayHello("ethan")
}
```

Jika sudah, coba jalankan lagi, harusnya error sudah lenyap.

![Contoh penerapan pemanggilan fungsi dari package berbeda](images/25_2_success.png)

## Penggunaan Public & Private Pada Struct Dan Propertinya

Modifier private & public bisa diterapkan di fungsi, struct, method, maupun property variabel. Dan cara penggunaannya sama seperti pada contoh di atas, yaitu dengan menentukan *case* atau huruf pertama dari nama, apakah huruf tersebut besar atau kecil.

Belajar tentang level akses di Golang akan lebih cepat jika langsung praktek. Oleh karena itu langsung saja kita praktekan. Hapus isi file `library.go`, lalu siapkan struct dengan nama `student` didalamnya.

```go
package library

type student struct {
    Name  string
    grade int
}
```

Buat contoh sederhana penerapan struct di atas pada file `main.go`.

```go
package main

import "belajar-golang-level-akses/library"
import "fmt"

func main() {
    var s1 = library.student{"ethan", 21}
    fmt.Println("name ", s1.Name)
    fmt.Println("grade", s1.grade)
}
```

Setelah itu jalankan program.

![Error saat menjalankan program](images/25_3_error.png)

Error muncul ketika program dijalankan. Penyebabnya adalah karena struct `student` masih di set sebagai private. Ganti menjadi public (dengan cara mengubah huruf awalnya menjadi huruf besar) lalu jalankan.

```go
// pada library/library.go
type Student struct {

// pada main.go
var s1 = library.Student{"ethan", 21}
```

Hasilnya:

![Error lain muncul saat menjalankan program](images/25_4_error.png)

Error masih tetap muncul, tapi kali ini berbeda. Error yang baru ini disebabkan karena salah satu properti dari struct `Student` bermodifier private. Properti yg dimaksud adalah `grade`. Ubah menjadi public, lalu jalankan lagi.

```go
// pada library/library.go
Grade int

// pada main.go
fmt.Println("grade", s1.Grade)
```

Dari contoh program di atas, bisa disimpulkan bahwa untuk menggunakan `struct` yang berada di package lain, selain nama stuct-nya harus bermodifier public, properti yang diakses juga harus public.

![Contoh penerapan pemanfaatan struct dan propertynya dari package berbeda](images/25_4_success.png)

## Import Dengan Tanda Titik

Seperti yang kita tahu, untuk mengakses fungsi/struct/variabel yg berada di package lain, nama package nya perlu ditulis, contohnya seperti pada penggunaan penggunaan `library.Student` dan `fmt.Println()`.

Di Golang, package bisa di-import setara dengan file peng-import, caranya dengan menambahkan titik pada saat penulisan keyword `import`. Maksud dari setara disini adalah, semua properti di package lain yg di-import bisa diakses tanpa perlu menuliskan nama package, seperti ketika mengakses sesuatu dari file yang sama.

```go
import (
    . "belajar-golang-level-akses/library"
    "fmt"
)

func main() {
    var s1 = Student{"ethan", 21}
    fmt.Println("name ", s1.Name)
    fmt.Println("grade", s1.Grade)
}
```

Pada kode di atas package `library` di-import menggunakan tanda titik. Dengan itu, pemanggilan struct `Student` tidak perlu dengan menuliskan nama package nya.

## Pemanfaatan Alias Ketika Import Package

Fungsi yang berada di package lain bisa diakses dengan cara menuliskan nama-package diikuti nama fungsi-nya, contohnya seperti `fmt.Println()`. Package yang sudah di-import tersebut bisa diubah namanya dengan cara menggunakan alias pada saat import. Contohnya bisa dilihat pada kode berikut.

```go
import (
    f "fmt"
)

func main() {
    f.Println("Hello World!")
}
```

Disiapkan alias untuk package `fmt` dengan nama `f`. Dengan ini untuk mengakses fungsi `Println()` cukup dengan `f.Println()`.

## Mengakses Properti Dalam File Yang Package-nya Sama

Jika properti yang ingin di akses masih dalam satu package tapi file nya saja yg berbeda, cara mengaksesnya bisa langsung dengan memanggil namanya. Hanya saja ketika eksekusi, file-file lain yang yang nama package-nya sama juga ikut dipanggil.

Langsung saja kita praktekan, buat file baru dalam `belajar-golang-level-akses` dengan nama `partial.go`.

![File `partial.go` disiapkan setara dengan file `main.go`](images/25_5_structure.png)

File `partial.go` ditentukan packagenya adalah **main** (sama dengan package `main.go`). Selanjutnya tulis kode berikut pada file tersebut.

```go
package main

import "fmt"

func sayHello(name string) {
    fmt.Println("halo", name)
}
```

Hapus semua isi file `main.go`, lalu silakan tulis kode berikut.

```go
package main

func main() {
    sayHello("ethan")
}
```

Sekarang terdapat 2 file berbeda yang package-nya sama-sama `main`, yaitu `main.go` dan `partial.go`. Pada saat **go build** atau **go run**, semua file yang nama package-nya `main` tersebut harus dituliskan sebagai argumen.

```
$ go run main.go partial.go
```

Fungsi `sayHello` pada file `partial.go` bisa dikenali meski level aksesnya adalah private. Hal ini karena kedua file tersebut (`main.go` dan `partial.go`) memiliki package yang sama.

![Pemanggilan fungsi private dari dalam package yang sama](images/25_6_multi_main.png)
