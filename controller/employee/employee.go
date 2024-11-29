package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

//GET
func GetEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET ALL Method!",
    })
}

//GET BY ID
func GetEmployeebyID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}

//POST
func Postemployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee POST Method!",
    })
}

func Putemployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee PUT Method!",
    })
}

func Deleteemployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee DELETE Method!",
    })
}