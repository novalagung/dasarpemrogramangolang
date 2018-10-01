# C.16. Send Mail (`net/smtp`, Gomail v2)

Pada bab ini kita akan belajar cara mengirim email dari aplikasi golang, menggunakan dua cara berikut.

 1. Dengan memanfaatkan package `net/smtp`.
 2. Menggunakan Library [gomail](https://gopkg.in/gomail.v2).

## C.16.1. Kirim Email Menggunakan `net/smtp`

Golang menyediakan package `net/smtp`, isinya banyak API untuk berkomunikasi via protokol SMTP. Lewat package ini kita bisa melakukan operasi kirim email.

Sebuah akun email diperlukan dalam mengirim email, silakan gunakan provider email apa saja. Pada bab ini kita gunakan Google Mail (gmail), jadi siapkan satu buah akun gmail untuk keperluan testing.

Mari kita praktekan. Buat folder projek baru, salin kode berikut.

```go
package main

import (
    "fmt"
    "log"
    "net/smtp"
    "strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "emailanda@gmail.com"
const CONFIG_PASSWORD = "passwordemailanda"

func main() {
    to := []string{"recipient1@gmail.com", "emaillain@gmail.com"}
    cc := []string{"tralalala@gmail.com"}
    subject := "Test mail"
    message := "Hello"

    err := sendMail(to, cc, subject, message)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println("Mail sent!")
}
```

Dalam implementasinya, untuk bisa mengirim email, dibutuhkan mail server. Karena kita menggunakan email google, maka mail server milik google digunakan.

Pada kode di atas, konstanta dengan prefix `CONFIG_` adalah konfigurasi yang diperlukan untuk terhubung dengan mail server. Sesuaikan `CONFIG_EMAIL` dan `CONFIG_PASSWORD` dengan akun yang digunakan.

Di dalam fungsi main bisa dilihat, fungsi `sendMail()` dipanggil untuk mengirim email, dengan empat buah parameter disisipkan.

 - Parameter `to`, adalah tujuan email.
 - Parameter `cc`, adalah cc tujuan.
 - Parameter `subject`, adalah subjek email.
 - Parameter `message`, adalah body email.

Email pada konstanta `CONFIG_EMAIL` menjadi sender email.

OK, selanjutnya buat fungsi `sendMail()` berikut.

```go
func sendMail(to []string, cc []string, subject, message string) error {
    body := "From: " + CONFIG_EMAIL + "\n" +
        "To: " + strings.Join(to, ",") + "\n" +
        "Cc: " + strings.Join(cc, ",") + "\n" +
        "Subject: " + subject + "\n\n" +
        message

    auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
    smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

    err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, append(to, cc...), []byte(body))
    if err != nil {
        return err
    }

    return nil
}
```

Fungsi `sendMail()` digunakan untuk mengirim email. Empat data yang disisipkan pada fungsi tersebut dijadikan satu dalam format tertentu, lalu disimpan ke variabel `body`. 

Statement yang ditampung oleh `body` akan menghasilkan string berikut (formatnya adalah baku).

```
From: recipient1@gmail.com, emaillain@gmail.com
Cc: tralalala@gmail.com
Subject: Test mail

Hello
```

Pengiriman email dilakukan lewat `smtp.SendMail()`. Dalam pemanggilannya 5 buah parameter disisipkan, berikut adalah penjelasan masing-masing parameter.

 - Parameter ke-1, `smtpAddr`, merupakan kombinasi host dan port mail server.
 - Parameter ke-2, `auth`, menampung credentials untuk keperluan otentikasi ke mail server. Objek ini dicetak lewat `smtp.PlainAuth()`.
 - Parameter ke-3, `CONFIG_EMAIL`, adalah alamat email yang digunakan untuk mengirim email.
 - Parameter ke-4, Isinya adalah semua email tujuan, termasuk `Cc`.
 - Parameter ke-5, isinya `body` email.

Jalankan aplikasi. Lihat di console, error muncul.

![Error gmail SMTP](images/C.16_1_gmail_error.png)

Error di atas hanya muncul pada pengiriman email menggunakan akun google mail. Untuk alasan keamanan, google men-disable akun gmail untuk digunakan mengirim email lewat kode program.

Aktifkan fasilitas **less secure apps** untuk meng-enable-nya. Login ke gmail masing-masing, kemudian buka link https://myaccount.google.com/lesssecureapps, lalu klik tombol toggle agar menjadi **OFF**.

![Google Less Secure Apps](images/C.16_2_less_secure.png)

Jalankan ulang aplikasi, email terkirim. Lihat di inbox email tujuan pengiriman untuk mengecek hasilnya.

![Email received](images/C.16_3_email_received.png)

## C.16.2. Kirim Email Menggunakan Gomail v2

Dengan library [gomail](https://gopkg.in/gomail.v2), pengiriman email bisa dilakukan dengan mudah. Beberapa operasi seperti membuat email dalam bentuk html, menambahkan attachment, menambahkan bcc, bisa dilakukan dengan mudah lewat gomail.

Mari langsung kita praktekan. Unduh terlebih dahulu library-nya.

```bash
$ go get -u gopkg.in/gomail.v2
```

Lalu tulis kode berikut.

```go
package main

import (
    "gopkg.in/gomail.v2"
    "log"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "emailanda@gmail.com"
const CONFIG_PASSWORD = "passwordemailanda"

func main() {
    mailer := gomail.NewMessage()
    mailer.SetHeader("From", CONFIG_EMAIL)
    mailer.SetHeader("To", "recipient1@gmail.com", "emaillain@gmail.com")
    mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
    mailer.SetHeader("Subject", "Test mail")
    mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
    mailer.Attach("./sample.png")

    dialer := gomail.NewDialer(
        CONFIG_SMTP_HOST,
        CONFIG_SMTP_PORT,
        CONFIG_EMAIL,
        CONFIG_PASSWORD,
    )

    err := dialer.DialAndSend(mailer)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println("Mail sent!")
}
```

Siapkan satu buah image bernama `sample.png`, simpan di dalam folder yang sama dengan file main. Untuk meng-attach file ke dalam email, gunakan method `.Attach()` milik `*gomail.Message`.

Pada contoh kali ini email isinya adalah HTML. Gunakan MIME html pada parameter pertama `.SetBody()` untuk mengaktifkan mode html email.

Jalankan aplikasi, lalu cek hasilnya email yang dikirim di inbox.

![Mail with attachment](images/C.16_4_mail_with_attachment.png)

---

 - [Gomail v2](https://gopkg.in/gomail.v2), by Alexandre Cesaro, MIT license
