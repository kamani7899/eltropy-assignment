package main

import (
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
  "eltropy-assignment/controllers"

)


func main() {
  r := gin.Default()

  models.ConnectDataBase() // new

  

  r.POST("/customers", controllers.CreateCustomer)

  r.POST("/accounts", controllers.CreateAccount)

  r.POST("/employees", controllers.CreateEmployee)

  r.POST("/transactions", controllers.CreateTransaction)


  r.POST("/signup", controllers.SignUp)

  r.POST("/signin", controllers.SignIn)

  r.Run()
}
