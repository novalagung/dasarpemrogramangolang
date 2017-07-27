# Instalasi Golang

Hal pertama yang perlu dilakukan sebelum bisa menggunakan Golang adalah meng-install-nya terlebih dahulu. Panduan instalasi sebenarnya sudah disediakan di situs official Golang [http://golang.org/doc/install#install](http://golang.org/doc/install#install).

Disini penulis mencoba meringkas pandian instalasi yang telah disediakan, agar lebih mudah untuk diikuti terutama untuk pembaca yang baru belajar.

> Golang yang digunakan adalah versi **1.8.3**. Direkomendasikan menggunakan versi tersebut, atau versi lain minimal **1.4.2** ke atas.<br /><br />Perbedaan signifikan antara versi [**1.4.2**](https://golang.org/doc/go1.4.2), [**1.5**](https://golang.org/doc/go1.5), [**1.6**](https://golang.org/doc/go1.6), [**1.7**](https://golang.org/doc/go1.7) kebanyakan adalah dibagian performa, hanya sedikit update dibagian sintaks bahasa.

Link untuk download installer golang: https://golang.org/dl/. Anda bisa langsung unduh dari URL tersebut lalu lakukan instalasi sendiri, atau bisa mengikuti petunjuk berikut.

## Instalasi Golang di Windows

 1. Download terlebih dahulu installer-nya. Pilih sesuai jenis bit yang digunakan.

    - 32bit => [go1.8.3.windows-386.msi](https://storage.googleapis.com/golang/go1.8.3.windows-386.msi)
    - 64bit => [go1.8.3.windows-amd64.msi](https://storage.googleapis.com/golang/go1.8.3.windows-amd64.msi)

 2. Setelah ter-download, jalankan installer, klik **next** sampai proses instalasi selesai. By default jika anda tidak merubah path pada saat instalasi, Golang akan terinstall di terinstal di `C:\go`. Path tersebut secara otomatis didaftarkan dalam **path variable**.

 3. Buka **Command Prompt** / **CMD**, eksekusi perintah untuk mengecek versi Golang.

    ```
    $ go version
    ```

 4. Jika output adalah sama dengan Golang yang ter-install, menandakan instalasi berhasil.

> Sering terjadi command `go version` tidak bisa dijalankan meskipun instalasi sukses. Solusinya bisa dengan restart CMD (close CMD, lalu buka kembali). Setelah itu coba jalankan sekali lagi command tersebut.

## Instalasi Golang di Mac OS

Cara termudah instalasi Golang di **Mac OS** adalah menggunakan [homebrew](http://brew.sh/).

 1. Install terlebih dahulu Homebrew (jika belum ada), jalankan perintah tersebut di **terminal**.

    ```
    $ ruby -e "$(curl -fsSL http://git.io/pVOl)"
    ```

 2. Install Golang menggunakan command `brew`.

    ```
    $ brew install go
    ```

 3. Tambahkan path ke environment variable.

    ```
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
    $ source ~/.bash_profile
    ```

 4. Jalankan command untuk mengecek versi Golang.

    ```
    $ go version
    ```

 5. Jika output adalah sama dengan Golang yang ter-install, menandakan instalasi berhasil.

## Instalasi Golang di Linux

 1. Download archive berikut, pilih sesuai jenis bit komputer anda.

     - 32bit => [go1.8.3.linux-386.msi](https://storage.googleapis.com/golang/go1.8.3.linux-386.tar.gz)
     - 64bit => [go1.8.3.linux-amd64.msi](https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz)

    Download bisa dilakukan lewat CLI, menggunakan `wget` atau `curl`.

    ```
    $ wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
    ```

 2. Buka **terminal**, extract archive tersebut ke `/usr/local`.

    ```
    $ tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
    ```

 3. Setelah itu export path-nya, gunakan command di bawah ini.

    ```
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
    $ source ~/.bashrc
    ```

 4. Selanjutnya, eksekusi perintah berikut untuk mengetes apakah Golang sudah terinstal dengan benar.

    ```
    $ go version
    ```

 5. Jika output adalah sama dengan Golang yang ter-install, menandakan instalasi berhasil.
