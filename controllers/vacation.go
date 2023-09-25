package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abhijithk1/vacation-planner/models"
	"github.com/abhijithk1/vacation-planner/utils/token"
	"github.com/gin-gonic/gin"
)

type AddVacationInput struct {
	UserId uint
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	FromDate time.Time `json:"fromDate" binding:"required"`
	EndDate time.Time `json:"endDate" binding:"required"`
}

func Add(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User",
			"error": err.Error(),
		})
		return
	}

	var input AddVacationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "request failed",
			"error":   err.Error(),
		})
		return
	}

	v := models.Vacation{}

	v.UserId = user_id
	v.Title = input.Title
	v.Description = input.Description
	v.FromDate = input.FromDate
	v.EndDate = input.EndDate

	fmt.Println("vacation : ", v)

	_,err = v.AddVacation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not add the vacation",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vacation Added"})
}

func ListUpcomingVacation(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User",
			"error": err.Error(),
		})
		return
	}

	fmt.Println("user id : ", user_id)

	v, err := models.ListUpcomingVacationsID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("v : ", v)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": v})
}

func ListTeamUpcomingVacation(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User",
			"error": err.Error(),
		})
		return
	}

	fmt.Println("user id : ", user_id)

	v, err := models.ListTeamUpcomingVacations(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("v : ", v)

	data, err := models.ListTeamMember(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("data : ", data)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}