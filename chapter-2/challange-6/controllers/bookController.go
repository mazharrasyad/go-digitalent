package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// create table books (
// 	id SERIAL PRIMARY KEY,
// 	judul varchar(50) NOT NULL,
// 	penulis varchar(50) NOT NULL,
// 	harga INT NOT NULL,
// 	penerbit varchar(20) NOT NULL
// 	);

type Book struct {
	BookID   int    `json:"book_id"`
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Harga    int    `json:"harga"`
	Penerbit string `json:"penerbit"`
}

var BookDatas = []Book{}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

var (
	db       *sql.DB
	err      error
	psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func init() {
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected to database")
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	// Open the database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var book Book
	sqlStatement := `
        INSERT INTO books (judul, penulis, harga, penerbit)
        VALUES ($1, $2, $3, $4)
        RETURNING *
    `

	// Execute the SQL statement and scan the result into the book variable
	if err := db.QueryRow(sqlStatement, newBook.Judul, newBook.Penulis, newBook.Harga, newBook.Penerbit).
		Scan(&book.BookID, &book.Judul, &book.Penulis, &book.Harga, &book.Penerbit); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	// Open the database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var updatedBook Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
        UPDATE books
        SET judul=$2, penulis=$3, harga=$4, penerbit=$5
        WHERE id=$1
        RETURNING *
    `
	var book Book
	if err := db.QueryRow(sqlStatement, bookID, updatedBook.Judul, updatedBook.Penulis, updatedBook.Harga, updatedBook.Penerbit).
		Scan(&book.BookID, &book.Judul, &book.Penulis, &book.Harga, &book.Penerbit); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	// Open the database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var book Book
	sqlStatement := `SELECT * FROM books WHERE id=$1`

	// Execute the SQL statement and scan the result into the book variable
	if err := db.QueryRow(sqlStatement, bookID).Scan(&book.BookID, &book.Judul, &book.Penulis, &book.Harga, &book.Penerbit); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetAllBook(ctx *gin.Context) {
	// Open the database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var results []Book
	sqlStatement := `SELECT * from books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookID, &book.Judul, &book.Penulis, &book.Harga, &book.Penerbit); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, gin.H{"books": results})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	// Open the database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Execute the SQL statement to delete the book
	sqlStatement := `DELETE FROM books WHERE id=$1`
	res, err := db.Exec(sqlStatement, bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Check if any row was affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// If no row was affected, return error
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfully deleted", bookID),
	})
}
