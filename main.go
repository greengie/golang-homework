package main

import (
	// "fmt"
	"log"
	"net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)

func ReverseString(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var startstring = params["string"]
  json.NewEncoder(w).Encode(Reverse(startstring))
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/reversestring/{string}", ReverseString).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
