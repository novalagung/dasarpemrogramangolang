# C.24. Advanced Client HTTP Request

Pada bab ini kita akan belajar tentang topik yang sedikit berbeda dibanding bab sebelumnya, yaitu cara untuk melakukan http request ke sebuah web server.

Dua aplikasi akan dibuat, server dan client. Server merupakan aplikasi web server kecil, memiliki  satu endpoint. Lalu dari client http request di-trigger, dengan tujuan adalah server.

## C.24.1. Aplikasi Server

Buat projek baru seperti biasa, lalu buat `server.go`. Import package yang dibutuhkan.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type M map[string]interface{}
```

Pada fungsi `main()`, siapkan mux dengan isi satu buah handler, jalankan pada port `:9000`.

```go
func main() {
    mux := new(http.ServeMux)
    mux.HandleFunc("/data", ActionData)

    server := new(http.Server)
    server.Handler = mux
    server.Addr = ":9000"

    log.Println("Starting server at", server.Addr)
    err := server.ListenAndServe()
    if err != nil {
        log.Fatalln("Failed to start web server", err)
    }
}
```

Buat fungsi `ActionData()` yang merupakan handler dari rute `/data`. Handler ini hanya menerima method `POST`, dan mewajibkan consumer endpoint untuk menyisipkan payload dalam request-nya, dengan isi adalah JSON.

```go
func ActionData(w http.ResponseWriter, r *http.Request) {
    log.Println("Incoming request with method", r.Method)

    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusBadRequest)
        return
    }

    payload := make(M)
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if _, ok := payload["Name"]; !ok {
        http.Error(w, "Payload `Name` is required", http.StatusBadRequest)
        return
    }

    // ...
}
```

Isi dari `r.Body` kita decode ke objek `payload` yang bertipe `map[string]interface{}`. Setelah proses decoding selesai, terdapat pengecekan ada tidaknya property `Name` dalam payload. Jika tidak ada maka dianggap bad request.

Setelah itu, buat objek `data` dengan 2 property, yang salah satunya berisi kombinasi string dari payload `.Name`.

```go
data := M{
    "Message": fmt.Sprintf("Hello %s", payload["Name"]),
    "Status":  true,
}

w.Header().Set("Content-Type", "application/json")
err = json.NewEncoder(w).Encode(data)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
}
```

Cara render output JSON biasanya kita lakukan menggunakan statement `w.Write()` dengan isi adalah `[]byte` milik JSON. Ada cara lain, yaitu menggunakan json encoder. Penerapannya bisa dilihat pada kode di atas.

Aplikasi server sudah siap. Selanjutnya kita masuk ke bagian pembuatan aplikasi client.

## C.24.2. Aplikasi Client

Tugas dari aplikasi client: melakukan http request ke aplikasi server, pada endpoint `/data` sesuai dengan spesifikasi yang sudah dijelaskan di atas (ber-method POST, dan memiliki JSON payload).

Buat file baru, `client.go`, import package yang diperlukan.

```go
package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
)

type M map[string]interface{}
```

Buat fungsi `doRequest()`. Fungsi ini kita gunakan men-trigger http request.

```go
func doRequest(url, method string, data interface{}) (interface{}, error) {
    var payload *bytes.Buffer = nil

    if data != nil {
        payload = new(bytes.Buffer)
        err := json.NewEncoder(payload).Encode(data)
        if err != nil {
            return nil, err
        }
    }

    request, err := http.NewRequest(method, url, payload)
    if err != nil {
        return nil, err
    }

    // ...
}
```

Fungsi tersebut menerima 3 buah parameter.

 - Parameter `url`, adalah alamat tujuan request.
 - Parameter `method`, bisa GET, POST, PUT, ataupun method valid lainnya sesuai spesifikasi [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html).
 - Parameter `data`, isinya boleh kosong. Jika ada isinya, data tersebut di-encode ke bentuk JSON untuk selanjutnya disisipkan pada body request.

Buat objek request lewat `http.NewRequest()`. Sisipkan ke-3 parameter tersebut.

Selanjutnya buat objek client. Dari client ini request di-trigger, menghasilkan objek response.

```go
client := new(http.Client)

response, err := client.Do(request)
if response != nil {
    defer response.Body.Close()
}
if err != nil {
    return nil, err
}

responseBody := make(M)
err = json.NewDecoder(response.Body).Decode(&responseBody)
if err != nil {
    return nil, err
}

return responseBody, nil
```

Decode response tersebut ke tipe `M`, lalu tampilkan hasilnya.

Buat fungsi `main()`. Panggil fungsi `doRequest()` yang sudah dibuat. Untuk payload silakan isi sesuka hati, asalkan ada item dengan key `Name`. Lalu tampilkan response body hasil pemanggilan fungsi `doRequest()`.

```go
func main() {
    baseURL := "http://localhost:9000"
    method := "POST"
    data := M{"Name": "Noval Agung"}

    responseBody, err := doRequest(baseURL+"/data", method, data)
    if err != nil {
        log.Println("ERROR", err.Error())
        return
    }

    log.Printf("%#v \n", responseBody)
}
```

## C.24.3. Testing

Jalankan aplikasi server, buka prompt terminal/CMD baru, lalu jalankan aplikasi client.

![Testing](images/C_client_http_request_advanced_1_test_client_request.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.24-client-http-request">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.24...</a>
</div>
