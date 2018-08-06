package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintf(w, "Server GO response")
// 	})

// 	server := http.ListenAndServer(":4200", nil)

// 	log.Fatal(server)
// }
func main() {
  http.HandleFunc("/", handler) // each request calls handler
  http.HandleFunc("/about", aboutHandler)
  http.HandleFunc("/red", redColor)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func aboutHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    //...
}

func redColor(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "roses are red")
}