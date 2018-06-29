package mysqlProc

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "debr979"
	password = "xxxxxxx"
	host     = "127.0.0.1"
	dbName   = "twStock"
)

func UserAction(param1 string, param2 map[string]string, actionID int) string {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", userName, password, host, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err.Error()
	}
	defer db.Close()
	switch actionID {
	case 0: //Register

		sqlCmd, err := db.Prepare("INSERT UserInf VALUE (?,?)")
		if err != nil {
			return err.Error()
		}
		defer sqlCmd.Close()
		pwd, _ := param2["PWD"]
		res, err := sqlCmd.Exec(param1, pwd)
		if err != nil {

			return err.Error()
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err.Error()
		}
		fmt.Println(id)

		return "REGISTER_SUCCESS"
	case 1: //Login
		pwd, _ := param2["PWD"]
		rows, err := db.Query("SELECT * FROM UserInf WHERE UID=? AND password=?", param1, pwd)
		if err != nil {
			return err.Error()
		}
		defer rows.Close()
		for rows.Next() {
			var UID string
			var password string
			if err := rows.Scan(&UID, &password); err != nil {
				return err.Error()
			}
			fmt.Printf("name:%s,id:is %s\n", UID, password)
			return UID + password
		}
		return ""

	case 2: //Delete
		sqlCmd, err := db.Prepare("DELETE FROM UserInf WHERE UID=?")
		if err != nil {
			return err.Error()
		}
		uid, _ := param2["UID"]
		res, err := sqlCmd.Exec(uid)
		if err != nil {
			return err.Error()
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err.Error()
		}
		fmt.Println(id)
		defer sqlCmd.Close()
		return "DELETE_SUCCESS"
	case 3:
		sqlCmd, err := db.Prepare("UPDATE UserInf SET password=? WHERE UID=? AND password=?")
		if err != nil {
			return err.Error()
		}
		pwd1, _ := param2["PWD"]
		pwd2, _ := param2["TOCHANGE"]
		res, err := sqlCmd.Exec(pwd2, param1, pwd1)
		if err != nil {
			return err.Error()
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err.Error()
		}
		fmt.Println(id)
		return "UPDATE_SUCCESS"
	}

	return ""
}

func DbControl(data map[string]interface{}) string {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", userName, password, host, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err.Error()
	}
	defer db.Close()
	actID, _ := data["actID"].(int)

	switch actID {
	case 0:
		UID := data["account"].(string)
		rows, err := db.Query("SELECT password FROM UserInf WHERE UID=?", UID)
		if err != nil {
			return err.Error()
		}
		defer rows.Close()
		for rows.Next() {
			var passwordx string
			if err := rows.Scan(&passwordx); err != nil {
				return err.Error()
			}
			return passwordx
		}
	case 1:
		break

	}
	return "NO_DATA"
}
