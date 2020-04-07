# 抢课系统

1. 学生部分

    * 登录
    * 课程列表
        * 可选课程(从redis里面获取库存)
        * 已选课程(从mqsql里面, 根据课程的id, 获取课程信息)
        * 课程详细信息(从redis里面获取信息)

2. 管理员部分

    * 学生信息部分导入(excel)
    * 课程结构(老师是谁, 此课程的学生总人数, 课程的介绍信息)
    * 打印课程结构(班级+老师+学生)

3. 流程

    管理员:
    * 首先, 管理员录入学生信息(姓名+学院.专业.班级), 并给与每个学生一个id
    * 然后管理员导入课程信息(课程时间, 名称, 授课老师, 学生数量, 可选这门课的班级范围)
    
    学生:
    * 获取可选课程(从redis里面获取, redis里面的值从mysql中拿过来)
    * 点击选课: 
        * 学生token+课程信息{名称: "课程名称", id: 课程标识符}, 发送给服务器, 服务器首先判断学生是否符合抢课条件(已经抢了吗?) 服务器根据课程标识符, (注意加锁sync.mutex)询问reids, 库存是否>0, 如果课程库存大于0, 则redis库存减一, 将用户id+课程id发送至mqtt, 然后写到数据库, 如果库存=0, 则服务器直接返回: 课程已经没有了, 前端在这个过程中, 直接卡窗口10s, 如果http response, 渲染response内容, 如果没有返回, 提示: "请查看已选课程窗口"
        * 学生点击: 已选课程||可选课程, 详情中可以看到: 根据课程id, 可以看到课程的详细信息
    
    管理员:
    * 打印选课结果{成功的结果+失败的结果}

4. 接口设计:

   * 登录(从student_info里面获取信息)
   ```
    method: POST
    url: /login
    request body: // 请求体
        usename: string // 学号
        password: string // 密码
    response: // 响应体
        token: string
        userId: int
   ```

   * 学生密码修改(暂时不做)
   ```
    method: POST
    url: /modifyUserPasswd
    request body: 
        oldPasswd: string
        newPasswd: string
        confirmNewPasswd: string
    response body:
    type resInfo struct:
        code: int
        msg: string
        result: interface{} // optional
   ```

   * 可选课程(从course_info里面获取课程)
   ```
   method: GET
   url: /optionalCourse
   request body: //请求体
    userId: int // 直接传token, 后台根据token去查找userId, 貌似也可以呀 :)
   response body: // 响应体
    courseList: []Course
    type Course struct {
        courseId int
        courseName string
        courseTeacher string
        courseDescription string
        courseImg string
        courseAllNum int
        courseStock int
    }
    // 可选课程的获取条件是: 管理员在导入课程的时候, 课程和班级相互关联. 学生查到班级, 班级查到关联课程, 然后将课程做成: []Course 返回
   ```
   
    * 已选课程(从表格students_info里面获取数据)
    ```
    method: GET
    url: /hadCourse
    request body:
        userId: int
    response body:
        courseList: []Course
        type Course struct {
            courseId int
            courseName string
            courseTeacher string
            courseDescription string
            courseImg string
            courseAllNum int
            courseStock int
        }
    ```

    * 课程详细信息(从表格course_info里面拿到数据)
    ```
    method: GET
    url: /courseInfo
    request body:
        courseId: int
    response body:
    type Course struct {
        courseName string
        courseTeacher string
        courseDescription string
        courseImg string
        courseAllNum int
        courseStock int
    }
    ```

    * 学生信息导入
    ```
    method: POST
    url: /importStudentsInfo
    request body:
        file: studentInfo.xlsx
    response body:
        type studentInfo struct {
            studentName string // 学生姓名
            studentNo string // 学号
            studentPasswd string // 密码, 直接用md5存储特征值
            studentClass string // 学生班级, 例如: "电子142"
            studentDepartment string // 学生所属部门
            studentCollege string // 学生所属学院
        }
    数据库表: students_info
    id 自增主键 (无符号自增int主键)
    studentNo 学号
    studentClass 学生所属班级
    studentDepartment 学生所属学系
    studentCollege string 学生所属学院
    studentCourse string 学生所选课程 例如: "1;34,56" //多种课程之间以;分割
    studentCourseRange string 学生可以选择的课程范围 // 不用第二张表, 怕影响性能
    ```

    * 课程结构导入
    ```
    method: POST
    url: /importCourseStructure
    request body:
        file: courseInfo.xlsx
        // 课程结构包括: 课程名称, 任课老师, 课程介绍, 关联的班级
    response body:
        type CourseInfo struct {
            courseName string 
            courseTeacher string
            courseBrief string
            courseAboutClass string
            courseAllNum int // 课程可以容纳的学生数量
            courseStock int // 课程剩余的可以容纳的学生数量
        }
    课程结构数据库表: course_info
    courseName string
    courseTeacher string
    courseBrief string
    courseAboutClass string
    ```
    上传"课程结构.xlsx"这个文件之后, 遍历一遍course_info这个表, 查到课程和班级的对应关系, 然后更新表: students_info里面studentCourseRange这个键的值







              