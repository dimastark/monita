package user

import (
	"net/http"

	"monita/config"
	"monita/pkg/response"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

var tokenAuth = jwtauth.New(
	"HS256",
	[]byte(config.JWTSecret),
	[]byte(config.JWTSecret),
)

// Verifier search for a jwt token in a http request
// Order: jwt query parameter, Authorization header, jwt cookie
func Verifier(next http.Handler) http.Handler {
	return jwtauth.Verifier(tokenAuth)(next)
}

// IsAuthentificated checks token authorization
func IsAuthentificated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := FromContext(r.Context())

		if u == nil {
			render.Render(w, r, response.Unauthorized())

			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsAdmin checks admin secret authorization
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ADMIN_SECRET")

		if err != nil || cookie.Value != config.AdminSecret {
			render.Render(w, r, response.Unauthorized())

			return
		}

		next.ServeHTTP(w, r)
	})
}
