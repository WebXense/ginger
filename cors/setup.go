package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config cors.Config

// Setup setup cors
func Setup(engine *gin.Engine, config Config) {
	engine.Use(cors.New(cors.Config(config)))
}

// SetupDefault setup cors using default config
func SetupDefault(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
}
