package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/zepyrshut/pet-clinic/internal/config"
)

var app *config.Application

func NewMiddleware(a *config.Application) {
	app = a
}

func Sessions(name string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(app.Config.Session.Secret))
	return sessions.Sessions(name, store)
}

// Cross Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
	config.AllowMethods = []string{"POST", "OPTIONS", "GET", "PUT"}

	return cors.New(config)
}

func Localize() gin.HandlerFunc {
	config := ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "./internal/i18n",
		AcceptLanguage:   []language.Tag{language.English, language.Spanish},
		DefaultLanguage:  language.Spanish,
		UnmarshalFunc:    yaml.Unmarshal,
		FormatBundleFile: "yaml",
	})

	return ginI18n.Localize(config)
}
