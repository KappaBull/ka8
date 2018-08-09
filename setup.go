<<<<<<< HEAD
package ka8

import (
	"crypto/md5"
	// "crypto/sha"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
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

func addSafeHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
}

//Welcome PageData
type Welcome struct {
	Name        string
	Message     string
	AvaterSize  string
	GravatarURL string
	Date        string
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	addSafeHeaders(w)

	//キャッシュ
	now := time.Now().AddDate(0, 0, 1)
	w.Header().Set("Cache-Control", "max-age=300, public")
	w.Header().Set("Last-Modified", now.Format(http.TimeFormat))

	iconSize := "300"
	message := Welcome{
		Name:        myname,
		Message:     "自宅にブレードサーバーラックがあるゲームのフロント作ってる人",
		AvaterSize:  iconSize,
		GravatarURL: getGravatarURL(iconSize),
		Date:        time.Now().Format("20060102"),
	}
	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
}

func getGravatarURL(size string) string {
	hasher := md5.New()
	hasher.Write([]byte(mail))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hasher.Sum(nil)) + ".jpg?s=" + size
}

// func sriHashGenerate(url_ string) string {
// 	hasher := sha.New()
// 	hasher.Write([]byte(url_))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }
=======
package ka8

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	myname = "KappaBull"
	mail   = "kappa8v11@gmail.com"
)

var templates = make(map[string]*template.Template)

func init() {
	initializeTemplates()
	r := gin.New()
	r.SetHTMLTemplate(templates["welcome.html"])
	r.Static("/css", "./css")
	r.Static("/js", "./js")

	defineRoutes(r)
}

func defineRoutes(r *gin.Engine) {
	iconSize := "300"
	r.GET("/", func(c *gin.Context) {
		//キャッシュ
		now := time.Now().AddDate(0, 0, 1)
		c.Header("Cache-Control", "max-age=300, public")
		c.Header("Last-Modified", now.Format(http.TimeFormat))

		c.HTML(http.StatusOK, "outerTheme", gin.H{
			"Name":        myname,
			"Message":     "自宅にブレードサーバーラックがあるゲームのフロント作ってる人",
			"AvaterSize":  iconSize,
			"GravatarURL": getGravatarURL(iconSize),
			"Date":        time.Now().Format("20060102"),
		})
	})

	http.Handle("/", r)
}

func initializeTemplates() {
	layouts, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "templates/theme.html"))
	}
}

func addSafeHeaders(c *gin.Context) {
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("X-Frame-Options", "DENY")
	c.Header("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
}

// //Welcome PageData
// type Welcome struct {
// 	Name        string
// 	Message     string
// 	AvaterSize  string
// 	GravatarURL string
// 	Date        string
// }

// func welcomeHandler(w http.ResponseWriter, r *http.Request) {
// 	addSafeHeaders(w)

// 	//キャッシュ
// 	now := time.Now().AddDate(0, 0, 1)
// 	w.Header().Set("Cache-Control", "max-age=300, public")
// 	w.Header().Set("Last-Modified", now.Format(http.TimeFormat))

// 	iconSize := "300"
// 	message := Welcome{
// 		Name:        myname,
// 		Message:     "自宅にブレードサーバーラックがあるゲームのフロント作ってる人",
// 		AvaterSize:  iconSize,
// 		GravatarURL: getGravatarURL(iconSize),
// 		Date:        time.Now().Format("20060102"),
// 	}
// 	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
// }

func getGravatarURL(size string) string {
	hasher := md5.New()
	hasher.Write([]byte(mail))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hasher.Sum(nil)) + ".jpg?s=" + size
}
>>>>>>> develop
