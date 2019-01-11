package engine

import "github.com/hvs-fasya/chusha/internal/models"

// DBInterface - stores common interface
type DBInterface interface {
	TabsGet(bool) ([]*models.Tab, error)
	TabsSet([]*models.Tab) error

	UserCheck(string, string) (*models.UserDB, error)
	UserCreate(*models.UserNewInput, string) error
	UserGetByName(string) (*models.UserDB, error)

	RoleGetByName(role string) (*models.RoleDB, error)
}
