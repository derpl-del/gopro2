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

	"github.com/derpl-del/gopro2/envcode/dbcode"
	"github.com/derpl-del/gopro2/envcode/dtacode"
	"github.com/derpl-del/gopro2/envcode/imgcode"
	"github.com/derpl-del/gopro2/envcode/jscode"
	"github.com/derpl-del/gopro2/envcode/logcode"
	"github.com/derpl-del/gopro2/envcode/logincode"
	"github.com/derpl-del/gopro2/envcode/strcode"
	"github.com/derpl-del/gopro2/envcode/txtcode"
	"github.com/gorilla/mux"
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
	_, user := logincode.GetUserName(r, name)
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
	var Dproduct = strcode.ProductData{Pid: productid, Powner: user, Tittle: tittle1, Pname: pname1, Pprice: price, Pamount: amount, Pquality: pquality1, Pcategory: pcategory1, PCreateDate: createdate, PLastUpdate: createdate}
	Dproduct.UsernameInfo.Username = user
	jscode.MakeProductData(Dproduct)
	dbcode.InsDataProduct(Dproduct.UsernameInfo.Username, Dproduct.Pid)
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
	_, user := logincode.GetUserName(r, name)
	logcode.LogW("BuySomeProduct")
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string(reqBody))
	buydata := jscode.BuyProductData(reqBody)
	jscode.UpdateAmount(buydata.InPid, buydata.InAmount)
	currentTime := time.Now()
	productid := currentTime.Format("2006010215-04_05")
	JSTittle := "hist/trx_" + productid + ".json"
	dbcode.InsTrxProduct(buydata, user, productid)
	jscode.MakeHistTrx(buydata, user, JSTittle)
}

//LoginHandler fo login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("LoginHandler"))
	logcode.LogW(string(reqBody))
	data := jscode.SignUpData(reqBody)
	name := data.Username
	pass := data.Password
	result := dbcode.ValidationUserData(name, pass)
	if result == true {
		// .. check credentials ..
		logincode.SetSession(name, w, r)
		dbcode.UpdateLoginData(data.Username)
	}
}

//SignUpHandler fo login
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string(reqBody))
	data := jscode.SignUpData(reqBody)
	dbcode.InsData(data.Username, data.Password)
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

//UserLoginVal func
func UserLoginVal(w http.ResponseWriter, r *http.Request) {
	var Articles strcode.UserLoginInfo
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("UserLoginVal"))
	logcode.LogW(string(reqBody))
	data := jscode.SignUpData(reqBody)
	result := dbcode.ValidationData("username", data.Username)
	if result == true {
		result = dbcode.ValidationUserData(data.Username, data.Password)
		if result == true {
			Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0000"}
		} else {
			Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0002"}
		}
	} else {
		Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0001"}
	}
	out, _ := json.Marshal(Articles)
	logcode.LogW(string(out))
	json.NewEncoder(w).Encode(Articles)
}

//GetProductByID func
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("GetProductByID"))
	logcode.LogW(string(reqBody))
	data := jscode.SignUpData(reqBody)
	result := dbcode.QueryProductByID(data.Username)
	ListProduct := []strcode.ProductData{}
	for _, dataDB := range result {
		tittle := "product_" + dataDB + ".json"
		article := jscode.GetProductData(tittle)
		ListProduct = append(ListProduct, article)
	}
	Product := strcode.AllProductData{ListProduct: ListProduct}
	Product.UsernameInfo.Username = data.Username
	json.NewEncoder(w).Encode(Product)
}

