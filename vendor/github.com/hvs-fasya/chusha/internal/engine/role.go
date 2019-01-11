package engine

import "github.com/hvs-fasya/chusha/internal/models"

//RoleGetByName get role db record by role name
func (db *PgDB) RoleGetByName(roleName string) (*models.RoleDB, error) {
	var role = new(models.RoleDB)
	q := `SELECT id, role FROM roles WHERE role=$1 LIMIT 1`
	e := db.Conn.QueryRow(q, roleName).Scan(
		&role.ID,
		&role.Role,
	)
	if e != nil {
		return role, e
	}
	return role, nil
}
