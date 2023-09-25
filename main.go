package main

import (
	"github.com/abhijithk1/vacation-planner/controllers"
	"github.com/abhijithk1/vacation-planner/middlewares"
	"github.com/abhijithk1/vacation-planner/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.MigrateDataBase()

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	user := router.Group("/api/user")
	user.Use(middlewares.JwtAuthMiddleware())
	user.POST("/add",controllers.Add)
	user.GET("/list/upcoming", controllers.ListUpcomingVacation)
	user.GET("/list/team/upcoming", controllers.ListTeamUpcomingVacation)

	admin := router.Group("/api/admin")
	admin.Use(middlewares.JwtAuthMiddleware())
	admin.GET("/users", controllers.ListUsers)
	// admin.GET("/user/:id",controllers.CurrentUser)

	router.Run(":8080")
}