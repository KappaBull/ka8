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
		addSafeHeaders(c)
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

func getGravatarURL(size string) string {
	hasher := md5.New()
	hasher.Write([]byte(mail))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hasher.Sum(nil)) + ".jpg?s=" + size
}
