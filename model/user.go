package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

// User role enum
const (
	RoleDonator = "DONATOR"
	RoleApplicant = "APPLICANT"
	RoleAdmin = "ADMIN"
)

// HasCreateNewUser handles new user creation.
type HasCreateNewUser interface {
	CreateNewUser(ctx context.Context, user *models.User) (*models.User, error)
}

// HasGetUserByEmail handles user retrieval by email given.
type HasGetUserByEmail interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

// HasGetUserByID handles user retrieval by user id given.
type HasGetUserByID interface {
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}

// UserDatastore holds db information.
type UserDatastore struct {
	*sql.DB
}

// CreateNewUser creates a new user with given username and password.
func (db *UserDatastore) CreateNewUser(ctx context.Context, data *models.User) (*models.User, error) {
	hash, _ := hashPassword(data.Password)

	user := &models.User{
		ID:            ksuid.New().String(),
		Password:      hash,
		Role:          data.Role,
		ContactNumber: data.ContactNumber,
		ContactPerson: data.ContactPerson,
		Email:         data.Email,
		Name:          data.Name,
	}

	if err := user.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// GetUserByEmail returns user by given email.
func (db *UserDatastore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return models.Users(models.UserWhere.Email.EQ(email)).One(ctx, db)
}

// GetUserByID returns user by given user id.
func (db *UserDatastore) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	return models.Users(models.UserWhere.ID.EQ(userID)).One(ctx, db)
}
