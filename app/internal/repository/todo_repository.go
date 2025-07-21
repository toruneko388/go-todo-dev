package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/toruneko388/todoapp/internal/models"
)

// TodoRepository はTodoリポジトリのインターフェースです。
// これにより、ハンドラは具体的な実装に依存せず、テストが容易になります。
type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	Insert(title string) error
}

// SQLiteRepository はtodosテーブルへのデータアクセスを担当する構造体です。
type SQLiteRepository struct {
	DB *sql.DB
}

// NewTodoRepository は新しいSQLiteRepositoryのインスタンスを生成します。
// main.goで生成されたデータベース接続を受け取ります。
func NewSQLiteTodoRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{DB: db}
}

// GetAllはデータベースからすべてのTodoを取得します。
// 作成日時の降順（新しいものが上）でソートして返します。
func (r *SQLiteRepository) GetAll() ([]models.Todo, error) {
	// todosテーブルから全件取得するクエリを実行
	rows, err := r.DB.Query("SELECT id, title, created_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		log.Printf("クエリの実行に失敗しました: %v", err)
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	// 取得した各行をループ処理
	for rows.Next() {
		var todo models.Todo
		// 行のデータをTodo構造体にスキャン
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.CreatedAt); err != nil {
			log.Printf("行のスキャンに失敗しました: %v", err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Createは新しいTodoをデータベースに保存します。
func (r *SQLiteRepository) Insert(title string) error {
	// INSERT文を準備
	stmt, err := r.DB.Prepare("INSERT INTO todos(title, created_at) VALUES(?, ?)")
	if err != nil {
		log.Printf("INSERT文の準備に失敗しました: %v", err)
		return err
	}
	defer stmt.Close()

	// 準備した文を実行
	_, err = stmt.Exec(title, time.Now())
	if err != nil {
		log.Printf("Todoの作成に失敗しました: %v", err)
		return err
	}
	return nil
}
