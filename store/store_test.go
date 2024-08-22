package store

import (
	"testing"

	"github.com/ankitg1706/zoopone/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestNewStore(t *testing.T) {
	// Set up the PostgreSQL database connection for testing
	dsn := "host=localhost user=ankit password=password dbname=manage port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Suppress logging during testing
	})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	tests := []struct {
		name          string
		prepare       func() error
		expectedError bool
		verify        func() bool
		cleanup       func() error
	}{
		{
			name: "New Store Success",
			prepare: func() error {
				// Ensure the database is empty before the test
				return db.Exec("DROP TABLE IF EXISTS users").Error
			},
			expectedError: false,
			verify: func() bool {
				// Check if the User table is created
				return db.Migrator().HasTable(&model.User{})
			},
			cleanup: func() error {
				// Clean up: Drop the table after the test
				return db.Exec("DROP TABLE IF EXISTS users").Error
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.prepare(); err != nil {
				t.Fatalf("failed to prepare test: %v", err)
			}

			store := &Postgress{DB: db}

			err := store.NewStore()
			if tt.expectedError {
				assert.Error(t, err, "Expected error while creating new store")
			} else {
				assert.NoError(t, err, "Expected no error while creating new store")
			}

			assert.True(t, tt.verify(), "Expected User table to be created")

			if err := tt.cleanup(); err != nil {
				t.Fatalf("failed to clean up after test: %v", err)
			}
		})
	}
}
