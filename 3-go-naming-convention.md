# A.3.C. Go Naming Convention
Golang naming convention jika diterjemahkan kira-kira tata cara penamaan. Seperti pemprograman yang lain, golang juga memiliki tata cara penamaan untuk path/folder/direktori, nama file golang, nama package golang dan nama fungsi golang. Salah satu tujuan dari mengikuti naming convention adalah supaya aplikasi kita mengikuti standart dan mudah dibaca oleh orang lain.

# A.3.C.1 Struktur path / folder
Struktur path / folder biasanya digunakan untuk mengelompokkan fungsi menjadi satu. Jika anda sudah terbiasa dengan konsep MVC (Model, View, Controller) maka fungsi fungsi yang sejenis misal semua fungsi yang berhubungan dengan controller dijadikan dalam satu folder.

Aturan penamaan struktur path / folder:
-	Nama path/folder berupa huruf kecil tanpa under score contoh dasarpemrogramangolang

# A.3.C.2 Package
Struktur aplikasi golang diorganisasikan / dikelompokkan berdasarkan package. package akan menjadi acuan / pedoman untuk memanggil fungsi yang berada dalam package tersebut.

Aturan penamaan package :
-	Nama package harus berupa huruf kecil tanpa under score 
-	Name package diusahakan pendek dan mengandung kata benda missal http, token
-	Nama package bisa berisi singkatan yang merepresentasikan kata benda misal package fmt     yang artinya formatted I/O
-	Nama package biasanya sama dengan nama folder dimana package itu berada
-	Jika dalam satu folder ada banyak file golang maka tiap tiap file memiliki nama package yang sama
-	Jangan membuat nama package yang terlalu umum misal package util, common dll. Detailkan nama package menjadi lebih spesifik. contoh :

package yang umum dengan nama util.
```
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}
```
koding pemanggil menjadi :
```
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```
detailkan package menjadi seperti berikut :
```
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}
```
koding pemanggil menjadi :
```
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

# A.3.C.3 Nama file golang
Nama file aplikasi golang sebenarnya tidak begitu berarti, dalam artian nama file golang tidak menjadi acuan untuk memanggil fungsi-fungsi yang ada dalam file tersebut. Sebenarnya golang tidak ada aturan khusus dalam penamaan file, tetapi untuk memudahkan dalam pembacaan nama file golang dibuatlah beberapa aturan umum.

Aturan penamaan nama file golang :
-	Nama file golang diusahakan berupa huruf kecil semua
-	Jika nama file lebih dari satu suku kata sebaiknya dipisahkan oleh dash(-) misal persegi-panjang.go,  service-transfer.go dan lain lain
-   Nama file golang yang merepresentasikan hal tertentu biasanya ada underscore contoh persegi_test.go (file ini bertujuan untuk unit testing), main_linux.go (file main ini untuk running di linux).

# A.3.C.3 Fungsi
Fungsi adalah sentral / jantung dari golang. Semua proses logic dilakukan dalam fungsi.

Aturan penamaan fungsi pada golang :
-	Fungsi tidak boleh dimulai dengan angka
-	Nama fungsi harus diawali dengan huruf alfabet, setalah diawali dengan huruf alfabet berikutnya dapat diikuti dengan angka atau huruf.  Contoh nama fungsi adalah Sum12.
-	Nama fungsi tidak boleh ada spasi
-	Jika fungsi diawali dengan huruf besar berarti fungsi tersebut dapat diexport ke luar package. Jika fungsi berawalan huruf kecil berarti fungsi tersebut tidak bisa dieksport ke luar package tetapi dapat dipanggil oleh fungsi lain dalam satu package.
-	Jika nama fungsi terdiri dari banyak suku kata maka tiap kata setelah kata pertama sebaiknya berupa huruf capital contoh hitungLuas, hitungVolume.
-	Nama fungsi adalah case sensitive (beda antara huruf kecil dan huruf besar) contoh hitungluas vs hitungLuas vs HitungLuas.

Contoh dari fungsi :

Kodingan persegi untuk menghitung segala sesuatu yang berhubungan dengan persegi misal luas, keliling, sisi, dll.
```
//nama packagenya adalah persegi
package persegi

import "fmt"

//exported function karena fungsi diawali huruf besar
func KelilingPersegi(sisi int) {
	kililing := sisi * 4
	fmt.Printf("Keliling = %d \n", kililing)
	printEnd()
}

//unexported fucntion karena fungsi diawali huruf kecil
func printEnd() {
	fmt.Println("End")
}
```

Kodingan fungsi pemanggil package persegi :
```
fmt.Println("Starting")
//memanggil fungsi keliling
persegi.KelilingPersegi(5)
```




