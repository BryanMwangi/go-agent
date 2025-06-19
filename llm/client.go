package llm

import (
	"net/http"

	"github.com/BryanMwangi/go-agent/config"
)

// Client is a wrapper around the http.Client
// Has methods for setting the request URI, method, headers, and body
// needed to send requests to the LLM API
//
// You can either get a streamed response that can be written as a std output
// as a stream or a static response that can be used for code formatting or any
// other operation such as running bash scripts or scripts in your code
type Client struct {
	apiKey  string
	baseURL string
	model   string
	client  *http.Client
}

func NewClient(cfg config.Config) *Client {
	return &Client{
		apiKey:  cfg.APIKey,
		baseURL: cfg.Session.API.BaseURL + cfg.Session.API.CompletionPath,
		model:   cfg.Session.Model,
		client:  http.DefaultClient,
	}
}
