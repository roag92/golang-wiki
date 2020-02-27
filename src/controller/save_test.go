package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wiki/src/test"
)

func TestPostSave(t *testing.T) {
	state := test.NewTestState(t)

	save := SaveController{}

	t.Run("Return an internal server error", func(t *testing.T) {
		title := "MyTitleTest_1"

		req, err := http.NewRequest("POST", "/title/"+title, nil)

		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(save.PostSave)

		handler.ServeHTTP(res, req)

		if status := res.Code; status != http.StatusInternalServerError {
			t.Errorf(
				"handler returned wrong status code: got %v expected %v",
				status,
				http.StatusInternalServerError,
			)
		}

		body := res.Body.String()
		expected := "Title not found"

		if !strings.Contains(body, expected) {
			t.Errorf(
				"expected body `%s` is not equal to `%s`",
				expected,
				body,
			)
		}
	})

	state.ResetState()
}
