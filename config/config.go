package config

import "time"

type (
	// LogSetting -.
	LogSetting struct {
		Enable       bool   `mapstructure:"LOG_ENABLE"`         // 是否啟用
		Level        string `mapstructure:"LOG_LEVEL"`          // log level
		FileSizeMega int    `mapstructure:"LOG_FILE_SIZE_MEGA"` // log rotate 的檔案大小 (MB)
		FileCount    int    `mapstructure:"LOG_FILE_COUNT"`     // log 檔的保留數量
		KeepDays     int    `mapstructure:"LOG_KEEP_DAYS"`      // log 檔名日期的保留天數
		Path         string `mapstructure:"LOG_PATH"`           // log 路徑; 若為空字串, 則不輸出到檔案
	}

	AppSetting struct {
		DBDriver      string `mapstructure:"APP_DB_DRIVER"`
		DBSource      string `mapstructure:"APP_DB_SOURCE"`
		ServerAddress string `mapstructure:"APP_SERVER_ADDRESS"`
	}
	TokenSetting struct {
		TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
		AccessTokenDuration time.Duration `mapstructure:"TOKEN_ACCESS_TOKEN_DURATION"`
	}
)

//mapstructure:"DB_DRIVER"
//mapstructure:"DB_Source"
//mapstructure:"SERVER_ADDRESS"
