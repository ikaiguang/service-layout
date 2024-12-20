package serviceexporter

import (
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	testingapi "github.com/ikaiguang/service-layout/api/testing-service"
	servicev1 "github.com/ikaiguang/service-layout/api/testing-service/v1/services"
	"github.com/ikaiguang/service-layout/app/testing-service/cmd/database-migration/migrate"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/conf"
)

func ExportServiceConfig() []configutil.Option {
	return conf.LoadServiceConfig()
}

func ExportAuthWhitelist() []map[string]middlewareutil.TransportServiceKind {
	return []map[string]middlewareutil.TransportServiceKind{
		testingapi.GetAuthWhiteList(),
	}
}

func ExportServices(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
	hs, err := serverManager.GetHTTPServer()
	if err != nil {
		return nil, err
	}
	gs, err := serverManager.GetGRPCServer()
	if err != nil {
		return nil, err
	}
	return exportServices(launcherManager, hs, gs)
	//return serverutil.MergeCleanup(exportServices(launcherManager, hs, gs))
}

func ExportServicesWithDatabaseMigration(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
	settingConfig := launcherManager.GetConfig().GetSetting()

	if settingConfig.GetEnableMigrateDb() {
		dbConn, err := setuputil.GetRecommendDBConn(launcherManager)
		if err != nil {
			return nil, err
		}
		logger, err := setuputil.GetLogger(launcherManager)
		if err != nil {
			return nil, err
		}
		dbmigrate.Run(dbConn, dbutil.WithLogger(logger))
	}
	return ExportServices(launcherManager, serverManager)
}

func ExportDatabaseMigration() []dbutil.MigrationFunc {
	return []dbutil.MigrationFunc{
		dbmigrate.Run,
	}
}

func ExportNodeIDV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvTestdataServer, error) {
	return exportTestdataServer(launcherManager)
}
