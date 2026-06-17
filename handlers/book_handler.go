package handlers

import (
	"api-booktracker/config"
	"api-booktracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	rows, err := config.DB.Query(
		"SELECT id, title, author FROM books",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	err := config.DB.QueryRow(
		"SELECT id, title, author FROM books WHERE id=$1",
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := config.DB.Exec(
		"INSERT INTO books (title, author) VALUES ($1, $2)",
		book.Title,
		book.Author,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
	})

}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := config.DB.Exec(
		"UPDATE books SET title=$1, author=$2 WHERE id=$3",
		book.Title,
		book.Author,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec(
		"DELETE FROM books WHERE id=$1",
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
