# B.17. Download File

Setelah sebelumnya belajar cara untuk handle upload file, kali ini kita akan belajar bagaimana cara membuat handler yang hasilnya adalah download file.

Sebenarnya download file bisa dengan mudah di-implementasikan menggunakan teknik routing static file, langsung akses url dari public assets di browser. Namun outcome dari teknik ini sangat tergantung pada browser. Tiap browser memiliki behaviour berbeda, ada yang file tidak di-download melainkan dibuka di tab, ada yang ter-download.

Dengan menggunakan teknik berikut, file pasti akan ter-download.

## B.17.1. Struktur Folder Proyek

OK, pertama siapkan terlebih dahulu proyek dengan struktur seperti gambar berikut.

![Project structure](images/B.17_1_structure.png)

File yang berada di folder `files` adalah dummy, jadi anda bisa gunakan file apapun dengan jumlah berapapun untuk keperluan belajar.

## B.17.2. Front End

Kali ini di bagian front end kita tidak menggunakan jQuery, cukup javascript saja tanpa library. 

Pertama siapkan dahulu template nya, isi file `view.html` dengan kode berikut.

```html
<!DOCTYPE html>
<html>
	<head>
		<title>Download file</title>
		<script>
			// javascript code goes here
		</script>
	</head>
	<body>
		<ul id="list-files"></ul>
	</body>
</html>
```

Tag `<ul />` nantinya akan berisikan list semua file yang ada dalam folder `files`. Data list file didapat dari back end. Diperlukan sebuah AJAX untuk pengambilan data tersebut.

Siapkan sebuah fungsi dengan nama `Yo` atau bisa lainnya, fungsi ini berisikan closure `renderData()`, `getAllListFiles()`, dan method `init()`. Buat instance object baru dari `Yo`, lalu akses method `init()`, tempatkan dalam event `window.onload`.

```js
function Yo() {
	var self = this;
	var $ul = document.getElementById("list-files");

	var renderData = function (res) {
		// do stuff
	};

	var getAllListFiles = function () {
		// do stuff
	};

	self.init = function () {
		getAllListFiles();
	};
};

window.onload = function () {
	new Yo().init();
};
```

Closure `renderData()` bertugas untuk melakukan rendering data JSON ke HTML. Berikut adalah isi dari fungsi ini.

```js
var renderData = function (res) {
	res.forEach(function (each) {
		var $li = document.createElement("li");
		var $a = document.createElement("a");

		$li.innerText = "download ";
		$li.appendChild($a);
		$ul.appendChild($li);

		$a.href = "/download?path=" + encodeURI(each.path);
		$a.innerText = each.filename;
		$a.target = "_blank";
	});
};
```

Sedangkan closure `getAllListFiles()`, memiliki tugas untuk request ke back end, mengambil data list semua file. Request dilakukan dalam bentuk AJAX, nilai baliknya adalah data JSON. Setelah data sudah di tangan, fungsi `renderData()` dipanggil.

```js
var getAllListFiles = function () {
	var xhr = new XMLHttpRequest();
	xhr.open("GET", "/list-files");
	xhr.onreadystatechange = function () {
		if (xhr.readyState == 4 && xhr.status == 200) {
			var json = JSON.parse(xhr.responseText);
			renderData(json);
		}
	};
	xhr.send();
};
```

## B.17.3. Back End

Pindah ke bagian back end. Siapkan beberapa hal pada `main.go`, import package, siapkan fungsi main, dan buat beberapa rute.

```go
package main

import "fmt"
import "net/http"
import "html/template"
import "path/filepath"
import "io"
import "encoding/json"
import "os"

type M map[string]interface{}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/list-files", handleListFiles)
	http.HandleFunc("/download", handleDownload)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
```

Buat handler untuk rute `/`.

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```


Lalu siapkan juga route handler `/list-files`. Isi dari handler ini adalah membaca semua file yang ada pada folder `files` untuk kemudian dikembalikan sebagai output berupa JSON. Endpoint ini akan diakses oleh AJAX dari front end.

```go
func handleListFiles(w http.ResponseWriter, r *http.Request) {
	files := []M{}
	basePath, _ := os.Getwd()
	filesLocation := filepath.Join(basePath, "files")
	
	err := filepath.Walk(filesLocation, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, M{"filename": info.Name(), "path": path})
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
```

Fungsi `os.Getwd()` mengembalikan informasi absolute path dimana aplikasi di-eksekusi. Path tersebut kemudian di gabung dengan folder bernama `files` lewat fungsi `filepath.Join`. 

> Fungsi `filepath.Join` akan menggabungkan item-item dengan path separator sesuai dengan sistem operasi dimana program dijalankan. `\` untuk Windows dan `/` untuk Linux/Unix.

Fungsi `filepath.Walk` berguna untuk membaca isi dari sebuah direktori, apa yang ada didalamnya (file maupun folder) akan di-loop. Dengan memanfaatkan callback parameter kedua fungsi ini (yang bertipe `filepath.WalkFunc`), kita bisa mengamil informasi tiap item satu-per satu.

Selanjutnya siapkan handler untuk `/download`. Implementasi teknik download pada dasarnya sama pada semua bahasa pemrograman, yaitu dengan memainkan header **Content-Disposition** pada response.

```go
func handleDownload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
```

**Content-Disposition** adalah salah satu ekstensi MIME protocol, berguna untuk menginformasikan browser bagaimana dia harus berinteraksi dengan output. Ada banyak jenis value content-disposition, salah satunya adalah `attachment`. Pada kode di atas, header `Content-Disposition: attachment; filename=filename.json` menghasilkan output response berupa attachment atau file, yang kemudian akan di-download oleh browser.

Objek file yang direpresentasikan variabel `f`, isinya di-copy ke objek response lewat statement `io.Copy(w, f)`.

## B.17.4. Testing

Jalankan program, akses rute `/`. List semua file dalam folder `files` muncul di sana. Klik salah satu file untuk men-download-nya.

![Download file](images/B.17_2_download.png)

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek pada bab ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-B.17-download-file">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-B.17...</a>
</div>
