# A.48. Arguments & Flag

**Arguments** adalah data opsional yang disisipkan ketika eksekusi program. Sedangkan **flag** merupakan ekstensi dari argument. Dengan flag, penulisan argument menjadi lebih rapi dan terstruktur.

Pada chapter ini kita akan belajar tentang penggunaan arguments dan flag.

## A.48.1. Penggunaan Arguments

Data arguments bisa didapat lewat variabel `os.Args` (package `os` perlu di-import terlebih dahulu). Data tersebut tersimpan dalam bentuk array dengan pemisah adalah tanda spasi.

Berikut merupakan contoh penggunaannya.

```go
package main

import "fmt"
import "os"

func main() {
    var argsRaw = os.Args
    fmt.Printf("-> %#v\n", argsRaw)
    // -> []string{".../bab45", "banana", "potato", "ice cream"}

    var args = argsRaw[1:]
    fmt.Printf("-> %#v\n", args)
    // -> []string{"banana", "potatao", "ice cream"}
}
```

Pada saat eksekusi program disisipkan juga argument-nya. Sebagai contoh disisipkan 3 buah data sebagai argumen, yaitu: `banana`, `potato`, dan `ice cream`.

Untuk eksekusinya sendiri bisa menggunakan `go run` ataupun dengan cara build-execute.

 - Menggunakan `go run`

    ```
    go run bab45.go banana potato "ice cream"
    ```

 - Menggunakan `go build`

    ```
    go build bab45.go
    $ ./bab45 banana potato "ice cream"
    ```

Variabel `os.Args` mengembalikan tak hanya arguments saja, tapi juga path file executable (jika eksekusi-nya menggunakan `go run` maka path akan merujuk ke folder temporary). Gunakan `os.Args[1:]` untuk mengambil slice arguments-nya saja.

![Pemanfaatan arguments](images/A_cli_flag_arg_1_argument.png)

Bisa dilihat pada kode di atas, bahwa untuk data argumen yang ada karakter spasi nya (<code> </code>), maka harus diapit tanda petik (`"`), agar tidak dideteksi sebagai 2 argumen.

## A.48.2. Penggunaan Flag

Flag memiliki kegunaan yang sama seperti arguments, yaitu untuk *parameterize* eksekusi program, dengan penulisan dalam bentuk key-value. Berikut merupakan contoh penerapannya.

```go
package main

import "flag"
import "fmt"

func main() {
    var name = flag.String("name", "anonymous", "type your name")
    var age = flag.Int64("age", 25, "type your age")

    flag.Parse()
    fmt.Printf("name\t: %s\n", *name)
    fmt.Printf("age\t: %d\n", *age)
}
```

Cara penulisan arguments menggunakan flag:

```
go run bab45.go -name="john wick" -age=28
```

Tiap argument harus ditentukan key, tipe data, dan nilai default-nya. Contohnya seperti pada `flag.String()` di atas. Agar lebih mudah dipahami, mari kita bahas kode berikut.

```go
var dataName = flag.String("name", "anonymous", "type your name")
fmt.Println(*dataName)
```

Kode tersebut maksudnya adalah, disiapkan flag bertipe `string`, dengan key adalah `name`, dengan nilai default `"anonymous"`, dan keterangan `"type your name"`. Nilai flag nya sendiri akan disimpan ke dalam variabel `dataName`.

Nilai balik fungsi `flag.String()` adalah string pointer, jadi perlu di-*dereference* terlebih dahulu agar bisa mendapatkan nilai aslinya (`*dataName`).

![Contoh penggunaan flag](images/A_cli_flag_arg_2_flag.png)

Flag yang nilainya tidak di set, secara otomatis akan mengembalikan nilai default.

Tabel berikut merupakan macam-macam fungsi flag yang tersedia untuk tiap jenis tipe data.

| Nama Fungsi                                | Return Value     |
|:------------------------------------------ |:---------------- |
| `flag.Bool(name, defaultValue, usage)`     | `*bool`          |
| `flag.Duration(name, defaultValue, usage)` | `*time.Duration` |
| `flag.Float64(name, defaultValue, usage)`  | `*float64`       |
| `flag.Int(name, defaultValue, usage)`      | `*int`           |
| `flag.Int64(name, defaultValue, usage)`    | `*int64`         |
| `flag.String(name, defaultValue, usage)`   | `*string`        |
| `flag.Uint(name, defaultValue, usage)`     | `*uint`          |
| `flag.Uint64(name, defaultValue, usage)`   | `*uint64`        |

## A.48.3. Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data

Sebenarnya ada 2 cara deklarasi flag yang bisa digunakan, dan cara di atas merupakan cara pertama.

Cara kedua mirip dengan cara pertama, perbedannya adalah kalau di cara pertama nilai pointer flag dikembalikan lalu ditampung variabel. Sedangkan pada cara kedua, nilainya diambil lewat parameter pointer.

Agar lebih jelas perhatikan contoh berikut.

```go
// cara ke-1
var data1 = flag.String("name", "anonymous", "type your name")
fmt.Println(*data1)

// cara ke-2
var data2 string
flag.StringVar(&data2, "gender", "male", "type your gender")
fmt.Println(data2)
```

Tinggal tambahkan akhiran `Var` pada pemanggilan nama fungsi flag yang digunakan (contoh `flag.IntVar()`, `flag.BoolVar()`, dll), lalu disisipkan referensi variabel penampung flag sebagai parameter pertama.

Kegunaan dari parameter terakhir method-method flag adalah untuk memunculkan hints atau petunjuk arguments apa saja yang bisa dipakai, ketika argument `--help` ditambahkan saat eksekusi program.

![Contoh penggunaan flag](images/A_cli_flag_arg_3_flag_info.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.48-cli-arguments-flag">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.48...</a>
</div>

---

<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>
