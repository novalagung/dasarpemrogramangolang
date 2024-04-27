# B.19. Middleware `http.Handler`

Pada chapter ini, topik yang dibahas adalah penerapan interface `http.Handler` untuk implementasi custom middleware. Kita gunakan sample proyek pada chapter sebelumnya [B.18. HTTP Basic Auth](/B-http-basic-auth.html) sebagai dasar bahan pembahasan chapter ini.

> Apa itu middleware? 
> 
> Istilah middleware berbeda-beda di tiap bahasa/framework. Di NodeJS dan Rails ada istilah middleware. Pada pemrograman Java Enterprise, istilah filters digunakan. Pada C# middleware disebut dengan delegate handlers. Definisi sederhana middleware adalah sebuah blok kode yang dipanggil sebelum ataupun sesudah http request di proses.

Pada chapter sebelumnya, terdapat beberapa proses yang dijalankan dalam handler rute `/student`, yaitu pengecekan otentikasi dan pengecekan HTTP method. Misalnya terdapat rute lagi, maka dua validasi tersebut juga harus dipanggil lagi dalam handlernya.

```go
func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if !Auth(w, r) {
        return
    }
    if !AllowOnlyGET(w, r) {
        return
    }

    // ...
}
```

Jika ada banyak rute, apa yang harus kita lakukan? salah satu solusi adalah dengan memanggil fungsi `Auth()` dan `AllowOnlyGet()` di setiap handler rute yang ada. Namun jelasnya ini bukan best practice karena mengharuskan penulisan kode yang berulang-ulang. Selain itu, bisa jadi ada jenis validasi lainnya yang harus diterapkan, misalnya misalnya pengecekan csrf, authorization, dan lainnya. Maka perlu ada desain penataan kode yang lebih efisien tanpa harus menuliskan validasi yang banyak tersebut berulang-ulang. 

Solusi yang pas adalah dengan membuat middleware baru untuk keperluan validasi.

## B.19.1. Interface `http.Handler`

Interface `http.Handler` merupakan tipe data paling populer di Go untuk keperluan manajemen middleware. Struct yang mengimplementasikan interface ini diwajibkan untuk memilik method dengan skema `ServeHTTP(ResponseWriter, *Request)`.

> Di Go, objek utama untuk keperluan routing web server adalah `mux` (kependekan dari multiplexer), dan `mux` ini mengimplementasikan interface `http.Handler`.

Kita akan buat beberapa middleware baru dengan memanfaatkan interface `http.Handler` untuk keperluan pengecekan otentikasi dan pengecekan HTTP method.

## B.19.2. Persiapan

OK, mari masuk ke bagian *coding*. Pertama duplikat folder project sebelumnya sebagai folder proyek baru. Lalu pada `main.go`, ubah isi fungsi `ActionStudent()` dan `main()`.

 - Fungsi`ActionStudent()`

    ```go
    func ActionStudent(w http.ResponseWriter, r *http.Request) {
        if id := r.URL.Query().Get("id"); id != "" {
            OutputJSON(w, SelectStudent(id))
            return
        }

        OutputJSON(w, GetStudents())
    }
    ```

 - Fungsi `main()`

    ```go
    func main() {
        mux := http.DefaultServeMux

        mux.HandleFunc("/student", ActionStudent)

        var handler http.Handler = mux
        handler = MiddlewareAuth(handler)
        handler = MiddlewareAllowOnlyGet(handler)

        server := new(http.Server)
        server.Addr = ":9000"
        server.Handler = handler

        fmt.Println("server started at localhost:9000")
        server.ListenAndServe()
    }
    ```


Perubahan pada kode `ActionStudent()` adalah penghapusan kode untuk pengecekan basic auth dan HTTP method. Selain itu, di fungsi `main()` juga terdapat cukup banyak perubahan, yang detailnya akan dijelaskan sebentar lagi.

## B.19.3. Mux / Multiplexer

Di Go, mux (kependekan dari multiplexer) adalah router. Semua routing pasti dilakukan lewat objek mux.

Apa benar? Routing `http.HandleFunc()` sepertinya tidak menggunakan mux? Begini, sebenarnya routing tersebut juga menggunakan mux. Go memiliki default objek mux yaitu `http.DefaultServeMux`. Routing yang langsung dilakukan dari fungsi `HandleFunc()` milik package `net/http` sebenarnya mengarah ke method default mux `http.DefaultServeMux.HandleFunc()`.

