package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Background.html")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
