package handlers

import (
	"net/http"

	"monita/config"
	"monita/handlers/observable"
	"monita/handlers/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Listen runs listen and serve
func Listen() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", user.Login)

		r.Route("/users", func(r chi.Router) {
			r.Use(user.IsAdmin)

			r.Post("/", user.Create)

			r.Delete("/{userName}", user.Delete)
		})

		r.Route("/user", func(r chi.Router) {
			r.Use(user.Verifier)
			r.Use(user.IsAuthentificated)

			r.Get("/", user.Current)

			r.Patch("/email", user.ChangeEmail)
			r.Patch("/password", user.ChangePassword)
			r.Patch("/mute", user.MuteNotifications)
			r.Patch("/unmute", user.UnmuteNotifications)

			r.Route("/observables", func(r chi.Router) {
				r.Post("/", observable.Create)

				r.Route("/{observableID:[0-9]+}", func(r chi.Router) {
					r.Get("/", observable.LoadData)

					r.Post("/", observable.Handle)

					r.Patch("/mute", observable.MuteNotifications)
					r.Patch("/unmute", observable.UnmuteNotifications)
					r.Patch("/reorder", observable.Reorder)

					r.Delete("/", observable.Delete)
				})
			})
		})
	})

	http.ListenAndServe(config.Entrypoint, r)
}
