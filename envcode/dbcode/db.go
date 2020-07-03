package dbcode

import (
	"database/sql"
	"fmt"
	"log"

	//for framework
	"github.com/derpl-del/gopro2/envcode/strcode"
	"github.com/derpl-del/gopro2/envcode/utilitycode"

	//for framework
	_ "github.com/godror/godror"
)

//OpenCon DB
func OpenCon() *sql.DB {
	db, err := sql.Open("godror", "BANANA/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return db
	}
	return db
}

//InsData insert data
func InsData(input1 string, input2 string) {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO MYCLIENT VALUES ('%s', '%s')", input1, input2)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
	statementSQL = fmt.Sprintf("INSERT INTO CLIENT_INFO VALUES ('%s', SYSDATE, SYSDATE)", input1)
	rows, err = db.Query(statementSQL)
	defer rows.Close()
}

//InsDataProduct insert data
func InsDataProduct(input1 string, input2 string) {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO CLIENT_PRODUCT VALUES ('%s', '%s')", input1, input2)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
}

//InsTrxProduct insert data
func InsTrxProduct(data strcode.BuyProduct, buyer string, pid string) {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO TRX_INFO VALUES ('%s', '%s', '%s', '%s', '%s' , SYSDATE)", data.InOwner, buyer, data.InName, data.InAmount, data.InTotalPay)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
}

//UpdateLoginData insert data
func UpdateLoginData(input1 string) {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("UPDATE CLIENT_INFO SET LAST_LOGIN = SYSDATE WHERE USERNAME =  '%s'", input1)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
}

//ValidationUserData insert data
func ValidationUserData(input1 string, input2 string) bool {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("Select Username, Password from MYCLIENT where username = '%s' and password = '%s' and rownum = '1'", input1, input2)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	var Rs1 string
	var Rs2 string
	for rows.Next() {
		rows.Scan(&Rs1, &Rs2)
	}
	if Rs1 != "" {
		return true
	}
	return false
}

//ValidationData insert data
func ValidationData(input1 string, input2 string) bool {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("Select %s from MYCLIENT where %s = '%s'", input1, input1, input2)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	var Rs1 string
	for rows.Next() {
		rows.Scan(&Rs1)
	}
	if Rs1 != "" {
		return true
	}
	return false
}

//QueryProductByID func
func QueryProductByID(input string) []string {
	names := make([]string, 0)
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("Select PID from CLIENT_PRODUCT where username = '%s'", input)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return names
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	return names
}

//QueryTrxByID func
func QueryTrxByID(input string) []strcode.TrxProduct {
	names := make([]strcode.TrxProduct, 0)
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("Select Owner,Buyer,Product,amount,pay,create_date from TRX_INFO where owner = '%s' order by create_date desc", input)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return names
	}
	defer rows.Close()
	num := 0
	for rows.Next() {
		var owner string
		var buyer string
		var product string
		var amount string
		var pay string
		var createdate string
		if err := rows.Scan(&owner, &buyer, &product, &amount, &pay, &createdate); err != nil {
			log.Fatal(err)
		}
		date1 := utilitycode.SubsBefore(createdate, "T")
		date2 := utilitycode.SubsAfer(utilitycode.SubsBefore(createdate, "+"), "T")
		datefix := date1 + " " + date2
		num = num + 1
		DataProduct := strcode.TrxProduct{No: num, Owner: owner, Amount: amount, Buyer: buyer, CreateDate: datefix, Pay: pay, Product: product}
		names = append(names, DataProduct)
	}
	return names
}

//InsChatProduct insert data
func InsChatProduct(chatid string, user1 string, user2 string, payload string) {
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO CHAT_DATA VALUES ('%s', '%s', '%s','%s', SYSDATE,SYSDATE)", chatid, user1, user2, payload)
	fmt.Println("Query :" + statementSQL)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
}

//QueryChatByID func
func QueryChatByID(user string) []strcode.ChatProduct {
	names := make([]strcode.ChatProduct, 0)
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("Select chat_id,user1,user2,payload,last_update from CHAT_DATA where (user1 = '%s' or user2 = '%s') order by last_update desc", user, user)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return names
	}
	defer rows.Close()
	num := 0
	for rows.Next() {
		var chatid string
		var user1 string
		var user2 string
		var payload string
		var lastupdate string
		var puser string
		var suser string
		if err := rows.Scan(&chatid, &user1, &user2, &payload, &lastupdate); err != nil {
			log.Fatal(err)
		}
		if user1 == user {
			puser = user1
			suser = user2
		} else {
			puser = user2
			suser = user1
		}
		date1 := utilitycode.SubsBefore(lastupdate, "T")
		date2 := utilitycode.SubsAfer(utilitycode.SubsBefore(lastupdate, "+"), "T")
		datefix := date1 + " " + date2
		num = num + 1
		DataProduct := strcode.ChatProduct{No: num, ChatID: chatid, User1: puser, User2: suser, LastUpdate: datefix, Payload: payload, Peek: "test"}
		names = append(names, DataProduct)
	}
	return names
}

//QueryChatByChatID func
func QueryChatByChatID(inchatid string, user string) []strcode.ChatProduct {
	names := make([]strcode.ChatProduct, 0)
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("Select chat_id,user1,user2,payload,last_update from CHAT_DATA where chat_id = '%s'", inchatid)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return names
	}
	defer rows.Close()
	num := 0
	for rows.Next() {
		var chatid string
		var user1 string
		var user2 string
		var payload string
		var lastupdate string
		var puser string
		var suser string
		if err := rows.Scan(&chatid, &user1, &user2, &payload, &lastupdate); err != nil {
			log.Fatal(err)
		}
		if user1 == user {
			puser = user1
			suser = user2
		} else {
			puser = user2
			suser = user1
		}
		date1 := utilitycode.SubsBefore(lastupdate, "T")
		date2 := utilitycode.SubsAfer(utilitycode.SubsBefore(lastupdate, "+"), "T")
		datefix := date1 + " " + date2
		num = num + 1
		DataProduct := strcode.ChatProduct{No: num, ChatID: chatid, User1: puser, User2: suser, LastUpdate: datefix, Payload: payload, Peek: "test"}
		names = append(names, DataProduct)
	}
	return names
}

//QueryChatExst func
func QueryChatExst(user1 string, user2 string) string {
	var chatid string
	db := OpenCon()
	defer db.Close()
	statementSQL := fmt.Sprintf("select chat_id from CHAT_DATA where (user1 = '%s' and user2 = '%s') or (user1 = '%s' and user2 = '%s')", user1, user2, user2, user1)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return ""
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&chatid); err != nil {
			log.Fatal(err)
		}
	}
	return chatid
}
