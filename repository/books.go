package repository

import (
	"database/sql"
	"APIBooks/structs"
)

func GetAllBooks(db *sql.DB) (result []structs.Book, err error){
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Book
		err = rows.Scan(&book.ID, & book.Title, & book.CategoryID, & book.Description, & book.ImageURL, & book.ReleaseYear, & book.Price, & book.TotalPages, & book.Thickness, & book.CreatedAt, & book.CreatedBy, & book.ModifiedAt, & book.ModifiedBy)
		if err != nil {
			return
		}
		result = append(result, book)
	}
	return
}

func InsertBooks(db *sql.DB, books []structs.Book) (err error) {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare("INSERT INTO books (title, category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, book := range books {
        _, err = stmt.Exec(book.Title, book.CategoryID, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPages, book.Thickness, book.CreatedAt, book.CreatedBy)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}

func GetBookByID(db *sql.DB, id int) (structs.Book, error) {
	var book structs.Book
    row := db.QueryRow("SELECT * FROM books WHERE id = $1", id)
    err := row.Scan(&book.ID, &book.Title, &book.CategoryID, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPages, &book.Thickness, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
    return book, err
}

func DeleteBook(db *sql.DB, id int) error {
    sql := "DELETE FROM books WHERE id = $1"

    _, err := db.Exec(sql, id)
    return err
}