package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

const (
	testEmail    = "TEST_EMAIL"
	testPassword = "TEST_PASSWORD"
)

func TestUser(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()

	t.Run("Create", testCreateUser(&UserDatastore{DB: db}))
}

func testCreateUser(repo *UserDatastore) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := repo.CreateNewUser(context.Background(), &models.User{
			Email:    testEmail,
			Password: testPassword,
			Name:     "TEST_NAME",
			Role:     "TEST_ROLE",
		})
		if err != nil {
			t.Error(err)
		}

		if user.ID == "" {
			t.Errorf("Want user id assigned, got %s", user.ID)
		}
		if user.Password == testPassword {
			t.Errorf("Want user password hashed, got %s", user.Password)
		}

		t.Run("Get by email", testGetUserByEmail(repo, user.Email))
		t.Run("Get by id", testGetUserByID(repo, user.ID))
	}
}

func testGetUserByEmail(repo *UserDatastore, email string) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := repo.GetUserByEmail(context.Background(), email)
		if err != nil {
			t.Error(err)
		}

		if got, want := user.Email, email; got != want {
			t.Errorf("Want user email %s, got %s", want, got)
		}
	}
}

func testGetUserByID(repo *UserDatastore, userID string) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := repo.GetUserByID(context.Background(), userID)
		if err != nil {
			t.Error(err)
		}

		if got, want := user.ID, userID; got != want {
			t.Errorf("Want user id %s, got %s", want, got)
		}
	}
}
