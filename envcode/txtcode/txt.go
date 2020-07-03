package txtcode

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/derpl-del/gopro2/envcode/logcode"
)

//CreateChat for write
func CreateChat(chatid string, payload string) {
	logTittle := "txtfile/" + chatid + ".txt"
	logtext := fmt.Sprintf("%v\n", payload)
	mydata := []byte(logtext)
	if FileExist(logTittle) {
		f, err := os.OpenFile(logTittle, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(logtext); err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(logTittle, mydata, 0777)
		// handle this error
		if err != nil {
			// print it out
			fmt.Println(err)
		}
	}
}

//ReadFile func
func ReadFile(chatid string) string {
	logTittle := "txtfile/" + chatid + ".txt"
	file, err := os.Open(logTittle)
	if err != nil {
		fmt.Println(err)
		logcode.LogE(err)
	}
	byteValue, _ := ioutil.ReadAll(file)
	return string(byteValue)
}

//FileExist validation
func FileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
