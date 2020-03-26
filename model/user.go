package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

// HasCreateNewUser handles new user creation.
type HasCreateNewUser interface {
	CreateNewUser(ctx context.Context, data *models.User) (*models.User, error)
}

// HasGetUser handles user retrieval.
type HasGetUser interface {
	GetUser(ctx context.Context, username string) (*models.User, error)
}

// UserDatastore holds db information.
type UserDatastore struct {
	*sql.DB
}

// CreateNewUser creates a new user with given username and password.
func (db *UserDatastore) CreateNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	hash, _ := hashPassword(user.Password)
	
	user.ID = ksuid.New().String()
	user.Password = hash

	if err := user.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// GetUser returns user by given email.
func (db *UserDatastore) GetUser(ctx context.Context, email string) (*models.User, error) {
	return models.Users(models.UserWhere.Email.EQ(email)).One(ctx, db)
}
