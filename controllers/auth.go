package controllers

import (
	"fmt"
	"net/http"

	"github.com/abhijithk1/vacation-planner/models"
	"github.com/abhijithk1/vacation-planner/utils/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var DB *gorm.DB

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "request failed",
			"error":   err.Error(),
		})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	fmt.Println("username: ", u.Username, "    Password : ", u.Password)

	_,err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not add the account",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account Added"})
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "request failed",
			"error":   err.Error(),
		})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": u.Username,
		"token":token,
	})
}

// func CurrentUser(c *gin.Context) {

// 	user_id, err := token.ExtractTokenID(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	u, err := models.GetUserID(user_id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
// }

func ListUsers(c *gin.Context) {

	err := token.ExtractAdminToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})

}