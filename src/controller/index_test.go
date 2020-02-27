package controller

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"wiki/src/test"
)

func TestGetIndex(t *testing.T) {
	state := test.NewTestState(t)

	fullPath, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	template := template.Must(template.ParseFiles(fullPath + "/templates/index.html"))

	index := IndexController{Template: template}

	t.Run("Retrieve 0 pages to list", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(index.GetIndex)

		handler.ServeHTTP(res, req)

		assertIndexView(t, res.Code, res.Body.String(), 0)
	})

	t.Run("Retrieve with 2 pages to list", func(t *testing.T) {
		expectedPages := 2

		base := fmt.Sprintf("%s/tmp/_tests/", fullPath)

		os.MkdirAll(base, os.ModePerm)

		for i := 0; i < expectedPages; i++ {
			ioutil.WriteFile(
				fmt.Sprintf("%s%s_%d.txt", base, "MyTest", i),
				[]byte("foo-fontent"),
				0600,
			)
		}

		req, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(index.GetIndex)

		handler.ServeHTTP(res, req)

		assertIndexView(t, res.Code, res.Body.String(), expectedPages)

		os.RemoveAll(base)
	})

	state.ResetState()
}

func assertIndexView(t *testing.T, code int, body string, totalPages int) {
	if status := code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusOK,
		)
	}

	expected := "Welcome to GoLang - Wiki"

	if !strings.Contains(body, expected) {
		t.Errorf(
			"not contains `%s` into the body `%s`",
			expected,
			body,
		)
	}

	actualPages := strings.Count(body, "<li>")

	if actualPages != totalPages {
		t.Error(actualPages)

		t.Error(body)

		t.Errorf(
			"got `%d` pages, expected `%d` pages",
			actualPages,
			totalPages,
		)
	}
}
