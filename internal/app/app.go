package app

import (
	"ebox-api/internal/config"
	"ebox-api/internal/db"
	"ebox-api/internal/middlewares"
	"ebox-api/internal/response"
	"ebox-api/pkg/auth"
	"ebox-api/pkg/boxes"
	"ebox-api/pkg/users"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct {
	Router *gin.Engine
	DB *db.DB
	Config *config.AppConfig
}

func notFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, response.Create(nil, errors.New("Page not found")))
}

func Create (config *config.Config) *App {
	r := gin.Default()
	r.Use(gin.Logger()) // Logs to console
	r.Use(gin.Recovery()) // Recovers from panics and returns 500
	r.NoRoute(notFoundHandler)

	apiRouter := r.Group("/api")

	dbStore, err := db.New(config.DB)
	if err != nil {
		panic(err)
	}

	usersRepository := users.NewUsersRepository(dbStore)

	usersService := users.NewUsersService(usersRepository)
	authService := auth.NewAuthService(usersRepository)
	boxesService := boxes.NewService(dbStore)

	usersHandlers := users.NewUsersHandlers(usersService)
	boxesHandlers := boxes.NewHandlers(boxesService)
	authHandlers := auth.NewAuthHandlers(authService)

	md := middlewares.Initialize(authService)

	apiRouter.POST("/auth/sign-in", authHandlers.SignIn)

	apiRouter.GET("/boxes/:boxID", md.AuthRequired, boxesHandlers.GetBoxById)
	apiRouter.PUT("/boxes", md.AuthRequired, boxesHandlers.PutBox)

	apiRouter.POST("/users", md.AuthRequired, usersHandlers.PostUser)
	apiRouter.GET("/users/me", md.AuthRequired, usersHandlers.GetMe)

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
