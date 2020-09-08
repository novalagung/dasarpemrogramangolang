# B.22. Simple Configuration

Dalam development, pasti banyak sekali variabel dan konstanta yang diperlukan. Mulai dari variabel yang dibutuhkan untuk start server seperti port, timeout, hingga variabel global dan variabel shared lainnya.

Pada bab ini, kita akan belajar cara membuat config file modular. 

## B.22.1. Struktur Aplikasi

Pertama-tama, buat project baru, siapkan dengan struktur seperti gambar berikut.

![Structure](images/B.22_1_structure.png)

Folder `conf` berisi 2 file.

 1. File `config.json`. Semua konfigurasi nantinya harus disimpan di file ini dalam struktur JSON.
 2. File `config.go`. Berisikan beberapa fungsi dan operasi untuk mempermudah pengaksesan konfigurasi dari file `config.json`.

## B.22.2. File Konfigurasi JSON `config.json`

Semua konfigurasi perlu dituliskan dalam file ini. Desain struktur JSON nya untuk bisa mudah dipahami. Tulis data berikut di file tersebut.

```json
{
    "server": {
        "port": 9000,
        "read_timeout": 5,
        "write_timeout": 5
    },

    "log": {
        "verbose": true
    }
}
```

Ada 4 buah konfigurasi disiapkan.

 1. Property `server.port`. Port yang digunakan saat start web server.
 2. Property `server.read_timeout`. Dijadikan sebagai timeout read.
 3. Property `server.write_timeout`. Dijadikan sebagai timeout write.
 4. Property `log.verbose`. Penentu apakah log di print atau tidak.

## B.22.3. Pemrosesan Konfigurasi

Pada file `config.go`, nantinya kita akan buat sebuah fungsi, isinya mengembalikan objek cetakan struct representasi dari `config.json`.

Siapkan struct nya terlebih dahulu.

```go
package conf

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

var shared *_Configuration

type _Configuration struct {
    Server struct {
        Port         int           `json:"port"`
        ReadTimeout  time.Duration `json:"read_timeout"`
        WriteTimeout time.Duration `json:"write_timeout"`
    } `json:"server"`

    Log struct {
        Verbose bool `json:"verbose"`
    } `json:"log"`
}
```

Bisa dilihat pada kode di atas, struct bernama `_Configuration` dibuat. Struct ini berisikan banyak property yang strukturnya sama persis dengan isi file `config.json`. Dengan desain seperti ini, akan sangat memudahkan developer dalam pengaksesan konfigurasi.

Dari struct tersebut tercetak private objek bernama `shared`. Variabel inilah yang nantinya akan dikembalikan lewat fungsi yang akan kita buat.

Selanjutnya, isi `init()` dengan beberapa proses: membaca file json, lalu di decode ke object `shared`. 

Dengan menuliskan proses barusan ke fungsi `init()`, pada saat package `conf` ini di import ke package lain maka file `config.json` akan otomatis di parsing. Dan dengan menambahkan sedikit validasi, parsing hanya akan terjadi sekali di awal.


```go
func init() {
    if shared != nil {
        return
    }

    basePath, err := os.Getwd()
    if err != nil {
        panic(err)
        return
    }

    bts, err := ioutil.ReadFile(filepath.Join(basePath, "conf", "config.json"))
    if err != nil {
        panic(err)
        return
    }

    shared = new(_Configuration)
    err = json.Unmarshal(bts, &shared)
    if err != nil {
        panic(err)
        return
    }
}
```

Lalu buat fungsi yang mengembalikan object `shared`.

```go
func Configuration() _Configuration {
    return *shared
}
```

## B.22.4. Routing & Server

Masuk ke bagian implementasi, buka `main.go`, lalu buat custom mux.

```go
package main

import (
    "chapter-B.22/conf"
    "fmt"
    "log"
    "net/http"
    "time"
)

type CustomMux struct {
    http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if conf.Configuration().Log.Verbose {
        log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
    }

    c.ServeMux.ServeHTTP(w, r)
}
```

Bisa dilihat dalam method `ServeHTTP()` di atas, ada pengecekan salah satu konfigurasi, yaitu `Log.Verbose`. Cara pengaksesannya cukup mudah, yaitu lewat fungsi `Configuration()` milik package `conf` yang telah di-import.

