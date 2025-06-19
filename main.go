package main

import (
	"context"

	"github.com/BryanMwangi/go-agent/terminal"
)

func main() {
	bgCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	terminal.StartUI(bgCtx)
}
