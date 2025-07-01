FROM golang:1.24-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 將 Go 模組文件複製到工作目錄
COPY go.mod go.sum ./

# 下載所有 Go 模組依賴
RUN go mod download

# 將應用程式原始碼複製到工作目錄
COPY . .

# 編譯應用程式，啟用 CGO_ENABLED=0 以建立靜態連結的二進位檔，適合從頭映像運行
# -o hello-world 指定輸出執行檔名
# ./... 表示編譯當前目錄及其子目錄下的所有 Go 程式碼
RUN CGO_ENABLED=0 go build -o hello-world .

# 使用一個更小的、基於 scratch 的映像來運行編譯好的應用程式
FROM alpine:latest

# 設定工作目錄
WORKDIR /root/

# 從 builder 階段複製編譯好的執行檔
COPY --from=builder /app/hello-world .

# 暴露應用程式可能使用的埠 (如果有的話，例如 HTTP 服務)
EXPOSE 8080

# 定義容器啟動時執行的命令
ENTRYPOINT ["./hello-world"]
