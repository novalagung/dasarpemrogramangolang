# C.7. Template Rendering in Echo

Pada chapter ini kita akan belajar cara render template html pada aplikasi yang routingnya menggunakan echo.

Pada dasarnya proses parsing dan rendering template tidak di-handle oleh echo sendiri, melainkan oleh API dari package `html/template`. Jadi bisa dibilang cara render template di echo adalah sama seperti pada aplikasi yang murni menggunakan golang biasa, seperti yang sudah dibahas pada chapter [Template: Render HTML Template](/B-template-render-html.html), [Template: Render Partial HTML Template](/B-template-render-partial-html.html), [Template: Render Specific HTML Template](/B-render-specific-html.html), dan [Template: Render HTML String](/B-render-html-string.html).

Echo menyediakan satu fasilitas yang bisa kita manfaatkan untuk standarisasi rendering template. Cara penggunaannya, dengan meng-override default `.Renderer` property milik echo menggunakan objek cetakan struct, yang di mana pada struct tersebut harus ada method bernama `.Render()` dengan skema sesuai dengan kebutuhan echo. Nah, di dalam method `.Render()` inilah kode untuk parsing dan rendering template ditulis.

## C.7.1. Praktek

Agar lebih mudah dipahami, mari langsung kita praktekan. Siapkan sebuah project, import package yang dibutuhkan.

```go
package main

import (
    "github.com/labstack/echo"
    "html/template"
    "io"
    "net/http"
)

type M map[string]interface{}
```

Buat sebuah struct bernama `Renderer`, struct ini mempunyai 3 buah property dan 2 buah method.

```go
type Renderer struct {
    template *template.Template
    debug    bool
    location string
}
```

Berikut adalah tugas dan penjelasan mengenai ketiga property di atas.

 - Property `.template` bertanggung jawab untuk parsing dan rendering template. 
 - Property `.location` mengarah ke path folder di mana file template berada.
 - Property `.debug` menampung nilai bertipe `bool`.
    - Jika `false`, maka parsing template hanya dilakukan sekali saja pada saat aplikasi di start. Mode ini sangat cocok untuk diaktifkan pada stage production.
    - Sedangkan jika nilai adalah `true`, maka parsing template dilakukan tiap pengaksesan rute. Mode ini cocok diaktifkan untuk stage development, karena perubahan kode pada file html sering pada stage ini.

Selanjutnya buat fungsi `NewRenderer()` untuk mempermudah inisialisasi objek renderer.

```go
func NewRenderer(location string, debug bool) *Renderer {
    tpl := new(Renderer)
    tpl.location = location
    tpl.debug = debug

    tpl.ReloadTemplates()

    return tpl
}
```

Siapkan dua buah method untuk struct renderer, yaitu `.ReloadTemplates()` dan `.Render()`.

```go
func (t *Renderer) ReloadTemplates() {
    t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
    w io.Writer, 
    name string, 
    data interface{}, 
    c echo.Context,
) error {
    if t.debug {
        t.ReloadTemplates()
    }

    return t.template.ExecuteTemplate(w, name, data)
}
```

Method `.ReloadTemplates()` bertugas untuk parsing template. Method ini wajib dipanggil pada saat inisialisasi objek renderer. Jika `.debug == true`, maka method ini harus dipanggil setiap kali rute diakses (jika tidak, maka perubahan pada view tidak akan muncul).

Method `.Render()` berguna untuk render template yang sudah diparsing sebagai output. Method ini harus dibuat dalam skema berikut.

```go
// skema method Render()
func (io.Writer, string, interface{}, echo.Context) error
```

Selanjutnya, buat echo router, override property renderer nya, dan siapkan sebuah rute.

```go
func main() {
    e := echo.New()

    e.Renderer = NewRenderer("./*.html", true)

    e.GET("/index", func(c echo.Context) error {
        data := M{"message": "Hello World!"}
        return c.Render(http.StatusOK, "index.html", data)
    })

    e.Logger.Fatal(e.Start(":9000"))
}
```

Saat pemanggilan `NewRenderer()` sisipkan path folder tempat file template html berada. Gunakan `./*.html` agar mengarah ke **semua file html pada current folder**.

Buat file `index.html` dengan isi kode di bawah ini.

```html
<!DOCTYPE html>
<html>
    <head>
        <title></title>
    </head>
    <body>
        Message from index: {{.message}}!
    </body>
</html>
```

Pada rute `/index`, sebuah variabel bernama `data` disiapkan, bertipe `map` dengan isi satu buah item. Data tersebut disisipkan pada saat view di-render, membuatnya bisa diakses dari dalam template html.

Syntax `{{.message}}` artinya menampilkan isi property yang namanya adalah `message` dari current context (yaitu objek data yang disisipkan). Lebih jelasnya silakan baca kembali chapter [B. Template Actions & Variables](/B-template-actions-variables.html).

Jalankan aplikasi untuk melihat hasilnya.

![Preview](images/C_echo_template_rendering_1_preview.png)

## C.7.2. Render Parsial dan Spesifik Template

Proses parsing dan rendering tidak di-handle oleh echo, melainkan menggunakan API dari `html/template`. Echo hanya menyediakan tempat untuk mempermudah pemanggilan fungsi rendernya. Nah dari sini berarti untuk render parsial, render spesifik template, maupun operasi template lainnya dilakukan seperti biasa, menggunakan `html/template`.

---

 - [Echo](https://github.com/labstack/echo), by Vishal Rana (Lab Stack), MIT license

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.7-echo-template-rendering">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.7...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="430px" frameborder="0" scrolling="no"></iframe>
