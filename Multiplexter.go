package main

import (
	"fmt"
	"net/http"
)

type database map[string]string

func (d database) kor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "kor: %s\n", d["kor"])
}

func (d database) eng(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "eng: %s\n", d["eng"])
}

func main() {
	db := database{
		"kor": "안녕!",
		"eng": "Hello!",
	}

	mux := http.NewServeMux()

	mux.Handle("/kor", http.HandlerFunc(db.kor))
	mux.HandleFunc("/eng", db.eng)

	http.ListenAndServe("localhost:8000", mux)
}
