package controllers

import "github.com/gin-gonic/gin"

func AddAdminController(c *gin.Context){
	c.HTML(200,"addadmin.html",nil)
}