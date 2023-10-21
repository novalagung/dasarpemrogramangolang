# A.54. Web Service API Server

Pada chapter ini kita akan mengkombinasikan pembahasan 2 chapter sebelumnya, yaitu web programming dan JSON, untuk membuat sebuah web service API dengan tipe data reponse berbentuk JSON.

> Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML.

## A.54.1. Pembuatan Web API

Pertama siapkan terlebih dahulu struct dan beberapa data sample.

```go
package main

import "encoding/json"
import "net/http"
import "fmt"

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

Selanjutnya buat fungsi `users()` untuk handle endpoint `/users`. Di dalam fungsi tersebut ada proses deteksi jenis request lewat property `r.Method()`, untuk mencari tahu apakah jenis request adalah **POST** atau **GET** atau lainnya.

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

Statement `w.Header().Set("Content-Type", "application/json")` digunakan untuk menentukan tipe response, yaitu sebagai JSON. Sedangkan `r.Write()` digunakan untuk mendaftarkan data sebagai response.

Selebihnya, jika request tidak valid, response di set sebagai error menggunakan fungsi `http.Error()`.

Siapkan juga handler untuk endpoint `/user`. Perbedaan endpoint ini dengan `/users` di atas adalah:

 - Endpoint `/users` mengembalikan semua sample data yang ada (array).
 - Endpoint `/user` mengembalikan satu buah data saja, diambel dari data sample berdasarkan `ID`-nya. Pada endpoint ini, client harus mengirimkan juga informasi `ID` data yang dicari.

```go
func user(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "GET" {
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

Dengan menggunakan `ID` tersebut dicarilah data yang relevan. Jika ada, maka dikembalikan sebagai response. Jika tidak ada maka error **400, Bad Request** dikembalikan dengan pesan **User Not Found**.

Terakhir, implementasikan kedua handler di atas.

```go
func main() {
    http.HandleFunc("/users", users)
    http.HandleFunc("/user", user)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
```

Jalankan program, sekarang web server sudah live dan bisa dikonsumsi datanya.

![Web API Server dijalankan](images/A_web_service_1_server.png)

## A.54.2. Test Web Service API via Postman

Setelah web server sudah berjalan, web service yang telah dibuat perlu untuk di-tes. Di sini saya menggunakan Google Chrome plugin bernama [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=en) untuk mengetes API yang sudah dibuat.

 - Test endpoint `/users`, apakah data yang dikembalikan sudah benar.

    ![Test <code>/users</code>](images/A_web_service_2_test_api_users.png)

 - Test endpoint `/user`, isi form data `id` dengan nilai `E001`.

    ![Test <code>/user</code>](images/A_web_service_3_test_api_user.png)

## A.54.3. Test Web Service API via `cURL`

Untuk testing bisa juga memanfaatkan cURL. Apabila pembaca menggunakan Windows 10, seharusnya sudah ter-include cURL. Jika bukan pengguna Windows 10, bisa menginstall-nya dan mendaftarkannya ke path variables (agar bisa diakses melalui terminal/cmd dari mana saja).

curl -X GET http://localhost:8080/users curl -X GET http://localhost:8080/user?id=B002 ```

![cURL test](images/A_web_service_4.png)

Data ID yang ingin dicari melalui endpoint /user, ditulis dengan ?id=B002 yang berarti dilewatkan melalui query parameters (umumnya data yang ingin dilampirkan melalui method GET adalah dengan query parameters).

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.54-web-service-api">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.54...</a>
</div>

---

<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>
