# B.2. Routing `http.HandleFunc`

Dalam Go, routing bisa dilakukan dengan beberapa cara, diantaranya:

 1. Dengan memanfaatkan fungsi `http.HandleFunc()`
 2. Mengimplementasikan interface `http.Handler` pada suatu struct, untuk kemudian digunakan pada fungsi `http.Handle()`
 3. Membuat multiplexer sendiri dengan memanfaatkan struct `http.ServeMux`
 4. Dan lainnya

Pada buku ini, kesemua cara tersebut akan dibahas, namun khusus pada bab ini saja, hanya `http.HandleFunc()` yang kita pelajari.

> Metode routing cara pertama dan cara kedua memiliki kesamaan yaitu sama-sama menggunakan `DefaultServeMux` untuk pencocokan rute/endpoint yang diregistrasikan. Mengenai apa itu `DefaultServeMux` akan kita bahas lebih mendetail pada bab lain.

## B.2.1. Penggunaan `http.HandleFunc()`

Seperti yang sudah dijelaskan sekilas pada bab sebelumnya, fungsi `http.HandleFunc()` digunakan untuk registrasi rute/endpoint dan handler-nya.

<div id="ads">&nbsp;</div>

Penggunaan fungsi ini cukup mudah, panggil saja fungsi lalu isi dua parameternya.

 1. Parameter ke-1, adalah rute (atau endpoint). Sebagai contoh: `/`, `/index`, `/about`.
 2. Parameter ke-2, berisikan handler untuk rute bersangkutan. Sebagai contoh handler untuk rute `/` bertugas untuk menampilkan output berupa html `<p>hello</p>`.

Agar lebih mudah dipahami mari langsung praktek. Siapkan file `main.go` dengan package adalah `main`, dan import package `net/http` didalamnya.

```go
package main

import "fmt"
import "net/http"
```

Buat fungsi `main()`, didalamnya siapkan sebuah closure `handlerIndex`, lalu gunakan closure tersebut sebagai handler dari dua rute baru yang diregistrasi, yaitu `/` dan `/index`.

```go
func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
}
```

Selanjutnya, masih dalam fungsi `main()`, tambahkan rute baru `/data` dengan handler adalah anonymous function.

```go
func main() {
    // ...

    http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
    	w.Write([]byte("hello again"))
    })
}
```

Terakhir, jalankan server.

```go
func main() {
    // ...

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
```

## B.2.2. Run & Test

Tes dan lihat hasilnya.

![Rute `/data` mengembalikan data json](images/B.2_1_routing.png)

Dalam routing, handler bisa berupa fungsi, closure, ataupun anonymous function; bebas, yang terpenting adalah skema fungsi-nya sesuai dengan `func (http.ResponseWriter, *http.Request)`.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang/tree/master/chapter-B.2-routing-http-handlefunc">https://github.com/novalagung/dasarpemrogramangolang/.../chapter-B.2...</a>
</div>
