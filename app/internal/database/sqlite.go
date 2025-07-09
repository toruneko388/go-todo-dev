package database

import (
	"database/sql"
	"log"

	// アンダースコアで始まるインポートは、パッケージの初期化関数のみを呼び出し、
	// 他の関数は直接使用しないことを示します。
	_ "modernc.org/sqlite" // SQLite ドライバ
	// _ "github.com/mattn/go-sqlite3" // SQLite ドライバ
	// これを使用する場合は、CGO_ENABLED=1 を有効にしてビルド・実行する
	// CGOはGoのコードからC言語のライブラリを使うための仕組み
	// sql.Open("sqlite3", path) として使用する
)

// DB はパッケージ全体で共有する *sql.DB インスタンス
var DB *sql.DB

// InitDB はアプリ起動時に呼び出して DB を初期化・マイグレーションする
func InitDB(path string) {
	var err error
	//DB, err = sql.Open("sqlite3", path)
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	if err = DB.Ping(); err != nil { // Ping は DB 接続が有効か確認する
		log.Fatalf("ping database: %v", err)
	}

	const createTable = `
	CREATE TABLE IF NOT EXISTS todos (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		title       TEXT    NOT NULL,
		created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := DB.Exec(createTable); err != nil { // _は最初の戻り値を捨てる
		log.Fatalf("create table: %v", err)
	}
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("database not initialized")
	}
	return DB
}
