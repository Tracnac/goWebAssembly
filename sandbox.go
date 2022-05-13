package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// //go:embed index.html
// var htmlFS embed.FS

// func mainHTML(w http.ResponseWriter, r *http.Request) {
// 	content, _ := htmlFS.ReadFile("index.html")
// 	fmt.Fprint(w, string(content))
// }

func main() {
	addr := flag.String("addr", ":5555", "server address:port")
	flag.Parse()

	fmt.Printf("listening on %v...\n", *addr)
	// http.HandleFunc("/", mainHTML)
	log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.Dir("."))))
}
