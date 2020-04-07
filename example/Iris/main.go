package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var app = iris.New()

func main() {
	//app := iris.New()
	//app.Logger().SetLevel("debug")
	//// 注册模板
	//app.RegisterView(iris.HTML("./web/views", ".html"))
	//// 注册控制器

	app.Get("/getId", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("request path is:", path)
		username := context.URLParam("username")
		app.Logger().Info("username is: " + username)
		passwd := context.URLParam("password")
		app.Logger().Info("password is: " + passwd)
		type User struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		u := User{
			Username: username,
			Password: passwd,
		}
		context.JSON(u)
		//context.HTML("<h1>" + username + passwd + "</h1>")
		//context.WriteString("username is: " + username + ", password is: " + passwd)
	})
	// handle方法, 大一统, 啦啦啦
	app.Handle("POST", "/postContent", func(c context.Context) {
		//path := c.Path()
		app.Logger().Info(c.URLParams())
	})

	app.Handle("POST", "/postHello", postHello)

	// 正则表达式
	// 查询天气
	// GET http://localhost:8080/weather/2020-04-07/Hangzhou
	app.Handle("GET", "/weather/{date:string}/{city:string}", func(c context.Context) {
		date := c.Params().Get("date")
		city := c.Params().Get("city")
		type Weather struct {
			Date string `json:"date"`
			City string `json:"city"`
		}
		w := &Weather{
			Date: date,
			City: city,
		}
		c.JSON(w) // 返回给client
		//c.WriteString(date+" "+city)
	})

	// 获取bool值
	app.Handle("POST", "/user/{status:bool}", func(c context.Context) {
		status, err := c.Params().GetBool("status")
		if err != nil {
			c.StatusCode(iris.StatusNonAuthoritativeInfo)
		}
		if status {
			c.WriteString("true")
		} else {
			c.WriteString("false")
		}
	})

	app.Run(
		iris.Addr("localhost:8080"),
	)

}

func postHello(c context.Context) {
	app.Logger().Info(c.Path())
}
