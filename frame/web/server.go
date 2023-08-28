package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func main() {

	e := echo.New()

	// 1. HelloWorld
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	// 2. 路由
	e.POST("/saveUser", saveUser)
	e.POST("/saveUser2", saveUser2)
	e.GET("/getUser/:id", getUser)
	e.PUT("/updateUser/:id", updateUser)
	e.DELETE("/deleteUser/:id", deleteUser)

	e.GET("/showName", showName)

	// 7. 处理请求
	// 根据 Content-Type 请求标头将 json，xml，form 或 query 负载绑定到 Go 结构中。
	// 通过状态码将响应渲染为 json 或者 xml 格式。
	e.POST("/createUser", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	// 8. 静态资源
	// 定义/static/*目录为静态资源文件目录
	// http://go-echo.org/guide/static-files
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":8000"))
}

func saveUser2(c echo.Context) error {
	// 6. 表单 multipart/form-data

	// Get name
	name := c.FormValue("name")

	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<b>Thank you! %s</b>", name))
}

func showName(c echo.Context) error {
	// 4. 请求参数
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "The name is "+name)
}

func deleteUser(c echo.Context) error {
	return nil
}

func updateUser(c echo.Context) error {
	return nil
}

func getUser(c echo.Context) error {
	// 3. URL 路径参数

	// User ID 来自于url `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	// 5. 表单 application/x-www-form-urlencoded
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, fmt.Sprintf("The name is %s, and the email is %s", name, email))
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}
