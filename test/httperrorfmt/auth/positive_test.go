package auth

import (
	m "app/pkg/model"
	"app/test/setup"
	"testing"
)

func TestPositiveAuth(t *testing.T) {
	app := setup.Run()
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("Login", func(t *testing.T) {
		// arrange
		credentials := m.NewCredentials("john_doe", "password1")

		// act
		token, err := app.HttpErrorFmts.Auth.Login(credentials)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if token == nil {
			t.Errorf("login failed, no tokens where provided")
			return
		}
	})

	t.Run("ResetPassword", func(t *testing.T) {
		// arrange
		credentials := m.NewCredentials("john_doe", "password2")

		// act
		err := app.HttpErrorFmts.Auth.ResetPassword(credentials.Username, credentials.Password)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		token, err := app.HttpErrorFmts.Auth.Login(credentials)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if token == nil {
			t.Errorf("login failed, no tokens where provided")
			return
		}
	})
}
