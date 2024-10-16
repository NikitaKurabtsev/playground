package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func defaultHttpPort() int {
	return 8080
}

func randomHttpPort() int {
	return rand.Intn(8080)
}

type configSettings struct {
	port    *int
	timeout time.Duration
}

type ConfigOption func(settings *configSettings) error

func WithPort(port int) ConfigOption {
	return func(settings *configSettings) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		settings.port = &port
		return nil
	}
}

func WithTimeout(timeout time.Duration) ConfigOption {
	return func(settings *configSettings) error {
		if timeout < 0 {
			return errors.New("timeout should be positive")
		}
		settings.timeout = timeout
		return nil
	}
}

func NewServer(address string, opts ...ConfigOption) (*http.Server, error) {
	settings := configSettings{
		timeout: 30 * time.Second, // Значение по умолчанию
	}

	for _, opt := range opts {
		err := opt(&settings)
		if err != nil {
			return nil, err
		}
	}

	var port int

	if settings.port == nil {
		port = defaultHttpPort()
	} else {
		if *settings.port == 0 {
			port = randomHttpPort()
		} else {
			port = *settings.port
		}
	}

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", address, port),
		ReadTimeout:  settings.timeout,
		WriteTimeout: settings.timeout,
	}, nil
}

func main() {
	server, err := NewServer(
		"127.0.0.1",
		WithPort(8080),
		WithTimeout(10*time.Second),
	)
	if err != nil {
		return
	}

	fmt.Printf("start server on \n%s \nwith port: \n%d", server.Addr, 8080)

	if err := server.ListenAndServe(); err != nil {
		err.Error()
	}
}
