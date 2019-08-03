package app

import (
	"ebox-api/internal/config"
	"ebox-api/internal/db"
	"ebox-api/pkg/auth"
	"ebox-api/pkg/boxes"
	"ebox-api/pkg/users"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB *db.DB
	Config *config.AppConfig
}

func Create (config *config.Config) *App {
	r := gin.Default()
	r.Use(gin.Logger())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiRouter := r.Group("/api")

	dbStore, err := db.New(config.DB)

	if err != nil {
		panic(err)
	}

	auth.Register(apiRouter, dbStore)
	users.Register(apiRouter, dbStore)
	boxes.Register(apiRouter, dbStore)

	return &App{
		Router: r,
		DB: dbStore,
		Config: config.App,
	}
}

func (app *App) Run (port string) {
	err := app.Router.Run(port)
	if err != nil {
		panic(err)
	}

	err = app.DB.Ping()
	if err != nil {
		panic(err)
	}
}
