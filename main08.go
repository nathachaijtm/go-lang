package main

import (
	AdminController "backend/api/controller/admin"
	AuthController "backend/api/controller/auth"
	EmployeeController "backend/api/controller/employee"
	"backend/api/db"
	"backend/api/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Get .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//get InitDB fuction
	db.InitDB()

	router := gin.Default()

	authorized := router.Group("/api", middleware.JwtAuthen())

	authorized.GET("/employeedb/:id", EmployeeController.GetEmployeeByID) 
	authorized.GET("/employeedb", EmployeeController.GetEmployeeDB) 

	authorized.POST("/employeedb", EmployeeController.PostEmployeeDB)

	router.POST("/register", AdminController.PostAdmin) 
	router.POST("/login", AuthController.Login)         

	authorized.PUT("/employeedb", EmployeeController.PutEmployeeDB) 

	authorized.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB) 

	//Customer API Method
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
