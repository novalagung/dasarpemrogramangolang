# GOPATH Dan Workspace

Ada hal yang perlu disiapkan sebelum bisa masuk ke sesi pembuatan aplikasi menggunakan Golang. Diantaranya adalah setup workspace untuk Project yang akan dibuat. Dan di bab ini kita akan belajar bagaimana caranya.

## Variabel `GOPATH`

**GOPATH** adalah variabel yang digunakan oleh Golang sebagai rujukan lokasi dimana semua folder project disimpan. Gopath berisikian 3 buah sub folder: `src`, `bin`, dan `pkg`.

Project di Golang harus ditempatkan dalam `$GOPATH/src`. Sebagai contoh anda ingin membuat project dengan nama `belajar`, maka harus dibuatkan sebuah folder dengan nama `belajar`, ditempatkan dalam `src` (`$GOPATH/src/belajar`).

> Path separator yang digunakan sebagai contoh di buku ini adalah slash `/`. Khusus pengguna Windows, path separator adalah backslah `\`.

## Setup Workspace

Lokasi folder yang akan dijadikan sebagai workspace bisa ditentukan sendiri. Anda bisa menggunakan alamat folder mana saja, bebas, tapi jangan gunakan path dimana golang di-install. Lokasi tersebut harus didaftarkan dalam path variable dengan nama `GOPATH`. Sebagai contoh, penulis memilih path `$HOME/Documents/go`, maka saya daftarkan alamat tersebut. Caranya:

 - Bagi pengguna **Windows**, tambahkan path folder tersebut ke **path variable** dengan nama `GOPATH`. Setelah variabel terdaftar, cek apakah path sudah terdaftar dengan benar.

    > Sering terjadi `GOPATH` tidak dikenali meskipun variabel sudah didaftarkan. Jika hal seperti ini terjadi, restart CMD, lalu coba lagi.

 - Bagi pengguna Mac OS, export path ke `~/.bash_profile`. Untuk Linux, export ke `~/.bashrc`

    ```
    $ echo "export GOPATH=$HOME/Documents/go" >> ~/.bash_profile
    $ source ~/.bash_profile
    ```

    Cek apakah path sudah terdaftar dengan benar.

    ![Pengecekan `GOPATH` di sistem operasi non-Wind\*ws](images/3_1_path.png)

Setelah `GOPATH` berhasil dikenali, perlu disiapkan 3 buah sub folder didalamnya, dengan kriteria sebagai berikut:

 - Folder `src`, adalah path dimana project golang disimpan
 - Folder `pkg`, berisi file hasil kompilasi
 - Folder `bin`, berisi file executable hasil build

![Struktur folder dalam worskpace](images/3_2_workspace.png)

Struktur diatas merupakan struktur standar workspace Golang. Jadi pastikan penamaan dan hirarki folder adalah sama.
