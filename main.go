package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonymimic1/go-transfer/api"
	db "github.com/sonymimic1/go-transfer/db/sqlc"
	"github.com/sonymimic1/go-transfer/global"
	"go.uber.org/zap"
)

const (
	dbDriver = "mysql"

	//如果配置了parseTime=true，MySQL中的DATE、DATETIME等时间类型字段将自动转换为golang中的time.Time类型。 类似的0000-00-00 00:00:00 ，会被转为time.Time的零值。
	dbSource      = "root:123456@tcp(127.0.0.1:3306)/transferDB?parseTime=true"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	logPrefix := "main.main()"

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		global.Logger.Fatal(logPrefix+": cannot connect to db", zap.Error(err))
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		global.Logger.Fatal(logPrefix+": cannot start server", zap.Error(err))
	}

}
