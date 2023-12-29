package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"todo/initializers"
	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// var ReqToken string
// func verifyToken(w http.ResponseWriter, r *http.Request) bool {
//     SecretKey := "SECRETKEY"
//     ReqToken = r.Header.Get("Authorization")
//     key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
//     if er != nil {
//         fmt.Println(er)

//         w.WriteHeader(http.StatusUnauthorized)
//         return false
//     }

//     token, err := jwt.Parse(ReqToken, func(token *jwt.Token) (interface{}, error) {

//         if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//             return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//         }
//         return key, nil
//     })

//     if err != nil {
//         fmt.Println(err)
//         w.WriteHeader(http.StatusUnauthorized)
//         return false
//     }
//     return true
//}
