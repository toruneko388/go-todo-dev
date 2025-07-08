package models

import "time"

// Todo は 1 レコードを表す
type Todo struct {
	ID        int
	Title     string
	CreatedAt time.Time
}

// GetAllTodos は todos テーブルを SELECT して一覧を返す
func GetAllTodos() ([]Todo, error) {
	rows, err := DB.Query(`
		SELECT id, title, created_at
		FROM todos
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

// InsertTodo は 1 件 INSERT する
func InsertTodo(title string) error {
	_, err := DB.Exec(`INSERT INTO todos(title) VALUES(?)`, title)
	return err
}
