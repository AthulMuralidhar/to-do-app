// original sauce: https://dev.to/matijakrajnik/1-starting-gin-server-g5k
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "index page"})
		})
		api.POST("/todo", todoPost)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	// router at http://localhost:8000/api/
	router.Run(":8000")

}