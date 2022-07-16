package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func f(c *gin.Context) {
	t := time.Now()

	// Set example variable
	c.Set("examplesssssssssssssss", "12333333333333345")

	// before request

	c.Next()

	// after request
	latency := time.Since(t)
	log.Print(latency)

	// access the status we are sending
	status := c.Writer.Status()
	log.Println(status)
}

func main() {
	r := gin.Default()
	r.GET("/", f)

	r.Run(":3000")
}
