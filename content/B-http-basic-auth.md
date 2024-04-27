# B.18. HTTP Basic Authentication

HTTP Basic Auth adalah salah satu spesifikasi yang mengatur otentikasi pada HTTP request. Metode ini mewajibkan client request untuk menyertakan username dan password dalam header request. Dengan menerapkan basic auth maka kita tidak perlu menggunakan token untuk mendapatkan session.

> Lebih jelasnya mengenai spesifikasi Basic Auth bisa di lihat di[RFC-7617](https://tools.ietf.org/html/rfc7617)

Informasi username dan password harus di-encode terlebih dahulu ke format yg sudah ditentukan sesuai spesifikasi, kemudian dijadikan value dari header `Authentication`.

Berikut adalah contoh format penulisan basic auth:

```js
// Request header
Authorization: Basic c29tZXVzZXJuYW1lOnNvbWVwYXNzd29yZA==
```

Informasi disisipkan dalam request header dengan key `Authorization`, dan value adalah `Basic` diikut karakter spasi dan hasil encode terhadap data username dan password. Data username dan password digabung dengan separator tanda titik dua (`:`) lalu di-encode dalam format encoding Base64.

```js
// Username password encryption
base64encode("someusername:somepassword")
// Hasilnya adalah c29tZXVzZXJuYW1lOnNvbWVwYXNzd29yZA==
```

Go menyediakan fasilitas untuk mengambil informasi basic auth dari suatu HTTP request dengan mudah, tanpa perlu untuk memparsing header request terlebih dahulu secara manual.

## B.18.1. Struktur Folder Proyek dan Endpoint

Pada chapter ini kita akan membuat sebuah web service sederhana, isinya hanya satu buah endpoint. Endpoint ini didesain untuk bisa menerima query parameter atau tanpa query parameter.

 - Endpoint `/student` menghasilkan response berisi semua data siswa
 - Endpoint `/student?id=s001` menghasilkan response berisi data siswa sesuai dengan id yang di minta

Data siswa sendiri merupakan slice object yang disimpan di variabel global.

OK, langsung saja kita praktekan. Siapkan 3 buah file berikut, tempatkan dalam satu folder proyek.

![Project structure](images/B_http_basic_auth_1_structure.png)

## B.18.2. Routing

Buka `main.go`, isi dengan kode berikut.

```go
package main

import "net/http"
import "fmt"
import "encoding/json"

func main() {
    http.HandleFunc("/student", ActionStudent)
    
    server := new(http.Server)
    server.Addr = ":9000"

    fmt.Println("server started at localhost:9000")
    server.ListenAndServe()
}
```

Lalu siapkan handler untuk rute `/student`.

```go
func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if !Auth(w, r)         {
        return
    }
    if !AllowOnlyGET(w, r) {
        return
    }

    if id := r.URL.Query().Get("id"); id != "" {
        OutputJSON(w, SelectStudent(id))
        return
    }

    OutputJSON(w, GetStudents())
}
```

Di dalam rute `/student` terdapat beberapa validasi.

 - Validasi `!Auth(w, r)`; Nantinya kita siapkan fungsi `Auth()` yang gunanya adalah untuk mengecek apakah request merupakan valid basic auth request atau tidak.
 - Validasi `!AllowOnlyGET(w, r)`; Akan dibuat juga fungsi `AllowOnlyGET()`, tugasnya memastikan hanya request dengan method `GET` yang diperbolehkan masuk.

Setelah request lolos dari 2 validasi di atas, lanjut ke pengecekan berikutnya yaitu mendeteksi apakah request memiliki parameter student id.

 - Ketika tidak ada parameter student id, maka endpoint ini mengembalikan semua data user yang ada. Fungsi `GetStudents()` dieksekusi.
 - Sedangkan jika ada parameter student id, maka hanya user dengan id yg diinginkan yg dijadikan nilai balik. Fungsi `SelectStudent(id)` dieksekusi.

Selanjutnya tambahkan satu fungsi lagi di `main()` yaitu `OutputJSON()`. Fungsi ini digunakan konversi data ke bentuk JSON string.

```go
func OutputJSON(w http.ResponseWriter, o interface{}) {
    res, err := json.Marshal(o)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}
```

Konversi dari objek atau slice ke JSON string dilakukan via `json.Marshal()`. Lebih jelasnya mengenai fungsi tersebut di bahas di chapter [A.53. JSON Data](/A-json.html).

## B.18.3. Data `Student`

Buka file `student.go`, siapkan struct `Student` dan variabel untuk menampung data yang bertipe `[]Student`. Data inilah yang nantinya dijadikan nilai balik endpoint `/student`.

```go
package main

var students = []*Student{}

type Student struct {
    Id    string
    Name  string
    Grade int32
}
```

Buat fungsi `GetStudents()`, fungsi ini mengembalikan semua data student. Buat juga fungsi `SelectStudent(id)`, fungsi ini mengembalikan data student sesuai dengan id terpilih.

```go
func GetStudents() []*Student {
    return students
}

func SelectStudent(id string) *Student {
    for _, each := range students {
        if each.Id == id {
            return each
        }
    }

    return nil
}
```

*Last but not least*, implementasikan fungsi `init()` yang didalamnya berisi pembuatan beberapa dummy data untuk ditampung variabel `students`. 

> Fungsi `init()` adalah fungsi yang secara otomatis dipanggil ketika package tersebut di import atau di run.

```go
func init() {
    students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
    students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
    students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
}
```

## B.18.4. Fungsi `Auth()` dan `AllowOnlyGET()`

Selanjutnya, ada dua fungsi lainnya yang perlu dipersiapkan yaitu `Auth()` dan `AllowOnlyGET()`.

#### ◉ Fungsi `Auth()`

Buka `middleware.go`, siapkan fungsi `Auth()`.

```go
package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
    username, password, ok := r.BasicAuth()
    if !ok {
        w.Write([]byte(`something went wrong`))
        return false
    }

    isValid := (username == USERNAME) && (password == PASSWORD)
    if !isValid {
        w.Write([]byte(`wrong username/password`))
        return false
    }

    return true
}
```

Tugas fungsi `Auth()` adalah memvalidasi apakah request merupakan valid basic auth request, dan juga apakah credentials yang dikirim cocok dengan data pada aplikasi kita. Informasi acuan credentials sendiri di hardcode pada konstanta `USERNAME` dan `PASSWORD`.

Fungsi `r.BasicAuth()` mengembalikan 3 informasi:

 1. Username
 2. Password
 3. Nilai balik ke-3 ini adalah representasi valid tidak nya basic auth request yang sedang berlangsung

Error dimunculkan ketika basic auth terdeteksi adalah tidak valid. Sedangkan jika ternyata valid, maka dilanjutkan ke proses otentikasi, mengecek apakah username dan password yang dikirim cocok dengan username dan password yang sudah di-*hardcode*.

#### ◉ Fungsi `AllowOnlyGET()`

Fungsi ini bertugas memastikan bahwa request yang diperbolehkan hanya yang ber-method `GET`. Selainnya, maka dianggap invalid request.

```go
func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
    if r.Method != "GET" {
        w.Write([]byte("Only GET is allowed"))
        return false
    }

    return true
}
```

## B.18.5. Testing

Semuanya sudah siap, sekarang jalankan aplikasi.

```bash
go run *.go
```

Jangan menggunakan `go run main.go`, dikarenakan dalam package `main` terdapat beberapa file lain yang harus diikutsertakan pada saat runtime.

![Run the server](images/B_http_basic_auth_2_run_server.png)

Test web service yang telah dibuat menggunakan command `curl`.

```bash
$ curl -X GET --user batman:secret http://localhost:9000/student
$ curl -X GET --user batman:secret http://localhost:9000/student?id=s001
```

![Consume API](images/B_http_basic_auth_3_test_api.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.18-http-basic-auth">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.18...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
