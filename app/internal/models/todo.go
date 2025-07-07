package models

// Todo は 1 件のタスクを表す
type Todo struct {
	ID    int
	Title string
}

var (
	todos  []Todo // メモリ上に保持
	nextID int    // ID 自動採番用
)

// GetAll は現在の Todo 一覧を返す
func GetAll() []Todo {
	return todos
}

// Add はタイトルを受け取り、新しい Todo を追加する
func Add(title string) {
	nextID++
	todos = append(todos, Todo{
		ID:    nextID,
		Title: title,
	})
}
