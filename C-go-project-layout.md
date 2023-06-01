# C.1. Project Layout

Pada chapter ini kita akan belajar cara melakukan pertukaran data lewat FTP (File Transfer Protocol) menggunakan Golang. 

> Definisi mengenai FTP sendiri adalah sebuah protokol network standar yang digunakan untuk pertukaran atau transfer data antar client dan server.

Sebelum memulai, ada satu hal penting yang perlu dipersiapkan, yaitu sebuah server dengan FTP server ter-install. Jika tidak ada, bisa menggunakan library [ftpd](https://github.com/goftp/ftpd) untuk set up ftp server di local (untuk keperluan belajar).

Dalam server tersebut, siapkan beberapa file dan folder dengan struktur sebagai berikut.

![FTP folder structure](images/C.26_1_ftp_folder_structure.png)

 - File `test1.txt`, isi dengan apapun.
 - File `test2.txt`, isi dengan apapun.
 - File `somefolder/test3.txt`, isi dengan apapun.
 - File `movie.mp4`, gunakan file seadanya.

Library FTP client yang kita gunakan adalah [github.com/jlaffaye/ftp](https://github.com/jlaffaye/ftp).

---

<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>
