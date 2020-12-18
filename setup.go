package ka8

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.New()
	r.Static("/asset", "./asset")
	r.GET("/", func(c *gin.Context) {
		//SafeHeaders
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
		//キャッシュ
		now := time.Now().AddDate(0, 0, 1)
		c.Header("Cache-Control", "max-age=300, public")
		c.Header("Last-Modified", now.Format(http.TimeFormat))
	})

	http.Handle("/", r)
}
