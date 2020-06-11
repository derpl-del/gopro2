package pgcode

import (
	"html/template"
	"net/http"
	"path"
)

//HomePage page
func HomePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "main.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Fprintf(w, "Hello World")
	err = tmpl.Execute(w, "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
