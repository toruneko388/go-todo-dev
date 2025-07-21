# go-todo-dev

# ツール
## goplantuml
- 出力をファイルに保存する
```
goplantuml -recursive ./ > diagram.puml
```
- サイトに内容を貼り付けて表示
https://www.plantuml.com/plantuml/

## godepgraph
- 出力をファイルに保存する
```
go list ./... | xargs godepgraph -s -o github.com/toruneko388/todoapp > deps.dot
```
 - サイトに内容を貼り付けて表示
[Webgraphviz](http://www.webgraphviz.com/)