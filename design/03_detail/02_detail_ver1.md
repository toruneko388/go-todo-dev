以下の内容のみを前提として、コードを作成してください

# 現時点の実装状況

## 1. 概要

### 1.1. アプリケーションの目的
Webブラウザを通じてTodo（タスク）の閲覧・追加ができるシンプルなTodo管理アプリケーションを開発する。

### 1.2. 技術スタック
- **言語**: Go
- **Webフレームワーク/ルーティング**: chi (`github.com/go-chi/chi/v5`)
- **データベース**: SQLite 3
- **DBドライバ**: go-sqlite3 (`github.com/mattn/go-sqlite3`)
- **テンプレートエンジン**: html/template（Go標準）
- **Goモジュール名**: `todoapp`

## 2. 機能要件

| No. | 機能名 | 概要 | エンドポイント | HTTPメソッド |
|:---:|:---|:---|:---|:---:|
| 1 | Todo一覧表示 | データベースに保存されている全てのTodoを作成日時の降順で表示する。 | `/todos` | GET |
| 2 | Todo追加 | フォームから送信されたタイトルのTodoをデータベースに追加する。 | `/todos` | POST |
| 3 | ルートリダイレクト | ルートパス(`/`)へのアクセスをTodo一覧表示(`/todos`)へリダイレクトする。 | `/` | GET |

## 3. アーキテクチャ設計

### 3.1. 設計方針
依存関係逆転の原則（DIP）に基づき、コンポーネントは具象ではなくインターフェースに依存する設計とする。これにより、データベースなどの下位モジュールの差し替えを容易にし、テスト容易性を高める。

### 3.2. ディレクトリ構成
```
app/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go       # エントリーポイント、依存関係の注入、サーバー起動
├── internal/
│   ├── database/
│   │   └── database.go   # データベース接続と初期化
│   ├── handlers/
│   │   └── todo_handler.go # HTTPリクエストの処理
│   ├── models/
│   │   └── todo.go       # データ構造の定義
│   └── repository/
│       └── todo_repository.go # データベース操作の抽象化と実装
└── templates/
    └── index.html          # HTMLテンプレート
```

## 4. データベース設計

### 4.1. データベース
- **種類**: SQLite
- **DBファイル名**: `todo.db` （アプリケーションルートに生成）

### 4.2. テーブル定義
- **テーブル名**: `todos`
- **テーブル作成**: アプリケーション起動時に、以下のSQL文を用いて自動でテーブルを作成する (`IF NOT EXISTS` を使用)。

  ```sql
  CREATE TABLE IF NOT EXISTS todos (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      created_at DATETIME NOT NULL
  );
  ```

| カラム名 | データ型 | 制約 | 説明 |
|:---|:---|:---|:---|
| `id` | `INTEGER` | `PRIMARY KEY`, `AUTOINCREMENT` | 一意な識別子（自動採番） |
| `title` | `TEXT` | `NOT NULL` | Todoのタイトル |
| `created_at` | `DATETIME` | `NOT NULL` | 作成日時 |

## 5. API（エンドポイント）詳細設計

### 5.1. `GET /todos`
- **機能**: Todo一覧表示
- **処理フロー**:
  1. データベースから全Todoを`created_at`の降順で取得する。
  2. 取得したTodoスライスをHTMLテンプレートに渡す。
  3. レンダリングされたHTMLをレスポンスとして返す。
- **成功レスポンス**:
  - **ステータスコード**: `200 OK`
  - **Content-Type**: `text/html; charset=utf-8`
  - **ボディ**: Todo一覧が描画されたHTML
- **失敗レスポンス**:
  - **ステータスコード**: `500 Internal Server Error`
  - **ボディ**: エラーメッセージ文字列

### 5.2. `POST /todos`
- **機能**: Todo追加
- **リクエスト**:
  - **Content-Type**: `application/x-www-form-urlencoded`
  - **フォームパラメータ**:
    - `title` (string, 必須): 追加するTodoのタイトル。
- **処理フロー**:
  1. リクエストフォームから`title`パラメータを取得する。
  2. `title`が空文字列の場合、何もせず`/todos`にリダイレクトする。
  3. `title`と現在時刻をデータベースの`todos`テーブルにINSERTする。
  4. 処理成功後、`/todos`にリダイレクトする。
- **成功レスポンス**:
  - **ステータスコード**: `302 Found`
  - **Locationヘッダ**: `/todos`
- **失敗レスポンス**:
  - **ステータスコード**: `500 Internal Server Error`
  - **ボディ**: エラーメッセージ文字列

## 6. コンポーネント詳細設計

