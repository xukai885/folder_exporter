FROM golang:alpine AS builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
# 将代码复制到容器中
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app .
FROM scratch
COPY --from=builder /build/app /
# 声明服务端口
EXPOSE 9101
# 启动容器时运行的命令
CMD ["/app"]