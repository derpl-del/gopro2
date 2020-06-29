package dbcode

import (
	"database/sql"
	"fmt"
	"log"

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
