# B.23. Handler Testing dengan `net/http/httptest`

Pada chapter-chapter sebelumnya kita telah membuat banyak handler, namun pengujiannya selalu dilakukan secara manual lewat browser atau `curl`. Go menyediakan package `net/http/httptest` yang memungkinkan penulisan _automated test_ untuk handler tanpa perlu menjalankan server sungguhan.

Pendekatan testing ini sering disebut dengan _mock testing_ pada layer HTTP, di mana kita menggantikan komponen infrastruktur (server dan koneksi TCP) dengan objek palsu (_mocked object_) yang berperilaku sama. Handler tidak perlu tahu apakah ia dipanggil dari server sungguhan atau dari test, karena ia hanya berinteraksi lewat interface `http.ResponseWriter` dan `*http.Request`. Package `httptest` menyediakan implementasi palsu dari kedua komponen tersebut sehingga handler bisa diuji secara terisolasi, cepat, dan tanpa ketergantungan jaringan.

## B.23.1. Pengenalan Package `httptest`

Package `net/http/httptest` menyediakan fungsi utama berikut ini:

- `httptest.NewRequest()`: membuat objek `*http.Request` untuk keperluan test.
- `httptest.NewRecorder()`: membuat objek `ResponseWriter` palsu yang merekam response. Cocok untuk test handler secara langsung tanpa server.
- `httptest.NewServer()`: membuat HTTP server sungguhan di port acak untuk keperluan test yang membutuhkan HTTP client nyata.

## B.23.2. Praktik

Mari coba praktikkan fungsi-fungsi tersebut untuk testing. Namun sebelum itu, siapkan dulu aplikasinya. Buat file `main.go` berisi dua buah handler sederhana berikut yang nantinya akan kita test.

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func HandlerHello(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "name is required", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello, " + name})
}

func HandlerPing(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("pong"))
}

func main() {
    http.HandleFunc("/hello", HandlerHello)
    http.HandleFunc("/ping", HandlerPing)

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

## B.23.3. Penggunaan `httptest.NewRequest()` dan `httptest.NewRecorder()`

Sekarang buat file testing `main_test.go`. Di dalamnya disiapkan kode unit test untuk `HandlerPing` dan `HandlerHello`.

```go
package main

import (
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandlerPing(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/ping", nil)
    w := httptest.NewRecorder()

    HandlerPing(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status 200, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if string(body) != "pong" {
        t.Errorf("expected body 'pong', got '%s'", string(body))
    }
}
```

`httptest.NewRequest()` membuat `*http.Request` tanpa perlu koneksi jaringan. `httptest.NewRecorder()` membuat `ResponseWriter` yang merekam semua data response ke dalam buffer internal.

Setelah handler dipanggil, `w.Result()` digunakan untuk mendapatkan objek `*http.Response` dari buffer tersebut. Dari sana status code, header, dan body bisa diinspeksi.

Jalankan command berikut untuk eksekusi file testing.

```bash
go test -v ./...
```

Hasilnya, testing **PASS**.

![Test passed](images/B_httptest_handler_testing_1.png)

## B.23.4. Handler Testing Menggunakan Table-Driven Test

Selanjutnya buat fungsi test lagi untuk testing handler `/hello`. Kali ini kita akan pakai _table-driven test_ dengan 2 skenario testing.

- Skenario nama valid
- Skenario nama kosong

Tulis kode berikut di file test yang telah dibuat.

```go
func TestHandlerHello(t *testing.T) {
    tests := []struct {
        name           string
        query          string
        expectedStatus int
        expectedMsg    string
    }{
        {"valid name", "?name=Batman", http.StatusOK, "Batman"},
        {"missing name", "", http.StatusBadRequest, "name is required"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodGet, "/hello"+tt.query, nil)
            w := httptest.NewRecorder()

            HandlerHello(w, req)

            res := w.Result()
            if res.StatusCode != tt.expectedStatus {
                t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
            }

            body, _ := io.ReadAll(res.Body)
            if tt.expectedStatus == http.StatusOK {
                var result map[string]string
                json.Unmarshal(body, &result)
                if result["message"] != "Hello, "+tt.expectedMsg {
                    t.Errorf("unexpected message: %s", result["message"])
                }
            }
        })
    }
}
```

Jalankan test menggunakan command ini:

```bash
go test -v ./...
```

Hasilnya, testing **PASS**.

![Test passed](images/B_httptest_handler_testing_2.png)

## B.23.5. Penggunaan `httptest.NewServer()`

Gunakan `httptest.NewServer()` untuk skenario testing yang membutuhkan HTTP client sungguhan (misalnya mengikuti redirect atau membaca cookie).

```go
func TestHandlerHelloWithServer(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", HandlerHello)

    ts := httptest.NewServer(mux)
    defer ts.Close()

    res, err := http.Get(ts.URL + "/hello?name=Batman")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected 200, got %d", res.StatusCode)
    }
}
```

Jalankan command berikut untuk eksekusi file testing.

```bash
go test -v ./...
```

Hasilnya, testing **PASS**.

![Test passed](images/B_httptest_handler_testing_3.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div><a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.23-httptest-handler-testing">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.23...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
