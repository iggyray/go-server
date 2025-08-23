package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	router "github.com/iggyray/go-server/internal/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	r.Mount("/posts", router.PostRouter())

	log.Printf("Listening on http://localhost:%d\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", r))
}
