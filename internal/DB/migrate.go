package db

import (
	logger "ArtAPI/util/logging"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	l := logger.GetLogger()
	l.Info().Msg("Checking for Database and migrating if missing")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///internal/DB/migration_scripts",
		"postgres",
		driver,
	)
	if err != nil {
		l.Debug().Err(err).Msg("")
		return err
	}

	if err:= m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up mgrations: %w", err)
		}
	}

	l.Info().Msg("successfully found or migrated database")
	return nil
}