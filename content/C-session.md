# C.13. Session (Gorilla Session)

Session adalah sebuah konsep penyimpanan data yang shared antar http request. Session umumnya menggunakan cookie untuk menyimpan identifier (kita sebut sebagai **SessionID**). Informasi SessionID tersebut ber-asosiasi dengan data (kita sebut sebagai **SessionData**) yang disimpan di sisi back end dalam media tertentu.

Di back end, SessionData disimpan dalam media database, atau memory, atau fasilitas penyimpanan lainnya. Bisa saja sebenarnya jika SessionData juga disimpan dalam cookie, dengan memanfaatkan secure cookie maka SessionData tersebut ter-enkripsi dan aman dari peretas. Memang aman, tapi jelasnya lebih aman kalau disimpan di sisi server.

Pada chapter ini kita akan mempelajari penerapan session di golang menggunakan beberapa jenis media penyimpanan, yaitu mongo db, postgres sql db, dan secure cookie.

## C.13.1. Manage Session Menggunakan Gorilla Sessions

[Gorilla Sessions](https://github.com/gorilla/sessions) adalah library untuk manajemen session di golang. 

Gorilla menyediakan interface `sessions.Store`, lewat interface ini kita bisa mengakses 3 buah method penting untuk manage session. Store sendiri adalah representasi dari media penyimpanan di back end, bisa berupa database, memory, atau lainnya. Objek store dibuat oleh library lain yang merupakan implementasi dari interface store itu sendiri.

Kembali ke pembahasan mengenai store, 3 buah method yang dimaksud adalah berikut:
    
 - Method `.Get(r *http.Request, name string) (*Session, error)`, mengembalikan objek session. Jika session yang dengan `name` yang dicari tidak ada, maka objek session baru dikembalikan.
 - Method `.New(r *http.Request, name string) (*Session, error)`, mengembalikan objek session baru.
 - Method `.Save(r *http.Request, w http.ResponseWriter, s *Session) error`, digunakan untuk menyimpan session baru.

Dari ketiga method di-atas saya rasa cukup jelas sekilas bagaimana cara mengakses, membuat, dan menyimpan session.

> Kita akan fokus membahas API milik interface `sessions.Store` dahulu, mengenai pembuatan store sendiri ada di pembahasan setelahnya.

Lalu bagaimana dengan operasi hapus/delete? Seperti yang sudah dijelaskan sebelumnya, informasi session dipisah menjadi dua, pertama adalah SessionID yang disimpan di cookie, dan kedua adalah SessionData yang disimpan di back end. Cara untuk menghapus session adalah cukup dengan meng-expired-kan cookie yang menyimpan SessionID.

Cookie merupakan salah satu header pada http request, operasi yang berhubungan dengan cookie pasti membutuhkan objek `http.Request` dan `http.ResponseWriter`. Jika menggunakan echo, kedua objek tersebut bisa diakses lewat objek http context `echo.Context`.

## C.13.2. Membuat Objek Session Baru

Berikut adalah contoh cara membuat session lewat store.

```go
e.GET("/set", func(c echo.Context) error {
    session, _ := store.Get(c.Request(), SESSION_ID)
    session.Values["message1"] = "hello"
    session.Values["message2"] = "world"
    session.Save(c.Request(), c.Response())

    return c.Redirect(http.StatusTemporaryRedirect, "/get")
})
```

Statement `store.Get()` mengembalikan dua objek dengan tipe `session.Session` dan `error`. Pemanggilan method ini memerlukan dua buah parameter untuk disisipkan, yaitu objek http request, dan nama/key SessionID yang disiapkan di konstanta `SESSION_ID`. Method `.Get()` ini akan selalu mengembalikan objek session, ada ataupun tidak ada session yang dicari, objek session tetap dikembalikan.

> Pembuatan objek session baru bisa dilakukan lewat `store.New()` maupun `store.Get()`.

Dari objek session, akses property mutable `.Values` untuk mengambil ataupun mengisi data session. Objek ini bertipe `map[interface{}]interface{}`, berarti SessionData yang akan disimpan juga harus memiliki identifier.

Pada contoh di atas, dua buah data bertipe string disimpan, dengan identifier data yang juga string.

 - SessionData `"hello"` disimpan dengan identifier adalah `message1`.
 - SessionData `"world"` disimpan dengan identifier adalah `message2`.

Cara menyimpan session adalah dengan memanggil method `.Save()` milik objek session, dengan parameter adalah http request dan response.

## C.13.3. Mengakses SessionData

SessionData diakses dari objek session, berikut merupakan contoh caranya.

```go
e.GET("/get", func(c echo.Context) error {
    session, _ := store.Get(c.Request(), SESSION_ID)

    if len(session.Values) == 0 {
        return c.String(http.StatusOK, "empty result")
    }

    return c.String(http.StatusOK, fmt.Sprintf(
        "%s %s",
        session.Values["message1"],
        session.Values["message2"],
    ))
})
```

Seperti yang sudah dibahas di atas, objek `session` kembalian `store.Get()` TIDAK akan pernah berisi `nil`. Ada atau tidak, objek session selalu dikembalikan. 

Dari objek session dilakukan pengecekan ada tidaknya SessionData, caranya dengan cara menghitung isi property `.Values` yang tipenya `map`. Jika isinya kosong maka session belum ada (atau mungkin ada hanya saja expired, atau bisa saja ada tapi invalid).

Pada kode di atas, jika SessionData kosong maka string `empty result` ditampilkan ke layar. Sedangkan jika ada, maka kedua SessionData (message1 dan message2) diambil lalu ditampilkan.

## C.13.4. Menghapus Session

Cara menghapus session adalah dengan meng-expired-kan max age cookie-nya. Property max age bisa diakses lewat `session.Options.MaxAge`.

```go
e.GET("/delete", func(c echo.Context) error {
    session, _ := store.Get(c.Request(), SESSION_ID)
    session.Options.MaxAge = -1
    session.Save(c.Request(), c.Response())

    return c.Redirect(http.StatusTemporaryRedirect, "/get")
})
```

Isi dengan `-1` agar expired, lalu simpan ulang kembali session-nya.

## C.13.5. Session Store dan Context Clear Handler

Session Store adalah representasi dari media tempat di mana data asli session disimpan. Gorilla menyediakan `CookieStore`, penyimpanan data asli pada store ini adalah juga di dalam cookie, namun di-encode dan di-enkripsi menggunakan [Securecookie](https://github.com/gorilla/securecookie).

Selain CookieStore, ada banyak store lain yang bisa kita gunakan. Komunitas begitu baik telah menyediakan berbagai macam store berikut.

 - [github.com/starJammer/gorilla-sessions-arangodb](https://github.com/starJammer/gorilla-sessions-arangodb) - ArangoDB
 - [github.com/yosssi/boltstore](https://github.com/yosssi/boltstore) - Bolt
 - [github.com/srinathgs/couchbasestore](https://github.com/srinathgs/couchbasestore) - Couchbase
 - [github.com/denizeren/dynamostore](https://github.com/denizeren/dynamostore) - Dynamodb on AWS
 - [github.com/savaki/dynastore](https://github.com/savaki/dynastore) - DynamoDB on AWS (Official AWS library)
 - [github.com/bradleypeabody/gorilla-sessions-memcache](https://github.com/bradleypeabody/gorilla-sessions-memcache) - Memcache
 - [github.com/dsoprea/go-appengine-sessioncascade](https://github.com/dsoprea/go-appengine-sessioncascade) - Memcache/Datastore/Context in AppEngine
 - [github.com/kidstuff/mongostore](https://github.com/kidstuff/mongostore) - MongoDB
 - [github.com/srinathgs/mysqlstore](https://github.com/srinathgs/mysqlstore) - MySQL
 - [github.com/EnumApps/clustersqlstore](https://github.com/EnumApps/clustersqlstore) - MySQL Cluster
 - [github.com/antonlindstrom/pgstore](https://github.com/antonlindstrom/pgstore) - PostgreSQL
 - [github.com/boj/redistore](https://github.com/boj/redistore) - Redis
 - [github.com/boj/rethinkstore](https://github.com/boj/rethinkstore) - RethinkDB
 - [github.com/boj/riakstore](https://github.com/boj/riakstore) - Riak
 - [github.com/michaeljs1990/sqlitestore](https://github.com/michaeljs1990/sqlitestore) - SQLite
 - [github.com/wader/gormstore](https://github.com/wader/gormstore) - GORM (MySQL, PostgreSQL, SQLite)
 - [github.com/gernest/qlstore](https://github.com/gernest/qlstore) - ql
 - [github.com/quasoft/memstore](https://github.com/quasoft/memstore) - In-memory implementation for use in unit tests
 - [github.com/lafriks/xormstore](https://github.com/lafriks/xormstore) - XORM (MySQL, PostgreSQL, SQLite, Microsoft SQL Server, TiDB)

Objek store dibuat sekali di awal (atau bisa saja berkali-kali di tiap handler, tergantung kebutuhan). Pada pembuatan objek store, umumya ada beberapa konfigurasi yang perlu disiapkan dan dua buah keys: authentication key dan encryption key.

Dari objek store tersebut, dalam handler, kita bisa mengakses objek session dengan menyisipkan context http request. Silakan lihat kode berikut untuk lebih jelasnya. Store direpresentasikan oleh variabel objek `store`.

```go
package main

import (
    "fmt"
    "github.com/gorilla/context"
    "github.com/gorilla/sessions"
    "github.com/labstack/echo"
    "net/http"
)

const SESSION_ID = "id"

func main() {
    store := newMongoStore()

    e := echo.New()

    e.Use(echo.WrapMiddleware(context.ClearHandler))

    e.GET("/set", func(c echo.Context) error {
        session, _ := store.Get(c.Request(), SESSION_ID)
        session.Values["message1"] = "hello"
        session.Values["message2"] = "world"
        session.Save(c.Request(), c.Response())

        return c.Redirect(http.StatusTemporaryRedirect, "/get")
    })

    // ...
```

Sesuai dengan README Gorilla Session, library ini jika digabung dengan library lain selain gorilla mux, akan berpotensi menyebabkan memory leak. Untuk mengcover isu ini maka middleware `context.ClearHandler` perlu diregistrasikan. Middleware tersebut berada dalam library [Gorilla Context](https://github.com/gorilla/context).

## C.13.6. Mongo DB Store

Kita akan mempelajari pembuatan session store dengan media adalah mongo db. Sebelum kita mulai, ada dua library yang perlu di `go get`.

 - [gopkg.in/mgo.v2](https://gopkg.in/mgo.v2)
 - [github.com/kidstuff/mongostore](https://github.com/kidstuff/mongostore)

Library pertama, `mgo.v2` merupakan driver mongo db untuk golang. Koneksi dari golang ke mongodb akan kita buat lewat API library ini.

Library kedua, merupakan implementasi dari interface `sessions.Store` untuk mongo db.

Silakan kombinasikan semua koding yang sudah kita tulis di atas agar menjadi satu aplikasi. Lalu buat fungsi `newMongoStore()`.

```go
import (
    "fmt"
    "github.com/gorilla/context"
    "github.com/kidstuff/mongostore"
    "github.com/labstack/echo"
    "gopkg.in/mgo.v2"
    "log"
    "net/http"
    "os"
)

// ...

func newMongoStore() *mongostore.MongoStore {
    mgoSession, err := mgo.Dial("localhost:27123")
    if err != nil {
        log.Println("ERROR", err)
        os.Exit(0)
    }

    dbCollection := mgoSession.DB("learnwebgolang").C("session")
    maxAge := 86400 * 7
    ensureTTL := true
    authKey := []byte("my-auth-key-very-secret")
    encryptionKey := []byte("my-encryption-key-very-secret123")

    store := mongostore.NewMongoStore(
        dbCollection,
        maxAge,
        ensureTTL,
        authKey,
        encryptionKey,
    )
    return store
}
```

Statement `mgo.Dial()` digunakan untuk terhubung dengan mongo db server. Method dial mengembalikan dua objek, salah satunya adalah mgo session.

> Pada saat pembuatan buku ini, penulis menggunakan mongo db server yang up pada port `27123`, silakan menyesuaikan connection string dengan credentials mongo db yang digunakan.

Dari mgo session akses database lewat method `.DB()`, lalu akses collection yang ingin digunakan sebagai media penyimpanan data asli session lewat method `.C()`.

Statement `mongostore.NewMongoStore()` digunakan untuk membuat mongo db store. Ada beberapa parameter yang diperlukan: objek collection mongo di atas, dan dua lagi lainnya adalah authentication key dan encryption key.

Jika pembaca merasa bingung, silakan langsung buka [source code untuk chapter ini di Github](https://github.com/novalagung/dasarpemrogramangolang-example/), mungkin membantu.

## C.13.7. Postgres SQL Store

Pembuatan postgres store caranya kurang lebih sama dengan mongo store. Library yang dipakai adalah [github.com/antonlindstrom/pgstore](https://github.com/antonlindstrom/pgstore).

Gunakan `pgstore.NewPGStore()` untuk membuat store. Isi parameter pertama dengan connection string postgres server, lalu authentication key dan encryption key.

```go
import (
    "fmt"
    "github.com/antonlindstrom/pgstore"
    "github.com/gorilla/context"
    "github.com/labstack/echo"
    "log"
    "net/http"
    "os"
)

// ...

func newPostgresStore() *pgstore.PGStore {
    url := "postgres://novalagung:@127.0.0.1:5432/novalagung?sslmode=disable"
    authKey := []byte("my-auth-key-very-secret")
    encryptionKey := []byte("my-encryption-key-very-secret123")

    store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
    if err != nil {
        log.Println("ERROR", err)
        os.Exit(0)
    }

    return store
}
```

## C.13.8. Secure Cookie Store

Penggunaan cookie store kurang penulis anjurkan, meski sebenarnya cukup aman. Implementasi store jenis ini adalah yang paling mudah, karena tidak butuh database server atau media lainnya; dan juga karena API untuk cookie store sudah tersedia dalam gorilla sessions secara default.

```go
import (
    "fmt"
    "github.com/gorilla/context"
    "github.com/gorilla/sessions"
    "github.com/labstack/echo"
    "net/http"
)

// ...

func newCookieStore() *sessions.CookieStore {
    authKey := []byte("my-auth-key-very-secret")
    encryptionKey := []byte("my-encryption-key-very-secret123")

    store := sessions.NewCookieStore(authKey, encryptionKey)
    store.Options.Path = "/"
    store.Options.MaxAge = 86400 * 7
    store.Options.HttpOnly = true

    return store
}
```

Tentukan path dan default max age cookie lewat `store.Options`.

## C.13.9. Test Aplikasi

Silakan gabung semua kode yang sudah kita pelajari (kecuali bagian store), lalu pilih salah satu implementasi store di atas. Jalankan aplikasi untuk testing.

Tujuan dari kode yang kita tulis kurang lebih sebagai berikut.

 1. Ketika `/get` diakses untuk pertama kali, `empty result` muncul, tidak ada data session yang disimpan sebelumnya.
 2. Rute `/set` diakses, lalu sebuah session disimpan, dari rute ini pengguna di-redirect ke `/get`, sebuah pesan muncul yang sumber datanya tak lain adalah dari session.
 3. Rute `/delete` diakses, session dihapus, lalu di-redirect lagi ke `/get`, pesan `empty result` muncul kembali karena session sudah tidak ada (dihapus).

![Session Test](images/C_session_1_test.png)

---

 - [Echo](https://github.com/labstack/echo), by Vishal Rana (Lab Stack), MIT license
 - [Gorilla Sessions](https://github.com/gorilla/sessions), by Gorilla web toolkit team, BSD-3-Clause license
 - [Gorilla Context](https://github.com/gorilla/context), by Gorilla web toolkit team, BSD-3-Clause license
 - [Gorilla Securecookie](https://github.com/gorilla/securecookie), by Gorilla web toolkit team, BSD-3-Clause license
 - [PG Store](https://github.com/antonlindstrom/pgstore), by Anton Lindström, MIT License
 - [Mongo Store](https://github.com/kidstuff/mongostore), by Nguyễn Văn Cao Nguyên, BSD-3-Clause License
 - [Mgo v2, Golang Mongo Driver](https://labix.org/mgo), by Gustavo Niemeyer, Simplified BSD License

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktik chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.13-session">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.13...</a>
</div>

---

<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>
