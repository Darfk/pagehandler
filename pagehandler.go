package pagehandler

import (
	"github.com/darfk/page"
	"net/http"
	"path"
)

type PageHandler struct{}

func (ph PageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()

	var exists bool

	templateFiles, exists := query["template"]

	var templateFile string = "main"
	if exists && len(templateFiles) > 0 {
		templateFile = templateFiles[0]
	}

	var pageFile string
	{
		var matched bool
		if matched, _ = path.Match("/", req.URL.Path); matched {
			pageFile = "index"
		} else if matched, _ = path.Match("/*", req.URL.Path); matched {
			pageFile = path.Base(req.URL.Path)
		} else {
			http.NotFound(res, req)
			return
		}
	}

	p := page.NewPage()

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
