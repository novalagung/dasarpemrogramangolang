# B.10. Template: Render HTML String

Output HTML yang muncul, selain bersumber dari template view bisa juga bersumber dari sebuah string. Dengan menggunakan method `Parse()` milik `*template.Template` kita bisa menjadikan HTML string sebagai output di web.

## B.10.1. Praktek

Langsung saja kita praktekkan, siapkan folder project baru beserta file `main.go`, isi dengan kode berikut. 

```go
package main

import "net/http"
import "fmt"
import "html/template"

const view string = `<html>
	<head>
		<title>Template</title>
	</head>
	<body>
		<h1>Hello</h1>
	</body>
</html>`
```

Konstanta bernama `view` dengan tipe `string` disiapkan, isinya HTML string yang nanbtinya kita jadikan sebagai output pengaksesan endpoint.

Kemudian buat fungsi `main()`, isinya adalah route handler `/index`. Dalam handler tersebut, string html `view` diparsing lalu dirender sebagai output.

Tambahkan juga rute `/` yang isinya adalah me-redirect request secara paksa ke `/index` (via fungsi `http.Redirect()`).

```go
func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
```

Pada kode di atas bisa dilihat sebuah template bernama `main-template` disiapkan. Template tersebut diisi dengan hasil parsing string html `view` lewat method `Parse()`.

## B.10.2. Testing

Lakukan tes dan lihat hasilnya.

![String html sebagai output](images/B_render_html_string_1_parse.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.10-render-html-string">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.10...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
