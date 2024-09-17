package endpoints

import (
	m "app/pkg/model"
	"app/test/setup"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_Role_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error(setup.UnexpectedErrorMsg, err)
		return
	}

	t.Run("Add", func(t *testing.T) {
		// arrange
		value := `{"name":"default user"}`
		reader := strings.NewReader(value)
		expect := &m.Role{Id: 2, Name: "default user"}
		got := &m.Role{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/roles", reader)
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

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("AddAll", func(t *testing.T) {
		// arrange
		value := `
			[
				{"name":"name1"}, 
				{"name":"name2"}
			]
		`
		reader := strings.NewReader(value)
		expect := &[]m.Role{
			{Id: 3, Name: "name1"},
			{Id: 4, Name: "name2"},
		}
		got := &[]m.Role{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/roles/all", reader)
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

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("Remove", func(t *testing.T) {
		// arrange
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodDelete, "/roles/4", nil)
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

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 3
		got := &[]m.Role{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/roles", nil)
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

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if expect != len(*got) {
			t.Errorf("expected %v, but got %v", expect, len(*got))
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := &m.Role{Id: 2, Name: "default user"}
		got := &m.Role{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/roles/2", nil)
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

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("Edit", func(t *testing.T) {
		// arrange
		value := `{"name":"name3"}`
		reader := strings.NewReader(value)
		expect := &m.Role{Id: 2, Name: "name3"}
		got := &m.Role{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPut, "/roles/2", reader)
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

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})
}
