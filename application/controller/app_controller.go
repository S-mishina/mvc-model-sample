package controller

import (
	"github.com/gin-gonic/gin"
	"mvc-model-sample/application/model"
	"strconv"
	"log"
)

func Main(c *gin.Context) {
	//dbからデータを取得する
	var test = model.GetTodo()
	c.HTML(200, "index.html", gin.H{"todo": test})
}

func CreatedTodo(c *gin.Context) {
	c.HTML(200, "create_todo.html", gin.H{})
}

func PostTodo(c *gin.Context){
	var title 		  = 	c.PostForm("title")
	var body 		  = 	c.PostForm("detail")
	var priority, err = strconv.Atoi(c.PostForm("priority"))
	if err != nil {
		log.Println(err)
		c.Redirect(301, "/create")
	}
	model.CreateTodo(title,body,priority)
	c.Redirect(301, "/")
}

func ShowDetail(c *gin.Context){
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Redirect(301, "/")
	}
	var detail = model.GetDetails(id)
	c.HTML(200, "todo_detail.html", gin.H{"detail": detail})
}

func UpdateDetail(c *gin.Context){
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Redirect(301, "/")
	}
	var detail = model.GetDetails(id)
	c.HTML(200, "update_form.html", gin.H{"detail": detail})
}

func UpdateDataDetail(c *gin.Context){
	var id, iderr = strconv.Atoi(c.Param("id"))
	var title 		  = 	c.PostForm("title")
	var body 		  = 	c.PostForm("detail")
	var priority, err = strconv.Atoi(c.PostForm("priority"))
	if err != nil || iderr != nil {
		c.Redirect(301, "/detail/"+strconv.Itoa(id))
	}
	model.UpdateDetail(id,title,body,priority)
	c.Redirect(301, "/detail/"+strconv.Itoa(id))
}

func  DeleteDataDetail(c *gin.Context){
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Redirect(301, "/")
	}
	model.DeleteDetail(id)
	c.Redirect(301, "/")
}