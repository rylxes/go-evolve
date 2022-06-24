package app

import (
	"github.com/gin-gonic/gin"
	"go-postgres/controllers/users"
	"go-postgres/domain"
)

var (
	router = gin.Default()
)

func StartApp() {
	DB := domain.Init()
	h := users.New(DB)

	router.GET("/books/find", h.FindUser)
	router.Run(":8080")
}
