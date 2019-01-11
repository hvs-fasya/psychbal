package models

//User user structure as per DB
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	PswdHash string
	Role     *Role
}

//Role role structure as per DB
type Role struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

//LoginInput login input structure
type LoginInput struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}
