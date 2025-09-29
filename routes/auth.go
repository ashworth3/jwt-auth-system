package routes

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"jwt-auth-system/controllers"
	"jwt-auth-system/middleware"
	"gorm.io/gorm"
)

func AuthRoutes(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.Register(db, w, r)
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(db, w, r)
	})
	r.With(middleware.JWTAuthMiddleware).Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected content here"))
	})

	return r
}
