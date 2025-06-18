package terminal

import (
	"fmt"

	"github.com/BryanMwangi/go-agent/auth"
	"github.com/BryanMwangi/go-agent/config"
	"github.com/rivo/tview"
)

var App *tview.Application

func StartUI() {
	fmt.Println(`
 ██████╗  ██████╗        █████╗  ██████╗ ███████╗███╗   ██╗████████╗
██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝ ██╔════╝████╗  ██║╚══██╔══╝
██║  ███╗██║   ██║█████╗███████║██║  ███╗█████╗  ██╔██╗ ██║   ██║   
██║   ██║██║   ██║╚════╝██╔══██║██║   ██║██╔══╝  ██║╚██╗██║   ██║   
╚██████╔╝╚██████╔╝      ██║  ██║╚██████╔╝███████╗██║ ╚████║   ██║   
 ╚═════╝  ╚═════╝       ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═══╝   ╚═╝
 `)
	cfg := config.InitConfig()

	fmt.Println("Welcome to the go-agent terminal")

	if cfg.APIKey == "" {
		auth.AuthenticateUser(cfg)
	}

	// run the terminal tool
	Run(cfg)
}
