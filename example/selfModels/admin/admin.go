package admin

import "github.com/kataras/iris/context"

func ImportStudentsInfoByExcel(c context.Context) {
	c.WriteString("导入学生excel信息成功!")
	c.Next()
}

func ImportCourseStructure(c context.Context) {
	c.WriteString("导入课程结构成功!")
	c.Next()
}
