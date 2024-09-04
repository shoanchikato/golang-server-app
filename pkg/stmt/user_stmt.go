package stmt

const (
	CREATE_USER_TABLE_STMT = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			first_name TEXT,
			last_name TEXT
		);
		CREATE TABLE IF NOT EXISTS auth (
			id INTEGER PRIMARY KEY,
			username TEXT UNIQUE,
			email TEXT UNIQUE,
			password TEXT,
			user_id TEXT UNIQUE
		);
		CREATE TABLE IF NOT EXISTS users_auth (
			user_id INT,
			auth_id INT,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(auth_id) REFERENCES auth(id)
		);

		-- create indexes
		CREATE INDEX IF NOT EXISTS idx_username ON auth (username);
		CREATE INDEX IF NOT EXISTS idx_users_auth_user_id_auth_id ON users_auth (user_id, auth_id);
	`
	ADD_USER_STMT = `
		INSERT INTO users (first_name, last_name) VALUES ($1, $2);

		INSERT INTO auth (username, email, password, user_id) VALUES ($3, $4, $5, LAST_INSERT_ROWId());

		INSERT INTO users_auth (user_id, auth_id) VALUES (LAST_INSERT_ROWId(), (SELECT LAST_INSERT_ROWId() FROM auth));
	`
	EDIT_USER_STMT = `
		UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;

		UPDATE auth SET username = $4, email = $5 WHERE user_id = $6;
	`
	GET_ALL_USER_STMT = `
		SELECT 
			u.id,
			u.first_name,
			u.last_name,
			a.username,
			a.email
		FROM 
			users u
		JOIN 
			users_auth ua ON u.id = ua.user_id
		JOIN 
			auth a ON ua.auth_id = a.id
		WHERE 
			u.id > $1 LIMIT $2;
	`
	GET_ONE_USER_STMT = `
		SELECT 
			u.id,
			u.first_name,
			u.last_name,
			a.username,
			a.email
		FROM 
			users u
		JOIN 
			users_auth ua ON u.id = ua.user_id
		JOIN 
			auth a ON ua.auth_id = a.id
		WHERE 
			u.id = $1;
	`
	REMOVE_USER_STMT = `
		DELETE FROM users WHERE id = $1;
		DELETE FROM auth WHERE user_id = $1;
		DELETE FROM users_auth WHERE user_id = $1;
	`
)
