package controller

import (
	"html/template"
	"net/http"
	"wiki/src/middleware"
	"wiki/src/model"
)

// EditController structure to handle /edit/ requests
type EditController struct {
	Template *template.Template
}

// GetEdit function to edit a post by title (e.g. GET /edit/my-title)
func (edit EditController) GetEdit(res http.ResponseWriter, req *http.Request) {
	title := req.Context().Value(middleware.TitleKey).(string)

	page, _ := model.GetPage(title)

	edit.Template.ExecuteTemplate(res, "edit.html", page)
}
