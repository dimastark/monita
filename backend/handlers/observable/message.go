package observable

import (
	"fmt"
	"net/http"

	"monita/pkg/response"
	"monita/pkg/validator"
	"monita/storage/observable"

	"github.com/go-chi/render"
)

// CreateObservableRequest represents create observable payload
type CreateObservableRequest struct {
	Name        string
	Periodicity string
	URL         string
	Selector    string
}

// Bind implements render.Binder interface
func (cr *CreateObservableRequest) Bind(r *http.Request) error {
	return response.Validate(
		validator.Name(cr.Name),
		validator.Periodicity(cr.Periodicity),
		validator.URL(cr.URL),
		validator.Selector(cr.Selector),
	)
}

// CreateObservableResponse represents create observable result
type CreateObservableResponse struct {
	*observable.Observable
}

func newCreateResponse(o *observable.Observable) *CreateObservableResponse {
	return &CreateObservableResponse{o}
}

// Render implements render.Renderer interface
func (cr *CreateObservableResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)

	return nil
}

// LoadDataResponse represents load data result
type LoadDataResponse struct {
	*observable.Observable

	NewData string `json:"newData"`
}

func newLoadDataResponse(o *observable.Observable, newData string) *LoadDataResponse {
	return &LoadDataResponse{o, newData}
}

// Render implements render.Renderer interface
func (ldr *LoadDataResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// Print prints all LoadDataResponse data
func (ldr *LoadDataResponse) Print() {
	fmt.Println("ID:      ", ldr.ID)
	fmt.Println("Name:    ", ldr.Name)
	fmt.Println("Old data:", ldr.LastData)
	fmt.Println("New data:", ldr.NewData)
}

// HandleResponse represents handle result
type HandleResponse struct {
	*observable.Observable
}

func newHandleResponse(o *observable.Observable) *HandleResponse {
	return &HandleResponse{o}
}

// Render implements render.Renderer interface
func (hr *HandleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)

	return nil
}

// ReorderRequest represents reorder payload
type ReorderRequest struct {
	New int
}

// Bind implements render.Binder interface
func (rr *ReorderRequest) Bind(r *http.Request) error {
	return nil
}
