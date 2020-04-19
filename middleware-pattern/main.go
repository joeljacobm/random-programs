package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "I am testing middleware in golang")

}

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Before servehttp call")
		h.ServeHTTP(w, r)
		fmt.Fprintln(w, "After servehttp call")
	})
}

func main() {
	http.HandleFunc("/", middleware(indexHandler))
	http.ListenAndServe(":8080", nil)
}
