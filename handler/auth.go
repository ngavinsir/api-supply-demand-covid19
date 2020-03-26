package handler

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"golang.org/x/crypto/bcrypt"
)

// AuthResource holds user data store information.
type AuthResource struct {
	*model.UserDatastore
}

// UserCtxKey extracts user information from context.
var UserCtxKey = &contextKey{"User_id"}

var jwtAuth *jwtauth.JWTAuth

func init() {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic(errors.New("env JWT_SECRET not provided"))
	}
	jwtAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}

func (store *AuthResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", Login(store))
	r.Post("/register", Register(store))
	
	return r
}

// Register new user handler
func Register(repo interface {
	model.HasCreateNewUser
	model.HasGetUserByEmail
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &RegisterRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		_, err := repo.CreateNewUser(r.Context(), data.User)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		loginResponse, err := loginLogic(r.Context(), repo, data.User)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, loginResponse)
	}
}

// Login handler
func Login(repo interface {model.HasGetUserByEmail}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &LoginRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		loginResponse, err := loginLogic(r.Context(), repo, data.User)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, loginResponse)
	}
}

func loginLogic(ctx context.Context, repo model.HasGetUserByEmail, data *models.User) (*LoginResponse, error) {
	user, err := repo.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(data.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	jwtClaims := jwt.MapClaims{
		"user_id": user.ID,
	}
	jwtClaims["exp"] = jwtauth.ExpireIn(3 * time.Hour)

	_, tokenString, _ := jwtAuth.Encode(jwtClaims)
	
	loginResponse := &LoginResponse{
		Email: data.Email,
		Name:  data.Name,
		Role: data.Role,
		ContactPerson: data.ContactPerson.String,
		ContactNumber: data.ContactNumber.String,
		JWT: tokenString,
	}

	return loginResponse, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AuthMiddleware to handle request jwt token
func AuthMiddleware(next http.Handler) http.Handler {
	return jwtauth.Verifier(jwtAuth)(next)
}

// UserCtx middleware is used to extract user information from jwt.
func UserCtx(repo interface {model.HasGetUserByID}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, claims, err := jwtauth.FromContext(r.Context())
	
			if err != nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}
	
			if token == nil || !token.Valid {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			
			userID, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			user, err := repo.GetUserByID(r.Context(), userID)
			if err != nil {
				render.Render(w, r, ErrRender(err))
			}

			ctx := context.WithValue(r.Context(), UserCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RegisterRequest struct
type RegisterRequest struct {
	*models.User
}

// Bind RegisterRequest (Username, Password) [Required]
func (req *RegisterRequest) Bind(r *http.Request) error {
	if req.Email == "" || req.Password == "" || req.Name == "" || req.Role == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

// LoginRequest struct
type LoginRequest struct {
	*models.User
}

// Bind LoginRequest (Username, Password) [Required]
func (req *LoginRequest) Bind(r *http.Request) error {
	if req.Email == "" || req.Password == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

// LoginResponse struct
type LoginResponse struct {
	Email         string `boil:"email" json:"email"`
	Name          string `boil:"name" json:"name"`
	Role          string `boil:"role" json:"role"`
	ContactPerson string `boil:"contact_person" json:"contactPerson,omitempty"`
	ContactNumber string `boil:"contact_number" json:"contactNumber,omitempty"`
	JWT           string `json:"jwt"`
}
