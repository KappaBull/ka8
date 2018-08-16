package ka8

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

//Profile YamlImportData
type Profile struct {
	Name  string  `yaml:"name"`
	Email string  `yaml:"email"`
	Bio   BioData `yaml:"bio"`
}

//BioData YamlImportData
type BioData struct {
	Short string `yaml:"short"`
	Long  string `yaml:"long"`
}

//SiteSettings サイト構成データ
type SiteSettings struct {
	profile Profile
}

var templates = make(map[string]*template.Template)
var setting SiteSettings

func init() {
	importYaml()
	initializeTemplates()
	r := gin.New()
	r.SetHTMLTemplate(templates["welcome.html"])
	r.Static("/css", "./css")
	r.Static("/js", "./js")

	defineRoutes(r)
}

func importYaml() {
	buf, err := ioutil.ReadFile("settings/profile.yaml")
	if err != nil {
		panic(err)
	}
	// structにUnmasrshal
	err = yaml.Unmarshal(buf, &setting.profile)
	if err != nil {
		panic(err)
	}
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
			"Name":         setting.profile.Name,
			"LongMessage":  setting.profile.Bio.Long,
			"ShortMessage": setting.profile.Bio.Short,
			"Email":        setting.profile.Email,
			"AvaterSize":   iconSize,
			"GravatarURL":  getGravatarURL(iconSize),
			"Date":         time.Now().Format("20060102"),
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
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles("templates/theme.html", layout))
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
	hasher.Write([]byte(setting.profile.Email))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hasher.Sum(nil)) + ".jpg?s=" + size
}

// func compileTemplates(filenames ...string) (*template.Template, error) {
// 	m := minify.New()
// 	m.AddFunc("text/html", html.Minify)

// 	var tmpl *template.Template
// 	for _, filename := range filenames {
// 		name := filepath.Base(filename)
// 		if tmpl == nil {
// 			tmpl = template.New(name)
// 		} else {
// 			tmpl = tmpl.New(name)
// 		}

// 		b, err := ioutil.ReadFile(filename)
// 		if err != nil {
// 			return nil, err
// 		}

// 		mb, err := m.Bytes("text/html", b)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tmpl.Parse(string(mb))
// 	}
// 	return tmpl, nil
// }
