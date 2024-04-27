# B.9. Template: Render Specific HTML Template

Pada chapter ini kita akan mempelajari cara render template html tertentu untuk dijadikan output pengaksesan endpoint. Sebuah file view bisa berisikan banyak template. Template mana yang ingin di-render bisa ditentukan.

## B.9.1. Front End

Siapkan folder project baru, buat file template bernama `view.html`, lalu isi dengan kode berikut.

```html
{{define "index"}}
<html>
	<head>
		<title>Learning html/template Functions</title>
	</head>
	<body>
		<h2>Index</h2>
	</body>
</html>
{{end}}

{{define "test"}}
<html>
	<head>
		<title>Other Template</title>
	</head>
	<body>
		<h2>Test</h2>
	</body>
</html>
{{end}}
```

Pada file view di atas, terlihat terdapat 2 template didefinisikan dalam 1 file, template `index` dan `test`.

- Template `index` ditampilkan ketika rute `/` diakses 
- Template `test` ditampilkan ketika rute `/test` diakses

## B.9.2. Back End

Selanjutnya siapkan kode di sisi back-end, buat file `main.go`, tulis kode berikut.

```go
package main

import "net/http"
import "fmt"
import "html/template"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("index").ParseFiles("view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("test").ParseFiles("view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
```

Pada kode di atas bisa dilihat, terdapat 2 rute yang masing-masing mem-parsing file yang sama, tapi spesifik template yang dipilih untuk di-render berbeda.

Contoh di rute `/`, sebuah template dialokasikan dengan nama `index`, kemudian di-parsing-lah view bernama `view.html` menggunakan method `ParseFiles()`. Go secara cerdas melakukan pencarian dalam file view tersebut, apakah ada template yang namanya adalah `index` atau tidak. Jika ada maka ditampilkan. Hal ini juga berlaku pada rute `/test`, jika isi dari template bernama `test` ada maka ditampilkan.

## B.9.3. Testing

Lakukan tes pada program yang telah kita buat, kurang lebih hasilnya seperti pada gambar berikut.

![Rute `/` dan `/test`](images/B_render_specific_html_template_1_preview.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.9-render-specific-html-template">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.9...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
