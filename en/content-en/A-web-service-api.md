# A.54. Web Service API Server

Pada chapter ini kita akan mencoba mengkombinasikan hasil pembelajaran di 2 chapter sebelumnya (yaitu web programming dan JSON), untuk membuat sebuah web service API yang memiliki endpoint dengan response data mengadopsi format JSON.

> Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML atau format lainnya.

## A.54.1. Pembuatan Web API

Pertama siapkan terlebih dahulu struct dan beberapa data sample.

```go
package main

import "encoding/json"
import "fmt"
import "log"
import "net/http"

type student struct {
    ID    string
    Name  string
    Grade int
}

var data = []student{
    student{"E001", "ethan", 21},
    student{"W001", "wick", 22},
    student{"B001", "bourne", 23},
    student{"B002", "bond", 23},
}
```

Struct `student` di atas digunakan sebagai tipe elemen slice sample data, ditampung variabel `data`.

Selanjutnya buat fungsi `users()` untuk handle endpoint `/users`. Di dalam fungsi tersebut ada proses deteksi jenis request lewat property `r.Method`, untuk mencari tahu apakah jenis request adalah **GET** atau **POST** atau lainnya.

```go
func users(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "GET" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}
```

Jika request adalah GET (mengambil data), maka data yang di-encode ke JSON dijadikan sebagai response.

Statement `w.Header().Set("Content-Type", "application/json")` digunakan untuk menentukan tipe response, yaitu sebagai JSON. Sedangkan `w.Write()` digunakan untuk mendaftarkan data sebagai response.

Selebihnya, jika request tidak valid, response diset sebagai error menggunakan fungsi `http.Error()`.

Siapkan juga handler untuk endpoint `/user`. Perbedaan endpoint ini dengan `/users` di atas adalah:

 - Endpoint `/users` mengembalikan semua sample data yang ada (array).
 - Endpoint `/user` mengembalikan satu buah data saja, diambil dari data sample berdasarkan `ID`-nya. Pada endpoint ini, client harus mengirimkan juga informasi `ID` data yang dicari.

```go
func user(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var id = r.FormValue("id")
        var result []byte
        var err error

        for _, each := range data {
            if each.ID == id {
                result, err = json.Marshal(each)

                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                w.Write(result)
                return
            }
        }

        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}
```

Method `r.FormValue()` digunakan untuk mengambil data form yang dikirim dari client, pada konteks ini data yang dimaksud adalah `ID`.

Dengan menggunakan `ID` tersebut dicarilah data yang relevan. Jika ada, maka dikembalikan sebagai response. Jika tidak ada maka error **404, Not Found** dikembalikan dengan pesan **User not found**.

Terakhir, implementasikan kedua handler di atas.

```go
func main() {
    http.HandleFunc("/users", users)
    http.HandleFunc("/user", user)

    fmt.Println("starting web server at http://localhost:8080/")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

Jalankan program, sekarang web server sudah live dan bisa dikonsumsi datanya.

![Web API Server dijalankan](images/A_web_service_1_server.png)

## A.54.2. Test Web Service API via Postman

Setelah web server sudah berjalan, web service yang telah dibuat perlu untuk di-tes. Di sini saya menggunakan aplikasi [Postman](https://www.postman.com/) untuk mengetes API yang sudah dibuat.

 - Test endpoint `/users`, apakah data yang dikembalikan sudah benar.

    ![Test <code>/users</code>](images/A_web_service_2_test_api_users.png)

 - Test endpoint `/user`, isi form data `id` dengan nilai `E001`.

    ![Test <code>/user</code>](images/A_web_service_3_test_api_user.png)

## A.54.3. Test Web Service API via `cURL`

Testing bisa juga dilakukan via cURL. Pastikan untuk menginstall cURL terlebih dahulu agar bisa menggunakan command berikut.

```
curl -X GET http://localhost:8080/users
curl -X POST -d "id=B002" http://localhost:8080/user
```

![cURL test](images/A_web_service_4.png)

Data user ID pada endpoint `/user` dikirim sebagai form data menggunakan flag `-d`.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.54-web-service-api">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.54...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
