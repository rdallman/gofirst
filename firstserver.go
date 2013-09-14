package main

import (
  "fmt"
  "net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
  wordbey := r.URL.Path[1:] // slice off the '/'
  fmt.Fprintf(w, "How much %s", wordbey)
  fmt.Fprintf(w, " would a %sbey chuck,", wordbey)
  fmt.Fprintf(w, "\nif a %sbey could chuck %s?", wordbey, wordbey)
  // a handler can be as simple as a print...
}

func main() {
  // takes a path and a func (http.ResponseWriter, r *http.Request)
  http.HandleFunc("/", simpleHandler)
  // our server is one line!
  http.ListenAndServe(":8080", nil)
}
