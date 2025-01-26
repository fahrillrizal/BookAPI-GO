package controllers

import (
    "github.com/gin-gonic/gin"
    "APIBooks/database"
    "APIBooks/repository"
    "APIBooks/structs"
    "net/http"
    "strconv"
	"time"
)

// Handler GetAllBooks godoc
// @Summary Get all books
// @Description Menampilkan Semua buku yang ada di database
// @Tags Books
// @Produce json
// @Success 200 {object} map[string]interface{} "List buku"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
    books, err := repository.GetAllBooks(database.DbConnection)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"result": books})
}

// Handler InsertBooks godoc
// @Summary Add multiple books
// @Description Menambahkan beberapa buku ke database
// @Tags Books
// @Accept json
// @Produce json
// @Param books body []structs.Book true "List of books"
// @Success 200 {object} structs.Book "Success message"
// @Failure 400 {object} map[string]string "Invalid request or validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /books/batch [post]
func InsertBooks(c *gin.Context) {
    var books []structs.Book
    if err := c.ShouldBindJSON(&books); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    username := c.GetString("username")
    for i := range books {
        books[i].CreatedAt = time.Now()
        books[i].CreatedBy = username

        if books[i].ReleaseYear < 1980 || books[i].ReleaseYear > 2024 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
            return
        }

        if books[i].TotalPages > 100 {
            books[i].Thickness = "tebal"
        } else {
            books[i].Thickness = "tipis"
        }
    }

    err := repository.InsertBooks(database.DbConnection, books)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Books added successfully"})
}

// Handler GetBookByID godoc
// @Summary Get Book By ID
// @Description Menampilkan buku sesuai ID yang diminta
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} structs.Book "Detail buku"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    book, err := repository.GetBookByID(database.DbConnection, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"result": book})
}

// Handler DeleteBook godoc
// @Summary Delete a buku by ID
// @Description Menghapus buku sesuai ID yang diminta
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} structs.Book "Success message"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := repository.DeleteBook(database.DbConnection, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}