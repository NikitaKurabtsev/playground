package main

import (
	"errors"
)

type Config struct {
	Port int
}

type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		cfg.Port = defaultHttpPort()
	} else if *b.port < 0 {
		return Config{}, errors.New("port should be positive")
	} else if *b.port == 0 {
		cfg.Port = randomHttpPort()
	} else {
		cfg.Port = *b.port
	}

	return cfg, nil
}

func main() {
	builder := ConfigBuilder{}
	builder.Port(9090)

	cfg, err := builder.Build()
	if err != nil {
		err.Error()
	}

}
