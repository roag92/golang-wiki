package wiki

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"wiki/src/controller"
	"wiki/src/middleware"
	"wiki/src/utils"
)

// Run HTTP server on 8080 port
func Run() {
	httpPort := 8080

	files, err := utils.Search("/src/views/", ".html", false, true)

	if err != nil {
		log.Fatal(err)
	}

	template := template.Must(template.ParseFiles(files...))

	index := controller.IndexController{Template: template}
	view := controller.ViewController{Template: template}
	edit := controller.EditController{Template: template}
	save := controller.SaveController{}

	http.HandleFunc("/view/", middleware.PageTitle(view.GetView))
	http.HandleFunc("/edit/", middleware.PageTitle(edit.GetEdit))
	http.HandleFunc("/save/", middleware.PageTitle(save.PostSave))
	http.HandleFunc("/", index.GetIndex)

	fmt.Printf("Listening on %d\n", httpPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), middleware.LogRequest(http.DefaultServeMux)))
}
