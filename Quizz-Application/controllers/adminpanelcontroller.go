package controllers

import "github.com/gin-gonic/gin"

func AdminpanelController(c *gin.Context) {
	c.HTML(200, "adminpanel.html", nil)
}
