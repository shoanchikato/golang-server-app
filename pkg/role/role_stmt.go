package role

const (
	ADD_ROLE_STMT     = `INSERT INTO roles (name) VALUES ($1);`
	EDIT_ROLE_STMT    = `UPDATE roles SET	name = $1 WHERE id = $2;`
	GET_ALL_ROLE_STMT = `SELECT * FROM roles;`
	GET_ONE_ROLE_STMT = `SELECT * FROM roles WHERE id = $1;`
	REMOVE_ROLE_STMT  = `DELETE FROM roles WHERE id = $1;`
)
