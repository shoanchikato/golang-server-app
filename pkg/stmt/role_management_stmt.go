package stmt

const (
	GET_ROLE_BY_USER_ID_STMT = `
		SELECT 
			r.id,
			r.name
		FROM 
				roles r
		JOIN 
				users_roles ur ON r.id = ur.role_id
		JOIN 
				users u ON ur.user_id = u.id
		WHERE 
				u.id = $1;
	`
	ADD_ROLE_TO_USER_STMT      = `INSERT INTO users_roles (role_id, user_id) VALUES ($1, $2);`
	REMOVE_ROLE_FROM_USER_STMT = `DELETE FROM users_roles WHERE role_id = $1 AND user_id = $2`
)
