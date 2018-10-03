package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type M map[string]interface{}

var SEARCH_MAX_DURATION = 4 * time.Second
var GOOGLE_SEARCH_API_KEY = "ASSVnHfjD_ltXXXXSyB6WWWWWWWWveMFgE"

func main() {
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareUtility)

	mux.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		keyword := r.URL.Query().Get("keyword")
		chanRes := make(chan []byte)
		chanErr := make(chan error)

		go doSearch(ctx, keyword, chanRes, chanErr)

		select {
		case res := <-chanRes:
			w.Header().Set("Content-type", "application/json")
			w.Write(res)
		case err := <-chanErr:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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

func doSearch(
	ctx context.Context,
	keyword string,
	chanRes chan []byte,
	chanErr chan error,
) {
	innerChanRes := make(chan []byte)
	innerChanErr := make(chan error)

	url := "https://www.googleapis.com/customsearch/v1"
	url = fmt.Sprintf("%s?key=%s", url, GOOGLE_SEARCH_API_KEY)
	url = fmt.Sprintf("%s&cx=017576662512468239146:omuauf_lfve", url)
	url = fmt.Sprintf("%s&callback=hndlr", url)
	url = fmt.Sprintf("%s&q=%s", url, keyword)

	from := ctx.Value("from").(string)

	ctx, cancel := context.WithTimeout(ctx, SEARCH_MAX_DURATION)
	defer cancel()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		innerChanErr <- err
		return
	}

	req = req.WithContext(ctx)
	req.Header.Set("Referer", from)

	transport := new(http.Transport)
	client := new(http.Client)
	client.Transport = transport

	go func() {
		resp, err := client.Do(req)
		if err != nil {
			innerChanErr <- err
			return
		}

		if resp != nil {
			defer resp.Body.Close()
			resData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				innerChanErr <- err
				return
			}
			innerChanRes <- resData
		} else {
			innerChanErr <- errors.New("No response")
		}
	}()

	select {
	case <-ctx.Done():
		transport.CancelRequest(req)
		chanErr <- errors.New("Search proccess exceed timeout")
		return
	case res := <-innerChanRes:
		chanRes <- res
		return
	case err := <-innerChanErr:
		transport.CancelRequest(req)
		chanErr <- err
		return
	}
}
