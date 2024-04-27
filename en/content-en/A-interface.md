# A.27. Interface

Interface adalah definisi suatu kumpulan method yang tidak memiliki isi, jadi hanya definisi header/schema-nya saja. Kumpulan method tersebut ditulis dalam satu block interface dengan nama tertentu.

Interface merupakan tipe data. Objek bertipe interface memiliki zero value yaitu `nil`. Variabel bertipe interface digunakan untuk menampung nilai objek konkret yang memiliki definisi method minimal sama dengan yang ada di interface.

## A.27.1. Penerapan Interface

Untuk menerapkan interface, pertama siapkan deklarasi tipe baru menggunakan keyword `type` dan tipe data `interface` lalu siapkan juga isinya (definisi method-nya).

```go
package main

import "fmt"
import "math"

type hitung interface {
    luas() float64
    keliling() float64
}
```

Di atas, interface `hitung` dideklarasikan memiliki 2 buah method yaitu `luas()` dan `keliling()`. Interface ini nantinya digunakan sebagai tipe data pada variabel untuk menampung objek bangun datar hasil dari struct yang akan dibuat.

Dengan adanya interface `hitung` ini, maka perhitungan luas dan keliling bangun datar bisa dilakukan tanpa perlu tahu jenis bangun datarnya sendiri itu apa.

Selanjutnya, siapkan struct bangun datar `lingkaran`, struct ini memiliki definisi method yang sebagian adalah ada di interface `hitung`.

```go
type lingkaran struct {
    diameter float64
}

func (l lingkaran) jariJari() float64 {
    return l.diameter / 2
}

func (l lingkaran) luas() float64 {
    return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
    return math.Pi * l.diameter
}
```

Struct `lingkaran` memiliki tiga buah method yaitu `jariJari()`, `luas()`, dan `keliling()`.

Berikutnya, siapkan struct bangun datar `persegi` berikut:

```go
type persegi struct {
    sisi float64
}

func (p persegi) luas() float64 {
    return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
    return p.sisi * 4
}
```

Perbedaan struct `persegi` dengan `lingkaran` terletak pada method `jariJari()`. Struct `persegi` tidak memiliki method tersebut. Tetapi meski demikian, variabel objek hasil cetakan 2 struct ini akan tetap bisa ditampung oleh variabel cetakan interface `hitung`, karena dua method yang ter-definisi di interface tersebut juga ada pada struct `persegi` dan `lingkaran`, yaitu method `luas()` dan `keliling()`.

Sekarang buat implementasi perhitungan di fungsi `main()`.

```go
func main() {
    var bangunDatar hitung

    bangunDatar = persegi{10.0}
    fmt.Println("===== persegi")
    fmt.Println("luas      :", bangunDatar.luas())
    fmt.Println("keliling  :", bangunDatar.keliling())

    bangunDatar = lingkaran{14.0}
    fmt.Println("===== lingkaran")
    fmt.Println("luas      :", bangunDatar.luas())
    fmt.Println("keliling  :", bangunDatar.keliling())
    fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}
```

Perhatikan kode di atas. Variabel objek `bangunDatar` bertipe interface `hitung`. Variabel tersebut digunakan untuk menampung objek konkrit buatan struct `lingkaran` dan `persegi`.

Dari variabel tersebut, method `luas()` dan `keliling()` diakses. Secara otomatis Golang akan mengarahkan pemanggilan method pada interface ke method asli milik struct yang bersangkutan.

![Pemanfaatan interface](images/A_interface_1_interface.png)

Method `jariJari()` pada struct `lingkaran` tidak akan bisa diakses karena tidak terdefinisi dalam interface `hitung`. Pengaksesannya secara paksa menyebabkan error.

Untuk mengakses method yang tidak ter-definisi di interface, variabel-nya harus di-casting terlebih dahulu ke tipe asli variabel konkritnya (pada kasus ini tipenya `lingkaran`), setelahnya method akan bisa diakses.

Cara casting objek interface sedikit unik, yaitu dengan menuliskan nama tipe tujuan dalam kurung, ditempatkan setelah nama interface dengan menggunakan notasi titik (seperti cara mengakses property, hanya saja ada tanda kurung nya). Contohnya bisa dilihat di kode berikut. Statement `bangunDatar.(lingkaran)` adalah contoh casting pada objek interface.

```go
var bangunDatar hitung = lingkaran{14.0}
var bangunLingkaran lingkaran = bangunDatar.(lingkaran)

bangunLingkaran.jariJari()
```

> Metode casting pada tipe data interface biasa disebut dengan **type assertion**

Perlu diketahui juga, jika ada interface yang menampung objek konkrit yang mana struct-nya tidak memiliki salah satu method yang terdefinisi di interface, maka error akan muncul. Intinya kembali ke aturan awal, variabel interface hanya bisa menampung objek yang minimal memiliki semua method yang terdefinisi di interface tersebut.

## A.27.2. Embedded Interface

Interface bisa di-embed ke interface lain, sama seperti struct. Cara penerapannya juga sama, cukup dengan menuliskan nama interface yang ingin di-embed ke dalam body interface tujuan.

Pada contoh berikut, disiapkan interface bernama `hitung2d` dan `hitung3d`. Kedua interface tersebut kemudian di-embed ke interface baru bernama `hitung`.

```go
package main

import "fmt"
import "math"

type hitung2d interface {
    luas() float64
    keliling() float64
}

type hitung3d interface {
    volume() float64
}

type hitung interface {
    hitung2d
    hitung3d
}
```

Interface `hitung2d` berisikan method untuk kalkulasi luas dan keliling, sedang `hitung3d` berisikan method untuk mencari volume bidang. Kedua interface tersebut embed ke interface `hitung`, menjadikannya memiliki kemampuan untuk mengakses method `luas()`, `keliling()`, dan `volume()`.

Next, siapkan struct baru bernama `kubus` yang memiliki method `luas()`, `keliling()`, dan `volume()`.

```go
type kubus struct {
    sisi float64
}

func (k *kubus) volume() float64 {
    return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
    return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
    return k.sisi * 12
}
```

Objek hasil cetakan struct `kubus` di atas, nantinya akan ditampung oleh objek cetakan interface `hitung` yang isinya merupakan gabungan interface `hitung2d` dan `hitung3d`.

Terakhir, buat implementasi-nya di fungsi `main()`.

```go
func main() {
    var bangunRuang hitung = &kubus{4}

    fmt.Println("===== kubus")
    fmt.Println("luas      :", bangunRuang.luas())
    fmt.Println("keliling  :", bangunRuang.keliling())
    fmt.Println("volume    :", bangunRuang.volume())
}
```

Bisa dilihat di kode di atas, lewat interface `hitung`, method `luas()`, `keliling()`, dan `volume()` bisa di akses.

Pada chapter [A.23. Pointer](/A-pointer.html) dijelaskan bahwa method pointer bisa diakses lewat variabel objek biasa dan variabel objek pointer. Variabel objek yang dicetak menggunakan struct yang memiliki method pointer, jika ditampung ke dalam variabel interface, harus diambil referensi-nya terlebih dahulu. Contohnya bisa dilihat pada kode di atas `var bangunRuang hitung = &kubus{4}`.

![Embedded interface](images/A_interface_2_embedded_interface.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.27-interface">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.27...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
