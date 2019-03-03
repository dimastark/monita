package response

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// Validation response with array of errors
type Validation struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}

func (v *Validation) Error() string {
	return strings.Join(v.Errors, "\n")
}

// Render writes status to response
func (v *Validation) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, v.Status)

	return nil
}

// Validate returns Validation response
func Validate(errors ...error) error {
	var errStrings []string

	for _, err := range errors {
		if err != nil {
			errStrings = append(errStrings, err.Error())
		}
	}

	if len(errStrings) == 0 {
		return nil
	}

	return &Validation{400, errStrings}
}

// ValidationError for array of errors
func ValidationError(err error) render.Renderer {
	if v, ok := err.(*Validation); ok {
		return v
	}

	return BadRequest(err)
}
