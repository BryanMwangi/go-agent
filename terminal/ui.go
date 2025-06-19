package terminal

import (
	"context"
	"fmt"

	"github.com/BryanMwangi/go-agent/auth"
	"github.com/BryanMwangi/go-agent/config"
	"github.com/rivo/tview"
)

var App *tview.Application

func StartUI(ctx context.Context) {
	fmt.Println(`
 ██████╗  ██████╗        █████╗  ██████╗ ███████╗███╗   ██╗████████╗
██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝ ██╔════╝████╗  ██║╚══██╔══╝
██║  ███╗██║   ██║█████╗███████║██║  ███╗█████╗  ██╔██╗ ██║   ██║   
██║   ██║██║   ██║╚════╝██╔══██║██║   ██║██╔══╝  ██║╚██╗██║   ██║   
╚██████╔╝╚██████╔╝      ██║  ██║╚██████╔╝███████╗██║ ╚████║   ██║   
 ╚═════╝  ╚═════╝       ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═══╝   ╚═╝
 `)
	cfg := config.InitConfig()

	if cfg.APIKey == "" {
		auth.AuthenticateUser(cfg)
	}
	fmt.Printf("Welcome %s to the go-agent terminal", cfg.Name)

	// run the terminal tool
	Run(cfg, ctx)
}
