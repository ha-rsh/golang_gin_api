package main

import (
	"gin-mongo-api/configs"
	"github.com/gin-gonic/gin"
	"gin-mongo-api/routes"
)

func main()  {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)
	

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"data" : "hello from gin and mongodb",
	// 	})
	// })

    router.Run("localhost:8080")
}