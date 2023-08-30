.PHONY: all build clean run check cover lint docker help logo

gotter ="   _____    ____    _______   _______   ______   _____\n"
gotter +="  / ____|  / __ \  |__   __| |__   __| |  ____| |  __ \ \n"
gotter +=" | |  __  | |  | |    | |       | |    | |__    | |__) |\n"
gotter +=" | | |_ | | |  | |    | |       | |    |  __|   |  _  /\n"
gotter +=" | |__| | | |__| |    | |       | |    | |____  | | \ \\n"
gotter +="  \_____|  \____/     |_|       |_|    |______| |_|  \_\\n"

#记录的时间
time = $(shell date "+%Y-%m-%d %H:%M:%S")
# 应用生成的位置
OUT_FILE_PATH=./bin/app
# docker网络名称
NETWORK_NAME="docker_network"



ifeq ($(OS),Windows_NT)
 	# Windows 系统下的命令
    SWAG_COMMAND = ./cmd/swag_win.exe init
else ifeq ($(shell uname),Darwin)
    # macOS 系统下的命令
    SWAG_COMMAND = ./cmd/swag_mac init
else ifeq ($(shell uname),Linux)
   # Linux 系统下的命令
     SWAG_COMMAND = ls
else
    $(error Unsupported operating system: $(shell uname))
endif

run:build logo
	${OUT_FILE_PATH}_mac
build:
	@echo "==============编译开始:"$(time)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUT_FILE_PATH}_linux main.go
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${OUT_FILE_PATH}_win main.go
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${OUT_FILE_PATH}_mac main.go
	@echo "==============编译结束:"$(time)

clean:
	rm -f ${OUT_FILE_PATH}_linux
	rm -f ${OUT_FILE_PATH}_win
	rm -f ${OUT_FILE_PATH}_mac

docker:build
#如果网络不存在则创建
ifeq ("$(shell docker network ls | grep "${NETWORK_NAME}")","")
	docker network create ${NETWORK_NAME}
endif
	docker-compose down
	docker-compose up -d
clean_docker:
	docker-compose down
logo:
	@echo  $(gotter)
help:
	@echo "run 编译并运行"
	@echo "build 编译"
	@echo "clean 删除编译的程序"
	@echo "docker 编译之后删除容器再创建docker容器"
	@echo "clean_docker 删除容器"
	@echo "logo 显示logo"
	@echo "release 发布编译文件到仓库"
	@echo "swag 根据注释自动生成文档"

release:build
	cd ../jdjy-go-build && git pull
	cp ${OUT_FILE_PATH}_linux ../jdjy-go-build/bin/app_linux
	cd ../jdjy-go-build && git add * && git commit -m "update" && git push

swag:
	$(SWAG_COMMAND)
deploy_test:
	./deploy_test.sh