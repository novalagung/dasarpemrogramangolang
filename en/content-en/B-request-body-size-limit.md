# B.27. Request Body Size Limit

Tanpa adanya limit ukuran payload, endpoint yang menerima payload (JSON, XML, file upload) akan rentan terhadap serangan _memory exhaustion_ di mana client mengirimkan data berukuran sangat besar sehingga server kehabisan memori. Go menyediakan `http.MaxBytesReader` untuk antisipasi masalah ini.

## B.27.1. Fungsi `http.MaxBytesReader()`

```go
r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
```

Fungsi `http.MaxBytesReader()` membungkus `r.Body` dengan reader yang berhenti membaca setelah jumlah byte yang ditentukan tercapai, kemudian mengembalikan error bertipe `*http.MaxBytesError`.

Parameter:

1. `w`: `http.ResponseWriter`, digunakan untuk menandai response sebagai error secara internal
2. `r.Body`: body request asli
3. `maxBytes`: batas ukuran dalam byte

> Pastikan untuk memanggil `http.MaxBytesReader` **sebelum** membaca body, bukan setelahnya.

## B.27.2. Implementasi pada JSON Payload

Pada kode berikut, fungsi `handleUpload()` merupakan HTTP handler untuk parsing payload.

```go
const maxBodyBytes = 1 << 20 // 1 MB

func handleUpload(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

    var payload map[string]any
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        var maxBytesErr *http.MaxBytesError
        if errors.As(err, &maxBytesErr) {
            http.Error(w, "request body too large", http.StatusRequestEntityTooLarge)
            return
        }
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]any{"received": payload})
}
```

Konstanta `maxBodyBytes` bernilai `1 << 20` yang nilainya adalah `1.048.576` atau setara 1 MB. Cara ini umum digunakan untuk ukuran berbasis **power-of-two** karena lebih eksplisit dibanding menulis `1048576` secara harfiah.

Di baris pertama handler, `r.Body` di-assign ulang dengan hasil `http.MaxBytesReader()` agar nantinya saat pembacaan payload limit `maxBodyBytes` diberlakukan sebagai batas. Perlu diperhatikan bahwa `MaxBytesReader` mengembalikan reader baru yang membungkus body asli, bukan memodifikasinya secara in-place, sehingga hasilnya harus di-assign kembali ke `r.Body`. Jika dilewatkan, limit tidak akan berlaku.

Statement `json.NewDecoder(r.Body).Decode()` bisa gagal karena dua hal:

1. Body terlalu besar
2. JSON tidak valid

Fungsi `errors.As()` digunakan untuk memeriksa apakah error bertipe `*http.MaxBytesError` atau bukan, sehingga kedua kondisi tersebut bisa dikembalikan dengan status code yang tepat: 413 untuk payload terlalu besar, 400 untuk JSON rusak.

## B.27.3. Implementasi pada File Upload

Masih sama seperti pada kode sebelumnya (`handleUpload()`), `http.MaxBytesReader()` diaplikasikan sebelum pembacaan payload (yang pada kode berikut adalah operasi `ParseMultipartForm`).

```go
func handleFileUpload(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

    if err := r.ParseMultipartForm(maxBodyBytes); err != nil {
        var maxBytesErr *http.MaxBytesError
        if errors.As(err, &maxBytesErr) {
            http.Error(w, "file too large (max 1MB)", http.StatusRequestEntityTooLarge)
            return
        }
        http.Error(w, "failed to parse form", http.StatusBadRequest)
        return
    }

    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "file is required", http.StatusBadRequest)
        return
    }
    defer file.Close()

    data, _ := io.ReadAll(file)
    log.Printf("received file: %s, size: %d bytes", header.Filename, len(data))
    w.Write([]byte("received: " + header.Filename))
}
```

Pola yang digunakan di sini sama seperti pada `handleUpload()`. Satu hal yang perlu dicatat: parameter `maxMemory` pada `ParseMultipartForm` bukan limit ukuran payload, melainkan hanya mengontrol seberapa besar data form disimpan di memori sebelum ditulis ke file temporary. Itulah mengapa `MaxBytesReader` tetap dibutuhkan.

`r.FormFile("file")` mengambil file yang dikirim via field bernama `"file"`. Fungsi ini mengembalikan tiga nilai: `multipart.File` (isi file yang bisa dibaca seperti `io.Reader`), `*multipart.FileHeader` (metadata seperti nama file dan ukuran), dan error jika field tidak ada atau bukan file.

`io.ReadAll(file)` membaca seluruh isi file ke dalam slice byte. Di sini hasilnya hanya digunakan untuk mengambil ukuran file via `len(data)` yang kemudian di-log. Pada aplikasi nyata, `data` bisa diteruskan ke proses berikutnya seperti menyimpan ke disk atau object storage.

## B.27.4. Implementasi Lengkap

Buat file `main.go`.

```go
package main

import (
    "encoding/json"
    "errors"
    "io"
    "log"
    "net/http"
)

const maxBodyBytes = 1 << 20 // 1 MB

func handleUpload(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

    var payload map[string]any
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        var maxBytesErr *http.MaxBytesError
        if errors.As(err, &maxBytesErr) {
            http.Error(w, "request body too large", http.StatusRequestEntityTooLarge)
            return
        }
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]any{"received": payload})
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

    if err := r.ParseMultipartForm(maxBodyBytes); err != nil {
        var maxBytesErr *http.MaxBytesError
        if errors.As(err, &maxBytesErr) {
            http.Error(w, "file too large (max 1MB)", http.StatusRequestEntityTooLarge)
            return
        }
        http.Error(w, "failed to parse form", http.StatusBadRequest)
        return
    }

    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "file is required", http.StatusBadRequest)
        return
    }
    defer file.Close()

    data, _ := io.ReadAll(file)
    log.Printf("received file: %s, size: %d bytes", header.Filename, len(data))
    w.Write([]byte("received: " + header.Filename))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("POST /upload/json", handleUpload)
    mux.HandleFunc("POST /upload/file", handleFileUpload)

    log.Println("server started at localhost:9000")
    err := http.ListenAndServe(":9000", mux)
    if err != nil {
        log.Fatal(err)
    }
}
```

## B.27.5. Testing

Jalankan server lalu coba keempat skenario berikut.

```bash
# Payload dengan ukuran normal
curl -X POST http://localhost:9000/upload/json \
  -H "Content-Type: application/json" \
  -d '{"name":"batman"}'
```

Response: `{"received":{"name":"batman"}}` dengan status 200.

```bash
# Payload dengan ukuran melebihi limit (generate 2MB data lalu pipe ke curl)
dd if=/dev/zero bs=1024 count=2048 | curl -X POST http://localhost:9000/upload/json \
  -H "Content-Type: application/json" \
  --data-binary @-
```

Response: `request body too large` dengan status `413 Request Entity Too Large`.

```bash
# Upload file dengan ukuran normal
echo "hello world" > test.txt
curl -X POST http://localhost:9000/upload/file \
  -F "file=@test.txt"
```

Response: `received: test.txt` dengan status 200.

```bash
# Upload file dengan ukuran melebihi limit (generate file 2MB)
dd if=/dev/zero of=bigfile.bin bs=1024 count=2048
curl -X POST http://localhost:9000/upload/file \
  -F "file=@bigfile.bin"
```

Response: `file too large (max 1MB)` dengan status `413 Request Entity Too Large`.

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div><a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.27-request-body-size-limit">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.27...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
