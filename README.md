# gopl 学习笔记

## start

### 安装

```bash
GOROOT=安装包的位置
PATH=%PATH%;%GOROOT%\bin
GOPATH=D:\gocode
```

### 配置

```bash
go env -w GO111MODULE=on
# GOPROXY=https://proxy.golang.com.cn,direct
# GOPROXY=https://goproxy.io,direct
# GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

#go get -v github.com/xiaobo9/gopl
# 更新
#go get -u -v github.com/xiaobo9/gopl

git config --global url."git@git.xiaobo9.top:".insteadOf "http://git.xiaobo9.top/"
go env -w GOPRIVATE=git.xiaobo9.top
go get -v -insecure git.xiaobo9.top/projectTest # http请求服务，https 不用 insecure

```

### project

```bash
go mod init github.com/xiaobo9/gopl
go mod download
go install github.com/xiaobo9/gopl
```

### import local modules

```bash
require "github.com/userName/otherModule" v0.0.0
replace "github.com/userName/otherModule" v0.0.0 => "local physical path to the otherModule"
```

### build

```bash
GOOS=linux GOARCH=amd64 go build gopl
# 压缩体积
go build -ldflags="-s -w" -o server main.go
# 用 upx 压缩
go build -o server main.go && upx -9 server
#用 upx 进一步压缩 
go build -ldflags="-s -w" -o server main.go && upx -9 server
```

### debug

```bash
go get -d github.com/go-delve/delve/cmd/dlv
go install github.com/go-delve/delve/cmd/dlv
```

## 包的使用

因为反射，不推荐 Logrus，推荐 zap，zerolog

因为反射，不推荐 encoding/json，推荐 Easyjson，研究下 json-iterator ?

进度条 progressbar github.com/schollz/progressbar

### Go embed

把文件以及目录中的内容都打包到 exe 中

### gin

### hugo

### tinygo

`https://tinygo.org/`

### gorilla/mux

### httprouter

### chi

## docker

```Dockerfile
# https://hub.docker.com/_/golang
# 基础镜像
FROM stretch
```

## 工具

`https://github.com/russross/blackfriday`

`https://github.com/posener/complete`

### vscode

```json
// .vscode/launch.json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}"
            // "program": "${workspaceFolder}"
            // "program": "${workspaceFolder}/ch1/server/main/"
        }
    ]
}
```

## go 接口 方法判断

```go
// if se, ok := v.Value.(SetOptioner); ok {
// 判断是否有 SetOptioner 方法的接口，然后就可以直接调用了？
// html.go L525 renderLink
// n := node.(*ast.Link) 直接获取接口实现类的属性
for i := l - 1; i >= 0; i-- {
    v := r.config.NodeRenderers[i]
    nr, _ := v.Value.(NodeRenderer)
    if se, ok := v.Value.(SetOptioner); ok {
        for oname, ovalue := range r.options {
            se.SetOption(oname, ovalue)
        }
    }
    nr.RegisterFuncs(r)
}
```
