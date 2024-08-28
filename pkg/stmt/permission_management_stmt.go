package stmt

const (
	GET_PERMISSIONS_BY_ROLE_Id_STMT = `
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
	GET_PERMISSIONS_BY_USER_Id = `
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
	ADD_PERMISSION_TO_ROLE_STMT      = `INSERT INTO roles_permissions (permission_id, role_id) VALUES ($1, $2);`
	REMOVE_PERMISSION_FROM_ROLE_STMT = `DELETE FROM roles_permissions WHERE role_id = $1 AND permission_id = $2;`
)
