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

func Test_Permission_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App
	_ = di.Validators

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	t.Run("Add", func(t *testing.T) {
		// arrange
		value := `{"name":"name", "entity":"entity", "operation":"operation"}`
		reader := strings.NewReader(value)
		expect := &m.Permission{Id: 49, Name: "name", Entity: "entity", Operation: "operation"}
		got := &m.Permission{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/permissions", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
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
				{"name":"name1", "entity":"entity1", "operation":"operation1"}, 
				{"name":"name2", "entity":"entity2", "operation":"operation2"}
			]
		`
		reader := strings.NewReader(value)
		expect := &[]m.Permission{
			{Id: 50, Name: "name1", Entity: "entity1", Operation: "operation1"},
			{Id: 51, Name: "name2", Entity: "entity2", Operation: "operation2"},
		}
		got := &[]m.Permission{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/permissions/all", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
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
		req := httptest.NewRequest(http.MethodDelete, "/permissions/51", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 50
		got := &[]m.Permission{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/permissions", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if expect != len(*got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := &m.Permission{Id: 50, Name: "name1", Entity: "entity1", Operation: "operation1"}
		got := &m.Permission{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/permissions/50", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
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
		value := `{"name":"name3", "entity":"entity3", "operation":"operation3"}`
		reader := strings.NewReader(value)
		expect := &m.Permission{Id: 1, Name: "name3", Entity: "entity3", Operation: "operation3"}
		got := &m.Permission{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPut, "/permissions/1", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
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