OK, kembali lagi ke contoh, dari mux di atas, buat object baru bernama `router`, lalu lakukan registrasi beberapa rute.

```go
func main() {
    router := new(CustomMux)
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("How are you?"))
    })

    // ...
}
```

Selanjutnya, kita akan start web server untuk serve mux di atas. Masih di dalam `main.go`, tambahkan kode berikut.


```go
server := new(http.Server)
server.Handler = router
server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
server.WriteTimeout = conf.Configuration().Server.WriteTimeout * time.Second
server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

if conf.Configuration().Log.Verbose {
    log.Printf("Starting server at %s \n", server.Addr)
}

err := server.ListenAndServe()
if err != nil {
    panic(err)
}
```

Objek baru bernama `server` telah dibuat dari struct `http.Server`. Untuk start server cukup panggil method `ListenAndServe()` milik objek tersebut.

Dengan memanfaatkan struct ini, kita bisa meng-custom beberapa konfigurasi default pada Go web server. Di antaranya seperti `ReadTimeout` dan `WriteTimeout`.

Pada kode di atas bisa kita lihat, ada 4 buah properti milik `server` di-isi.

 - `server.Handler`. Properti ini wajib di isi dengan custom mux yang dibuat.
 - `server.ReadTimeout`. Adalah timeout ketika memproses sebuah request. Kita isi dengan nilai dari configurasi.
 - `server.WriteTimeout`. Adalah timeout ketika memproses response.
 - `server.Addr`. Port yang digunakan web server pada saat start.

Terakhir jalankan aplikasi, akses dua buah endpoint yang sudah dibuat, lalu coba cek di console.

![Structure](images/B.22_2_log.png)

Coba ubah konfigurasi pada `config.json` nilai `log.verbose` menjadi `false`. Lalu restart aplikasi, maka log tidak muncul.

## B.22.5. Kekurangan Konfigurasi File

Ok, kita telah selesai belajar tentang cara membuat file konfigurasi yang mudah dibaca dan praktis. Namun penerapan kontrol konfigurasi dengan metode ini kurang dianjurkan karena beberapa hal:

#### • Tidak mendukung komentar

Komentar sangat penting karena untuk aplikasi besar yang konfigurasi item-nya sangat banyak - akan susah untuk dipahami. Sebenarnya perihal ini bisa di-*resolve* menggunakan jenis konfigurasi lain seperti `YAML`, `.env`, atau lainnya.

#### • Nilai konfigurasi harus diketahui di awal

Kita harus tau semua value tiap-tiap konfigurasi terlebih dahulu, dan dituliskan ke file, sebelum aplikasi di-up. Dari sini akan sangat susah jika misal ada beberapa konfigurasi yang kita tidak tau nilainya tapi tau cara pengambilannya.

Contohnya pada beberapa kasus, seperti di AWS, database server yang di-setup secara automated akan meng-generate connection string yang host-nya bisa berganti-ganti tiap start-up, dan tidak hanya itu, bisa saja username, password dan lainnya juga tidak statis.

Dengan ini akan sangat susah jika kita harus cari terlebih dahulu value konfigurasi tersebut untuk kemudian dituliskan ke file. Memakan waktu dan kurang baik dari banyak sisi.

#### • Tidak terpusat

Dalam pengembangan aplikasi, banyak konfigurasi yang nilai-nya akan didapat lewat jalan lain, seperti *environment variables* atau *command arguments*.

Akan lebih mudah jika hanya ada satu sumber konfigurasi saja untuk dijadikan acuan.

#### • Statis (tidak dinamis)

Konfigurasi umumnya dibaca hanya jika diperlukan. Penulisan konfigurasi dalam file membuat proses pembacaan file harus dilakukan di awal, haru kemudian kita bisa ambil nilai konfigurasi dari data yang sudah ada di memori.

Hal tersebut memiliki beberapa konsekuensi, untuk aplikasi yang di-manage secara automated, sangat mungkin adanya perubahan nilai konfigurasi. Dari sini berarti pembacaan konfigurasi file tidak boleh hanya dilakukan di awal saja. Tapi juga tidak boleh dilakukan di setiap waktu, karena membaca file itu ada *cost*-nya dari prespektif I/O.

#### • Solusi

Kita akan membahas solusi dari beberapa masalah di atas pada chapter terpisah, yaitu ketika masuk ke C.8.

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.22-simple-configuration">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.22...</a>
</div>
