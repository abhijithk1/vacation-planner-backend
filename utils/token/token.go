package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/abhijithk1/vacation-planner/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.GetAppConfig().SecretString),nil
	})
	if err != nil {
		return err
	}

	return nil
}

func GenerateToken(user_id uint, a ...string) (string, error) {
	tokenLifespan := utils.GetAppConfig().TokenLifespan

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["expire"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()

	if(a != nil) {
		claims["permission"] = a[0]
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString([]byte(utils.GetAppConfig().SecretString))
}

func ExtractToken(c *gin.Context) string {

	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if (len(strings.Split(bearerToken, " ")) == 2) {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.GetAppConfig().SecretString),nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f",claims["user_id"]),10,32)

		if err != nil {
			return 0, err
		}

		return uint(uid),nil
	}

	return 0, nil
}

func ExtractAdminToken(c *gin.Context) error {

	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.GetAppConfig().SecretString),nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		if claims["permission"] != "admin" {
			return fmt.Errorf(" Invalid Access Token")
		}
	}
	
	return nil
}