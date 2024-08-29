package stmt

const (
	CREATE_PERMISSION_TABLE_STMT = `
		CREATE TABLE IF NOT EXISTS permissions (
			id INTEGER PRIMARY KEY,
			name TEXT,
			entity TEXT,
			operation TEXT
		);
		CREATE TABLE IF NOT EXISTS roles (
			id INTEGER PRIMARY KEY,
			name TEXT
		);
		CREATE TABLE IF NOT EXISTS roles_permissions (
			role_id INT,
			permission_id INT,
			FOREIGN KEY(role_id) REFERENCES roles(id),
			FOREIGN KEY(permission_id) REFERENCES permissions(id)
		);
		CREATE TABLE IF NOT EXISTS users_roles (
			user_id INT,
			role_id INT,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(role_id) REFERENCES roles(id)
		);

		-- create indexes
		CREATE INDEX IF NOT EXISTS idx_roles_permissions_role_id_permission_id ON roles_permissions (role_id, permission_id);
		CREATE INDEX IF NOT EXISTS idx_users_roles_user_id_role_id ON users_roles (user_id, role_id);
	`

	ADD_PERMISSION_STMT     = `INSERT INTO permissions (name, entity, operation) VALUES ($1, $2, $3);`
	EDIT_PERMISSION_STMT    = `UPDATE permissions SET	name = $1, entity = $2, operation = $3 WHERE id = $4;`
	GET_ALL_PERMISSION_STMT = `SELECT * FROM permissions WHERE id > $1 LIMIT $2;`
	GET_BY_ENTITY_PERMISSION_STMT = `SELECT * FROM permissions WHERE entity = $1;`
	GET_ONE_PERMISSION_STMT = `SELECT * FROM permissions WHERE id = $1;`
	REMOVE_PERMISSION_STMT  = `
		DELETE FROM permissions WHERE id = $1;
		DELETE FROM roles_permissions WHERE permission_id = $2;
	`
)
