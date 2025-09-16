# gingerale

使用 Gin 框架的最小示例，提供简单的 HTTP 路由。

## 运行环境
- Go 1.20+（推荐）

## 快速开始
1. 初始化并下载依赖（首次）：
   ```bash
   go mod tidy
   ```
2. 运行服务（本地端口 8080）：
   ```bash
   # 在受限环境下建议为 Go 构建缓存指定本地目录
   GOCACHE=$(pwd)/.gocache go run .
   ```

## 示例请求
- 根路由：
  ```bash
  curl -i http://localhost:8080/
  ```
  响应：`Welcome to gingerale with Gin!`

- 健康检查：
  ```bash
  curl -i http://localhost:8080/ping
  ```
  响应：`{"message":"pong"}`

## 目录结构
```
.
├── LICENSE
├── README.md
├── go.mod
└── main.go
```

## 备注
- 如果你的环境对网络访问有限制，首次 `go get` 或 `go mod tidy` 可能需要代理或放行。
- 生产环境建议添加优雅关闭、日志配置、配置管理等。
