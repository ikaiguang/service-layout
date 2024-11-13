//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	servicev1 "github.com/ikaiguang/service-layout/api/testing-service/v1/services"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/biz/biz"
	bizrepos "github.com/ikaiguang/service-layout/app/testing-service/internal/biz/repo"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/data/data"
	datarepos "github.com/ikaiguang/service-layout/app/testing-service/internal/data/repo"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/service/service"
)

func exportTestingData(launcherManager setuputil.LauncherManager) (datarepos.TestingDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		//setuputil.GetRecommendDBConn,
		data.NewTestingData,
	))
	return nil, nil
}

func exportTestingBiz(launcherManager setuputil.LauncherManager) (bizrepos.TestingBizRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		exportTestingData,
		// biz
		biz.NewTestingBiz,
	))
	return nil, nil
}

func exportTestdataServer(launcherManager setuputil.LauncherManager) (servicev1.SrvTestdataServer, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// biz
		exportTestingBiz,
		// service
		service.NewTestingV1Service,
	))
	return nil, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		// service
		exportTestdataServer,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
