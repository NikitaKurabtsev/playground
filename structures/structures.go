package main

import (
	"io"
	"os"
)

type Logger struct {
	io.WriteCloser
}

func main() {
	logger := &Logger{WriteCloser: os.Stdout}

	_, err := logger.Write([]byte("Hello World"))
	if err != nil {
		err.Error()
	}

	err = logger.Close()
	if err != nil {
		err.Error()
	}
}
