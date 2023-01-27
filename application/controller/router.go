package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/", Main)
	r.GET("/create", CreatedTodo)
	r.GET("/detail/:id", ShowDetail)
	r.GET("/update/:id", UpdateDetail)
	r.POST("/post_create", PostTodo)
	r.POST("update/detail/:id", UpdateDataDetail)
	r.GET("delete/detail/:id", DeleteDataDetail)
	return r
}
