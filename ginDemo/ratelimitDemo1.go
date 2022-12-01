package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	engine := gin.Default()
	engine.Use(RateLimitMiddleware(10*time.Second, 1))
	engine.GET("version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "2000",
			"message": "rate limit test",
		})
	})
	if err := engine.Run("0.0.0.0:8000"); err != nil {
		fmt.Println(err)
		return
	}
}
