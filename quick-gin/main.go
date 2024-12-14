package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// 路径参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Query 参数
	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, "Your ID is %s, %s", id, name)
	})

	// Form 参数
	// curl http://localhost:9999/form  -X POST -d 'username=Bob'
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// Query、Form 混合参数
	// curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// Map 参数（字典参数）
	// curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}

	v1 := r.Group("/v1")
	v1.GET("/posts", defaultHandler)
	v1.GET("/series", defaultHandler)

	v2 := r.Group("/v2")
	v2.GET("/posts", defaultHandler)
	v2.GET("/series", defaultHandler)

	// 上传文件
	r.POST("/uploadSingle", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	r.POST("/uploadMulti", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}

		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	// HTML 模板
	type student struct {
		Name string
		Age  int
	}

	r.LoadHTMLGlob("templates/*")

	s1 := &student{Name: "Jack", Age: 25}
	s2 := &student{Name: "Alice", Age: 23}
	r.GET("/students", func(c *gin.Context) {
		c.HTML(http.StatusOK, "students.tmpl", gin.H{
			"title":    "Students",
			"students": []*student{s1, s2},
		})
	})

	// 中间件
	// 作用于全局
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 作用于单个路由
	r.GET("/benchmark", LatencyLogger())

	// 作用于组
	auth := r.Group("/auth")
	auth.Use(AuthRequired())
	{
		auth.POST("/login", func(c *gin.Context) {})
		auth.POST("/submit", func(c *gin.Context) {})
	}

	// 热加载调试 Hot Reload
	// go get -v -u github.com/pilu/fresh

	// 运行
	err := r.Run(":9999")
	if err != nil {
		return
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func LatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("test", "1")
		c.Next()
		latency := time.Since(start)
		log.Print(latency)
	}
}
