package util

import (
	"log"
	"net/http"
	"os"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	csrf "github.com/utrack/gin-csrf"
)

// GoDotEnvVariable returns the value of the environment variable
func GoDotEnvVariable(key string) string {
	err := godotenv.Load("../../environment.env")
	if err != nil {
		log.Fatal("error loading environment variables")
	}
	return os.Getenv(key)
}

// Test session
func Increment(c *gin.Context) {
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
}

// Test CSRF protection
func GetToken(c *gin.Context) {
	c.String(http.StatusOK, csrf.GetToken(c))
}

func PostToken(c *gin.Context) {
	c.String(http.StatusOK, "CSRF token is valid")
}

// Test i18n
func GetLocalize(c *gin.Context) {
	c.String(http.StatusOK, ginI18n.MustGetMessage("welcome"))
}
