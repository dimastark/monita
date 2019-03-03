package user

import (
	"net/http"

	"monita-backend/pkg/response"
	"monita-backend/pkg/validator"
	"monita-backend/storage/observable"
	"monita-backend/storage/user"

	"github.com/go-chi/render"
)

// Create

type createRequest struct {
	Name     string
	Password string
	Email    string
}

func (cr *createRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Name(cr.Name),
		validator.Password(cr.Password),
		validator.Email(cr.Email),
		validator.NameOrEmailNotExist(cr.Name, cr.Email),
	)
}

type createResponse struct {
	*user.User
}

func newCreateResponse(u *user.User) *createResponse {
	return &createResponse{u}
}

func (er *createResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)

	return nil
}

// Login

type loginRequest struct {
	Login    string
	Password string
}

func (lr *loginRequest) Bind(r *http.Request) error {
	return nil
}

type loginResponse struct {
	Token string `json:"token"`
}

func newLoginResponse(token string) *loginResponse {
	return &loginResponse{token}
}

func (lr *loginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// Current

type currentResponse struct {
	*user.User

	Observable []observable.Observable `json:"observables"`
}

func newCurrentResponse(u *user.User) *currentResponse {
	return &currentResponse{u, u.GetObservables()}
}

func (cr *currentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// ChangePassword

type changePasswordRequest struct {
	New string
}

func (cpr *changePasswordRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Password(cpr.New),
	)
}

// ChangeEmail

type changeEmailRequest struct {
	New string
}

func (cer *changeEmailRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Email(cer.New),
		validator.EmailNotExist(cer.New),
	)
}
