# 詳細設計（実装依頼用）
## ✅ 概要

- ユーザーがTodoを1つ追加できる
- 追加されたTodoは一覧に表示される
- データはメモリに保持（[]Todo スライス）
- Webブラウザ上でフォーム入力・一覧表示ができる
  - アプリは `http://localhost:8080` でアクセス

## 構造体
`Todo` 構造体 ： `ID (int), Title (string)` がある


## 💾 データ保持方法
- `[]Todo` スライス：グローバル変数として定義
- IDは自動で連番を振ってもらえるとよい（方法は任せる）

## ✳️ 追加機能
### 機能：Todoを一覧表示する

- get /todos
- やりたいこと：
  1. []Todo から現在のTodo一覧を取得
  2. 一覧を表示

### 機能：Todoを追加する

- post /todos
- 入力：フォームから取得
- やりたいこと：
  1. フォームから文字列を受け取る
  2. 空欄でなければ新しいTodo構造体を作成
  3. []Todo スライスに append
  4. `/todos` にリダイレクト


## 🖼️ テンプレート要件（`index.html`）

- タイトルを入力するテキストフィールド
- 送信するフォームボタン
-  `[]Todo` スライスをテンプレートに渡してTodo一覧を表示
- テンプレートエンジン：Go標準の `html/template` を想定

## ディレクトリ構成
ルートディレクトリ app 以下に、Goの標準的な構成に沿って以下の役割ごとにディレクトリを分けてください：
エントリーポイント：cmd/server/
データ構造：internal/models/
ハンドラ処理：internal/handlers/
テンプレート：templates/

## 使用パッケージ
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)
- 使用言語：Go（Golang）
- 使用パッケージ：net/http, html/template
- モジュール名：`todoapp`
- ビルドや実行に必要なGoモジュール定義は自動生成でOK

