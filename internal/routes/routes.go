package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zepyrshut/pet-clinic/internal/config"
	"github.com/zepyrshut/pet-clinic/internal/handlers"
	"github.com/zepyrshut/pet-clinic/internal/middleware"
	"github.com/zepyrshut/pet-clinic/internal/util"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() *gin.Engine {

	router := gin.Default()

	// CORS and CSRF protection
	router.Use(cors.Default())
	router.Use(middleware.Localize())
	router.Use(middleware.Sessions("session"))
	router.Use(middleware.CORSMiddleware())

	// Status and test
	router.GET("/status", handlers.Repo.GetStatusHandler)
	router.GET("/localize", util.GetLocalize)

	// Pets
	router.GET("/pets", handlers.Repo.GetAllPets)
	router.GET("/one-pet/:id", handlers.Repo.GetOnePet)
	router.POST("/new-pet", handlers.Repo.InsertNewPet)

	// People
	router.GET("/one-person/:id", handlers.Repo.GetOnePerson)
	router.POST("/new-person", handlers.Repo.InsertNewPerson)

	// PetOwner
	router.POST("/bind-pet-owner", handlers.Repo.BindPetWithOwner)

	//Testing sessions and CSRF protection
	router.GET("/incr", util.Increment)
	router.GET("/protected", util.GetToken)
	router.POST("/protected", util.PostToken)

	return router

}
