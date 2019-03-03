package response

import (
	"fmt"
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
	return &Generic{200, http.StatusText(200)}
}

// BadRequest for 400 status
func BadRequest(err error) *Generic {
	return &Generic{400, err.Error()}
}

// Unauthorized for 401 status
func Unauthorized() *Generic {
	return &Generic{401, http.StatusText(401)}
}

// NotFound for 404 status
func NotFound(entity string) *Generic {
	return &Generic{404, fmt.Sprint(entity, "not found")}
}

// ServerError for 500 status
func ServerError() *Generic {
	return &Generic{500, http.StatusText(500)}
}

// Render writes status to response
func (gr *Generic) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, gr.Status)

	return nil
}
