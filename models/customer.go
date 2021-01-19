package models

import "time"

type Customer struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Name  string `json:"name"`
  Age  int `json:"age"`
  Address string `json:"address"`
  Contact string `json:"contact"`
  Kyc string `json:"kyc"`
  Active bool `json:"active"`
  Created time.Time `json:"createdat"`
}