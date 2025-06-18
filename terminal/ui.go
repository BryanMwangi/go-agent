package terminal

import (
	"fmt"
	"os"
	"time"

	"github.com/BryanMwangi/go-agent/auth"
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
	fmt.Println("Welcome to the go-agent terminal")
	session := auth.AuthenticateUser()

	// show loader
	ShowLoader("Verifying API Key...")
	err := auth.ValidateUser(session)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	StopLoader(1 * time.Second)
}
