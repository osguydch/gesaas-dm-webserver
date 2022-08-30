# RmServer

WebUI / Run Manager

### 环境
```
golang版本1.19
```

### 准备
```shell
#拉取子模块
git submodule update --init 
#更新子模块
git submodule update --remote
#编译thrift
thrift -r --gen go:package_prefix=rm/gen-go/ dm_idl/header.thrift 
```

### 单元测试
```shell
go test ./...
```

### 打包
```shell
git tag 0.0.1
make
#make会将tag和时间打进程序中，使用rmserver version查看
rmserver version
#module: RmServer
#desc: RunManager Server repo
#version: 0.0.1-0-g8f63ff0
#build: 2022-03-31 09:59:19
```

### 执行
```shell
#默认使用配置文件./config/config.toml启动
rmserver server 
#查看帮助
rmserver server -h 
#使用指定配置文件夹
rmserver server -c config/
#release模式
rmserver server -c config/ -p 8089 -m release 
```
