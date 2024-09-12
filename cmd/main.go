package main

import (
	"fmt"
	"groupie/pkg"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static"))))
	http.HandleFunc("/", pkg.IndexHandler)
	http.HandleFunc("/artist/", pkg.ArtistHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
