package controllers

import "github.com/gin-gonic/gin"


func AddTeacherController(c *gin.Context){
	c.HTML(200,"addteacher.html",nil)
}