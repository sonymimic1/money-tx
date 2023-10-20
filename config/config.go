package config

type (
	// Log -.
	LogSetting struct {
		Enable       bool   `json:"enable"`       // 是否啟用
		Level        string `json:"level"`        // log level
		FileSizeMega int    `json:"fileSizeMega"` // log rotate 的檔案大小 (MB)
		FileCount    int    `json:"fileCount"`    // log 檔的保留數量
		KeepDays     int    `json:"keepDays"`     // log 檔名日期的保留天數
		Path         string `json:"path"`         // log 路徑; 若為空字串, 則不輸出到檔案
	}

	// App -.
	AppSetting struct {
		GameServerHost      string   //	後端連線位置
		GameCodes           []string // 遊戲代碼列表
		WaitConnectInterval int      // 等待連線的間隔期間 (ms)
	}
)
