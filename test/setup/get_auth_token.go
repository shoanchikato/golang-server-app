package setup

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	m "app/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func GetAuthTokens(app *fiber.App) (*m.Tokens, error) {
	body := strings.NewReader(`{"username":"john_doe", "password":"password1"}`)
	req := httptest.NewRequest(http.MethodPost, "/login", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tokens := &m.Tokens{}
	err = json.Unmarshal(bytes, tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
