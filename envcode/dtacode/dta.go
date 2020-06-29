package dtacode

import (
	"fmt"
	"io/ioutil"

	"github.com/derpl-del/gopro2/envcode/jscode"
	"github.com/derpl-del/gopro2/envcode/logcode"
	"github.com/derpl-del/gopro2/envcode/strcode"
)

//ListProduct array
var ListProduct []strcode.ProductData

//ReturnAllProduct for homepage
func ReturnAllProduct() strcode.AllProductData {
	ListProduct = []strcode.ProductData{}
	fileInfo, err := ioutil.ReadDir("product_list/")
	if err != nil {
		fmt.Println(err)
		logcode.LogE(err)
	}
	for i, info := range fileInfo {
		var article = jscode.GetProductData(info.Name())
		article.No = i + 1
		ListProduct = append(ListProduct, article)
	}
	Articles := strcode.AllProductData{ListProduct: ListProduct}
	return Articles
}
