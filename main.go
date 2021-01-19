package main

import (
  
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
  "eltropy-assignment/controllers"
)


func main() {
  r := gin.Default()

  models.ConnectDataBase() // new

  // r.GET("/books", controllers.FindBooks)
  
  // r.POST("/books", controllers.CreateBook)

  // r.DELETE("/books/:id", controllers.DeleteBook)

  // r.PATCH("/books/:id", controllers.UpdateBook)

  // r.GET("/books/:id", controllers.FindBook)

  r.POST("/customers", controllers.CreateCustomer)

  r.POST("/accounts", controllers.CreateAccount)

  r.Run()
}
