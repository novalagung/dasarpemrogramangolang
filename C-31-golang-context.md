# C.31. Context

Pada bab ini kita akan belajar pemanfaatan `context.Context` untuk keperluan menyisipkan data ke dalam objek `*http.Request`, dan untuk mempermudah handling timeout pada operasi yang membutuhkan waktu lama (seperti http client request).

## C.31.1. Persiapan

Buat folder 

```bash
novalagung:1-simple $ tree .
.
├── main.go
└── middleware.go

0 directories, 2 files
```

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
