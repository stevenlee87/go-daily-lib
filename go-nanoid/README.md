## 使用说明

1.mkdir go-nanoid & && cd go-nanoid  

2.go mod init github.com/stevenlee87/go-daily-lib/go-nanoid

3.go get -u github.com/matoous/go-nanoid/v2

4.example1
```go
package main

import (
    "fmt"
    "github.com/matoous/go-nanoid/v2"
)

func main() {
    id, err := gonanoid.New()
    if err != nil {
        fmt.Println("Error generating ID:", err)
        return
    }

    fmt.Println("Generated ID:", id)
}
```

运行后你将得到类似如下的输出：
Generated ID: VDAwsOqCM6EaXvCYapHUd

5.example2  
指定长度生成随机 ID  
你还可以通过传递长度参数来生成指定长度的 ID，比如生成 10 个字符的随机 ID：
```go
package main

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" // 自定义字母表
	id, err := gonanoid.Generate(alphabet, 10)
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return
	}

	fmt.Println("Generated ID with length 10:", id)
}

## go-nanoid 的优点
- 高效性：与传统的 UUID 不同，go-nanoid 生成的随机字符串更短、更加高效（默认 21 个字符），但仍然具有唯一性。
- 灵活性：你可以自定义字符集和生成 ID 的长度，以适应不同的使用场景。
- 安全性：基于密码学安全的伪随机数生成器（crypto/rand），在需要唯一性和安全性的场景中非常适用。

## 示例应用场景
- 数据库主键：你可以使用 go-nanoid 生成唯一的字符串作为数据库中的主键，替代传统的自增 ID。
- 短链生成：可以用它来生成短链服务中的唯一链接 ID。
- 文件名生成：如果你需要给上传的文件生成唯一的文件名，可以使用 go-nanoid。
- API 密钥生成：可以用于生成随机的 API 密钥。

## 性能考量
go-nanoid 采用了高效的算法，在大多数场景下生成 ID 的速度是非常快的。默认使用的字符集包含大小写字母和数字，生成的 21 字符长度的随机 ID 足够保证在海量数据下的唯一性。

在性能测试中，nanoid 在生成 100,000 个 ID 时表现出极高的速度，尤其适合需要高效生成唯一标识符的场景。

## 总结
go-nanoid 是一个轻量级但强大的库，能够帮助我们高效地生成唯一随机的 ID。与 UUID 相比，nanoid 的 ID 更短，且足够安全，适合用于多种场景，比如数据库主键、文件名、短链和 API 密钥等。通过灵活的 API，我们可以轻松定制生成的 ID 字符集和长度，从而满足不同的需求。


```