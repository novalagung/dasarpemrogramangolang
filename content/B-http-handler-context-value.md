# B.29. HTTP Handler Context Value

Dalam arsitektur web server yang menggunakan middleware, ada kebutuhan umum untuk meneruskan data dari satu lapisan ke lapisan berikutnya tanpa harus mengubah signature fungsi handler. Misalnya, middleware autentikasi yang sudah memvalidasi user perlu meneruskan informasi user tersebut ke handler, atau middleware logging yang men-generate request ID perlu memastikan ID tersebut tersedia di seluruh chain.

Go menyediakan mekanisme untuk penanganan kasus ini melalui context yang tertanam di `*http.Request`. Context bisa diisi dengan data di middleware lalu dibaca kembali di handler tanpa parameter tambahan. Pada chapter ini kita akan belajar cara menggunakannya.

## B.29.1. Konsep Context Value

Setiap `*http.Request` membawa sebuah context yang bisa diisi dengan data menggunakan `context.WithValue()`. Data tersebut kemudian bisa dibaca di handler lewat `r.Context().Value(key)`.

```go
// Di middleware: simpan data ke context
ctx := context.WithValue(r.Context(), key, value)
next.ServeHTTP(w, r.WithContext(ctx))

// Di handler: baca data dari context
value := r.Context().Value(key)
```

## B.29.2. Context Typed Key

Penting untuk **tidak menggunakan string biasa** sebagai key context. Jika ada 2 package berbeda yang sama-sama menggunakan key `"user"`, yang terjadi adalah value context dengan key `"user"` tersebut akan saling ditimpa.

Best practice-nya adalah menggunakan tipe custom, bisa dengan mendefinisikan tipe baru dari `string` atau `int`.

```go
type contextKey string

const (
    contextKeyUser      contextKey = "user"
    contextKeyRequestID contextKey = "request_id"
)
```

Tipe `contextKey` berbeda dari `string`, sehingga nilai key ini tidak akan pernah bentrok dengan key dari package lain.

## B.29.3. Implementasi

Pada contoh berikut kita akan membuat web server dengan dua middleware: satu untuk autentikasi yang menyimpan data user ke context, dan satu untuk request ID yang digunakan untuk logging dan tracing. Handler membaca kedua nilai tersebut dari context untuk menyusun response dan log.

Tulis ini di `main.go`.

```go
package main

import (
    "context"
    "log"
    "net/http"
)

type contextKey string

const (
    contextKeyUser      contextKey = "user"
    contextKeyRequestID contextKey = "request_id"
)

type User struct {
    Username string
    Role     string
}

func middlewareAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username, _, ok := r.BasicAuth()
        if !ok || username == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }

        user := User{Username: username, Role: "admin"}
        ctx := context.WithValue(r.Context(), contextKeyUser, user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func middlewareRequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := r.Header.Get("X-Request-ID")
        if requestID == "" {
            requestID = "auto-generated-001"
        }
        ctx := context.WithValue(r.Context(), contextKeyRequestID, requestID)
        w.Header().Set("X-Request-ID", requestID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
    user, ok := r.Context().Value(contextKeyUser).(User)
    if !ok {
        http.Error(w, "user not found in context", http.StatusInternalServerError)
        return
    }

    requestID, _ := r.Context().Value(contextKeyRequestID).(string)
    log.Printf("[%s] profile accessed by %s (%s)", requestID, user.Username, user.Role)

    w.Write([]byte("Hello, " + user.Username + " [" + user.Role + "]"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /profile", handleProfile)

    handler := middlewareRequestID(middlewareAuth(mux))

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", handler)
    if err != nil {
        log.Fatal(err)
    }
}
```

`middlewareAuth()` menggunakan `r.BasicAuth()` untuk membaca kredensial dari header `Authorization`. Jika valid, informasi user dibungkus ke dalam struct `User` lalu disimpan ke context via `context.WithValue()`. Alasan menggunakan struct bukan sekadar string adalah agar data yang dibawa lebih terstruktur dan bisa langsung digunakan tanpa parsing tambahan di handler.

`middlewareRequestID()` membaca header `X-Request-ID` dari request. Jika header tersebut tidak dikirim client, nilai fallback `"auto-generated-001"` digunakan sebagai gantinya. Request ID kemudian disimpan ke context sekaligus dikirim kembali ke client via response header `X-Request-ID`, sehingga client bisa menggunakannya untuk keperluan tracing.

Kedua middleware menggunakan pola `r.WithContext(ctx)` untuk meneruskan context yang sudah diperkaya ke handler berikutnya. `r.WithContext()` tidak memodifikasi request asli, melainkan mengembalikan salinan request baru dengan context yang diberikan.

Di `main()`, middleware dirantai dengan pola `middlewareRequestID(middlewareAuth(mux))`. Urutan nesting ini menentukan urutan eksekusi: `middlewareRequestID()` dieksekusi pertama, lalu `middlewareAuth()`, baru handler. Urutan ini penting karena jika dibalik, request ID belum ada di context saat auth berjalan.

## B.29.4. Penjelasan Alur

Urutan eksekusi middleware dan handler pada setiap request adalah sebagai berikut.

 1. `middlewareRequestID()`: membaca atau men-generate request ID, menyimpannya ke context
 2. `middlewareAuth()`: memvalidasi basic auth, membuat objek `User`, menyimpannya ke context
 3. `handleProfile()`: membaca `User` dan request ID dari context, menggunakannya untuk response

Di `handleProfile()`, nilai diambil dari context menggunakan *type assertion*:

```go
user, ok := r.Context().Value(contextKeyUser).(User)
```

Nilai yang disimpan di context bertipe `any`, sehingga perlu di-*assert* ke tipe yang tepat. Selalu gunakan bentuk dua nilai (`value, ok`) agar aman jika key tidak ditemukan atau nilainya bukan tipe yang diharapkan.

> Lebih jelasnya mengenai context dibahas di chapter [A.65. Concurrency Pattern: Context Cancellation Pipeline](/A-pipeline-context-cancellation.html)

## B.29.5. Testing

Jalankan server lalu coba tiga skenario berikut.

```bash
# Request tanpa auth header
curl http://localhost:9000/profile
```

Response: `unauthorized` dengan status 401.

```bash
# Request dengan basic auth
curl --user batman:secret http://localhost:9000/profile
```

Response: `Hello, batman [admin]` dengan status 200. Server juga mencetak log seperti `[auto-generated-001] profile accessed by batman (admin)`.

```bash
# Request dengan basic auth dan custom request ID
curl --user batman:secret \
  -H "X-Request-ID: req-abc-123" \
  http://localhost:9000/profile
```

Response tetap sama, namun log di server mencetak `[req-abc-123] profile accessed by batman (admin)`. Response header `X-Request-ID: req-abc-123` juga dikembalikan ke client.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.29-http-handler-context-value">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.29...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
