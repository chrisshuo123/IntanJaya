package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

// Web is about making certain amout of Requests
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("emailUser")
	password := r.FormValue("passwordUser")

	if password != "blyatiful" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		d := struct {
			Email    string
			Password string
		}{
			Email:    email,
			Password: password,
		}
		tpl.ExecuteTemplate(w, "processor.html", d)
	}
}
