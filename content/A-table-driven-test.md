# A.59. Table-Driven Test

*Table-driven test* adalah teknik penulisan test di Go yang menggunakan **tabel data** (slice of struct) sebagai test case. Teknik ini sangat populer di komunitas Go karena memungkinkan kita menulis banyak test case dengan kode yang lebih bersih dan mudah di-maintain.

Pada chapter ini kita akan belajar cara membuat table-driven test menggunakan `t.Run()` untuk sub-tests, dan juga cara membuat test yang lebih terstruktur.

## A.59.1. Masalah dengan Test Konvensional

Pada chapter sebelumnya (A.58. Unit Test), kita sudah belajar membuat test seperti ini:

```go
func TestHitungVolume(t *testing.T) {
    kubus := Kubus{Sisi: 4}
    if kubus.Volume() != 64 {
        t.Errorf("Volume kubus salah, seharusnya 64, dapat %f", kubus.Volume())
    }
}
```

Masalahnya, jika kita ingin mengetes banyak skenario (sisi = 0, sisi negatif, sisi desimal, dll), kita harus menulis banyak fungsi test terpisah atau banyak blok `if` di dalam satu fungsi. Kode menjadi panjang dan repetitif.

## A.59.2. Solusi: Table-Driven Test

Dengan table-driven test, kita mendefinisikan semua test case dalam sebuah slice, lalu mengiterasinya. Contoh:

```go
package main

import (
    "math"
    "testing"
)

type Kubus struct {
    Sisi float64
}

func (k Kubus) Volume() float64 {
    return math.Pow(k.Sisi, 3)
}

func (k Kubus) Luas() float64 {
    return math.Pow(k.Sisi, 2) * 6
}

func TestKubus_TableDriven(t *testing.T) {
    // definisikan test cases dalam bentuk tabel
    tests := []struct {
        name           string
        sisi           float64
        expectedVolume float64
        expectedLuas   float64
    }{
        {
            name:           "sisi 0",
            sisi:           0,
            expectedVolume: 0,
            expectedLuas:   0,
        },
        {
            name:           "sisi 1",
            sisi:           1,
            expectedVolume: 1,
            expectedLuas:   6,
        },
        {
            name:           "sisi 4",
            sisi:           4,
            expectedVolume: 64,
            expectedLuas:   96,
        },
        {
            name:           "sisi 2.5",
            sisi:           2.5,
            expectedVolume: 15.625,
            expectedLuas:   37.5,
        },
        {
            name:           "sisi 10",
            sisi:           10,
            expectedVolume: 1000,
            expectedLuas:   600,
        },
    }

    for _, tt := range tests {
        // gunakan t.Run untuk membuat sub-test
        t.Run(tt.name, func(t *testing.T) {
            k := Kubus{Sisi: tt.sisi}

            if k.Volume() != tt.expectedVolume {
                t.Errorf("Volume kubus sisi %f = %f, want %f", tt.sisi, k.Volume(), tt.expectedVolume)
            }

            if k.Luas() != tt.expectedLuas {
                t.Errorf("Luas kubus sisi %f = %f, want %f", tt.sisi, k.Luas(), tt.expectedLuas)
            }
        })
    }
}
```

Jalankan test di atas:

```bash
go test -v -run TestKubus_TableDriven
```

Output:

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

Perhatikan bagaimana setiap test case dijalankan sebagai **sub-test** terpisah, memudahkan kita melihat mana yang pass dan mana yang fail.

## A.59.3. Sub-Test dengan `t.Run()`

Fungsi `t.Run()` menerima dua parameter:
- **nama sub-test** (string)
- **fungsi test** (`func(t *testing.T)`)

Keuntungan menggunakan `t.Run()`:
- Setiap test case punya output terpisah saat `-v` flag digunakan
- Jika satu test case gagal, yang lain tetap dijalankan
- Bisa menjalankan sub-test tertentu saja: `go test -v -run TestKubus_TableDriven/sisi_4`

## A.59.4. Contoh Lain: Fungsi dengan Banyak Kondisi

Berikut contoh table-driven test untuk fungsi yang mengkategorikan nilai:

```go
package main

import "testing"

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

func TestKategorikanNilai(t *testing.T) {
    tests := []struct {
        name     string
        nilai    int
        expected string
    }{
        {"nilai sempurna", 100, "A"},
        {"nilai A minimum", 90, "A"},
        {"nilai A batas", 89, "B"},
        {"nilai B minimum", 80, "B"},
        {"nilai B batas", 79, "C"},
        {"nilai C minimum", 70, "C"},
        {"nilai C batas", 69, "D"},
        {"nilai D minimum", 60, "D"},
        {"nilai D batas", 59, "E"},
        {"nilai nol", 0, "E"},
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

## A.59.5. Menjalankan Sub-Test Tertentu

Salah satu keuntungan besar dari table-driven test adalah bisa menjalankan sub-test tertentu saja menggunakan pattern matching:

```bash
# Jalankan semua sub-test yang mengandung kata "batas"
go test -v -run TestKategorikanNilai/batas

# Jalankan sub-test spesifik
go test -v -run TestKategorikanNilai/nilai_sempurna
```

## A.59.6. Best Practices

Beberapa best practices dalam menulis table-driven test di Go:

1. **Gunakan nama yang deskriptif** untuk setiap test case, agar mudah diidentifikasi saat gagal
2. **Sertakan edge cases** seperti nilai 0, nilai negatif, empty string, nil, dll
3. **Sertakan test case yang diharapkan gagal** jika fungsi memiliki validasi input
4. **Gunakan struct anonim** untuk test case agar kode tetap bersih
5. **Jangan lupa `t.Parallel()`** jika test case bersifat independen dan ingin dijalankan paralel

Contoh dengan parallel test:

```go
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
        tt := tt // capture range variable
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

> Perhatikan baris `tt := tt` di dalam loop. Ini diperlukan untuk menangkap variabel range agar tidak terjadi race condition saat test dijalankan secara paralel. Pada Go 1.22+, hal ini tidak lagi diperlukan karena variabel loop sudah di-scoped per iterasi.

---

- [Go Testing Documentation](https://pkg.go.dev/testing)
- [TableDrivenTests - Go Wiki](https://github.com/golang/go/wiki/TableDrivenTests)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.59-table-driven-test">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.59...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
