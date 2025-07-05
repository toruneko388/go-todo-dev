# 🛠️ 詳細設計（Todoアプリ・初期版）

---

## 1. モジュール構成（ファイル単位）

| ファイル名        | 役割                       |
|-------------------|----------------------------|
| `main.go`         | 全体のエントリポイント（起動処理・ルーティング定義） |
| `handler.go`      | HTTPハンドラ群（一覧表示・追加処理）     |
| `model.go`        | Todo構造体およびデータ保持               |
| `templates/`      | HTMLテンプレート格納ディレクトリ         |

---

## 2. データ構造（構造体定義）

```go
// model.go
type Todo struct {
    ID    int    // 自動採番（連番）
    Title string // 1〜100文字
    Done  bool   // 初期値はfalse
}

var todos []Todo      // メモリ保持
var nextID int = 1    // 自動採番用
```

---

## 3. ルーティングと処理

| メソッド | パス     | 関数名         | 処理概要                                |
|----------|----------|----------------|-----------------------------------------|
| `GET`    | `/todos` | `handleList`   | Todo一覧を取得しテンプレート表示        |
| `POST`   | `/todos` | `handleAdd`    | フォームからタイトルを受け取り、Todo追加 |

---

## 4. 処理フロー

### 🟩 Todo一覧表示（GET /todos）

```text
1. `handleList` が呼ばれる
2. `todos` スライスの内容を取得
3. `templates/todos.html` に渡してHTMLレンダリング
4. ブラウザに返す
```

---

### 🟦 Todo追加（POST /todos）

```text
1. `handleAdd` が呼ばれる
2. フォーム値 `title` を取得
3. 空欄チェック・文字数制限
4. 新しい `Todo{ID: nextID, Title: title, Done: false}` を作成
5. `todos` に追加し、`nextID++`
6. `/todos` にリダイレクト
```

---

## 5. テンプレートの設計（todos.html）

```html
<h1>Todo List</h1>
<ul>
  {{range .Todos}}
    <li>{{.Title}}</li>
  {{end}}
</ul>

<form method="POST" action="/todos">
  <input type="text" name="title" maxlength="100" required>
  <button type="submit">Add</button>
</form>
```

---

## 6. バリデーション仕様（POST）

- `title` は必須（空欄不可）
- 最大100文字（それ以上はエラー）
- エラー時は一覧画面上部に赤字で表示
