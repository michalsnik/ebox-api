package boxes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Register (router *gin.RouterGroup, db *sql.DB) {
	svc := NewService(db)
	handlers := NewHandlers(svc)

	router.GET("/boxes/:boxID", handlers.GetBoxById)
	router.PUT("/boxes", handlers.PutBox)
}
