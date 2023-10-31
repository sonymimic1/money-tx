package config

type (
	//Log -.
	LogSetting struct {
		Enable       bool   `json:"enable"`       // 是否啟用
		Level        string `json:"level"`        // log level
		FileSizeMega int    `json:"fileSizeMega"` // log rotate 的檔案大小 (MB)
		FileCount    int    `json:"fileCount"`    // log 檔的保留數量
		KeepDays     int    `json:"keepDays"`     // log 檔名日期的保留天數
		Path         string `json:"path"`         // log 路徑; 若為空字串, 則不輸出到檔案
	}

	AppSetting struct {
		DBDriver      string `json:"dbDriver"`
		DBSource      string `json:"dbSource"`
		ServerAddress string `json:"serverAddress"`
	}
	// LogSetting struct {
	// 	Enable       bool   `mapstructure:"ENABLE"`       // 是否啟用
	// 	Level        string `mapstructure:"LEVEL"`        // log level
	// 	FileSizeMega int    `mapstructure:"FILESIZEMEGA"` // log rotate 的檔案大小 (MB)
	// 	FileCount    int    `mapstructure:"FILECOUNT"`    // log 檔的保留數量
	// 	KeepDays     int    `mapstructure:"KEEPDAYS"`     // log 檔名日期的保留天數
	// 	Path         string `mapstructure:"PATH"`         // log 路徑; 若為空字串, 則不輸出到檔案
	// }

	// AppSetting struct {
	// 	DBDriver      string `mapstructure:"DB_DRIVER"`
	// 	DBSource      string `mapstructure:"DB_SOURCE"`
	// 	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	// }
)

//mapstructure:"DB_DRIVER"
//mapstructure:"DB_Source"
//mapstructure:"SERVER_ADDRESS"
