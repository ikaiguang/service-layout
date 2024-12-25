//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/biz/biz"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/data/data"
	"github.com/ikaiguang/service-layout/app/testing-service/internal/service/service"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		//setuputil.GetRecommendDBConn,
		// data
		data.NewTestingData,
		// biz
		biz.NewTestingBiz,
		// service
		service.NewTestingV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