Agar lebih jelas perbedaannya, silakan perhatikan dua kode berikut.

```go
http.HandleFunc("/student", ActionStudent)

// vs

mux := http.DefaultServeMux
mux.HandleFunc("/student", ActionStudent)
```

Dua kode di atas melakukan prosees yang ekuivalen.

Mux sendiri adalah bentuk nyata struct yang mengimplementasikan interface `http.Handler`. Di kode setelah routing, bisa dilihat objek `mux` ditampung ke variabel baru bertipe `http.Handler`. Seperti ini adalah valid karena memang struct multiplexer memenuhi kriteria interface `http.Handler`, yaitu memiliki method `ServeHTTP()`. 

> Untuk lebih jelasnya silakan baca dokumentasi package net/http di [https://golang.org/pkg/net/http/#Handle](https://golang.org/pkg/net/http/#Handle)

Lalu dari objek `handler` tersebut, ke-dua middleware dipanggil dengan argument parameter diisi objek `handler` itu sendiri, dan nilai baliknya ditampung pada objek yang sama.

```go
var handler http.Handler = mux
handler = MiddlewareAuth(handler)
handler = MiddlewareAllowOnlyGet(handler)
```

Fungsi `MiddlewareAuth()` dan `MiddlewareAllowOnlyGet()` adalah middleware yang akan kita buat sebentar lagi. Cara registrasi middleware yang paling populer adalah dengan memanggilnya secara sekuensial atau berurutan, seperti pada kode di atas.

 - `MiddlewareAuth()` bertugas melakukan pengencekan credentials, basic auth.
 - `MiddlewareAllowOnlyGet()` bertugas melakukan pengecekan method.

> Silakan lihat source code beberapa library middleware yang sudah terkenal seperti gorilla, gin-contrib, echo middleware, dan lainnya; Semua metode implementasi middleware-nya adalah sama, atau minimal mirip. Point plus nya, beberapa di antara library tersebut mudah diintegrasikan dan *compatible* satu sama lain.

Kedua middleware yang akan kita buat tersebut mengembalikan fungsi bertipe `http.Handler`. Eksekusi middleware sendiri terjadi pada saat ada http request masuk.

Setelah semua middleware diregistrasi. Masukan objek `handler` ke property `.Handler` milik server.

```go
server := new(http.Server)
server.Addr = ":9000"
server.Handler = handler
```

## B.19.3. Pembuatan Middleware

Di dalam `middleware.go` ubah fungsi `Auth()` (hasil salinan project pada chapter sebelumnya) menjadi fungsi `MiddlewareAuth()`. Parameternya objek bertipe `http.Handler`, dan nilai baliknya juga sama.

```go
func MiddlewareAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()
        if !ok {
            w.Write([]byte(`something went wrong`))
            return
        }

        isValid := (username == USERNAME) && (password == PASSWORD)
        if !isValid {
            w.Write([]byte(`wrong username/password`))
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

Idealnya fungsi middleware harus mengembalikan struct yang implements `http.Handler`. Beruntungnya, Go sudah menyiapkan fungsi ajaib untuk mempersingkat pembuatan struct yang implement `http.Handler`, yaitu fungsi `http.HandlerFunc()`. Cukup bungkus callback `func(http.ResponseWriter,*http.Request)` sebagai tipe `http.HandlerFunc()` maka semuanya beres.

Isi dari `MiddlewareAuth()` sendiri adalah pengecekan basic auth (sama seperti pada chapter sebelumnya).

Tak lupa, ubah juga `AllowOnlyGet()` menjadi `MiddlewareAllowOnlyGet()`.

```go
func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            w.Write([]byte("Only GET is allowed"))
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

## B.19.4. Testing

Jalankan aplikasi.

![Run the server](images/B_http_basic_auth_2_run_server.png)

Lalu test menggunakan `curl`, hasilnya adalah sama dengan pada chapter sebelumnya.

![Consume API](images/B_http_basic_auth_3_test_api.png)

Dibanding metode pada chapter sebelumnya, dengan teknik ini kita lebih mudah mengontrol lalu lintas routing aplikasi, karena semua rute pasti melewati layer middleware terlebih dahulu sebelum sampai ke handler tujuan. Cukup maksimalkan saja penerapan middleware tanpa perlu menambahkan validasi di masing-masing handler.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.19-middleware-using-http-handler">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.19...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
