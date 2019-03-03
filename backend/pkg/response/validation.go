package response

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// Validation response with array of errors
type Validation struct {
	Status int     `json:"status"`
	Errors []error `json:"errors"`
}

func (v *Validation) Error() string {
	var errStrings []string

	for _, err := range v.Errors {
		errStrings = append(errStrings, err.Error())
	}

	return strings.Join(errStrings, "\n")
}

// Render writes status to response
func (v *Validation) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, v.Status)

	return nil
}

// Validate returns Validation response
func Validate(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return &Validation{400, errors}
		}
	}

	return nil
}

// ValidationError for array of errors
func ValidationError(err error) render.Renderer {
	if v, ok := err.(*Validation); ok {
		return v
	}

	return BadRequest(err)
}
