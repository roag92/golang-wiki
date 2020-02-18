package controller

import (
	"html/template"
	"net/http"
	"wiki/src/middleware"
	"wiki/src/model"
)

// ViewController structure to handle /view/ requests
type ViewController struct {
	Template *template.Template
}

// GetView function to view a post by title (e.g. GET /view/my-title)
func (view ViewController) GetView(res http.ResponseWriter, req *http.Request) {
	title, _ := req.Context().Value(middleware.TitleKey).(string)

	page, err := model.GetPage(title)

	if err != nil {
		http.Redirect(res, req, "/edit/"+title, http.StatusFound)
		return
	}

	view.Template.ExecuteTemplate(res, "view.html", page)
}
