package handlers

import (
	"net/http"

	"monita-backend/config"
	"monita-backend/handlers/observable"
	"monita-backend/handlers/user"

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

			r.Post("/mute", user.MuteNotifications)
			r.Post("/unmute", user.UnmuteNotifications)

			r.Patch("/email", user.ChangeEmail)
			r.Patch("/password", user.ChangePassword)

			r.Route("/observables", func(r chi.Router) {
				r.Post("/", observable.Create)

				r.Route("/{observableID}", func(r chi.Router) {
					r.Get("/", observable.LoadData)

					r.Post("/", observable.Handle)
					r.Post("/mute", observable.MuteNotifications)
					r.Post("/unmute", observable.UnmuteNotifications)
					r.Post("/reorder", observable.Reorder)

					r.Delete("/", observable.Delete)
				})
			})
		})
	})

	http.ListenAndServe(config.Entrypoint, r)
}
