package observable

import (
	"net/http"
	"strconv"

	"monita/storage/observable"

	"github.com/go-chi/chi"
)

// FromRequest returns Observable struct from Request
// returns `nil` if url param is not provided or some error
func FromRequest(r *http.Request) *observable.Observable {
	param := chi.URLParam(r, "observableID")

	if param == "" {
		return nil
	}

	id, err := strconv.ParseUint(param, 10, 0)

	if err != nil {
		return nil
	}

	return observable.GetByID(uint(id))
}
