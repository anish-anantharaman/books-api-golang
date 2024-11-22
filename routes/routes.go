package routes

import (
	"database/sql"
	"go-gin-crud/controllers"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/v1/books", func(c *gin.Context) {
		controllers.GetBooks(c, db)
	})

	router.GET("/v1/books/:id", func(c *gin.Context) {
		controllers.GetBookByID(c, db)
	})

	router.POST("/v1/books", func(c *gin.Context) {
		controllers.CreateBook(c, db)
	})

	router.PUT("/v1/books/:id", func(c *gin.Context) {
		controllers.UpdateBook(c, db)
	})

	router.DELETE("/v1/books/:id", func(c *gin.Context) {
		controllers.DeleteBook(c, db)
	})
}

