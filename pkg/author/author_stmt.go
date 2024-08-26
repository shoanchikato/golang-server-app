package author

const (
	CREATE_AUTHOR_TABLE_STMT = `
		CREATE TABLE IF NOT EXISTS authors (
			id INTEGER PRIMARY KEY,
			first_name TEXT,
			last_name TEXT
		);
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY,
			name TEXT,
			year INT
		);
		CREATE TABLE IF NOT EXISTS authors_books (
			author_id INT,
			book_id INT,
			FOREIGN KEY(author_id) REFERENCES authors(id),
			FOREIGN KEY(book_id) REFERENCES books(id)
		);

		-- create indexes
		CREATE INDEX IF NOT EXISTS idx_authors_books_author_id_book_id ON authors_books (author_id, book_id);
	`
	GET_BOOKS_BY_AUTHOR_ID_STMT = `
		SELECT 
			b.id,
			b.name,
			b.year,
			ab.author_id
		FROM 
				books b
		JOIN 
				authors_books ab ON b.id = ab.book_id
		JOIN 
				authors a ON ab.author_id = a.id
		WHERE 
				a.id = $1;
	`
	ADD_AUTHOR_STMT     = `INSERT INTO authors (first_name, last_name) VALUES ($1, $2);`
	EDIT_AUTHOR_STMT    = `UPDATE authors SET	first_name = $1, last_name = $2 WHERE id = $3;`
	GET_ALL_AUTHOR_STMT = `SELECT * FROM authors;`
	GET_ONE_AUTHOR_STMT = `SELECT * FROM authors WHERE id = $1;`
	REMOVE_AUTHOR_STMT  = `DELETE FROM authors WHERE id = $1;`
)
