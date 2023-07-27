package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_ADDRESS, MYSQL_ADDRESS_PORT, MYSQL_DB)
	// conn, err := sql.Open(DBDriver, dsn)

	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// } else {
	// 	fmt.Printf("OK??\n")
	// 	fmt.Printf("conn:%v\n", conn)
	// }

	// db, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// testQueries = New(db)

	os.Exit(m.Run())
}
