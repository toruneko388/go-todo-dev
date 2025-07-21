package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/toruneko388/todoapp/internal/service"
)

// TodoHandler はTodo関連のHTTPリクエストを処理します。
// テンプレートとリポジトリを保持します。
type TodoHandler struct {
	Tmpl    *template.Template
	Service service.TodoService // 具体的な実装ではなくインターフェースに依存
}

// テンプレートをパースして返す
func NewTodoHandler(svc service.TodoService) *TodoHandler {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	return &TodoHandler{
		Tmpl:    tmpl,
		Service: svc,
	}
}

// ListTodos : GET /todos
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetAll()
	if err != nil {
		log.Printf("Todoの取得に失敗しました: %v", err)
		http.Error(w, "failed to fetch todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// テンプレートにデータを渡して実行
	err = h.Tmpl.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"Todos": todos,
	})
	if err != nil {
		log.Printf("テンプレートの実行に失敗しました: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// AddTodo : POST /todos
func (h *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("フォームの解析に失敗しました: %v", err)
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")

	// バリデーションロジックはServiceレイヤーに移動したため、ここでは単純に呼び出すだけ
	if err := h.Service.Create(title); err != nil {
		log.Printf("Todoの追加に失敗しました: %v", err)
		http.Error(w, "insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
