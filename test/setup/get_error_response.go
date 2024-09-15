package setup

import (
	"encoding/json"
	"io"
	"testing"
)

func GetErrorResponse(t *testing.T, reader io.Reader) *map[string]any {
	value := &map[string]any{}

	err := json.NewDecoder(reader).Decode(value)
	if err != nil {
		t.Error("error decoding error message: ", err)
		return nil
	}

	return value
}
