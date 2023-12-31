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
	AppSetting    *config.AppSetting
	TokenSetting  *config.TokenSetting
	Logger        *zap.Logger
)

func InitConfig() {
	logPrefix := "global.init()"

	s, err := util.NewSetting(".", "app", "env")
	//err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	//err = s.ReadSection("Log", &LoggerSetting)
	err = s.Vp.Unmarshal(&LoggerSetting)
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

	//err = s.ReadSection("App", &AppSetting)
	err = s.Vp.Unmarshal(&AppSetting)
	if err != nil {
		Logger.Fatal(logPrefix+": load AppSetting fail", zap.Error(err))
	}

	//err = s.ReadSection("Token", &TokenSetting)
	err = s.Vp.Unmarshal(&TokenSetting)
	if err != nil {
		Logger.Fatal(logPrefix+": load AppSetting fail", zap.Error(err))
	}

}
