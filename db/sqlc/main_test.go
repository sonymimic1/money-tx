package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	MYSQL_USER         = "root1"
// 	MYSQL_PASSWORD     = "123456"
// 	MYSQL_ADDRESS      = "127.0.0.11"
// 	MYSQL_ADDRESS_PORT = "3306"
// 	MYSQL_DB           = "transferDB"
// )
const (
	DBDriver = "mysql"
	DBSource = "mysql://" + MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_ADDRESS + ":" + MYSQL_ADDRESS_PORT + ")/" + MYSQL_DB + ""
)

var testQueries *Queries

func TestMain(m *testing.M) {

	// 設定連線資訊
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_ADDRESS, MYSQL_ADDRESS_PORT, MYSQL_DB)
	fmt.Printf("dsn=>%s \n", dsn)
	db, err := sql.Open("mysql", dsn)
	fmt.Println("MySQL連線中")
	if err != nil {
		panic(err.Error())
	}
	// 測試連線
	err = db.Ping()
	if err != nil {
		fmt.Printf("無法與MySQL建立連線： %v \n", err)
		panic(err.Error())
	}

	fmt.Println("MySQL連線ok")
	testQueries = New(db)

	os.Exit(m.Run())
}
