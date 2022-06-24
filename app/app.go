package app

import (
	"github.com/gin-gonic/gin"
	"go-postgres/controllers/books"
	"go-postgres/domain"
)

var (
	router = gin.Default()
)

func StartApp() {
	DB := domain.Init()
	h := books.New(DB)

	router.GET("/books/find", h.FindBook)
	router.Run(":8080")
}
