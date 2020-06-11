package ctrcode

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/gopro2/envcode/pgcode"
	"github.com/gorilla/mux"
)

//Funchandler Controller
func Funchandler() {
	fmt.Println("morning")
	r := mux.NewRouter()
	r.HandleFunc("/", pgcode.HomePage)
	r.HandleFunc("/product", pgcode.ProductPage)
	r.HandleFunc("/add_product", pgcode.AddProductPage)
	r.PathPrefix("/envstyle/").Handler(http.StripPrefix("/envstyle/", http.FileServer(http.Dir("envstyle"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
