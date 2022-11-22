package ginger

import (
	"os"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
	Engine = gin.Default()
}

// RUN starts the engine
// If GIN_MODE == release, the host will be 0.0.0.0
// Otherwise, the host will be 127.0.0.1
// The port is always set to 5000
func Run(engine *gin.Engine) {
	if os.Getenv("GIN_MODE") == "release" {
		engine.Run("0.0.0.0:5000")
	} else {
		engine.Run("127.0.0.1:5000")
	}
}
