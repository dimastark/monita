package user

import (
	"log"
	"net/http"

	"monita-backend/pkg/response"
	"monita-backend/storage/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// Create returns new User
func Create(w http.ResponseWriter, r *http.Request) {
	data := &createRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ValidationError(err))

		return
	}

	u, err := user.Create(data.Name, data.Password, data.Email)

	if err != nil {
		log.Println(err)

		render.Render(w, r, response.ServerError())

		return
	}

	render.Render(w, r, newCreateResponse(u))
}

// Delete removes User by name
func Delete(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "userName")

	if user.GetByName(name) == nil {
		render.Render(w, r, response.NotFound("user"))
	}

	user.DeleteByName(name)

	render.Render(w, r, response.OK())
}

// Login returns token or 401
func Login(w http.ResponseWriter, r *http.Request) {
	data := &loginRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ValidationError(err))

		return
	}

	u := user.Login(data.Login, data.Password)

	if u == nil {
		render.Render(w, r, response.Unauthorized())

		return
	}

	_, token, _ := tokenAuth.Encode(jwtauth.Claims{"id": u.ID})

	render.Render(w, r, newLoginResponse(token))
}

// Current returns current User data
func Current(w http.ResponseWriter, r *http.Request) {
	u := FromContext(r.Context())

	render.Render(w, r, newCurrentResponse(u))
}

// ChangePassword changes authenticated User password handler
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	u := FromContext(r.Context())

	data := &changePasswordRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ValidationError(err))

		return
	}

	u.ChangePassword(data.New)

	render.Render(w, r, response.OK())
}

// ChangeEmail changes authenticated User email
func ChangeEmail(w http.ResponseWriter, r *http.Request) {
	u := FromContext(r.Context())

	data := &changeEmailRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ValidationError(err))

		return
	}

	u.ChangeEmail(data.New)

	render.Render(w, r, response.OK())
}

// MuteNotifications mute all notifications
func MuteNotifications(w http.ResponseWriter, r *http.Request) {
	u := FromContext(r.Context())

	u.MuteNotifications()

	render.Render(w, r, response.OK())
}

// UnmuteNotifications unmute all notifications
func UnmuteNotifications(w http.ResponseWriter, r *http.Request) {
	u := FromContext(r.Context())

	u.UnmuteNotifications()

	render.Render(w, r, response.OK())
}
