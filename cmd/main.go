package main

import (
	"first/api"
	"first/utils"
)

func main() {
	utils.ConnectDB()
	r := api.NewRouter()
	r.Spin()
}

//
//package main
//
//import (
//	"context"
//	"database/sql"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/app/server"
//	"github.com/cloudwego/hertz/pkg/common/json"
//	"github.com/cloudwego/hertz/pkg/common/utils"
//	"github.com/cloudwego/hertz/pkg/protocol/consts"
//	_ "github.com/go-sql-driver/mysql"
//	"log"
//	"net/http"
//)
//
//type Student struct {
//	Studentid string  `json:"studentid"`
//	Name      string  `json:"name"`
//	Age       int     `json:"age"`
//	Score     float64 `json:"score"`
//}
//type User struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//
//// var user map[string]User
//var students map[string]*Student
//var Db *sql.DB
//
//func main() {
//	h := server.New()
//	dns := "root:dzf244106_F@tcp(127.0.0.1:3306)/mytest"
//	var err error
//	Db, err = sql.Open("mysql", dns)
//	if err != nil {
//		panic(err)
//	}
//	err = Db.Ping()
//	if err != nil {
//		panic(err)
//	}
//	students = make(map[string]*Student)
//	a := func(ctx context.Context, c *app.RequestContext) {
//		path := c.Request.Path()
//		if string(path) == "/register" {
//			c.Next(ctx)
//			return
//		}
//		username, password, ok := c.Request.BasicAuth()
//		if !ok {
//			c.JSON(http.StatusUnauthorized, utils.H{
//				"massage": "401",
//			})
//			return
//		}
//		var usermassage User
//		err := Db.QueryRow("select password from usermassage where account=?", username).Scan(&usermassage.Password)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, utils.H{
//				"massage": "401",
//			})
//			return
//		}
//		if usermassage.Password != password {
//			c.JSON(http.StatusUnauthorized, utils.H{
//				"massage": "密码错误",
//			})
//			return
//		}
//		c.Next(ctx)
//	}
//	h.Use(a)
//
//	h.GET("/search", searchstudent)
//	h.POST("/profile", updateSF)
//	h.POST("/add", addstudent)
//	h.DELETE("/delete", deleteStudent)
//	h.POST("/register", register)
//	h.Spin()
//}
//
//func register(ctx context.Context, c *app.RequestContext) {
//	puser := &User{}
//	date, err := c.Body()
//	if err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"massage": "400",
//		})
//		return
//	}
//	if err = json.Unmarshal(date, puser); err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"massage": "400",
//		})
//		return
//	}
//	tx, err := Db.Begin()
//	if err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"massage": "400",
//		})
//		return
//	}
//	defer tx.Rollback()
//	var count int
//	err = tx.QueryRow("select count(*) from usermassage where account=?", puser.Username).Scan(&count)
//	if err != nil || count > 0 {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"massage": "400",
//		})
//		return
//	}
//	_, err = tx.Exec("insert into usermassage (account,password) values (?, ?)", puser.Username, puser.Password)
//	if err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"massage": "400",
//		})
//		return
//	}
//	err = tx.Commit()
//	if err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"massage": "400",
//		})
//		return
//	}
//	c.JSON(consts.StatusOK, utils.H{
//		"massage": "200",
//	})
//}
//
//func addstudent(ctx context.Context, c *app.RequestContext) {
//	nstudent := &Student{}
//	date, err := c.Body()
//	if err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"request": "无效的请求数据",
//		})
//	}
//
//	if err = json.Unmarshal(date, nstudent); err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"request": "无效的请求数据",
//		})
//		return
//	}
//	if _, ok := students[nstudent.Studentid]; ok {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"request": "数据已存在",
//		})
//		return
//	}
//	tx, err := Db.Begin()
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "操作失败",
//		})
//		return
//	}
//	defer tx.Rollback()
//	var cid int
//	err = tx.QueryRow("select count(*) from student where id=?", nstudent.Studentid).Scan(&cid)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//	}
//	if cid > 0 {
//		c.JSON(consts.StatusConflict, utils.H{
//			"massage": "409",
//		})
//		return
//	}
//	_, err = tx.Exec("insert into student (id,name,age,score) values (?,?,?,?)",
//		nstudent.Studentid, nstudent.Name, nstudent.Age, nstudent.Score)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	err = tx.Commit()
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	c.JSON(consts.StatusCreated, utils.H{
//		"massge": "201",
//	})
//}
//
//func updateSF(ctx context.Context, c *app.RequestContext) {
//	studentdate := &Student{}
//	date, err := c.Body()
//	if err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"resquest": "无效的请求数据",
//		})
//	}
//	if err := json.Unmarshal(date, studentdate); err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"resquest": "无效的请求数据",
//		})
//		return
//	}
//	tx, err := Db.Begin()
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	defer tx.Rollback()
//	var cid int
//	err = tx.QueryRow("select count(*) from student where id=?", studentdate.Studentid).Scan(&cid)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	if cid != 1 {
//		c.JSON(consts.StatusNotFound, utils.H{
//			"massage": "404",
//		})
//		return
//	}
//	_, err = tx.Exec("update student set name=?,age=?,score=? where id=?", studentdate.Name, studentdate.Age, studentdate.Score, studentdate.Studentid)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	err = tx.Commit()
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	c.JSON(consts.StatusOK, utils.H{
//		"message": "数据更新成功",
//	})
//}
//
//func searchstudent(ctx context.Context, c *app.RequestContext) {
//	idstudents := c.Query("studentid")
//	var studentmassdge Student
//	err := Db.QueryRow("select * from student where id=?", idstudents).Scan(&studentmassdge.Studentid, &studentmassdge.Name, &studentmassdge.Score, &studentmassdge.Age)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	c.JSON(consts.StatusOK, studentmassdge)
//}
//
//func deleteStudent(ctx context.Context, c *app.RequestContext) {
//	id := c.Query("studentid")
//	tx, err := Db.Begin()
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	defer tx.Rollback()
//	var cid int
//	err = tx.QueryRow("select count(*) from student where id=?", id).Scan(&cid)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	if cid == 0 {
//		c.JSON(consts.StatusNotFound, utils.H{
//			"massage": "404",
//		})
//		return
//	}
//	_, err = tx.Exec("delete from student where id=?", id)
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	err = tx.Commit()
//	if err != nil {
//		c.JSON(consts.StatusInternalServerError, utils.H{
//			"massage": "500",
//		})
//		return
//	}
//	c.JSON(consts.StatusOK, utils.H{
//		"massage": "200",
//	})
//}
//
//func loginMiddleware(next app.HandlerFunc) app.HandlerFunc {
//	return func(ctx context.Context, c *app.RequestContext) {
//		username, password, ok := c.Request.BasicAuth()
//		if !ok {
//			c.JSON(http.StatusUnauthorized, utils.H{
//				"massage": "401",
//			})
//		}
//		log.Println(username, password)
//		var usermassage User
//		err := Db.QueryRow("select password from usermassge where account=?", username).Scan(&usermassage.Password)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, utils.H{
//				"massage": "401",
//			})
//			return
//		}
//		if usermassage.Password != password {
//			c.JSON(http.StatusUnauthorized, utils.H{
//				"massage": "密码错误",
//			})
//			return
//		}
//		next(ctx, c)
//	}
//}

//package main
//
//import (
//	"context"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/app/server"
//	"github.com/cloudwego/hertz/pkg/common/utils"
//	"github.com/cloudwego/hertz/pkg/protocol/consts"
//)
//
//type Student struct {
//	StudentID string `json:"student_id"`
//	Birthday  string `json:"birthday"`
//	Gender    string `json:"gender"`
//}
//
//var students map[string]Student
//
//func main() {
//	students = make(map[string]Student)
//
//	h := server.Default()
//
//	h.POST("/add", addStudent)
//
//	h.POST("/profile", updateStudentProfile)
//
//	h.GET("/search", searchStudent)
//
//	h.Spin()
//}
//
//func addStudent(ctx context.Context, c *app.RequestContext) {
//	var newStudent Student
//
//	if err := c.BindAndValidate(&newStudent); err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"message": "无效的请求数据",
//		})
//		return
//	}
//
//	if _, ok := students[newStudent.StudentID]; ok {
//		c.JSON(consts.StatusConflict, utils.H{
//			"message": "该学号的学生已存在",
//		})
//		return
//	}
//
//	students[newStudent.StudentID] = newStudent
//	c.JSON(consts.StatusCreated, utils.H{
//		"message": "学生添加成功",
//	})
//}
//
//func updateStudentProfile(ctx context.Context, c *app.RequestContext) {
//	var updateData Student
//
//	if err := c.BindAndValidate(&updateData); err != nil {
//		c.JSON(consts.StatusBadRequest, utils.H{
//			"message": "无效的请求数据",
//		})
//		return
//	}
//
//	if _, ok := students[updateData.StudentID]; !ok {
//		c.JSON(consts.StatusNotFound, utils.H{
//			"message": "未找到该学号的学生",
//		})
//		return
//	}
//
//	students[updateData.StudentID] = updateData
//	c.JSON(consts.StatusOK, utils.H{
//		"message": "学生信息更新成功",
//	})
//}
//
//func searchStudent(ctx context.Context, c *app.RequestContext) {
//	massagerequest()
//	studentID := c.Query("student_id")
//
//	if student, ok := students[studentID]; ok {
//		c.JSON(consts.StatusOK, utils.H{
//			"student": student,
//		})
//	} else {
//		c.JSON(consts.StatusNotFound, utils.H{
//			"message": "未找到该学号的学生",
//		})
//	}
//}
//
//func massagerequest() app.HandlerFunc {
//	return func(c context.Context, ctx *app.RequestContext) {
//		token := ctx.GetHeader("Authorization")
//		password := string(token)
//		pass := password == "123456"
//		if pass {
//			ctx.Next(c)
//		} else {
//			ctx.JSON(consts.StatusUnauthorized, utils.H{
//				"error": "Unauthorized",
//			})
//			ctx.Abort()
//		}
//
//	}
//}

//package main
//
//import (
//	"context"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/app/server"
//	"net/http"
//)
//
//func main() {
//
//	h := server.New()
//	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
//		ctx.JSON(http.StatusOK, map[string]string{"response": "pong!"})
//	})
//
//	h.GET("/echo", func(c context.Context, ctx *app.RequestContext) {
//		message := ctx.Query("message")
//		ctx.JSON(http.StatusOK, map[string]string{"response": message})
//	})
//	h.Spin()
//}

//package main
//
//import (
//	"context"
//	"net/http"
//
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/app/server"
//)
//
//func main() {
//	// 创建Hertz服务器实例
//	h := server.New()
//
//	// 定义处理 /ping 路由的函数
//	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
//		ctx.JSON(http.StatusOK, map[string]string{"response": "pong!"})
//	})
//
//	// 定义处理 /echo 路由的函数
//	h.GET("/echo", func(c context.Context, ctx *app.RequestContext) {
//		// 获取查询参数中的message值
//		message := ctx.Query("message")
//		ctx.JSON(http.StatusOK, map[string]string{"response": message})
//	})
//
//	// 在指定端口启动服务
//	h.ListenAndServe(":8888")
//}

//package main

//func main() {
//	// 创建Hertz服务器实例
//	h := server.New()
//
//	// 定义处理 /ping 路由的函数
//	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
//		// 直接响应 "pong!"
//		ctx.String(http.StatusOK, "pong!")
//	})
//
//	// 定义处理 /echo 路由的函数
//	h.GET("/echo", func(c context.Context, ctx *app.RequestContext) {
//		// 获取查询参数中的message值
//		message := ctx.Query("message")
//		// 响应查询参数传入的内容
//		ctx.String(http.StatusOK, message)
//	})
//
//	// 在指定端口启动服务
//	h.ListenAndServe(":8000")
//}

//package main
//
//import (
//	"context"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/app/server"
//	"net/http"
//)
//
//func main() {
//	// 创建Hertz服务器实例
//	h := server.New()
//
//	// 定义处理 /ping 路由的函数
//	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
//		ctx.JSON(http.StatusOK, map[string]string{"response": "pong!"})
//	})
//
//	// 定义处理 /echo 路由的函数
//	h.GET("/echo", func(c context.Context, ctx *app.RequestContext) {
//		// 获取查询参数中的message值
//		message := ctx.Query("message")
//		ctx.JSON(http.StatusOK, map[string]string{"response": message})
//	})
//
//	// 在指定端口启动服务
//	h.ListenAndServe(":8888")
//}

//
//import "fmt"
//
//type Jiekou interface {
//	Qiuhe() int
//	Qiuji() int
//}
//
//type mys int
//
//func (s mys) Qiuhe() (n int) {
//	n = 0
//	for i := 1; i <= int(s); i++ {
//		n += i
//	}
//	return n
//}
//func (s mys) Qiuji() (n int) {
//	n = 1
//	for i := 1; i <= int(s); i++ {
//		n *= i
//	}
//	return n
//}
//func jh(r Jiekou) {
//	fmt.Println(r.Qiuhe())
//}
//func qj(r Jiekou) {
//	fmt.Println(r.Qiuji())
//}
//func main() {
//	var jsz mys = 5
//	//var z Jiekou = jsz
//	jh(jsz)
//	qj(jsz)
//}

//package main
//
//import (
//	"fmt"
//)
//
//// 1. 定义接口
//// Shape接口定义了计算面积和周长的方法 c
//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}
//
//// Drawable接口定义了绘制图形的方法
//type Drawable interface {
//	Draw()
//}
//
//// 2. 结构体定义及实现接口
//// Rectangle结构体表示矩形
//type Rectangle struct {
//	width  float64
//	height float64
//}
//
//// 实现Shape接口的Area方法
//func (r Rectangle) Area() float64 {
//	return r.width * r.height
//}
//
//// 实现Shape接口的Perimeter方法
//func (r Rectangle) Perimeter() float64 {
//	return 2 * (r.width + r.height)
//}
//
//// 实现Drawable接口的Draw方法
//func (r Rectangle) Draw() {
//	fmt.Printf("绘制一个矩形，宽为%.2f，高为%.2f\n", r.width, r.height)
//}
//
//// Circle结构体表示圆形
//type Circle struct {
//	radius float64
//}
//
//// 实现Shape接口的Area方法
//func (c Circle) Area() float64 {
//	return 3.14 * c.radius * c.radius
//}
//
//// 实现Shape接口的Perimeter方法
//func (c Circle) Perimeter() float64 {
//	return 2 * 3.14 * c.radius
//}
//
//// 实现Drawable接口的Draw方法
//func (c Circle) Draw() {
//	fmt.Printf("绘制一个圆形，半径为%.2f\n", c.radius)
//}
//
//// 3. 函数使用接口作为参数
//// PrintShapeInfo函数接受Shape接口类型的参数，用于打印图形的面积和周长信息
//func PrintShapeInfo(s Shape) {
//	fmt.Printf("图形的面积是：%.2f\n", s.Area())
//	fmt.Printf("图形的周长是：%.2f\n", s.Perimeter())
//}
//func DrawShape(s Drawable) {
//	s.Draw()
//}
//
//// DrawShapes函数接受Drawable接口类型的参数切片，用于绘制多个图形
//func DrawShapes(ds []Drawable) {
//	for _, d := range ds {
//		d.Draw()
//	}
//}
//
//// 4. 函数使用接口作为返回值
//// NewShape函数根据传入的类型标识创建并返回相应的Shape接口类型的对象
//func NewShape(shapeType string) (a Shape, b Drawable) {
//	if shapeType == "rectangle" {
//		return Rectangle{width: 5, height: 3}, Rectangle{width: 5, height: 3}
//	} else if shapeType == "circle" {
//		return Circle{radius: 2}, Circle{radius: 2}
//	}
//	return nil, nil
//}
//
//func main() {
//	// 5. 在main函数中使用接口
//	// 创建矩形和圆形结构体实例
//	rectangle := Rectangle{width: 4, height: 6}
//	circle := Circle{radius: 3}
//	// 将矩形和圆形实例赋值给Shape接口类型变量
//	//var shape1 Shape = rectangle
//	var shape2 Shape = circle
//
//	// 调用PrintShapeInfo函数打印图形信息
//	PrintShapeInfo(rectangle)
//	PrintShapeInfo(shape2)
//	// 创建可绘制图形的切片
//	drawables := []Drawable{rectangle, circle}
//
//	// 调用DrawShapes函数绘制图形
//	DrawShapes(drawables)
//	DrawShape(rectangle)
//	DrawShape(circle)
//	// 通过NewShape函数创建Shape接口类型的对象并打印信息
//	newShape, newshape := NewShape("rectangle")
//	if newShape != nil {
//		PrintShapeInfo(newShape)
//		DrawShape(newshape)
//	}
//
//	newShape, newshape = NewShape("circle")
//	if newShape != nil {
//		PrintShapeInfo(newShape)
//		DrawShape(newshape)
//
//	}
//}

//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//// ping 响应函数
//
//func ping(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "pong!,abc")
//}
//
//func main() {
//	http.HandleFunc("/ping", ping)    // 创建路由
//	http.ListenAndServe(":8000", nil) // 监听端口及启动服务
//
//}

//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//func ping(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "pong!,abc")
//}
//func main() {
//	http.HandleFunc("/ping", ping)
//	http.ListenAndServe(":8000", nil)
//	// 定义请求的URL
//	url := "http://127.0.0.1:8000/ping"
//	// 创建一个HTTP客户端
//	client := &http.Client{}
//	// 发送GET请求
//	resp, err := client.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close() // 确保在函数返回时关闭响应体
//	// 读取响应内容
//	body, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//		panic(err)
//	}
//	// 打印响应内容
//	fmt.Println(string(body))
//}

//package main
//
//import "fmt"
//
//func Transpose(a [][]int, m int, n int) {
//	for i := 0; i < m; i++ {
//		for j := i; j < n; j++ {
//			a[i][j], a[j][i] = a[j][i], a[i][j]
//		}
//	}
//}
//func main() {
//	var m, n, z int
//	fmt.Scan(&m, &n)
//	if m > n {
//		z = m
//	} else {
//		z = n
//	}
//	a := make([][]int, z)
//	for i := range a {
//		a[i] = make([]int, z)
//	}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			fmt.Scan(&a[i][j])
//		}
//	}
//	Transpose(a, m, n)
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			fmt.Printf("%d\t", a[i][j])
//		}
//		fmt.Println()
//	}
//
//}

//func main() {
//	var n int
//	fmt.Scan(&n)
//	a := make([][]int, n)
//	for i := range a {
//		a[i] = make([]int, n)
//	}
//	for j := 0; j < n; j++ {
//		for i := 0; i < n; i++ {
//			a[j][i] = j*n + i + 1
//		}
//	}
//	for j := 0; j < n; j++ {
//		for i := 0; i < n; i++ {
//			a[j][i], a[i][j] = a[i][j], a[j][i]
//		}
//	}
//	for j := 0; j < n; j++ {
//		for i := 0; i < n; i++ {
//			fmt.Printf("%4d ", a[i][j])
//		}
//		fmt.Println()
//	}
//
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//	"strings"
//	"time"
//)
//
//func main() {
//	// 打开文件
//	file, err := os.OpenFile("1.txt", os.O_RDONLY, 0644)
//	if err != nil {
//		fmt.Println("打开文件出错：", err)
//		return
//	}
//	defer file.Close()
//
//	// 当前日期
//	now := time.Now()
//
//	// 用于存储最近事件的相关信息
//	var closestEvent string
//	var closestDiff time.Duration
//
//	// 创建读取文件的缓冲读取器
//	reader := bufio.NewReader(file)
//
//	// 逐行读取文件内容
//	for {
//		line, err := reader.ReadString('\n')
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			fmt.Println("读取文件出错：", err)
//			return
//		}
//
//		// 去除换行符
//		line = strings.TrimSuffix(line, "\n")
//
//		// 分割日期和事件
//		parts := strings.Split(line, " ")
//		if len(parts) != 2 {
//			continue
//		}
//
//		eventDateStr := parts[0]
//		eventName := parts[1]
//
//		// 解析日期
//		eventDate, err := time.Parse("2006-01-02", eventDateStr)
//		if err != nil {
//			fmt.Println("解析日期出错：", err)
//			continue
//		}
//
//		// 计算当前日期与事件日期的时间差
//		diff := eventDate.Sub(now)
//		if diff < 0 {
//			diff = -diff
//		}
//		if closestEvent == "" || diff < closestDiff {
//			closestEvent = eventName
//			closestDiff = diff
//			if diff < 0 {
//				diff = -diff
//			}
//		}
//	}
//
//	if closestEvent != "" {
//		fmt.Printf("最近的一个事件是：%s\n", closestEvent)
//		fmt.Printf("还有 %d天\n", int(closestDiff.Hours()/24))
//	} else {
//		fmt.Println("没有找到未来的事件。")
//	}
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	// 打开一个只读的文件
//	f, err := os.OpenFile("D:/GoProjects/src/Project_2/test.txt", os.O_RDONLY, 0)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close() // 延迟关闭
//	// 读取文件内容
//	b := make([]byte, 1024)
//	n, err := f.Read(b)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(b[:n]))
//}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Person struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//func main() {
//	jsonData := `{"name":"Jane Doe","age":25}`
//	var p Person
//	err := json.Unmarshal([]byte(jsonData), &p)
//	if err != nil {
//		// 处理错误
//	}
//	fmt.Println(p.Name)
//}

//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Person struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//func main() {
//	p := Person{Name: "John Doe", Age: 30}
//	jsonData, err := json.Marshal(p)
//	if err != nil {
//		// 处理错误
//	}
//	fmt.Println(string(jsonData)) // 输出: {"name":"John Doe","age":30}
//}

/*package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 98
	str := string(x)
	fmt.Println(str)
	fmt.Println("------")
	// 字符串转换为整型
	str = "15"
	fmt.Println(strconv.Atoi(str))
	// 整型转换为字符串
	x = 85
	fmt.Println(strconv.Itoa(x))
}
*/
//
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	s := "Hello,世界"
//	// 判断字符串是否包含某个字串
//	fmt.Println(strings.Contains(s, "世界"))
//	// 判断字符串是否以某个字串开头或结尾
//	fmt.Println(strings.HasPrefix(s, "Hello"))
//	fmt.Println(strings.HasSuffix(s, "Hello"))
//	// 统计字符串中某个字串出现的次数  *
//	fmt.Println(strings.Index(s, "l"))
//	// 替换字符串中某个字串为另一个字串
//	fmt.Println(strings.Replace(s, "世界", "world", -1))
//	// 将字符串全部转换为大写或者小写
//	fmt.Println(strings.ToUpper(s))
//	fmt.Println(strings.ToLower(s))
//	// 将字符串按照某个分隔符分隔为一个切片
//	b := strings.Split(s, ",")
//	fmt.Println(b)
//	// 将一个切片按照某个分隔符拼接为一个字符串
//	fmt.Println(strings.Join([]string{"a", "b", "c"}, " "))
//	// 返回将字符串按照空白分割的多个字符串。
//	// 如果字符串全部是空白或者是空字符串的话，会返回空切片。
//	a := strings.Fields(s)
//	fmt.Println(a)
//}

//
////package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	//注意传入和返回参数类型一般都是float64（int类型自己强转一下）
//	// math.Max math.Min 对比两个数字取最大值或最小值，
//	x, y := 1.1, 2.2
//	fmt.Println(math.Max(x, y))
//
//	fmt.Println(math.Min(x, y))
//	// math.Abs 取绝对值
//	z := -1.3
//	fmt.Println(math.Abs(z))
//	// math.Sqrt 返回x的二次方根
//	x = 2.0
//	fmt.Println(math.Sqrt(x))
//	// math.Pow 返回x^y
//	x, y = 2.0, 3.0
//	fmt.Println(math.Pow(x, y))
//}

//package main
//
//import "fmt"
//
//func main() {
//	s := "Hello, 世界"
//	fmt.Println(s)
//	b := []byte(s)
//	fmt.Println(b)
//	b[7] = 230
//	fmt.Println(b)
//	s = string(b)
//	fmt.Println(s)
//}

//package main
//
//func main() {
//	//
//
//}

//func main() {
//	type student struct {
//		Name  string
//		Age   int
//		Score float64
//	}
//	stuent_s := []student{
//		{"张三", 18, 90},
//		{"lisi", 19, 85.5},
//		{"wangwu", 19, 87.5},
//	}
//	for _, stuent := range stuent_s {
//		fmt.Println(stuent)
//	}
//}

//type People struct {
//	Name  string
//	Age   int
//	Books []Book
//}
//type Book struct {
//	Name string
//}
//
//func (w People) PrintName() {
//	fmt.Println(w.Name)
//}
//func (w People) PrintAge() {
//	fmt.Println(w.Age)
//}
//func (w People) PrintBook() {
//	for _, book := range w.Books {
//		fmt.Println(book.Name) // 打印每个Book的Name字段
//	}
//}
//func (b Book) PrintBookName() {
//	fmt.Println(b.Name)
//}
//func main() {
//	a := People{}
//	c := Book{}
//	a.Name = "abc"
//	a.Age = 18
//	c.Name = "po"
//	a.Books = append(a.Books, Book{Name: "POO"})
//	a.PrintName()
//	a.PrintAge()
//	a.PrintBook()
//	c.PrintBookName()
//}
//
////type person struct {
////	name string
////	city string
////	age  int
////}
//
//type Book
//
//func main() {
//	//var p1 person
//	//p1.name = "DYT"
//	//p1.city = "重庆"
//	//p1.age = 18
//	//
//	//fmt.Printf("p1=%v\n", p1.age)
//	//fmt.Printf("p1=%#v\n", p1)
//	//
//	//var p2 = new(person)
//	//fmt.Printf("%T\n", p2)
//	//fmt.Printf("p2=%v\n", &p2)
//	//fmt.Printf("p2=%#v\n", p2)
//	//p2.name = "DYT"
//	//p2.age = 18
//	//p2.city = "CQ"
//	//fmt.Printf("p2=%#v\n", p2)
//
//	//p5 := person{
//	//	name: "DYT",
//	//	city: "CQ",
//	//	age:  18,
//	//}
//	//fmt.Printf("p5=%#v\n", p5)
//	//
//	//var a *int //指针——>引用类型
//	//a = new(int)
//	//*a = 10
//	//fmt.Println(*a)
//
//	type People struct {
//		Name string
//		Age int
//		Books []Book
//	}
//	type Book struct {
//		Name string
//	}
//	func (w People) PrintName() {
//		fmt.Println(w.Name)
//	}
//	func (w People) PrintAge() {
//		fmt.Println(w.Age)
//	}
//	func (w People) PrintBook() {
//		fmt.Println(w.Books)
//	}
//	func (b Book) PrintBookName() {
//		fmt.Println(b.Name)
//	}
