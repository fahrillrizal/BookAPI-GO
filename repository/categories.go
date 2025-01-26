package repository

import (
    "database/sql"
    "APIBooks/structs"
)

func GetAllCategories(db *sql.DB) (result []structs.Category, err error){
    rows, err:= db.Query("SELECT * FROM categories")
    if err != nil{
        return
    }
    defer rows.Close()

    for rows.Next() {
        var category structs.Category
        err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
        if err != nil {
            return nil, err
        }
        result = append(result, category)
    }
    return result, nil
}

func InsertCategories(db *sql.DB, categories []structs.Category) (err error) {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare("INSERT INTO categories (name, created_at, created_by) VALUES ($1, $2, $3)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, category := range categories {
        _, err = stmt.Exec(category.Name, category.CreatedAt, category.CreatedBy)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}

func GetCategoryByID(db *sql.DB, id int) (structs.Category, error) {
    var category structs.Category
    row := db.QueryRow("SELECT * FROM categories WHERE id = $1", id)
    err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
    return category, err
}

func DeleteCategory(db *sql.DB, id int) error {
    sql := "DELETE FROM categories WHERE id = $1"

    _, err := db.Exec(sql, id)
    return err
}

func GetBooksByCategory(db *sql.DB, categoryID int) ([]structs.Book, error) {
    var books []structs.Book
    rows, err := db.Query("SELECT * FROM books WHERE category_id = $1", categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var book structs.Book
        err := rows.Scan(&book.ID, &book.Title, &book.CategoryID, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPages, &book.Thickness, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
        if err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}