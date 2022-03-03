// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type database map[string]string

// func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	for country, greeting := range d {
// 		fmt.Fprintf(w, "%s: %s\n", country, greeting)
// 	}
// }

// func main() {
// 	db := database{
// 		"kor": "안녕!",
// 		"eng": "Hello!",
// 	}

// 	http.ListenAndServe("localhost:8000", db)
// }
