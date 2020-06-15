package pgcode

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/derpl-del/gopro2/envcode/dtacode"
	"github.com/derpl-del/gopro2/envcode/jscode"
	"github.com/derpl-del/gopro2/envcode/logcode"
	"github.com/derpl-del/gopro2/envcode/strcode"
)

//HomePage page
func HomePage(w http.ResponseWriter, r *http.Request) {
	logcode.LogW("HomePage")
	data := dtacode.ReturnAllProduct()
	out, _ := json.Marshal(data)
	logcode.LogW(string(out))
	var filepath = path.Join("views", "main.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
	}
}

//ProductPage page
func ProductPage(w http.ResponseWriter, r *http.Request) {
	logcode.LogW("ProductPage")
	var filepath = path.Join("views", "product.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
		return
	}
	//fmt.Fprintf(w, "Hello World")
	err = tmpl.Execute(w, "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
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
	var pcategory1 = r.FormValue("category")
	price, _ := strconv.Atoi(pprice1)
	amount, _ := strconv.Atoi(pamount1)
	var Dproduct = strcode.ProductData{Pid: productid, Tittle: tittle1, Pname: pname1, Pprice: price, Pamount: amount, Pquality: pquality1, Pcategory: pcategory1}
	jscode.MakeProductData(Dproduct)
	r.ParseMultipartForm(10 << 20)
	file, handler, err1 := r.FormFile("myFile")
	if err1 != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err1)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	filetittle := "data_img/upload_" + productid + "_1.png"
	img, _ := os.Create(filetittle)
	defer img.Close()
	b, _ := io.Copy(img, file)
	fmt.Println("File size: ", b)
	var filepath = path.Join("views", "add_product.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
		return
	}
	//fmt.Fprintf(w, "Hello World")
	err = tmpl.Execute(w, "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
	}
}
