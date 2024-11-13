package dbv1_0_0_testing

import (
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"

	stdlog "log"
	"os"
	"testing"
)

// upHandler handler
var (
	launcherManager setuputil.LauncherManager
	upHandler       *Migrate
)

func TestMain(m *testing.M) {
	// 初始化
	configPath := "./../../../../configs"
	launcher, err := setuputil.NewLauncherManager(configPath)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	//defer func() { _ = launcher.Close() }()

	// 数据库链接
	//dbConn, err := setuputil.GetMysqlDBConn(launcher)
	dbConn, err := setuputil.GetPostgresDBConn(launcher)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// migrateHandler 迁移手柄
	migrateRepo := migrationpkg.NewMigrateRepo(dbConn)

	// handler
	launcherManager = launcher
	upHandler = NewMigrateHandler(dbConn, migrateRepo)

	os.Exit(m.Run())
}
