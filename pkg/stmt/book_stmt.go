package stmt

const (
	ADD_BOOK_STMT             = `
		INSERT INTO books (name, year) VALUES ($1, $2);
		INSERT INTO authors_books (author_id, book_id) VALUES ($3, LAST_INSERT_ROWId());
	`
	EDIT_BOOK_STMT            = `UPDATE books SET	name = $1, year = $2 WHERE id = $3;`
	GET_ALL_BOOK_STMT         = `SELECT * FROM books WHERE id > $1 LIMIT $2;`
	GET_ONE_BOOK_STMT         = `SELECT * FROM books WHERE id = $1;`
	REMOVE_BOOK_STMT          = `DELETE FROM books WHERE id = $1;`
)
