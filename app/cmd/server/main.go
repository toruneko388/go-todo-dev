package main

import (
	"log"
	"net/http"

	"github.com/toruneko388/todoapp/internal/database"
	"github.com/toruneko388/todoapp/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// --- DB 初期化
	database.InitDB("todo.db")

	r := chi.NewRouter()

	// リクエストログ
	r.Use(middleware.Logger)

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
