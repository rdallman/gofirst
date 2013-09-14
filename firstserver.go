package main

import (
  "fmt"
  "net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
  wordbey := r.URL.Path[1:]
  fmt.Fprintf(w, "How much %s", wordbey)
  fmt.Fprintf(w, " would a %sbey chuck,", wordbey)
  fmt.Fprintf(w, "\nif a %sbey could chuck %s?", wordbey, wordbey)
}

func main() {
  http.HandleFunc("/", simpleHandler)
  http.ListenAndServe(":8080", nil)
}
