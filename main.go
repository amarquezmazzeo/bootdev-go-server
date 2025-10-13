package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Add("Content-Type", "text/plain; charset=utf-8")
	respWriter.WriteHeader(200)
	respWriter.Write([]byte("OK\n"))

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", healthHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("starting serv")
	log.Fatal(server.ListenAndServe())

}
