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

3). 整体校验过程

4. 缓存查询结果

go命令每次访问GOSUMDB得到的数据都会缓存在本地的$GOPATH/pkg/mod/cache/download/sumdb/sum.golang.org目录中。
缓存目录结构如下所示：
```golang
~/project/golang/pkg/mod/cache/download/sumdb/sum.golang.org$ tree *
lookup
├── cloud.google.com
│   ├── go@v0.0.0-20170206221025-ce650573d812
│   ├── go@v0.26.0
│   └── go@v0.34.0
├── fyne.io
│   └── fyne@v1.4.2
├── github.com
│   ├── !burnt!sushi
│   │   └── toml@v0.3.1
│   ├── !geert!johan
│   │   ├── go.incremental@v1.0.0
│   │   └── go.rice@v1.0.0
│   ├── !google!cloud!platform
│   │   └── cloudsql-proxy@v0.0.0-20190129172621-c8b1d7a94ddf
│   ├── !jeffail
│   │   ├── gabs
│   │   │   └── v2@v2.6.0
│   │   └── gabs@v1.4.0
│   ├── !knetic
...
```

### 12.3.14 第三方代理

#### 1.模块代理  

1). 如何配置第三方代理

不管使用哪个代理服务，都需要将代理服务器地址配置到GOPROXY环境变量中，假如第三方代理服务地址为
https:/proxy.example.org,Go版本为1.13及以上，则可以使用如下命令：
```text
export GOPROXY="https://proxy.example.org,direct"
```
可以同时指定多个代理服务器，使用逗号分隔即可，额外指定一个direct可以保证在代理 无法工作时使用go命令仍能从源版本控制系统中获取模块。

2). 可用的第三方代理

目前国内用得较多的公共代理服务主要有以下几个：  
goproxy.io(开源项目)，国内最早的代理服务；  
goproxy.cn(开源项目)，目前由七牛云提供服务；  
mirrors.tencent.com/go,由腾讯云提供服务；  
mirrors.aliyun.com/goproxy,由阿里云提供服务。  

#### 2.校验和数据库代理  
校验和数据库通常也是使用模块代理服务器代理的，为了能让模块代理服务器代理校验和数据库，模块代理服务器需要额外实现一个supported接口：
```text
<proxyURL>/sumdb/<sumdb-name>/supported
```
go命令在访问GOSUMDB环境变量指示的数据库时，会先通过此接口询问模块代理是否能够代理校验和数据库，
如果该接口的HTTP返回码为200，那么g0命令将通过该代理校验模块的Hash值。

读者可以使用该接口测试第三方代理是否支持代理校验和数据库，以goproxy.io为例：
```text
curl -w %{http_code} https://goproxy.io/sumdb/sum.golang.org/supported
200
```

使用代理goproxy.io查询官方数据库中的最新信息：
```text
curl https://goproxy.io/sumdb/sum.golang.org/latest
go.sum database tree
8629467
5Wj34KHPtaJKm+Hu6xowQDZp3vOty1Bw3qyAM5GZBHA=

— sum.golang.org Az3grsClmczmqBdbN8BJx0NkKwbvDyBs6npE6qJDVEp5Xed207UJZfb3x16rzNyAn1HXf4gN3uRurmqBS9ELpK6zKQU=
```

#### 3.安全性


### 12.3.15 私有模块

#### 1.私有模块配置

#### 2.私有代理场景

### 12.3.16 Go Module的演进

Go Module特性旨在提供一种官方的包管理系统，用于替代GOPATH。同时Go Module的出现也终结了第三方包管理系统百家争鸣的时代，
越来越多的组织、项目转而采用Go Module来管理依赖包。

然而，该特性并不是一蹴而就的，而是历经多个版本的打磨才逐渐走向成熟。2018年8月，Go团队发布1.11版本，Go Module作为一个实验特性被引入，
并在1.12、1.13版本中不断迭代，最终于2020年2月发布的1.14版本中走向成熟，经过一年半左右的开发，Go团队才自信地
宣布Go Module可以在生产环境中使用了。

若要使用Go Module特性，笔者推荐直接使用Go1.l4+,从而获得最终的Go Module体验。

