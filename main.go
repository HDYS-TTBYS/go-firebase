package main

import (
	"app/tool"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	tool.LoadEnv()
	tool.FirebaseInit()

	r := chi.NewRouter()

	//panic
	r.Use(middleware.Recoverer)

	//logging
	r.Use(middleware.Logger)

	//timeout
	r.Use(middleware.Timeout(time.Second * 10))

	//Httpレート制限
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	//Cors
	tool.Cors(r)

	//application/json; charset=UTF-8 のみ許可
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.ContentCharset("UTF-8"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":8000", r)
}
