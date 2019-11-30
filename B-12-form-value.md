# B.12. Form Value

Pada bab ini kita akan belajar bagaimana cara untuk submit data, dari form di layer front end, ke back end. 

# B.12.1. Front End

Pertama siapkan folder projek baru, dan sebuah file template view `view.html`. Pada file ini perlu didefinisikan 2 buah template, yaitu `form` dan `result`. Template pertama (`form`) dijadikan landing page program, isinya beberapa inputan untuk submit data.

```html
{{define "form"}}
<!DOCTYPE html>
<html>
	<head>
		<title>Input Message</title>
	</head>
	<body>
		<form method="post" action="/process">
			<label>Name :</label>
			<input type="text" placeholder="Type name here" name="name" required />
			<br />
			
			<label>Message :</label>
			<input type="text" placeholder="Type message here" name="message" required />
			<br />
			
			<button type="submmit">Print</button>
		</form>
	</body>
</html>
{{end}}
```

Aksi dari form di atas adalah `/process`, yang dimana url tersebut nantinya akan mengembalikan output berupa html hasil render template `result`. Silakan tulis template result berikut dalam `view.html` (jadi file view ini berisi 2 buah template).

```html
{{define "result"}}
<!DOCTYPE html>
<html>
	<head>
		<title>Show Message</title>
	</head>
	<body>
		<h1>Hello {{.name}}</h1>
		<p>{{.message}}</p>
	</body>
</html>
{{end}}
```

## B.12.2. Back End

Buat file `main.go`. Dalam file ini 2 buah route handler diregistrasikan.

 - Route `/` adalah landing page, menampilkan form input.
 - Route `/process` sebagai action dari form input, menampilkan text.

```go
package main

import "net/http"
import "fmt"
import "html/template"

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
```

Handler route `/` dibungkus dalam fungsi bernama `routeIndexGet`. Di dalamnya, template `form` dalam file template `view.html` akan di-render ke view. Request dalam handler ini hanya dibatasi untuk method GET saja, request dengan method lain akan menghasilkan response 400 Bad Request.

```go
func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
```

Fungsi `routeSubmitPost` yang merupakan handler route `/process`, berisikan proses yang mirip seperti handler route `/`, yaitu parsing `view.html` untuk di ambil template `result`-nya. Selain itu, pada handler ini ada proses pengambilan data yang dikirim dari form ketika di-submit, untuk kemudian disisipkan ke template view.

```go
func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = r.FormValue("name")
		var message = r.Form.Get("message")

		var data = map[string]string{"name": name, "message": message}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
```

Ketika user submit ke `/process`, maka data-data yang ada di form input dikirim. Method `ParseForm()` pada statement `r.ParseForm()` berguna untuk parsing form data yang dikirim dari view, sebelum akhirnya bisa diambil data-datanya. Method tersebut mengembalikan data `error` jika proses parsing gagal (kemungkinan karena data yang dikirim ada yang tidak valid).

Pengambilan data yang dikirim dilakukan lewat method `FormValue()`. Contohnya seperti pada kode di atas, `r.FormValue("name")`, akan mengembalikan data inputan `name` (data dari inputan `<input name="name" />`).

Selain lewat method `FormValue()`, pengaksesan data juga bisa dilakukan dengan cara mengakses property `Form` terlebih dahulu, kemudian mengakses method `Get()`. Contohnya seperti `r.Form.Get("message")`, yang akan menghasilkan data inputan `message`. Hasil dari kedua cara di atas adalah sama.

Setelah data dari form sudah ditangkap oleh back-end, data ditampung dalam variabel `data` yang bertipe `map[string]string`. Variabel `data` tersebut kemudian disisipkan ke view, lewat statement `tmpl.Execute(w, data)`.

## B.12.3. Test

OK, sekarang coba jalankan program yang telah kita buat, dan cek hasilnya.

![Form Value](images/B.12_1_form.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.12-form-value">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.12...</a>
</div>
