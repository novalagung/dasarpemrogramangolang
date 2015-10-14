# Fungsi String

Golang menyediakan package `strings` yang didalamnya tersedia cukup banyak fungsi untuk keperluan pengolahan data bertipe string. Bab ini berisikan pembahasan beberapa fungsi yang ada di dalam package tersebut.

## Fungsi `strings.Contains()`

Dipakai untuk deteksi apakah string (parameter kedua) merupakan bagian dari string lain (parameter pertama). Nilai kembaliannya berupa `bool`.

```go
var isExists = strings.Contains("john wick", "wick")
// true
```

Variabel `isExists` akan berisikan `true`, karena string `"wick"` merupakan bagian dari `"john wick"`.

## Fungsi `strings.HasPrefix()`

Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).

```go
var isPrefix1 = strings.HasPrefix("john wick", "jo") 
// true

var isPrefix2 = strings.HasPrefix("john wick", "wi")
// false
```

## Fungsi `strings.HasSuffix()`

Digunakan untuk deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua).

```go
var isSuffix1 = strings.HasSuffix("john wick", "ic") 
// false

var isSuffix2 = strings.HasSuffix("john wick", "ck")
// true
```

## Fungsi `strings.Count()`

Memiliki kegunaan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama). Nilai kembalian fungsi ini adalah jumlah karakternya.

```go
var howMany = strings.Count("ethan hunt", "t")
// 2
```

Nilai yang dikembalikan `2`, karena pada string `"ethan hunt"` terdapat dua buah karakter `"t"`.

## Fungsi `strings.Index()`

Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).

```go
var index = strings.Index("ethan hunt", "ha")
// 2
```

String `"ha"` berada pada posisi ke `2` dalam string `"ethan hunt"` (indeks dimulai dari 0).

Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:

```go
var index = strings.Index("ethan hunt", "n")
// 4
```

String `"n"` berada pada indeks `4` dan `8`. Yang dikembalikan adalah yang paling kiri (paling kecil), yaitu `4`.

## Fungsi `strings.Replace()`

Fungsi ini digunakan untuk replace/mengganti bagian dari string dengan string tertentu.

Jumlah substring yang di-replace bisa ditentukan, apakah hanya 1 string pertama, 2 string, atau kesemuanya.

```go
var text = "banana"
var find = "a"
var replaceWith = "o"

var newText1 = strings.Replace(text, find, replaceWith, 1)
// "bonana"

var newText2 = strings.Replace(text, find, replaceWith, 2)
// "bonona"

var newText3 = strings.Replace(text, find, replaceWith, -1)
// "bonono"
```

Pada contoh di atas, substring `"a"` pada string `"banana"` akan di-replace dengan string `"o"`.

Pada `newText1`, hanya 1 huruf `o` saja yang tereplace karena maksimal substring yang ingin di-replace ditentukan 1. 

Angka `-1` akan menjadikan proses replace berlaku pada semua substring. Contoh bisa dilihat pada `newText3`.

## Fungsi `strings.Repeat()`

Digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).

```go
var str = strings.Repeat("na", 4)
// "nananana"
```

Pada contoh di atas, string `"na"` diulang sebanyak 4 kali. Hasilnya adalah: `"nananana"`

## Fungsi `strings.Split()`

Digunakan untuk memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya berupa array string.

```go
var string1 = strings.Split("the dark knight", " ")
// ["the", "dark", "knight"]

var string2 = strings.Split("batman", "")
// ["b", "a", "t", "m", "a", "n"]
```

String `"the dark knight"` dipisah menggunakan pemisah string spasi `" "`, hasilnya kemudian ditampung oleh `string1`.

Untuk memisah string menjadi array tiap 1 string, gunakan pemisah string kosong `""`. Bisa dilihat contohnya pada variabel `string2`.

## Fungsi `strings.Join()`

Memiliki kegunaan berkebalikan dengan `strings.Split()`. Digunakan untuk menggabungkan array string (parameter pertama) menjadi sebuah string dengan pemisah tertentu (parameter kedua.

```go
var data = []string{"banana", "papaya", "tomato"}
var str = strings.Join(data, "-")
// "banana-papaya-tomato"
```

Array `data` digabungkan menjadi satu dengan pemisah tanda *dash* (`-`).

## Fungsi `strings.ToLower()`

Mengubah huruf-huruf string menjadi huruf kecil.

```go
var str = strings.ToLower("aLAy")
// "alay"
```

## Fungsi `strings.ToUpper()`

Mengubah huruf-huruf string menjadi huruf besar.

```go
var str = strings.ToUpper("eat!")
// "EAT!"
```

