// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type database map[string]string

// func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/foo":
// 		fmt.Fprintf(w, "kor: %s\n", d["kor"])
// 	case "/bar":
// 		fmt.Fprintf(w, "eng: %s\n", d["eng"])
// 	default:
// 		http.NotFound(w, r)
// 	}
// }

// func main() {
// 	db := database{
// 		"kor": "안녕!",
// 		"eng": "Hello!",
// 	}

// 	http.ListenAndServe("localhost:8000", db)
// }
