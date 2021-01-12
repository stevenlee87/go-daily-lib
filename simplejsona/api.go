package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenlee87/go-daily-lib/simplejsona/middleware"
)

//type Status struct {
//	ErrorNo int `json:"error_no"`
//}

type Navigation struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category int    `json:"category"`
}

func main() {
	//navs := []Navigation{{ID: 1, Name: "Python", Category: 1}, {ID: 2, Name: "Linux基础", Category: 4},
	//	{ID: 3, Name: "前端", Category: 2}, {ID: 4, Name: "Python进阶", Category: 1}, {ID: 6, Name: "UI", Category: 5},
	//	{ID: 7, Name: "工具", Category: 1}}

	var navs []Navigation
	nav1 := Navigation{
		ID: 1, Name: "Python", Category: 1,
	}
	nav2 := Navigation{
		ID: 2, Name: "Linux基础", Category: 4,
	}
	nav3 := Navigation{
		ID: 3, Name: "前端", Category: 2,
	}
	nav4 := Navigation{
		ID: 4, Name: "Python进阶", Category: 1,
	}
	nav6 := Navigation{
		ID: 6, Name: "UI", Category: 5,
	}
	nav7 := Navigation{
		ID: 7, Name: "工具", Category: 1,
	}

	navs = append(navs, nav1)
	navs = append(navs, nav2)
	navs = append(navs, nav3)
	navs = append(navs, nav4)
	navs = append(navs, nav6)
	navs = append(navs, nav7)

	cnnJson := make(map[string]interface{})
	cnnJson["error_no"] = 0
	cnnJson["data"] = navs

	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/category", func(c *gin.Context) {
		c.JSON(200, cnnJson)
	})
	r.Run(":8080")
}
