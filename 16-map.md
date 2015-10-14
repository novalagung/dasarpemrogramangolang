# Map

**Map** adalah tipe data asosiatif yang ada di Golang. Bentuknya *key-value*, artinya setiap data (atau value) yang disimpan, disiapkan juga key-nya. Key tersebut harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan data atau item yang tersimpan.

Kalau dilihat, map mirip seperti slice, hanya saja indeks yang digunakan untuk pengaksesan bisa ditentukan sendiri tipe item-nya (indeks tersebut adalah key).

## Penggunaan Map

Cara menggunakan map adalah dengan menuliskan keyword `map` diikuti tipe data key dan value-nya. Agar lebih mudah dipahami, silakan perhatikan contoh di bawah ini.

```go
var chicken map[string]int
chicken = map[string]int{}

chicken["januari"] = 50
chicken["februari"] = 40

fmt.Println("januari", chicken["januari"]) // januari 50
fmt.Println("mei",     chicken["mei"])     // mei 0
```

Variabel `chicken` dideklarasikan sebagai map, dengan tipe data key adalah `string` dan value nya `int`. Dari kode tersebut bisa dilihat bagaimana cara penulisan tipe data map. `map[string]int` maknanya adalah tipe data `map` dengan key bertipe `string` dan value bertipe `int`.

Default nilai variabel `map` adalah `nil`. Oleh karena itu perlu dilakukan inisialisasi nilai default di awal, caranya cukup dengan tambahkan kurung kurawal pada akhir tipe, contoh seperti pada kode di atas: `map[string]int{}`.

Cara menge-set nilai pada variabel map cukup mudah, tinggal panggil variabel-nya, sisipkan `key` pada kurung siku variabel (mirip seperti cara pengaksesan elemen slice), lalu isi nilainya. Sedangkan cara pengambilan value adalah cukup dengan menyisipkan `key` pada kurung siku variabel.

Pengisian data pada map bersifat **overwrite**, ketika variabel sudah memiliki item dengan key yang sama, maka value-nya akan ditimpa dengan yang baru. 

![Pengaksesan data map](images/16_1_map_set_get.png)

Pada pengaksesan item menggunakan key yang belum tersimpan di map, akan dikembalikan nilai default tipe data value-nya. Contohnya seperti pada kode di atas, `chicken["mei"]` menghasilkan nilai 0 (nilai default tipe `int`), karena belum ada item yang tersimpan menggunakan key `"mei"`.

## Inisialisasi Nilai Map

Nilai variabel bertipe map bisa didefinisikan di awal, caranya dengan menambahkan kurung kurawal setelah tipe data, lalu menuliskan key dan value didalamnya. Cara ini sekilas mirip dengan definisi nilai array/slice namun dalam bentuk key-value.

```go
// cara vertikal
var chicken = map[string]int{"januari": 50, "februari": 40}

// cara horizontal
var chicken = map[string]int{
    "januari":  50,
    "februari": 40,
}
```

Key dan value dituliskan dengan pembatas tanda titik dua (`:`). Sedangkan tiap itemnya dituliskan dengan pembatas tanda koma (`,`). Khusus deklarasi dengan gaya vertikal, setelah item terakhir perlu dituliskan tanda koma (`,`).

Variabel map bisa diinisialisasi dengan tanpa nilai awal, caranya cukup menggunakan tanda kurung kurawal, contoh: `map[string]int{}`. Atau bisa juga dengan menggunakan keyword `make` dan `new`. Contohnya bisa dilihat pada kode berikut. Ketiga cara di bawah ini intinya adalah sama.

```go
var chicken1 = map[string]int{}
var chicken2 = make(map[string]int)
var chicken3 = *new(map[string]int)
```

Khusus inisialisasi data menggunakan keyword `new`, yang dihasilkan adalah data pointer. Untuk mengambil nilai aslinya bisa dengan menggunakan tanda asterisk (`*`). Topik pointer akan dibahas lebih detail ketika sudah masuk bab 22.

## Iterasi Item Map Menggunakan `for` - `range`

Item variabel `map` bisa di iterasi menggunakan `for range`. Cara penerapannya masih sama seperti penerapan `for` - `range` pada slice, pembedanya data yang dikembalikan di tiap perlangan adalah key dan value, bukan indeks dan elemen. Contohnya bisa dilihat di kode berikut.

```go
var chicken = map[string]int{
    "januari":  50,
    "februari": 40,
    "maret":    34,
    "april":    67,
}

for key, val := range chicken {
    fmt.Println(key, "  \t:", val)
}
```

![Perulangan Map](images/16_2_map_for_range.png)

## Menghapus Item Map

Fungsi `delete()` digunakan untuk menghapus item dengan key tertentu pada variabel map. Cara penggunaannya, dengan memasukan objek map dan key item yang ingin dihapus sebagai parameter.

```go
fmt.Println(len(chicken), "items") // 4 items
fmt.Println(chicken)

delete(chicken, "januari")

fmt.Println(len(chicken), "items") // 3 items
fmt.Println(chicken)
```

Item yang memiliki key `"januari"` dalam variabel `chicken` akan dihapus.

![Hapus item Map](images/16_3_map_delete_item.png)

Fungsi `len()` jika digunakan pada map akan mengembalikan jumlah elemen.

## Deteksi Keberadaan Item Menggunakan Key

Ada cara untuk mengetahui apakah dalam sebuah variabel map terdapat item dengan key tertentu atau tidak, yaitu dengan memanfaatkan 2 variabel sebagai penampung nilai kembalian pengaksesan item. Variabel ke-2 akan menampung nilai `bool` yang menunjukkan ada atau tidaknya item yang dicari.

```go
var value, isExist = chicken["mei"]

if isExist {
    fmt.Println(value)
}
```
