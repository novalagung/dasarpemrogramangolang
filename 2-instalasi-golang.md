# Instalasi Golang

Hal pertama yang perlu dilakukan sebelum bisa menggunakan Golang adalah menginstalnya terlebih dahulu. Cara instalasinya berbeda-beda untuk tiap jenis sistem operasi. Panduan instalasi Golang sebenarnya sudah disediakan, bisa dilihat di situs official-nya [http://golang.org/doc/install#install](http://golang.org/doc/install#install).

Bab ini merupakan ringkasan dari panduan instalasi yang disediakan oleh Golang.

> Di buku ini versi Golang yang digunakan adalah **1.4.2** 

## Instalasi Golang di Windows

 1. Download terlebih dahulu installer-nya. Pilih sesuai jenis bit prosesor yang digunakan.

    - 32bit => [go1.4.2.wind\*ws-386.msi](https://storage.googleapis.com/golang/go1.4.2.windows-386.msi)
    - 64bit => [go1.4.2.wind\*ws-amd64.msi](https://storage.googleapis.com/golang/go1.4.2.windows-amd64.msi)


 2. Setelah ter-download, jalankan installer, klik **next** sampai proses instalasi selesai. Default-nya Golang akan terinstal di `c:\go\bin`. Path tersebut akan secara otomatis terdaftar di **path variable**.

 3. Buka **Command Prompt** / **CMD**, eksekusi perintah berikut untuk mengetes apakah Golang sudah terinstal dengan benar.

    ```
    $ go version
    ```

 4. Jika output command di atas adalah versi Golang yang di-instal, maka instalasi berhasil.

> Sering terjadi command `go version` tidak bisa dijalankan meskipun Golang sudah terinstal. Solusinya adalah dengan restart CMD (close CMD, lalu buka kembali). Setelah itu coba jalankan sekali lagi command tersebut.

## Instalasi Golang di OSX

Cara termudah instalasi Golang di **\*SX** adalah dengan menggunakan [homebrew](http://brew.sh/). Homebrew sendiri adalah **package manager** khusus untuk O\*SX (mirip seperti `apt-get` milik Ubuntu).

Di bawah ini merupakan langkah instalasi Golang di \*SX menggunakan homebrew.

 1. Install terlebih dahulu homebrew (jika belum ada), dengan cara mengeksekusi perintah berikut di **terminal**.

    ```
    $ ruby -e "$(curl -fsSL http://git.io/pVOl)"
    ```

 2. Install Golang menggunakan command `brew`.

    ```
    $ brew install go
    ```

 3. Selanjutnya, eksekusi perintah di bawah ini untuk mengetes apakah golang sudah terinstal dengan benar.

    ```
    $ go version
    ```

 4. Jika output command di atas adalah versi Golang yang di-instal, maka instalasi berhasil.

## Instalasi Golang di Ubuntu

Cara menginstal Golang di **Ub\*ntu** bisa dengan memanfaatkan `apt-get`. Silakan ikuti petunjuk di bawah ini.

 1. Jalankan command berikut di **terminal**.

    ```
    $ sudo add-apt-repository ppa:gophers/go
    $ sudo apt-get update
    $ sudo apt-get install Golang-stable
    ```

 2. Setelah instalasi selesai, eksekusi perintah di bawah ini untuk mengetes apakah Golang sudah terinstal dengan benar.

    ```
    $ go version
    ```

 3. Jika output command di atas adalah versi Golang yang di-instal, maka instalasi berhasil.

## Instalasi Golang di Distro Linux Lain

 1. Download archive berikut, pilih sesuai jenis bit komputer anda.

    - 32bit => [go1.4.2.lin\*x-386.tar.gz](https://storage.googleapis.com/golang/go1.4.2.linux-386.tar.gz)
    - 64bit => [go1.4.2.lin\*x-amd64.tar.gz](https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz)


 2. Buka **terminal**, ekstrak archive tersebut ke /usr/local. Setelah itu export path-nya. Gunakan command di bawah ini untuk melakukan hal tersebut.

    ```
    $ tar zxvf go1.4.2.lin*x-....tar.gz -C /usr/local
    $ export PATH=$PATH:/usr/local/go/bin
    ```

 3. Selanjutnya, eksekusi perintah berikut untuk mengetes apakah Golang sudah terinstal dengan benar.

    ```
    $ go version
    ```

 4. Jika output command di atas adalah versi Golang yang di-instal, maka instalasi berhasil.
