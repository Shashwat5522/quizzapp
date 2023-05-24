package Router

import (
	"quizz-application/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	Authentication := r.Group("/authentication")
	{
		Authentication.GET("/Registration", controllers.RegistrationController)
		Authentication.POST("/PostRegistration", controllers.PostRegistrationDataController)
		Authentication.GET("/Login", controllers.LoginController)
		Authentication.POST("/Login", controllers.PostLoginDataController)
	}

	Admin:=r.Group("/admin")
	{
		Admin.GET("/adminpanel",controllers.AdminpanelController)
		Admin.GET("/addteacher",controllers.AddTeacherController)
		Admin.GET("/addadmin",controllers.AddAdminController)
		Admin.POST("/addteacher",controllers.CreateTeacherController)
		Admin.POST("/addadmin",controllers.CreateAdminController)
		
	}

	r.Run(":8080")

}
