package auth

import (
	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/utils"
)

func validateUser(cfg config.Config) error {
	utils.Welcome(cfg.Name)
	return nil
}
