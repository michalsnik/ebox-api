package app

import (
	"database/sql"
	"ebox-api/internal/config"
	"ebox-api/pkg/boxes"
	"ebox-api/pkg/users"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type App struct {
	Router *gin.Engine
	DB *sql.DB
	Config *config.AppConfig
}

func Create (config *config.Config) *App {
	r := gin.Default()
	r.Use(gin.Logger())

	apiRouter := r.Group("/api")

	dbConnInfo := fmt.Sprintf(`
		host=%s
		port=%d
		user=%s
		password=%s
		dbname=%s
		sslmode=disable`,
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.DBName)

	db, err := sql.Open("postgres", dbConnInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	boxes.Register(apiRouter, db)
	users.Register(apiRouter, db)

	return &App{
		Router: r,
		DB: db,
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
