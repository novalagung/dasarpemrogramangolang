# Web API JSON

Pada bab ini kita akan mengkombinasikan pembahasan 2 bab sebelumnya, yaitu web dan JSON, untuk membuat sebuah web API dengan tipe data reponse berbentuk JSON.

> Web API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML. 

## Pembuatan Web API

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

Struct `student` di atas digunakan sebagai tipe elemen array sample data yang ditampung oleh variabel `data`.

Selanjutnya buat fungsi `users()` untuk handle rute `/users`. Didalam fungsi tersebut ada proses deteksi jenis request lewat property `r.Method()`, untuk mencari tahu apakah jenis request adalah **POST** atau **GET** atau lainnya.

```go
func users(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
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

Jika request adalah POST, maka data yang di-encode ke JSON dijadikan sebagai response.

Statement `w.Header().Set("Content-Type", "application/json")` digunakan untuk menentukan tipe response, yaitu sebagai JSON. Sedangkan `r.Write()` digunakan untuk mendaftarkan data sebagai response.

Selebihnya, jika request tidak valid, response di set sebagai error menggunakan fungsi `http.Error()`.

Siapkan juga handler untuk rute `/user`. Perbedaan rute ini dengan rute `/users` di atas adalah:

 - `/users` menghasilkan semua sample data yang ada (array).
 - `/user` menghasilkan satu buah data saja, diambel dari data sample berdasarkan `ID`-nya. Pada rute ini, client harus mengirimkan juga informasi `ID` data yang dicari.

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

        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}
```

Method `r.FormValue()` digunakan untuk mengambil data payload yang dikirim dari client, pada konteks ini data yang dimaksud adalah `ID`.

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

Jalankan program, maka web server sudah live dan bisa dikonsumsi datanya.

![Web API Server dijalankan](images/51_1_server.png)

## Test API

Setelah web server sudah berjalan, web API yang telah dibuat perlu untuk di tes. Di sini saya menggunakan Go\*gle Chrome plugin bernama [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=en) untuk mengetes API yang sudah dibuat.

 - Test `/users`, apakah data yang dikembalikan sudah benar.
   
    ![Test `/users`](images/51_2_test_api_users.png)

 - Test `/user` dengan data payload `id` adalah `E001`.

    ![Test `/user`](images/51_3_test_api_user.png)
