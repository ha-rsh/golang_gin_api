package routes

import (
	"github.com/gin-gonic/gin"
	"gin-mongo-api/controllers"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.POST("/cvedata", controllers.GetCvesFrom())
}
