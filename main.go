package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"net/http"
	"time"
)

func main() {

	//gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 创建一个 cron 实例
	c := cron.New(cron.WithSeconds())

	// 添加定时任务，每隔30s执行一次
	_, err := c.AddFunc("*/30 * * * * *", func() {
		fmt.Println("执行定时任务：", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		return
	}

	// 启动 cron 服务
	c.Start()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
