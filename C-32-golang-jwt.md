# C.32. JSON Web Token (JWT)

Pada bab ini kita akan belajar JSON Web Token (JWT) dan cara penerapannya. 

## C.31.1. Definisi

JWT adalah salah satu standar JSON (RFC 7519) mengenai akses token. Token tersebut merupakan kombinasi dari beberapa informasi yang di-encode. Informasi yang dimaksud adalah header, payload, dan signature.

![JWT scheme](images/C.32_1_jwt_scheme.png)

 - **Header**, isinya adalah jenis algoritma yang digunakan untuk generate signature.
 - **Payload**, isinya adalah data penting untuk keperluan otentikasi, seperti *grant*, *group*, kapan login terjadi, atau lainnya. Data ini dalam konteks JWT biasa disebut dengan **claim**. 
 - **Signature**, isinya adalah data hasil dari enkripsi data menggunakan algoritma kriptografi. Data yang dimaksud adalah (encoded) header, (encoded) payload, dan secret key.

Umumnya pada aplikasi yang menerapkan teknik otorisasi menggunakan token, data token tersebut di-generate di back end secara acak (menggunakan algoritma tertentu) lalu disimpan di database bersamaan dengan data user. Token tersebut tidak ada isinya, hanya random string (atau mungkin saja tidak). Nantinya ketika ada request, token yang disertakan di request dicocokan dengan token yang ada di database, dilanjutkan dengan pengecekan grant/group/sejenisnya, untuk menentukan apakah request tersebut adalah verified request yang memang berhak mengakses endpoint.

Pada aplikasi yang menerapkan JWT, yang terjadi berbeda. Token adalah hasil dari proses kombinasi, encoding, dan enkripsi terhadap beberapa informasi. Nantinya pada sebuah request, pengecekan token tidak dilakukan dengan membandingkan token yang ada di request vs token yang tersimpan di database, karena token yang disimpan di database. Pengecekan token dilakukan dengan men-decode dan mendekripsi token tersebut, untuk kemudian dicek **claim** nya.

Mungkin sekilas terlihat mengerikan, karena sangat gampang sekali untuk di-retas, buktinya informasi untuk otorisasi sudah bisa diambil lewat token. Ini adalah presepsi yang salah.

Informasi yang ada dalam token, selain di-encode, juga di-enkripsi. Dalam enkripsi diperlukan private key atau secret key, dan hanya pengembang yang tau. Jadi pastikan simpan baik-baik key tersebut.

Selama peretas tidak tau secret key yang digunakan, hasil decoding dan dekripsi data **PASTI TIDAK VALID**.

## C.32.2. Persiapan Praktek

Kita akan buat sebuah aplikasi RESTful web service sederhana, isinya dua buah endpoint `/index` dan `/login`.

 - Pengaksesan `/index` memerlukan token.
 - Token didapat dari proses otentikasi ke endpoint `/login` dengan menyisipkan username dan password.
 - Pada aplikasi yang sudah kita buat, sudah ada data list user yang tersimpan di database (sebenarnya bukan di-database, tapi di file json).

Siapkan folder projek dengan skema seperti berikut.

```bash
$ tree .
.
├── main.go
├── middleware.go
└── users.json

0 directories, 3 files
```

Isi `middleware.go` dengan kode middleware yang sudah biasa kita gunakan.

```go
package main

import "net/http"

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}
```

File `users.json` merupakan database aplikasi, informasi user disimpan dalam file ini. Silakan tambahkan data JSON berikut.

```json
[{
	"username": "noval",
	"password": "kaliparejaya123",
	"email": "terpalmurah@gmail.com",
	"group": "admin"
}, {
	"username": "farhan",
	"password": "masokpakeko",
	"email": "cikrakbaja@gmail.com",
	"group": "publisher"
}]
```

