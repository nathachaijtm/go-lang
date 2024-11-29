package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API Method
	router.GET("/employee", EmployeeController.GetEmployee)
	router.GET("/employee/:id", EmployeeController.GetEmployeebyID)
	router.POST("/employee", EmployeeController.Postemployee)
	router.PUT("/employee", EmployeeController.Putemployee)
	router.DELETE("/employee", EmployeeController.Deleteemployee)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080/employee/")
}