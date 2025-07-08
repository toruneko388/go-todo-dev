package handlers

import (
	"html/template"
	"net/http"

	"todoapp/internal/models"
)

type TodoHandler struct {
	tmpl *template.Template
}

// NewTodoHandler はテンプレートをパースして返す
func NewTodoHandler() *TodoHandler {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	return &TodoHandler{tmpl: tmpl}
}

// ListTodos : GET /todos
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAllTodos()
	if err != nil {
		http.Error(w, "failed to fetch todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.tmpl.Execute(w, todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddTodo : POST /todos
func (h *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	if err := models.InsertTodo(title); err != nil {
		http.Error(w, "insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
