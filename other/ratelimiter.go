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

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	timeElapsed := now.Sub(tb.lastUpdate)                          // find duration between lastUpdate and now
	tb.lastUpdate = now                                            // update lastUpdate to now
	tb.tokens += int(timeElapsed.Seconds() * float64(tb.fillRate)) // update tokens to timeElapsed.Seconds + rate

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	if tb.tokens < 1 {
		return false
	}

	tb.tokens--
	return true
}

func (tb *TokenBucket) Limit() int {
	return tb.capacity
}

func (tb *TokenBucket) Remaining() int {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	return tb.tokens
}

func (tb *TokenBucket) Reset() time.Time {
	return time.Now().Add(tb.interval)
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
