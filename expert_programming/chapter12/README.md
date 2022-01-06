# Chapter 12 Go语言依赖管理

### 12.3.11 模块代理

#### 1.什么是GOPROXY

#### 2.代理协议

1). 获取模块列表

2). 获取模块元素数据

3). 获取go.mod文件

4). 获取代码压缩包

5). 获取模块的最新可用版本

6). 下载过程

#### 3.下载过程步骤  
使用go mod download命令观察请求代理服务器的信息：
```text
go mod download -x -json github.com/google/uuid@latest

# get https://goproxy.cn/github.com/google/uuid/@v/list
# get https://goproxy.cn/github.com/google/uuid/@v/list: 200 OK (0.358s)
{
	"Path": "github.com/google/uuid",
	"Version": "v1.3.0",
	"Info": "/Users/steven/project/golang/pkg/mod/cache/download/github.com/google/uuid/@v/v1.3.0.info",
	"GoMod": "/Users/steven/project/golang/pkg/mod/cache/download/github.com/google/uuid/@v/v1.3.0.mod",
	"Zip": "/Users/steven/project/golang/pkg/mod/cache/download/github.com/google/uuid/@v/v1.3.0.zip",
	"Dir": "/Users/steven/project/golang/pkg/mod/github.com/google/uuid@v1.3.0",
	"Sum": "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
	"GoModSum": "h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo="
}
```

#### 4.小结


### 12.3.12 GOSUMDB的工作机制

在Go 1.13中，引入了新的环境变量GOSUMDB，其默认值为官方的checksum database，即校验和数据库sum.golang.org：  
```text
GOSUMDB="sum.golang.org"
```

#### 1.环境变量


#### 2.引入GOSUMDB的背景

1). 正确的模块版本  

2). go.sum的不足

#### 2.引入GOSUMDB的背景


#### 3.GOSUMDB的总做机制

1). GOSUMDB是什么

校验和数据库是公开的，读者可以使用/lookup/M@v接口来查询特定的模块版本的Hash值：  
```text
curl https://sum.golang.org/lookup/github.com/google/uuid@v1.3.0
5994992
github.com/google/uuid v1.3.0 h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=
github.com/google/uuid v1.3.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=

go.sum database tree
8617558
f5KSnshtcJa62/LNDuI7gZcszmfvqVvcxkPMjNf3xZs=

— sum.golang.org Az3grq9hzFjHaUJ4ohlzx2zGTtq8x4UOHTEd6EJyvB5Zd1/9vNOznd6OMy9KHrwN3CcPvzbxaFPjUPRV+qFvpA2xqQo=
```

2). 校验流程

3). 数据库数据来源


### 12.3.13 GOSUMDB的实现原理

#### 1.数据库的数据结构

#### 2.数据库的查询接口

1). 查询数据库的整体信息

使用/latest 接口（下面简称latest 接口）可以查询数据库的整体信息，读者也可以使用如下命令查询：
```text
curl https://sum.golang.org/latest
go.sum database tree
8618520
DbQn8MpDm8eo3AYTIcu7kG8AOwJxoSv1bQjPzvXjG60=

— sum.golang.org Az3grq5WfO+7Y2YiAfTc8w/Vh/2lB7RDsxCYssRnr4UbtORI1SyHUNEX1KG5JnJccaJEZCr9FnTyD4XJy594u65WeQI=
```

- 当前树的总结点数，随着新模块的不断加入，节点数也会随之增大，上例中说明树当前包含8618520个节点；
- 当前树顶节点的Hash值；
- 响应该请求的服务器信息，以便在出错时能够定位；

2). 查询模块版本和Hash值

3). 查询任意节点

#### 3.客户端审计

1). 确保对端是规范的

2). 确保对端没有被篡改
