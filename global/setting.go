package global

import (
	"log"

	"github.com/sonymimic1/go-transfer/config"
	"github.com/sonymimic1/go-transfer/pkg/logger"
	"github.com/sonymimic1/go-transfer/pkg/util"
	"go.uber.org/zap"
)

// common.
var (
	LoggerSetting *config.LogSetting
	Logger        *zap.Logger
)

func init() {

	s, err := util.NewSetting("./")
	if err != nil {
		log.Fatal(err)
	}

	err = s.ReadSection("Log", &LoggerSetting)
	if err != nil {
		log.Fatal(err)
	}

	Logger = logger.NewLogger(&logger.Config{
		Enable:     LoggerSetting.Enable,
		Filename:   "./log/money-tx.log",
		MaxSize:    LoggerSetting.FileSizeMega,
		MaxAge:     LoggerSetting.KeepDays,
		MaxBackups: LoggerSetting.FileCount,
		Level:      LoggerSetting.Level,
	}, nil)
}
