package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"
	"todo/initializers"
	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	bind := c.Bind(&body)

	if bind != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error getting the request body",
		})
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash the password",
		})
		return
	}

	//create the user

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User Creation failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "User created successfully",
	})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	bind := c.Bind(&body)

	if bind != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error reading the request body",
		})
		return
	}

	//fetch the user

	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		"Username": user.Email,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create token",
		})
		return
	}

	c.Copy().SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 360*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetAllUsers(db *gin.Context) {
	var users []models.User
	initializers.DB.Preload("Tasks").Find(&users)

	db.JSON(200, gin.H{
		"Message": "Users fetched successfully",
		"users":   users,
	})
}

func GetTokenFromRequest(c *gin.Context) { //for testing

	bearerToken, _ := c.Cookie("Authorization")
	splitToken := strings.Split(bearerToken, ".")

	if len(splitToken) > 1 {
		c.JSON(200, gin.H{
			"Message": "Splited",
			"Token":   splitToken[1],
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "one length",
		"Token":   bearerToken,
	})
}
