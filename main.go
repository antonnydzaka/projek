package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	route := app
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"hello": "word",
		})
	})
	app.Run(":8000")
}
