# A.2. Instalasi Golang (Stable & Unstable)

Hal pertama yang perlu dilakukan sebelum bisa menggunakan Go adalah meng-*install*-nya terlebih dahulu. Panduan instalasi sebenarnya sudah disediakan di situs resmi Go [http://golang.org/doc/install#install](http://golang.org/doc/install#install).

Di sini penulis mencoba meringkas petunjuk instalasi pada *link* di atas, agar lebih mudah untuk diikuti terutama untuk pembaca yang baru belajar.

> Go yang digunakan adalah versi **1.20**, direkomendasikan menggunakan versi tersebut.

URL untuk mengunduh *installer* Go: https://golang.org/dl/. Silakan langsung unduh dari *link* tersebut lalu lakukan proses instalasi, atau bisa mengikuti petunjuk pada chapter ini.

## A.2.1. Instalasi Go *Stable*

#### • Instalasi Go di Windows

 1. Download terlebih dahulu *installer*-nya di [https://golang.org/dl/](https://golang.org/dl/). Pilih *installer* untuk sistem operasi Windows sesuai jenis bit yang digunakan.

 2. Setelah ter-*download*, jalankan *installer*, klik *next* hingga proses instalasi selesai. *By default* jika anda tidak merubah path pada saat instalasi, Go akan ter-*install* di `C:\go`. *Path* tersebut secara otomatis akan didaftarkan dalam `PATH` *environment variable*.

 3. Buka *Command Prompt* / *CMD*, eksekusi perintah berikut untuk mengecek versi Go.

    ```bash
    go version
    ```

 4. Jika output adalah sama dengan versi Go yang ter-*install*, menandakan proses instalasi berhasil.

> Sering terjadi, command `go version` tidak bisa dijalankan meskipun instalasi sukses. Solusinya bisa dengan restart CMD (tutup CMD, kemudian buka lagi). Setelah itu coba jalankan ulang command di atas.

#### • Instalasi Go di MacOS

Cara termudah instalasi Go di MacOS adalah menggunakan [Homebrew](http://brew.sh/).

 1. *Install* terlebih dahulu Homebrew (jika belum ada), caranya jalankan perintah berikut di **terminal**.

    ```bash
    $ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```

 2. *Install* Go menggunakan command `brew`.

    ```bash
    $ brew install go
    ```

 3. Tambahkan path binary Go ke `PATH` *environment variable*.

    ```bash
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
    $ source ~/.bash_profile
    ```

 4. Jalankan perintah berikut mengecek versi Go.

    ```bash
    go version
    ```

 5. Jika output adalah sama dengan versi Go yang ter-*install*, menandakan proses instalasi berhasil.

#### • Instalasi Go di Linux

 1. Unduh arsip *installer* dari [https://golang.org/dl/](https://golang.org/dl/), pilih installer untuk Linux yang sesuai dengan jenis bit komputer anda. Proses download bisa dilakukan lewat CLI, menggunakan `wget` atau `curl`.

    ```bash
    $ wget https://storage.googleapis.com/golang/go1...
    ```

 2. Buka terminal, *extract* arsip tersebut ke `/usr/local`.

    ```bash
    $ tar -C /usr/local -xzf go1...
    ```

 3. Tambahkan path binary Go ke `PATH` *environment variable*.

    ```bash
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
    $ source ~/.bashrc
    ```

 4. Selanjutnya, eksekusi perintah berikut untuk mengetes apakah Go sudah terinstal dengan benar.

    ```bash
    go version
    ```

 5. Jika output adalah sama dengan versi Go yang ter-*install*, menandakan proses instalasi berhasil.

## A.2.2. Variabel `GOROOT`

*By default*, setelah proses instalasi Go selesai, secara otomatis akan muncul *environment variable* `GOROOT`. Isi dari variabel ini adalah lokasi di mana Go ter-*install*.

Sebagai contoh di Windows, ketika Go di-*install* di `C:\go`, maka path tersebut akan menjadi isi dari `GOROOT`.

Silakan gunakan command `go env` untuk melihat informasi konfigurasi *environment* yang ada.

## A.2.3. Instalasi Go *Unstable*/*Development*

Jika pembaca tertarik untuk mencoba versi development Go, ingin mencoba fitur yang belum dirilis secara official, ada beberapa cara:

- Instalasi dengan *build from source* https://go.dev/doc/install/source
- Gunakan command `go install`, contohnya seperti `go install golang.org/dl/go1.18beta1@latest`. Untuk melihat versi unstable yang bisa di-install silakan merujuk ke https://go.dev/dl/#unstable

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
