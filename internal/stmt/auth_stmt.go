package stmt

const (
	GET_AUTH_DETAILS_BY_USERNAME = `SELECT username, email, password, user_id FROM auth WHERE username = $1;`
	RESET_PASSWORD               = `UPDATE auth SET password = $1 WHERE username = $2;`
)
