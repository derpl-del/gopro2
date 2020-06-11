package jscode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/derpl-del/gopro2/envcode/strcode"
)

//MakeProductData JSON
func MakeProductData(input strcode.ProductData) {
	CacheTittle := "product_list/product_" + input.Pid + ".json"
	out, _ := json.Marshal(input)
	err := ioutil.WriteFile(CacheTittle, out, 0777)
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}
