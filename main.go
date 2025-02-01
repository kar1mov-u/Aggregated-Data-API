package main

import (
	"Aggregated-Data-API/src"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Working")
	})
	mux.HandleFunc("/aggregate/", src.MainHandler)
	fmt.Println("Listening on port 8080:")
	http.ListenAndServe(":8080", mux)

}
