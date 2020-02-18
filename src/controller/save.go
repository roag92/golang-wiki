package controller

import (
	"net/http"
	"wiki/src/middleware"
	"wiki/src/model"
)

// SaveController structure to handle /save/ requests
type SaveController struct {
}

// PostSave handler to store a new post (e.g. POST /save/my-title -d 'body=my content')
func (save SaveController) PostSave(res http.ResponseWriter, req *http.Request) {
	title, ok := req.Context().Value(middleware.TitleKey).(string)

	if !ok {
		http.Error(res, "Title not found", http.StatusInternalServerError)
		return
	}

	body := req.FormValue("body")

	page := model.Page{Title: title, Body: []byte(body)}

	err := page.Save()

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(res, req, "/view/"+page.Title, http.StatusFound)
}
