# A.2. Instalasi Golang

Hal pertama yang perlu dilakukan sebelum bisa menggunakan Golang adalah meng-install-nya terlebih dahulu. Panduan instalasi sebenarnya sudah disediakan di situs official Golang [http://golang.org/doc/install#install](http://golang.org/doc/install#install).

Disini penulis mencoba meringkas petunjuk instalasi di link tersebut, agar lebih mudah untuk diikuti terutama untuk pembaca yang baru belajar.

> Golang yang digunakan adalah versi **1.8.3**. Direkomendasikan menggunakan versi tersebut, atau versi lain minimal **1.4.2** ke atas.

Link untuk download installer golang: https://golang.org/dl/. Anda bisa langsung unduh dari URL tersebut lalu lakukan instalasi sendiri, atau bisa mengikuti petunjuk di bab ini.

## A.2.1. Instalasi Golang di Windows

 1. Download terlebih dahulu installer-nya. Pilih sesuai jenis bit yang digunakan.

    - 32bit => [go1.8.3.windows-386.msi](https://storage.googleapis.com/golang/go1.8.3.windows-386.msi)
    - 64bit => [go1.8.3.windows-amd64.msi](https://storage.googleapis.com/golang/go1.8.3.windows-amd64.msi)

 2. Setelah ter-download, jalankan installer, klik **next** sampai proses instalasi selesai. By default jika anda tidak merubah path pada saat instalasi, Golang akan terinstall di terinstal di `C:\go`. Path tersebut secara otomatis didaftarkan dalam **path variable**.

 3. Buka **Command Prompt** / **CMD**, eksekusi perintah untuk mengecek versi Golang.

    ```bash
    $ go version
    ```

 4. Jika output adalah sama dengan Golang yang ter-install, menandakan instalasi berhasil.

> Sering terjadi command `go version` tidak bisa dijalankan meskipun instalasi sukses. Solusinya bisa dengan restart CMD (close CMD, lalu buka kembali). Setelah itu coba jalankan sekali lagi command tersebut.

## A.2.2. Instalasi Golang di Mac OS

Cara termudah instalasi Golang di **Mac OS** adalah menggunakan [homebrew](http://brew.sh/).

 1. Install terlebih dahulu Homebrew (jika belum ada), jalankan perintah tersebut di **terminal**.

    ```bash
    $ ruby -e "$(curl -fsSL http://git.io/pVOl)"
    ```

 2. Install Golang menggunakan command `brew`.

    ```bash
    $ brew install go
    ```

 3. Tambahkan path ke environment variable.

    ```bash
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
    $ source ~/.bash_profile
    ```

 4. Jalankan command untuk mengecek versi Golang.

    ```bash
    $ go version
    ```

 5. Jika output adalah sama dengan Golang yang ter-install, menandakan instalasi berhasil.

## A.2.3. Instalasi Golang di Linux

 1. Download archive berikut, pilih sesuai jenis bit komputer anda.

     - 32bit => [go1.8.3.linux-386.tar.gz](https://storage.googleapis.com/golang/go1.8.3.linux-386.tar.gz)
     - 64bit => [go1.8.3.linux-amd64.tar.gz](https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz)

    Download bisa dilakukan lewat CLI, menggunakan `wget` atau `curl`.

    ```bash
    $ wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
    ```

 2. Buka **terminal**, extract archive tersebut ke `/usr/local`.

    ```bash
    $ tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
    ```

 3. Setelah itu export path-nya, gunakan command di bawah ini.

    ```bash
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
    $ source ~/.bashrc
    ```

 4. Selanjutnya, eksekusi perintah berikut untuk mengetes apakah Golang sudah terinstal dengan benar.

    ```bash
    $ go version
    ```

 5. Jika output adalah sama dengan Golang yang ter-install, menandakan instalasi berhasil.

## A.2.4. Variabel `GOROOT`

By default, setelah proses instalasi Golang selesai, secara otomatis akan muncul environment variabel bernama `GOROOT`. Isi dari environment variabel ini adalah path/folder/lokasi dimana golang di install.

Sebagai contoh di Windows, ketika golang di-install di `C:\go`, maka path tersebut akan menjadi isi dari `GOROOT`.
