package main

import (
	"flag"
	"log"
	"os"

	"github.com/zepyrshut/pet-clinic/internal/config"
	"github.com/zepyrshut/pet-clinic/internal/driver"
	"github.com/zepyrshut/pet-clinic/internal/handlers"
	"github.com/zepyrshut/pet-clinic/internal/middleware"
	"github.com/zepyrshut/pet-clinic/internal/routes"
	"github.com/zepyrshut/pet-clinic/internal/util"
)

// Application properties
const version = "0.1.1-beta.3"
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
	server := routes.Routes()

	// Start server
	err = server.Run("localhost:" + app.Config.Port)
	if err != nil {
		app.ErrorLog.Println(err)
	}

}
