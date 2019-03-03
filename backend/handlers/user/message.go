package user

import (
	"net/http"

	"monita/pkg/response"
	"monita/pkg/validator"
	"monita/storage/observable"
	"monita/storage/user"

	"github.com/go-chi/render"
)

// CreateUserRequest represents create user payload
type CreateUserRequest struct {
	Name     string
	Password string
	Email    string
}

// Bind implements render.Binder interface
func (cr *CreateUserRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Name(cr.Name),
		validator.Password(cr.Password),
		validator.Email(cr.Email),
		validator.NameOrEmailNotExist(cr.Name, cr.Email),
	)
}

// CreateUserResponse represents create user result
type CreateUserResponse struct {
	*user.User
}

func newCreateResponse(u *user.User) *CreateUserResponse {
	return &CreateUserResponse{u}
}

// Render implements render.Renderer interface
func (er *CreateUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)

	return nil
}

// LoginRequest represents login payload
type LoginRequest struct {
	Login    string
	Password string
}

// Bind implements render.Binder interface
func (lr *LoginRequest) Bind(r *http.Request) error {
	return nil
}

// LoginResponse represents login result
type LoginResponse struct {
	Token string `json:"token"`
}

func newLoginResponse(token string) *LoginResponse {
	return &LoginResponse{token}
}

// Render implements render.Renderer interface
func (lr *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// CurrentUserResponse represents current user result
type CurrentUserResponse struct {
	*user.User

	Observables []observable.Observable `json:"observables"`
}

func newCurrentResponse(u *user.User) *CurrentUserResponse {
	return &CurrentUserResponse{u, u.GetObservables()}
}

// Render implements render.Renderer interface
func (cr *CurrentUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// ChangePasswordRequest represents change password payload
type ChangePasswordRequest struct {
	New string
}

// Bind implements render.Binder interface
func (cpr *ChangePasswordRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Password(cpr.New),
	)
}

// ChangeEmailRequest represents change email payload
type ChangeEmailRequest struct {
	New string
}

// Bind implements render.Binder interface
func (cer *ChangeEmailRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Email(cer.New),
		validator.EmailNotExist(cer.New),
	)
}
