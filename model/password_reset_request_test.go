package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

func TestPasswordResetRequest(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()
	
	userDatastore := &UserDatastore{DB: db}
	user, err := userDatastore.CreateNewUser(context.Background(), &models.User{
		Email:    testUserEmail,
		Name:     testUserName,
		Role:     testUserRole,
		Password: testUserPassword,
	})
	if err != nil {
		panic(err)
	}

	t.Run("Create", testCreatePasswordResetRequest(&PasswordResetRequestDatastore{DB: db}, user))
}

func testCreatePasswordResetRequest(repo *PasswordResetRequestDatastore, user *models.User) func(t *testing.T) {
	return func(t *testing.T) {
		id, err := repo.CreatePasswordResetRequest(context.Background(), user.ID, "NEW_PASSWORD")
		if err != nil {
			t.Error(err)
		}

		request, err := models.PasswordResetRequests().All(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		if len(request) == 0 {
			t.Error("password reset request hasn't been created")
		}

		if request[0].NewPassword == "NEW_PASSWORD" {
			t.Error("Want password reset request new password hashed, got not hashed")
		}

		t.Run("Confirm", testConfirmPasswordReset(repo, id, user))
	}
}

func testConfirmPasswordReset(repo *PasswordResetRequestDatastore, requestID string, user *models.User) func(t *testing.T) {
	return func(t *testing.T) {
		err := repo.ConfirmPasswordResetRequest(context.Background(), user.ID, requestID)
		if err != nil {
			t.Error(err)
		}

		requestCount, err := models.PasswordResetRequests().Count(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		if requestCount > 0 {
			t.Error("Want request deleted, got request not deleted")
		}

		oldPassword := user.Password

		if err = user.Reload(context.Background(), repo.DB); err != nil {
			t.Error(err)
		}
		
		if user.Password == oldPassword {
			t.Error("Want user password changed, got old password")
		}
	}
}