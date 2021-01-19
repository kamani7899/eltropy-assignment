
package controllers

import (
  "fmt"
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
)

type CreateAccountInput struct {
	AccountType  string `json:"accounttype" binding:"required"`
	CurrentBalance  float32 `json:"currentbalance" binding:"required"`
	Details  string `json:"details" binding:"required"`
	CustomerId uint `json:"customerid"`
	Status bool `json:"status"`
}


func CreateAccount(c *gin.Context) {
	// Validate input
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("input", input)

	currentTime := time.Now()
	
	
	account := models.Account{AccountType: input.AccountType, CurrentBalance: input.CurrentBalance, Status: input.Status, Details: input.Details, CreatedAt: currentTime, CustomerId: input.CustomerId}
	models.DB.Create(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})
}





