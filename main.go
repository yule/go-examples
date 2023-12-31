package main

import (
  "fmt"
  "log"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello World")
  fmt.Println("Home page!")
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
  handleRequests()
}
