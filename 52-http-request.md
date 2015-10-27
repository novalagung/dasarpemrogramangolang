# HTTP Request

Setelah di bab sebelumnya kita belajar tentang bagaimana membuat Web API yang menyediakan data dalam bentuk JSON, pada bab ini kita akan belajar mengenai cara untuk mengkonsumsi data tersebut.

Pastikan anda sudah mempraktekkan apa-apa yang ada pada bab sebelumnya (bab 51), karena web api server yang sudah dibahas pada bab tersebut digunakan juga pada bab ini.

![Jalankan web server](images/51_1_server.png)

## Penggunaan HTTP Request

Package `net/http`, selain berisikan tools untuk keperluan pembuatan web, juga berisikan fungsi-fungsi untuk melakukan http request. Salah satunya adalah `http.NewRequest()` yang akan kita bahas di sini.

Sebelumnya, import package yang dibutuhkan. Dan siapkan struct `student` yang nantinya akan dipakai sebagai tipe data reponse dari web API. Struk tersebut skema nya sama dengan yang ada pada bab 51.

```go
package main

import "fmt"
import "net/http"
import "encoding/json"

var baseURL = "http://localhost:8080"

type student struct {
    ID    string
    Name  string
    Grade int
}
```

Setelah itu buat fungsi `fetchUsers()`. Fungsi ini bertugas melakukan request ke [http://localhost:8080/users](http://localhost:8080/users), menerima response dari request tersebut, lalu menampilkannya.

```go
func fetchUsers() []*student {
    var err error
    var client = &http.Client{}
    var data []*student

    request, err := http.NewRequest("POST", baseURL+"/users", nil)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }

    response, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }

    return data
}
```

Statement `&http.Client{}` menghasilkan instance `http.Client`. Objek ini nantinya diperlukan untuk eksekusi request.

Fungsi `http.NewRequest()` digunakan untuk membuat request baru. Fungsi tersebut memiliki 3 parameter yang wajib diisi.

 1. Parameter pertama, berisikan tipe request **POST** atau **GET** atau lainnya
 2. Parameter kedua, adalah URL tujuan request
 3. Parameter ketiga, payload request (jika ada)

Fungsi tersebut menghasilkan instance bertipe `http.Request`. Objek tersebut nantinya disisipkan pada saat eksekusi request.

Cara eksekusi request sendiri adalah dengan memanggil method `Do()` pada instance `http.Client` yang sudah dibuat, dengan parameter adalah instance request-nya. Contohnya seperti pada `client.Do(request)`.

Method tersebut mengembalikan instance bertipe `http.Response`, yang didalamnya berisikan informasi yang dikembalikan dari web API.

Data response bisa diambil lewat property `Body` dalam bentuk string. Gunakan JSON Decoder untuk mengkonversinya menjadi bentuk JSON. Contohnya bisa dilihat di kode di atas, `json.NewDecoder(response.Body).Decode(&data)`. Setelah itu barulah kita bisa menampilkannya.

Perlu diketahui, data response perlu di-**close** setelah tidak dipakai. Caranya seperti pada kode `defer response.Body.Close()`.

Implementasikan fungsi `fetchUsers()` di atas pada `main`.

```go
func main() {
    var users = fetchUsers()

    for _, each := range users {
        fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
    }
}
```

Jalankan program untuk mengetes hasilnya.

![HTTP Request](images/52_1_http_request.png)

## HTTP Request Dengan Payload

Untuk menyisipkan data pada sebuah request, ada beberapa hal yang perlu ditambahkan. Yang pertama, import beberapa package lagi, `bytes` dan `net/url`.

```go
import "bytes"
import "net/url"
```

Buat fungsi baru dengan isi adalah sebuah request ke [http://localhost:8080/user](http://localhost:8080/user) dengan data yang disisipkan adalah `ID`.


```go
func fetchUser(ID string) student {
    var err error
    var client = &http.Client{}
    var data student

    var param = url.Values{}
    param.Set("id", ID)
    var payload = bytes.NewBufferString(param.Encode())

    request, err := http.NewRequest("POST", baseURL+"/user", payload)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    response, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println(err.Error())
        return data
    }

    return data
}
```

Isi fungsi di atas bisa dilihat memiliki beberapa kemiripan dengan fungsi `fetchUsers()` sebelumnya. 

Statement `url.Values{}` akan menghasilkan objek yang nantinya digunakan sebagai payload request. Pada objek tersebut perlu di set data apa saja yang ingin dikirimkan menggunakan fungsi `Set()` seperti pada `param.Set("id", ID)`.

Statement `bytes.NewBufferString(param.Encode())` maksudnya, objek payload di-encode lalu diubah menjadi bentuk `bytes.Buffer`, yang nantinya disisipkan pada parameter ketiga pemanggilan fungsi `http.NewRequest()`.

Karena data yang akan dikirim di-encode, maka pada header perlu di set tipe konten request-nya. Kode `request.Header.Set("Content-Type", "application/x-www-form-urlencoded")` artinya tipe konten request di set sebagai `application/x-www-form-urlencoded`.

> Pada konteks HTML, HTTP Request yang di trigger dari tag `<form></form>` secara default tipe konten-nya sudah di set `application/x-www-form-urlencoded`. Lebih detailnya bisa merujuk ke spesifikasi HTML form [http://www.w3.org/TR/html401/interact/forms.html#h-17.13.4.1](http://www.w3.org/TR/html401/interact/forms.html#h-17.13.4.1)

Response dari rute `/user` bukan berupa array objek, melainkan sebuah objek. Maka pada saat decode pastikan tipe variabel penampung hasil decode data response adalah `student` bukan `[]*student`.

Terakhir buat implementasinya pada fungsi `main`.

```go
func main() {
    var user1 = fetchUser("E001")
    fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}
```

Pada kode di atas `ID` ditentukan nilainya adalah `"E001"`. Jalankan program untuk mengetes apakah data yang dikembalikan sesuai.

![HTTP request payload](images/52_2_http_request_payload.png)
