package main

import (
	"log"
	"net/http"
	"strings"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	code := r.URL.Path
	code = strings.TrimPrefix(code, "/")
	if url, ok := urls[code]; ok {
		http.Redirect(w, r, url, 302)
		log.Printf("status: redirected,  code: %v, url: %v", code, url)
	} else {
		log.Printf("action: failed,  code: %v", code)
		http.Error(w, "Code not found", 404)
	}
}
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/", redirect)
	log.Printf("Starting")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Oops")
		panic(err)
	}
	log.Printf("Bye")
}
