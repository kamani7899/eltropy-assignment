package main

import (
  "os"
  "fmt"
  "net/http"
  "strings"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
  "github.com/dgrijalva/jwt-go"
  "eltropy-assignment/controllers"
  "github.com/gin-contrib/sessions"
  "github.com/go-redis/redis/v7"
  

)


var client *redis.Client

func init() {
  //Initializing redis
  dsn := os.Getenv("REDIS_DSN")
  if len(dsn) == 0 {
     dsn = "localhost:6379"
  }
  client = redis.NewClient(&redis.Options{
     Addr: dsn, //redis port
  })
  _, err := client.Ping().Result()
  if err != nil {
     panic(err)
  }
}


func main() {
  r := gin.New()

  r.Use(gin.Logger())

  r.Use(gin.Recovery())

  models.ConnectDataBase() // new

  r.POST("/signup", controllers.SignUp)

  r.POST("/signin", controllers.SignIn)

  r.POST("/signout", controllers.SignOut)


// impose the token for the APIS below

  r.Use(AuthRequired)
    
    
  r.POST("/customers", controllers.CreateCustomer)

  r.POST("/accounts", controllers.CreateAccount)

  r.POST("/employees", controllers.CreateEmployee)

  r.POST("/transactions", controllers.CreateTransaction)

  r.Run()

 
}


func AuthRequired(c *gin.Context)  {
    //validate the token and retun 401 in case of error
    user, err := ExtractTokenMetadata(c.Request)
    if err != nil {
      c.JSON(http.StatusUnauthorized, "unauthorized")
      c.Abort()
      return
    }

    session := sessions.Default(c)
    session.Set("user", user)
    c.Next()
  }


func ExtractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  //normally Authorization:Bearer token 
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
  
  tokenString := ExtractToken(r)
  token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      //need to improve this
     if x, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     }
     
     return []byte(os.Getenv("ACCESS_SECRET")), nil
     
  })
  return token, nil
}

func TokenValid(r *http.Request) error {
  token, err := VerifyToken(r)
  if err != nil {
     return err
  }
  if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
     return err
  }
  return nil
}


func ExtractTokenMetadata(r *http.Request) (string, error) {
  token, err := VerifyToken(r)
  if err != nil {
     return "", err
  }
  claims, _ := token.Claims.(jwt.MapClaims)
  username, ok := claims["username"].(string)
  if !ok || len(username) < 1 {
      return "", fmt.Errorf("invalid token ")
  }
  
  return username, err
}

