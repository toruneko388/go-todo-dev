以下の内容のみを前提として、コードを作成してください

# Go言語で作っているTodoアプリについての依頼

## アプリ概要

現在、Go言語でWebベースのTodoアプリを開発しています。  
ブラウザ上でタスク（Todo）を一覧表示・追加できる仕組みです。

## 現時点の実装状況

- Todo一覧表示機能（GET `/todos`）
  - メモリ上（`[]Todo` スライス）に保持したタスクを一覧表示
  - HTMLテンプレートを使用して表示
- Todo追加機能（POST `/todos`）
  - フォームからの入力により、タイトルを持つTodoを1件追加
  - データはまだ永続化されておらず、サーバー再起動で消える

## 現状の使用パッケージ一覧
- ルーティング：`chi`
- テンプレート：`html/template`
- データベース：なし
- セッション：なし
- バリデーション：なし
- ミドルウェア：`chi/v5/middleware`
- ORM/SQLラッパー：なし
- JSON：標準
- テスト：未導入

## 今回依頼したい内容（詳細仕様）

以下の仕様に基づいて、SQLiteを使ったTodoの永続化機能を追加してほしいです。

- SQLiteデータベースを使ってTodoを保存するようにする
- アプリ起動時にDBファイル（例：`todo.db`）を作成・接続
- Todo用のテーブルを作成（まだなければ自動で）
- 保存したい情報は以下の3つです：
  - 一意なID（自動採番される数値）
  - タイトル（文字列、必須）
  - 作成日時（日付時刻）
- テーブル名は `todos` にしたいです。
- カラム名や型は任せます。
- Todo追加時はこのDBに `INSERT` する
- Todo一覧表示時はDBから `SELECT` して表示する
- 今のメモリベースのスライス管理は削除する
- 可能であれば、テーブル作成SQLもコード内で自動生成してほしい

## ディレクトリ構成
ルートディレクトリ app 以下に、Goの標準的な構成に沿って以下の役割ごとにディレクトリを分けてください：
エントリーポイント：cmd/server/
データ構造：internal/models/
ハンドラ処理：internal/handlers/
テンプレート：templates/

## モジュール構成（現在の実装）
- すでに動作している部分はこの構成の通りです。実装に合わせてパッケージを増やし構いません
- 今回の依頼では、必要な部分のみを差し替え・追加してもらえれば構いません。
- 未定義の関数・構造体などは新たに提案してもらって構いません。
- モジュール名：`todoapp`

### cmd/server パッケージ
- main.go
  - エントリーポイント

### models パッケージ
- todo.go
  - `type Todo` ID・タイトルを持つTodo構造体
  - `func GetAll() []Todo` 現在のTodo一覧を返す（メモリ上に保持）

### handlers パッケージ
- todo_handler.go
  - `type TodoHandler` Goの標準テンプレート構造体を内部に保持し、テンプレートを使ってTodoの一覧などを表示する役割を持つハンドラ構造体
  - `func NewTodoHandler() *TodoHandler` テンプレートをパースして返す
  - `func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request)` Todo一覧をテンプレートで表示
  - `func (h *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request)` フォーム送信されたTodoを追加

## テンプレート
- `index.html` はすでに作成済みで、一覧表示に使用
- 今回の依頼では `index.html` を新しく作成・変更する必要はありません。

## 注意点・補足

- 今後の拡張を見越して、DB接続処理はなるべく分離したい