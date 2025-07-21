# ツールの使い方
## goplantuml
- 出力をファイルに保存する
```
goplantuml -recursive ./ > tools/diagram.puml
```
- サイトに内容を貼り付けて表示
https://www.plantuml.com/plantuml/

## godepgraph
- 出力をファイルに保存する
```
go list ./... | xargs godepgraph -s -o github.com/toruneko388/todoapp > tools/deps.dot
```
 - サイトに内容を貼り付けて表示
http://www.webgraphviz.com/