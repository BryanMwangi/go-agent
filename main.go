package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/BryanMwangi/go-agent/terminal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		cancel()
	}()

	terminal.StartUI(ctx)
}
