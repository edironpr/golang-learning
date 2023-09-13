package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare(c *gin.Context) {
	// 获取客户端cookie并校验
	if cookie, err := c.Cookie("token"); err == nil && cookie == "yes" {
		c.Next()
		return
	}
	// 返回错误
	c.JSON(http.StatusUnauthorized, gin.H{"message": "UNAUTHORIZED"})
	// 认证不通过，中止后续函数的调用
	c.Abort()
}

func main() {
	// 创建路由
	r := gin.Default()

	// 登录接口
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("token", "yes", 60, "/", "127.0.0.1", false, true)
		// 返回消息
		c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
	})

	r.GET("/home", AuthMiddleWare, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "SUCCESS", "data": "home"})
	})

	r.Run(":8000")
}
