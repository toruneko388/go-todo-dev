package main

import (
	"log"
	"net/http"

	"github.com/toruneko388/todoapp/internal/database"
	"github.com/toruneko388/todoapp/internal/handlers"
	"github.com/toruneko388/todoapp/internal/repository"
	"github.com/toruneko388/todoapp/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := database.InitDB(database.GetDBPath())
	defer db.Close()

	// 1. リポジトリのインスタンスを作成
	todoRepo := repository.NewSQLiteTodoRepository(db)

	// 2. サービスレイヤーのインスタンスを作成
	//    リポジトリをサービスに注入する
	todoService := service.NewTodoService(todoRepo)

	// 3. ハンドラのインスタンスを作成 (注入するものを変更)
	//    サービスをハンドラに注入する
	todoHandler := handlers.NewTodoHandler(todoService)

	// ルーティングの設定
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
