# B.22. Simple Configuration

Dalam development, pastinya kita programmer akan berurusan dengan banyak sekali variabel dan konstanta untuk keperluan konfigurasi. Misalnya, variabel berisi informasi port web server, timeout, variabel global, dan lainnya.

Pada chapter ini, kita akan belajar dasar pengelolahan variabel konfigurasi dengan memanfaatkan file JSON. 

## B.22.1. Struktur Aplikasi

Pertama-tama, buat project baru dengan struktur seperti gambar berikut.

![Structure](images/B_simple_configuration_1_structure.png)

Folder `conf` berisi 2 file.

 1. File `config.json`. Semua konfigurasi nantinya harus disimpan di file ini dalam struktur JSON.
 2. File `config.go`. Berisikan beberapa fungsi dan operasi untuk mempermudah pengaksesan konfigurasi dari file `config.json`.

## B.22.2. File Konfigurasi JSON `config.json`

Semua konfigurasi dituliskan dalam file ini. Desain struktur JSON-nya untuk bisa mudah dipahami, contoh:

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

Data JSON di atas berisi 4 buah data konfigurasi.

 1. Property `server.port`. Port yang digunakan saat start web server.
 2. Property `server.read_timeout`. Dijadikan sebagai timeout read.
 3. Property `server.write_timeout`. Dijadikan sebagai timeout write.
 4. Property `log.verbose`. Penentu apakah log di print atau tidak.

## B.22.3. Pemrosesan Konfigurasi

Pada file `config.go` kita akan siapkan sebuah fungsi yang isinya mengembalikan objek cetakan struct didapat dari konten file `config.json`.

Siapkan struct nya terlebih dahulu.

```go
package conf

import (
    "encoding/json"
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

Bisa dilihat pada kode di atas, struct bernama `_Configuration` dibuat. Struct ini berisikan banyak property yang strukturnya sama persis dengan isi file `config.json`. Dengan skema seperti itu akan cukup mempermudah developer dalam pengaksesan data konfigurasi.

Dari struct tersebut disiapkan objek bernama `shared`. Variabel ini berisi informasi konfigurasi hasil baca `config.json`, dan nantinya isinya bisa diakses via fungsi fungsi yang sebentar lagi akan dibuat.

Selanjutnya, siapkan fungsi `init()` dengan isi operasi baca file `config.json` serta operasi decode data JSON dari isi file tersebut ke variabel `shared`. 

Dengan adanya fungsi `init()` maka pada saat package `conf` ini di-import ke package lain otomatis file `config.json` dibaca dan di-parse untuk disimpan di variabel `shared`. Tambahkan juga validasi untuk memastikan kode hanya di-parse sekali saja.

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

    bts, err := os.ReadFile(filepath.Join(basePath, "conf", "config.json"))
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

Kemudian buat fungsi `Configuration()` yang isinya menjembatani pengaksesan object `shared`.

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

OK, kembali lagi ke contoh, dari mux di atas dibuatkan object baru bernama `router`, lalu beberapa rute didaftarkan ke object mux tersebut.

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

Selanjutnya, siapkan kode untuk start web server. Tulis kode berikut di dalam fungsi `main()` tepat setelah kode deklarasi route handler.

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

Di atas, ada objek baru dibuat dari struct `http.Server`, yaitu `server`. Untuk start server, panggil method `ListenAndServe()` milik objek tersebut.

Dengan memanfaatkan struct ini, kita bisa meng-custom beberapa konfigurasi default pada Go web server. Di antaranya seperti `ReadTimeout` dan `WriteTimeout`.

Bisa dilihat di contoh ada 4 buah properti milik `server` yang diisi nilainya dengan data konfigurasi.

 - `server.Handler`. Properti ini wajib di isi dengan custom mux yang dibuat.
 - `server.ReadTimeout`. Adalah timeout ketika memproses sebuah request. Kita isi dengan nilai dari configurasi.
 - `server.WriteTimeout`. Adalah timeout ketika memproses response.
 - `server.Addr`. Port yang digunakan web server pada saat start.

Ok. Sekarang jalankan aplikasi, akses dua buah endpoint yang sudah dibuat, kemudian cek di console.

![Structure](images/B_simple_configuration_2_log.png)

Coba ubah konfigurasi pada `config.json` nilai `log.verbose` menjadi `false`. Lalu restart aplikasi, maka log tidak muncul.

## B.22.5. Kekurangan Konfigurasi File

Ok, kita telah selesai belajar tentang cara membuat file konfigurasi yang terpusat dan mudah dibaca. Metode konfigurasi seperti ini umum digunakan, tapi dalam penerapannya memiliki beberapa *cons* yang mungkin akan mulai terasa ketika aplikasi arsitektur aplikasi berkembang dan arsitektur sistemnya menjadi kompleks. *Cons* yang dimaksud diantaranya adalah:

#### ◉ Tidak mendukung komentar

Komentar sangat penting karena untuk aplikasi besar yang konfigurasi item-nya sangat banyak, konfigurasi seperti pada contoh ini akan cukup susah untuk dikelola. Sebenarnya masalah ini bisa diselesaikan dengan mudah dengan cara mengadopsi file format lainnya, misalnya `YAML`, `.env`, atau lainnya.

#### ◉ Nilai konfigurasi harus diketahui di awal

Kita harus tau semua value tiap-tiap konfigurasi terlebih dahulu sebelum dituliskan ke file, dan sebelum aplikasi di-up. Dari sini akan sangat susah jika misal ada beberapa konfigurasi yang kita tidak tau nilainya tapi tau cara pengambilannya.

Contohnya seperti ini, di beberapa kasus, misalnya di AWS, database server yang di-setup secara automated akan meng-generate connection string yang host-nya bisa berganti-ganti tiap start-up, dan tidak hanya itu saja, bisa saja username, password dan lainnya juga tidak statis.

Dengan ini akan cukup merepotkan jika kita harus cari terlebih dahulu value konfigurasi tersebut untuk kemudian dituliskan ke file secara manual.

#### ◉ Tidak terpusat

Dalam pengembangan aplikasi, banyak konfigurasi yang nilai-nya akan didapat lewat jalan lain, seperti *environment variables* atau *command arguments*. Menyimpan konfigurasi file itu sudah cukup bagus, cuman untuk *case* dimana terdapat banyak sekali services, agak merepotkan pengelolahannya.

Ketika ada perubahan konfigurasi, semua services harus direstart.

#### ◉ Statis (tidak dinamis)

Konfigurasi umumnya dibaca hanya jika diperlukan. Penulisan konfigurasi dalam file membuat proses pembacaan file harus dilakukan di awal, haru kemudian kita bisa ambil nilai konfigurasi dari data yang sudah ada di memori.

Hal tersebut memiliki beberapa konsekuensi, untuk aplikasi yang di-manage secara automated, sangat mungkin adanya perubahan nilai konfigurasi. Dari sini berarti pembacaan konfigurasi file tidak boleh hanya dilakukan di awal saja. Tapi juga tidak boleh dilakukan di setiap waktu, karena membaca file itu ada *cost*-nya dari prespektif I/O.

#### ◉ Solusi

Kita akan membahas solusi dari beberapa masalah di atas (tidak semuanya) pada chapter terpisah, yaitu [C.11. Best Practice Configuration Menggunakan Environment Variable](/C-best-practice-configuration-env-var.html)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.22-simple-configuration">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.22...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
