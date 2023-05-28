package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func CreateTeacherController(c *gin.Context) {
	fmt.Println("hello world")
	name := c.PostForm("teachername")
	fmt.Println(name)

	var newuser models.User
	var count int64
	dbconnection.DB.Debug().Where("name=?", name).Find(&newuser).Count(&count)
	fmt.Println(count)
	if count == 0 {
		c.HTML(404, "addteacher.html", gin.H{
			"error": "user not found!!!",
		})
		return
	}
	var tcount int64
	dbconnection.DB.Model(&models.Role{}).Where("role_name=?","teacher").Count(&tcount)
	if tcount!=1{
		dbconnection.DB.Create(&models.Role{Role_name: "teacher"})
	}
	dbconnection.DB.Debug().Model(&models.User{}).Where("name=?",name).Update("role_name","teacher")
	

	c.HTML(200, "addteacher.html", gin.H{
		"error": "Role changed successfully!!!",
	})
}
