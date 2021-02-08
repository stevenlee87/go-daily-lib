package main

import (
	"fmt"
	"net/http"

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

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func nav() map[string]interface{} {
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

	navJson := make(map[string]interface{})
	navJson["error_no"] = 0
	navJson["data"] = navs
	return navJson
}

func course0() map[string]interface{} {

	var courses []Course
	course1 := Course{
		ID: 1, Name: "Python常用模块",
	}
	course2 := Course{
		ID: 2, Name: "Shell脚本编写",
	}
	course3 := Course{
		ID: 3, Name: "Web开发之HTML篇",
	}
	course4 := Course{
		ID: 4, Name: "Python进阶开发",
	}
	course6 := Course{
		ID: 6, Name: "Sketch-UI设计的利器",
	}
	course7 := Course{
		ID: 7, Name: "Git入门",
	}

	courses = append(courses, course1)
	courses = append(courses, course2)
	courses = append(courses, course3)
	courses = append(courses, course4)
	courses = append(courses, course6)
	courses = append(courses, course7)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func course1() map[string]interface{} {

	var courses []Course
	course1 := Course{
		ID: 1, Name: "Python常用模块",
	}

	courses = append(courses, course1)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func course2() map[string]interface{} {

	var courses []Course
	course2 := Course{
		ID: 2, Name: "Shell脚本编写",
	}

	courses = append(courses, course2)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func course3() map[string]interface{} {

	var courses []Course
	course3 := Course{
		ID: 3, Name: "Web开发之HTML篇",
	}
	courses = append(courses, course3)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func course4() map[string]interface{} {

	var courses []Course
	course4 := Course{
		ID: 4, Name: "Python进阶开发",
	}

	courses = append(courses, course4)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func course6() map[string]interface{} {

	var courses []Course
	course6 := Course{
		ID: 6, Name: "Sketch-UI设计的利器",
	}

	courses = append(courses, course6)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func course7() map[string]interface{} {

	var courses []Course
	course7 := Course{
		ID: 7, Name: "Git入门",
	}

	courses = append(courses, course7)

	courseJson := make(map[string]interface{})
	courseJson["error_no"] = 0
	courseJson["data"] = courses
	return courseJson

}

func getRouteParams(c *gin.Context) {
	// 获取路由参数为name的值
	// http://127.0.0.1:8888/route/card  输出 card
	name := c.Param("name")
	c.String(http.StatusOK, name)
}

func getRoutePregParams(c *gin.Context) {

	// 获取路由参数为name的值 和job的值 job为无限长,忽略/,并且包含前一个/
	// http://127.0.0.1:8888/route/card/
	// 输出 card /

	// http://127.0.0.1:8888/route/card/name
	// 输出 card /name

	// http://127.0.0.1:8888/route/card/name/hello
	// 输出 card /name/hello

	name := c.Param("name")
	job := c.Param("job")
	c.String(http.StatusOK, name+" "+job)
}

func GetName(c *gin.Context) {
	//获取name参数, 通过Query获取的参数值是String类型。
	name := c.Query("name")
	fmt.Println("name is", name)
	//获取name参数, 跟Query函数的区别是，可以通过第二个参数设置默认值。
	//name := c.DefaultQuery("name", "sockstack")

	//获取id参数, 通过GetQuery获取的参数值也是String类型,
	// 区别是GetQuery返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在。
	//id, ok := c.GetQuery("id")
	//if !ok {
	//	参数不存在
	//}
}

func print_func(funcName string) {
	fmt.Println(funcName)
}

func fire() {
	fmt.Println("fire")
}

func main() {

	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/category", func(c *gin.Context) {
		c.JSON(200, nav())
	})

	// 获取路由参数
	r.GET("/route/:name", getRouteParams)

	// 获取路由参数
	r.GET("/test/", GetName)

	//r.GET("/courses/", func(c *gin.Context) {
	//	sub_category := c.Query("sub_category")
	//	sub_category_int, _ := strconv.Atoi(sub_category)
	//	if sub_category_int == 1 {
	//		fmt.Println("sub_category is", sub_category)
	//		c.JSON(200, course1())
	//	} else {
	//		c.JSON(200, course0())
	//	}
	//
	//})

	r.GET("/courses/", func(c *gin.Context) {
		sub_category := c.Query("sub_category")

		func_name := "course" + sub_category
		//fmt.Printf("func_name is", func_name)
		var (
			fm = map[string]func() map[string]interface{}{
				"course0": course0,
				"course1": course1,
				"course2": course2,
				"course3": course3,
				"course4": course4,
				"course6": course6,
				"course7": course7,
			}
		)
		c1 := fm[func_name]
		c.JSON(200, c1())

	})

	r.Run(":8090")
}
