package pgcode

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/derpl-del/gopro2/envcode/dtacode"
	"github.com/derpl-del/gopro2/envcode/imgcode"
	"github.com/derpl-del/gopro2/envcode/jscode"
	"github.com/derpl-del/gopro2/envcode/logcode"
	"github.com/derpl-del/gopro2/envcode/logincode"
	"github.com/derpl-del/gopro2/envcode/strcode"
)

var name string

//HomePage page
func HomePage(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		logcode.LogW("HomePage")
		data := dtacode.ReturnAllProduct()
		data.UsernameInfo.Username = user
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
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//ProductPage page
func ProductPage(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		logcode.LogW("ProductPage")
		data := strcode.AllProductData{}
		data.UsernameInfo.Username = user
		var filepath = path.Join("views", "product.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
			return
		}
		//fmt.Fprintf(w, "Hello World")
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//AddProductPage page
func AddProductPage(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	productid := currentTime.Format("200601021504-05")
	createdate := currentTime.Format("2006-01-02 15:04:05")
	var tittle1 = r.FormValue("tittle")
	var pname1 = r.FormValue("name")
	var pprice1 = r.FormValue("price")
	var pamount1 = r.FormValue("amount")
	var pquality1 = r.FormValue("quality")
	var pcategory1 = r.FormValue("category")
	price, _ := strconv.Atoi(pprice1)
	amount, _ := strconv.Atoi(pamount1)
	var Dproduct = strcode.ProductData{Pid: productid, Tittle: tittle1, Pname: pname1, Pprice: price, Pamount: amount, Pquality: pquality1, Pcategory: pcategory1, PCreateDate: createdate, PLastUpdate: createdate}
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
	imgcode.ResizeImg(productid)
	fmt.Println("File size: ", b)
	http.Redirect(w, r, "/", 302)
}

//BuyProduct page
func BuyProduct(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		logcode.LogW("BuyProduct")
		var pname1 = "product_" + r.FormValue("Pid") + ".json"
		data := jscode.GetProductData(pname1)
		data.UsernameInfo.Username = user
		var filepath = path.Join("views", "buy_product.html")
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
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//BuySomeProduct page
func BuySomeProduct(w http.ResponseWriter, r *http.Request) {
	logcode.LogW("BuySomeProduct")
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string(reqBody))
	buydata := jscode.BuyProductData(reqBody)
	jscode.UpdateAmount(buydata.InPid, buydata.InAmount)
}

//LoginHandler fo login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("password")
	redirectTarget := "/login"
	if name != "" && pass != "" {
		// .. check credentials ..
		logincode.SetSession(name, w, r)
		redirectTarget = "/"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

//LoginPage fo login
func LoginPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "login.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
		return
	}
	err = tmpl.Execute(w, "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logcode.LogE(err)
	}
}

//LogoutHandler logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logincode.ClearSession(w, r, "banana")
	http.Redirect(w, r, "/", 302)
}
