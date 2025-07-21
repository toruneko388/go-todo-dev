package database

import (
	"database/sql"
	"log"
	"path/filepath"

	// アンダースコアで始まるインポートは、パッケージの初期化関数のみを呼び出し、
	// 他の関数は直接使用しないことを示します。
	_ "modernc.org/sqlite" // SQLite ドライバ
	// _ "github.com/mattn/go-sqlite3" // SQLite ドライバ
	// これを使用する場合は、CGO_ENABLED=1 を有効にしてビルド・実行する
	// CGOはGoのコードからC言語のライブラリを使うための仕組み
	// sql.Open("sqlite3", path) として使用する
)

// InitDB はアプリ起動時に呼び出して DB を初期化・マイグレーションする
func InitDB(path string) *sql.DB {
	//db, err := sql.Open("sqlite3", path)
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}
	if err = db.Ping(); err != nil { // Ping は DB 接続が有効か確認する
		log.Fatalf("データベースへのPingに失敗しました: %v", err)
	}

	const createTable = `
	CREATE TABLE IF NOT EXISTS todos (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		title       TEXT    NOT NULL,
		created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := db.Exec(createTable); err != nil { // _は最初の戻り値を捨てる
		log.Fatalf("テーブルの作成に失敗しました: %v", err)
	}

	log.Println("データベースの準備が完了し、todosテーブルが利用可能です。")
	return db
}

// GetDBPath はdataディレクトリ下のパスを返す
func GetDBPath() string {
	return filepath.Join("data", "todo.db")
}
