package controller

import (
	"html/template"
	"net/http"
	"wiki/src/utils"
)

// IndexController structure to handle / requests
type IndexController struct {
	Template *template.Template
}

// GetIndex function to list all posts (e.g. GET /)
func (index IndexController) GetIndex(res http.ResponseWriter, req *http.Request) {
	posts, err := utils.Search(utils.GetTmpPath(), ".txt", true, false)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	err = index.Template.ExecuteTemplate(res, "index.html", posts)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
