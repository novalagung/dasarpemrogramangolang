# C.20. Write PDF File (gofpdf)

Reporting pada aplikasi web, selain ke bentuk file excel biasanya ke bentuk file pdf. Pada chapter ini kita akan mempelajari cara membuat file pdf di golang menggunakan [gofpdf](https://github.com/jung-kurt/gofpdf).

**gofpdf** adalah library yang berguna untuk membuat dokumen PDF dari golang. Penggunannya tidak terlalu sulit. Jadi mari belajar sambil praktek seperti biasanya. 

## C.20.1. Membuat PDF Menggunakan gofpdf

Pertama `go get` library-nya.

```bash
go get -u github.com/jung-kurt/gofpdf
```

Buat folder projek baru, isi main dengan kode berikut.

```go
package main

import (
    "github.com/jung-kurt/gofpdf"
    "log"
)

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Text(40, 10, "Hello, world")
    pdf.Image("./sample.png", 56, 40, 100, 0, false, "", 0, "")

    err := pdf.OutputFileAndClose("./file.pdf")
    if err != nil {
        log.Println("ERROR", err.Error())
    }
}
```

Statement `gofpdf.New()` digunakan untuk membuat objek dokumen baru. Fungsi `.New()` tersebut membutuhkan 4 buah parameter.

 1. Parameter ke-1, orientasi dokumen, apakah portrait (`P`) atau landscape (`L`).
 2. Parameter ke-2, satuan ukuran yang digunakan, `mm` berarti milimeter.
 3. Parameter ke-3, ukuran dokumen, kira pilih A4.
 4. Parameter ke-4, path folder font.

Fungsi `.New()` mengembalikan objek PDF. Dari situ kita bisa mengakses banyak method sesuai kebutuhan, beberapa diantaranya adalah 4 buah method yang dicontohkan di atas.

#### • Method `.AddPage()`
 
Method ini digunakan untuk menambah halaman baru. Defaultnya, objek dokumen yang baru dibuat tidak memiliki halaman. Dengan memanggil `.AddPage()` maka halaman baru dibuat.

Setelah *at least* satu halaman tersedia, kita bisa lanjut ke proses tulis menulis.

#### • Method `.SetFont()`

Method ini digunakan untuk menge-set konfigurasi font dokumen. Font Family, Font Style, dan Font Size disisipkan dalam parameter secara berurutan.
 
#### • Method `.Text()`

Digunakan untuk menulis text pada koordinat tertentu. Pada kode di atas, `40` artinya `40mm` dari kiri, sedangkan `10` artinya `10mm` dari atas. Satuan milimeter digunakan karena pada saat penciptaan objek dipilih `mm` sebagai satuan.

Method ini melakukan penulisan text pada current page.

#### • Method `.Image()` 

Digunakan untuk menambahkan image. Method ini memerlukan beberapa parameter.

- Parameter ke-1 adalah path image.
- Paraketer ke-2 adalah x offset. Nilai `56` artinya `56mm` dari kiri.
- Parameter ke-3 adalah y offset. Nilai `40` artinya `40mm` dari atas.
- Parameter ke-4 adalah width gambar. Jika diisi dengan nilai lebih dari 0 maka gambar akan di-resize secara proporsional sesuai angka. Jika di-isi `0`, maka gambar akan muncul sesuai ukuran aslinya. Pada kode di atas, gambar `sample.png` digunakan, silakan gunakan gambar apa saja bebas.
- Parameter ke-5 adalah height gambar.

Sebenarnya masih banyak lagi method yang tersedia, selengkapnya cek saja di https://godoc.org/github.com/jung-kurt/gofpdf#Fpdf.

Setelah selesai bermain dengan objek pdf, gunakan `.OutputFileAndClose()` untuk menyimpan hasil sebagai file fisik PDF.

Coba jalankan aplikasi untuk melihat hasilnya. Buka generated file `file.pdf`, isinya kurang lebih seperti gambar berikut.

![Write PDF file](images/C_write_pdf_file_1_write_pdf_file.png)


---

- [gofpdf](https://github.com/jung-kurt/gofpdf), by Kurt Jung, MIT license

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.20-write-pdf-file">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.20...</a>
</div>
