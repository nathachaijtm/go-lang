package main

import (
	EmployeeController "backend/api/controller/employee"
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
	//Employee API Method
	router.GET("/employees", EmployeeController.GetEmployee)       //GET
	router.GET("/employeesdb", EmployeeController.GetEmployeeDB)   //GET
	router.POST("/employees", EmployeeController.PostEmployee)     //POST
	router.PUT("/employees", EmployeeController.PutEmployee)       //PUT
	router.DELETE("/employees", EmployeeController.DeleteEmployee) //DELETE

	//Customer API Method

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
