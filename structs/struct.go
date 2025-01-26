package structs

import (
	"time"
	"database/sql"
)

// User represents a user entity.
// @Description User entity with ID, username, password, and timestamps.
type User struct {
	ID         int64        `json:"id"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	CreatedAt  time.Time    `json:"created_at"`
	CreatedBy  string       `json:"created_by"`
	ModifiedAt sql.NullTime `json:"modified_at"`
	ModifiedBy sql.NullString `json:"modified_by"`
}

// Category represents a category entity.
// @Description Category entity with ID, name, and timestamps.
type Category struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	CreatedAt  time.Time    `json:"created_at"`
	CreatedBy  string       `json:"created_by"`
	ModifiedAt sql.NullTime `json:"modified_at"`
	ModifiedBy sql.NullString `json:"modified_by"`
}

// Book represents a book entity.
// @Description Book entity with ID, title, category ID, description, image URL, release year, price, total pages, thickness, and timestamps.
type Book struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	CategoryID  int64        `json:"category_id"`
	Description string       `json:"description"`
	ImageURL    string       `json:"image_url"`
	ReleaseYear int64        `json:"release_year"`
	Price       int          `json:"price"`
	TotalPages  int          `json:"total_pages"`
	Thickness   string       `json:"thickness"`
	CreatedAt   time.Time    `json:"created_at"`
	CreatedBy   string       `json:"created_by"`
	ModifiedAt  sql.NullTime `json:"modified_at"`
	ModifiedBy  sql.NullString `json:"modified_by"`
}