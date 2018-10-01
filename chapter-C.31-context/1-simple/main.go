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
		domain := r.Context().Value("domain").(string)
		w.Write([]byte(domain))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":8080"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

func MiddlewareUtility(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		domain := "mysampletestapp.com"
		ctx = context.WithValue(ctx, "domain", domain)

		requestWithContext := r.WithContext(ctx)
		next.ServeHTTP(w, requestWithContext)
	})
}
