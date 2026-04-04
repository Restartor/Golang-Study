package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)
var (
	limiters = make(map[string]*rate.Limiter)
	mu sync.Mutex
)

func getLimiter(ip string) *rate.Limiter{
	mu.Lock()
	defer mu.Lock()

	if limiter, exist := limiters[ip]; exist{
		return limiter
	}

	limiter := rate.NewLimiter(5, 10)
	limiters[ip] = limiter
	return limiter
}

func RateLimiter() gin.HandlerFunc{
	return func(c *gin.Context) {
		limiter := getLimiter(c.ClientIP())
		if !limiter.Allow() {
			respondError(c, http.StatusTooManyRequests, "Too many request, please calm down!")
			c.Abort()
			return			
		}
		c.Next()
	}
}