
package controllers

import (
  "fmt"
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
)

type CreateEmployeeInput struct {
	AccountId  uint `json:"accountid" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Age int `json:"age" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Contact  string `json:"contact" binding:"required"`
	Salary  float32 `json:"salary" binding:"required"`
	
}


func CreateEmployee(c *gin.Context) {
	// Validate input
	var input CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("input", input)

	currentTime := time.Now()
	
	
	employee := models.Employee{AccountId: input.AccountId, Age: input.Age, Address: input.Address, Contact: input.Contact, Name: input.Name, Salary: input.Salary, CreatedAt: currentTime, Active: true, Admin: false}
	models.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}





