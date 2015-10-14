# GOPATH Dan Workspace

Setelah Golang berhasil di-instal, ada hal yang perlu disiapkan sebelum bisa masuk ke sesi pembuatan aplikasi. Yaitu, setup workspace untuk proyek-proyek yang akan dibuat. Dan di bab ini kita akan belajar bagaimana caranya.

## Variabel `GOPATH`

**GOPATH** adalah variabel yang digunakan oleh Golang sebagai rujukan lokasi dimana semua folder proyek disimpan. Gopath berisikian 3 buah sub folder: `src`, `bin`, dan `pkg`.

Proyek di Golang harus ditempatkan pada path `$GOPATH/src`. Sebagai contoh kita ingin membuat proyek dengan nama `belajar`, maka harus dibuatkan sebuah folder dengan nama `belajar`, yang folder tersebut ditempatkan di `src` (`$GOPATH/src/belajar`). Nantinya semua file untuk keperluan proyek yang bersangkutan ditempatkan disana.

> Path separator yang digunakan sebagai contoh di buku ini adalah slash `/`. Khusus pengguna Windows, path separator adalah backslah `\`.

## Setup Workspace

Lokasi atau alamat folder yang akan dijadikan sebagai workspace bisa ditentukan sendiri. Anda bisa menggunakan alamat folder mana saja, bebas. Lokasi tersebut perlu disimpan kedalam path variable dengan nama `GOPATH`. Sebagai contoh, saya memilih path `$HOME/Documents/go`, maka saya daftarkan alamat tersebut. Caranya:

 - Bagi pengguna **Windows**, tambahkan path folder tersebut ke **path variable** dengan nama `GOPATH`. Setelah variabel didaftarkan, cek apakah path sudah terdaftar dengan benar.

    > Sering terjadi `GOPATH` tidak dikenali meskipun variabel sudah didaftarkan. Jika hal seperti ini terjadi, restart command prompt anda, lalu coba lagi.

 - Bagi pengguna OSX, BSD, dan linux, gunakan keyword `export` untuk mendaftarkan `GOPATH`.

    ```
    $ export GOPATH=$HOME/Documents/go
    ```

    Setelah variabel didaftarkan, cek apakah path sudah terdaftar dengan benar.

    ![Pengecekan `GOPATH` di sistem operasi non-Windows](images/3_1_path.png)

Setelah `GOPATH` berhasil dikenali, perlu disiapkan 3 buah sub folder didalamnya dengan kriteria sebagai berikut:

 - Folder `src`, adalah path dimana proyek golang disimpan
 - Folder `pkg`, berisi file hasil kompilasi
 - Folder `bin`, berisi file executable hasil build

![Struktur folder dalam worskpace](images/3_2_workspace.png)

Struktur diatas merupakan struktur standar workspace Golang. Jadi pastikan penamaan dan hirarki folder adalah sama.
