package controllers

import (
  "fmt"
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "eltropy-assignment/models"
  "golang.org/x/crypto/bcrypt"
  "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")


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
	fmt.Println("inside SignUp")
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("failed to bind SignUp")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("input", input)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	
	
	user := models.User{Username: input.Username, Password: string(hashedPassword), EmployeeId: 1}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func SignIn(c *gin.Context) {
	

	fmt.Println("inside signin")
	var input SigninCredentials
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("failed to bind SignUp")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("inside signin", input.Username)

	var user models.User


	x :=models.DB.Where("username = ?", input.Username).First(&user)

	fmt.Println("x",x)


	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    	return
  	}

  
	fmt.Println("inside signin password is ", input.Password)

	fmt.Println("inside signin password is ", user.Password)

	fmt.Println("inside signin password is ", input.Password)
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

		expirationTime := time.Now().Add(5 * time.Minute)

		credsWithToken := &CredsWithToken{
			Username: input.Username,
			StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credsWithToken)


	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	http.SetCookie(c.Writer, &http.Cookie{
        Name:     "token",
        Value:    tokenString,
        MaxAge:   300,
        HttpOnly: false,
    })
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	coo,err :=c.Cookie("token")
	fmt.Println("err", err)
	fmt.Println("coo", coo)

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}

