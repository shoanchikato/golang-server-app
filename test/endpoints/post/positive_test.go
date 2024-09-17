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

func Test_Post_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error(setup.UnexpectedErrorMsg, err)
		return
	}

	t.Run("Add", func(t *testing.T) {
		// arrange
		value := `{"title":"title one", "body":"body one", "user_id":1}`
		reader := strings.NewReader(value)
		expect := &m.Post{Id: 1, Title: "title one", Body: "body one", UserId: 1}
		got := &m.Post{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/posts", reader)
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
				{"title":"title two", "body":"body two", "user_id":1},
				{"title":"title three", "body":"body three", "user_id":1}
			]
		`
		reader := strings.NewReader(value)
		expect := &[]m.Post{
			{Id: 2, Title: "title two", Body: "body two", UserId: 1},
			{Id: 3, Title: "title three", Body: "body three", UserId: 1},
		}
		got := &[]m.Post{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/posts/all", reader)
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
		req := httptest.NewRequest(http.MethodDelete, "/posts/51", nil)
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
		got := &[]m.Post{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/posts", nil)
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
		expect := &m.Post{Id: 2, Title: "title two", Body: "body two", UserId: 1}
		got := &m.Post{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/posts/2", nil)
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
		value := `{"title":"title new", "body":"body new", "user_id":1}`
		reader := strings.NewReader(value)
		expect := &m.Post{Id: 1, Title: "title new", Body: "body new", UserId: 1}
		got := &m.Post{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPut, "/posts/1", reader)
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
