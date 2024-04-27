# B.23. Server Handler HTTP Request Cancellation

Dalam konteks web application, kadang kala sebuah HTTP request butuh waktu cukup lama untuk selesai, bisa jadi karena kode yang kurang dioptimasi atau prosesnya memang lama, atau mungkin ada faktor lainnya. Dari sisi client, biasanya ada handler untuk cancel request ketika request melebihi batas timeout yang sudah ditentukan.

Berbeda dengan handler di back end-nya, by default request yang sudah di-cancel oleh client tidak mempengaruhi yang terjadi di back-end, proses di back end akan tetap lanjut hingga selesai. Umumnya hal ini bukan merupakan masalah, tapi untuk beberapa *case* ada baiknyakita perlu men-*treat* *cancelled request* dengan baik. Dan pada chapter ini kita akan belajar caranya.

> Chapter ini fokus terhadap cancellation pada client http request di sisi back-end. Untuk topik cancellation pada proses konkuren silakan pembahasannya ada di chapter [A.64. Concurrency Pattern: Context Cancellation Pipeline](/A-pipeline-context-cancellation.html).

## B.32.1. Praktek

Dari objek `*http.Request` informasi objek context bisa diakses lewat method `.Context()`, dan dari context tersebut kita bisa mendeteksi apakah sebuah request di-cancel atau tidak oleh client.

> Pada chapter ini kita tidak membahas secara rinci apa itu context karena sudah ada pembahasan terpisah mengenai topik tersebut di chapter [A.64. Concurrency Pattern: Context Cancellation Pipeline](/A-pipeline-context-cancellation.html).

Object context memiliki method `.Done()` yang nilai baliknya berupa channel. Dari channel tersebut kita bisa deteksi apakah request di-cancel atau tidak oleh client, jika ada data yang diterima via channel tersebut dan error yang didapat ada keterangan `"cancelled"` maka bisa diasumsikan request tersebut dibatalkan oleh client.

Mari kita praktekan langsung. Silakan mulai dengan menulis kode berikut.

```go
package main

import (
	"log"
	"net/http"
	"strings"
	"time"
	"log"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// do something here
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
```

Di dalam `handleIndex()` disimulasikan sebuah proses membutuhkan waktu lama untuk selesai (kita gunakan `time.Sleep()` untuk ini). Umumnya kode dituliskan langsung dalam handler tersebut, tapi pada kasus ini tidak. Untuk bisa mendeteksi sebuah request di-cancel atau tidak kita akan manfaatkan goroutine baru.

Dalam penerapannya ada dua pilihan opsi:

- Cara ke-1: Dengan menaruh proses utama di dalam gorutine tersebut, dan menaruh kode untuk deteksi di luar (di dalam handler-nya).
- Cara ke-2: Atau sebaliknya. Menaruh proses utama di dalam handler, dan menempatkan deteksi cancelled request dalam goroutine baru.

Pada contoh berikut, kita gunakan cara pertama. Tulis kode berikut dalam handler.

```go
done := make(chan bool)
go func() {
    // do the process here
    // simulate a long-time request by putting 10 seconds sleep
    time.Sleep(10 * time.Second)

    done <- true
}()

select {
case <-r.Context().Done():
    if err := r.Context().Err(); err != nil {
        if strings.Contains(strings.ToLower(err.Error()), "canceled") {
            log.Println("request canceled")
        } else {
            log.Println("unknown error occured.", err.Error())
        }
    }
case <-done:
    log.Println("done")
}
```

Pada kode di atas terlihat, proses utama dibungkus dalam goroutine. Ketika selesai, maka back-end akan menerima data via channel `done`.

Keyword `select` di situ disiapkan untuk pendeteksian dua kondisi berikut:

- Channel `r.Context().Done()`. Jika channel ini menerima data maka diasumsikan request selesai. Selanjutnya lakukan pengecekan pada objek error milik context untuk deteksi apakah selesai-nya request ini karena memang selesai, atau di-cancel oleh client, atau faktor lainnya.
- Channel `<-done`. Jika channel ini menerima data, maka proses utama adalah selesai.

Sekarang coba jalankan kode lalu test hasilnya.

```bash
curl -X GET http://localhost:8080/
```

![Cancelled client http request](images/B_server_handler_http_request_cancellation_1_cancelled_request_get.png)

Pada gambar di atas terdapat dua request, yg pertama sukses dan yang kedua adalah cancelled. Pesan `request cancelled` muncul ketika client http request dibatalkan.

> Di CMD/terminal bisa cukup dengan `ctrl + c` untuk cancel request

## B.32.2. Handle Cancelled Request yang ada Payload-nya

Khusus untuk request dengan HTTP method yang memiliki request body (payload), maka channel `r.Context().Done()` tidak akan menerima data hingga terjadi proses read pada body payload.

Silakan coba saja, misalnya dengan menambahkan kode berikut.

```go
go func() {
    // do the process here
    // simulate a long-time request by putting 10 seconds sleep
    
    body, err := io.ReadAll(r.Body)
    // ...

    time.Sleep(10 * time.Second)

    done <- true
}()
```

Hasilnya:

```go
curl -X POST http://localhost:8080/ -H 'Content-Type: application/json' -d '{}'
```

![Cancelled client http request](images/B_server_handler_http_request_cancellation_2_cancelled request_with_payload.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.23-server-handler-http-request-cancellation">https://github.com/novalagung/dasarpemrogramangolang-example/...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
