package repository

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	USERS_TABLE                    = "users"
	REFRESH_TOKENS_TABLE           = "refresh_tokens"
	BUILDINGS_TABLE                = "buildings"
	BUILDING_UNITS_TABLE           = "building_units"
	USERS_AND_BUILDING_UNITS_TABLE = "users_and_building_units"
	ORGANIZATIONS_TABLE            = "organizations"
	ORGANIZATIONS_STAFF_TABLE      = "organizations_staff"
)

func NewPostgresDB(ctx context.Context, dsn string) (*gorm.DB, *sql.DB, error) {
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
