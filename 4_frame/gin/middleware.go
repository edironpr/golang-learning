package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func runTime(c *gin.Context) {
	now := time.Now()
	c.Next()
	fmt.Println("用时", time.Since(now))
}

func main() {
	// 创建路由
	// 默认使用 Logger() 和 Recovery() 两个中间件
	r := gin.Default()
	r.Use(runTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopHandler)
		shoppingGroup.GET("/home", shopHandler)
	}
	_ = r.Run(":8000")
}

func shopHandler(c *gin.Context) {
	time.Sleep(1 * time.Second)
	c.JSON(200, gin.H{"message": "SUCCESS", "data": "test"})
}
