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

### 1.3 创建go工程

* 创建go module
```
go mod init hello-world    
```

### 1.4 运行go工程

```
go run main.go   
```

### 1.5 编译go工程

```
go build  
```

## 2. go web 框架 - gin

文档 =》 https://gin-gonic.com/zh-cn/docs/

### 2.1 下载并安装Gin
```
go get -u github.com/gin-gonic/gin
```
会安装在gopath/pkg/mod下，看gopath用go env命令。

### 2.2 引入Gin框架
```
import "github.com/gin-gonic/gin"
```

### 2.3 Gin框架hello world实例

```
package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func main() {
	r := gin.Default() //返回默认的路由引擎

	r.GET("/hello", sayHello) //GET请求执行sayHello函数

	r.Run(":9090") //启动服务
}
```
go build之后运行可执行文件，浏览器访问http://localhost:9090/hello


### 2.4 Gin框架引入前端dist目录实例

将dist目录下的static目录，index.html...拷贝到go project根目录里。在go 项目中新建templates文件夹，将index.html, favicon.ico放入。然后编码如下，在go build，之后运行生成的可执行文件。
```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()            //返回默认的路由引擎
	r.Static("/static", "static") //静态文件去哪找
	r.LoadHTMLGlob("templates/*") //去哪找html模板文件
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}) //GET请求执行匿名函数返回index.html

	r.Run(":9090") //启动服务
}
```
