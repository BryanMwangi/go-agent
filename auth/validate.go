package auth

import (
	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/utils"
)

func ValidateUser(session config.Session) error {
	utils.Welcome(session.Name)
	return nil
}
