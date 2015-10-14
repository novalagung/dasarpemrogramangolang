# Encode - Decode Base64

Di bab ini kita akan belajar tentang **encode** dan **decode** data menggunakan `base64`.

Golang memiliki package `encoding/base64`, yang berisikan fungsi-fungsi untuk kebutuhan *encode-decode* data ke base64 dan sebaliknya. 

Data yang akan di-encode harus bertipe `[]byte`, perlu dilakukan casting untuk data-data yang belum sesuai tipenya.

Secara umum ada 3 macam cara untuk encode dan decode data, kita akan pelajari kesemuanya.

## Penerapan Dungsi `EncodeToString` & `DecodeString`

Kedua fungsi ini digunakan untuk encode decode data dari bentuk `string` ke `[]byte` atau sebaliknya.

Berikut adalah contoh penerapan encode dan decode menggunakan `EncodeToString` dan `DecodeString`.

```go
import "encoding/base64"
import "fmt"

func main() {
    var data = "john wick"

    var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("encoded:", encodedString)

    var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
    var decodedString = string(decodedByte)
    fmt.Println("decoded:", decodedString)
}
```

Variabel `data` bertipe `string`, harus di-casting terlebih dahulu kedalam bentuk `[]byte` sebelum bisa di-encode. Nilai balik hasil `base64.StdEncoding.EncodeToString` berupa `string`.

Sedangkan pada fungsi decode `base64.StdEncoding.DecodeString`, data string yang di-decode akan dikembalikan dalam bentuk `[]byte`. Ekspresi `string(decodedByte)` menjadikan data `[]byte` tadi berubah menjadi string.

![Encode & decode data string](images/43_1_encode_decode.png)

## Penerapan Fungsi `Encode` & `Decode`

Kedua fungsi ini digunakan untuk decode dan encode data dari `[]byte` ke `[]byte`. Cara ini cukup panjang karena variabel penyimpan hasil encode maupun decode harus memiliki lebar elemen yang sama dengan hasilnya (yang nilainya bisa dicari menggunakan `EncodedLen` dan `DecodedLen`).

```go
var data = "john wick"

var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
base64.StdEncoding.Encode(encoded, []byte(data))
var encodedString = string(encoded)
fmt.Println(encodedString)

var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
var _, err = base64.StdEncoding.Decode(decoded, encoded)
if err != nil {
    fmt.Println(err.Error())
}
var decodedString = string(decoded)
fmt.Println(decodedString)
```

Ekspresi `base64.StdEncoding.EncodedLen(len(data))` menghasilkan nilai lebar data setelah di encode. Nilai tersebut kemudian digunakan sebagai lebar elemen variabel `encoded`.

Variabel tersebut bertipe `[]byte`, menjadikannya tidak perlu untuk di-pass-by-reference, ketika fungsi `base64.StdEncoding.Encode` dipanggil.

Fungsi `base64.StdEncoding.DecodedLen` pada dasarnya memiliki kegunaan sama dengan `EncodedLen`, dan digunakan untuk keperluan proses encode. Fungsi ini mengembalikan data, yaitu lebar element `[]byte` hasil encoding dan error.

## Encode & Decode Data URL

Khusus untuk data string yang bentuknya URL, akan lebih efektif menggunakan `URLEncoding` dibanding `StdEncoding`.

Cara penerapannya kurang lebih sama, bisa menggunakan metode pertama maupun metode kedua yang sudah dibahas di atas. Cukup dengan mengganti `StdEncoding` dengan `URLEncoding`.

```go
var data = "http://google.com/"

var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
fmt.Println(encodedString)

var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
var decodedString = string(decodedByte)
fmt.Println(decodedString)
```
