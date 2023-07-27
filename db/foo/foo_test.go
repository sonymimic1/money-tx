package foo

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

const (
	MYSQL_USER         = "root"
	MYSQL_PASSWORD     = "123456"
	MYSQL_ADDRESS      = "127.0.0.1"
	MYSQL_ADDRESS_PORT = "3306"
	MYSQL_DB           = "transferDB"
)

func TestFooBasic(t *testing.T) {
	expect := 22
	actual := fooBasic(1, 1)
	assert.Equal(t, expect, actual)
}

func TestMysqlConnection(t *testing.T) {

	// 設定連線資訊
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_ADDRESS, MYSQL_ADDRESS_PORT, MYSQL_DB)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		t.Fatalf("無法連線至MySQL： %v", err)
	}
	defer db.Close()

	// 測試連線
	err = db.Ping()
	if err != nil {
		t.Fatalf("無法與MySQL建立連線： %v", err)
	}

	fmt.Println("MySQ連線正常")
	t.Log("MySQL連線正常")
}
