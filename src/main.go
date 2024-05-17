package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hawa130/computility-cloud/config"
)

func main() {
	startServer()

	config.OnConfigChange(func() {
		restartServer()
	})

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // waiting for interrupt or terminate signal

	stopServer()
}
