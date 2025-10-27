package postgres

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(ctx context.Context, dsn string) (*gorm.DB, *sql.DB, error) {
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, nil, err
	}
	return gdb, sqlDB, nil
}
