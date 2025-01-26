-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT NOT NULL REFERENCES categories(id),
    description TEXT,
    image_url TEXT,
    release_year INT NOT NULL,
    price INT NOT NULL,
    total_page INT NOT NULL,
    thickness VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255)
);

-- +migrate StatementEnd