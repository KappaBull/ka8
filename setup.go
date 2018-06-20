package bootstrapquickstart

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates = make(map[string]*template.Template)

func init() {
	initializeTemplates()
	defineRoutes()
}

func defineRoutes() {
	http.HandleFunc("/", welcomeHandler)
}

func initializeTemplates() {
	layouts, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "templates/layouts/theme.html"))
	}
}

type Welcome struct {
	Title   string
	Message string
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	message := Welcome{Title: "KappaBull", Message: "Hosting to GAE/go"}
	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
}
