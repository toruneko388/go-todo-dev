package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"todoapp/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	h := handlers.NewTodoHandler()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	})

	r.Get("/todos", h.ListTodos)
	r.Post("/todos", h.AddTodo)

	log.Println("⇢ http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
