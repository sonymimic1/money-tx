package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonymimic1/go-transfer/api"
	db "github.com/sonymimic1/go-transfer/db/sqlc"
	"github.com/sonymimic1/go-transfer/global"
	"go.uber.org/zap"
)

// :"如果配置了parseTime=true，MySQL中的DATE、DATETIME等时间类型字段将自动转换为golang中的time.Time类型。 类似的0000-00-00 00:00:00 ，会被转为time.Time的零值。",
func main() {

	logPrefix := "main.main()"
	global.Logger.Info(logPrefix + ": start !")
	conn, err := sql.Open(global.AppSetting.DBDriver, global.AppSetting.DBSource)
	if err != nil {
		global.Logger.Fatal(logPrefix+": cannot connect to db", zap.Error(err))
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(global.AppSetting.ServerAddress)
	if err != nil {
		global.Logger.Fatal(logPrefix+": cannot start server", zap.Error(err))
	}

}
