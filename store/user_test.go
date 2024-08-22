package store

import (
	"testing"

	"github.com/ankitg1706/zoopone/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=ankit password=password dbname=manage port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Suppress logging during testing
	})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}

	// Migrate the User model for testing
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("failed to auto-migrate User model: %v", err)
	}

	return db
}

func cleanupTestDB(db *gorm.DB, t *testing.T) {
	err := db.Exec("DROP TABLE IF EXISTS users").Error
	if err != nil {
		t.Fatalf("failed to drop table: %v", err)
	}
}

// TestCreateUser tests the CreateUser function.
func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db, t)

	store := Postgress{DB: db}

	user := &model.User{
		ID:        uuid.New(),
		FirstName: "Test User",
		Email:     "test@example.com",
	}

	err := store.CreateUser(user)
	assert.NoError(t, err, "Expected no error while creating user")

	// Check if the user was created
	var foundUser model.User
	err = db.First(&foundUser, "email = ?", user.Email).Error
	assert.NoError(t, err, "Expected no error while finding user")
	assert.Equal(t, user.FirstName, foundUser.FirstName, "Expected user name to match")
}

// TestGetUsers tests the GetUsers function.
func TestGetUsers(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db, t)

	store := Postgress{DB: db}

	// Create a user
	user := &model.User{
		ID:        uuid.New(),
		FirstName: "Test User",
		Email:     "test@example.com",
	}
	store.CreateUser(user)

	users, err := store.GetUsers()
	assert.NoError(t, err, "Expected no error while getting users")
	assert.Len(t, users, 1, "Expected to find one user")
	assert.Equal(t, user.Email, users[0].Email, "Expected user email to match")
}

// TestGetUser tests the GetUser function.
func TestGetUser(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db, t)

	store := Postgress{DB: db}

	// Create a user
	user := &model.User{
		ID:        uuid.New(),
		FirstName: "TestUser",
		Email:     "test@example.com",
	}
	store.CreateUser(user)

	// Test fetching the user
	foundUser, err := store.GetUser(user.ID)
	assert.NoError(t, err, "Expected no error while getting user")
	assert.Equal(t, user.Email, foundUser.Email, "Expected user email to match")
}

// TestDeleteUser tests the DeleteUser function.
func TestDeleteUser(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db, t)

	store := Postgress{DB: db}

	// Create a user
	user := &model.User{
		ID:        uuid.New(),
		FirstName: "Test User",
		Email:     "test@example.com",
	}
	store.CreateUser(user)

	// Delete the user
	err := store.DeleteUser(user.ID.String())
	assert.NoError(t, err, "Expected no error while deleting user")

	// Try to find the deleted user
	var foundUser model.User
	err = db.First(&foundUser, "email = ?", user.Email).Error
	assert.Error(t, err, "Expected error while finding deleted user")
	assert.Equal(t, gorm.ErrRecordNotFound, err, "Expected record not found error")
}

// TestUpdateUser tests the UpdateUser function.
func TestUpdateUser(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db, t)

	store := Postgress{DB: db}

	// Create a user
	user := &model.User{
		ID:        uuid.New(),
		FirstName: "TestUser",
		Email:     "test@example.com",
	}
	store.CreateUser(user)

	// Update the user's name
	user.FirstName = "Updated Test User"
	err := store.UpdateUser(user)
	assert.NoError(t, err, "Expected no error while updating user")

	// Check if the user was updated
	var updatedUser model.User
	err = db.First(&updatedUser, "email = ?", user.Email).Error
	assert.NoError(t, err, "Expected no error while finding updated user")
	assert.Equal(t, user.FirstName, updatedUser.FirstName, "Expected user name to match updated name")
}

// Additional test cases for SignUp, SignIn, GetUserByFilter can be written similarly.
