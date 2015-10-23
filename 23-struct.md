# Struct

Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus dengan nama tertentu.

Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti `map`, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

Dengan memanfaatkan struct, data akan terbungkus lebih rapi dan mudah di-maintain.

Struct merupakan cetakan, digunakan untuk mencetak variabel objek (istilah untuk variabel yang memiliki property). Variabel objek memiliki behaviour atau sifat yang sama sesuai struct pencetaknya. Konsep ini sama dengan konsep **class** pada pemrograman berbasis objek. Sebuah buah struct bisa dimanfaatkan untuk mencetak banyak objek. 

## Deklarasi Struct

Keyword `type` digunakan untuk deklarasi struct. Di bawah ini merupakan contoh cara penggunaannya.

```go
type student struct {
    name string
    grade int
}
```

Struct `student` dideklarasikan memiliki 2 property, yaitu `name` dan `grade`. Objek yang dicetak dengan struct ini nantinya akan memiliki sifat yang sama.

## Penerapan Struct

Struct `student` yang sudah disiapkan di atas akan kita manfaatkan untuk mencetak sebuah variabel objek. Property variabel tersebut nantinya diisi kemudian ditampilkan.

```go
func main() {
    var s1 student
    s1.name = "john wick"
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("grade :", s1.grade)
}
```

Cara membuat variabel objek sama seperti pembuatan variabel biasa. Tinggal tulis saja nama variabel diikuti nama struct, contoh: `var s1 student`. 

Semua property variabel objek pada awalnya memiliki nilai default sesuai tipe datanya.

Property variabel objek bisa diakses nilainya menggunakan notasi titik, contohnya `s1.name`. Nilai property-nya juga bisa diubah, contohnya pada kode `s1.grade = 2`.

![Pengaksesan property variabel objek](images/23_1_struct.png)

## Inisialisasi Object Struct

Cara inisialisasi variabel objek adalah dengan menambahkan kurung kurawal setelah nama struct. Nilai masing-masing property bisa diisi pada saat inisialisasi.

Pada contoh berikut, terdapat 3 buah variabel objek yang dideklarasikan dengan cara berbeda.

```go
var s1 = student{}
s1.name = "wick"
s1.grade = 2

var s2 = student{"ethan", 2}

var s3 = student{name: "jason"}

fmt.Println("student 1 :", s1.name)
fmt.Println("student 2 :", s2.name)
fmt.Println("student 3 :", s3.name)
```

Pada kode di atas, variabel `s1` menampung objek cetakan `student`. Vartiabel tersebut kemudian di-set nilai property-nya.

Variabel objek `s2` dideklarasikan dengan metode yang sama dengan `s1`, pembedanya di `s2` nilai propertinya di isi langsung ketika deklarasi. Nilai pertama akan menjadi nilai property pertama (yaitu `name`), dan selanjutnya berurutan.

Pada deklarasi `s3`, dilakukan juga pengisian property ketika pencetakan objek. Hanya saja, yang diisi hanya `name` saja. Cara ini cukup efektif jika digunakan untuk membuat objek baru yang nilai property-nya tidak semua harus disiapkan di awal. Keistimewaan lain menggunakan cara ini adalah penentuan nilai property bisa dilakukan dengan tidak berurutan. Contohnya:

```go
var s4 = student{name: "wayne", grade: 2}
var s5 = student{grade: 2, name: "bruce"}
```

## Variabel Objek Pointer

Objek hasil cetakan struct bisa diambil nilai pointer-nya, dan bisa disimpan pada variabel objek yang bertipe struct pointer. Contoh penerapannya:

```go
var s1 = student{name: "wick", grade: 2}

var s2 *student = &s1
fmt.Println("student 1, name :", s1.name)
fmt.Println("student 4, name :", s2.name)

s2.name = "ethan"
fmt.Println("student 1, name :", s1.name)
fmt.Println("student 4, name :", s2.name)
```

`s2` adalah variabel pointer hasil cetakan struct `student`. `s2` menampung nilai referensi `s1`, mengakibatkan setiap perubahan pada property variabel tersebut, akan juga berpengaruh pada variabel objek `s1`.

Meskipun `s2` bukan variabel asli, property nya tetap bisa diakses seperti biasa. Inilah keunikan variabel objek pointer, tanpa perlu di-dereferensi nilai asli property tetap bisa diakses. Pengisian nilai pada property tersebut juga bisa langsung menggunakan nilai asli, contohnya seperti `s2.name = "ethan"`.

![Variabel objek pointer](images/23_2_pointer_object.png)

## Embedded Struct

**Embedded** struct adalah penurunan properti dari satu struct ke struct lain, sehingga properti struct yang diturunkan bisa digunakan. Agar lebih mudah dipahami, mari kita bahas kode berikut.

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    grade int
    person
}

func main() {
    var s1 = student{}
    s1.name = "wick"
    s1.age = 21
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}

```

Pada kode di atas, disiapkan struct `person` dengan properti yang tersedia adalah `name` dan `age`. Disiapkan juga struct `student` dengan property `grade`. Struct `person` di-embed kedalam struct `student`. Caranya cukup mudah, yaitu dengan menuliskan nama struct yang ingin di-embed ke dalam body `struct` target.

Embedded struct adalah **mutable**, nilai property-nya nya bisa diubah.

Khusus untuk properti yang bukan properti asli (properti turunan dari struct lain), bisa diakses dengan cara mengakses struct *parent*-nya terlebih dahulu. Contoh `s1.person.age`. Nilai yang dikembalikan memiliki referensi yang sama dengan `s1.age`.

## Embedded Struct Dengan Nama Property Yang Sama

Jika salah satu nama properti sebuah struct memiliki kesamaan dengan properti milik struct lain yang di-embed, maka pengaksesan property-nya harus dilakukan dengan jelas. Contoh bisa dilihat di kode berikut.

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    person
    age   int
    grade int
}

func main() {
    var s1 = student{}
    s1.name = "wick"
    s1.age = 21        // age of student
    s1.person.age = 22 // age of person

    fmt.Println(s1.name)
    fmt.Println(s1.age)
    fmt.Println(s1.person.age)
}
```

