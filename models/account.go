package models

import "time"

type Account struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  AccountType  string `json:"accounttype"`
  Details string `json:"details"`
  CurrentBalance float32 `json:"currentbalance"`
  Status bool `json:"status"`
  CreatedAt time.Time `json:"createdat"`
  CustomerId uint `json:"customerid" `
}