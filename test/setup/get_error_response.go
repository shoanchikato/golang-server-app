package setup

import (
	"encoding/json"
	"net/http"
	"testing"
)

func GetErrorResponse(t *testing.T, resp *http.Response) *map[string]any {
	value := &map[string]any{}

	err := json.NewDecoder(resp.Body).Decode(value)
	if err == nil {
		return value
	}

	stringValue := ""

	switch {
	case resp.StatusCode >= 400 && resp.StatusCode < 600:
		stringValue = http.StatusText(resp.StatusCode)
	default:
		t.Error("error decoding error message: ", err)
	}

	return &map[string]any{
		"error": map[string]string{
			"message": stringValue,
		},
	}
}
