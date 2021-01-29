package controllers

import (
  "fmt"
  "net/http"
  "os"
  "time"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
  "golang.org/x/crypto/bcrypt"
  "github.com/gin-contrib/sessions"
  "github.com/dgrijalva/jwt-go"
  "github.com/go-redis/redis/v7"
)

var jwtKey = []byte("my_secret_key")

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

type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	EmployeeId uint `json:"employeeid" `
}


type SigninCredentials struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type CredsWithToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


func SignUp(c *gin.Context) {
	// Validate input
	
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	//Encrypt the password before storing in the database
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	
	
	user := models.User{Username: input.Username, Password: string(hashedPassword), EmployeeId: 1}
	
	//store the user
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func SignIn(c *gin.Context) {
	

	
	var input SigninCredentials
	if err := c.ShouldBindJSON(&input); err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	

	var user models.User


	x :=models.DB.Where("username = ?", input.Username).First(&user)

	


	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    	return
  	}

  
	
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	


	tokenInfo, err := CreateToken(input.Username)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = CreateAuth(input.Username,tokenInfo)
  	if err != nil {
     			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
     			return
  	}

  	session := sessions.Default(c)
	session.Set("username", input.Username)

	c.JSON(http.StatusOK, gin.H{"tokenInfo": tokenInfo})
}

func CreateToken(username string) (*models.TokenDetails, error) {
  tokenInfo := &models.TokenDetails{}
  tokenInfo.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
  

  tokenInfo.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
  

  var err error
  //Creating Access Token
  os.Setenv("ACCESS_SECRET", "chandra") //this should be in an env file
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true

  atClaims["username"] = username
  atClaims["exp"] = tokenInfo.AccessTokenExpires
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  tokenInfo.AccessToken, err = at.SignedString([]byte("ACCESS_SECRET"))
  if err != nil {
     return nil, err
  }
  //Creating Refresh Token
  os.Setenv("REFRESH_SECRET", "kamani") //this should be in an env file
  rtClaims := jwt.MapClaims{}
  rtClaims["username"] = username
  rtClaims["exp"] = tokenInfo.RefreshTokenExpires
  rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
  tokenInfo.RefreshToken, err = rt.SignedString([]byte("REFRESH_SECRET"))
  if err != nil {
     return nil, err
  }
  
  return tokenInfo, nil
}

func CreateAuth(username string, td *models.TokenDetails) error {
	at := time.Unix(td.AccessTokenExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RefreshTokenExpires, 0)
	now := time.Now()

	errAccess := client.Set(username, td.AccessToken, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	key := "username"+"_refreshtoken"
	errRefresh := client.Set(key, td.RefreshToken, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func SignOut(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("username")
	usernameStr := username.(string)
	client.Del(usernameStr).Result()

	c.JSON(http.StatusOK, gin.H{"logout": " successfully"})
}
