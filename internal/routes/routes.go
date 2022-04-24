package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zepyrshut/go-web-starter-gin/internal/config"
	"github.com/zepyrshut/go-web-starter-gin/internal/handlers"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() *gin.Engine {

	router := gin.Default()

	store := cookie.NewStore([]byte(app.Config.Session.Secret))

	// CORS and CSRF protection
	router.Use(cors.Default())
	router.Use(sessions.Sessions("session", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret:    app.Config.Session.Secret,
		ErrorFunc: func(c *gin.Context) { c.String(400, "CSRF token mismatch") },
	}))

	// Status
	router.GET("/status", handlers.Repo.GetStatusHandler)

	// Pets
	router.GET("/one-pet/:id", handlers.Repo.GetOnePet)

	// Testing CSRF and sessions
	router.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	router.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	return router

}
