# Go 项目构建脚本
# 支持: 编译二进制、构建前端、清理产物、运行测试

# 定义变量
BINARY_NAME := app
WEB_DIR := web
DIST_DIR := ./web/dist

# 开发模式: 同时运行Go服务和前端热更新
dev:
	@echo "启动服务（按 Ctrl+C 关闭所有服务）..."
	trap 'kill -TERM $$PID1 $$PID2; wait' INT TERM; \
	(cd $(WEB_DIR) && npm run dev) & PID1=$$!; \
	(GIN_MODE=debug go run main.go) & PID2=$$!; \
	wait $$PID1 $$PID2

build:
	@echo "启动前端后端打包..."
	rm -rf ./bin
	(cd $(WEB_DIR) && pnpm i && npm run build)
	mkdir -p ./bin/web
	cp -r $(DIST_DIR)/* ./bin/web
	go build -o ./bin/$(BINARY_NAME) main.go

release:
	@echo "release模式启动项目..."
	trap 'kill -TERM $$PID1; wait' INT TERM; \
	(GIN_MODE=release ./bin/$(BINARY_NAME)) & PID1=$$!; \
	wait $$PID1

run: build release

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make dev       - 本地运行"
	@echo "  make build     - 编译前后端"
	@echo "  make release   - release模式启动项目"
	@echo "  make run		- 编译前后端并且启动项目"