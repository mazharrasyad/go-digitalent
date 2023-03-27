package controllers

import (
	"fmt"
	"net/http"
	"project-2/database"
	"project-2/models"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	BookID    int       `json:"id"`
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	var maxBookID int
	if err := db.Table("books").Select("COALESCE(MAX(id), 0)").Scan(&maxBookID).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	book.BookID = maxBookID + 1
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	Book := models.Book{
		NameBook: book.NameBook,
		Author:   book.Author,
	}

	err := db.Create(&Book).Error

	if err != nil {
		fmt.Println("Error creating book data:", err)
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	db := database.GetDB()

	book := models.Book{}

	err := db.First(&book, "id = ?", bookID).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func GetAllBook(ctx *gin.Context) {
	db := database.GetDB()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}

	if len(books) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "No books found",
		})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	db := database.GetDB()

	var updatedBook Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&updatedBook).Where("id = ?", bookID).Updates(models.Book{
		NameBook:  updatedBook.NameBook,
		Author:    updatedBook.Author,
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		fmt.Println("Error updating book data:", err)
		return
	}

	book := models.Book{}

	err = db.First(&book, "id = ?", bookID).Error

	if err != nil {
		fmt.Println("Error book data not found:", err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	db := database.GetDB()

	book := models.Book{}

	err := db.Where("id = ?", bookID).Delete(&book).Error

	if err != nil {
		fmt.Println("Error deleting book:", err.Error())
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfully deleted", bookID),
	})
}
