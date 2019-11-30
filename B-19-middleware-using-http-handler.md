# B.19. Middleware `http.Handler`

Pada bab ini, kita akan belajar penggunaan interface `http.Handler` untuk implementasi custom middleware. Kita akan menggunakan sample proyek pada bab sebelumnya [Bab B.18. HTTP Basic Authentication](/B-18-http-basic-auth.html) sebagai dasar bahan pembahasan bab ini.

> Apa itu middleware? Istilah middleware berbeda-beda di tiap bahasa/framework. NodeJS dan Rails ada istilah middleware. Pada pemrograman Java Enterprise, istilah filters digunakan. Pada C# istilahnya adalah delegate handlers. Definisi dari middleware sendiri versi penulis, sebuah blok kode yang dipanggil sebelum ataupun sesudah http request di proses.

Di bab sebelumnya, kalau dilihat, ada beberapa proses yang dijalankan dalam handler rute `/student`, yaitu pengecekan otentikasi dan pengecekan method. Misalnya terdapat rute lagi, maka dua validasi tersebut juga harus dipanggil lagi dalam handlernya.

```go
func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if !Auth(w, r) { return }
    if !AllowOnlyGET(w, r) { return }

    // ...
}
```

Jika ada banyak rute, apa yang harus kita lakukan? salah satu solusi yang bisa digunakan adalah dengan memanggil fungsi `Auth()` dan `AllowOnlyGet()` di semua handler rute yang ada. Namun jelasnya ini bukan best practice. Dan juga belum tentu di tiap rute hanya ada dua validasi ini, bisa saja ada lebih banyak proses, misalnya pengecekan csrf, authorization, dan lainnya. 

Solusi dari masalah tersebut adalah, mengkonversi fungsi-fungsi di atas menjadi middleware.

## B.19.1. Interface `http.Handler`

Interface `http.Handler` merupakan tipe data paling populer di Go untuk keperluan manajemen middleware. Struct yang mengimplementasikan interface ini diwajibkan memilik method dengan skema `ServeHTTP(ResponseWriter, *Request)`.

Di Go sendiri objek utama untuk keperluan routing yaitu `mux` atau multiplexer, adalah mengimplementasikan interface `http.Handler` ini.

Dengan memanfaatkan interface ini, kita akan membuat beberapa middleware. Fungsi  pengecekan otentikasi dan pengecekan method akan kita ubah menjadi middleware terpisah.

## B.19.2. Persiapan

OK, mari kita praktekan. Pertama duplikat folder projek sebelumnya sebagai folder proyek baru.  Lalu pada `main.go`, ubah isi fungsi `ActionStudent` dan `main`.

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


Perubahan pada kode `ActionStudent()` adalah, pengecekan basic auth dan pengecekan method dihapus. Selain itu di fungsi `main()` juga terdapat cukup banyak perubahan, yang detailnya akan kita bahas di bawah ini.

## B.19.3. Mux / Multiplexer

Di Go, mux (kependekan dari multiplexer) adalah router. Semua routing pasti dilakukan lewat objek mux.

Apa benar? routing `http.HandleFunc()` sepertinya tidak menggunakan mux? Begini, sebenarnya routing tersebut juga menggunakan mux. Go memiliki default objek mux yaitu `http.DefaultServeMux`. Routing yang langsung dilakukan dari fungsi `HandleFunc()` milik package `net/http` sebenarnya mengarah ke method default mux `http.DefaultServeMux.HandleFunc()`. Agar lebih jelas, silakan perhatikan dua kode berikut.

```go
http.HandleFunc("/student", ActionStudent)

// vs

mux := http.DefaultServeMux
mux.HandleFunc("/student", ActionStudent)
```

Dua kode di atas menghasilkan hasil yang sama persis.

Mux sendiri adalah bentuk nyata struct yang mengimplementasikan interface `http.Handler`. Untuk lebih jelasnya silakan baca dokumentasi package net/http di https://golang.org/pkg/net/http/#Handle.

Kembali ke pembahasan source code. Di kode setelah routing, bisa dilihat objek `mux` ditampung ke variabel baru bertipe `http.Handler`. Seperti ini adalah valid karena memang struct multiplexer memenuhi kriteria interface `http.Handler`, yaitu memiliki method `ServeHTTP()`. 

Lalu dari objek `handler` tersebut, ke-dua middleware dipanggil dengan parameter adalah objek `handler` itu sendiri dan nilai baliknya ditampung pada objek yang sama.

```go
var handler http.Handler = mux
handler = MiddlewareAuth(handler)
handler = MiddlewareAllowOnlyGet(handler)
```

Fungsi `MiddlewareAuth()` dan `MiddlewareAllowOnlyGet()` adalah middleware yang akan kita buat setelah ini. Cara registrasi middleware yang paling populer adalah dengan memanggilnya secara sekuensial atau berurutan, seperti pada kode di atas.

 - `MiddlewareAuth()` bertugas untuk melakukan pengencekan credentials, basic auth.
 - `MiddlewareAllowOnlyGet()` bertugas untuk melakukan pengecekan method.

> Silakan lihat source code beberapa library middleware yang sudah terkenal seperti gorilla, gin-contrib, echo middleware, dan lainnya; kesemuanya metode implementasi middleware-nya adalah sama, atau paling tidak mirip. Point plus nya, beberapa diantara library tersebut mudah diintegrasikan dan compatible satu sama lain.

Kedua middleware yang akan kita buat tersebut mengembalikan fungsi bertipe `http.Handler`. Eksekusi middleware sendiri terjadi pada saat ada http request masuk.

Setelah semua middleware diregistrasi. Masukan objek `handler` ke property `.Handler` milik server.

```go
server := new(http.Server)
server.Addr = ":9000"
server.Handler = handler
```

## B.19.3. Pembuatan Middleware

Di dalam `middleware.go` ubah fungsi `Auth()` (hasil salinan projek pada bab sebelumnya) menjadi fungsi `MiddlewareAuth()`. Parameternya objek bertipe `http.Handler`, dan nilai baliknya juga sama.

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

Idealnya fungsi middleware harus mengembalikan struct yang implements `http.Handler`. Beruntungnya, Go sudah menyiapkan fungsi ajaib untuk mempersingkat pembuatan struct-yang-implemenets-`http.Handler`. Fungsi tersebut adalah `http.HandlerFunc`, cukup bungkus callback `func(http.ResponseWriter,*http.Request)` sebagai tipe `http.HandlerFunc` dan semuanya beres.

Isi dari `MiddlewareAuth()` sendiri adalah pengecekan basic auth, sama seperti pada bab sebelumnya.

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

![Run the server](images/B.18_2_run_server.png)

Lalu test menggunakan `curl`, hasilnya adalah sama dengan pada bab sebelumnya.

![Consume API](images/B.18_3_test_api.png)

Dibanding metode pada bab sebelumnya, dengan teknik ini kita bisa sangat mudah mengontrol lalu lintas routing aplikasi, karena semua rute pasti melewati middleware terlebih dahulu sebelum sampai ke tujuan. Cukup maksimalkan middleware tersebut tanpa menggangu fungsi callback masing-masing rute.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.19-middleware-using-http-handler">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.19...</a>
</div>
