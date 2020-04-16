package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/badoux/checkmail"
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
	*model.PasswordResetRequestDatastore
}

// Common ctx key.
var (
	UserIDCtxKey = &contextKey{"User_id"}
	UserCtxKey   = &contextKey{"User"}
	PageCtxKey   = &contextKey{"Pagination"}
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

func (res *AuthResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", Login(res))
	r.Post("/register", Register(res))
	r.Post("/reset", ResetPassword(res))
	r.Put("/reset/{requestID}/confirm", ConfirmPasswordReset(res))
	
	r.With(AuthMiddleware).Post("/refresh", RefreshToken)

	return r
}

func (res *AuthResource) cmd(args []string) {
	// admin {email} {password}
	loginResponse, err := registerLogic(
		context.Background(),
		res.UserDatastore,
		&models.User{
			Email:    args[0],
			Name:     args[0],
			Password: args[1],
			Role:     model.RoleAdmin,
		},
	)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(loginResponse)
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

		if data.User.Role == model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		loginResponse, err := registerLogic(r.Context(), repo, data.User)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, loginResponse)
	}
}

func registerLogic(ctx context.Context, repo interface {
	model.HasCreateNewUser
	model.HasGetUserByEmail
}, data *models.User) (*LoginResponse, error) {
	err := checkmail.ValidateFormat(data.Email)
	if err != nil {
		return nil, err
	}

	_, err = repo.CreateNewUser(ctx, data)
	if err != nil {
		return nil, err
	}

	loginResponse, err := loginLogic(ctx, repo, data)
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}

// Login handler
func Login(repo interface{ model.HasGetUserByEmail }) http.HandlerFunc {
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

	user.Password = ""
	loginResponse := &LoginResponse{
		User: user,
		JWT:  tokenString,
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

// ResetPassword creates new password reset request.
func ResetPassword(repo interface{ model.HasCreatePasswordResetRequest }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &ResetPasswordRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		id, err := repo.CreatePasswordResetRequest(r.Context(), data.Email)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		if err = SendPasswordResetConfirmationMail(data.Email, id); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
}

// ConfirmPasswordReset confirms password reset request.
func ConfirmPasswordReset(repo interface{ model.HasConfirmPasswordResetRequest }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := chi.URLParam(r, "requestID")
		if requestID == "" {
			render.Render(w, r, ErrInvalidRequest(ErrMissingReqFields))
			return
		}

		data := &ConfirmResetPasswordRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		if err := repo.ConfirmPasswordResetRequest(r.Context(), requestID, data.NewPassword); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
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
func UserCtx(repo interface{ model.HasGetUserByID }) func(http.Handler) http.Handler {
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
	User *models.User `json:"user"`
	JWT  string       `json:"jwt,omitempty"`
}

// ResetPasswordRequest struct
type ResetPasswordRequest struct {
	Email string `json:"email"`
}

// Bind ResetPasswordRequest (email) [Required]
func (req *ResetPasswordRequest) Bind(r *http.Request) error {
	if req.Email == "" {
		return ErrMissingReqFields
	}

	return nil
}

// ConfirmResetPasswordRequest struct
type ConfirmResetPasswordRequest struct {
	NewPassword string `json:"newPassword"`
}

// Bind ConfirmResetPasswordRequest (newPassword) [Required]
func (req *ConfirmResetPasswordRequest) Bind(r *http.Request) error {
	if req.NewPassword == "" {
		return ErrMissingReqFields
	}

	return nil
}
