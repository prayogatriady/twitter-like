package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/entities"
)

func ValidateToken(signedtoken string) (*entities.SignedDetails, error) {
	token, err := jwt.ParseWithClaims(
		signedtoken,
		&entities.SignedDetails{},
		func(t *jwt.Token) (any, error) {
			return []byte("secret"), nil

		},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	getClaims, ok := token.Claims.(*entities.SignedDetails)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return getClaims, nil
}

func AuthMiddleware(c *gin.Context) {
	// get signed token from cookie
	signedToken, err := c.Cookie("token")
	if err != nil || signedToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "401 - Unauthorized",
			"msg":    "Unauthorized" + err.Error() + signedToken,
		})
		c.Abort()
		return
	}

	getClaims, err := ValidateToken(signedToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "401 - Unauthorized",
			"msg":    err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("username", getClaims.User.Username)
	c.Set("email", getClaims.User.Email)
	c.Set("password", getClaims.User.Password)

	c.Next()
}

func GenerateAllToken(user entities.User) (string, string, error) {
	claims := &entities.SignedDetails{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &entities.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return token, "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}
