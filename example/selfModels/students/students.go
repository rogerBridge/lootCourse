package students

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type StudentLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type StudentLoginResponse struct {
	Code    int                 `json:"code"`
	Student StudentLoginRequest `json:"student"`
}

func Login(c context.Context) {
	//c.Application().Logger().Info("Login ...")
	s := StudentLoginRequest{}
	if err := c.ReadJSON(&s); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
	}
	r := StudentLoginResponse{
		Code:    8001,
		Student: s,
	}
	c.JSON(r)
	c.Next()
}

func ModifyPasswd(c context.Context) {
	c.WriteString("你刚刚已经修改了密码")
	c.Next()
}

func OptionalCourse(c context.Context) {
	c.WriteString("获取到的可选课程表格是:")
	c.Next()
}

func HadCourse(c context.Context) {
	c.WriteString("获取到已选择的课程表格是:")
	c.Next()
}

func CourseInfo(c context.Context) {
	c.WriteString("获取到课程的详细信息")
	c.Next()
}
