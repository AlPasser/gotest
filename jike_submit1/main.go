package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	version := os.Getenv("VERSION")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		for k, v := range r.Header {
			h.Set(k, strings.Join(v, ";"))
			// fmt.Printf("%+v=%+v\n%+v\n", k, v, h.Get(k))
		}
		h.Set("VERSION", version)
		// fmt.Printf("%+v\n", h)
		fmt.Printf("client=%+v, rsp code=200\n", r.Host)
		w.WriteHeader(200)
	})
	http.ListenAndServe("0.0.0.0:80", nil)
}
