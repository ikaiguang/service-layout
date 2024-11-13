package main

import (
	"flag"
	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"github.com/ikaiguang/service-layout/app/testing-service/cmd/database-migration/migrate"
	stdlog "log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

// go run ./cmd/store-configuration/... -conf=./configs
func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	launcher, err := setuputil.NewLauncherManager(flagconf)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// 数据库链接
	//dbConn, err := setuputil.GetMysqlDBConn(launcherManager)
	//dbConn, err := setuputil.GetPostgresDBConn(launcher)
	dbConn, err := setuputil.GetRecommendDBConn(launcher)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	dbmigrate.Run(dbConn, dbutil.WithClose())
}
