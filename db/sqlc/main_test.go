package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var TConn *sql.DB
var TestQuri *Queries

const (
	dbDriver = "mysql"

	//如果配置了parseTime=true，MySQL中的DATE、DATETIME等时间类型字段将自动转换为golang中的time.Time类型。 类似的0000-00-00 00:00:00 ，会被转为time.Time的零值。
	dbSource = "root:123456@tcp(127.0.0.1:3306)/transferDB?parseTime=true"
)

func TestMain(m *testing.M) {
	setup()
	Conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	TConn = Conn
	TestQuri = New(TConn)
	fmt.Printf("\033[1;33m%s\033[0m", "> Setup completed\n")

	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Do something here.

}

func teardown() {
	// Do something here.
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
