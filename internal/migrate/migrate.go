package migrate

import (
	"database/sql"

	"github.com/rubenv/sql-migrate"
)

// MigrationService service for rollup/down db migrations
type MigrationService struct {
	Conn    *sql.DB
	source  *migrate.MemoryMigrationSource
	dialect string
}

//NewMigrationService creates new instance of migration service
func NewMigrationService(c *sql.DB, dialect string) *MigrationService {
	srv := new(MigrationService)
	srv.Conn = c
	srv.source = getSource()
	srv.dialect = dialect
	return srv
}

//MigrateUP process database migrations
func (ms *MigrationService) MigrateUP() (int, error) {
	n, err := migrate.Exec(ms.Conn, ms.dialect, ms.source, migrate.Up)
	if err != nil {
		return 0, err
	}
	return n, nil
}

//MigrateDown process database migrations DOWN
func (ms *MigrationService) MigrateDown() (int, error) {
	n, err := migrate.Exec(ms.Conn, ms.dialect, ms.source, migrate.Down)
	if err != nil {
		return 0, err
	}
	return n, nil
}
