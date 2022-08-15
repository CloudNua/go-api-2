package main

import (
	"fmt"
	"os"

	"github.com/CloudNua/go-api-2/pkg/comments"
	"github.com/CloudNua/go-api-2/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CloudNua: Comment Service
// @version         1.0
// @description     This is a sample CRUD HTTP service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1/comment

// @securityDefinitions.basic  BasicAuth

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

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// port := viper.Get("PORT").(string)
	// dbUrl := viper.Get("DB_URL").(string)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	r := gin.Default()
	h := db.InitDatabase(dsn)

	comments.RegisterRoutes(r, h)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "CloudNua Comment Service",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
