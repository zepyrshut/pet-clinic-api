package main

import (
	"flag"
	"log"
	"os"

	"github.com/zepyrshut/go-web-starter-gin/internal/config"
	"github.com/zepyrshut/go-web-starter-gin/internal/driver"
	"github.com/zepyrshut/go-web-starter-gin/internal/handlers"
	"github.com/zepyrshut/go-web-starter-gin/internal/middleware"
	"github.com/zepyrshut/go-web-starter-gin/internal/routes"
	"github.com/zepyrshut/go-web-starter-gin/internal/util"
)

// Application properties
const version = "0.0.1-beta.1"
const environment = "development"
const inProduction = false

// Initalize application
var app config.Application

func main() {
	// Environment variables
	dsn := util.GoDotEnvVariable("DATA_SOURCE_NAME")
	apiPort := util.GoDotEnvVariable("API_PORT")
	csrfToken := util.GoDotEnvVariable("CSRF_TOKEN")

	// Application flags
	// Port
	flag.StringVar(&app.Config.Port, "port", apiPort, "Port to listen")
	// Version and environment
	flag.StringVar(&app.Status.Version, "version", version, "Version")
	flag.StringVar(&app.Status.Environment, "env", environment, "Environment")
	flag.BoolVar(&app.InProduction, "production", inProduction, "Production")
	flag.StringVar(&app.Config.Session.Secret, "secret", csrfToken, "Secret")
	// Database
	flag.StringVar(&app.Config.DB.DSN, "dsn", dsn, "Database DSN")
	flag.Parse()

	// Logging format
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize database
	db, err := driver.ConnectSQL(app.Config.DB.DSN)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	defer db.SQL.Close()

	// Initialize handlers and routes
	routes.NewRoutes(&app)
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	middleware.NewMiddleware(&app)
	srv := routes.Routes()

	// Start server
	err = srv.Run(":" + app.Config.Port)
	if err != nil {
		app.ErrorLog.Println(err)
	}

}
