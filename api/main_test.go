package api

import (
	"os"
	"testing"
	"time"

	"github.com/sonymimic1/go-transfer/config"
	db "github.com/sonymimic1/go-transfer/db/sqlc"
	"github.com/sonymimic1/go-transfer/global"

	"github.com/sonymimic1/go-transfer/pkg/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	global.TokenSetting = &config.TokenSetting{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(store)
	require.NoError(t, err)

	return server

}
func TestMain(m *testing.M) {
	//gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