### 6.1. `models` パッケージ
- **ファイル**: `internal/models/todo.go`
- **構造体**: `Todo`
  ```go
  package models
  
  import "time"
  
  type Todo struct {
      ID        int
      Title     string
      CreatedAt time.Time
  }
  ```

### 6.2. `repository` パッケージ
- **ファイル**: `internal/repository/todo_repository.go`
- **インターフェース**: `TodoRepository`
  - **責務**: Todoデータストアへのアクセス方法を定義する。
  - **メソッド**:
    - `GetAll() ([]models.Todo, error)`: 全てのTodoを取得する。
    - `Insert(title string) error`: タイトルを指定してTodoを1件作成する。

- **構造体**: `SQLiteTodoRepository`
  - **責務**: `TodoRepository`インターフェースをSQLiteで実装する。
  - **フィールド**: `DB (*sql.DB)`
  - **メソッド実装**:
    - `GetAll()`: `SELECT id, title, created_at FROM todos ORDER BY created_at DESC` を実行し、結果を`[]models.Todo`にマッピングして返す。
    - `Insert(title string)`: `INSERT INTO todos (title, created_at) VALUES (?, ?)` を実行する。プレースホルダには引数の`title`と現在時刻を渡す。

### 6.3. `handlers` パッケージ
- **ファイル**: `internal/handlers/todo_handler.go`
- **構造体**: `TodoHandler`
  - **責務**: Todo関連のHTTPリクエストを処理し、レスポンスを生成する。
  - **フィールド**:
    - `Tmpl (*template.Template)`: HTMLテンプレート
    - `Repo (repository.TodoRepository)`: リポジトリのインターフェース
- **関数**: `NewTodoHandler(repo repository.TodoRepository) *TodoHandler`
  - `templates/*.html`を`template.ParseGlob`でパースし、引数のリポジトリと共に`TodoHandler`のインスタンスを生成して返す。
- **メソッド**:
  - `ListTodos(w http.ResponseWriter, r *http.Request)`:
    1. `h.Repo.GetAll()`を呼び出す。
    2. 取得したTodoデータを`map[string]interface{}{"Todos": todos}`の形で`index.html`テンプレートに渡して実行する。
  - `AddTodo(w http.ResponseWriter, r *http.Request)`:
    1. `r.ParseForm()`と`r.FormValue("title")`でタイトルを取得する。
    2. タイトルが空でなければ`h.Repo.Insert(title)`を呼び出す。
    3. `http.Redirect`で`/todos`へリダイレクトする。

### 6.4. `database` パッケージ
- **ファイル**: `internal/database/database.go`
- **関数**: `InitDB(filepath string) *sql.DB`
  - 引数で指定されたファイルパスでSQLiteデータベースに接続する (`sql.Open`)。
  - セクション4.2で定義された`CREATE TABLE`文を実行する。
  - `*sql.DB`のインスタンスを返す。致命的なエラーの場合はプログラムを終了する。

### 6.5. `cmd/server` パッケージ
- **ファイル**: `cmd/server/main.go`
- **main関数処理フロー**:
  1. `database.InitDB("todo.db")`を呼び出し、DB接続を初期化する。
  2. `repository.NewSQLiteTodoRepository(db)`を呼び出し、リポジトリのインスタンスを生成する。
  3. `handlers.NewTodoHandler(repo)`を呼び出し、ハンドラのインスタンスを生成する（依存性の注入）。
  4. `chi.NewRouter()`でルーターを初期化し、ロガー等のミドルウェアを設定する。
  5. セクション2の機能要件に従い、各エンドポイントとハンドラのメソッドをルーティングする。
  6. `http.ListenAndServe(":8080", r)`でWebサーバーを起動する。

## 7. 画面（テンプレート）設計

- **ファイル**: `templates/index.html`
- **渡されるデータ構造**:
  ```
  map[string]interface{}{
      "Todos": []models.Todo,
  }
  ```
- **画面要素**:
  1. **Todo一覧**:
     - `{{range .Todos}}`を使い、`Todos`スライスをループ処理する。
     - 各ループ内で`{{.Title}}`と`{{.CreatedAt}}`を表示する。
     - Todoが1件もない場合は`{{else}}`ブロック内のメッセージを表示する。
  2. **Todo追加フォーム**:
     - `action="/todos"` `method="post"` を持つ`<form>`タグ。
     - `name="title"` を持つ`<input type="text">`フィールド。
     - `<button type="submit">`の送信ボタン。

## リファクタリング案
1. Service (Usecase) レイヤーの導入
2. エラーハンドリングの強化
3. 設定の外部化
4. バリデーションの導入
5. データベースマイグレーションの導入
6. 本格的なテストの実装