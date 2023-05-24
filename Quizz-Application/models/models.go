package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	RoleName string `json:"rolename" form:"rolename"`
	Role     Role   `gorm:"foreignKey:RoleName;References:Role_name"`
}
type Role struct {
	Role_name string `gorm:"primaryKey"`
}