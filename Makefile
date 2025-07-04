# Go 项目构建脚本
# 支持: 编译二进制、构建前端、清理产物、运行测试

# 定义变量
BINARY_NAME := app.exe
WEB_DIR := web
DIST_DIR := ./web/dist

# 默认目标: 编译完整项目
all: build-frontend build-go copy-web

# 构建前端（React/Vite项目）
build-frontend:
	cd $(WEB_DIR) && pnpm build

# 将dist目录复制到bin目录并且改名为web
copy-web:
ifeq ($(OS),Windows_NT)
	robocopy $(DIST_DIR) .\bin\web /E /NDL /NFL /NJH /NJS /NP
else
	cp -r $(DIST_DIR)/ ./bin/web
endif

# 编译Go二进制
build-go:
	go build -o ./bin/$(BINARY_NAME) main.go

# 运行单元测试
test:
	go test -v ./...

# 清理产物
clean:
	- rm -f $(BINARY_NAME)
	- rm -rf $(DIST_DIR)

# 开发模式: 同时运行Go服务和前端热更新
dev:
	@echo "启动前端热更新服务..."
	cd $(WEB_DIR) && start pnpm dev
	@echo "启动Go开发服务..."
	go run main.go

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make all       - 完整构建（前端+Go）"
	@echo "  make build     - 仅编译Go二进制"
	@echo "  make test      - 运行单元测试"
	@echo "  make clean     - 清理构建产物"
	@echo "  make dev       - 启动开发模式（前端热更新+Go服务）"
	@echo "  make help      - 显示帮助信息"