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
)

type M map[string]interface{}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func main() {
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareAuth)

	mux.HandleFunc("/login", HandlerLogin)
	mux.HandleFunc("/index", HandlerIndex)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":8080"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

func authenticateUser(username, password string) (bool, M) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "registered_users.txt")
	buf, _ := ioutil.ReadFile(dbPath)

	for _, each := range strings.Split(string(buf), "\n") {
		if strings.TrimSpace(each) == "" {
			continue
		}

		parts := strings.Split(each, ":")
		if username == parts[0] && password == parts[1] {
			return true, M{
				"username": username,
				"email":    parts[2],
				"grant":    parts[3],
			}
		}
	}

	return false, nil
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Username string
		Password string
	}{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, userInfo := authenticateUser(payload.Username, payload.Password)
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	claims := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		jwt.MapClaims(userInfo),
	)

	token, err := claims.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(M{"token": token})
	w.Write([]byte(tokenString))
}

func MiddlewareAuth(next http.Handler) http.Handler {
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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", M(claims))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(M)
	message := fmt.Sprintf("hello %s (%s)", userInfo["username"], userInfo["grant"])
	w.Write([]byte(message))
}
