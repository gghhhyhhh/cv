package main

import (
	"log"
	"net/http"
)

func main() {

	// Fichiers CSS, JavaScript et image
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Fichier PDF
	http.Handle(
		"/downloads/",
		http.StripPrefix(
			"/downloads/",
			http.FileServer(http.Dir("downloads")),
		),
	)

	// Page de téléchargement
	http.HandleFunc("/downloads", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/downloads" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "templates/downloads.html")
	})

	// Page principale du CV
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "templates/index.html")
	})

	log.Println("CV disponible sur http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
