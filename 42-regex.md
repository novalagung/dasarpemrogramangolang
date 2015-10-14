# Regexp

Regexp atau *regular expression* adalah suatu teknik yang digunakan untuk mencocokan string dengan pola tertentu. Regexp biasa dimanfaatkan untuk pencarian dan pengubahan data string.

Pencarian string dilakukan dengan mengkombinasikan literal-literal regex. Golang mengadopsi standar regex **RE2**, untuk melihat sintaks yang di-support engine ini bisa langsung merujuk ke dokumentasinya di [https://github.com/google/re2/wiki/Syntax](https://github.com/google/re2/wiki/Syntax).

Pada bab ini kita akan belajar mengenai pengaplikasian regex. Golang menyediakan package `regexp`, berisikan banyak sekali fungsi untuk keperluan regex.

## Penerapan Regexp

Fungsi `regexp.Compile()` digunakan untuk mengkompilasi ekspresi regex yang dimasukkan. Fungsi tersebut mengembalikan objek bertipe `regexp.*Regexp`. 

Berikut merupakan contoh penerapan regex untuk pencarian karakter.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

if err != nil {
    fmt.Println(err.Error())
}

var res1 = regex.FindAllString(text, 2)
// ["banana", "orange juice"]

var res2 = regex.FindAllString(text, -1)
// ["banana", "orange juice", "burger", "soup"]
```

Ekspresi `b[a-zA-Z]+` maknanya adalah, semua string yang merupakan alphabet dan diawali huruf `b`. Ekspresi tersebut di-compile oleh `regexp.Compile()` lalu disimpan ke variabel objek `regex` yang tipenya adalah `regexp.*Regexp`.

Struct `regexp.Regexp` memiliki banyak method, salah satunya adalah `FindAllString()`, yang digunakan untuk mencari semua string yang sesuai dengan ekspresi regex, dengan kembalian berupa array string.

Jumlah hasil pencarian dari `regex.FindAllString()` bisa ditentukan. Contohnya pada `res1`, ditentukan maksimal `2` data saja pada nilai kembalian. Jika batas di set `-1`, maka akan mengembalikan semua data. 

Ada cukup banyak method struct `regexp.*Regexp` yang bisa kita manfaatkan untuk keperluan pengelolaan string. Berikut merupakan pembahasan tiap method-nya.

## Method `MatchString()`

Method ini digunakan untuk mendeteksi apakah string memenuhi sebuah pola regexp.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var isMatch = regex.MatchString(text)
// true
```

Pada contoh di atas `isMatch` bernilai `true` karena string `"banana,orange juice,burger,soup"` memenuhi pola regex `b[a-zA-Z]+`.

## Method `FindString()`

Digunakan untuk mencari string yang memenuhi kriteria regexp yang telah ditentukan.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var str = regex.FindString(text)
// banana
```

Fungsi ini hanya mengembalikan 1 buah hasil saja. Jika ada banyak substring yang sesuai dengan ekspresi regexp, akan dikembalikan yang pertama saja.

## Method `FindStringIndex()`


Digunakan untuk mencari index string kembalian hasil dari operasi regexp.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var idx = regex.FindStringIndex(text)
// [0, 6]

var str = text[0:6]
// "banana"
```

Method ini sama dengan `FindString()` hanya saja yang dikembalikan indeknya.

## Method `FindAllString()`

Digunakan untuk mencari banyak string yang memenuhi kriteria regexp yang telah ditentukan.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var str1 = regex.FindAllString(text, -1)
// ["banana", "burger"]

var str2 = regex.FindAllString(text, 1)
// ["banana"]
```

Jumlah data yang dikembalikan bisa ditentukan. Jika diisi dengan `-1`, maka akan mengembalikan semua data.

## Method `ReplaceAllString()`

Berguna untuk me-replace semua string yang memnuhi kriteri regexp, dengan string lain. 

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var str = regex.ReplaceAllString(text, "potato")
// "potato,potato,orange juice,soup"
```

## Method `ReplaceAllStringFunc()`

Digunakan untuk me-replace semua string yang memnuhi kriteri regexp, dengan kondisi yang bisa ditentukan untuk setiap substring yang akan di replace.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var str = regex.ReplaceAllStringFunc(text, func(each string) string {
    if each == "burger" {
        return "potato"
    }
    return each
})
// "banana,potato,orange juice,soup"
```

Pada contoh di atas, jika salah satu substring yang *match* adalah `"burger"` maka akan diganti dengan `"potato"`, string selainnya tidak di replace.

## Method `Split()`

Digunakan untuk memsih string dengan separator atau pemisah adalah substring yang memnuhi kriteria regexp yang telah ditentukan.

```go
var text = "banana,orange juice,burger,soup"
var regex, err = regexp.Compile(`b[a-zA-Z]+`)

var str = regex.Split(text)
// ["", ",", ",orange juice,soup"]
```
