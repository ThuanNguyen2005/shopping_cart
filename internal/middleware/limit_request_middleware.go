package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	clients = map[string]*Client{}
	mu      sync.Mutex
)

func getIPUsers(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func getLimit(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exist := clients[ip]
	if !exist {
		newClient := &Client{
			limiter:  rate.NewLimiter(5, 15), //5 request/s, 10 brust
			lastSeen: time.Now(),
		}
		clients[ip] = newClient
		return clients[ip].limiter
	}
	client.lastSeen = time.Now()
	return client.limiter
}

// hey -n 20 -c 1 -H "X-API-Key: 789fd5b7-f072-4c18-acac-db3c0e78d462" http://localhost:8080/api/v1/user/
func CleanUpClient() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func LimitRequestMiddleware(ratelimitLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getIPUsers(ctx)
		limiter := getLimit(ip)
		if !limiter.Allow() {
			if shouldRatelimit(ip) {
				ratelimitLogger.Error().
					Str("path", ctx.Request.URL.Path).
					Str("method", ctx.Request.Method).
					Str("client_IP", ctx.ClientIP()).
					Msg("panic occurred.")
			}

			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many request",
			})
			return
		}
		ctx.Next()
	}
}

var maplogRatelimit sync.Map
var timelogRatelimit = 10 * time.Minute

func shouldRatelimit(ip string) bool {
	now := time.Now()
	if val, ok := maplogRatelimit.Load(ip); ok {
		if time := val.(time.Time); now.Sub(time) < timelogRatelimit {
			return false
		}
	}
	maplogRatelimit.Store(ip, now)
	return true
}
