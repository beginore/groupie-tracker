package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	artists     []Artist
	TemplateDir = "../web/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	err := GetAPI()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	tmplPath := filepath.Join(TemplateDir, "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, artists); err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id_str := r.URL.Path[len("/artist/"):]
	if id_str == "" {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(id_str)
	if err != nil || id < 1 {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if id > len(artists) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	artist := artists[id-1]
	err = fetchRelationsForArtist(&artist)
	if err != nil {
		fmt.Print(artist.Name)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join(TemplateDir, "artist.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, artist); err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
