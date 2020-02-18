package middleware

import (
	"context"
	"net/http"
	"regexp"
)

type key string

// TitleKey key used for context
const TitleKey key = "title"

// PageTitle handler function to validate if the edit, view or save path has a valid title
func PageTitle(next func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

		m := validPath.FindStringSubmatch(req.URL.Path)

		if m == nil {
			http.NotFound(res, req)
			return
		}

		ctx := context.WithValue(req.Context(), TitleKey, m[2])

		next(res, req.WithContext(ctx))
	}
}
