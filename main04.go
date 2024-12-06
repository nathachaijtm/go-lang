package main

import (
  "fmt"
  EmployeeController "backend/api/controller/employee"
  "github.com/gin-gonic/gin"
  "backend/api/db"
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
  router.GET("/employee", EmployeeController.GetEmployee) //GET
  router.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID
  router.GET("/employeedb", EmployeeController.GetEmployeeDB) //GET FROM DB
  router.POST("/employee", EmployeeController.PostEmployee) //POST
  router.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST TO DB
  router.PUT("/employee", EmployeeController.PutEmployee) //PUT
  router.DELETE("/employee", EmployeeController.DeleteEmployee) //DELETE

  //Customer API Method

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}