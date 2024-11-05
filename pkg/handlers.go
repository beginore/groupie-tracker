package pkg

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
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

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var results []map[string]string

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, map[string]string{
				"type": "artist/band", "name": artist.Name, "url": fmt.Sprintf("/artist/%d", artist.Id),
			})
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				results = append(results, map[string]string{
					"type": "member", "name": member, "url": fmt.Sprintf("/artist/%d", artist.Id),
				})
			}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, map[string]string{
				"type": "first album date", "name": artist.FirstAlbum, "url": fmt.Sprintf("/artist/%d", artist.Id),
			})
		}
		if strings.Contains(strconv.Itoa(artist.CreationDate), query) {
			results = append(results, map[string]string{
				"type": "creation date", "name": strconv.Itoa(artist.CreationDate), "url": fmt.Sprintf("/artist/%d", artist.Id),
			})
		}
		for location := range artist.DatesLocations {
			if strings.Contains(strings.ToLower(location), query) {
				results = append(results, map[string]string{
					"type": "location", "name": location, "url": fmt.Sprintf("/artist/%d", artist.Id),
				})
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
