package controller

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"wiki/src/middleware"
	"wiki/src/test"
)

func TestGetView(t *testing.T) {
	state := test.NewTestState(t)

	fullPath, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	template := template.Must(template.ParseFiles(fullPath + "/templates/view.html"))

	view := ViewController{Template: template}

	t.Run("Redirect to edit page", func(t *testing.T) {
		title := "MyTitleTest_1"

		req, err := http.NewRequest("GET", "/view/"+title, nil)

		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), middleware.TitleKey, title)

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(view.GetView)

		handler.ServeHTTP(res, req.WithContext(ctx))

		if status := res.Code; status != http.StatusFound {
			t.Errorf(
				"handler returned wrong status code: got %v expected %v",
				status,
				http.StatusFound,
			)
		}
	})

	t.Run("Retrieve the view page", func(t *testing.T) {
		title := "MyTitleTest_1"
		content := "my-content-body"

		base := fmt.Sprintf("%s/tmp/_tests/", fullPath)

		os.MkdirAll(base, os.ModePerm)

		ioutil.WriteFile(
			fmt.Sprintf("%s%s.txt", base, title),
			[]byte(content),
			0600,
		)

		req, err := http.NewRequest("GET", "/view/"+title, nil)

		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), middleware.TitleKey, title)

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(view.GetView)

		handler.ServeHTTP(res, req.WithContext(ctx))

		if status := res.Code; status != http.StatusOK {
			t.Errorf(
				"handler returned wrong status code: got %v want %v",
				status,
				http.StatusOK,
			)
		}

		body := res.Body.String()

		if !strings.Contains(body, title) {
			t.Errorf(
				"not contains `%s` into the body `%s`",
				title,
				body,
			)
		}

		if !strings.Contains(body, content) {
			t.Errorf(
				"not contains `%s` into the body `%s`",
				content,
				body,
			)
		}

		os.RemoveAll(base)
	})

	state.ResetState()
}
