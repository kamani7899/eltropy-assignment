
package controllers

import (
  "fmt"
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
)

type CreateCustomerInput struct {
	Name  string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Contact  string `json:"contact" binding:"required"`
	Kyc  string `json:"kyc" binding:"required"`
	Active bool `json:"active"`
	Age int `json:"age" binding:"required"`
}


func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("input", input)

	currentTime := time.Now()
	
	
	customer := models.Customer{Name: input.Name, Age: input.Age, Address: input.Address, Contact: input.Contact, Active: input.Active, Kyc: input.Kyc, Created: currentTime}
	models.DB.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}





