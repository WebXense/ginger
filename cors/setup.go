package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Setup setup cors
func Setup(engine *gin.Engine, config cors.Config) {
	engine.Use(cors.New(config))
}

// SetupDefault setup cors using default config
func SetupDefault(engine *gin.Engine) {
	engine.Use(cors.New(cors.DefaultConfig()))
}
