package pagehandler

import (
	"net/http"
	"github.com/darfk/page"
	"github.com/gorilla/mux"
)

type PageHandler struct {}

func (ph PageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()

	var exists bool

	templateFiles, exists := query["template"]

	var templateFile string = "main"
	if exists && len(templateFiles) > 0 {
		templateFile = templateFiles[0]
	}

	vars := mux.Vars(req)

	pageFile, exists := vars["page"]

	if ! exists {
		pageFile = "index"
	}

	if len(pageFile) == 0 {
		pageFile = "index"
	}

	p:= page.NewPage()

	var err error
	err = p.LoadBody(pageFile)

	if err != nil {
		http.NotFound(res, req)
		return
	}

	err = p.LoadTemplate(templateFile)

	if err != nil {
		http.Error(res, "template not found error", http.StatusInternalServerError)
		return
	}

	err = p.Execute(res)

	if err != nil {
		http.Error(res, "template execution error", http.StatusInternalServerError)
		return
	}

	return
	
}

