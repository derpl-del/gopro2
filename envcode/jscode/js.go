package jscode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/derpl-del/gopro2/envcode/logcode"
	"github.com/derpl-del/gopro2/envcode/strcode"
)

//MakeProductData JSON
func MakeProductData(input strcode.ProductData) {
	JSTittle := "product_list/product_" + input.Pid + ".json"
	out, _ := json.Marshal(input)
	logcode.LogW(string(out))
	err := ioutil.WriteFile(JSTittle, out, 0777)
	if err != nil {
		// print it out
		fmt.Println(err)
		logcode.LogE(err)
	}
}

//GetProductData JSON
func GetProductData(input string) strcode.ProductData {
	JSTittle := "product_list/" + input
	//logcode.LogW(fmt.Sprintf("input read : %v", string(byteValue)))
	jsonFile, err := os.Open(JSTittle)
	if err != nil {
		fmt.Println(err)
		logcode.LogE(err)
	}
	//fmt.Println("Successfully Opened users.json")
	now := time.Now()
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	logcode.LogW(fmt.Sprintf("result read : %v", string(byteValue)))
	var struct1 strcode.ProductData
	json.Unmarshal(byteValue, &struct1)
	date := struct1.PCreateDate + "+0700"
	t, _ := time.Parse("2006-01-02 15:04:05-0700", date)
	timeis := now.Sub(t).Minutes()
	ptime := fmt.Sprintf("%v minute", math.Round(timeis))
	if timeis > 60 && timeis <= 1440 {
		timeis = math.Round(timeis / 60)
		ptime = fmt.Sprintf("%v hour", timeis)
	} else if timeis > 1440 {
		timeis = math.Round(timeis / 1440)
		ptime = fmt.Sprintf("%v days", timeis)
	}
	struct1.Ptime = ptime
	return struct1
}

//ReadFileJS JSON
func ReadFileJS(input string) strcode.ProductData {
	JSTittle := "product_list/product_" + input + ".json"
	jsonFile, err := os.Open(JSTittle)
	if err != nil {
		fmt.Println(err)
		logcode.LogE(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var struct1 strcode.ProductData
	json.Unmarshal(byteValue, &struct1)
	return struct1
}

//BuyProductData JSON
func BuyProductData(input []byte) strcode.BuyProduct {
	var struct1 strcode.BuyProduct
	json.Unmarshal(input, &struct1)
	return struct1
}

//UpdateAmount JSON
func UpdateAmount(pid string, amount string) {
	currentTime := time.Now()
	lastupdate := currentTime.Format("2006-01-02 15:04:05")
	FileData := ReadFileJS(pid)
	inamount, _ := strconv.Atoi(amount)
	diff := FileData.Pamount - inamount
	FileData.Pamount = diff
	FileData.PLastUpdate = lastupdate
	MakeProductData(FileData)
	if diff == 0 {
		DeleteProductData(pid)
	}
}

//DeleteProductData JSON
func DeleteProductData(pid string) {
	oldJSTittle := "product_list/product_" + pid + ".json"
	newJSTittle := "tmp_file/data_product/product_" + pid + ".json"
	err := os.Rename(oldJSTittle, newJSTittle)
	if err != nil {
		logcode.LogE(err)
	}
	oldfiletittle := "data_img/upload_" + pid + "_1.png"
	newfiletittle := "tmp_file/data_img/upload_" + pid + "_1.png"
	err2 := os.Rename(oldfiletittle, newfiletittle)
	if err != nil {
		logcode.LogE(err2)
	}

}
