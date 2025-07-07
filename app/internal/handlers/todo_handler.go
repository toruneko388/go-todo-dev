package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"todoapp/internal/models"
)

// TodoHandler は Todo 関連のハンドラを束ねる
type TodoHandler struct {
	tmpl *template.Template
}

// NewTodoHandler はテンプレートをパースして返す
func NewTodoHandler() *TodoHandler {
	t := template.Must(template.ParseFiles("templates/index.html"))
	return &TodoHandler{tmpl: t}
}

// ListTodos は /todos GET
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.Execute(w, models.GetAll()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddTodo は /todos POST
func (h *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(r.FormValue("title"))
	if title != "" {
		models.Add(title)
	}
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
