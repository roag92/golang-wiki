package controller

import (
	"context"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"wiki/src/middleware"
	"wiki/src/test"
)

func TestGetEdit(t *testing.T) {
	state := test.NewTestState(t)

	fullPath, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	template := template.Must(template.ParseFiles(fullPath + "/templates/edit.html"))

	edit := EditController{Template: template}

	t.Run("Retrieve the edit page", func(t *testing.T) {
		title := "MyTitleTest"

		req, err := http.NewRequest("GET", "/edit/"+title, nil)

		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), middleware.TitleKey, title)

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(edit.GetEdit)

		handler.ServeHTTP(res, req.WithContext(ctx))

		if status := res.Code; status != http.StatusOK {
			t.Errorf(
				"handler returned wrong status code: got %v expected %v",
				status,
				http.StatusOK,
			)
		}

		expected := "/save/MyTitleTest"

		body := res.Body.String()

		if !strings.Contains(body, expected) {
			t.Errorf(
				"not contains `%s` into the body `%s`",
				expected,
				body,
			)
		}
	})

	state.ResetState()
}
