package models

import "time"

type Transaction struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  AccountId uint `json:"accountid" `
  EmployeeId uint `json:"employeeid" `
  Merchant string `json:"merchant"`
  amount float32 `json:"amount"`
  IsSuccess bool `json:"issuccess"`
  TransactionTime time.Time `json:"transactiontime"`
  TransactionType string `json:"transactiontype"`
}