package response

import (
	"net/http"

	"github.com/go-chi/render"
)

// Generic response with status and message
type Generic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// OK for 200 status
func OK() *Generic {
	return &Generic{http.StatusOK, http.StatusText(http.StatusOK)}
}

// BadRequest for 400 status
func BadRequest(err error) *Generic {
	return &Generic{http.StatusBadRequest, err.Error()}
}

// Unauthorized for 401 status
func Unauthorized() *Generic {
	return &Generic{http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)}
}

// NotFound for 404 status
func NotFound(entity string) *Generic {
	return &Generic{http.StatusNotFound, entity + " not found"}
}

// ServerError for 500 status
func ServerError() *Generic {
	return &Generic{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)}
}

// Render writes status to response
func (gr *Generic) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, gr.Status)

	return nil
}
