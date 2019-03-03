package observable

import (
	"net/http"

	"monita-backend/pkg/response"
	"monita-backend/pkg/validator"
	"monita-backend/storage/observable"

	"github.com/go-chi/render"
)

// Create

type createRequest struct {
	Name        string
	Periodicity string
	URL         string
	Selector    string
}

func (cr *createRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Name(cr.Name),
		validator.Periodicity(cr.Periodicity),
		validator.URL(cr.URL),
		validator.Selector(cr.Selector),
	)
}

type createResponse struct {
	*observable.Observable
}

func newCreateResponse(o *observable.Observable) *createResponse {
	return &createResponse{o}
}

func (cr *createResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)

	return nil
}

// LoadData

type loadDataResponse struct {
	*observable.Observable

	NewData string `json:"newData"`
}

func newLoadDataResponse(o *observable.Observable, newData string) *loadDataResponse {
	return &loadDataResponse{o, newData}
}

func (ldr *loadDataResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// Handle

type handleResponse struct {
	*observable.Observable
}

func newHandleResponse(o *observable.Observable) *handleResponse {
	return &handleResponse{o}
}

func (hr *handleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// Reorder

type reorderRequest struct {
	New int
}

func (rr *reorderRequest) Bind(r *http.Request) error {
	return nil
}

type reorderResponse struct {
	*observable.Observable
}

func newReorderResponse(o *observable.Observable) *reorderResponse {
	return &reorderResponse{o}
}

func (rr *reorderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}
