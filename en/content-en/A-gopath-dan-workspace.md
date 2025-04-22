# A.4. GOPATH dan Workspace

Pada chapter ini kita akan belajar tentang apa itu GOPATH beserta cara setupnya.

> ⚠️ INFORMASI ⚠️
> 
> Setup Go project menggunakan GOPATH kurang dianjurkan untuk Go versi terbaru. Lebih baik gunakan [A.3. Setup Go Modules](/A-setup-go-project-dengan-go-modules.html).
> 
> Namun meski demikian, bukan berarti GOPATH tidak berguna sama sekali, jadi silakan ikuti panduan berikut jika diperlukan.

## A.4.1. Variabel `GOPATH`

**GOPATH** adalah variabel yang digunakan oleh Go sebagai rujukan lokasi di mana semua folder project disimpan (kecuali untuk yg diinisialisasi menggunakan Go Modules). GOPATH berisikan 3 buah sub-folder: `src`, `bin`, dan `pkg`.

Project di Go bisa ditempatkan dalam `$GOPATH/src`. Sebagai contoh anda ingin membuat project dengan nama `belajar`, maka **harus** dibuatkan sebuah folder dengan nama `belajar`, ditempatkan dalam `src` (`$GOPATH/src/belajar`).

> Path separator yang digunakan sebagai contoh di buku ini adalah slash `/`. Khusus pengguna Windows, path separator adalah backslash `\`.

## A.4.2. Setup Workspace

Lokasi folder yang akan dijadikan sebagai workspace bisa ditentukan sendiri. Anda bisa menggunakan alamat folder mana saja, bebas, tapi jangan gunakan path tempat di mana Go ter-*install* (tidak boleh sama dengan `GOROOT`). Lokasi tersebut harus didaftarkan dalam path variable dengan nama `GOPATH`. Sebagai contoh, penulis memilih path `$HOME/Documents/go`, maka saya daftarkan alamat tersebut. Caranya:

 - Bagi pengguna **Windows**, tambahkan path folder tersebut ke **path variable** dengan nama `GOPATH`. Setelah variabel terdaftar, cek apakah path sudah terdaftar dengan benar.

    > Sering terjadi `GOPATH` tidak dikenali meskipun variabel sudah didaftarkan. Jika hal seperti ini terjadi, restart CMD, lalu coba lagi.

 - Bagi pengguna Mac OS, export path ke `~/.bash_profile`. Untuk Linux, export ke `~/.bashrc`

    ```bash
    $ echo "export GOPATH=$HOME/Documents/go" >> ~/.bash_profile
    $ source ~/.bash_profile
    ```

    Cek apakah path sudah terdaftar dengan benar.

    ![Pengecekan <code>GOPATH</code> di sistem operasi non-Windows](images/A_gopath_workspace_1_path.png)

Setelah `GOPATH` berhasil dikenali, perlu disiapkan 3 buah sub folder di dalamnya, dengan kriteria sebagai berikut:

 - Folder `src`, adalah path di mana project Go disimpan
 - Folder `pkg`, berisi file hasil kompilasi
 - Folder `bin`, berisi file executable hasil build

![Struktur folder dalam worskpace](images/A_gopath_workspace_2_workspace.png)

Struktur di atas merupakan struktur standar workspace Go. Jadi pastikan penamaan dan hirarki folder adalah sama.

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
