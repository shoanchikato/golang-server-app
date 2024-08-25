package post

const (
	CREATE_POST_TABLE_STMT = `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY,
			title TEXT,
			body TEXT,
			user_id INT
		);
	`
	ADD_POST_STMT     = `INSERT INTO posts (title, body, user_id) VALUES ($1, $2, $3);`
	EDIT_POST_STMT    = `UPDATE posts SET	title = $1, body = $2, user_id = $3 WHERE id = $4;`
	GET_ALL_POST_STMT = `SELECT * FROM posts;`
	GET_ONE_POST_STMT = `SELECT * FROM posts WHERE id = $1;`
	REMOVE_POST_STMT  = `DELETE FROM posts WHERE id = $1;`
)
