package main

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", ping)
	err := r.Run(":8084")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
