package models

import "time"


type Employee struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  AccountId uint `json:"accountid" `	
  Name  string `json:"name"`
  Age  string `json:"age"`
  Address string `json:"address"`
  Contact string `json:"contact"`
  Salary float32 `json:"salary"`
  Active bool `json:"active"`
  Admin bool `json:"admin"`
  CreatedAt time.Time `json:"createdat"`
}