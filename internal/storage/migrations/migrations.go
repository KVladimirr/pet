package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migration(folderPath string, dbHost string) error {
    m, err := migrate.New(
        "file://" + folderPath,
        dbHost,
    )

    if err != nil {
		return fmt.Errorf("migration failed: %w", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
    }

	return nil
}