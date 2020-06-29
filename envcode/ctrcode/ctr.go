package ctrcode

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/gopro2/envcode/logincode"
	"github.com/derpl-del/gopro2/envcode/pgcode"
	"github.com/gorilla/mux"
)

//Funchandler Controller
func Funchandler() {
	logincode.CreateStore()
	fmt.Println("morning")
	r := mux.NewRouter()
	r.HandleFunc("/", pgcode.HomePage)
	r.HandleFunc("/login", pgcode.LoginPage)
	r.HandleFunc("/login_page", pgcode.LoginHandler).Methods("POST")
	r.HandleFunc("/signup_page", pgcode.SignUpHandler).Methods("POST")
	r.HandleFunc("/logout", pgcode.LogoutHandler).Methods("POST")
	r.HandleFunc("/UserLoginVal", pgcode.UserLoginVal).Methods("POST")
	r.HandleFunc("/SignLoginVal", pgcode.SignUpVal).Methods("POST")
	r.HandleFunc("/QueryProductByID", pgcode.GetProductByID).Methods("POST")
	r.HandleFunc("/product", pgcode.ProductPage)
	r.HandleFunc("/productlist", pgcode.ListProductPage)
	r.HandleFunc("/add_product", pgcode.AddProductPage)
	r.HandleFunc("/buy_product", pgcode.BuyProduct)
	r.HandleFunc("/EditHandle", pgcode.EditProduct)
	r.HandleFunc("/buy_someproduct", pgcode.BuySomeProduct)
	r.PathPrefix("/envstyle/").Handler(http.StripPrefix("/envstyle/", http.FileServer(http.Dir("envstyle"))))
	r.PathPrefix("/dataimg/").Handler(http.StripPrefix("/dataimg/", http.FileServer(http.Dir("data_img"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
