package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ShowQuizController(c *gin.Context) {
	fmt.Println("i am called")
	c.HTML(200, "showquiz.html", nil)
}
