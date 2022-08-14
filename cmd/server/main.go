package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Run - is going to be responsible for
// the instantiation and startup of the
// go application
func Run() error {
	fmt.Println("starting up api service")

	return nil
}

func setupRouter() *gin.Engine {

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

	log, _ := zap.NewDevelopment()
	defer log.Sync()

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
