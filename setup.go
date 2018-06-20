package bootstrapquickstart

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	myname  = "KappaBull"
	mail    = "kappa8v11@gmail.com"
	twitter = myname
	qiita   = myname
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
	Title       string
	Message     string
	GravatarURL string
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	message := Welcome{
		Title:       myname,
		Message:     "Hosting to GAE/go",
		GravatarURL: _GetGravatarURL("128"),
	}
	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
}

func _GetGravatarURL(size string) string {
	hasher := md5.New()
	hasher.Write([]byte(mail))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hasher.Sum(nil)) + ".jpg?s=" + size
}
