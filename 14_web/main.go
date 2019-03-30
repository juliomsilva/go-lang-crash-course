package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1><span style='font-family:sans-serif;color:#b9f2ff'>Shining Bright Like Diamonds</span><h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
