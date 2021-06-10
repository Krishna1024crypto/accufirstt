package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("ping", func(ctx *gin.context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.Run()
}
