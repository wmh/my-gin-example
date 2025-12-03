package services

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

var limiter *rateLimiter

func InitRateLimiter(limit int, window time.Duration) {
	limiter = &rateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	go limiter.cleanup()
}

func (rl *rateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for ip, times := range rl.requests {
			var validTimes []time.Time
			for _, t := range times {
				if now.Sub(t) < rl.window {
					validTimes = append(validTimes, t)
				}
			}
			if len(validTimes) == 0 {
				delete(rl.requests, ip)
			} else {
				rl.requests[ip] = validTimes
			}
		}
		rl.mu.Unlock()
	}
}

func RateLimitMiddleware() gin.HandlerFunc {
	if limiter == nil {
		InitRateLimiter(100, time.Minute)
	}

	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		now := time.Now()
		if times, exists := limiter.requests[ip]; exists {
			var validTimes []time.Time
			for _, t := range times {
				if now.Sub(t) < limiter.window {
					validTimes = append(validTimes, t)
				}
			}

			if len(validTimes) >= limiter.limit {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "Rate limit exceeded",
				})
				c.Abort()
				return
			}

			limiter.requests[ip] = append(validTimes, now)
		} else {
			limiter.requests[ip] = []time.Time{now}
		}

		c.Next()
	}
}
