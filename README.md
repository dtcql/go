# GO学习

## 1. 安装

### 1.1 安装go

golang.org或者网络搜索golang，直接下载安装。

### 1.2 go env配置环境变量

* 开启go module
```
go env -w GO111MODULE=on
```

* 开启go代理
```
go env -w GOPROXY=https://goproxy.cn,direct
```

### 1.3 设置go工程

* 开启go module
```
go mod init hello-world    
```

## 2. go web 框架 - gin

文档 =》 https://gin-gonic.com/zh-cn/docs/

### 2.1 下载并安装Gin
```
go get -u github.com/gin-gonic/gin
```
