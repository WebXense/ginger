package ratelimit

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Setup General Rate Limit
func GeneralRateLimit(engine *gin.Engine, period time.Duration, limit int64) {
	engine.Use(RateLimit(period, limit))
}
