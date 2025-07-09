package models

import (
	"database/sql"
	"time"
)

// Todo は 1 レコードを表す
type Todo struct {
	ID        int
	Title     string
	CreatedAt time.Time
}

// GetAllTodos は todos テーブルを SELECT して一覧を返す
func GetAllTodos(DB *sql.DB) ([]Todo, error) {
	rows, err := DB.Query(`
		SELECT id, title, created_at
		FROM todos
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // この関数（GetAllTodos）が終了する直前に rows.Close() を実行する

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt); err != nil { // Scan は SELECT した結果を t に格納していき、Goの変数に読み込むときのエラーをチェック
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

// InsertTodo は 1 件 INSERT する
func InsertTodo(DB *sql.DB, title string) error {
	_, err := DB.Exec(`INSERT INTO todos(title) VALUES(?)`, title)
	return err
}
