package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func RunMigrations(dbURL, migrationsPath string) error {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func MigrateWithLock(ctx context.Context, db *sqlx.DB, dsn, migrationsPath string) error {
	const lockKey = 3141592653589793238
	var locked bool
	err := db.GetContext(ctx, &locked,
		"SELECT pg_try_advisory_lock($1)", lockKey)
	if err != nil {
		return fmt.Errorf("failed to check lock: %w", err)
	}

	if !locked {
		log.Println("Another migration is in progress, waiting...")
		// Wait for lock with timeout
		ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		defer cancel()

		_, err = db.ExecContext(ctx,
			"SELECT pg_advisory_lock($1)", lockKey)
		if err != nil {
			return fmt.Errorf("failed to acquire lock: %w", err)
		}
	}

	// release the lock when done
	defer db.Exec("SELECT pg_advisory_unlock($1)", lockKey)

	// Run migrations
	dbURL := db.DriverName() + "://" + dsn
	return RunMigrations(dbURL, migrationsPath)
}
