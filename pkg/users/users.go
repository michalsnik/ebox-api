package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Register (router *gin.RouterGroup, db *sql.DB) {
	router.GET("/users", GetUsers)
}