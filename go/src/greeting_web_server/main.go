package main

import (
	"fmt"
	"net/http"
)

func greeting(value string) string {
	return "<b>" + value + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
