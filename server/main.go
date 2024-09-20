package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, mTLS example!"))
	})

	server := &http.Server{
		Addr:    ":8443",
		Handler: mux,
	}

	log.Println("Starting server on https://localhost:8443")
	err := server.ListenAndServeTLS("tls-certs/server-cert.pem", "tls-certs/server-key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