Di file `main.go`, isi dengan kode berikut.

 - Import package yang dibutuhkan. Pada bab ini kita menggunakan library [jwt-go](https://github.com/dgrijalva/jwt-go) untuk handle bagian JWT.

	```go
	package main

	import (
		"context"
		"encoding/json"
		"fmt"
		"io/ioutil"
		"net/http"
		"os"
		"path/filepath"
		"strings"

		jwt "github.com/dgrijalva/jwt-go"
		"github.com/novalagung/gubrak"
	)
	```

 - Siapkan 4 konstanta: nama aplikasi, durasi login, metode enkripsi token, dan secret key.

	```go
	type M map[string]interface{}

	var APPLICATION_NAME = "My Simple JWT App"
	var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")
	```

 - Buat `mux` baru, daftarkan middlrewa `MiddlewareJWTAuthorization` dan dua buah rute. 

	```go
	func main() {
		mux := new(CustomMux)
		mux.RegisterMiddleware(MiddlewareJWTAuthorization)

		mux.HandleFunc("/login", HandlerLogin)
		mux.HandleFunc("/index", HandlerIndex)

		server := new(http.Server)
		server.Handler = mux
		server.Addr = ":8080"

		fmt.Println("Starting server at", server.Addr)
		server.ListenAndServe()
	}
	```

Middleware `MiddlewareJWTAuthorization` tugasnya adalah memvalidasi tiap request yang masuk, dengan mengecek token yang disertakan. Middleware ini hanya berguna pada request ke selain endpoint `/login` (karena pada endpoint ini proses otentikasi terjadi).

## C.32.3. Otentikasi & Generate Token

Siapkan handler `HandlerLogin`.

```go
func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	// ...
}
```

Handler ini bertugas untuk mengotentikasi client/consumer. Username dan Password dikirimkan dalam skema [Basic Auth](https://dasarpemrogramangolang.novalagung.com/B-18-http-basic-auth.html).

Data payload kemudian disisipkan dalam pemanggilan fungsi otentikasi `authenticateUser()`.


```go
ok, userInfo := authenticateUser(username, password)
if !ok {
	http.Error(w, "Invalid username or password", http.StatusBadRequest)
	return
}
```

Fungsi tersebut nantinya akan kita buat. Dua data dikembalikan dalam pemanggilan fungsi:

 - Data `ok`, penanda sukses tidaknya otentikasi.
 - Data `userInfo`, isinya informasi user yang login, yang didapat dari `data.json` (tanpa data password).

Selanjutnya kita akan buat objek claims. Objek ini harus memenuhi persyaratan interface `jwt.Claims`. Objek claims bisa dibuat dari tipe `map` dengan cara membungkusnya dalam fungsi `jwt.MapClaims()`; atau dengan meng-embed interface `jwt.StandardClaims`, dan cara inilah yang akan kita pakai.

Di dalam aturan JWT, ada istilah [Standard Fields](https://en.wikipedia.org/wiki/JSON_Web_Token#Standard_fields). Sebuah claims harus berisikan fields tersebut.

Pada aplikasi yang akan kita buat, claims selain menampung standard fields, juga menampung beberapa informasi lain, oleh karena itu kita perlu buat `struct` baru yang meng-embed `jwt.StandardClaims`.

```
type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Group    string `json:"Group"`
}
```

Ok, struct `MyClaims` sudah siap, sekarang buat objek baru dari struct tersebut.

```go
claims := MyClaims{
	StandardClaims: jwt.StandardClaims{
		Issuer:    APPLICATION_NAME,
		ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
	},
	Username: userInfo["username"].(string),
	Email:    userInfo["email"].(string),
	Group:    userInfo["group"].(string),
}
```

Ada beberapa standard claims, tapi pada contoh yang kita buat hanya dua yang kita isi, `Issuer` dan `ExpiresAt`. Selebihnya, fields seperti username, email, dan group kita isi menggunakan data dari `userInfo`.

Ok, claims sudah siap, sekarang buat token baru. Pembuatannya menggunakan fungsi `jwt.NewWithClaims()`. Parameter pertama adalah metode enkripsi yang digunakan, yaitu `JWT_SIGNING_METHOD`, dan parameter kedua adalah `claims`.

```go
token := jwt.NewWithClaims(
	JWT_SIGNING_METHOD,
	claims,
)
```

Kemudian tanda-tangani token tersebut menggunakan secret key yang sudah didefinisikan di `JWT_SIGNATURE_KEY`, caranya dengan memanggil method `SignedString()` milik token. Pemanggilan method ini mengembalikan token string yang dibutuhkan consumer untuk mengakses aplikasi.

```go
token, err := token.SignedString(JWT_SIGNATURE_KEY)
if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

tokenString, _ := json.Marshal(M{"token": token})
w.Write([]byte(tokenString))
```

Bagian otentikasi dan generate token sebenarnya cukup sampai disini. Tapi sebenarnya ada yang kurang, yaitu fungsi `authenticateUser()`. Silakan siapkan fungsi tersebut.

Isi fungsi ini cukup jelas sesuai namanya, yaitu melakukan pencocokan username dan password dengan data di `users.json`.

```go
func authenticateUser(username, password string) (bool, M) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "users.json")
	buf, _ := ioutil.ReadFile(dbPath)

	data := make([]M, 0)
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return false, nil
	}

	res, _ := gubrak.Find(data, func(each M) bool {
		return each["username"] == username && each["password"] == password
	})

	if res != nil {
		resM := res.(M)
		delete(resM, "password")
		return true, resM
	}

	return false, nil
}
```

## C.32.4. JWT Authorization Middleware

Sekarang kita perlu menyiapkan `MiddlewareJWTAuthorization`, yang tugasnya adalah mengecek setiap request yang masuk ke endpoint selain `/login`, apakah ada akses token yang dibawa atau tidak. Dan jika ada, akan kita verifikasi valid tidaknya token tersebut.

```go
func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// ...
	})
}
```

Kita gunakan skema header `Authorization: Bearer <token>` untuk mengambil token.

Setelah token di tangan, parsing dan lakukan validasi menggunakan fungsi `jwt.Parse()`.

```go
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signing method invalid")
	} else if method != JWT_SIGNING_METHOD {
		return nil, fmt.Errorf("Signing method invalid")
	}

	return JWT_SIGNATURE_KEY, nil
})
if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
```

Parameter ke-2 fungsi `jwt.Parse()` berisikan callback untuk pengecekan valid tidak-nya signing method, jika valid maka secret key dikembalikan. Fungsi `jwt.Parse()` ini sendiri mengembalikan objek token.

Dari objek token, informasi claims diambil, lalu dilakukan pengecekan valid-tidaknya claim tersebut.

```go
claims, ok := token.Claims.(jwt.MapClaims)
if !ok || !token.Valid {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
```

O iya, mungkin ada pertanyaan kenapa objek claims tipenya bukan `MyClaims`. Hal ini karena setelah objek claims dimasukan dalam proses pembentukan token (lewat fungsi `jwt.NewWithClaims()`), objek claim akan selalu di-encode ke tipe `jwt.MapClaims`.

Data claim kemudian disisipkan dalam context. Jadi di endpoint, informasi `userInfo` ini bisa dengan mudah diambil.

```go
ctx := context.WithValue(context.Background(), "userInfo", claims)
r = r.WithContext(ctx)

next.ServeHTTP(w, r)
```

#### Handler Index

Terakhir, kita perlu menyiapkan handler untuk `HandlerIndex`.

```go
func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	message := fmt.Sprintf("hello %s (%s)", userInfo["Username"], userInfo["Group"])
	w.Write([]byte(message))
}
```

Informasi `userInfo` diambil dari context, lalu ditampilkan sebagai response endpoint.

## C.32.5. Testing

Jalankan aplikasi, lalu test menggunakan curl.

 1. Otentikasi:

	```bash
	curl -X POST \
		--header "Content-Type: application/json" \
		--data '{"username":"noval","password":"kaliparejaya123"}' \
		http://localhost:8080/login
	```

	Output:

	![JWT Authentication](images/C.32_2_jwt_authentication.png)

 2. Test endpoint `/index`. Sisipkan token yang dikembalikan pada saat otentikasi, sebagai value header otorisasi dengan skema `Authorization: Bearer <token>`. 

	```bash
	curl -X GET \
		--header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzg2NTI0NzksImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwiVXNlcm5hbWUiOiJub3ZhbCIsIkVtYWlsIjoidGVycGFsbXVyYWhAZ21haWwuY29tIiwiR3JvdXAiOiJhZG1pbiJ9.JREJgUAYs2R5zuquqyX8tk23QlajVVe19susm6JsZq8" \
		http://localhost:8080/index
	```

	Output:

	![JWT Authorization](images/C.32_3_jwt_authorization.png)

Semua berjalan sesuai harapan. Agar lebih meyakinkan, coba lakukan beberapa test dengan skenario yg salah, seperti:

![JWT Error](images/C.32_4_invalid_jwt_token.png)