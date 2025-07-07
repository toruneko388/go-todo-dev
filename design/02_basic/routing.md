## ルーティング図
URLと処理（GET/POSTやハンドラ）の対応関係を整理する

```mermaid
graph TD
    A[ブラウザ] -->|GET /todos| B[HandleList]
    A -->|POST /todos| C[HandleAdd]
    B --> D[テンプレート todos.html]
    C --> E[Todo を生成し TodoStore に追加]
    C --> F[リダイレクト /todos]
```