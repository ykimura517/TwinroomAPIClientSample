FROM golang:1.21

EXPOSE 8083

# ワーキングディレクトリの設定
WORKDIR /app

# 依存関係のコピー
COPY go.mod .
COPY go.sum .

# 依存関係のインストール
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN go build -buildvcs=false -o main .

# アプリケーションの実行
CMD ["./main"]
