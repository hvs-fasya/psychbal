package migrate

import (
	"github.com/rubenv/sql-migrate"
	//"github.com/hvs-fasya/psychbal/internal/utils"
)

func getSource() (migrations *migrate.MemoryMigrationSource) {
	//var h string
	//h, _ = utils.HashAndSalt([]byte("12345678"))
	migrations = &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{},
	}
	return
}
