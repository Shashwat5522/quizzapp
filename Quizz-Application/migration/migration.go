package main

import (
	"quizz-application/dbconnection"
	"quizz-application/models"
)

func init() {
	dbconnection.DBconnection()
}
func main() {

	// dbconnection.DB.Migrator().DropTable(&models.User{},&models.Role{})
	// dbconnection.DB.AutoMigrate(&models.User{}, &models.Role{})
	dbconnection.DB.Migrator().DropTable(&models.Questions{},&models.Quiz{})
	dbconnection.DB.AutoMigrate(&models.Questions{},&models.Quiz{})

	// dbconnection.DB.AutoMigrate(&models.Teacher_Student{})
}
