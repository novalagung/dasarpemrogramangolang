# A.49. Exec

**Exec** digunakan untuk eksekusi perintah command line lewat kode program. Command yang bisa dieksekusi adalah semua command yang bisa dieksekusi di command line sesuai sistem operasinya (Linux-distros, Windows, MacOS, dan lainnya).

## A.49.1. Penggunaan Exec

Go menyediakan package `exec` isinya banyak sekali API atau fungsi untuk keperluan eksekusi perintah command line.

Cara eksekusi command adalah menggunakan fungsi `exec.Command()` dengan argument pemanggilan fungsi diisi command CLI yang diinginkan. Contoh:

```go
package main

import "fmt"
import "os/exec"

func main() {
    var output1, _ = exec.Command("ls").Output()
    fmt.Printf(" -> ls\n%s\n", string(output1))

    var output2, _ = exec.Command("pwd").Output()
    fmt.Printf(" -> pwd\n%s\n", string(output2))

    var output3, _ = exec.Command("git", "config", "user.name").Output()
    fmt.Printf(" -> git config user.name\n%s\n", string(output3))
}
```

Fungsi `exec.Command()` menjalankan command yang dituliskan pada argument pemanggilan fungsi.

Untuk mendapatkan outputnya, chain saja langsung dengan method `Output()`. Output yang dihasilkan berbentuk `[]byte`, maka pastikan cast ke string terlebih dahulu untuk membaca isi outputnya.

![Ekeskusi command menggunakan exec](images/A_exec_1_exec.png)

## A.49.2. Rekomendasi Penggunaan Exec

Ada kalanya saat eksekusi command yang sudah jelas-jelas ada (seperti `ls`, `dir`, atau lainnya), error muncul menginformasikan bahwa command tidak ditemukan (command not found). Hal ini biasanya terjadi karena executable dari command-command tersebut tidak ada. Seperti di windows tidak ada `cmd` atau `cmd.exe`, di Linux tidak ditentukan apakah memakai `bash` atau `shell`, dan lainnya

Untuk mengatasi masalah ini, tambahkan `bash -c` pada sistem operasi berbasi Linux, MacOS, Unix, atau `cmd /C` untuk OS Windows.

```go
if runtime.GOOS == "windows" {
    output, err = exec.Command("cmd", "/C", "git config user.name").Output()
} else {
    output, err = exec.Command("bash", "-c", "git config user.name").Output()
}
```

Statement `runtime.GOOS` penggunaannya mengembalikan informasi sistem operasi dalam bentuk string. Manfaatkan seleksi kondisi untuk memastikan command yang ingin dieksekusi sudah sesuai dengan OS atau belum.

## A.49.3. Method Exec Lainnya

Selain `.Output()` ada sangat banyak sekali API untuk keperluan komunikasi dengan OS/CLI yang bisa dipergunakan. Lebih detailnya silakan langsung melihat dokumentasi package tersebut di [https://golang.org/pkg/os/exec/](https://golang.org/pkg/os/exec/)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-A.49-exec">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-A.49...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
