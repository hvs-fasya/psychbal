package utils

import "golang.org/x/crypto/bcrypt"

//common app constants
const (
	ClientRoleName = "client"
	AdminRoleName  = "admin"
)

//HashAndSalt get salted hash from password string
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
