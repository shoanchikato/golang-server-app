package stmt

const (
	GET_PERMISSIONS_BY_ROLE_ID_STMT = `
		SELECT 
			p.id,
			p.name
		FROM 
				permissions p
		JOIN 
				roles_permissions rp ON p.id = rp.permission_id
		JOIN 
				roles r ON rp.role_id = r.id
		WHERE 
				r.id = $1;
	`
	GET_PERMISSIONS_BY_USER_ID = `
		SELECT 
			p.id,
			p.name
		FROM 
			permissions p
		JOIN 
			roles_permissions rp ON p.id = rp.permission_id
		JOIN 
			roles r ON rp.role_id = r.id
		JOIN 
			users_roles ur ON r.id = ur.role_id
		WHERE 
			ur.user_id = $1;
	`
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
	ADD_PERMISSION_TO_ROLE_STMT      = `INSERT INTO roles_permissions (permission_id, role_id) VALUES ($1, $2);`
	ADD_ROLE_TO_USER_STMT            = `INSERT INTO users_roles (role_id, user_id) VALUES ($1, $2);`
	REMOVE_PERMISSION_FROM_ROLE_STMT = `DELETE FROM roles_permissions WHERE role_id = $1 AND permission_id = $2;`
	REMOVE_ROLE_FROM_USER_STMT       = `DELETE FROM users_roles WHERE role_id = $1 AND user_id = $2`
)
