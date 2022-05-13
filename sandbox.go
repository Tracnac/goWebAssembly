package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// //go:embed wasm/*
// var htmlFS embed.FS

// func mainHTML(w http.ResponseWriter, r *http.Request) {
// 	content, _ := htmlFS.ReadFile("wasm/index.html")
// 	fmt.Fprint(w, string(content))
// }

func main() {
	server := flag.String("server", ":5555", "server address:port")
	flag.Parse()

	fmt.Printf("listening on %v...\n", *server)
	// http.HandleFunc("/", mainHTML)
	log.Fatal(http.ListenAndServe(*server, http.FileServer(http.Dir("."))))
}
