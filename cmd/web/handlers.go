package web

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var files = []string{
	"./ui/html/home.page.html",
	"./ui/html/base.layout.html",
	"./ui/html/footer.partial.html",
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, "Hello")
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
	w.Write([]byte("hello!!!"))
}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Show snippet with ID %d...", id)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	w.Write([]byte("create snippet!!!"))
}
