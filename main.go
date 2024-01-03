package main

import (
	"vgblog/global"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库指针
var db *sqlx.DB

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//卧槽，先写这些，把这些实现
	r.POST("/user/login", userLogin)
	r.POST("/user/register", userRegister)
	r.POST("/user/ban", userBan)
	r.POST("/user/avatar", userAvatar)
	r.POST("/admin", admin)
	r.POST("/blog", blogIndex)
	r.POST("/blog/edit", blogEdit)
	r.POST("/blog/read", blogRead)
	r.POST("/blog/comment", blogComment)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

	global.ConnectDB(db)
}

func userLogin(ctx *gin.Context) {

}

func userRegister(ctx *gin.Context) {

}

func userBan(ctx *gin.Context) {

}

func userAvatar(ctx *gin.Context) {

}

func admin(ctx *gin.Context) {

}

func blogIndex(ctx *gin.Context) {

}

func blogEdit(ctx *gin.Context) {

}

func blogRead(ctx *gin.Context) {

}

func blogComment(ctx *gin.Context) {

}
