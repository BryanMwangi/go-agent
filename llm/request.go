package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/BryanMwangi/go-agent/utils"
)

type message struct {
	Role    string `json:"role"`    // "user", "system", "assistant"
	Content string `json:"content"` // the user's prompt or request
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream,omitempty"`
}

const (
	formatQuery = "The user requests to format the code below and has included a message in their request: "
)

func SystemPrompt(content string) message {
	return message{Role: "system", Content: content}
}

func UserPrompt(content string) message {
	return message{Role: "user", Content: content}
}

// used to send the request to the LLM API
func (c *Client) send(messages []message) ([]byte, error) {
	reqBody := chatRequest{
		Model:    c.model,
		Messages: messages,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body) //ignore errors if body is nil or malformed
		apiErr := fmt.Errorf("failed: %s\n%s", resp.Status, string(body))
		return nil, handleErrorCodes(resp.StatusCode, apiErr, *c.Cfg)
	}

	return io.ReadAll(resp.Body)
}

// Called by auth.ValidateAPIKey to ensure that the API key is valid
func OnStartQuery(c *Client) ([]byte, error) {
	messages := []message{
		UserPrompt("This is a start query to verify the API key is valid. Simply respond with 'OK'."),
	}
	return c.send(messages)
}

func Query(c *Client, query string) ([]byte, error) {
	messages := []message{
		UserPrompt(query),
		SystemPrompt("You are a helpful assistant."),
	}

	return c.send(messages)
}

func Format(c *Client, query, code, lang string) ([]byte, error) {
	parsedCode := utils.ParseCode(code, lang)
	newContent := utils.ParseCodeRequest(formatQuery, query, parsedCode)

	messages := []message{
		UserPrompt(newContent),
		SystemPrompt("You are a code formatter. Output only the cleaned code."),
	}

	return c.send(messages)
}
