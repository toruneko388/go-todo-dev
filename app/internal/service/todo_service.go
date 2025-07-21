package service

import (
	"errors"

	"github.com/toruneko388/todoapp/internal/models"
	"github.com/toruneko388/todoapp/internal/repository"
)

// TodoServiceはTodoに関するビジネスロジックを定義するインターフェースです。
type TodoService interface {
	GetAll() ([]models.Todo, error)
	Create(title string) error
}

// todoServiceはTodoServiceインターフェースの具象実装です。
type todoService struct {
	repo repository.TodoRepository
}

// NewTodoServiceは新しいTodoServiceのインスタンスを生成します。
// 依存性注入により、具体的なリポジトリ実装を受け取ります。
func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

// GetAllは全てのTodoを取得するためのビジネスロジックを実装します。
// (現状はリポジトリを呼び出すだけですが、将来的にロジックを追加できます)
func (s *todoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

// Createは新しいTodoを作成するためのビジネスロジックを実装します。
func (s *todoService) Create(title string) error {
	// ここがビジネスロジック。ハンドラから移動してきました。
	if title == "" {
		return errors.New("title cannot be empty")
	}
	return s.repo.Insert(title)
}
