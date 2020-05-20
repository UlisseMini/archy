package main

import (
	"fmt"
	"net/http"
)

func main() {
	hits := 0 // TODO: Prevent race
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hits++
		fmt.Fprintf(w, "The api has been called %d times!\n", hits)
	})

	panic(http.ListenAndServe(":80", nil))
}
