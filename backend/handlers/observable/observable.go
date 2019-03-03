package observable

import (
	"log"
	"net/http"

	"monita/handlers/user"
	"monita/pkg/response"
	"monita/storage/observable"

	"github.com/go-chi/render"
)

// Create creates new Observable
func Create(w http.ResponseWriter, r *http.Request) {
	u := user.FromContext(r.Context())

	data := &CreateObservableRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ValidationError(err))

		return
	}

	o, err := observable.Create(observable.CreatePayload{
		Name:        data.Name,
		Periodicity: data.Periodicity,
		URL:         data.URL,
		Selector:    data.Selector,
		UserID:      u.ID,
	})

	if err != nil {
		log.Println(err)

		render.Render(w, r, response.ServerError())

		return
	}

	render.Render(w, r, newCreateResponse(o))
}

// Delete deletes Observable by id
func Delete(w http.ResponseWriter, r *http.Request) {
	o := FromRequest(r)
	u := user.FromContext(r.Context())

	if o == nil || o.UserID != u.ID {
		render.Render(w, r, response.NotFound("observable"))

		return
	}

	observable.DeleteByID(o.ID)

	render.Render(w, r, response.OK())
}

// LoadData returns Observable with new data
func LoadData(w http.ResponseWriter, r *http.Request) {
	o := FromRequest(r)
	u := user.FromContext(r.Context())

	if o == nil || o.UserID != u.ID {
		render.Render(w, r, response.NotFound("observable"))

		return
	}

	newData, err := o.LoadData()

	if err != nil {
		log.Println(err)

		render.Render(w, r, response.ServerError())

		return
	}

	render.Render(w, r, newLoadDataResponse(o, newData))
}

// Handle returns handled Observable
func Handle(w http.ResponseWriter, r *http.Request) {
	o := FromRequest(r)
	u := user.FromContext(r.Context())

	if o == nil || o.UserID != u.ID {
		render.Render(w, r, response.NotFound("observable"))

		return
	}

	err := o.Handle()

	if err != nil {
		log.Println(err)

		render.Render(w, r, response.ServerError())

		return
	}

	render.Render(w, r, newHandleResponse(o))
}

// MuteNotifications mute notifications for Observable
func MuteNotifications(w http.ResponseWriter, r *http.Request) {
	o := FromRequest(r)
	u := user.FromContext(r.Context())

	if o == nil || o.UserID != u.ID {
		render.Render(w, r, response.NotFound("observable"))

		return
	}

	o.MuteNotifications()

	render.Render(w, r, response.OK())
}

// UnmuteNotifications unmute notifications for Observable
func UnmuteNotifications(w http.ResponseWriter, r *http.Request) {
	o := FromRequest(r)
	u := user.FromContext(r.Context())

	if o == nil || o.UserID != u.ID {
		render.Render(w, r, response.NotFound("observable"))

		return
	}

	o.UnmuteNotifications()

	render.Render(w, r, response.OK())
}

// Reorder change Observable order
func Reorder(w http.ResponseWriter, r *http.Request) {
	o := FromRequest(r)
	u := user.FromContext(r.Context())

	if o == nil || o.UserID != u.ID {
		render.Render(w, r, response.NotFound("observable"))

		return
	}

	data := &ReorderRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ValidationError(err))

		return
	}

	o.Reorder(data.New)

	render.Render(w, r, response.OK())
}
