package models


type User struct {
  Username     string   `json:"username" gorm:"primary_key"`
  Password  string `json:"password"`
  EmployeeId uint `json:"employeeid" `
  Admin bool `json:"admin" gorm:->`
  
}