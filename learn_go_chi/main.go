package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		//dd : chi.
		dta := chi.URLParam(r, "id")
		log.Println(dta)
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}
