package models

import (
	"time"
)

// Todo は 1 レコードを表す
type Todo struct {
	ID        int
	Title     string
	CreatedAt time.Time
}