Struct `person` di-embed ke dalam struct `student`, dan kedua struct tersebut kebetulan salah satu nama property-nya ada yg sama, yaitu `age`. Cara mengakses property `age` milik struct `person` lewat objek struct `student`, adalah dengan menuliskan nama struct yg di-embed kemudian nama property-nya, contohnya: `s1.person.age = 22`.

## Pengisian Nilai Sub-Struct

Pengisian nilai property sub-struct bisa dilakukan dengan langsung memasukkan variabel objek yang tercetak dari struct yang sama.

```go
var p1 = person{name: "wick", age: 21}
var s1 = student{person: p1, grade: 2}

fmt.Println("name  :", s1.name)
fmt.Println("age   :", s1.age)
fmt.Println("grade :", s1.grade)
```

Pada deklarasi `s1`, property `person` diisi variabel objek `p1`.

## Anonymous Struct

Anonymous struct adalah struct yang tidak dideklarasikan di awal, melainkan ketika dibutuhkan saja, langsung pada saat penciptaan objek. Teknik ini cukup efisien untuk pembuatan variabel objek yang struct nya hanya dipakai sekali.

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {
    var s1 = struct {
        person
        grade int
    }{}
    s1.person = person{"wick", 21}
    s1.grade = 2

    fmt.Println("name  :", s1.person.name)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}
```

Pada kode di atas, variabel `s1` langsung diisi objek anonymous struct yang memiliki sebuah property `grade`, dan property lain yang diturunkan dari struct `person`. 

Salah satu aturan yang perlu diingat dalam pembuatan anonymous struct adalah, deklarasi harus diikuti dengan inisialisasi. Bisa dilihat pada `s1` setelah deklarasi struktur struct, terdapat kurung kurawal untuk inisialisasi objek. Meskipun nilai tidak diisikan di awal, kurung kurawal tetap harus ditulis.

```go
// anonymous struct tanpa inisialisasi
var s1 = struct {
    person
    grade int
}{}

// anonymous struct dengan inisialisasi
var s2 = struct {
    person
    grade int
}{
    person: person{"wick", 21},
    grade:  2,
}
```

## Inisialisasi Langsung Array Anonymous Struct

Anonymous struct bisa dijadikan sebagai tipe sebuah array. Dan nilai awalnya juga bisa diinisialisasi langsung pada saat deklarasi. Berikut adalah contohnya:

```go
var allStudents = []struct {
    person
    grade int
}{
    {person: person{"wick", 21}, grade: 2},
    {person: person{"ethan", 22}, grade: 3},
    {person: person{"bond", 21}, grade: 3},
}

for _, student := range allStudents {
    fmt.Println(student)
}
```

## Deklarasi Anonymous Struct Menggunakan Keyword **var**

Cara lain untuk deklarasi anonymous struct adalah dengan menggunakan keyword `var`.

```go
var student struct {
    person
    grade int
}

student.person = person{"wick", 21}
student.grade = 2
```

Statement `type student struct` adalah contoh bagaimana struct dideklrasikan. Maknanya akan berbeda ketika keyword `type` disitu diganti `var`, seperti pada contoh di atas `var student struct`, yang artinya akan dicetak sebuah objek dari anonymous struct dan disimpan pada variabel bernama `student`.

Kelemahan metode ini, nilai tidak bisa diinisialisasi langsung pada saat deklarasi. Contohnya bisa dilihat pada kode di bawah ini.

```go
// dekalrasi saja
var student struct {
    grade int
}

// dekalrasi sekaligus inisialisasi
var student = struct {
    grade int
} {
    12
}
```

## Nested struct

Nested struct adalah anonymous struct yang di-embed ke sebuah struct. Deklarasinya langsung didalam struct peng-embed. Contoh:

```go
type student struct {
    person struct {
        name string
        age  int
    }
    grade   int
    hobbies []string
}
```

Teknik ini biasa digunakan ketika decoding data **json** yang struktur datanya cukup kompleks dengan proses decode hanya sekali.

## Deklarasi Dan Inisialisasi Struct Secara Horizontal

Deklarasi struct bisa dituliskan secara horizontal, caranya bisa dilihat pada kode berikut:

```go
type person struct { name string; age int; hobbies []string }
```

Tanda semi-colon (`;`) digunakan sebagai pembatas deklarasi poperty yang dituliskan secara horizontal. Inisialisasi nilai juga bisa dituliskan dengan metode ini. Contohnya:

```go
var p1 = struct { name string; age int } { age: 22, name: "wick" }
var p2 = struct { name string; age int } { "ethan", 23 }
```

Bagi pengguna editor Sublime yang terinstal plugin GoSublime didalamnya, dekalrasi dengan cara ini tidak bisa dilakukan, karena setiap kali save isi file program akan dirapikan. Jadi untuk mengetesnya bisa dengan menggunakan editor lain.

## Tag property dalam struct

Tag merupakan informasi opsional yang bisa ditambahkan pada masing-masing property struct. Cara penggunaannya:

```go
type person struct {
    name string `tag1`
    age  int    `tag2`
}
```

Tag biasa dimanfaatkan untuk keperluan encode/decode data json. Informasi tag juga bisa diakses lewat reflect. Nantinya akan ada pembahasan yang lebih detail mengenai pemanfaatan tag dalam struct, yaitu ketika sudah masuk bab json.
