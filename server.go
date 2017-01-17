package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	listen := flag.String("listen", ":80", "addr:port or :port to listen on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(assetFS()))

	logger := http.NewServeMux()
	logger.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		mux.ServeHTTP(w, r)
	})

	log.Printf("starting server at %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, logger))
}
