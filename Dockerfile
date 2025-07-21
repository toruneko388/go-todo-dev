FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest && \
    go install github.com/kisielk/godepgraph@latest

# Go モジュールのキャッシュを有効にするための設定
# --mount=type=cache,target=/go/pkg/mod : Go モジュールのキャッシュを /go/pkg/mod に保存します（Go の公式イメージではモジュールキャッシュのデフォルト場所が /go/pkg/mod）
# sharing=locked : マウントの共有モードを locked にすることで、実行している間は単独でロックすることで競合を避けます
# --mount=type=bind,source=app/go.mod,target=go.mod : ホストの go.mod をコンテナの go.mod にマウントします
# go mod download -x : Go モジュールの依存関係をダウンロードして、モジュールキャッシュに帆損する（デバッグ用の verbose（詳細ログ））。
RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked \
    --mount=type=bind,source=app/go.mod,target=go.mod \
    --mount=type=bind,source=app/go.sum,target=go.sum \
    go mod download -x              # 依存だけ落としてキャッシュ

# アプリケーションのビルドを行う場合は以下のコマンドを有効にします
#RUN --mount=type=cache,target=/go/pkg/mod/ \
#    --mount=type=bind,source=./app, target=./app \
#    go build -o /bin/myapp

# Air が /go/bin/air に入るので PATH を通す
ENV PATH="/go/bin:$PATH"

COPY ./app .

# 非rootユーザーを作成
#RUN addgroup -g 1000 appgroup && \
#    adduser -D -u 1000 -G appgroup appuser && \
#    chown -R appuser:appgroup /app
#USER appuser

# 現在はvscodeのタスクでairを実行するので、CMDはコメントアウトしています。
# CMD ["air", "-c", ".air.toml"]
CMD ["sleep", "infinity"] 