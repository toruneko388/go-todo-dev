package main

import (
	"log"
	"net/http"

	"github.com/toruneko388/todoapp/internal/database"
	"github.com/toruneko388/todoapp/internal/handlers"
	"github.com/toruneko388/todoapp/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// 1. データベースの初期化
	db := database.InitDB(database.GetDBPath())

	// 2. 依存関係の注入 (Dependency Injection)
	// リポジトリを初期化し、データベース接続を渡す
	todoRepo := repository.NewTodoRepository(db)
	// ハンドラを初期化し、リポジトリを渡す
	todoHandler := handlers.NewTodoHandler(todoRepo)

	// 3. ルーティングの設定
	r := chi.NewRouter()

	// リクエストログ
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	})

	r.Get("/todos", todoHandler.ListTodos)
	r.Post("/todos", todoHandler.AddTodo)

	// 4. サーバーの起動
	port := ":8080"
	log.Printf("サーバーを http://localhost%s で起動します", port)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
