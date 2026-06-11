# A.59. Table-Driven Test

*Table-driven test* adalah teknik penulisan test di Go menggunakan **tabel data** (slice of struct) sebagai kumpulan test case. Dengan teknik ini, banyak skenario pengujian bisa ditulis dalam satu fungsi test yang ringkas dan mudah dikembangkan.

> Teknik ini juga direkomendasikan secara resmi oleh tim Go. Baca lebih lanjut di [https://go.dev/wiki/TableDrivenTests](https://go.dev/wiki/TableDrivenTests).

## A.59.1. Masalah dengan Test Konvensional

Pada chapter sebelumnya ([A.58. Unit Test](/A-unit-test.html)), kita sudah belajar menulis test seperti ini:

```go
func TestHitungVolume(t *testing.T) {
    kubus := Kubus{4}
    if kubus.Volume() != 64 {
        t.Errorf("Volume salah: dapat %f, seharusnya 64", kubus.Volume())
    }
}
```

Masalahnya, kalau kita ingin mengetes banyak skenario (sisi = 0, sisi desimal, sisi besar, dan lainnya), kita harus menulis banyak fungsi test terpisah. Kode menjadi panjang dan repetitif, padahal logika pengujiannya sama persis.

## A.59.2. Solusi: Table-Driven Test

Dengan table-driven test, semua test case didefinisikan dalam sebuah slice, lalu diiterasikan satu per satu. Ini menghilangkan duplikasi dan memudahkan penambahan skenario baru cukup dengan menambah satu baris ke tabel.

Contoh berikut melanjutkan project dari chapter A.58. Struct `Kubus` dan method-nya sudah ada di `main.go`:

```go
package main

import "math"

type Kubus struct {
    Sisi float64
}

func (k Kubus) Volume() float64 {
    return math.Pow(k.Sisi, 3)
}

func (k Kubus) Luas() float64 {
    return math.Pow(k.Sisi, 2) * 6
}
```

Tambahkan file `main_test.go` dengan isi berikut:

```go
package main

import "testing"

func TestKubus_TableDriven(t *testing.T) {
    tests := []struct {
        name           string
        sisi           float64
        expectedVolume float64
        expectedLuas   float64
    }{
        {"sisi 0",   0,   0,      0},
        {"sisi 1",   1,   1,      6},
        {"sisi 4",   4,   64,     96},
        {"sisi 2.5", 2.5, 15.625, 37.5},
        {"sisi 10",  10,  1000,   600},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := Kubus{tt.sisi}

            if k.Volume() != tt.expectedVolume {
                t.Errorf("Volume(%v) = %v, want %v", tt.sisi, k.Volume(), tt.expectedVolume)
            }

            if k.Luas() != tt.expectedLuas {
                t.Errorf("Luas(%v) = %v, want %v", tt.sisi, k.Luas(), tt.expectedLuas)
            }
        })
    }
}
```

Jalankan test:

```bash
go test -v -run TestKubus_TableDriven
```

Output ketika semua test case lolos:

```
=== RUN   TestKubus_TableDriven
=== RUN   TestKubus_TableDriven/sisi_0
=== RUN   TestKubus_TableDriven/sisi_1
=== RUN   TestKubus_TableDriven/sisi_4
=== RUN   TestKubus_TableDriven/sisi_2.5
=== RUN   TestKubus_TableDriven/sisi_10
--- PASS: TestKubus_TableDriven (0.00s)
    --- PASS: TestKubus_TableDriven/sisi_0 (0.00s)
    --- PASS: TestKubus_TableDriven/sisi_1 (0.00s)
    --- PASS: TestKubus_TableDriven/sisi_4 (0.00s)
    --- PASS: TestKubus_TableDriven/sisi_2.5 (0.00s)
    --- PASS: TestKubus_TableDriven/sisi_10 (0.00s)
PASS
```

Setiap test case dijalankan sebagai **sub-test** terpisah lewat `t.Run()`. Fungsi ini menerima dua parameter: nama sub-test (string) dan fungsi test-nya (`func(t *testing.T)`). Keuntungannya, kalau satu test case gagal, yang lain tetap dijalankan dan hasilnya dilaporkan secara terpisah, sehingga mudah melacak bagian mana yang bermasalah.

Contoh output ketika satu test case gagal (misalnya expected volume untuk `sisi 4` diubah ke nilai yang salah):

```
=== RUN   TestKubus_TableDriven/sisi_4
    main_test.go:24: Volume(4) = 64, want 999
--- FAIL: TestKubus_TableDriven/sisi_4 (0.00s)
```

Sub-test lain tetap dijalankan dan dilaporkan secara terpisah. Ini kelebihan utama dibanding test konvensional yang langsung berhenti saat ada error pertama.

## A.59.3. Contoh Lain: Fungsi dengan Banyak Kondisi

Table-driven test sangat berguna untuk menguji fungsi yang punya banyak cabang kondisi. Mari praktikkan dengan membuat project baru.

File `main.go`:

```go
package main

func KategorikanNilai(nilai int) string {
    switch {
    case nilai >= 90:
        return "A"
    case nilai >= 80:
        return "B"
    case nilai >= 70:
        return "C"
    case nilai >= 60:
        return "D"
    default:
        return "E"
    }
}
```

File `main_test.go`:

```go
package main

import "testing"

func TestKategorikanNilai(t *testing.T) {
    tests := []struct {
        name     string
        nilai    int
        expected string
    }{
        {"nilai sempurna", 100, "A"},
        {"nilai A minimum", 90,  "A"},
        {"nilai A batas",   89,  "B"},
        {"nilai B minimum", 80,  "B"},
        {"nilai B batas",   79,  "C"},
        {"nilai C minimum", 70,  "C"},
        {"nilai C batas",   69,  "D"},
        {"nilai D minimum", 60,  "D"},
        {"nilai D batas",   59,  "E"},
        {"nilai nol",       0,   "E"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := KategorikanNilai(tt.nilai)
            if result != tt.expected {
                t.Errorf("KategorikanNilai(%d) = %s, want %s", tt.nilai, result, tt.expected)
            }
        })
    }
}
```

Perhatikan betapa mudahnya menambah skenario baru: cukup tambahkan satu baris ke slice `tests`, tanpa perlu membuat fungsi test baru.

## A.59.4. Menjalankan Sub-Test Tertentu

Flag `-run` pada `go test` mendukung *pattern matching* hingga ke level nama sub-test. Nama sub-test dibentuk dari `NamaFungsiTest/nama_sub_test`, dengan spasi otomatis diganti `_`. Pattern yang digunakan adalah regex.

```bash
# jalankan semua sub-test yang mengandung kata "batas"
go test -v -run TestKategorikanNilai/batas

# jalankan sub-test spesifik
go test -v -run TestKategorikanNilai/nilai_sempurna
```

## A.59.5. Parallel Test

Sub-test bisa dijalankan secara paralel dengan memanggil `t.Parallel()` di awal fungsi sub-test. Ini berguna untuk mempercepat eksekusi test ketika antar test case tidak saling bergantung dan tidak mengakses data yang sama.

```go
package main

import "testing"

func TestKategorikanNilai_Parallel(t *testing.T) {
    tests := []struct {
        name     string
        nilai    int
        expected string
    }{
        {"nilai A", 95, "A"},
        {"nilai B", 85, "B"},
        {"nilai C", 75, "C"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            result := KategorikanNilai(tt.nilai)
            if result != tt.expected {
                t.Errorf("KategorikanNilai(%d) = %s, want %s", tt.nilai, result, tt.expected)
            }
        })
    }
}
```

> Hindari penggunaan `t.Parallel()` jika antar test case berbagi data yang bisa diubah (misalnya variabel global atau koneksi database yang sama), karena bisa menyebabkan *race condition*.

> Sejak Go 1.22, variabel loop sudah di-scope per iterasi, sehingga tidak perlu lagi menulis `tt := tt` di dalam loop untuk menghindari masalah pada parallel test.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.59-table-driven-test">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.59...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