//ListProductPage fo login
func ListProductPage(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		result := dbcode.QueryProductByID(user)
		ListProduct := []strcode.ProductData{}
		for i, dataDB := range result {
			tittle := "product_" + dataDB + ".json"
			article := jscode.GetProductData(tittle)
			article.No = i + 1
			ListProduct = append(ListProduct, article)
		}
		Product := strcode.AllProductData{ListProduct: ListProduct}
		Product.UsernameInfo.Username = user
		var filepath = path.Join("views", "productlist.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
			return
		}
		err = tmpl.Execute(w, Product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//ListTrxPage fo login
func ListTrxPage(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		result := dbcode.QueryTrxByID(user)
		Product := strcode.HistPage{TrxProduct: result}
		Product.UsernameInfo.Username = user
		var filepath = path.Join("views", "trxlist.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
			return
		}
		err = tmpl.Execute(w, Product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//EditProduct page
func EditProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("EditProduct"))
	logcode.LogW(string(reqBody))
	data := jscode.EditProductData(reqBody)
	jscode.UpdateProduct(data)
	response := strcode.Response{ErrorCode: "0000"}
	json.NewEncoder(w).Encode(response)
}

//SignUpVal func
func SignUpVal(w http.ResponseWriter, r *http.Request) {
	var Articles strcode.UserLoginInfo
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("SignUpVal"))
	logcode.LogW(string(reqBody))
	data := jscode.SignUpData(reqBody)
	result := dbcode.ValidationData("username", data.Username)
	if result == false && len(data.Password) >= 6 && len(data.Username) >= 6 {
		Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0000"}
	} else if result == true {
		Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0001"}
	} else if len(data.Username) <= 6 {
		Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0002"}
	} else if len(data.Password) <= 6 {
		Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0003"}
	} else {
		Articles = strcode.UserLoginInfo{Username: data.Username, Password: data.Password, Message: "0004"}
	}
	out, _ := json.Marshal(Articles)
	logcode.LogW(string(out))
	json.NewEncoder(w).Encode(Articles)
}

//ChatPage fo login
func ChatPage(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		vars := mux.Vars(r)
		key := vars["id"]
		result := dbcode.QueryChatByChatID(key, user)
		ChatList := strcode.ChatIdle{ChatProduct: result}
		ChatList.UsernameInfo.Username = user
		var filepath = path.Join("views", "chat.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
			return
		}
		err = tmpl.Execute(w, ChatList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//ChatPageIdle fo login
func ChatPageIdle(w http.ResponseWriter, r *http.Request) {
	userName, user := logincode.GetUserName(r, name)
	if userName == true {
		result := dbcode.QueryChatByID(user)
		Product := strcode.ChatIdle{ChatProduct: result}
		Product.UsernameInfo.Username = user
		var filepath = path.Join("views", "chatidle.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
			return
		}
		err = tmpl.Execute(w, Product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logcode.LogE(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//CreateChat page
func CreateChat(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("CreateChat"))
	logcode.LogW(string(reqBody))
	data := jscode.ChatData(reqBody)
	currentTime := time.Now()
	productid := currentTime.Format("200601021504-05")
	user1 := data.User1
	user2 := data.User2
	chatid := user1[0:3] + "_" + user2[0:3] + "_" + productid
	dbcode.InsChatProduct(chatid, user1, user2, "")
	response := strcode.Content{Message: chatid}
	json.NewEncoder(w).Encode(response)
}

//GetChatByID func
func GetChatByID(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("GetChatByID"))
	logcode.LogW(string(reqBody))
	data := jscode.SignUpData(reqBody)
	result := dbcode.QueryChatByID(data.Username)
	ChatList := strcode.ChatIdle{ChatProduct: result}
	ChatList.UsernameInfo.Username = data.Username
	json.NewEncoder(w).Encode(ChatList)
}

//QueryChatByChatID fo login
func QueryChatByChatID(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("QueryChatByChatID"))
	logcode.LogW(string(reqBody))
	data := jscode.ChatData(reqBody)
	chatid := data.ChatID
	user := data.User1
	result := dbcode.QueryChatByChatID(chatid, user)
	ChatList := strcode.ChatIdle{ChatProduct: result}
	ChatList.UsernameInfo.Username = user
	json.NewEncoder(w).Encode(result)
}

//ChatVal fo login
func ChatVal(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("ChatVal"))
	logcode.LogW(string(reqBody))
	data := jscode.ChatData(reqBody)
	chatid := data.ChatID
	tittle := "txtfile/" + chatid + ".txt"
	fileexists := txtcode.FileExist(tittle)
	var msg string
	if fileexists == true {
		msg = "0000"
	} else {
		msg = "0001"
	}
	response := strcode.Response{ErrorCode: msg}
	json.NewEncoder(w).Encode(response)
}

//ChatGet fo login
func ChatGet(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("ChatVal"))
	logcode.LogW(string(reqBody))
	data := jscode.ChatData(reqBody)
	chatid := data.ChatID
	file := txtcode.ReadFile(chatid)
	response := strcode.Content{Message: file}
	json.NewEncoder(w).Encode(response)
}

//GetChatVal page
func GetChatVal(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	logcode.LogW(string("GetChatVal"))
	logcode.LogW(string(reqBody))
	data := jscode.ChatData(reqBody)
	user1 := data.User1
	user2 := data.User2
	responseDb := dbcode.QueryChatExst(user1, user2)
	response := strcode.Content{Message: responseDb}
	if responseDb == "" {
		response.Message = "0001"
	}
	json.NewEncoder(w).Encode(response)
}
