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

// GetAllCategories godoc
// @Summary Get all categories
// @Description Menampilkan semua kategori yang ada di database
// @Tags Categories
// @Produce json
// @Success 200 {object} structs.Category "List of categories"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
    categories, err := repository.GetAllCategories(database.DbConnection)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"result": categories})
}

// InsertCategories godoc
// @Summary Add multiple categories
// @Description Menambahkan beberapa kategori ke database
// @Tags Categories
// @Accept json
// @Produce json
// @Param categories body []structs.Category true "List of categories"
// @Success 200 {object} structs.Category "Success message"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories/batch [post]
func InsertCategories(c *gin.Context) {
    var categories []structs.Category
    if err := c.ShouldBindJSON(&categories); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    username := c.GetString("username")
    for i := range categories {
        categories[i].CreatedAt = time.Now()
        categories[i].CreatedBy = username
    }

    err := repository.InsertCategories(database.DbConnection, categories)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Categories added successfully"})
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Menampilkan kategori sesuai ID yang diminta
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} structs.Book "Category details"
// @Failure 404 {object} map[string]string "Category not found"
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    category, err := repository.GetCategoryByID(database.DbConnection, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"result": category})
}

// GetBooksByCategoryID godoc
// @Summary Get books by category ID
// @Description Menampilkan semua buku yang ada di kategori tertentu
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} structs.Book "List of books"
// @Failure 404 {object} map[string]string "Category not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories/{id}/books [get]
func GetBooksByCategoryID(c *gin.Context) {
    categoryID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    books, err := repository.GetBooksByCategory(database.DbConnection, categoryID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(books) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No books found for this category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"result": books})
}

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description Menghapus kategori sesuai ID
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} structs.Category "Success message"
// @Failure 404 {object} map[string]string "Category not found"
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := repository.DeleteCategory(database.DbConnection, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}