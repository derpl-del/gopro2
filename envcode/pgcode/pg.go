package pgcode

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/derpl-del/gopro2/envcode/jscode"
	"github.com/derpl-del/gopro2/envcode/strcode"
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

//ProductPage page
func ProductPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "product.html")
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

//AddProductPage page
func AddProductPage(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	productid := currentTime.Format("200601021504-05")
	var tittle1 = r.FormValue("tittle")
	var pname1 = r.FormValue("name")
	var pprice1 = r.FormValue("price")
	var pamount1 = r.FormValue("amount")
	var pquality1 = r.FormValue("quality")
	price, _ := strconv.Atoi(pprice1)
	amount, _ := strconv.Atoi(pamount1)
	var Dproduct = strcode.ProductData{Pid: productid, Tittle: tittle1, Pname: pname1, Pprice: price, Pamount: amount, Pquality: pquality1}
	jscode.MakeProductData(Dproduct)
	var filepath = path.Join("views", "add_product.html")
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
