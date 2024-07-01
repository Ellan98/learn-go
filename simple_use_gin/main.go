package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义登录结构体

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/json", func(ctx *gin.Context) {
		// data := map[string]interface{}{
		// 	"name": "测试",
		// 	"age":  18,
		// }

		data := gin.H{"name": "小王子", "age": 18}
		ctx.JSON(http.StatusOK, data)
	})
	// 结构体内的字段 必须首字母大写 才能 导出使用 因此返回的字段是大写的
	//如果需要小写 可以 给字段 加上tag 即 ``  灵活使用tag 来对数据结构进行定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}

	r.GET("/another_json", func(ctx *gin.Context) {
		data := msg{
			Name:    "小王子",
			Message: "hello golang",
			Age:     18,
		}

		ctx.JSON(http.StatusOK, data)
	})

	// 通过querystring参数
	r.GET("/web", func(ctx *gin.Context) {
		//获取参数 第一种方式
		// name := ctx.Query("query") //通过query 获取请求携带的参数
		// 第二种方式
		// name := ctx.DefaultQuery("queryparam", "somebody")
		// 第三种方式
		name, ok := ctx.GetQuery("ok")
		if !ok {
			name = "somebody"
		}
		ctx.JSON(http.StatusOK, gin.H{"name": name})

	})

	// form 表单提交的参数

	r.POST("/login", func(c *gin.Context) {
		// 第一种 通过 PostForm
		// userName := c.PostForm("userName")
		// passWord := c.PostForm("passWord")
		//第二种 DefaultPostForm
		// userName := c.DefaultPostForm("userName", "somebody")
		// passWord := c.DefaultPostForm("xxx", "***") //如果没有填写 则为 空  可以用作初始化字段
		// 第三种
		userName, ok := c.GetPostForm("userName")

		if !ok {
			userName = "default"
		}

		passWord, ok := c.GetPostForm("passWord")

		if !ok {
			userName = "default"
		}

		c.JSON(http.StatusOK, gin.H{"Name": userName, "Word": passWord})

	})

	// path 参数
	r.GET("/user/search/:username/:address", func(ctx *gin.Context) {
		username := ctx.Param("username")
		address := ctx.Param("address")

		ctx.JSON(http.StatusOK, gin.H{"username": username, "address": address})
	})

	//参数绑定
	r.POST("/loginJSON", func(ctx *gin.Context) {
		var login Login

		if err := ctx.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)

			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}

	})

	// 路由组 通常将路由分组用在划分业务逻辑或划分API版本时。
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"group": "user/index"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"group": "user/login"})
		})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"group": "shop/index"})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"group": "shop/index"})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"group": "shop/index"})
		})
	}

	r.Run(":8080")
}
