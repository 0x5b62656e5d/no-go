package util

import (
	"encoding/json"
	"main/util/types"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var clients = make(map[string]*types.Client)
var mu sync.Mutex

func GetClientLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if client, exists := clients[ip]; exists {
		return client.Limiter
	}

	limiter := rate.NewLimiter(5, 1)
	clients[ip] = &types.Client{Limiter: limiter}
	return limiter
}
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
			ip = forwarded
		}

		limiter := GetClientLimiter(ip)

		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(types.StandardResponse[any]{
				Data:    nil,
				Message: nil,
				Error:   StringPtr("Too Many Requests"),
				Success: false,
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}
