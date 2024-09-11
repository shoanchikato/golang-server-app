package stmt

const (
	ADD_BOOK_STMT = `
		INSERT INTO books (name, year) VALUES ($1, $2);
		INSERT INTO authors_books (author_id, book_id) VALUES ($3, LAST_INSERT_ROWId());
	`
	EDIT_BOOK_STMT = `
		UPDATE books SET	name = $1, year = $2 WHERE id = $4;
		INSERT INTO authors_books (author_id, book_id) VALUES ($3, $4);
	`
	GET_ALL_BOOK_STMT = `
		SELECT 
			b.id,
			b.name,
			b.year,
			ab.author_id
		FROM 
				books b
		JOIN 
				authors_books ab ON b.id = ab.book_id
		WHERE b.id > $1 LIMIT $2;
	`
	GET_ONE_BOOK_STMT = `
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
				b.id = $1;
	`
	REMOVE_BOOK_STMT = `
		DELETE FROM books WHERE id = $1;
		DELETE FROM authors_books WHERE book_id = $2;
	`
)
