package boxes

import (
	"ebox-api/internal/db"
	"github.com/gin-gonic/gin"
)

func Register (router *gin.RouterGroup, db *db.DB) {
	svc := NewService(db)
	handlers := NewHandlers(svc)

	router.GET("/boxes/:boxID", handlers.GetBoxById)
	router.PUT("/boxes", handlers.PutBox)
}
