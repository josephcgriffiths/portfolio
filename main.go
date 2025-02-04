package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := map[string]string{
		"/":         "home.html",
		"/about":    "about.html",
		"/projects": "projects.html",
		"/contact":  "contact.html",
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	for path, tmpl := range templates {
		tmpl := tmpl // capture range variable
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			t := template.Must(template.ParseFiles("base.html", "navbar.html", tmpl))
			err := t.Execute(w, map[string]string{"Title": "Home"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}

	log.Fatal(http.ListenAndServe(":7000", nil))
}
