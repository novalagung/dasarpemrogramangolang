package main

import (
	"context"
	"fmt"
	"net/http"
)

type M map[string]interface{}

func main() {
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareUtility)

	mux.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		from := r.Context().Value("from").(string)
		w.Write([]byte(from))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":80"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

func MiddlewareUtility(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		from := r.Header.Get("Referer")
		if from == "" {
			from = r.Host
		}

		ctx = context.WithValue(ctx, "from", from)

		requestWithContext := r.WithContext(ctx)
		next.ServeHTTP(w, requestWithContext)
	})
}
