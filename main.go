package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	url := os.Getenv("URL")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if url == "" {
		log.Fatal("$URL must be set")
	}

	log.Printf("Listen on port %s and proxy all requests %s", port, url)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Method != http.MethodGet {
				http.Error(w, "Only GET!", http.StatusMethodNotAllowed)
				return
			}
			resp, err := http.Get(url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}
			defer resp.Body.Close()
			copyHeader(w.Header(), resp.Header)
			w.WriteHeader(resp.StatusCode)
			io.Copy(w, resp.Body)
		}),
	}
	log.Fatal(server.ListenAndServe())
}
