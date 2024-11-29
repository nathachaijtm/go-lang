package employee

import (
	"net/http"

	"backend/api/db"

	"github.com/gin-gonic/gin"
)

type Tbl_employee struct {
	Emp_id        int
	Emp_firstname string
	Emp_lastname  string
}

func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employees GET Method!",
	})
}

func GetEmployeeDB(c *gin.Context) {
	var employees []Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employees})
}

// GET By ID
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employees POST Method!",
	})
}

func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employees PUT Method!",
	})
}

func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employees DELETE Method!",
	})
}
