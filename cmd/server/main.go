package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Run - is going to be responsible for
// the instantiation and startup of the
// go application
func Run() error {
	fmt.Println("starting up api service")

	return nil
}

func main() {
	fmt.Println("Go-Gin REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
