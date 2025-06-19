package llm

import (
	"fmt"
	"net/http"

	"github.com/BryanMwangi/go-agent/config"
)

func handleErrorCodes(code int, err error, cfg config.Config) error {
	switch code {
	case http.StatusBadRequest:
		return fmt.Errorf("bad request: %w", err)
	case http.StatusUnauthorized:
		return config.InvalidateSession(cfg)
	case http.StatusTooManyRequests:
		return fmt.Errorf("too many requests: %w", err)
	case http.StatusInternalServerError:
		return fmt.Errorf("internal server error: %w", err)
	case http.StatusServiceUnavailable:
		return fmt.Errorf("service unavailable: %w", err)
	default:
		return fmt.Errorf("unknown error: %w", err)
	}
}
