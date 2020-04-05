package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

// HasCreatePasswordResetRequest handles reset password request creation.
type HasCreatePasswordResetRequest interface {
	CreatePasswordResetRequest(ctx context.Context, userID string, newPassword string) (string, error)
}

// HasConfirmPasswordResetRequest handles reset password request confirmation.
type HasConfirmPasswordResetRequest interface {
	ConfirmPasswordResetRequest(ctx context.Context, userID string, requestID string) error
}

// PasswordResetRequestDatastore holds db information
type PasswordResetRequestDatastore struct {
	*sql.DB
}

// CreatePasswordResetRequest creates new reset password request and deletes past request with the same user id.
func (db *PasswordResetRequestDatastore) CreatePasswordResetRequest(ctx context.Context, userID string, newPassword string) (string, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	pastRequest, err := models.PasswordResetRequests(
		models.PasswordResetRequestWhere.UserID.EQ(userID),
	).All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if len(pastRequest) > 0 {
		_, err = pastRequest[0].Delete(ctx, tx)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	requestID := ksuid.New().String()
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	passwordResetRequest := &models.PasswordResetRequest{
		ID: requestID,
		Date: time.Now(),
		NewPassword: string(hashedPassword),
		UserID: userID,
	}

	if err := passwordResetRequest.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}

	return requestID, nil
}

// ConfirmPasswordResetRequest confirms password reset request.
func (db *PasswordResetRequestDatastore) ConfirmPasswordResetRequest(ctx context.Context, userID string, requestID string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	request, err := models.FindPasswordResetRequest(ctx, tx, requestID)
	if err != nil {
		return err
	}

	if request.UserID != userID {
		return errors.New("invalid user id")
	}

	if request.Date.Before(time.Now().Add(-3 * time.Hour)) {
		return errors.New("password reset request has expired")
	}

	user, err := models.FindUser(ctx, tx, request.UserID)
	if err != nil {
		return err
	}

	user.Password = request.NewPassword
	if _, err = user.Update(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}
	
	if _, err = request.Delete(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}