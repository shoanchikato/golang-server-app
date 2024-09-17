package endpoints

import (
	m "app/pkg/model"
	"app/test/setup"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Role_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	t.Run("Login", func(t *testing.T) {
		// arrange
		value := `{"username":"john_doe", "password":"password1"}`
		reader := strings.NewReader(value)
		got := &m.Tokens{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodPost, "/login", reader)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}
	})

	t.Run("ResetPassword", func(t *testing.T) {
		// arrange
		value := `{"username":"john_doe", "password":"password3"}`
		reader := strings.NewReader(value)
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/reset-password", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})
}
