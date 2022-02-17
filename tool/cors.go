package tool

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Cors(r *chi.Mux) {
	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{frontendOrigin},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{""},
		//これを追加すると Cookie が取得できる
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}
