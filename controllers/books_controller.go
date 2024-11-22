package controllers

import (
	"database/sql"
	"go-gin-crud/models"
	"go-gin-crud/utils"
	"go-gin-crud/utils/sql_queries"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetBooks(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(sql_queries.GetBooks)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to fetch books")
		return
	}
	defer rows.Close()

	var books []models.Books

	for rows.Next() {
		var book models.Books

		if err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Category, &book.Price, &book.Isbn); err != nil {
			utils.SendErrorResponse(c, http.StatusMultiStatus, "Error scanning books")
			return
		} 
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var book models.Books
	err := db.QueryRow(sql_queries.GetBookByID, id).Scan(&book.Id, &book.Name, &book.Author, &book.Category, &book.Price, &book.Isbn)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Book not found")
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context, db *sql.DB) {
	var book models.Books

	if err := c.ShouldBindJSON(&book); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := db.Exec(sql_queries.CreateBook, book.Name, book.Author, book.Category, book.Price, book.Isbn)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create book")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Book created successfully"})
}

func UpdateBook(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var book models.Books
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := db.Exec(sql_queries.UpdateBook, book.Name, book.Author, book.Category, book.Price, book.Isbn, id)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update book")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Book updated successfully"})
}

func DeleteBook(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec(sql_queries.DeleteBook, id)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Book deleted successfully"})
}