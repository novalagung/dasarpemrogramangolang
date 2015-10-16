# Instalasi Editor

Pembuatan program menggunakan bahasa Golang, akan lebih maksimal jika didukung editor atau **IDE** yang pas. Ada cukup banyak pilihan bagus yang bisa dipertimbangkan, diantaranya: Intellij, Netbeans, Atom, Brackets, dan lainnya.

Pada saat menulis buku ini, editor yang saya gunakan adalah **Sublime Text 3**. Editor ini ringan, mudah didapat, dan memiliki cukup banyak plugin. Anda bisa memilih editor yang sama dengan yang digunakan di buku ini, atau editor lainnya, bebas, yang penting nyaman ketika digunakan.

Bagi yang memilih Sublime, saya sarankan untuk menginstall plugin bernama **GoSublime**. Plugin ini menyediakan banyak sekali fitur yang sangat membantu proses pengembangan aplikasi menggunakan Golang. Diantaranya seperti *code completion*, *lint* (deteksi kesalahan di level sintaks), perapian kode otomatis, dan lainnya.

Di bab ini akan dijelaskan bagaimana cara instalasi editor Sublime, package control, dan plugin GoSublime.

## Instalasi Editor Sublime Text

 1. Download **Sublime Text versi 3** di [http://www.sublimetext.com/3](http://www.sublimetext.com/3), pilih sesuai dengan sistem operasi yang digunakan.
 2. Setelah ter-download, buka file tersebut untuk memulai instalasi.
 3. Setelah instalasi selesai, jalankan aplikasi.

![Tampilan Sublime Text 3](images/4_1_sublime_text.png)

## Instalasi Package Control

Package control merupakan aplikasi *3rd party* untuk Sublime, digunakan untuk mempermudah instalasi plugin. Default-nya Sublime tidak menyediakan aplikas ini, kita perlu menginstalnya sendiri. Silakan ikuti petunjuk berikut untuk cara instalasinya.

 1. Buka situs [https://packagecontrol.io/installation](https://packagecontrol.io/installation), **copy** script yang ada di tab Sublime Text 3 (tab bagian kiri).

    ![Copy script instalasi plugin](images/4_5_plugin_control_code.png)

 2. Selanjutnya, jalankan aplikasi Sublime, klik menu **View > Show Console**, lalu **paste** script yang sudah di-copy tadi, ke inputan kecil di bagian bawah editor. Lalu tekan Enter.

    ![Show console, paste script instalasi package control](images/4_2_install_package_control.png)

 3. Tunggu hingga proses instalasi selesai. Perhatikan karakter sama dengan *(=)* di bagian kiri bawah editor yang bergerak-gerak. Jika karakter tersebut menghilang, menandakan bahwa proses instalasi sudah selesai.

 4. Setelah selesai, tutup aplikasi, lalu buka kembali. Package control sudah berhasil di-install.

## Instalasi Plugin GoSublime

Dengan memanfaatkan package control, instalasi plugin akan menjadi lebih mudah. Berikut merupakan langkah instalasi plugin GoSublime.

 1. Buka Sublime, tekan **ctrl+shift+p** (atau **cmd+shift+p** untuk pengguna \*SX), akan muncul sebuah input dialog. Ketikan disana `install`, lalu tekan enter.

    ![Cara menjalankan package control](images/4_3_install_plugin.png)

 2. Akan muncul lagi input dialog lainnya, ketikkan `GoSublime` lalu tekan enter. Tunggu hingga proses selesai (acuan instalasi selesai adalah karakter sama dengan *(=)* di bagian kiri bawah editor yang sebelumnya bergerak-gerak. Ketika karakter tersebut sudah hilang, menandakan bahwa instalasi selesai).

    ![Cara meng-install GoSublime](images/4_4_install_gosublime.png)

 3. Setelah selesai, restart Sublime, plugin GoSublime sudah berhasil ter-install.
