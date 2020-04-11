package main

import (
	"example/selfModels/admin"
	"example/selfModels/students"
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

	// 首先, 学生模块
	studentsParty := app.APIBuilder.Party("/students", func(c context.Context) {
		c.Next()
	})
	studentsParty.Done(func(c context.Context) {
		c.Application().Logger().Info("刚才访问了: ", c.Path())
	})
	// 学生登录部分的验证
	studentsParty.Handle("POST", "/login", students.Login)

	// 学生修改密码部分的验证
	studentsParty.Handle("POST", "/modifyPasswd", students.ModifyPasswd)

	// 控制学生可选课程的部分
	studentsParty.Handle("POST", "/optionalCourse", students.OptionalCourse)

	// 学生已选择课程的部分
	studentsParty.Handle("POST", "/hadCourse", students.HadCourse)

	// 学生获取到课程的详细信息
	studentsParty.Handle("POST", "/courseInfo", students.CourseInfo)

	// 然后, 管理员模块
	adminParty := app.Party("/admin", func(c context.Context) {
		c.Next()
	})
	adminParty.Done(func(c context.Context) {
		c.Application().Logger().Info("admin Just visit:", c.Path())
	})
	// 学生部分信息导入
	adminParty.Handle("POST", "/importStudentsInfo", admin.ImportStudentsInfoByExcel)

	// 课程结构导入
	adminParty.Handle("POST", "/importCourseStructure", admin.ImportCourseStructure)

	// 下面是瞎写的
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

	// 模块化, 分割模块, 用户模块, 管理员模块, 货物模块etc
	// /user/{login/register/info}
	userParty := app.Party("/user", func(c context.Context) {
		c.Next()
	})

	userParty.Done(func(c context.Context) {
		c.Application().Logger().Info("finish run /user/login")
	})
	userParty.Handle("POST", "/login", func(c context.Context) {
		type User struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		type Response struct {
			Code int  `json:"code"`
			User User `json:"user"`
		}
		app.Logger().Info("visit path: ", c.Path())
		//u := new(User)
		u := User{}
		if err := c.ReadJSON(&u); err != nil {
			c.StatusCode(iris.StatusBadRequest)
			c.WriteString(err.Error())
		}
		r := Response{
			Code: 8001,
			User: u,
		}
		c.JSON(r)
		c.Next()
	})

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithOptimizations,
		iris.WithCharset("UTF-8"),
		iris.WithConfiguration(iris.Configuration{
			IgnoreServerErrors:                nil,
			DisableStartupLog:                 false,
			DisableInterruptHandler:           false,
			DisablePathCorrection:             false,
			DisablePathCorrectionRedirection:  false,
			EnablePathEscape:                  false,
			EnableOptimizations:               false,
			FireMethodNotAllowed:              false,
			DisableBodyConsumptionOnUnmarshal: false,
			DisableAutoFireStatusCode:         false,
			TimeFormat:                        "",
			Charset:                           "",
			PostMaxMemory:                     0,
			TranslateFunctionContextKey:       "",
			TranslateLanguageContextKey:       "",
			ViewLayoutContextKey:              "",
			ViewDataContextKey:                "",
			RemoteAddrHeaders:                 nil,
			Other:                             nil,
		}),
	)
}

func postHello(c context.Context) {
	app.Logger().Info(c.Path())
}
