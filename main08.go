package main

import (
	AdminController "backend/api/controller/admin"
	EmployeeController "backend/api/controller/employee"
	AuthController "backend/api/controller/auth"
	"backend/api/middleware" 
	"backend/api/db"
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
	authorized.GET("/employeedb", EmployeeController.GetEmployeeDB)

	//Employee API Method
	router.GET("/employee", EmployeeController.GetEmployee)         //GET
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID
	router.GET("/employeedb", AdminController.GetAdmin)
	router.GET("/register", AdminController.GetAdmin)

	router.POST("/employee", EmployeeController.PostEmployee)     //POST
	router.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST TO DB
	router.POST("/register", AdminController.PostAdmin)           //POST TO DB

	router.POST("/login", AuthController.Login) 				  //POST LOGIN

	router.PUT("/employee", EmployeeController.PutEmployee)     //PUT
	router.PUT("/employeedb", EmployeeController.PutEmployeeDB) //PUT TO DB

	router.DELETE("/employee", EmployeeController.DeleteEmployee)         //DELETE
	router.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB) //DELETE DB

	//Customer API Method
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
