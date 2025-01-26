package main

import (
    "database/sql"
    "fmt"
    "APIBooks/database"
    "APIBooks/controllers"
    "APIBooks/middleware"
    "os"
    
    _ "APIBooks/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var (
    DB  *sql.DB
    err error
)

// @title API Books
// @version 1.0
// @description API untuk manajemen buku dan kategori.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@apibooks.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
    // Muat file .env
    err = godotenv.Load(".env")
    if err != nil {
        panic("Error loading .env file")
    }

    // Buat koneksi ke database PostgreSQL
    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
        os.Getenv("PGHOST"),
        os.Getenv("PGPORT"),
        os.Getenv("PGUSER"),
        os.Getenv("PGPASSWORD"),
        os.Getenv("PGDATABASE"),
    )

    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer DB.Close()

    // Ping database untuk memastikan koneksi berhasil
    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    // Jalankan migrasi database
    database.DBMigrate(DB)

    // Inisialisasi router Gin
    router := gin.Default()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    // Endpoint untuk login
    router.POST("/api/users/login", controllers.Login)

    // Grup route untuk API
    api := router.Group("/api")
    {
        // Route untuk categories
        categories := api.Group("/categories")
        {
            categories.GET("/", middleware.AuthMiddleware(), controllers.GetAllCategories)
            categories.POST("/", middleware.AuthMiddleware(), controllers.InsertCategories)
            categories.GET("/:id", middleware.AuthMiddleware(), controllers.GetCategoryByID)
            categories.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteCategory)
            categories.GET("/:id/books", middleware.AuthMiddleware(), controllers.GetBooksByCategoryID)
        }

        // Route untuk books
        books := api.Group("/books")
        {
            books.GET("/", middleware.AuthMiddleware(), controllers.GetAllBooks)
            books.POST("/", middleware.AuthMiddleware(), controllers.InsertBooks)
            books.GET("/:id", middleware.AuthMiddleware(), controllers.GetBookByID)
            books.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteBook)
        }
    }

    // Jalankan server di port 8080
    router.Run(":8080")

    fmt.Println("Berhasil menyambung ke DB!")
}