# A.60. Go Modules

Pada bagian ini kita akan belajar cara pemanfaatan Go Modules manajemen project dan dependency.

Sebelumnya sudah di bahas dalam [bab A.3](/3-gopath-dan-workspace.html) bahwa project Go harus ditempatkan didalam workspace, lebih spesifiknya dalam folder `$GOPATH/src/`. Aturan ini cukup memicu perdebatan di komunitas, karena menghasilkan efek negatif terhadap beberapa hal, yang salah satunya adalah: dependency management yang dirasa susah.

Dimisalkan, ada dua buah projek yang sedang di-develop, `ProjectA` dan `ProjectB`. Keduanya depend terhadap salah satu 3rd party library yg sama, [gubrak](https://github.com/novalagung/gubrak). Di dalam `ProjectA`, versi gubrak yang digunakan adalah `v0.9.1-alpha`, sedangkan di `ProjectB` versi `v1.0.0` digunakan. Pada `ProjectA` versi yang digunakan cukup tua karena proses pengembangannya sudah agak lama, dan aplikasinya sendiri sudah stabil, jika di upgrade paksa ke gubrak versi `v1.0.0` pasti terjadi banyak error dan panic.

Kedua projek tersebut pastinya akan lookup gubrak ke direktori yang sama, yaitu `$GOPATH/src/github.com/novalagung/gubrak`. Efeknya, ketika sedang bekerja pada `ProjectA`, harus dipastikan current revision pada repository gubrak di lokal adalah sesuai dengan `v1.0.0`. Dan, ketika mengerjakan `ProjectB` maka 

tanpa `$GOPATH`.
Go Modules adalah fasilitas baru yang disediakan oleh 
Sebelum kita membahas mengenai apa itu **mutex**? ada baiknya untuk mempelajari terlebih dahulu apa itu **race condition**, karena kedua konsep ini berhubungan erat satu sama lain.
