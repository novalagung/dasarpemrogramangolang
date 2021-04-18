# C.35. Redis

Pada bab ini kita akan belajar cara menggunakan Redis, dari cara koneksi Redis, cara menyimpan data dan cara mengambil data. Untuk command lain lebih lengkap bisa di lihat di [Command Redis](https://redis.io/commands).

## C.35.1 Apa itu Redis?
Redis, singkatan dari Remote Dictionary Server, adalah penyimpanan data nilai utama di dalam memori yang super cepat dengan sumber terbuka untuk digunakan sebagai database, cache, broker pesan, dan antrean. Proyek ini dimulai ketika Salvatore Sanfilippo, pengembang awal Redis, mencoba meningkatkan skalabilitas startup Italia miliknya. Redis kini memberikan respons dalam waktu di bawah satu milidetik yang memungkinkan jutaan permintaan per detik untuk aplikasi real-time pada Permainan, Ad-Tech, Layanan Finansial, Layanan Kesehatan, dan IoT. Redis adalah pilihan populer untuk caching, manajemen sesi, permainan, papan peringkat, analisis real-time, geospasial, tumpangan berkendara, obrolan/perpesanan, streaming media, dan aplikasi pub/sub.
> ~amazon/redis

Library Golang Redis yang populer ketika artikel ini dibuat adalah
- GoModule [gomodule/redigo](https://github.com/gomodule/redigo)
- GoRedis [go-redis/redis](https://github.com/go-redis/redis)

## C.35.2 Cara koneksi ke Redis

Sebelum kita mulai, pastikan Redis sudah diinstall dan sudah berjalan di sistem operasi dengan baik.
Buatlah file baru dengan nama `main.go`. Di dalam file main kita akan menambahkan perintah-perintah untuk operasi Redis seperti koneksi, create, get data.

Pertama, deklarasikan konfigurasi Redis di dalam variabel:
```
	var host = "127.0.0.1"
	var port = 6379
	var username = ""
	var password = ""
```

Secara default konfigurasi `username` dan `password` Redis itu masih kosong, jika konfigurasi Redis sudah diubah silahkan isi `username` dan `password` dengan yang digunakan.

Buatlah fungsi `connect` dengan menerima parameter dari konfigurasi Redis, return dari fungsi ini adalah interface koneksi (`redis.Conn`) dan error:

```
func connect(host string, port int, username, password string) (redis.Conn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", address, redis.DialUsername(username), redis.DialPassword(password))
	if err != nil {
		return nil, fmt.Errorf("redigo/connection: error redis connection, %v", err)
	}

	return conn, nil
}

```

Didalam fungsi connect kita format alamatnya menjadi `127.0.0.1:6379` dengan bantuan `fmt` library dan kita simpan dengan nama variabel baru `address`. Selanjutnya kita melakukan Dial ke Redis dengan konfigurasi yang kita isi sebelumya. Opsi parameter dari fungsi Dial yang kita gunakan  adalah `DialUsername(username string)` dan `DialPassword(password string)`, selain itu ada opsi-opsi lainnya yang bisa kita memanfaatkan diantaranya:

```
func DialTLSHandshakeTimeout(d time.Duration) DialOption {...}
func DialDatabase(db int) DialOption {...}
func DialTLSConfig(c *tls.Config) DialOption {...}
func DialUseTLS(useTLS bool) DialOption {...}
... etc
```

Opsi - opsi dari fungsi `Dial` ini tidak akan kita bahas semua, hanya mengenalkan bahwa ketika melakukan koneksi ke Redis kita bisa menambahkan parameter-parameter tertentu yang bisa kita gunakan sesuai dengan konfigurasi Redis yang digunakan.

Sekarang kita balik ke fungsi `main`, buatlah perintah untuk memanggil fungsi `connect` dengan membawa parameter `host`, `port`,`username`, `password`. Jangan lupe untuk menangani `error` ketika memanggil fungsi tersebut.

```
	conn, err := connect(host, port, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

```

Maka keseluruhkan kode yang kita buat kurang lebih seperti ini:

```

package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// main function
func main() {
	var host = "127.0.0.1"
	var port = 6379
	var username = ""
	var password = ""

	conn, err := connect(host, port, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Close()
	fmt.Println("Redis connected.")
}

// connect to redis with redigo library
func connect(host string, port int, username, password string) (redis.Conn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", address, redis.DialUsername(username), redis.DialPassword(password))
	if err != nil {
		return nil, fmt.Errorf("redigo/connection: error redis connection, %v", err)
	}

	return conn, nil
}
```

## C.35.3 Cara menyimpan data ke Redis dengan perintah SET.

Pada bab ini kita akan belajar cara menyimpan data ke Redis, perintah dari Redis yang akan kita gunakan adalah `SET`. Sebagai pengenalan perintah untuk menyimpan data itu tidak hanya `SET` tapi ada  banyak di antaranya: `SETTEX`, `HMSET`, `SETBIT`, `LSET`, RSET, `HSETNX`, masing-masing perintah memiliki maksud sendiri. Kita bisa melihat dokumentasi lengkapnya di [Redis Command](https://redis.io/commands).

Sebelumnya kita sudah membuat fungsi untuk melakukan koneksi ke Redis, selanjutnya ubah kode `conn.Close()` di `main` menjadi seperti berikut:

```
	fmt.Println("Redis connected.")

	reply, err := set(conn, "SET", "key-1", "Hello Redis", "NX")
	if err != nil {
		fmt.Println(err)
		return
	}

	if reply == "OK" {
		fmt.Println("data berhasil disimpan.")
	}
```

Perintah `SET` memiliki beberapa opsi yang bisa kita gunakan, yaitu:
- EX seconds -- SET dengan spesifikasi waktu kadaluarsa dalam detik.
- PX milliseconds -- SET dengan spesifikasi waktu kadaluarsa dalam milidetik.
- EXAT timestamp-seconds -- SET dengan spesifikasi Unix waktu kapan key akan kadaluarsa dalam detik.
- PXAT timestamp-milliseconds -- SET dengan spesifikasi Unix waktu kapan key akan kadaluarsa dalam milidetik.
- NX -- SET Hanya jika key tidak ada.
- XX -- SET Hanya jika key ada.

Jadi maksud dari `"SET", "key-1", "Hello Redis", "NX"` adalah menyimpan data baru "Hello Redis" dengan nama key adalah `key-1` dengan kondisi jika key dari `key-1` sebelumnya tidak ada.

Dari kode diatas ketika kita simpan akan memunculkan error karena fungsi `set` belom kita buat, selanjutnya buat fungsi set dengan parameter `Conn` dan array interface:

```
func set(conn redis.Conn, options ...interface{}) (string, error) {
	if err := conn.Close(); err != nil {
		return "", err
	}

	reply, err := redis.String(conn.Do("SET", options))
	if err != nil {
		return "", err
	}

	return reply, nil
}
```

Didalam fungsi `set` ini kita fokus hanya melakukan panggilan perintah `SET` dari Redis. Sehingga seluruh kode yang kita buat menjadi seperti berikut:

```
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// main function
// this describes how to connect to redis with redigo library
func main() {
	var host = "127.0.0.1"
	var port = 6379
	var username = ""
	var password = ""

	conn, err := connect(host, port, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Redis connected.")

	key := "key-1"
	reply, err := set(conn, key, "Hello Redis", "NX")
	if err != nil {
		fmt.Printf("redigo/set: error set value, %v", err)
		return
	}

	if reply != "OK" {
		fmt.Println("data sudah ada.")
		return
	}
}

func connect(host string, port int, username, password string) (redis.Conn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", address, redis.DialUsername(username), redis.DialPassword(password))
	if err != nil {
		return nil, fmt.Errorf("redigo/connection: error redis connection, %v", err)
	}

	return conn, nil
}

func set(conn redis.Conn, options ...interface{}) (string, error) {
	if err := conn.Close(); err != nil {
		return "", err
	}

	reply, err := redis.String(conn.Do("SET", options...))
	if err != nil {
		return "", err
	}

	return reply, nil
}
```

## C.35.4 Cara mengambil data dari Redis dengan perintah GET.

Perintah `GET` di Redis tidak memiliki argument seperti perintah `SET`, yang dibutuhkan ketika menjalankan perintah ini hanya key saja. Kondisi jika key tidak ada maka Redis akan mengembalikan nilain `nil` dan jika ada akan mengembalikan nilai string. Langsung saja buat fungsi get dengan param `Conn` dan key, seperti berikut:

```
func get(conn redis.Conn, key string) (string, error) {
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
```

Selanjutnya, kita panggil di fungsi `main`, seperti berikut:

```
	data, err := get(conn, key)
	if err != nil {
		fmt.Printf("redigo/get: error get value of key %s, %v", key, err)
		return
	}

	fmt.Println("Data of Key:", data)
```

Jika kita lihat secara keseluruhan kode akan seperti berikut:

```
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// main function
// this describes how to connect to redis with redigo library
func main() {
	var host = "127.0.0.1"
	var port = 6379
	var username = ""
	var password = ""

	conn, err := connect(host, port, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Redis connected.")

	key := "key-1"
	reply, err := set(conn, key, "Hello Redis", "NX")
	if err != nil {
		fmt.Printf("redigo/set: error set value, %v", err)
		return
	}

	if reply != "OK" {
		fmt.Println("data sudah ada.")
		return
	}

	data, err := get(conn, key)
	if err != nil {
		fmt.Printf("redigo/get: error get value of key %s, %v", key, err)
		return
	}

	fmt.Println("Data of Key:", data)
}

func connect(host string, port int, username, password string) (redis.Conn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", address, redis.DialUsername(username), redis.DialPassword(password))
	if err != nil {
		return nil, fmt.Errorf("redigo/connection: error redis connection, %v", err)
	}

	return conn, nil
}

func set(conn redis.Conn, options ...interface{}) (string, error) {
	reply, err := redis.String(conn.Do("SET", options...))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func get(conn redis.Conn, key string) (string, error) {
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
```

> Note
Response data yang dikembalikan dari perintah Redis adalah interface, kita bisa mengubahnya sesuai dengan kebutuhan aplikasi yang dibuat. Seperti perintah `GET` bisa diubah responsenya menjadi type `[]byte` dengan helper `redis.Bytes` kodenya menjadi  `redis.Bytes(conn.Do("GET", key))`.
