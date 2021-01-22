
package controllers

import (
  "fmt"
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
)

type CreateTransactionInput struct {
	AccountId  uint `json:"accountid" binding:"required"`
	EmployeeId  uint `json:"employeeid" binding:"required"`
	Merchant string `json:"merchant"`
	Amount float32 `json:"amount"  binding:"required"`
	TransactionType string `json:"transactiontype" binding:"required"`
}


func CreateTransaction(c *gin.Context) {
	// Validate input
	var input CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("input", input)

	currentTime := time.Now()
	
	
	transaction := models.Transaction{AccountId: input.AccountId, EmployeeId: input.EmployeeId, Merchant: input.Merchant, Amount: input.Amount, TransactionType: input.TransactionType, TransactionTime: currentTime}
	models.DB.Create(&transaction)

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}





