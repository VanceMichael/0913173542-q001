# 城市马拉松成绩复核服务

这是一个 Go 后端工程起点。服务端口为 8080，数据文件默认位于 `/data/race.db`。

运行：`docker build -t marathon-review .`，随后执行 `docker run --rm -p 8080:8080 marathon-review`。

接口目前提供健康检查和成绩复核资源的基础入口，后续业务数据应通过 HTTP 接口写入并持久化。测试命令为 `go test ./...`。
