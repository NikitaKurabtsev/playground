package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// for i in {1..5}; do curl -i http://localhost:8080/ & done
// ab -n 10 -c 5 http://localhost:8080/

type RateLimiter interface {
	Allow() bool
	Limit() int
	Remaining() int
	Reset() time.Time
}

type TokenBucket struct {
	fillRate   int
	capacity   int
	tokens     int
	lastUpdate time.Time
	interval   time.Duration
	mu         sync.Mutex
}

func NewTokenBucket(rate int, per time.Duration) *TokenBucket {
	return &TokenBucket{
		fillRate:   rate,
		capacity:   rate,
		tokens:     rate,
		lastUpdate: time.Now(),
		interval:   per,
	}
}

func (rl *TokenBucket) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(rl.lastUpdate)                          // find duration between lastUpdate and now
	rl.lastUpdate = now                                            // update lastUpdate to now
	rl.tokens += int(elapsedTime.Seconds() * float64(rl.fillRate)) // update tokens to elapsedTime.Seconds + rate

	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}

	if rl.tokens < 1 {
		return false
	}

	rl.tokens--
	return true
}

func (rl *TokenBucket) Limit() int {
	return rl.capacity
}

func (rl *TokenBucket) Remaining() int {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.tokens
}

func (rl *TokenBucket) Reset() time.Time {
	return time.Now().Add(rl.interval)
}

func RateLimitMiddleware(rl RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rl.Allow() {
				w.Header().Set("X-RateLimit-Limit", strconv.Itoa(rl.Limit()))
				w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(rl.Remaining()))
				w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(rl.Reset().Unix(), 10))
				w.Header().Set("Retry-After", strconv.FormatInt(rl.Reset().Unix()-time.Now().Unix(), 10))
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			w.Header().Set("X-RateLimit-Limit", strconv.Itoa(rl.Limit()))
			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(rl.Remaining()))
			next.ServeHTTP(w, r)
		})
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World\n")
}

func main() {
	rateLimiter := NewTokenBucket(2, time.Second)

	http.Handle("/", RateLimitMiddleware(rateLimiter)(http.HandlerFunc(helloHandler)))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
