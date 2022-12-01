# 伪目标，防止命令与实际工作目录下的文件名冲突
.PHONY: all build run gotool clean help

# 指定输出的文件名
BINARY="golang-Makefile"

all: gotool build

# build 指令
build:
	set GOOS=linux
	set GOARCH=amd64
	go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY}; fi

# 命令提示
help:
	@echo "make - 格式化Go代码，并编译生成二进制文件"
	@echo "make build - 编译Go代码，生成二进制文件"
	@echo "make run - 直接运行Go代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行GO工具 'fmt'和'vet' "