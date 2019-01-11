package engine

import (
	"golang.org/x/crypto/bcrypt"

	"database/sql"
	"github.com/hvs-fasya/chusha/internal/models"
	"github.com/hvs-fasya/chusha/internal/utils"
)

//UserCreate create new user in db
func (db *PgDB) UserCreate(user *models.UserNewInput, role string) error {
	var e error
	user.PswdHash, e = utils.HashAndSalt([]byte(user.Password))
	if e != nil {
		return e
	}
	q := `INSERT INTO users (email, phone, nickname, name, lastname, role_id, pswd_hash)
			VALUES ($1, $2, $3, $4, $5, (SELECT id FROM roles WHERE role=$6 LIMIT 1), $7)
			RETURNING id`
	e = db.Conn.QueryRow(q, user.Email, user.Phone, user.Nickname, user.Name, user.LastName, role, user.PswdHash).Scan(
		&user.ID,
	)
	if e != nil {
		return e
	}
	return e
}

//UserCheck check username/password pair exists
func (db *PgDB) UserCheck(login string, pwd string) (*models.UserDB, error) {
	user := new(models.UserDB)
	user.Role = new(models.RoleDB)
	q := `SELECT u.id, u.email, u.phone, u.nickname, u.name, u.lastname, u.pswd_hash,
			r.id, r.role
		FROM users u
		JOIN roles r ON r.id=u.role_id
		WHERE (nickname=$1 OR email=$1) LIMIT 1`
	err := db.Conn.QueryRow(q, login).Scan(
		&user.ID,
		&user.Email,
		&user.Phone,
		&user.Nickname,
		&user.Name,
		&user.LastName,
		&user.PswdHash,
		&user.Role.ID,
		&user.Role.Role,
	)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PswdHash), []byte(pwd))
	if err != nil {
		return user, sql.ErrNoRows
	}
	return user, nil //if err != nil => user not authorized
}

//UserGetByName get user object by nickname
func (db *PgDB) UserGetByName(nickname string) (*models.UserDB, error) {
	user := new(models.UserDB)
	user.Role = new(models.RoleDB)
	q := `SELECT u.id, u.email, u.phone, u.nickname, u.name, u.lastname,
				r.id, r.role
			FROM users u
			JOIN roles r ON u.role_id=r.id
			WHERE u.nickname=$1 LIMIT 1`
	err := db.Conn.QueryRow(q, nickname).Scan(
		&user.ID,
		&user.Email,
		&user.Phone,
		&user.Nickname,
		&user.Name,
		&user.LastName,
		&user.Role.ID,
		&user.Role.Role,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
