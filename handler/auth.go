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
	"github.com/joho/godotenv"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"golang.org/x/crypto/bcrypt"
)

// AuthResource holds user data store information.
type AuthResource struct {
	*model.UserDatastore
}

// Common ctx key.
var (
	UserIDCtxKey = &contextKey{"User_id"}
	UserCtxKey = &contextKey{"User"}
)

var jwtAuth *jwtauth.JWTAuth

func init() {
	godotenv.Load()

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
	r.With(AuthMiddleware).Post("/refresh", RefreshToken)
	
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
		Email: user.Email,
		Name:  user.Name,
		Role: user.Role,
		ContactPerson: user.ContactPerson.String,
		ContactNumber: user.ContactNumber.String,
		JWT: tokenString,
	}

	return loginResponse, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RefreshToken refreshes jwt token.
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		render.Render(w, r, ErrUnauthorized(err))
		return
	}

	if token == nil || !token.Valid {
		render.Render(w, r, ErrUnauthorized(errors.New("token is invalid")))
		return
	}

	claims["exp"] = jwtauth.ExpireIn(3 * time.Hour)
	_, tokenString, _ := jwtAuth.Encode(claims)

	render.PlainText(w, r, tokenString)
}

// AuthMiddleware to handle request jwt token
func AuthMiddleware(next http.Handler) http.Handler {
	return jwtauth.Verifier(jwtAuth)(extractUserID(next))
}

func extractUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			render.Render(w, r, ErrUnauthorized(err))
			return
		}

		if token == nil || !token.Valid || claims["user_id"] == nil {
			render.Render(w, r, ErrUnauthorized(errors.New("token is invalid")))
			return
		}

		ctx := context.WithValue(r.Context(), UserIDCtxKey, claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserCtx middleware is used to extract user information from jwt.
func UserCtx(repo interface {model.HasGetUserByID}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, _ := r.Context().Value(UserIDCtxKey).(string)

			user, err := repo.GetUserByID(r.Context(), userID)
			if err != nil {
				render.Render(w, r, ErrRender(err))
				return
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
	if req.User == nil || req.Email == "" || req.Password == "" || req.Name == "" || req.Role == "" {
		return ErrMissingReqFields
	}

	return nil
}

// LoginRequest struct
type LoginRequest struct {
	*models.User
}

// Bind LoginRequest (Username, Password) [Required]
func (req *LoginRequest) Bind(r *http.Request) error {
	if req.User == nil || req.Email == "" || req.Password == "" {
		return ErrMissingReqFields
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