本节简要地记录Go Module特性在演进过程中的变化，以期帮助读者更好地理解该特性。

#### 1. Go 1.11

该版本以实验性质推出Go Module特性，引人一个新的环境变量GO111MODULE用于控制特性开关。

- GO111MODULE=auto:自动模式； 
- GO111MODULE=on:显式开启模式；
- GO111MODULE=off:显式关闭模式。

GO111MODULE的默认值为auto,在此模式下，go命令将根据当前项目是否在GOPATH中决定是否启用GOMODULE。
为了不让原用户受Go Module影响，如果项目位于GOPATH目录中，那么Go Module特性将不开启、只有位于非GOPATH目录中的项目才会自动启用Go Module

- Go111MODULE=on表示任何情况下都启用Go Module,即便项目在GOPATH目录中也启用。
- Go111MODULE=off表示任何情况下都不启用Go Module,即便项目在GOPATH外面也不启用。

另外，该版本已支持模块代理，由于官方的代理服务器(proxy.golang.org)还未推出，所以环境变量GOPROXY为空：
```text
GOPROXY=""
```
作为初出茅庐的特性，Go Module在这个版本中主要考虑的是不影响既有用户。只有愿意体验新特性的用户，
通过显式的操作才可以使用，比如将项目移到GOPAT外面或显式地设置 GO111MODULE=on。

#### 2. GO 1.12

该版本正式推出Go Module,可以理解为Alpha阶段。  

该版本相较于1.11版本并无明显的变化，但有一个值得注意的小特性被引入、并且保留 到了最终的1.14版本中。
那就是在项目的go.mod文件中记录Go的版本，比如Go 1.12、它指示构建项目需要使用的最低Go版本，
如果使用低于此版本的Go来构建，则在出错时会给出额外的提醒。

#### 3. Go 1.13

Module体验。
该版本可理解为Go Module特性走向Beta阶段。

该版本GO111MODULE的默认值仍为auto,该模式有一个显著的变化、那就是它变得更加智能，
即如果项目中包含go.mod文件，那么默认将开启Go Module,即便项目位于GOPATH目录中。

另外，该版本重点加强了模块的下载、校验及对私有模块支持。  
- 推出官方的模块代理服务器，并把GOPROXY默认值设为官方代理服务器地址：  
https://proxy.golang.org,direct
- 推出官方的模块校验服务器，并把GOSUMDB默认值设为官方校验服务器地址：  
sum.golang.org。
- 引人环境变量GOPRIVATE,用于指定私有模块列表。

#### 4. Go 1.14
该版本可理解为Go Module特性正式走向GA(General Availability)。

笔者认为这个版本最大的亮点在于体现了对vendor的认可，在Go Module设计阶段，Go
团队曾想废弃vendor，在众多开发者的提议下才决定保留vendor。并且在Go1.14之前的版本
中，只要开启了Go Module、那么vendor默认是被忽视的，如果需要使用vendor，则需要提供
显式的参数，如go build -mod=vendor。在1.14版本中则变成了只要项中包含vendor,并
且go.sum中指定的Go版本大于或等于1.14、则默认启用vendor。

另外，该版本对待不规范的模块的态度有所转变。所谓不规范的模块，即模块Màjor的版
本号大于，但其module名字并没有体现版本号。不规范的模块也称为不兼容的模块，更多的
细节可参考前面的章节，这里不再赘述。在Go Module早期，为了兼容既有的模块，Go
Module能够容忍不兼容的模块、这些不兼容的模块会在go.mod中用+incompatible标注。在
1.14版本中，Go命令行工具将默认忽略不兼容版本，除非用户显式地指定依赖。

#### 5. Go Module实践

Go Module虽然是设计精良的模块管理特性、但在工程领域它也面临着一些挑战，比如：
- 如果构建环境无法访问外网怎么办？
- 如果使用的开源模块闭源或删除了怎么办（历史上曾经发生过类似的案例）？
- 如果版本控制系统宕机了怎么办？

尽管上述问题都可以使用vendor来解决、但使用vendor也有不足，比如会增大仓库体积，vendor变得更难以“review”等问题。

目前业界使用比较多的方案是Go Module + vendor,即使用Go Module来管理模块依赖，最后把依赖模块保存到vendor中。