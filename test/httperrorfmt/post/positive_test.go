package post

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	v "app/pkg/validation"
	"app/test/setup"
	"reflect"
	"testing"
)

func addPost(t *testing.T, validation v.Validators) {
	post := m.NewPost("title one", "body one", 1)

	err := validation.Post.Add(post)
	if err != nil {
		t.Error("error adding post")
		return
	}
}

func TestPositivePost(t *testing.T) {
	app := setup.Run()
	addPost(t, app.Validators)
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 1

		// act
		posts, err := app.HttpErrorFmts.Post.GetAll(1, 0, 50)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}
		got := len(*posts)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := m.Post{Id: 1, Title: "title one", Body: "body one", UserId: 1}

		// act
		got, err := app.HttpErrorFmts.Post.GetOne(1, 1)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("Add", func(t *testing.T) {
		// arrange
		got := m.NewPost("title two", "body two", 1)
		expect := m.Post{Id: 2, Title: "title two", Body: "body two", UserId: 1}

		// act
		err := app.HttpErrorFmts.Post.Add(1, got)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("Edit", func(t *testing.T) {
		// arrange
		got := m.NewPost("title 2", "body 2", 1)
		expect := m.Post{Id: 2, Title: "title 2", Body: "body 2", UserId: 1}

		// act
		err := app.HttpErrorFmts.Post.Edit(1, 2, got)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("Remove", func(t *testing.T) {
		// arrange
		expect := e.NewErrRepoNotFound("post id", "2")

		// act
		err := app.HttpErrorFmts.Post.Remove(1, 2)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}
		_, got := app.HttpErrorFmts.Post.GetOne(1, 2)
		if got == nil {
			t.Error("expected an error, but none was returned")
			return
		}

		// assert
		if !(got.Error() == expect.Error()) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("AddAll", func(t *testing.T) {
		// arrange
		got := &[]*m.Post{
			m.NewPost("title two", "body two", 1),
			m.NewPost("title three", "body three", 1),
		}
		expect := &[]*m.Post{
			{Id: 2, Title: "title two", Body: "body two", UserId: 1},
			{Id: 3, Title: "title three", Body: "body three", UserId: 1},
		}

		// act
		err := app.HttpErrorFmts.Post.AddAll(1, got)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})
}
